package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceToBeSwitchedDLList */
/* PDUSessionResourceToBeSwitchedDLItem */
type PDUSessionResourceToBeSwitchedDLList struct {
	List []PDUSessionResourceToBeSwitchedDLItem `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
