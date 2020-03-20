package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CompletedCellsInTAINRItem struct {
	NRCGI        NRCGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs `aper:"optional"`
}
