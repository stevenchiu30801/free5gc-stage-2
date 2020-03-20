package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs `aper:"optional"`
}
