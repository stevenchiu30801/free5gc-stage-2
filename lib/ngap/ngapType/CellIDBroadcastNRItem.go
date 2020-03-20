package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CellIDBroadcastNRItem struct {
	NRCGI        NRCGI                                                  `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs `aper:"optional"`
}
