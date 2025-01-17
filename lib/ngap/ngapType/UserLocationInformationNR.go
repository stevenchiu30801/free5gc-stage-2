package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type UserLocationInformationNR struct {
	NRCGI        NRCGI                                                      `aper:"valueExt"`
	TAI          TAI                                                        `aper:"valueExt"`
	TimeStamp    *TimeStamp                                                 `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerUserLocationInformationNRExtIEs `aper:"optional"`
}
