package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct NR_CGIListForWarning */
/* NRCGI */
type NRCGIListForWarning struct {
	List []NRCGI `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
