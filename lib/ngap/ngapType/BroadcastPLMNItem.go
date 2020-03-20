package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type BroadcastPLMNItem struct {
	PLMNIdentity        PLMNIdentity
	TAISliceSupportList SliceSupportList
	IEExtensions        *ProtocolExtensionContainerBroadcastPLMNItemExtIEs `aper:"optional"`
}
