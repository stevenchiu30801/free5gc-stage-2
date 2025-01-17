package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	NGAPPDUPresentNothing int = iota /* No components present */
	NGAPPDUPresentInitiatingMessage
	NGAPPDUPresentSuccessfulOutcome
	NGAPPDUPresentUnsuccessfulOutcome
	/* Extensions may appear below */
)

type NGAPPDU struct {
	Present             int
	InitiatingMessage   *InitiatingMessage
	SuccessfulOutcome   *SuccessfulOutcome
	UnsuccessfulOutcome *UnsuccessfulOutcome
}
