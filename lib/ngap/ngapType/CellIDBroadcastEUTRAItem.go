package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CellIDBroadcastEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                  `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs `aper:"optional"`
}
