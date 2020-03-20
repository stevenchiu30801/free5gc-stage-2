package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AllowedNSSAIItem struct {
	SNSSAI       SNSSAI                                            `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerAllowedNSSAIItemExtIEs `aper:"optional"`
}
