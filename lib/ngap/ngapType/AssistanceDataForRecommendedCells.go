package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AssistanceDataForRecommendedCells struct {
	RecommendedCellsForPaging RecommendedCellsForPaging                                          `aper:"valueExt"`
	IEExtensions              *ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs `aper:"optional"`
}
