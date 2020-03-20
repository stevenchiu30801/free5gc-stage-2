package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type GTPTEID struct {
	Value aper.OctetString `aper:"sizeLB:4,sizeUB:4"`
}
