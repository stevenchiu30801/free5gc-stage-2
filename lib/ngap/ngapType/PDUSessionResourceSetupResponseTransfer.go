package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceSetupResponseTransfer struct {
	QosFlowPerTNLInformation           QosFlowPerTNLInformation                                                 `aper:"valueExt"`
	AdditionalQosFlowPerTNLInformation *QosFlowPerTNLInformation                                                `aper:"valueExt,optional"`
	SecurityResult                     *SecurityResult                                                          `aper:"valueExt,optional"`
	QosFlowFailedToSetupList           *QosFlowList                                                             `aper:"optional"`
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs `aper:"optional"`
}
