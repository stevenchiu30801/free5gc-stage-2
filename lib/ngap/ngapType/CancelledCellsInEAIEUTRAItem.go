package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs `aper:"optional"`
}
