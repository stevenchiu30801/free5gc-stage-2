package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	IntegrityProtectionResultPresentPerformed    aper.Enumerated = 0
	IntegrityProtectionResultPresentNotPerformed aper.Enumerated = 1
)

type IntegrityProtectionResult struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
