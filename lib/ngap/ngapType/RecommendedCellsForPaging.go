package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type RecommendedCellsForPaging struct {
	RecommendedCellList RecommendedCellList
	IEExtensions        *ProtocolExtensionContainerRecommendedCellsForPagingExtIEs `aper:"optional"`
}
