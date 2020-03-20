package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceItemCxtRelCpl struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs `aper:"optional"`
}
