package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowSetupResponseListSURes */
/* QosFlowSetupResponseItemSURes */
type QosFlowSetupResponseListSURes struct {
	List []QosFlowSetupResponseItemSURes `aper:"valueExt,sizeLB:1,sizeUB:64"`
}
