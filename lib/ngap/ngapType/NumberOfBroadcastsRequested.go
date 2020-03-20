package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type NumberOfBroadcastsRequested struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}
