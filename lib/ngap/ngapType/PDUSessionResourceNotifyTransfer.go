package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   *QosFlowNotifyList                                                `aper:"optional"`
	QosFlowReleasedList *QosFlowList                                                      `aper:"optional"`
	IEExtensions        *ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs `aper:"optional"`
}
