package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type QosFlowItem struct {
	QosFlowIdentifier QosFlowIdentifier
	Cause             Cause                                        `aper:"valueLB:0,valueUB:5"`
	IEExtensions      *ProtocolExtensionContainerQosFlowItemExtIEs `aper:"optional"`
}
