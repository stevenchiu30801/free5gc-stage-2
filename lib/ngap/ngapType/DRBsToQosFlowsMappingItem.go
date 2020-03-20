package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type DRBsToQosFlowsMappingItem struct {
	DRBID                 DRBID
	AssociatedQosFlowList AssociatedQosFlowList
	IEExtensions          *ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs `aper:"optional"`
}
