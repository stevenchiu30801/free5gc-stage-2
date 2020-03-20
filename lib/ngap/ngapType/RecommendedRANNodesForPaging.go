package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList RecommendedRANNodeList
	IEExtensions           *ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs `aper:"optional"`
}
