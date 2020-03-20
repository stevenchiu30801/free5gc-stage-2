package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type TAIBroadcastNRItem struct {
	TAI                   TAI `aper:"valueExt"`
	CompletedCellsInTAINR CompletedCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAIBroadcastNRItemExtIEs `aper:"optional"`
}
