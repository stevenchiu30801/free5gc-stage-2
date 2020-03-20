package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceFailedToSetupItemPSReq struct {
	PDUSessionID                         PDUSessionID
	PathSwitchRequestSetupFailedTransfer aper.OctetString
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs `aper:"optional"`
}
