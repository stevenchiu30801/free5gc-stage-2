package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	CriticalityPresentReject aper.Enumerated = 0
	CriticalityPresentIgnore aper.Enumerated = 1
	CriticalityPresentNotify aper.Enumerated = 2
)

type Criticality struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}
