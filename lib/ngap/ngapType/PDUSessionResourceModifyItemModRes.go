package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionResourceModifyItemModRes struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceModifyResponseTransfer *aper.OctetString                                                   `aper:"optional"`
	IEExtensions                             *ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs `aper:"optional"`
}
