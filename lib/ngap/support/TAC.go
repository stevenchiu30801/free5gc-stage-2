package ngap

import (
	"gofree5gc/lib/aper"
)

// TAC Type
type TAC struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
