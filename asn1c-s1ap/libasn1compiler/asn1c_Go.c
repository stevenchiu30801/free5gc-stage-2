#include "asn1c_internal.h"
#include "asn1c_Go.h"
#include "asn1c_constraint.h"
#include "asn1c_out.h"
#include "asn1c_misc.h"
#include "asn1c_ioc.h"
#include "asn1c_naming.h"
#include <asn1print.h>
#include <asn1fix_crange.h>	
#include <asn1fix_export.h>	
#include <asn1parser.h>

enum onc_flags {
	ONC_noflags		= 0x00,
	ONC_avoid_keywords	= 0x01,
	ONC_force_compound_name	= 0x02,
};

static abuf origin_struct_name;

const char *get_ref_field_val_from_IEs(const char *id_name);
const char *get_ref_field_val_from_Procedures(const char *id_name);
int asn1c_lang_Go_emit_constraint_checking_code(arg_t *arg, asn1p_expr_t *target);
char *emit_single_member_PER_constraint(arg_t *arg, asn1cnst_range_t *range, int alphabetsize, const char *type);
asn1p_expr_t *get_topmost_expr(arg_t *arg, asn1p_expr_t *expr);
static int expr_elements_count(arg_t *arg, asn1p_expr_t *expr);
static int expr_extensible_conut(arg_t *arg, asn1p_expr_t *expr);
static asn1p_expr_type_e _find_terminal_type(arg_t *arg);
static asn1p_expr_type_e expr_get_type(arg_t *arg, asn1p_expr_t *expr);
static void abuf_oint(abuf *ab, asn1c_integer_t v);
static abuf *emit_range_comparison_code(asn1cnst_range_t *range, const char *varname,
                           asn1c_integer_t natural_start,
                           asn1c_integer_t natural_stop);

static int is_open_type(arg_t *arg, asn1p_expr_t *expr, asn1c_ioc_table_and_objset_t *opt_ioc);
static ssize_t find_column_index(arg_t *arg, asn1c_ioc_table_and_objset_t *opt_ioc, const char *column_name);

static int asn1c_lang_Go_OpenType(arg_t *arg, asn1c_ioc_table_and_objset_t *opt_ioc, const char *column_name);
static int asn1c_lang_Go_OpenType_extra_ref(arg_t *arg, asn1c_ioc_table_and_objset_t *opt_ioc, const char *column_name);

const char *asn1c_get_ioc_value(arg_t *arg, struct asn1p_ioc_cell_s *cell);
char *asn1c_get_ioc_cell(arg_t *arg, struct asn1p_ioc_cell_s *cell);

const char *asn1c_get_member_type_selector(arg_t *arg, asn1p_expr_t *expr);

enum tvm_compat {
	_TVM_SAME	= 0,	/* tags and all_tags are same */
	_TVM_SUBSET	= 1,	/* tags are subset of all_tags */
	_TVM_DIFFERENT	= 2,	/* tags and all_tags are different */
};

enum etd_spec {
	ETD_NO_SPECIFICS,
	ETD_HAS_SPECIFICS
};


#define	MKID(expr)	(asn1c_make_identifier(AMI_USE_PREFIX, expr, 0))
#define is_spec_index(expr) (expr->spec_index != -1 && expr->_lineno)


int asn1c_lang_Go_type_REFERENCE(arg_t *arg) {
	printf("[Debug] asn1c_lang_Go_type = REFERENCE, %s\n", c_name(arg).type.asn_name);
	asn1p_ref_t *ref;
	asn1p_expr_t *expr = arg->expr;
	asn1p_expr_t *v;
	int saved_target = arg->target->target;
	static abuf ab;

	abuf_clear(&ab);
    
	ref = arg->expr->reference;
	if(ref->components[ref->comp_count-1].name[0] == '&') {
//		printf("[Debug] The first char is &\n");
		asn1p_expr_t *extract;
		arg_t tmp;
		int ret;

        extract = WITH_MODULE_NAMESPACE(
            arg->expr->module, expr_ns,
            asn1f_class_access_ex(arg->asn, arg->expr->module, expr_ns,
                                  arg->expr, arg->expr->rhs_pspecs, ref));
        if(extract == NULL)
			return -1;

		extract = asn1p_expr_clone(extract, 0);
		if(extract) {
			free(extract->Identifier);
			extract->Identifier = strdup(arg->expr->Identifier);
			if(extract->Identifier == NULL) {
				asn1p_expr_free(extract);
				return -1;
			}
		} else {
			return -1;
		}

		tmp = *arg;
		tmp.asn = arg->asn;
		tmp.expr = extract;

		ret = arg->default_cb(&tmp, NULL);

		asn1p_expr_free(extract);

		return ret;
	}

	REDIR(OT_CODE);
	OUT("// Open type declare \n");
	struct c_names struct_name = c_name(arg);

    if(arg->embed) {
    	REDIR(OT_CODE);
        abuf_printf(&ab, "%s\t", c_name(arg).title_name);
        if(arg->expr->parent_expr->expr_type == ASN_CONSTR_CHOICE ||
            arg->expr->parent_expr->expr_type == ASN_CONSTR_OPEN_TYPE) abuf_str(&ab, "*");
            if(strcmp(c_name(arg).title_name, c_name(arg).type.go_name))
                abuf_str(&ab, origin_struct_name.buffer);
            abuf_str(&ab, c_name(arg).type.go_name);
		if(expr->marker.flags & EM_OPTIONAL) 
			OUT("*");
        OUT("%s\n", ab.buffer);
        REDIR(saved_target);
        return 0;
    }

    OUT("type %s%s struct {\n", struct_name.title_name, struct_name.type.go_name);
    INDENTED(
	    TQ_FOR(v, &(expr->members), next) {
		    if(v->expr_type == A1TC_EXTENSIBLE) continue;
   			arg_t tmp_arg = *arg;
    		tmp_arg.embed ++;
	    	tmp_arg.expr = v;
   			asn1c_lang_Go_type_SIMPLE_TYPE(&tmp_arg);
    		tmp_arg.embed --;
	    }
	);	
	OUT("}\n");
	OUT("\n");
	REDIR(saved_target);

	return 0;
}

int asn1c_lang_Go_type_EXTENSIBLE(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = EXTENSIBLE\n");
    return asn1c_lang_Go_type_SIMPLE_TYPE(arg);
}

static int
is_open_type(arg_t *arg, asn1p_expr_t *expr, asn1c_ioc_table_and_objset_t *opt_ioc) {

    (void)arg;

    if(!opt_ioc) {
        return 0;
    }

    if(expr->meta_type == AMT_TYPEREF
       && expr->expr_type == A1TC_REFERENCE
       && expr->reference->comp_count == 2
       && expr->reference->components[1].lex_type
              == RLT_AmpUppercase) {
        DEBUG("%s is a true open type", MKID(expr));
        return 1;
    }

    return 0;
}

int asn1c_lang_Go_type_SEQUENCE(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = SEQUENCE\n");

	asn1p_expr_t *expr = arg->expr;
	asn1p_expr_t *v;
	int saved_target = arg->target->target;
    asn1c_ioc_table_and_objset_t ioc_tao;


	ioc_tao = asn1c_get_ioc_table(arg);

	REDIR(OT_CODE);
	if(arg->embed) {
        OUT("/* Sequence Embed but no implement */\n");
        return 0;
	}
	else {
//		printf("[Debug] asn1c_lang_Go_type_SEQUENCE = not embed\n");
		OUT("type %s struct {\n", c_name(arg).title_name);
        abuf_clear(&origin_struct_name);
        abuf_str(&origin_struct_name, c_name(arg).title_name);
	}
    INDENTED(
	TQ_FOR(v, &(expr->members), next) {
		if(v->expr_type == A1TC_EXTENSIBLE) continue;
        arg_t tmp_arg = *arg;
        tmp_arg.embed++;
        tmp_arg.expr = v;
        if(is_open_type(arg, v, ioc_tao.ioct ? &ioc_tao : 0)) {
            const char *column_name = v->reference->components[1].name;
            if(asn1c_lang_Go_OpenType(&tmp_arg, &ioc_tao, column_name)) {
                return -1;
            }
        } else {
            if(tmp_arg.expr->expr_type != A1TC_REFERENCE) {
                tmp_arg.default_cb(&tmp_arg, (ioc_tao.ioct ? &ioc_tao : 0));
            }
            else
    			asn1c_lang_Go_type_SIMPLE_TYPE(&tmp_arg);
		}
		tmp_arg.embed--;
	}
    );
	OUT("}\n");
    OUT("\n");

    TQ_FOR(v, &(expr->members), next) {
        arg_t tmp_arg = *arg;
        tmp_arg.embed++;
        tmp_arg.expr = v;
        if(is_open_type(arg, v, ioc_tao.ioct ? &ioc_tao : 0)) {
             const char *column_name = v->reference->components[1].name;
            if(asn1c_lang_Go_OpenType_extra_ref(&tmp_arg, &ioc_tao, column_name)) {
                return -1;
            } 
        }
        tmp_arg.embed--;
    }

	REDIR(saved_target);
    return 0;
}

int asn1c_lang_Go_type_SET(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = SET\n");
    return asn1c_lang_Go_type_SIMPLE_TYPE(arg);
}

int asn1c_lang_Go_type_SEx_OF(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = SEx_OF\n");

	asn1p_expr_t *expr = arg->expr;
	asn1p_expr_t *memb = TQ_FIRST(&expr->members);
	int saved_target = arg->target->target;

	if(arg->expr->expr_type == ASN_CONSTR_SET_OF) {
		printf("[Debug] Set OF does not implement yet\n");
		return 0;
	}

	REDIR(OT_CODE);
	OUT("/* Sequence of = %d, FULL Name = %s */\n", expr->expr_type,  c_name(arg).full_name);
    OUT("/* %s */\n", asn1go_type_name(arg, memb, TNF_RSAFE));

	if(arg->embed) {
        OUT("/* Sequence Of Embed but not implement */\n");
        return 0;
	}
	else {
        OUT("type ");
        if(is_spec_index(expr)) {  /* reference to same struct but need to use different name */ 
            OUT("%s", c_name(arg).title_name);
            OUT("%s", asn1go_type_name(arg, memb, TNF_RSAFE));
        }
        else
    		OUT("%s", c_name(arg).title_name);
        OUT(" struct {\n");
	}
	INDENTED(
        asn1p_expr_t *v;
		TQ_FOR(v, &(expr->members), next) {
			if(v->expr_type == A1TC_EXTENSIBLE) {
				OUT("/* Extensions may appear below */\n");
				continue;
			}
			OUT("List []%s ", asn1go_type_name(arg, memb, TNF_RSAFE));
            asn1c_lang_Go_emit_constraint_checking_code(arg, expr);
            OUT("\n");
		}
	);
	OUT("}\n");
	REDIR(saved_target);
    return 0;
}

int asn1c_lang_Go_type_CHOICE(arg_t *arg) {
//	printf("[Debug] asn1c_lang_Go_type = CHOICE\n");
    asn1p_expr_t *expr = arg->expr;
	asn1p_expr_t *v;
	int saved_target = arg->target->target;

    REDIR(OT_CODE);

//	OUT("/* Choice Type */\n");

    OUT("const (\n");
    INDENTED(
        if(arg->embed)
            OUT("%sPresentNothing\tint = iota\t/* No components present */\n", origin_struct_name.buffer);
        else 
            OUT("%s\tint = iota\t/* No components present */\n", go_presence_name(arg, NULL));
		TQ_FOR(v, &(expr->members), next) {
            int skipComma = 1;
        	if(skipComma) skipComma = 0;
	   		else if (v->expr_type == A1TC_EXTENSIBLE && !TQ_NEXT(v, next)) OUT("\n");
	    	else OUT(",\n");
		    if(v->expr_type == A1TC_EXTENSIBLE) {
			    OUT("/* Extensions may appear below */\n");
	    	    skipComma = 1;
			    continue;
		    }
            if(arg->embed) 
                OUT("%sPresent%s\n", origin_struct_name.buffer, asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, v, 0));
            else
    		    OUT("%s\n", go_presence_name(arg, v));
	    }
    );
	OUT(")\n");
    OUT("\n");

	asn1c_ioc_table_and_objset_t ioc_tao;
	ioc_tao = asn1c_get_ioc_table(arg);
	

    if(arg->embed) {
        arg_t tmp_arg = *arg;
        tmp_arg.expr = expr->parent_expr;
        OUT("type %s", c_expr_name(&tmp_arg, tmp_arg.expr).title_name);
        OUT("%s struct {\n", c_name(arg).title_name);
    }
    else {
        OUT("type %s struct {\n", c_name(arg).title_name);
    }
    abuf_clear(&origin_struct_name);
    abuf_str(&origin_struct_name, c_name(arg).title_name);

	INDENTED(
		OUT("Present\tint\n");
		TQ_FOR(v, &(expr->members), next) {
            if(v->expr_type == A1TC_EXTENSIBLE) continue;
			arg_t tmp_arg = *arg;
			tmp_arg.embed ++;
			tmp_arg.expr = v;
			if(is_open_type(arg, v, ioc_tao.ioct ? &ioc_tao : 0)) {
				OUT("/* Choice member type is Open Type */\n");
				const char *column_name = v->reference->components[1].name;
            	if(asn1c_lang_Go_OpenType(&tmp_arg, &ioc_tao, column_name)) {
                	return -1;
				}
            } else {
                if(tmp_arg.expr->expr_type != A1TC_REFERENCE)
                    tmp_arg.default_cb(&tmp_arg, (ioc_tao.ioct ? &ioc_tao : 0));
                else
    		        asn1c_lang_Go_type_SIMPLE_TYPE(&tmp_arg);
			}
			tmp_arg.embed --;
		}
	);
	OUT("}\n");
    OUT("\n");
	REDIR(saved_target);
    return 0;
}
 
int asn1c_lang_Go_type_common_INTEGER(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = common_INTEGER\n");
    asn1p_expr_t *expr = arg->expr;
   	asn1p_expr_t *v;
	int saved_target = arg->target->target; 

	if(expr->expr_type == ASN_BASIC_ENUMERATED && arg->embed == 0) {
//		printf("ASN_BASIC_ENUMERATED\n");
        REDIR(OT_CODE);
        OUT("const (\n");
        INDENT(+1);
        TQ_FOR(v, &(expr->members), next) {
            switch(v->expr_type) {
                case A1TC_UNIVERVAL:
                    OUT("%s\taper.Enumerated = %s\n", go_presence_name(arg, v), asn1p_itoa(v->value->value.v_integer));
                    break;
                case A1TC_EXTENSIBLE:
//                    printf("[Debug] ENUMERATED is extensible, but not implement\n");
                    break;
                default:
                    printf("[Debug] ENUMERATED = Unimplement type\n");
                    return -1;
            }
        }
        INDENT(-1);
        OUT(")\n");
        OUT("\n");
	    REDIR(saved_target);
	}
	
	if(expr->expr_type == ASN_BASIC_INTEGER 
	&& asn1c_type_fits_long(arg, expr) == FL_FITS_UNSIGN) {
		printf("ASN_BASIC_INTEGER && FL_FITS_UNSIGN\n");
	}

    return asn1c_lang_Go_type_SIMPLE_TYPE(arg);
}

int asn1c_lang_Go_type_BIT_STRING(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = BIT_STRING\n");
    return asn1c_lang_Go_type_SIMPLE_TYPE(arg);
}

int asn1c_lang_Go_type_REAL(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = REAL\n");
    return asn1c_lang_Go_type_SIMPLE_TYPE(arg);
}

int asn1c_lang_Go_type_SIMPLE_TYPE(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = SIMPLE_TYPE\n");
	asn1p_expr_t *expr = arg->expr;
	int saved_target = arg->target->target;	
	static abuf ab;

	abuf_clear(&ab);

	if(arg->embed) {
//		printf("[Debug] asn1c_lang_Go_type_SIMPLE_TYPE = embed\n");
        struct c_names obj = c_expr_name(arg, expr);
		abuf_printf(&ab, "%s\t", obj.title_name);	
		if(arg->expr->parent_expr->expr_type == ASN_CONSTR_CHOICE ||
		   arg->expr->parent_expr->expr_type == ASN_CONSTR_OPEN_TYPE ||
		   expr->marker.flags & EM_OPTIONAL) abuf_str(&ab, "*");

		
        /* Like Sequence Of */
        arg_t tmp = *arg;
        asn1p_expr_t *ref_type = NULL;
        if(expr->reference) {
            ref_type = WITH_MODULE_NAMESPACE(
                    tmp.expr->module, expr_ns,
                    asn1f_lookup_symbol_ex(tmp.asn, expr_ns, tmp.expr,
                    arg->expr->reference));
        }
        if(ref_type && is_spec_index(ref_type)) {
            tmp.expr = ref_type;
            abuf_str(&ab, c_name(&tmp).title_name);
            if(ref_type->expr_type == ASN_CONSTR_SEQUENCE_OF) {
                abuf_str(&ab, asn1go_type_name(&tmp, TQ_FIRST(&ref_type->members), TNF_RSAFE));
            }
            else {
				asn1p_expr_t *terminal = asn1f_find_terminal_type_ex(arg->asn, arg->ns, expr);
				abuf_str(&ab, c_expr_name(arg, terminal).title_name);
            }
        } else {
		    abuf_str(&ab, obj.type.go_name);
        }

		OUT("%s ", ab.buffer);
		asn1c_lang_Go_emit_constraint_checking_code(arg, expr);
		OUT("\n");
	}
	else {
//		printf("[Debug] asn1c_lang_Go_type_SIMPLE_TYPE = not embed\n");
        REDIR(OT_CODE);
        OUT("type %s struct {\n", c_name(arg).title_name);
		INDENTED(
        	OUT("Value\t");
			if(expr->marker.flags & EM_OPTIONAL)
				OUT("*");
			OUT("%s ", asn1go_type_name(arg, expr, TNF_RSAFE));
			asn1c_lang_Go_emit_constraint_checking_code(arg, expr);
			OUT("\n");
		);
		OUT("}\n");
		OUT("\n");
	}

	REDIR(saved_target);
    return 0;
}

int asn1c_lang_Go_type_REFERENCE_Value(arg_t *arg) {
//    printf("[Debug] asn1c_lang_Go_type = REFERENCE_Value\n");
	arg_t tmp = *arg;
	asn1p_expr_t *expr, *ref_type;
	int saved_target;

	expr = arg->expr;
	ref_type = WITH_MODULE_NAMESPACE(
			tmp.expr->module, expr_ns,
			asn1f_lookup_symbol_ex(tmp.asn, expr_ns, tmp.expr,
			arg->expr->reference));
	if(!ref_type)
		return 0;

	if(!ref_type->data)
		asn1c_attach_streams(ref_type);

	arg->target = ref_type->data;
	saved_target = arg->target->target;
	REDIR(OT_CODE);

    if(arg->embed) {
        OUT("/* Reference Value Embed, but not implement */\n");
        return 0;
    }

	if((ref_type->expr_type == ASN_BASIC_INTEGER) ||
		(ref_type->expr_type == ASN_BASIC_ENUMERATED)) {
        struct c_names struct_name = c_name(arg);
        if(strncmp(struct_name.title_name, "Id", 2) == 0)
    		OUT("const %s%s ", struct_name.type.go_name, struct_name.title_name + 2);
        else
            OUT("const %s%s ", struct_name.type.go_name, struct_name.title_name);

        OUT("int64 = %s\n", asn1p_itoa(expr->value->value.v_integer));
	}

	REDIR(saved_target);
	arg->target = tmp.target;
	return 0;
}

static ssize_t
find_column_index(arg_t *arg, asn1c_ioc_table_and_objset_t *opt_ioc, const char *column_name) {
    (void)arg;

    if(!opt_ioc || !opt_ioc->ioct || !column_name) {
        return -1;
    }

    if(opt_ioc->ioct->rows == 0) {
        return 0;   /* No big deal. Just no data */
    } else {
        for(size_t clmn = 0; clmn < opt_ioc->ioct->row[0]->columns; clmn++) {
            if(strcmp(opt_ioc->ioct->row[0]->column[clmn].field->Identifier,
                      column_name) == 0) {
                return clmn;
            }
        }
        return -1;
    }

}

static int asn1c_lang_Go_OpenType(arg_t *arg, asn1c_ioc_table_and_objset_t *opt_ioc, 
								  const char *column_name) {
	printf("[Debug] asn1c_lang_Go_type = OpenType %s\n", column_name);
    arg_t tmp_arg = *arg;

    ssize_t column_index = find_column_index(arg, opt_ioc, column_name);
    if(column_index < 0) {
        FATAL("Open type generation attempted for %s, incomplete", column_name);
        return -1;
    }

    asn1p_expr_t *open_type_choice =
        asn1p_expr_new(arg->expr->_lineno, arg->expr->module);

    open_type_choice->Identifier = strdup(arg->expr->Identifier);
    open_type_choice->meta_type = AMT_TYPE;
    open_type_choice->expr_type = ASN_CONSTR_OPEN_TYPE;
    open_type_choice->_type_unique_index = arg->expr->_type_unique_index;
    open_type_choice->parent_expr = arg->expr->parent_expr;

    for(size_t row = 0; row < opt_ioc->ioct->rows; row++) {
        struct asn1p_ioc_cell_s *cell =
            &opt_ioc->ioct->row[row]->column[column_index];

        if(!cell->value) continue;

        asn1p_expr_t *m = asn1p_expr_clone(cell->value, 0);
        asn1p_expr_add(open_type_choice, m);
    }


    tmp_arg.expr = open_type_choice;
	struct c_names struct_name = c_name(&tmp_arg);
    const char *parent_name = asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, tmp_arg.expr->parent_expr, 0);
	static abuf new_name;
	abuf_clear(&new_name);
	abuf_printf(&new_name, "%s%s", parent_name, struct_name.title_name);

//	OUT("/* Open Type Selector = %s */\n", asn1c_get_member_type_selector(arg, arg->expr));

	OUT("%s\t", struct_name.title_name);
	if(tmp_arg.expr->marker.flags & EM_OPTIONAL) 
		OUT(".");
	OUT("%s ", new_name.buffer);
	asn1p_expr_type_e type_store = arg->expr->expr_type;
	arg->expr->expr_type = ASN_CONSTR_OPEN_TYPE;
    asn1c_lang_Go_emit_constraint_checking_code(arg, arg->expr);
	arg->expr->expr_type = type_store;
    OUT("\n");

    asn1p_expr_free(tmp_arg.expr);
    return 0;
}

static int asn1c_lang_Go_OpenType_extra_ref(arg_t *arg, asn1c_ioc_table_and_objset_t *opt_ioc, const char *column_name) {
	printf("[Debug] asn1c_lang_Go_type = OpenType Extra Ref %s\n", column_name);
    arg_t tmp_arg = *arg;
    static abuf ab;

    abuf_clear(&ab);

    ssize_t column_index = find_column_index(arg, opt_ioc, column_name);
    if(column_index < 0) {
        FATAL("Open type generation attempted for %s, incomplete", column_name);
        return -1;
    }

    asn1p_expr_t *open_type_choice =
        asn1p_expr_new(arg->expr->_lineno, arg->expr->module);

    open_type_choice->Identifier = strdup(arg->expr->Identifier);
    open_type_choice->meta_type = AMT_TYPE;
    open_type_choice->expr_type = ASN_CONSTR_OPEN_TYPE;
    open_type_choice->_type_unique_index = arg->expr->_type_unique_index;
    open_type_choice->parent_expr = arg->expr->parent_expr;

    for(size_t row = 0; row < opt_ioc->ioct->rows; row++) {
        struct asn1p_ioc_cell_s *cell =
            &opt_ioc->ioct->row[row]->column[column_index];

        if(!cell->value) continue;

        asn1p_expr_t *m = asn1p_expr_clone(cell->value, 0);
        asn1p_expr_add(open_type_choice, m);
    }

    tmp_arg.expr = open_type_choice;

    if(opt_ioc->ioct) {
//		OUT("/* ----- Open Type with ioc tao ----- */\n");
/*
		for(size_t rn = 0; rn < opt_ioc->ioct->rows; rn++) {
			asn1p_ioc_row_t *row = opt_ioc->ioct->row[rn];
			for(size_t cn = 0; cn < row->columns; cn++) {
				struct asn1p_ioc_cell_s *cell = &row->column[cn];
				asn1c_get_ioc_cell(arg, cell);
			}
		}
*/
//		OUT("/* ----- End of ioc tao ----- */\n");
		
		asn1p_expr_t *v;
		int member_cnt = 0;
        abuf_str(&ab, c_expr_name(&tmp_arg, tmp_arg.expr->parent_expr).title_name);
		TQ_FOR(v, &(tmp_arg.expr->members), next) {
            for(int procedureCodeExist = 0; opt_ioc->ioct->row[member_cnt] && !procedureCodeExist; member_cnt++) {
                asn1p_ioc_row_t *row = opt_ioc->ioct->row[member_cnt];
                procedureCodeExist = 1;
    			for(size_t cn = 0; cn < row->columns; cn++) {
	    			struct asn1p_ioc_cell_s *cell = &row->column[cn];
//                    asn1c_get_ioc_cell(arg, cell);

//                    OUT("/* What is my ID = %s */\n", cell->field->Identifier + 1); 
                    if(strcmp(cell->field->Identifier + 1, ab.buffer) == 0 && cell->value == NULL) {
                        procedureCodeExist = 0;
                    }
                    if(strcmp(cell->field->Identifier + 1, "procedureCode") == 0) {
                        if(procedureCodeExist) {
                            asn1p_expr_t *m = asn1p_expr_clone(cell->value, 0);
                            asn1p_expr_add(v, m);
//                            OUT("/* Procedure ID = %s, ", asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, m, 0));
//                            OUT("Type = %s */\n", asn1go_type_name(arg, m, TNF_RSAFE));
                        }
                        continue;
                    } else if(strcmp(cell->field->Identifier + 1, "id") == 0) {
                        asn1p_expr_t *m = asn1p_expr_clone(cell->value, 0); 
                        asn1p_expr_add(v, m);
//                      OUT("/* Id ID = %s, ", asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, m, 0));
//                      OUT("Type = %s */\n", asn1go_type_name(arg, m, TNF_RSAFE));
                    }

    				if(!(strcmp(cell->field->Identifier + 1, "id") == 0 || 
	    				 strcmp(cell->field->Identifier + 1, "Value") == 0 ||
		    			 strcmp(cell->field->Identifier + 1, c_name(arg).title_name) == 0))
			    		continue;
               
    				if(!cell->value) {
	    				continue;
    				}
	    			else if(cell->value->meta_type == AMT_VALUE){
		    			free(v->Identifier);
			    		v->Identifier = strdup(asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, cell->value, 0) + 2);
				    }
    				else if(cell->value->meta_type == AMT_TYPEREF || cell->value->meta_type == AMT_TYPE) {
	    				/* New Type = asn1go_type_name(arg, cell->value, TNF_RSAFE) */
		    		}
    			}	
			}
		}
		asn1c_lang_Go_type_CHOICE(&tmp_arg);
	}

    asn1p_expr_free(tmp_arg.expr);
    return 0;
}

const char *asn1c_get_member_type_selector(arg_t *arg, asn1p_expr_t *expr) {
	static abuf ab;	
	
	abuf_clear(&ab);

	const asn1p_constraint_t *crc =
        asn1p_get_component_relation_constraint(expr->combined_constraints);

	/* Not an Open Type, it seems. */
	if(!crc || crc->el_count <= 1) return 0;
	
	const asn1p_ref_t *cref = crc->elements[1]->value->value.reference;
	const char *cname = cref->components[0].name;
	if(cname[0] == '@' && cname[1] != '.') {
        cname += 1;
    } else if(cname[0] == '@' && cname[1] == '.' && cname[2] != '.') {
        cname += 2;
    } else {
        FATAL("Complex IoS reference %s can not be processed",
              asn1p_ref_string(cref));
        return NULL;
    }
	if(strlen(cname) == 0) return NULL;
	abuf_printf(&ab,"%c", (char)toupper(cname[0]));
	abuf_str(&ab, ++cname);
	return ab.buffer;
}

const char *
asn1c_get_ioc_value(arg_t *arg, struct asn1p_ioc_cell_s *cell) {
    static abuf ab;

    if(cell->value && cell->value->meta_type == AMT_VALUE) {
        int primitive_representation = 0;
        
        abuf_clear(&ab);

        asn1p_expr_t *cv_type =
            asn1f_find_terminal_type_ex(arg->asn, arg->ns, cell->value);

        switch(cv_type->expr_type) {
        case ASN_BASIC_INTEGER:
        case ASN_BASIC_ENUMERATED:
            switch(asn1c_type_fits_long(arg, cell->value /* sic */)) {
            case FL_NOTFIT:
                break;
            case FL_PRESUMED:
            case FL_FITS_SIGNED:
                primitive_representation = 1;
                break;
            case FL_FITS_UNSIGN:
                primitive_representation = 1;
                break;
            }
            break;
        case ASN_BASIC_OBJECT_IDENTIFIER:
            break;
        case ASN_BASIC_RELATIVE_OID:
            break;
        default: {
            char *p = strdup(MKID(cell->value));
            FATAL("Unsupported type %s for value %s",
                  asn1c_type_name(arg, cell->value, TNF_UNMODIFIED), p);
            free(p);
            return NULL;
        }
        }
        abuf_printf(&ab, "%s ", asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, cell->value, 0));

        asn1p_expr_t *expr_value = cell->value;
        while(expr_value->value->type == ATV_REFERENCED) {
            expr_value = WITH_MODULE_NAMESPACE(
                expr_value->module, expr_ns,
                asn1f_lookup_symbol_ex(arg->asn, expr_ns, expr_value,
                                       expr_value->value->value.reference));
            if(!expr_value) {
                FATAL("Unrecognized value type for %s", MKID(cell->value));
                return NULL;
            }
        }

        switch(expr_value->value->type) {
        case ATV_INTEGER:
            if(primitive_representation) {
                abuf_printf(&ab, "%s ", asn1p_itoa(expr_value->value->value.v_integer));
                break;
            }
        case ATV_UNPARSED:
            FATAL("Inappropriate value %s for type %s",
                  asn1f_printable_value(expr_value->value), MKID(cell->value));
            return NULL;   /* TEMPORARY FIXME FIXME */
        default:
            FATAL("Inappropriate value %s for type %s",
                  asn1f_printable_value(expr_value->value), MKID(cell->value));
            return NULL;
        }
    }
    return ab.buffer;
}

char *asn1c_get_ioc_cell(arg_t *arg, struct asn1p_ioc_cell_s *cell) {
//    if(strcmp(cell->field->Identifier, "&id") || strcmp(cell->field->Identifier, "&Value")) return NULL;
    OUT("/* %s ", cell->field->Identifier);
    if(!cell->value) {
    
    } else if(cell->value->meta_type == AMT_VALUE) {
        OUT("AMT_VALUE %s ", asn1go_type_name(arg, cell->value, TNF_RSAFE));
        OUT("%s", asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, cell->value, 0));
    } else if(cell->value->meta_type == AMT_TYPEREF) {
        OUT("AMT_TYPEREF %s", MKID(cell->value));
    } else if(cell->value->meta_type == AMT_TYPE){
        OUT("AMT_TYPE %s", asn1go_type_name(arg, cell->value, TNF_RSAFE));
    } else {
        OUT("Mata Type no right");
    }
    OUT(" */\n");
	return cell->field->Identifier;
}

static asn1p_expr_type_e
_find_terminal_type(arg_t *arg) {
	asn1p_expr_t *expr;
	expr = asn1f_find_terminal_type_ex(arg->asn, arg->ns, arg->expr);
	if(expr) return expr->expr_type;
	return A1TC_INVALID;
}

static asn1p_expr_type_e
expr_get_type(arg_t *arg, asn1p_expr_t *expr) {
	asn1p_expr_t *terminal;
	terminal = asn1f_find_terminal_type_ex(arg->asn, arg->ns, expr);
	if(terminal) return terminal->expr_type;
	return A1TC_INVALID;
}

asn1p_expr_t *
get_topmost_expr(arg_t *arg, asn1p_expr_t *expr) {
    asn1p_expr_t *topmost_parent;

    topmost_parent = WITH_MODULE_NAMESPACE(
        expr->module, expr_ns,
        asn1f_find_terminal_type_ex(arg->asn, expr_ns, expr));
    if(!topmost_parent) return NULL;

	if(!(topmost_parent->expr_type & ASN_CONSTR_MASK)
	&& !(topmost_parent->expr_type == ASN_BASIC_INTEGER)
	&& !(topmost_parent->expr_type == ASN_BASIC_ENUMERATED)
	&& !(topmost_parent->expr_type == ASN_BASIC_BIT_STRING))
		return NULL;

    return topmost_parent;
}

static int
expr_elements_count(arg_t *arg, asn1p_expr_t *expr) {
	asn1p_expr_t *topmost_parent;
	asn1p_expr_t *v;
	int elements = 0;

    topmost_parent = WITH_MODULE_NAMESPACE(
        expr->module, expr_ns,
        asn1f_find_terminal_type_ex(arg->asn, expr_ns, expr));
    if(!topmost_parent) return 0;

	if(!(topmost_parent->expr_type & ASN_CONSTR_MASK)
	&& !(topmost_parent->expr_type == ASN_BASIC_INTEGER)
	&& !(topmost_parent->expr_type == ASN_BASIC_ENUMERATED)
	&& !(topmost_parent->expr_type == ASN_BASIC_BIT_STRING))
		return 0;

	TQ_FOR(v, &(topmost_parent->members), next) {
		if(v->expr_type != A1TC_EXTENSIBLE)
			elements++;
	}

	return elements;
}

static int
expr_extensible_conut(arg_t *arg, asn1p_expr_t *expr) {
    asn1p_expr_t *topmost_parent;
	asn1p_expr_t *v;
	int elements = 0;

    topmost_parent = WITH_MODULE_NAMESPACE(
        expr->module, expr_ns,
        asn1f_find_terminal_type_ex(arg->asn, expr_ns, expr));
    if(!topmost_parent) return 0;

	if(!(topmost_parent->expr_type & ASN_CONSTR_MASK)
	&& !(topmost_parent->expr_type == ASN_BASIC_INTEGER)
	&& !(topmost_parent->expr_type == ASN_BASIC_ENUMERATED)
	&& !(topmost_parent->expr_type == ASN_BASIC_BIT_STRING))
		return 0;

	TQ_FOR(v, &(topmost_parent->members), next) {
		if(v->expr_type == A1TC_EXTENSIBLE)
			elements++;
	}

	return elements;
}

const char *get_ref_field_val_from_IEs(const char *id_name) {
    char *ptr;
    static abuf ab;
    
    abuf_clear(&ab);
    abuf_printf(&ab, "%s ", id_name);

    for(int i = 0; IEsToID[i]; i++) {
        ptr = strstr(IEsToID[i], ab.buffer);
        if(ptr) {
            int ptr_cnt = 0;
            ptr += strlen(id_name);
            while(*ptr == ' ') ptr++;
            while(isdigit(*(ptr + ptr_cnt))) ptr_cnt++;
            abuf_clear(&ab);
            abuf_add_bytes(&ab, ptr, ptr_cnt);

            if(ab.length) return ab.buffer;
        }
    }

    return NULL;
}

const char *get_ref_field_val_from_Procedures(const char *id_name) {
    char *ptr;
    static abuf ab;
    
    abuf_clear(&ab);
    abuf_printf(&ab, "%s ", id_name);

    for(int i = 0; ProceduresToID[i]; i++) {
        ptr = strstr(ProceduresToID[i], ab.buffer);
        if(ptr) {
            int ptr_cnt = 0;
            ptr += strlen(id_name);
            while(*ptr == ' ') ptr++;
            while(isdigit(*(ptr + ptr_cnt))) ptr_cnt++;
            abuf_clear(&ab); 
            abuf_add_bytes(&ab, ptr, ptr_cnt);

            if(ab.length) return ab.buffer;
        }
    }
    return NULL;
}



int 
asn1c_lang_Go_emit_constraint_checking_code(arg_t *arg, asn1p_expr_t *target) {
	asn1cnst_range_t *r_size;
	asn1cnst_range_t *r_value;
	asn1p_expr_t *expr = arg->expr;
	asn1p_expr_type_e etype;
	asn1p_constraint_t *ct;
	static abuf ab;
	int value_unsigned = 0;
    int cnt = 0;
	
	abuf_clear(&ab);
   
    asn1p_expr_t *terminal_expr = asn1f_find_terminal_type_ex(arg->asn, arg->ns, target);

//	OUT("/* Embed = %d, Type = %d, Terminal Type = %d */\n", arg->embed, expr->expr_type, terminal_expr->expr_type);
	if(!(arg->embed && terminal_expr->expr_type >= ASN_BASIC_MASK) && expr_extensible_conut(arg, expr)) {
		if(cnt++) abuf_str(&ab, ",");
		abuf_str(&ab, "valueExt");
	}

    if(expr->parent_expr != NULL && expr->parent_expr->expr_type == ASN_CONSTR_OPEN_TYPE) {
        
        asn1p_expr_t *v;
        const char *ref_field_val_ptr = NULL;
        TQ_FOR(v, &(target->members), next) {
            if(strcmp(asn1go_type_name(arg, v, TNF_RSAFE), "ProcedureCode") == 0) {
                ref_field_val_ptr = get_ref_field_val_from_Procedures(asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, v, 0));
            }
            else if(strcmp(asn1go_type_name(arg, v, TNF_RSAFE), "ProtocolIEID") == 0) {
                ref_field_val_ptr = get_ref_field_val_from_IEs(asn1c_make_identifier(AMI_TITLE | AMI_NODELIMITER, v, 0));
            } else {
                ref_field_val_ptr = NULL;
            }
        }

        if(ref_field_val_ptr) {
            if(cnt++) abuf_str(&ab, ",");
            abuf_str(&ab, "referenceFieldValue:");
            abuf_str(&ab, ref_field_val_ptr);
        } else {
            OUT("/* Do not find the Reference Field Value */ "); 
        }

        int member_cnt = expr_elements_count(arg, terminal_expr);
        if(member_cnt >= 0 && terminal_expr->expr_type == ASN_CONSTR_CHOICE) {
            if(cnt++) abuf_str(&ab, ",");
            abuf_printf(&ab, "valueLB:0,valueUB:%d", (member_cnt - 1) > 0 ? (member_cnt - 1) : 0);
        }
        
        if(ab.length) OUT("`aper:\"%s\"`", ab.buffer);
        return 0;
    }

    if(expr->expr_type == ASN_CONSTR_OPEN_TYPE) {
        abuf_str(&ab, "openType,referenceFieldName:");
		abuf_str(&ab, asn1c_get_member_type_selector(arg, arg->expr));
        cnt++;
    }

  
	ct = expr->combined_constraints;
	if(ct == NULL) {        /* No additional constraints defined */
//		printf("[Debug] ct == NULL : %d, %d\n", expr->expr_type, A1TC_REFERENCE);
		if(expr->expr_type == A1TC_REFERENCE) {
			arg_t tmp_arg = *arg;
			asn1p_expr_t *ref_type;
			
			ref_type = WITH_MODULE_NAMESPACE(
				target->module, expr_ns,
				asn1f_lookup_symbol_ex(tmp_arg.asn, expr_ns, tmp_arg.expr,
				target->reference));
			if(!ref_type) return 0;
//			printf("[Debug] ct == NULL : %d, %d\n", ref_type->expr_type, A1TC_REFERENCE);
			tmp_arg.expr = ref_type;
			int flag_backup = tmp_arg.expr->marker.flags;
			tmp_arg.expr->marker.flags = expr->marker.flags;
			asn1c_lang_Go_emit_constraint_checking_code(&tmp_arg, ref_type);
			tmp_arg.expr->marker.flags = flag_backup;
			return 0;
		}
        if(expr->expr_type == ASN_BASIC_ENUMERATED && arg->embed == 0) {
//	        printf("[Debug] emit_constraint_checking_code = ASN_BASIC_ENUMERATED\n");
			int member_cnt = expr_elements_count(arg, target);
            /* Zero base */
            if(cnt++) abuf_str(&ab, ",");
            abuf_printf(&ab, "valueLB:0,valueUB:%d", (member_cnt - 1) > 0 ? (member_cnt - 1) : 0);
            cnt++;
        } else if(expr->expr_type == ASN_CONSTR_CHOICE) {
//	        printf("[Debug] emit_constraint_checking_code = ASN_CONSTR_CHOICE\n");
			int member_cnt = expr_elements_count(arg, target);
            if(cnt++) abuf_str(&ab, ",");
            abuf_printf(&ab, "valueLB:0,valueUB:%d", (member_cnt - 1) > 0 ? (member_cnt - 1) : 0);
			cnt++;
		}
        if(target->marker.flags & EM_OPTIONAL) {
            if(cnt++) abuf_str(&ab, ",");
            abuf_str(&ab, "optional");
        }
        if(ab.length) OUT("`aper:\"%s\"`", ab.buffer);
        return 0; 
    }

	etype = _find_terminal_type(arg);

    /* Check the type is reference to other basic type */
    if(arg->embed && arg->expr->expr_type != etype && etype >= 0x40 && etype < 0x50) {
		if(target->marker.flags & EM_OPTIONAL) {
			if(cnt++) abuf_str(&ab, ",");
			abuf_str(&ab, "optional");
		}
		if(ab.length) OUT("`aper:\"%s\"`", ab.buffer);
		return 0;
	}

    asn1p_expr_t *topmost_expr = get_topmost_expr(arg, target);

    if(is_spec_index(expr) && topmost_expr) {
//        OUT("/* expr->spec_index != -1 && expr->_lineno && topmost_expr */\n");
        expr = topmost_expr;
        etype = expr_get_type(arg, topmost_expr);
		ct = topmost_expr->combined_constraints;
    }

	if(terminal_expr->expr_type != ASN_CONSTR_SEQUENCE_OF || (arg->embed == 0)) {
		r_value=asn1constraint_compute_constraint_range(expr->Identifier, etype, ct, ACT_EL_RANGE,0,0,0);
		r_size =asn1constraint_compute_constraint_range(expr->Identifier, etype, ct, ACT_CT_SIZE, 0,0,0);

		if(r_value) {
			if(r_value->incompatible
			|| r_value->empty_constraint
			|| (r_value->left.type == ARE_MIN
				&& r_value->right.type == ARE_MAX)
			|| (etype == ASN_BASIC_BOOLEAN
				&& r_value->left.value == 0
				&& r_value->right.value == 1)
			) {
				asn1constraint_range_free(r_value);
				r_value = 0;
			}
		}
		if(r_size) {
			if(r_size->incompatible
			|| r_size->empty_constraint
			|| (r_size->left.value == 0	/* or .type == MIN */
				&& r_size->right.type == ARE_MAX)
			) {
				asn1constraint_range_free(r_size);
				r_size = 0;
			}
		}

	    if(expr_extensible_conut(arg, expr) || 
		   (expr->expr_type == ASN_CONSTR_SEQUENCE_OF && expr_extensible_conut(arg, TQ_FIRST(&expr->members)))) {
	        abuf_str(&ab, "valueExt");
        	cnt++;
    	}

		abuf *tmp;
		if(r_size) {
	        if(cnt++) abuf_str(&ab, ",");
			tmp = emit_range_comparison_code(r_size, "size", -1, -1);
			if(tmp->length) {
				if(r_size->extensible) abuf_str(&ab, "sizeExt,");
				abuf_str(&ab, tmp->buffer);
			}
		}
		if(r_value) {
	        if(cnt++) abuf_str(&ab, ",");
			if(etype == ASN_BASIC_BOOLEAN)
				tmp = emit_range_comparison_code(r_value, "value", 0, 1);
			else
				tmp = emit_range_comparison_code(r_value, "value",
    	                                        value_unsigned ? 0 : -1, -1);
			if(tmp->length) {
				if(r_value->extensible) abuf_str(&ab, "valueExt,");
				abuf_str(&ab, tmp->buffer);
			}
		}
	}
    if(target->marker.flags & EM_OPTIONAL) {
        if(cnt++) abuf_str(&ab, ",");
        abuf_str(&ab, "optional");
    }

	if(ab.length) {
		OUT("`aper:\"%s\"`", ab.buffer);
	}
	return 0;
}

static void
abuf_oint(abuf *ab, asn1c_integer_t v) {
    if(v == (-2147483647L - 1)) {
        abuf_printf(ab, "(-2147483647L - 1)");
    } else {
        abuf_printf(ab, "%s", asn1p_itoa(v));
    }
}

static abuf *
emit_range_comparison_code(asn1cnst_range_t *range, const char *varname,
                           asn1c_integer_t natural_start,
                           asn1c_integer_t natural_stop) {
    static abuf ab;
	
	abuf_clear(&ab);

    if(range->el_count == 0) {
        int ignore_left =
            (range->left.type == ARE_MIN)
            || (natural_start != -1 && range->left.value <= natural_start);
        int ignore_right =
            (range->right.type == ARE_MAX)
            || (natural_stop != -1 && range->right.value >= natural_stop);

        if(ignore_left && ignore_right) {
            /* Empty constraint comparison */
        } else if(ignore_left) {	
			abuf_printf(&ab, "%sLB:", varname);
			abuf_oint(&ab, range->right.value);
			abuf_printf(&ab, ",%sUB:", varname);
			abuf_oint(&ab, range->right.value);
        } else if(ignore_right) {
			abuf_printf(&ab, "%sLB:", varname);
			abuf_oint(&ab, range->left.value);
			abuf_printf(&ab, ",%sUB:", varname);
			abuf_oint(&ab, range->left.value);
        } else if(range->left.value == range->right.value) {
			abuf_printf(&ab, "%sLB:", varname);
			abuf_oint(&ab, range->right.value);
			abuf_printf(&ab, ",%sUB:", varname);
			abuf_oint(&ab, range->right.value);
        } else {
			abuf_printf(&ab, "%sLB:", varname);
			abuf_oint(&ab, range->left.value);
			abuf_printf(&ab, ",%sUB:", varname);
			abuf_oint(&ab, range->right.value);
        }
    } else {
//		printf("[Debug] Unpredictable constraint\n");
		/*
        for(int i = 0; i < range->el_count; i++) {
            asn1cnst_range_t *r = range->elements[i];

            abuf *rec = emit_range_comparison_code(r, varname, natural_start,
                                                   natural_stop);
            if(rec->length) {
                if(ab->length) {
                    abuf_str(ab, " || ");
                }
                abuf_str(ab, "(");
                abuf_buf(ab, rec);
                abuf_str(ab, ")");
            } else {
                // Ignore this part 
            }
            abuf_free(rec);
        }
		*/
    }

    return &ab;
}
