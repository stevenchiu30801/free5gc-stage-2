package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AreaOfInterestCellItem struct {
	NGRANCGI     NGRANCGI                                                `aper:"valueLB:0,valueUB:2"`
	IEExtensions *ProtocolExtensionContainerAreaOfInterestCellItemExtIEs `aper:"optional"`
}
