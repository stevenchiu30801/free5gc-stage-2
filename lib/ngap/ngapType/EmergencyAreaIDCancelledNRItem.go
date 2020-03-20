package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       EmergencyAreaID
	CancelledCellsInEAINR CancelledCellsInEAINR
	IEExtensions          *ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs `aper:"optional"`
}
