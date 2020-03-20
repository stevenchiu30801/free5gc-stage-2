package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceFailedToModifyItemModRes struct {
	PDUSessionID                                 PDUSessionID
	PDUSessionResourceModifyUnsuccessfulTransfer aper.OctetString
	IEExtensions                                 *ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs `aper:"optional"`
}
