package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type TAC struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
