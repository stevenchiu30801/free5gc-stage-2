package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct AMF_TNLAssociationToAddList */
/* AMFTNLAssociationToAddItem */
type AMFTNLAssociationToAddList struct {
	List []AMFTNLAssociationToAddItem `aper:"valueExt,sizeLB:1,sizeUB:32"`
}
