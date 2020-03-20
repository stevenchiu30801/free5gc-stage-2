package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceSwitchedItem struct {
	PDUSessionID                         PDUSessionID
	PathSwitchRequestAcknowledgeTransfer aper.OctetString
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs `aper:"optional"`
}
