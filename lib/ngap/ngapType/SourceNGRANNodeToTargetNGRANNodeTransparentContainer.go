package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      RRCContainer
	PDUSessionResourceInformationList *PDUSessionResourceInformationList `aper:"optional"`
	ERABInformationList               *ERABInformationList               `aper:"optional"`
	TargetCellID                      NGRANCGI                           `aper:"valueLB:0,valueUB:2"`
	IndexToRFSP                       *IndexToRFSP                       `aper:"optional"`
	UEHistoryInformation              UEHistoryInformation
	IEExtensions                      *ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs `aper:"optional"`
}
