package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        PriorityLevelARP
	PreEmptionCapability    PreEmptionCapability
	PreEmptionVulnerability PreEmptionVulnerability
	IEExtensions            *ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs `aper:"optional"`
}
