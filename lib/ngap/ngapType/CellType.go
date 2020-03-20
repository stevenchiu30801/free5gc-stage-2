package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CellType struct {
	CellSize     CellSize
	IEExtensions *ProtocolExtensionContainerCellTypeExtIEs `aper:"optional"`
}
