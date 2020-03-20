package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct TAICancelledNR */
/* TAICancelledNRItem */
type TAICancelledNR struct {
	List []TAICancelledNRItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
