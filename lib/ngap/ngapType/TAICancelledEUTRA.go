package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct TAICancelledEUTRA */
/* TAICancelledEUTRAItem */
type TAICancelledEUTRA struct {
	List []TAICancelledEUTRAItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
