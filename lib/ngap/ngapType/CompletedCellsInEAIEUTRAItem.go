package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs `aper:"optional"`
}
