package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type GUAMI struct {
	PLMNIdentity PLMNIdentity
	AMFRegionID  AMFRegionID
	AMFSetID     AMFSetID
	AMFPointer   AMFPointer
	IEExtensions *ProtocolExtensionContainerGUAMIExtIEs `aper:"optional"`
}
