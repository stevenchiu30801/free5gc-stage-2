package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct BroadcastPLMNList */
/* BroadcastPLMNItem */
type BroadcastPLMNList struct {
	List []BroadcastPLMNItem `aper:"valueExt,sizeLB:1,sizeUB:12"`
}
