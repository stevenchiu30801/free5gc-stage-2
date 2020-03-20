package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type TAIListForPagingItem struct {
	TAI          TAI                                                   `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerTAIListForPagingItemExtIEs `aper:"optional"`
}
