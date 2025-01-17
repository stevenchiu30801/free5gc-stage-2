package ngapConvert

import (
	"gofree5gc/lib/aper"
	"gofree5gc/lib/ngap/ngapType"
	"gofree5gc/lib/openapi/models"
)

// TS 38.413 9.3.1.85
func RATRestrictionInformationToNgap(ratType models.RatType) (ratResInfo ngapType.RATRestrictionInformation) {

	// Only support EUTRA & NR in version15.2.0
	switch ratType {
	case models.RatType_EUTRA:
		ratResInfo.Value = aper.BitString{
			Bytes:     []byte{0x80},
			BitLength: 8,
		}
	case models.RatType_NR:
		ratResInfo.Value = aper.BitString{
			Bytes:     []byte{0x40},
			BitLength: 8,
		}
	}
	return
}
