package ngapConvert

import (
	"encoding/hex"
	"gofree5gc/lib/ngap/ngapType"
	"gofree5gc/lib/openapi/models"
)

func TaiToModels(tai ngapType.TAI) (modelsTai models.Tai) {
	plmnID := PlmnIdToModels(tai.PLMNIdentity)
	modelsTai.PlmnId = &plmnID
	modelsTai.Tac = hex.EncodeToString(tai.TAC.Value)
	return
}

func TaiToNgap(tai models.Tai) (ngapTai ngapType.TAI) {
	ngapTai.PLMNIdentity = PlmnIdToNgap(*tai.PlmnId)
	tac, _ := hex.DecodeString(tai.Tac)
	ngapTai.TAC.Value = tac
	return
}
