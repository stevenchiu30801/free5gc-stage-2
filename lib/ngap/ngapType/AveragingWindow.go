package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type AveragingWindow struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:4095"`
}
