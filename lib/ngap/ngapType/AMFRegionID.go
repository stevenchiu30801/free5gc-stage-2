package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AMFRegionID struct {
	Value aper.BitString `aper:"sizeLB:8,sizeUB:8"`
}
