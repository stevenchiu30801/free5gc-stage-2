package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type MultipleTNLInformation struct {
	TNLInformationList TNLInformationList
	IEExtensions       *ProtocolExtensionContainerMultipleTNLInformationExtIEs `aper:"optional"`
}
