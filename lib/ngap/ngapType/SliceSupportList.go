package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct SliceSupportList */
/* SliceSupportItem */
type SliceSupportList struct {
	List []SliceSupportItem `aper:"valueExt,sizeLB:1,sizeUB:1024"`
}
