package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type TNLMappingItem struct {
	DLNGUUPTNLInformation UPTransportLayerInformation                     `aper:"valueLB:0,valueUB:1"`
	ULNGUUPTNLInformation UPTransportLayerInformation                     `aper:"valueLB:0,valueUB:1"`
	IEExtensions          *ProtocolExtensionContainerTNLMappingItemExtIEs `aper:"optional"`
}
