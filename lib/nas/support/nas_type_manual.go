// ExtendedProtocolDiscriminator 9.2
// ExtendedProtocolDiscriminator Row, sBit, len = [0, 0], 8 , 8
type ExtendedProtocolDiscriminator struct {
	Octet uint8
}

// SpareHalfOctetAndSecurityHeaderType 9.3 9.5
// SpareHalfOctet Row, sBit, len = [0, 0], 8 , 4
// SecurityHeaderType Row, sBit, len = [0, 0], 4 , 4
type SpareHalfOctetAndSecurityHeaderType struct {
	Octet uint8
}

// AuthenticationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationRequestMessageIdentity struct {
	Octet uint8
}

// SpareHalfOctetAndNgksi 9.11.3.32 9.5
// SpareHalfOctet Row, sBit, len = [0, 0], 8 , 4
// TSC Row, sBit, len = [0, 0], 4 , 1
// NasKeySetIdentifiler Row, sBit, len = [0, 0], 3 , 3
type SpareHalfOctetAndNgksi struct {
	Octet uint8
}

// ABBA 9.11.3.10
// ABBAContents Row, sBit, len = [0, 0], 8 , INF
type ABBA struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// AuthenticationParameterRAND 9.11.3.16
// RANDValue Row, sBit, len = [0, 15], 8 , 128
type AuthenticationParameterRAND struct {
	Iei   uint8
	Octet [16]uint8
}

// AuthenticationParameterAUTN 9.11.3.15
// AUTN Row, sBit, len = [0, 15], 8 , 128
type AuthenticationParameterAUTN struct {
	Iei   uint8
	Len   uint8
	Octet [16]uint8
}

// EAPMessage 9.11.2.2
// EAPMessage Row, sBit, len = [0, 0], 8 , INF
type EAPMessage struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// AuthenticationResponseMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationResponseMessageIdentity struct {
	Octet uint8
}

// AuthenticationResponseParameter 9.11.3.17
// RES Row, sBit, len = [0, 15], 8 , 128
type AuthenticationResponseParameter struct {
	Iei   uint8
	Len   uint8
	Octet [16]uint8
}

// AuthenticationResultMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationResultMessageIdentity struct {
	Octet uint8
}

// AuthenticationFailureMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationFailureMessageIdentity struct {
	Octet uint8
}

// Cause5GMM 9.11.3.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
type Cause5GMM struct {
	Iei   uint8
	Octet uint8
}

// AuthenticationFailureParameter 9.11.3.14
// AuthenticationFailureParameter Row, sBit, len = [0, 13], 8 , 112
type AuthenticationFailureParameter struct {
	Iei   uint8
	Len   uint8
	Octet [14]uint8
}

// AuthenticationRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationRejectMessageIdentity struct {
	Octet uint8
}

// RegistrationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RegistrationRequestMessageIdentity struct {
	Octet uint8
}

// NgksiAndRegistrationType5GS 9.11.3.7 9.11.3.32
// TSC Row, sBit, len = [0, 0], 8 , 1
// NasKeySetIdentifiler Row, sBit, len = [0, 0], 7 , 3
// FOR  Row, sBit, len = [0, 0], 4 , 1
// RegistrationType5GS Row, sBit, len = [0, 0], 3 , 3
type NgksiAndRegistrationType5GS struct {
	Octet uint8
}

// MobileIdentity5GS 9.11.3.4
// MobileIdentity5GSContents Row, sBit, len = [0, 0], 8 , INF
type MobileIdentity5GS struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// NoncurrentNativeNASKeySetIdentifier 9.11.3.32
// Iei  Row, sBit, len = [0, 0], 8 , 4
// Tsc  Row, sBit, len = [0, 0], 4 , 1
// NasKeySetIdentifiler  Row, sBit, len = [0, 0], 3 , 3
type NoncurrentNativeNASKeySetIdentifier struct {
	Octet uint8
}

// Capability5GMM 9.11.3.1
// LPP  Row, sBit, len = [0, 0], 3 , 1
// HOAttach Row, sBit, len = [0, 0], 2 , 1
// S1Mode Row, sBit, len = [0, 0], 1 , 1
// Spare Row, sBit, len = [1, 12], 8 , 96
type Capability5GMM struct {
	Iei   uint8
	Len   uint8
	Octet [13]uint8
}

// UESecurityCapability 9.11.3.54
// EA0_5G Row, sBit, len = [0, 0], 8 , 1
// EA1_128_5G Row, sBit, len = [0, 0], 7 , 1
// EA2_128_5G Row, sBit, len = [0, 0], 6 , 1
// EA3_128_5G Row, sBit, len = [0, 0], 5 , 1
// EA4_5G Row, sBit, len = [0, 0], 4 , 1
// EA5_5G Row, sBit, len = [0, 0], 3 , 1
// EA6_5G Row, sBit, len = [0, 0], 2 , 1
// EA7_5G Row, sBit, len = [0, 0], 1 , 1
// IA0_5G Row, sBit, len = [1, 1], 8 , 1
// IA1_128_5G Row, sBit, len = [1, 1], 7 , 1
// IA2_128_5G Row, sBit, len = [1, 1], 6 , 1
// IA3_128_5G Row, sBit, len = [1, 1], 5 , 1
// IA4_5G Row, sBit, len = [1, 1], 4 , 1
// IA5_5G Row, sBit, len = [1, 1], 3 , 1
// IA6_5G Row, sBit, len = [1, 1], 2 , 1
// IA7_5G Row, sBit, len = [1, 1], 1 , 1
// EEA0 Row, sBit, len = [2, 2], 8 , 1
// EEA1_128 Row, sBit, len = [2, 2], 7 , 1
// EEA2_128 Row, sBit, len = [2, 2], 6 , 1
// EEA3_128 Row, sBit, len = [2, 2], 5 , 1
// EEA4 Row, sBit, len = [2, 2], 4 , 1
// EEA5 Row, sBit, len = [2, 2], 3 , 1
// EEA6 Row, sBit, len = [2, 2], 2 , 1
// EEA7 Row, sBit, len = [2, 2], 1 , 1
// EIA0 Row, sBit, len = [3, 3], 8 , 1
// EIA1_128 Row, sBit, len = [3, 3], 7 , 1
// EIA2_128 Row, sBit, len = [3, 3], 6 , 1
// EIA3_128 Row, sBit, len = [3, 3], 5 , 1
// EIA4 Row, sBit, len = [3, 3], 4 , 1
// EIA5 Row, sBit, len = [3, 3], 3 , 1
// EIA6 Row, sBit, len = [3, 3], 2 , 1
// EIA7 Row, sBit, len = [3, 3], 1 , 1
// Spare Row, sBit, len = [4, 7], 8 , 32
type UESecurityCapability struct {
	Iei   uint8
	Len   uint8
	Buffer []uint8
}

// RequestedNSSAI 9.11.3.37
// SNSSAIValue Row, sBit, len = [0, 0], 0 , INF
type RequestedNSSAI struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// LastVisitedRegisteredTAI 9.11.3.8
// MCCDigit2 Row, sBit, len = [0, 0], 8 , 4
// MCCDigit1 Row, sBit, len = [0, 0], 4 , 4
// MNCDigit3 Row, sBit, len = [1, 1], 8 , 4
// MCCDigit3 Row, sBit, len = [1, 1], 4 , 4
// MNCDigit2 Row, sBit, len = [2, 2], 8 , 4
// MNCDigit1 Row, sBit, len = [2, 2], 4 , 4
// TAC Row, sBit, len = [3, 5], 8 , 24
type LastVisitedRegisteredTAI struct {
	Iei   uint8
	Octet [7]uint8
}

// S1UENetworkCapability 9.11.3.48
// EEA0 Row, sBit, len = [0, 0], 8 , 1
// EEA1_128 Row, sBit, len = [0, 0], 7 , 1
// EEA2_128 Row, sBit, len = [0, 0], 6 , 1
// EEA3_128 Row, sBit, len = [0, 0], 5 , 1
// EEA4 Row, sBit, len = [0, 0], 4 , 1
// EEA5 Row, sBit, len = [0, 0], 3 , 1
// EEA6 Row, sBit, len = [0, 0], 2 , 1
// EEA7 Row, sBit, len = [0, 0], 1 , 1
// EIA0 Row, sBit, len = [1, 1], 8 , 1
// EIA1_128 Row, sBit, len = [1, 1], 7 , 1
// EIA2_128 Row, sBit, len = [1, 1], 6 , 1
// EIA3_128 Row, sBit, len = [1, 1], 5 , 1
// EIA4 Row, sBit, len = [1, 1], 4 , 1
// EIA5 Row, sBit, len = [1, 1], 3 , 1
// EIA6 Row, sBit, len = [1, 1], 2 , 1
// EIA7 Row, sBit, len = [1, 1], 1 , 1
// UEA0 Row, sBit, len = [2, 2], 8 , 1
// UEA1 Row, sBit, len = [2, 2], 7 , 1
// UEA2 Row, sBit, len = [2, 2], 6 , 1
// UEA3 Row, sBit, len = [2, 2], 5 , 1
// UEA4 Row, sBit, len = [2, 2], 4 , 1
// UEA5 Row, sBit, len = [2, 2], 3 , 1
// UEA6 Row, sBit, len = [2, 2], 2 , 1
// UEA7 Row, sBit, len = [2, 2], 1 , 1
// UCS2 Row, sBit, len = [3, 3], 8 , 1
// UIA1 Row, sBit, len = [3, 3], 7 , 1
// UIA2 Row, sBit, len = [3, 3], 6 , 1
// UIA3 Row, sBit, len = [3, 3], 5 , 1
// UIA4 Row, sBit, len = [3, 3], 4 , 1
// UIA5 Row, sBit, len = [3, 3], 3 , 1
// UIA6 Row, sBit, len = [3, 3], 2 , 1
// UIA7 Row, sBit, len = [3, 3], 1 , 1
// ProSedd Row, sBit, len = [4, 4], 8 , 1
// ProSe Row, sBit, len = [4, 4], 7 , 1
// H245ASH Row, sBit, len = [4, 4], 6 , 1
// ACCCSFB Row, sBit, len = [4, 4], 5 , 1
// LPP Row, sBit, len = [4, 4], 4 , 1
// LCS Row, sBit, len = [4, 4], 3 , 1
// xSRVCC Row, sBit, len = [4, 4], 2 , 1
// NF Row, sBit, len = [4, 4], 1 , 1
// EPCO Row, sBit, len = [5, 5], 8 , 1
// HCCPCIOT Row, sBit, len = [5, 5], 7 , 1
// ERwoPDN Row, sBit, len = [5, 5], 6 , 1
// S1UData Row, sBit, len = [5, 5], 5 , 1
// UPCIot Row, sBit, len = [5, 5], 4 , 1
// CPCIot Row, sBit, len = [5, 5], 3 , 1
// Proserelay Row, sBit, len = [5, 5], 2 , 1
// ProSedc Row, sBit, len = [5, 5], 1 , 1
// Bearer15 Row, sBit, len = [6, 6], 8 , 1
// SGC Row, sBit, len = [6, 6], 7 , 1
// N1mode Row, sBit, len = [6, 6], 6 , 1
// DCNR Row, sBit, len = [6, 6], 5 , 1
// CPbackoff Row, sBit, len = [6, 6], 4 , 1
// RestrictEC Row, sBit, len = [6, 6], 3 , 1
// V2XPC5 Row, sBit, len = [6, 6], 2 , 1
// MulitpeDRB Row, sBit, len = [6, 6], 1 , 1
// Spare Row, sBit, len = [7, 12], 8 , INF
type S1UENetworkCapability struct {
	Iei   uint8
	Len   uint8
	Buffer []uint8
}

// UplinkDataStatus 9.11.3.57
// PSI7 Row, sBit, len = [0, 0], 8 , 1
// PSI6 Row, sBit, len = [0, 0], 7 , 1
// PSI5 Row, sBit, len = [0, 0], 6 , 1
// PSI4 Row, sBit, len = [0, 0], 5 , 1
// PSI3 Row, sBit, len = [0, 0], 4 , 1
// PSI2 Row, sBit, len = [0, 0], 3 , 1
// PSI1 Row, sBit, len = [0, 0], 2 , 1
// PSI0 Row, sBit, len = [0, 0], 1 , 1
// PSI15 Row, sBit, len = [1, 1], 8 , 1
// PSI14 Row, sBit, len = [1, 1], 7 , 1
// PSI13 Row, sBit, len = [1, 1], 6 , 1
// PSI12 Row, sBit, len = [1, 1], 5 , 1
// PSI11 Row, sBit, len = [1, 1], 4 , 1
// PSI10 Row, sBit, len = [1, 1], 3 , 1
// PSI9 Row, sBit, len = [1, 1], 2 , 1
// PSI8 Row, sBit, len = [1, 1], 1 , 1
// Spare Row, sBit, len = [2, 2], 1 , INF
type UplinkDataStatus struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// PDUSessionStatus 9.11.3.44
// PSI7 Row, sBit, len = [0, 0], 8 , 1
// PSI6 Row, sBit, len = [0, 0], 7 , 1
// PSI5 Row, sBit, len = [0, 0], 6 , 1
// PSI4 Row, sBit, len = [0, 0], 5 , 1
// PSI3 Row, sBit, len = [0, 0], 4 , 1
// PSI2 Row, sBit, len = [0, 0], 3 , 1
// PSI1 Row, sBit, len = [0, 0], 2 , 1
// PSI0 Row, sBit, len = [0, 0], 1 , 1
// PSI15 Row, sBit, len = [1, 1], 8 , 1
// PSI14 Row, sBit, len = [1, 1], 7 , 1
// PSI13 Row, sBit, len = [1, 1], 6 , 1
// PSI12 Row, sBit, len = [1, 1], 5 , 1
// PSI11 Row, sBit, len = [1, 1], 4 , 1
// PSI10 Row, sBit, len = [1, 1], 3 , 1
// PSI9 Row, sBit, len = [1, 1], 2 , 1
// PSI8 Row, sBit, len = [1, 1], 1 , 1
// Spare Row, sBit, len = [2, 2], 1 , INF
type PDUSessionStatus struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// MICOIndication 9.11.3.31
// Iei Row, sBit, len = [0, 0], 8 , 4
// RAAI Row, sBit, len = [0, 0], 1 , 1
type MICOIndication struct {
	Octet uint8
}

// UEStatus 9.11.3.56
// N1ModeReg Row, sBit, len = [0, 0], 2 , 1
// S1ModeReg Row, sBit, len = [0, 0], 1 , 1
type UEStatus struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// AdditionalGUTI 9.11.3.4
// Spare Row, sBit, len = [0, 0], 4 , 1
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
// MCCDigit2 Row, sBit, len = [1, 1], 8 , 4
// MCCDigit1 Row, sBit, len = [1, 1], 4 , 4
// MNCDigit3 Row, sBit, len = [2, 2], 8 , 4
// MCCDigit3 Row, sBit, len = [2, 2], 4 , 4
// MNCDigit2 Row, sBit, len = [3, 3], 8 , 4
// MNCDigit1 Row, sBit, len = [3, 3], 4 , 4
// AMFRegionID Row, sBit, len = [4, 4], 8 , 8
// AMFSetID Row, sBit, len = [5, 6], 8 , 10
// AMFPointer Row, sBit, len = [6, 6], 6 , 6
// TMSI5G Row, sBit, len = [7, 10], 8 , 32
type AdditionalGUTI struct {
	Iei   uint8
	Len   uint16
	Octet [11]uint8
}

// AllowedPDUSessionStatus 9.11.3.13
// PSI7 Row, sBit, len = [0, 0], 8 , 1
// PSI6 Row, sBit, len = [0, 0], 7 , 1
// PSI5 Row, sBit, len = [0, 0], 6 , 1
// PSI4 Row, sBit, len = [0, 0], 5 , 1
// PSI3 Row, sBit, len = [0, 0], 4 , 1
// PSI2 Row, sBit, len = [0, 0], 3 , 1
// PSI1 Row, sBit, len = [0, 0], 2 , 1
// PSI0 Row, sBit, len = [0, 0], 1 , 1
// PSI15 Row, sBit, len = [1, 1], 8 , 1
// PSI14 Row, sBit, len = [1, 1], 7 , 1
// PSI13 Row, sBit, len = [1, 1], 6 , 1
// PSI12 Row, sBit, len = [1, 1], 5 , 1
// PSI11 Row, sBit, len = [1, 1], 4 , 1
// PSI10 Row, sBit, len = [1, 1], 3 , 1
// PSI9 Row, sBit, len = [1, 1], 2 , 1
// PSI8 Row, sBit, len = [1, 1], 1 , 1
// Spare Row, sBit, len = [2, 2], 1 , INF
type AllowedPDUSessionStatus struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// UesUsageSetting 9.11.3.55
// UesUsageSetting Row, sBit, len = [0, 0], 1 , 1
type UesUsageSetting struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// RequestedDRXParameters 9.11.3.2A
// DRXValue Row, sBit, len = [0, 0], 4 , 4
type RequestedDRXParameters struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// EPSNASMessageContainer 9.11.3.24
// EPANASMessageContainer Row, sBit, len = [0, 0], 8 , INF
type EPSNASMessageContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// LADNIndication 9.11.3.29
// LADNDNNValue Row, sBit, len = [0, 0], 8 , INF
type LADNIndication struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// PayloadContainer 9.11.3.39
// PayloadContainerContents Row, sBit, len = [0, 0], 8 , INF
type PayloadContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// NetworkSlicingIndication 9.11.3.36
// Iei Row, sBit, len = [0, 0], 8 , 4
// DCNI Row, sBit, len = [0, 0], 2 , 1
// NSSCI Row, sBit, len = [0, 0], 1 , 1
type NetworkSlicingIndication struct {
	Octet uint8
}

// UpdateType5GS 9.11.3.9A
// NGRanRcu Row, sBit, len = [0, 0], 2 , 1
// SMSRequested Row, sBit, len = [0, 0], 1 , 1
type UpdateType5GS struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// NASMessageContainer 9.11.3.33
// NASMessageContainerContents Row, sBit, len = [0, 0], 8 , INF
type NASMessageContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// RegistrationAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RegistrationAcceptMessageIdentity struct {
	Octet uint8
}

// RegistrationResult5GS 9.11.3.6
// SMSAllowed Row, sBit, len = [0, 0], 4 , 1
// RegistrationResultValue5GS Row, sBit, len = [0, 0], 3 , 3
type RegistrationResult5GS struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// GUTI5G 9.11.3.4
// Spare Row, sBit, len = [0, 0], 4 , 1
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
// MCCDigit2 Row, sBit, len = [1, 1], 8 , 4
// MCCDigit1 Row, sBit, len = [1, 1], 4 , 4
// MNCDigit3 Row, sBit, len = [2, 2], 8 , 4
// MCCDigit3 Row, sBit, len = [2, 2], 4 , 4
// MNCDigit2 Row, sBit, len = [3, 3], 8 , 4
// MNCDigit1 Row, sBit, len = [3, 3], 4 , 4
// AMFRegionID Row, sBit, len = [4, 4], 8 , 8
// AMFSetID Row, sBit, len = [5, 6], 8 , 10
// AMFPointer Row, sBit, len = [6, 6], 6 , 6
// TMSI5G Row, sBit, len = [7, 10], 8 , 32
type GUTI5G struct {
	Iei   uint8
	Len   uint16
	Octet [11]uint8
}

// EquivalentPlmns 9.11.3.45
// MCCDigit2PLMN1 Row, sBit, len = [0, 0], 8 , 4
// MCCDigit1PLMN1 Row, sBit, len = [0, 0], 4 , 4
// MNCDigit3PLMN1 Row, sBit, len = [1, 1], 8 , 4
// MCCDigit3PLMN1 Row, sBit, len = [1, 1], 4 , 4
// MNCDigit2PLMN1 Row, sBit, len = [2, 2], 8 , 4
// MNCDigit1PLMN1 Row, sBit, len = [2, 2], 4 , 4
// MCCDigit2PLMN2 Row, sBit, len = [3, 3], 8 , 4
// MCCDigit1PLMN2 Row, sBit, len = [3, 3], 4 , 4
// MNCDigit3PLMN2 Row, sBit, len = [4, 4], 8 , 4
// MCCDigit3PLMN2 Row, sBit, len = [4, 4], 4 , 4
// MNCDigit2PLMN2 Row, sBit, len = [5, 5], 8 , 4
// MNCDigit1PLMN2 Row, sBit, len = [5, 5], 4 , 4
// MCCDigit2PLMN3 Row, sBit, len = [6, 6], 8 , 4
// MCCDigit1PLMN3 Row, sBit, len = [6, 6], 4 , 4
// MNCDigit3PLMN3 Row, sBit, len = [7, 7], 8 , 4
// MCCDigit3PLMN3 Row, sBit, len = [7, 7], 4 , 4
// MNCDigit2PLMN3 Row, sBit, len = [8, 8], 8 , 4
// MNCDigit1PLMN3 Row, sBit, len = [8, 8], 4 , 4
// MCCDigit2PLMN4 Row, sBit, len = [9, 9], 8 , 4
// MCCDigit1PLMN4 Row, sBit, len = [9, 9], 4 , 4
// MNCDigit3PLMN4 Row, sBit, len = [10, 10], 8 , 4
// MCCDigit3PLMN4 Row, sBit, len = [10, 10], 4 , 4
// MNCDigit2PLMN4 Row, sBit, len = [11, 11], 8 , 4
// MNCDigit1PLMN4 Row, sBit, len = [11, 11], 4 , 4
// MCCDigit2PLMN5 Row, sBit, len = [12, 12], 8 , 4
// MCCDigit1PLMN5 Row, sBit, len = [12, 12], 4 , 4
// MNCDigit3PLMN5 Row, sBit, len = [13, 13], 8 , 4
// MCCDigit3PLMN5 Row, sBit, len = [13, 13], 4 , 4
// MNCDigit2PLMN5 Row, sBit, len = [14, 14], 8 , 4
// MNCDigit1PLMN5 Row, sBit, len = [14, 14], 4 , 4
// MCCDigit2PLMN6 Row, sBit, len = [15, 15], 8 , 4
// MCCDigit1PLMN6 Row, sBit, len = [15, 15], 4 , 4
// MNCDigit3PLMN6 Row, sBit, len = [16, 16], 8 , 4
// MCCDigit3PLMN6 Row, sBit, len = [16, 16], 4 , 4
// MNCDigit2PLMN6 Row, sBit, len = [17, 17], 8 , 4
// MNCDigit1PLMN6 Row, sBit, len = [17, 17], 4 , 4
// MCCDigit2PLMN7 Row, sBit, len = [18, 18], 8 , 4
// MCCDigit1PLMN7 Row, sBit, len = [18, 18], 4 , 4
// MNCDigit3PLMN7 Row, sBit, len = [19, 19], 8 , 4
// MCCDigit3PLMN7 Row, sBit, len = [19, 19], 4 , 4
// MNCDigit2PLMN7 Row, sBit, len = [20, 20], 8 , 4
// MNCDigit1PLMN7 Row, sBit, len = [20, 20], 4 , 4
// MCCDigit2PLMN8 Row, sBit, len = [21, 21], 8 , 4
// MCCDigit1PLMN8 Row, sBit, len = [21, 21], 4 , 4
// MNCDigit3PLMN8 Row, sBit, len = [22, 22], 8 , 4
// MCCDigit3PLMN8 Row, sBit, len = [22, 22], 4 , 4
// MNCDigit2PLMN8 Row, sBit, len = [23, 23], 8 , 4
// MNCDigit1PLMN8 Row, sBit, len = [23, 23], 4 , 4
// MCCDigit2PLMN9 Row, sBit, len = [24, 24], 8 , 4
// MCCDigit1PLMN9 Row, sBit, len = [24, 24], 4 , 4
// MNCDigit3PLMN9 Row, sBit, len = [25, 25], 8 , 4
// MCCDigit3PLMN9 Row, sBit, len = [25, 25], 4 , 4
// MNCDigit2PLMN9 Row, sBit, len = [26, 26], 8 , 4
// MNCDigit1PLMN9 Row, sBit, len = [26, 26], 4 , 4
// MCCDigit2PLMN10 Row, sBit, len = [27, 27], 8 , 4
// MCCDigit1PLMN10 Row, sBit, len = [27, 27], 4 , 4
// MNCDigit3PLMN10 Row, sBit, len = [28, 28], 8 , 4
// MCCDigit3PLMN10 Row, sBit, len = [28, 28], 4 , 4
// MNCDigit2PLMN10 Row, sBit, len = [29, 29], 8 , 4
// MNCDigit1PLMN10 Row, sBit, len = [29, 29], 4 , 4
// MCCDigit2PLMN11 Row, sBit, len = [30, 30], 8 , 4
// MCCDigit1PLMN11 Row, sBit, len = [30, 30], 4 , 4
// MNCDigit3PLMN11 Row, sBit, len = [31, 31], 8 , 4
// MCCDigit3PLMN11 Row, sBit, len = [31, 31], 4 , 4
// MNCDigit2PLMN11 Row, sBit, len = [32, 32], 8 , 4
// MNCDigit1PLMN11 Row, sBit, len = [32, 32], 4 , 4
// MCCDigit2PLMN12 Row, sBit, len = [33, 33], 8 , 4
// MCCDigit1PLMN12 Row, sBit, len = [33, 33], 4 , 4
// MNCDigit3PLMN12 Row, sBit, len = [34, 34], 8 , 4
// MCCDigit3PLMN12 Row, sBit, len = [34, 34], 4 , 4
// MNCDigit2PLMN12 Row, sBit, len = [35, 35], 8 , 4
// MNCDigit1PLMN12 Row, sBit, len = [35, 35], 4 , 4
// MCCDigit2PLMN13 Row, sBit, len = [36, 36], 8 , 4
// MCCDigit1PLMN13 Row, sBit, len = [36, 36], 4 , 4
// MNCDigit3PLMN13 Row, sBit, len = [37, 37], 8 , 4
// MCCDigit3PLMN13 Row, sBit, len = [37, 37], 4 , 4
// MNCDigit2PLMN13 Row, sBit, len = [38, 38], 8 , 4
// MNCDigit1PLMN13 Row, sBit, len = [38, 38], 4 , 4
// MCCDigit2PLMN14 Row, sBit, len = [39, 39], 8 , 4
// MCCDigit1PLMN14 Row, sBit, len = [39, 39], 4 , 4
// MNCDigit3PLMN14 Row, sBit, len = [40, 40], 8 , 4
// MCCDigit3PLMN14 Row, sBit, len = [40, 40], 4 , 4
// MNCDigit2PLMN14 Row, sBit, len = [41, 41], 8 , 4
// MNCDigit1PLMN14 Row, sBit, len = [41, 41], 4 , 4
// MCCDigit2PLMN15 Row, sBit, len = [42, 42], 8 , 4
// MCCDigit1PLMN15 Row, sBit, len = [42, 42], 4 , 4
// MNCDigit3PLMN15 Row, sBit, len = [43, 43], 8 , 4
// MCCDigit3PLMN15 Row, sBit, len = [43, 43], 4 , 4
// MNCDigit2PLMN15 Row, sBit, len = [44, 44], 8 , 4
// MNCDigit1PLMN15 Row, sBit, len = [44, 44], 4 , 4
type EquivalentPlmns struct {
	Iei   uint8
	Len   uint8
	Octet [45]uint8
}

// TAIList 9.11.3.9
// PartialTrackingAreaIdentityList Row, sBit, len = [0, 0], 8 , INF
type TAIList struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// AllowedNSSAI 9.11.3.37
// SNSSAIValue Row, sBit, len = [0, 0], 0 , INF
type AllowedNSSAI struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// RejectedNSSAI 9.11.3.46
// RejectedNSSAIContents Row, sBit, len = [0, 0], 0 , INF
type RejectedNSSAI struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// ConfiguredNSSAI 9.11.3.37
// SNSSAIValue Row, sBit, len = [0, 0], 0 , INF
type ConfiguredNSSAI struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// NetworkFeatureSupport5GS 9.11.3.5
// MPSI Row, sBit, len = [0, 0], 8 , 1
// IWKN26 Row, sBit, len = [0, 0], 7 , 1
// EMF Row, sBit, len = [0, 0], 6 , 2
// EMC Row, sBit, len = [0, 0], 4 , 2
// IMSVoPSN3GPP Row, sBit, len = [0, 0], 2 , 1
// IMSVoPS3GPP Row, sBit, len = [0, 0], 1 , 1
// MCSI Row, sBit, len = [1, 1], 2 , 1
// EMCN Row, sBit, len = [1, 1], 1 , 1
// Spare Row, sBit, len = [2, 2], 8 , 8
type NetworkFeatureSupport5GS struct {
	Iei   uint8
	Len   uint8
	Octet [3]uint8
}

// PDUSessionReactivationResult 9.11.3.42
// PSI7 Row, sBit, len = [0, 0], 8 , 1
// PSI6 Row, sBit, len = [0, 0], 7 , 1
// PSI5 Row, sBit, len = [0, 0], 6 , 1
// PSI4 Row, sBit, len = [0, 0], 5 , 1
// PSI3 Row, sBit, len = [0, 0], 4 , 1
// PSI2 Row, sBit, len = [0, 0], 3 , 1
// PSI1 Row, sBit, len = [0, 0], 2 , 1
// PSI0 Row, sBit, len = [0, 0], 1 , 1
// PSI15 Row, sBit, len = [1, 1], 8 , 1
// PSI14 Row, sBit, len = [1, 1], 7 , 1
// PSI13 Row, sBit, len = [1, 1], 6 , 1
// PSI12 Row, sBit, len = [1, 1], 5 , 1
// PSI11 Row, sBit, len = [1, 1], 4 , 1
// PSI10 Row, sBit, len = [1, 1], 3 , 1
// PSI9 Row, sBit, len = [1, 1], 2 , 1
// PSI8 Row, sBit, len = [1, 1], 1 , 1
// Spare Row, sBit, len = [2, 2], 1 , INF
type PDUSessionReactivationResult struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// PDUSessionReactivationResultErrorCause 9.11.3.43
// PDUSessionIDAndCauseValue Row, sBit, len = [0, 0], 8 , INF
type PDUSessionReactivationResultErrorCause struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// LADNInformation 9.11.3.30
// LADND Row, sBit, len = [0, 0], 8 , INF
type LADNInformation struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// ServiceAreaList 9.11.3.49
// PartialServiceAreaList Row, sBit, len = [0, 0], 8 , INF
type ServiceAreaList struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// T3512Value 9.11.2.5
// Unit Row, sBit, len = [0, 0], 8 , 3
// TimerValue Row, sBit, len = [0, 0], 5 , 5
type T3512Value struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// Non3GppDeregistrationTimerValue 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
type Non3GppDeregistrationTimerValue struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// T3502Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
type T3502Value struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// EmergencyNumberList 9.11.3.23
// Lengthof1EmergencyNumberInformation Row, sBit, len = [0, 0], 8 , 8
// EmergencyServiceCategoryValue Row, sBit, len = [1, 1], 5 , 5
// EmergencyInformation Row, sBit, len = [0, 0], 8 , INF
type EmergencyNumberList struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// ExtendedEmergencyNumberList 9.11.3.26
// EENL Row, sBit, len = [0, 0], 1 , 1
// EmergencyInformation Row, sBit, len = [0, 0], 8 , INF
type ExtendedEmergencyNumberList struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// SORTransparentContainer 9.11.3.51
// SORContent Row, sBit, len = [0, 0], 8 , INF
type SORTransparentContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// NSSAIInclusionMode 9.11.3.37A
// Iei Row, sBit, len = [0, 0], 8 , 4
// NSSAIInclusionMode Row, sBit, len = [0, 0], 2 , 2
type NSSAIInclusionMode struct {
	Octet uint8
}

// OperatordefinedAccessCategoryDefinitions 9.11.3.38
// OperatorDefinedAccessCategoryDefintiion Row, sBit, len = [0, 0], 8 , INF
type OperatordefinedAccessCategoryDefinitions struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// NegotiatedDRXParameters 9.11.3.2A
// DRXValue Row, sBit, len = [0, 0], 4 , 4
type NegotiatedDRXParameters struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// RegistrationCompleteMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RegistrationCompleteMessageIdentity struct {
	Octet uint8
}

// RegistrationRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RegistrationRejectMessageIdentity struct {
	Octet uint8
}

// T3346Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
type T3346Value struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// ULNASTRANSPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ULNASTRANSPORTMessageIdentity struct {
	Octet uint8
}

// SpareHalfOctetAndPayloadContainerType 9.11.3.40 9.5
// PayloadContainerType Row, sBit, len = [0, 0], 4 , 4
type SpareHalfOctetAndPayloadContainerType struct {
	Octet uint8
}

// PduSessionID2Value 9.11.3.41
// PduSessionID2Value Row, sBit, len = [0, 0], 8 , 8
type PduSessionID2Value struct {
	Iei   uint8
	Octet uint8
}

// OldPDUSessionID 9.11.3.41
// OldPDUSessionID Row, sBit, len = [0, 0], 8 , 8
type OldPDUSessionID struct {
	Iei   uint8
	Octet uint8
}

// RequestType 9.11.3.47
// Iei Row, sBit, len = [0, 0], 8 , 4
// RequestTypeValue Row, sBit, len = [0, 0], 3 , 3
type RequestType struct {
	Octet uint8
}

// SNSSAI 9.11.2.8
// SST Row, sBit, len = [0, 0], 8 , 8
// SD Row, sBit, len = [1, 3], 8 , 24
// MappedHPLMNSST Row, sBit, len = [4, 4], 8 , 8
// MappedHPLMNSD Row, sBit, len = [5, 7], 8 , 24
type SNSSAI struct {
	Iei   uint8
	Len   uint8
	Octet [8]uint8
}

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
type DNN struct {
Iei    uint8
Len    uint8
Buffer []uint8
}

// AdditionalInformation 9.11.2.1
// AdditionalInformationValue Row, sBit, len = [0, 0], 8 , INF
type AdditionalInformation struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// DLNASTRANSPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type DLNASTRANSPORTMessageIdentity struct {
	Octet uint8
}

// BackoffTimerValue 9.11.2.5
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
// TimerValue Row, sBit, len = [0, 0], 5 , 5
type BackoffTimerValue struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// DeregistrationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type DeregistrationRequestMessageIdentity struct {
	Octet uint8
}

// NgksiAndDeregistrationType 9.11.3.20 9.11.3.32
// TSC Row, sBit, len = [0, 0], 8 , 1
// NasKeySetIdentifiler Row, sBit, len = [0, 0], 7 , 3
// SwitchOff Row, sBit, len = [0, 0], 4 , 1
// ReRegistrationRequired Row, sBit, len = [0, 0], 3 , 1
// AccessType Row, sBit, len = [0, 0], 2 , 2
type NgksiAndDeregistrationType struct {
	Octet uint8
}

// DeregistrationAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type DeregistrationAcceptMessageIdentity struct {
	Octet uint8
}

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// SwitchOff Row, sBit, len = [0, 0], 4 , 1
// ReRegistrationRequired Row, sBit, len = [0, 0], 3 , 1
// AccessType Row, sBit, len = [0, 0], 2 , 2
type SpareHalfOctetAndDeregistrationType struct {
	Octet uint8
}

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ServiceRequestMessageIdentity struct {
	Octet uint8
}

// ServiceTypeAndNgksi 9.11.3.32 9.11.3.50
// ServiceTypeValue Row, sBit, len = [0, 0], 8 , 4
// TSC Row, sBit, len = [0, 0], 4 , 1
// NasKeySetIdentifiler Row, sBit, len = [0, 0], 3 , 3
type ServiceTypeAndNgksi struct {
	Octet uint8
}

// TMSI5GS 9.11.3.4
// Spare Row, sBit, len = [0, 0], 4 , 1
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
// AMFSetID Row, sBit, len = [1, 1], 8 , 8
// AMFSetID Row, sBit, len = [1, 2], 8 , 10
// AMFPointer Row, sBit, len = [2, 2], 6 , 6
// TMSI5G Row, sBit, len = [3, 6], 8 , 32
type TMSI5GS struct {
	Iei   uint8
	Len   uint16
	Octet [7]uint8
}

// ServiceAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ServiceAcceptMessageIdentity struct {
	Octet uint8
}

// ServiceRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ServiceRejectMessageIdentity struct {
	Octet uint8
}

// ConfigurationUpdateCommandMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ConfigurationUpdateCommandMessageIdentity struct {
	Octet uint8
}

// ConfigurationUpdateIndication 9.11.3.18
// Iei Row, sBit, len = [0, 0], 8 , 4
// RED Row, sBit, len = [0, 0], 2 , 1
// ACK Row, sBit, len = [0, 0], 1 , 1
type ConfigurationUpdateIndication struct {
	Octet uint8
}

// FullNameForNetwork 9.11.3.35
// Ext Row, sBit, len = [0, 0], 8 ,1
// CodingScheme Row, sBit, len = [0, 0], 7 , 3
// AddCI Row, sBit, len = [0, 0], 4 , 1
// NumberOfSpareBitsInLastOctet Row, sBit, len = [0, 0], 3 , 3
// TextString Row, sBit, len = [1, 1], 4 , INF
type FullNameForNetwork struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// ShortNameForNetwork 9.11.3.35
// Ext Row, sBit, len = [0, 0], 8 , 1
// CodingScheme Row, sBit, len = [0, 0], 7 , 3
// AddCI Row, sBit, len = [0, 0], 4 , 1
// NumberOfSpareBitsInLastOctet Row, sBit, len = [0, 0], 3 , 3
// TextString Row, sBit, len = [1, 1], 4 , INF
type ShortNameForNetwork struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// LocalTimeZone 9.11.3.52
// TimeZone Row, sBit, len = [0, 0], 8 , 8
type LocalTimeZone struct {
	Iei   uint8
	Octet uint8
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
// Year Row, sBit, len = [0, 0], 8 , 8
// Month Row, sBit, len = [1, 1], 8 , 8
// Day Row, sBit, len = [2, 2], 8 , 8
// Hour Row, sBit, len = [3, 3], 8 , 8
// Minute Row, sBit, len = [4, 4], 8 , 8
// Second Row, sBit, len = [5, 5], 8 , 8
// TimeZone Row, sBit, len = [6, 6], 8 , 8
type UniversalTimeAndLocalTimeZone struct {
	Iei   uint8
	Octet [7]uint8
}

// NetworkDaylightSavingTime 9.11.3.19
// value Row, sBit, len = [0, 0], 2 , 2
type NetworkDaylightSavingTime struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// SMSIndication 9.10.3.50A
// Iei Row, sBit, len = [0, 0], 8 , 4
// SAI Row, sBit, len = [0, 0], 1 , 1
type SMSIndication struct {
	Octet uint8
}

// ConfigurationUpdateCompleteMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ConfigurationUpdateCompleteMessageIdentity struct {
	Octet uint8
}

// IdentityRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type IdentityRequestMessageIdentity struct {
	Octet uint8
}

// SpareHalfOctetAndIdentityType 9.11.3.3 9.5
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
type SpareHalfOctetAndIdentityType struct {
	Octet uint8
}

// IdentityResponseMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type IdentityResponseMessageIdentity struct {
	Octet uint8
}

// MobileIdentity 9.11.3.4
// MobileIdentityContents Row, sBit, len = [0, 0], 8 , INF
type MobileIdentity struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// NotificationMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type NotificationMessageIdentity struct {
	Octet uint8
}

// SpareHalfOctetAndAccessType 9.11.3.11 9.5
// AccessType Row, sBit, len = [0, 0], 2 , 2
type SpareHalfOctetAndAccessType struct {
	Octet uint8
}

// NotificationResponseMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type NotificationResponseMessageIdentity struct {
	Octet uint8
}

// SecurityModeCommandMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type SecurityModeCommandMessageIdentity struct {
	Octet uint8
}

// SelectedNASSecurityAlgorithms 9.11.3.34
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 8 , 4
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 4 , 4
type SelectedNASSecurityAlgorithms struct {
	Iei   uint8
	Octet uint8
}

// ReplayedUESecurityCapabilities 9.11.3.54
// EA0_5G Row, sBit, len = [0, 0], 8 , 1
// EA1_128_5G Row, sBit, len = [0, 0], 7 , 1
// EA2_128_5G Row, sBit, len = [0, 0], 6 , 1
// EA3_128_5G Row, sBit, len = [0, 0], 5 , 1
// EA4_5G Row, sBit, len = [0, 0], 4 , 1
// EA5_5G Row, sBit, len = [0, 0], 3 , 1
// EA6_5G Row, sBit, len = [0, 0], 2 , 1
// EA7_5G Row, sBit, len = [0, 0], 1 , 1
// IA0_5G Row, sBit, len = [1, 1], 8 , 1
// IA1_128_5G Row, sBit, len = [1, 1], 7 , 1
// IA2_128_5G Row, sBit, len = [1, 1], 6 , 1
// IA3_128_5G Row, sBit, len = [1, 1], 5 , 1
// IA4_5G Row, sBit, len = [1, 1], 4 , 1
// IA5_5G Row, sBit, len = [1, 1], 3 , 1
// IA6_5G Row, sBit, len = [1, 1], 2 , 1
// IA7_5G Row, sBit, len = [1, 1], 1 , 1
// EEA0 Row, sBit, len = [2, 2], 8 , 1
// EEA1_128 Row, sBit, len = [2, 2], 7 , 1
// EEA2_128 Row, sBit, len = [2, 2], 6 , 1
// EEA3_128 Row, sBit, len = [2, 2], 5 , 1
// EEA4 Row, sBit, len = [2, 2], 4 , 1
// EEA5 Row, sBit, len = [2, 2], 3 , 1
// EEA6 Row, sBit, len = [2, 2], 2 , 1
// EEA7 Row, sBit, len = [2, 2], 1 , 1
// EIA0 Row, sBit, len = [3, 3], 8 , 1
// EIA1_128 Row, sBit, len = [3, 3], 7 , 1
// EIA2_128 Row, sBit, len = [3, 3], 6 , 1
// EIA3_128 Row, sBit, len = [3, 3], 5 , 1
// EIA4 Row, sBit, len = [3, 3], 4 , 1
// EIA5 Row, sBit, len = [3, 3], 3 , 1
// EIA6 Row, sBit, len = [3, 3], 2 , 1
// EIA7 Row, sBit, len = [3, 3], 1 , 1
// Spare Row, sBit, len = [4, 7], 8 , 32
type ReplayedUESecurityCapabilities struct {
	Iei   uint8
	Len   uint8
	Buffer []uint8
}

// IMEISVRequest 9.11.3.28
// Iei Row, sBit, len = [0, 0], 8 , 4
// IMEISVRequestValue Row, sBit, len = [0, 0], 3 , 3
type IMEISVRequest struct {
	Octet uint8
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 7 , 3
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 3 , 3
type SelectedEPSNASSecurityAlgorithms struct {
	Iei   uint8
	Octet uint8
}

// Additional5GSecurityInformation 9.11.3.12
// RINMR Row, sBit, len = [0, 0], 2 , 1
// HDP Row, sBit, len = [0, 0], 1 , 1
type Additional5GSecurityInformation struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA0 Row, sBit, len = [0, 0], 8 , 1
// EEA1_128 Row, sBit, len = [0, 0], 7 , 1
// EEA2_128 Row, sBit, len = [0, 0], 6 , 1
// EEA3_128 Row, sBit, len = [0, 0], 5 , 1
// EEA4 Row, sBit, len = [0, 0], 4 , 1
// EEA5 Row, sBit, len = [0, 0], 3 , 1
// EEA6 Row, sBit, len = [0, 0], 2 , 1
// EEA7 Row, sBit, len = [0, 0], 1 , 1
// EIA0 Row, sBit, len = [1, 1], 8 , 1
// EIA1_128 Row, sBit, len = [1, 1], 7 , 1
// EIA2_128 Row, sBit, len = [1, 1], 6 , 1
// EIA3_128 Row, sBit, len = [1, 1], 5 , 1
// EIA4 Row, sBit, len = [1, 1], 4 , 1
// EIA5 Row, sBit, len = [1, 1], 3 , 1
// EIA6 Row, sBit, len = [1, 1], 2 , 1
// EIA7 Row, sBit, len = [1, 1], 1 , 1
// UEA0 Row, sBit, len = [2, 2], 8 , 1
// UEA1 Row, sBit, len = [2, 2], 7 , 1
// UEA2 Row, sBit, len = [2, 2], 6 , 1
// UEA3 Row, sBit, len = [2, 2], 5 , 1
// UEA4 Row, sBit, len = [2, 2], 4 , 1
// UEA5 Row, sBit, len = [2, 2], 3 , 1
// UEA6 Row, sBit, len = [2, 2], 2 , 1
// UEA7 Row, sBit, len = [2, 2], 1 , 1
// UIA1 Row, sBit, len = [3, 3], 7 , 1
// UIA2 Row, sBit, len = [3, 3], 6 , 1
// UIA3 Row, sBit, len = [3, 3], 5 , 1
// UIA4 Row, sBit, len = [3, 3], 4 , 1
// UIA5 Row, sBit, len = [3, 3], 3 , 1
// UIA6 Row, sBit, len = [3, 3], 2 , 1
// UIA7 Row, sBit, len = [3, 3], 1 , 1
// GEA1 Row, sBit, len = [4, 4], 7 , 1
// GEA2 Row, sBit, len = [4, 4], 6 , 1
// GEA3 Row, sBit, len = [4, 4], 5 , 1
// GEA4 Row, sBit, len = [4, 4], 4 , 1
// GEA5 Row, sBit, len = [4, 4], 3 , 1
// GEA6 Row, sBit, len = [4, 4], 2 , 1
// GEA7 Row, sBit, len = [4, 4], 1 , 1
type ReplayedS1UESecurityCapabilities struct {
	Iei   uint8
	Len   uint8
	Buffer []uint8
}

// SecurityModeCompleteMessageIdentity 9.6
// MessageType Row, sBit, len = [0, 0], 8 , 8
type SecurityModeCompleteMessageIdentity struct {
	Octet uint8
}

// IMEISV 9.11.3.4
// IdentityDigit1 Row, sBit, len = [0, 0], 8 , 4
// OddEvenIdic Row, sBit, len = [0, 0], 4 , 1
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
// IdentityDigitP_1 Row, sBit, len = [1, 1], 8 , 4
// IdentityDigitP Row, sBit, len = [1, 1], 4 , 4
// IdentityDigitP_3 Row, sBit, len = [2, 2], 8 , 4
// IdentityDigitP_2 Row, sBit, len = [2, 2], 4 , 4
// IdentityDigitP_5 Row, sBit, len = [3, 3], 8 , 4
// IdentityDigitP_4 Row, sBit, len = [3, 3], 4 , 4
// IdentityDigitP_7 Row, sBit, len = [4, 4], 8 , 4
// IdentityDigitP_6 Row, sBit, len = [4, 4], 4 , 4
// IdentityDigitP_9 Row, sBit, len = [5, 5], 8 , 4
// IdentityDigitP_8 Row, sBit, len = [5, 5], 4 , 4
// IdentityDigitP_11 Row, sBit, len = [6, 6], 8 , 4
// IdentityDigitP_10 Row, sBit, len = [6, 6], 4 , 4
// IdentityDigitP_13 Row, sBit, len = [7, 7], 8 , 4
// IdentityDigitP_12 Row, sBit, len = [7, 7], 4 , 4
// IdentityDigitP_15 Row, sBit, len = [8, 8], 8 , 4
// IdentityDigitP_14 Row, sBit, len = [8, 8], 4 , 4
type IMEISV struct {
	Iei   uint8
	Len   uint16
	Octet [9]uint8
}

// SecurityModeRejectMessageIdentity 9.6
// MessageType Row, sBit, len = [0, 0], 8 , 8
type SecurityModeRejectMessageIdentity struct {
	Octet uint8
}

// MessageAuthenticationCode MAC 9.8
// MAC Row, sBit, len = [0, 3], 8 , 32
type MessageAuthenticationCode struct {
	Octet [4]uint8
}

// SequenceNumber 9.10
// SQN Row, sBit, len = [0, 0], 8 , 8
type SequenceNumber struct {
	Octet uint8
}

// Plain5GSNASMessage 9.9
type Plain5GSNASMessage struct {
}

// STATUSMessageIdentity5GMM 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type STATUSMessageIdentity5GMM struct {
	Octet uint8
}

// PTI 9.6
// PTI Row, sBit, len = [0, 0], 8 , 8
type PTI struct {
	Octet uint8
}

// PDUSESSIONESTABLISHMENTREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTREQUESTMessageIdentity struct {
	Octet uint8
}

// IntegrityProtectionMaximumDataRate 9.11.4.7
// MaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink Row, sBit, len = [0, 0], 8 , 8
// MaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink Row, sBit, len = [1, 1], 8 , 8
type IntegrityProtectionMaximumDataRate struct {
	Iei   uint8
	Octet [2]uint8
}

// PDUSessionType 9.11.4.11
// Iei Row, sBit, len = [0, 0], 8 , 4
// Spare Row, sBit, len = [0, 0], 4 , 1
// PDUSessionTypeValue Row, sBit, len = [0, 0], 3 , 3
type PDUSessionType struct {
	Octet uint8
}

// SSCMode 9.11.4.16
// Iei Row, sBit, len = [0, 0], 8 , 4
// Spare Row, sBit, len = [0, 0], 4 , 1
// SSCMode Row, sBit, len = [0, 0], 3 , 3
type SSCMode struct {
	Octet uint8
}

// Capability5GSM 9.11.4.1
// MH6PDU Row, sBit, len = [0, 0], 2 , 1
// RqoS Row, sBit, len = [0, 0], 1 , 1
// Spare Row, sBit, len = [1, 12], 8 , 96
type Capability5GSM struct {
	Iei   uint8
	Len   uint8
	Octet [13]uint8
}

// MaximumNumberOfSupportedPacketFilters 9.11.4.9
// MaximumNumberOfSupportedPacketFilters Row, sBit, len = [0, 1], 8 , 10
type MaximumNumberOfSupportedPacketFilters struct {
	Iei   uint8
	Octet [2]uint8
}

// AlwaysonPDUSessionRequested 9.11.4.4
// Iei Row, sBit, len = [0, 0], 8 , 4
// APSR Row, sBit, len = [0, 0], 1 , 1
type AlwaysonPDUSessionRequested struct {
	Octet uint8
}

// SMPDUDNRequestContainer 9.11.4.15
// DNSpecificIdentity Row, sBit, len = [0, 0], 8 , INF
type SMPDUDNRequestContainer struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// ExtendedProtocolConfigurationOptions 9.11.4.6
// ExtendedProtocolConfigurationOptionsContents Row, sBit, len = [0, 0], 8 , INF
type ExtendedProtocolConfigurationOptions struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// PDUSESSIONESTABLISHMENTACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTACCEPTMessageIdentity struct {
	Octet uint8
}

// SelectedSSCModeAndSelectedPDUSessionType 9.11.4.11 9.11.4.16
// SSCMode Row, sBit, len = [0, 0], 7 , 3
// PDUSessionType  Row, sBit, len = [0, 0], 3 , 3
type SelectedSSCModeAndSelectedPDUSessionType struct {
	Octet uint8
}

// AuthorizedQosRules 9.11.4.13
// QosRule Row, sBit, len = [0, 0], 3 , INF
type AuthorizedQosRules struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// SessionAMBR 9.11.4.14
// UnitForSessionAMBRForDownlink Row, sBit, len = [0, 0], 8 , 8
// SessionAMBRForDownlink Row, sBit, len = [1, 2], 8 , 16
// UnitForSessionAMBRForUplink Row, sBit, len = [3, 3], 8 , 8
// SessionAMBRForUplink Row, sBit, len = [4, 5], 8 , 16
type SessionAMBR struct {
	Iei   uint8
	Len   uint8
	Octet [6]uint8
}

// Cause5GSM 9.11.4.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
type Cause5GSM struct {
	Iei   uint8
	Octet uint8
}

// PDUAddress 9.11.4.10
// PDUSessionTypeValue Row, sBit, len = [0, 0], 3 , 3
// PDUAddressInformation Row, sBit, len = [1, 12], 8 , 96
type PDUAddress struct {
	Iei   uint8
	Len   uint8
	Octet [13]uint8
}

// RQTimerValue 9.11.2.3
// Unit Row, sBit, len = [0, 0], 8 , 3
// TimerValue Row, sBit, len = [0, 0], 5 , 5
type RQTimerValue struct {
	Iei   uint8
	Octet uint8
}

// AlwaysonPDUSessionIndication 9.11.4.3
// Iei Row, sBit, len = [0, 0], 8 , 4
// APSI Row, sBit, len = [0, 0], 1 , 1
type AlwaysonPDUSessionIndication struct {
	Octet uint8
}

// MappedEPSBearerContexts 9.11.4.8
// MappedEPSBearerContext Row, sBit, len = [0, 0], 8 , INF
type MappedEPSBearerContexts struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// AuthorizedQosFlowDescriptions 9.11.4.12
// QoSFlowDescriptions Row, sBit, len = [0, 0], 8 , INF
type AuthorizedQosFlowDescriptions struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// PDUSESSIONESTABLISHMENTREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTREJECTMessageIdentity struct {
	Octet uint8
}

// AllowedSSCMode 9.11.4.5
// Iei Row, sBit, len = [0, 0], 8 , 4
// SSC3 Row, sBit, len = [0, 0], 3 , 1
// SSC2 Row, sBit, len = [0, 0], 2 , 1
// SSC1 Row, sBit, len = [0, 0], 1 , 1
type AllowedSSCMode struct {
	Octet uint8
}

// PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONAUTHENTICATIONRESULTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONAUTHENTICATIONRESULTMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONMODIFICATIONREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONREQUESTMessageIdentity struct {
	Octet uint8
}

// RequestedQosRules 9.11.4.13
// QoSRules Row, sBit, len = [0, 0], 8 , INF
type RequestedQosRules struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

// RequestedQosFlowDescriptions 9.11.4.12
// QoSFlowDescriptions Row, sBit, len = [0, 0], 8 , INF
type RequestedQosFlowDescriptions struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

// PDUSESSIONMODIFICATIONREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONREJECTMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONMODIFICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONCOMMANDMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONMODIFICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONCOMPLETEMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONRELEASEREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONRELEASEREQUESTMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONRELEASEREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONRELEASEREJECTMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONRELEASECOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONRELEASECOMMANDMessageIdentity struct {
	Octet uint8
}

// PDUSESSIONRELEASECOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONRELEASECOMPLETEMessageIdentity struct {
	Octet uint8
}

// STATUSMessageIdentity5GSM 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type STATUSMessageIdentity5GSM struct {
	Octet uint8
}

// PDUSessionID 9.4
// PDUSessionID Row, sBit, len = [0, 0], 8 , 8
type PDUSessionID struct {
    Octet uint8
}