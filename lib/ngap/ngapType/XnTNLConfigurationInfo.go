package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type XnTNLConfigurationInfo struct {
	XnTransportLayerAddresses         XnTLAs
	XnExtendedTransportLayerAddresses *XnExtTLAs                                              `aper:"optional"`
	IEExtensions                      *ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs `aper:"optional"`
}
