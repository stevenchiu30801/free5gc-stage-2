package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowAcceptedList */
/* QosFlowAcceptedItem */
type QosFlowAcceptedList struct {
	List []QosFlowAcceptedItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}
