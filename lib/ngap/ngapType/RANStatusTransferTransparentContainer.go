package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type RANStatusTransferTransparentContainer struct {
	DRBsSubjectToStatusTransferList DRBsSubjectToStatusTransferList
	IEExtensions                    *ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs `aper:"optional"`
}
