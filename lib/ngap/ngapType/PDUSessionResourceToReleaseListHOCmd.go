package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceToReleaseListHOCmd */
/* PDUSessionResourceToReleaseItemHOCmd */
type PDUSessionResourceToReleaseListHOCmd struct {
	List []PDUSessionResourceToReleaseItemHOCmd `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
