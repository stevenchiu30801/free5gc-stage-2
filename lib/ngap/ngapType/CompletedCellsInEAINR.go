package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct CompletedCellsInEAI_NR */
/* CompletedCellsInEAINRItem */
type CompletedCellsInEAINR struct {
	List []CompletedCellsInEAINRItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
