package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type UESecurityCapabilities struct {
	NRencryptionAlgorithms             NRencryptionAlgorithms
	NRintegrityProtectionAlgorithms    NRintegrityProtectionAlgorithms
	EUTRAencryptionAlgorithms          EUTRAencryptionAlgorithms
	EUTRAintegrityProtectionAlgorithms EUTRAintegrityProtectionAlgorithms
	IEExtensions                       *ProtocolExtensionContainerUESecurityCapabilitiesExtIEs `aper:"optional"`
}
