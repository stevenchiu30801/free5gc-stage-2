package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList  QosFlowModifyConfirmList
	TNLMappingList            *TNLMappingList                                                          `aper:"optional"`
	QosFlowFailedToModifyList *QosFlowList                                                             `aper:"optional"`
	IEExtensions              *ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs `aper:"optional"`
}
