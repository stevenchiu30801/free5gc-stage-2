package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	UENGAPIDsPresentNothing int = iota /* No components present */
	UENGAPIDsPresentUENGAPIDPair
	UENGAPIDsPresentAMFUENGAPID
	UENGAPIDsPresentChoiceExtensions
)

type UENGAPIDs struct {
	Present          int
	UENGAPIDPair     *UENGAPIDPair `aper:"valueExt"`
	AMFUENGAPID      *AMFUENGAPID
	ChoiceExtensions *ProtocolIESingleContainerUENGAPIDsExtIEs
}
