package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct CancelledCellsInEAI_NR */
/* CancelledCellsInEAINRItem */
type CancelledCellsInEAINR struct {
	List []CancelledCellsInEAINRItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
