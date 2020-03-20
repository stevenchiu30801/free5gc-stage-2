package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type PDUSessionAggregateMaximumBitRate struct {
	PDUSessionAggregateMaximumBitRateDL BitRate
	PDUSessionAggregateMaximumBitRateUL BitRate
	IEExtensions                        *ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs `aper:"optional"`
}
