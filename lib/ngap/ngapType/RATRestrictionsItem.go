package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type RATRestrictionsItem struct {
	PLMNIdentity              PLMNIdentity
	RATRestrictionInformation RATRestrictionInformation
	IEExtensions              *ProtocolExtensionContainerRATRestrictionsItemExtIEs `aper:"optional"`
}
