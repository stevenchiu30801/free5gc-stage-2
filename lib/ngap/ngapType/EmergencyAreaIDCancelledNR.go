package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct EmergencyAreaIDCancelledNR */
/* EmergencyAreaIDCancelledNRItem */
type EmergencyAreaIDCancelledNR struct {
	List []EmergencyAreaIDCancelledNRItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
