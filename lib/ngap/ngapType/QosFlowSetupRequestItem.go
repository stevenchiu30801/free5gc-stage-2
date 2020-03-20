package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters QosFlowLevelQosParameters                                `aper:"valueExt"`
	ERABID                    *ERABID                                                  `aper:"optional"`
	IEExtensions              *ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs `aper:"optional"`
}
