package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PLMNSupportItem struct {
	PLMNIdentity     PLMNIdentity
	SliceSupportList SliceSupportList
	IEExtensions     *ProtocolExtensionContainerPLMNSupportItemExtIEs `aper:"optional"`
}
