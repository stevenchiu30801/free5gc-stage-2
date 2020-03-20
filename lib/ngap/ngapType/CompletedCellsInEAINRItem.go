package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CompletedCellsInEAINRItem struct {
	NRCGI        NRCGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs `aper:"optional"`
}
