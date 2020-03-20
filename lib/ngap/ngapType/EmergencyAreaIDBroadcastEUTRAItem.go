package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type EmergencyAreaIDBroadcastEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID
	CompletedCellsInEAIEUTRA CompletedCellsInEAIEUTRA
	IEExtensions             *ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs `aper:"optional"`
}
