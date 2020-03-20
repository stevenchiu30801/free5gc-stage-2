package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        QosFlowIdentifier
	QosFlowMappingIndication *aper.Enumerated                                       `aper:"optional"`
	IEExtensions             *ProtocolExtensionContainerAssociatedQosFlowItemExtIEs `aper:"optional"`
}
