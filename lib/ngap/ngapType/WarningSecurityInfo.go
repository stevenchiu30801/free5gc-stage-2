package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type WarningSecurityInfo struct {
	Value aper.OctetString `aper:"sizeLB:50,sizeUB:50"`
}
