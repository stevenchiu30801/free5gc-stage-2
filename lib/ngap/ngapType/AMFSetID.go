package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AMFSetID struct {
	Value aper.BitString `aper:"sizeLB:10,sizeUB:10"`
}
