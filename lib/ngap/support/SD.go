package ngap

import (
	"gofree5gc/lib/aper"
)

// SD Type
type SD struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
