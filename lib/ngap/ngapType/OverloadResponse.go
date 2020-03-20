package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	OverloadResponsePresentNothing int = iota /* No components present */
	OverloadResponsePresentOverloadAction
	OverloadResponsePresentChoiceExtensions
)

type OverloadResponse struct {
	Present          int
	OverloadAction   *OverloadAction
	ChoiceExtensions *ProtocolIESingleContainerOverloadResponseExtIEs
}
