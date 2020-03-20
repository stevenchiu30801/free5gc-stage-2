package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            PDUSessionID
	SNSSAI                  SNSSAI `aper:"valueExt"`
	HandoverRequestTransfer aper.OctetString
	IEExtensions            *ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs `aper:"optional"`
}
