package ngap

import (
	"gofree5gc/lib/aper"
)

// SST Type
type SST struct {
	Value aper.OctetString `aper:"sizeLB:1,sizeUB:1"`
}
