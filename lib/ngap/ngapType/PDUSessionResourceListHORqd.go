package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceListHORqd */
/* PDUSessionResourceItemHORqd */
type PDUSessionResourceListHORqd struct {
	List []PDUSessionResourceItemHORqd `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
