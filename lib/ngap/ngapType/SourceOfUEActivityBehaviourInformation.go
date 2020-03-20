package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	SourceOfUEActivityBehaviourInformationPresentSubscriptionInformation aper.Enumerated = 0
	SourceOfUEActivityBehaviourInformationPresentStatistics              aper.Enumerated = 1
)

type SourceOfUEActivityBehaviourInformation struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
