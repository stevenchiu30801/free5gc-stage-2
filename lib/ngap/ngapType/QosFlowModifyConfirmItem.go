package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs `aper:"optional"`
}
