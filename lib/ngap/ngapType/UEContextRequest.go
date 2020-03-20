package ngapType

import "gofree5gc/lib/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	UEContextRequestPresentRequested aper.Enumerated = 0
)

type UEContextRequest struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
