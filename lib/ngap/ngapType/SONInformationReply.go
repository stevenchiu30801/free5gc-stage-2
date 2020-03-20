package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SONInformationReply struct {
	XnTNLConfigurationInfo *XnTNLConfigurationInfo                              `aper:"valueExt,optional"`
	IEExtensions           *ProtocolExtensionContainerSONInformationReplyExtIEs `aper:"optional"`
}
