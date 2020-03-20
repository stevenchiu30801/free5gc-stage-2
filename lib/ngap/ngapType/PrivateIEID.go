package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	PrivateIEIDPresentNothing int = iota /* No components present */
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

type PrivateIEID struct {
	Present int
	Local   *int64 `aper:"valueLB:0,valueUB:65535"`
	Global  *aper.ObjectIdentifier
}
