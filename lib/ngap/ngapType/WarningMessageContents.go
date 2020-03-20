package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type WarningMessageContents struct {
	Value aper.OctetString `aper:"sizeLB:1,sizeUB:9600"`
}
