package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SliceOverloadItem struct {
	SNSSAI       SNSSAI                                             `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerSliceOverloadItemExtIEs `aper:"optional"`
}
