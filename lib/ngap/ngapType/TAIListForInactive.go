package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct TAIListForInactive */
/* TAIListForInactiveItem */
type TAIListForInactive struct {
	List []TAIListForInactiveItem `aper:"valueExt,sizeLB:1,sizeUB:16"`
}
