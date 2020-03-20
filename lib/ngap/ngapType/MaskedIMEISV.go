package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type MaskedIMEISV struct {
	Value aper.BitString `aper:"sizeLB:64,sizeUB:64"`
}
