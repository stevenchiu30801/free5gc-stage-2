package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality
	IEID          ProtocolIEID
	TypeOfError   TypeOfError
	IEExtensions  *ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs `aper:"optional"`
}
