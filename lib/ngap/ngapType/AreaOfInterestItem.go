package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest `aper:"valueExt"`
	LocationReportingReferenceID LocationReportingReferenceID
	IEExtensions                 *ProtocolExtensionContainerAreaOfInterestItemExtIEs `aper:"optional"`
}
