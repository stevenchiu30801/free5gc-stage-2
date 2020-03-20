package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SliceSupportItem struct {
	SNSSAI       SNSSAI                                            `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerSliceSupportItemExtIEs `aper:"optional"`
}
