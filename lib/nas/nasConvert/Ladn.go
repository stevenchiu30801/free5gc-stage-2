package nasConvert

import (
	"gofree5gc/src/amf/amf_context"
)

func LadnToModels(buf []uint8) (dnnValues []string) {

	for bufOffset := 1; bufOffset < len(buf); {
		lenOfDnn := int(buf[bufOffset])
		dnn := string(buf[bufOffset : bufOffset+lenOfDnn])
		dnnValues = append(dnnValues, dnn)
		bufOffset += lenOfDnn
	}

	return
}

func LadnToNas(ladn amf_context.LADN) (ladnNas []uint8) {

	dnnNas := []byte(ladn.Ladn)

	ladnNas = append(ladnNas, uint8(len(dnnNas)))
	ladnNas = append(ladnNas, dnnNas...)

	taiListNas := TaiListToNas(ladn.TaiLists)
	ladnNas = append(ladnNas, uint8(len(taiListNas)))
	ladnNas = append(ladnNas, taiListNas...)
	return
}
