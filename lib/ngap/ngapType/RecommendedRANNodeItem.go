package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type RecommendedRANNodeItem struct {
	AMFPagingTarget AMFPagingTarget                                         `aper:"valueLB:0,valueUB:2"`
	IEExtensions    *ProtocolExtensionContainerRecommendedRANNodeItemExtIEs `aper:"optional"`
}
