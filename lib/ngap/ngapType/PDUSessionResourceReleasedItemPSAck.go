package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceReleasedItemPSAck struct {
	PDUSessionID                          PDUSessionID
	PathSwitchRequestUnsuccessfulTransfer aper.OctetString
	IEExtensions                          *ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs `aper:"optional"`
}
