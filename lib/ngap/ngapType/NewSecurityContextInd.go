package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	NewSecurityContextIndPresentTrue aper.Enumerated = 0
)

type NewSecurityContextInd struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
