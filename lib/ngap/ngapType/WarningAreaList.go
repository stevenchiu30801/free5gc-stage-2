package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	WarningAreaListPresentNothing int = iota /* No components present */
	WarningAreaListPresentEUTRACGIListForWarning
	WarningAreaListPresentNRCGIListForWarning
	WarningAreaListPresentTAIListForWarning
	WarningAreaListPresentEmergencyAreaIDList
	WarningAreaListPresentChoiceExtensions
)

type WarningAreaList struct {
	Present                int
	EUTRACGIListForWarning *EUTRACGIListForWarning
	NRCGIListForWarning    *NRCGIListForWarning
	TAIListForWarning      *TAIListForWarning
	EmergencyAreaIDList    *EmergencyAreaIDList
	ChoiceExtensions       *ProtocolIESingleContainerWarningAreaListExtIEs
}
