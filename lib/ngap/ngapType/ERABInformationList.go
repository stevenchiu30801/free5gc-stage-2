package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct E_RABInformationList */
/* ERABInformationItem */
type ERABInformationList struct {
	List []ERABInformationItem `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
