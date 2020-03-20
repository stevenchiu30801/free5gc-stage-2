package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AreaOfInterestTAIItem struct {
	TAI          TAI                                                    `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs `aper:"optional"`
}
