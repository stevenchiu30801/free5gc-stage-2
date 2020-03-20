package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR                             `aper:"optional"`
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA                          `aper:"optional"`
	IEExtensions                      *ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs `aper:"optional"`
}
