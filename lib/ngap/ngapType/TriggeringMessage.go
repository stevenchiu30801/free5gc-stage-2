package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	TriggeringMessagePresentInitiatingMessage    aper.Enumerated = 0
	TriggeringMessagePresentSuccessfulOutcome    aper.Enumerated = 1
	TriggeringMessagePresentUnsuccessfullOutcome aper.Enumerated = 2
)

type TriggeringMessage struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}
