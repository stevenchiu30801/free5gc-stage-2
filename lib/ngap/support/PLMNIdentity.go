package ngap

import (
	"gofree5gc/lib/aper"
)

// PLMNIdentity Type
type PLMNIdentity struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
