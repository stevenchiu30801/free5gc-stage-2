package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct OverloadStartNSSAIList */
/* OverloadStartNSSAIItem */
type OverloadStartNSSAIList struct {
	List []OverloadStartNSSAIItem `aper:"valueExt,sizeLB:1,sizeUB:1024"`
}
