package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	AdditionalQosFlowInformationPresentMoreLikely aper.Enumerated = 0
)

type AdditionalQosFlowInformation struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
