package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs `aper:"optional"`
}
