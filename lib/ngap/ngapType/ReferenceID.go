package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type ReferenceID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:64"`
}
