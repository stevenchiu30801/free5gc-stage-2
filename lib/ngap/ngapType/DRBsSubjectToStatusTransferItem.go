package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type DRBsSubjectToStatusTransferItem struct {
	DRBID       DRBID
	DRBStatusUL DRBStatusUL                                                      `aper:"valueLB:0,valueUB:2"`
	DRBStatusDL DRBStatusDL                                                      `aper:"valueLB:0,valueUB:2"`
	IEExtension *ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs `aper:"optional"`
}
