package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type EPSTAC struct {
	Value aper.OctetString `aper:"sizeLB:2,sizeUB:2"`
}