package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs `aper:"optional"`
}
