package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	CellIDListForRestartPresentNothing int = iota /* No components present */
	CellIDListForRestartPresentEUTRACGIListforRestart
	CellIDListForRestartPresentNRCGIListforRestart
	CellIDListForRestartPresentChoiceExtensions
)

type CellIDListForRestart struct {
	Present                int
	EUTRACGIListforRestart *EUTRACGIList
	NRCGIListforRestart    *NRCGIList
	ChoiceExtensions       *ProtocolIESingleContainerCellIDListForRestartExtIEs
}
