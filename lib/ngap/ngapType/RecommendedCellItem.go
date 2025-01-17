package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type RecommendedCellItem struct {
	NGRANCGI         NGRANCGI                                             `aper:"valueLB:0,valueUB:2"`
	TimeStayedInCell *int64                                               `aper:"valueLB:0,valueUB:4095,optional"`
	IEExtensions     *ProtocolExtensionContainerRecommendedCellItemExtIEs `aper:"optional"`
}
