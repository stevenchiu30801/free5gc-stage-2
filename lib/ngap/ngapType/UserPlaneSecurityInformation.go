package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type UserPlaneSecurityInformation struct {
	SecurityResult     SecurityResult                                                `aper:"valueExt"`
	SecurityIndication SecurityIndication                                            `aper:"valueExt"`
	IEExtensions       *ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs `aper:"optional"`
}
