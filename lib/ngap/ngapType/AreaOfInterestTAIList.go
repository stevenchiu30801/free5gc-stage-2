package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct AreaOfInterestTAIList */
/* AreaOfInterestTAIItem */
type AreaOfInterestTAIList struct {
	List []AreaOfInterestTAIItem `aper:"valueExt,sizeLB:1,sizeUB:16"`
}
