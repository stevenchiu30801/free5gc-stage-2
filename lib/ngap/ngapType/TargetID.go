package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	TargetIDPresentNothing int = iota /* No components present */
	TargetIDPresentTargetRANNodeID
	TargetIDPresentTargeteNBID
	TargetIDPresentChoiceExtensions
)

type TargetID struct {
	Present          int
	TargetRANNodeID  *TargetRANNodeID `aper:"valueExt"`
	TargeteNBID      *TargeteNBID     `aper:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerTargetIDExtIEs
}
