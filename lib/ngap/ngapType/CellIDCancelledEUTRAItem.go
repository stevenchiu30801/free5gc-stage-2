package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CellIDCancelledEUTRAItem struct {
	EUTRACGI           EUTRACGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs `aper:"optional"`
}
