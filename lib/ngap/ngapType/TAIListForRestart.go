package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct TAIListForRestart */
/* TAI */
type TAIListForRestart struct {
	List []TAI `aper:"valueExt,sizeLB:1,sizeUB:2048"`
}
