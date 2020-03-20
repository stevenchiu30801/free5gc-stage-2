package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type GBRQosInformation struct {
	MaximumFlowBitRateDL    BitRate
	MaximumFlowBitRateUL    BitRate
	GuaranteedFlowBitRateDL BitRate
	GuaranteedFlowBitRateUL BitRate
	NotificationControl     *NotificationControl                               `aper:"optional"`
	MaximumPacketLossRateDL *PacketLossRate                                    `aper:"optional"`
	MaximumPacketLossRateUL *PacketLossRate                                    `aper:"optional"`
	IEExtensions            *ProtocolExtensionContainerGBRQosInformationExtIEs `aper:"optional"`
}
