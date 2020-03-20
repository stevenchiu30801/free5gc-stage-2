package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters *QosFlowLevelQosParameters                                     `aper:"valueExt,optional"`
	ERABID                    *ERABID                                                        `aper:"optional"`
	IEExtensions              *ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs `aper:"optional"`
}
