package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct TNLInformationList */
/* TNLInformationItem */
type TNLInformationList struct {
	List []TNLInformationItem `aper:"valueExt,sizeLB:1,sizeUB:4"`
}
