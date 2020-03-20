package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type EUTRACellIdentity struct {
	Value aper.BitString `aper:"sizeLB:28,sizeUB:28"`
}
