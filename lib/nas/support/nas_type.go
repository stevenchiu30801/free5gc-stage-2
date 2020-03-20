// ExtendedProtocolDiscriminator 9.2
type ExtendedProtocolDiscriminator struct {
    Iei uint8
    Octet uint8
}

// SecurityHeaderTypeAndSpareHalfOctet 9.3 9.5
type SecurityHeaderTypeAndSpareHalfOctet struct {
    Octet uint8
}

// AuthenticationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationRequestMessageIdentity struct {
    Octet uint8
}

// NgksiAndSpareHalfOctet 9.11.3.32 9.5
type NgksiAndSpareHalfOctet struct {
    Octet uint8
}

// ABBA 9.11.3.10
type ABBA struct {
}

// AuthenticationParameterRAND 9.11.3.16
type AuthenticationParameterRAND struct {
}

// AuthenticationParameterAUTN 9.11.3.15
type AuthenticationParameterAUTN struct {
}

// EAPMessage 9.11.2.2
type EAPMessage struct {
}

// AuthenticationResponseMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationResponseMessageIdentity struct {
    Octet uint8
}

// AuthenticationResponseParameter 9.11.3.17
type AuthenticationResponseParameter struct {
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
type Cause5GMM struct {
    Iei uint8
    Octet uint8
}

// AuthenticationFailureParameter 9.11.3.14
type AuthenticationFailureParameter struct {
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

// RegistrationType5GSAndNgksi 9.11.3.7 9.11.3.32
type RegistrationType5GSAndNgksi struct {
    Octet uint8
}

// MobileIdentity5GS 9.11.3.4
type MobileIdentity5GS struct {
}

// NoncurrentNativeNASKeySetIdentifier 9.11.3.32
type NoncurrentNativeNASKeySetIdentifier struct {
}

// Capability5GMM 9.11.3.1
type Capability5GMM struct {
}

// UESecurityCapability 9.11.3.54
type UESecurityCapability struct {
}

// RequestedNSSAI 9.11.3.37
type RequestedNSSAI struct {
}

// LastVisitedRegisteredTAI 9.11.3.8
type LastVisitedRegisteredTAI struct {
}

// S1UENetworkCapability 9.11.3.48
type S1UENetworkCapability struct {
}

// UplinkDataStatus 9.11.3.57
type UplinkDataStatus struct {
}

// PDUSessionStatus 9.11.3.44
type PDUSessionStatus struct {
}

// MICOIndication 9.11.3.31
type MICOIndication struct {
}

// UEStatus 9.11.3.56
type UEStatus struct {
}

// AdditionalGUTI 9.11.3.4
type AdditionalGUTI struct {
}

// AllowedPDUSessionStatus 9.11.3.13
type AllowedPDUSessionStatus struct {
}

// UesUsageSetting 9.11.3.55
type UesUsageSetting struct {
}

// RequestedDRXParameters 9.11.3.2A
type RequestedDRXParameters struct {
}

// EPSNASMessageContainer 9.11.3.24
type EPSNASMessageContainer struct {
}

// LADNIndication 9.11.3.29
type LADNIndication struct {
}

// PayloadContainer 9.11.3.39
type PayloadContainer struct {
}

// NetworkSlicingIndication 9.11.3.36
type NetworkSlicingIndication struct {
}

// UpdateType5GS 9.11.3.9A
type UpdateType5GS struct {
}

// NASMessageContainer 9.11.3.33
type NASMessageContainer struct {
}

// RegistrationAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RegistrationAcceptMessageIdentity struct {
    Octet uint8
}

// RegistrationResult5GS 9.11.3.6
type RegistrationResult5GS struct {
}

// GUTI5G 9.11.3.4
type GUTI5G struct {
}

// EquivalentPlmns 9.11.3.45
type EquivalentPlmns struct {
}

// TAIList 9.11.3.9
type TAIList struct {
}

// AllowedNSSAI 9.11.3.37
type AllowedNSSAI struct {
}

// RejectedNSSAI 9.11.3.46
type RejectedNSSAI struct {
}

// ConfiguredNSSAI 9.11.3.37
type ConfiguredNSSAI struct {
}

// NetworkFeatureSupport5GS 9.11.3.5
type NetworkFeatureSupport5GS struct {
}

// PDUSessionReactivationResult 9.11.3.42
type PDUSessionReactivationResult struct {
}

// PDUSessionReactivationResultErrorCause 9.11.3.43
type PDUSessionReactivationResultErrorCause struct {
}

// LADNInformation 9.11.3.30
type LADNInformation struct {
}

// ServiceAreaList 9.11.3.49
type ServiceAreaList struct {
}

// T3512Value 9.11.2.5
type T3512Value struct {
}

// Non3GppDeregistrationTimerValue 9.11.2.4
type Non3GppDeregistrationTimerValue struct {
}

// T3502Value 9.11.2.4
type T3502Value struct {
}

// EmergencyNumberList 9.11.3.23
type EmergencyNumberList struct {
}

// ExtendedEmergencyNumberList 9.11.3.26
type ExtendedEmergencyNumberList struct {
}

// SORTransparentContainer 9.11.3.51
type SORTransparentContainer struct {
}

// NSSAIInclusionMode 9.11.3.37A
type NSSAIInclusionMode struct {
}

// OperatordefinedAccessCategoryDefinitions 9.11.3.38
type OperatordefinedAccessCategoryDefinitions struct {
}

// NegotiatedDRXParameters 9.11.3.2A
type NegotiatedDRXParameters struct {
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
type T3346Value struct {
}

// ULNASTRANSPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ULNASTRANSPORTMessageIdentity struct {
    Octet uint8
}

// PayloadContainerTypeAndSpareHalfOctet 9.11.3.40 9.5
type PayloadContainerTypeAndSpareHalfOctet struct {
    Octet uint8
}

// PDUSessionID 9.11.3.41
type PDUSessionID struct {
}

// OldPDUSessionID 9.11.3.41
type OldPDUSessionID struct {
}

// RequestType 9.11.3.47
type RequestType struct {
}

// SNSSAI 9.11.2.8
type SNSSAI struct {
}

// DNN 9.11.3.21
type DNN struct {
}

// AdditionalInformation 9.11.2.1
type AdditionalInformation struct {
}

// DLNASTRANSPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type DLNASTRANSPORTMessageIdentity struct {
    Octet uint8
}

// BackoffTimerValue 9.11.2.5
type BackoffTimerValue struct {
}

// DeregistrationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type DeregistrationRequestMessageIdentity struct {
    Octet uint8
}

// DeregistrationTypeAndNgksi 9.11.3.20 9.11.3.32
type DeregistrationTypeAndNgksi struct {
    Octet uint8
}

// DeregistrationAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type DeregistrationAcceptMessageIdentity struct {
    Octet uint8
}

// DeregistrationTypeAndSpareHalfOctet 9.11.3.20 9.5
type DeregistrationTypeAndSpareHalfOctet struct {
    Octet uint8
}

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ServiceRequestMessageIdentity struct {
    Octet uint8
}

// NgksiAndServiceType 9.11.3.32 9.11.3.50
type NgksiAndServiceType struct {
    Octet uint8
}

// TMSI5GS 9.11.3.4
type TMSI5GS struct {
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
type ConfigurationUpdateIndication struct {
}

// FullNameForNetwork 9.11.3.35
type FullNameForNetwork struct {
}

// ShortNameForNetwork 9.11.3.35
type ShortNameForNetwork struct {
}

// LocalTimeZone 9.11.3.52
type LocalTimeZone struct {
}

// UniversalTimeAndLocalTimeZone 9.11.3.53
type UniversalTimeAndLocalTimeZone struct {
}

// NetworkDaylightSavingTime 9.11.3.19
type NetworkDaylightSavingTime struct {
}

// SMSIndication 9.10.3.50A
type SMSIndication struct {
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

// IdentityTypeAndSpareHalfOctet 9.11.3.3 9.5
type IdentityTypeAndSpareHalfOctet struct {
    Octet uint8
}

// IdentityResponseMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type IdentityResponseMessageIdentity struct {
    Octet uint8
}

// MobileIdentity 9.11.3.4
type MobileIdentity struct {
}

// NotificationMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type NotificationMessageIdentity struct {
    Octet uint8
}

// AccessTypeAndSpareHalfOctet 9.11.3.11 9.5
type AccessTypeAndSpareHalfOctet struct {
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
type SelectedNASSecurityAlgorithms struct {
    Iei uint8
    Octet uint8
}

// ReplayedUESecurityCapabilities 9.11.3.54
type ReplayedUESecurityCapabilities struct {
}

// IMEISVRequest 9.11.3.28
type IMEISVRequest struct {
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
type SelectedEPSNASSecurityAlgorithms struct {
}

// Additional5GSecurityInformation 9.11.3.12
type Additional5GSecurityInformation struct {
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
type ReplayedS1UESecurityCapabilities struct {
}

// SecurityModeCompleteMessageIdentity 9.6
type SecurityModeCompleteMessageIdentity struct {
    Iei uint8
    Octet uint8
}

// IMEISV 9.11.3.4
type IMEISV struct {
}

// SecurityModeRejectMessageIdentity 9.6
type SecurityModeRejectMessageIdentity struct {
    Iei uint8
    Octet uint8
}

// MessageAuthenticationCode 9.8
type MessageAuthenticationCode struct {
    Iei uint8
    Octet [4]uint8
}

// SequenceNumber 9.10
type SequenceNumber struct {
    Iei uint8
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
type PTI struct {
    Iei uint8
    Octet uint8
}

// PDUSESSIONESTABLISHMENTREQUESTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTREQUESTMessageIdentity struct {
    Octet uint8
}

// IntegrityProtectionMaximumDataRate 9.11.4.7
type IntegrityProtectionMaximumDataRate struct {
    Iei uint8
    Octet [2]uint8
}

// PDUSessionType 9.11.4.11
type PDUSessionType struct {
}

// SSCMode 9.11.4.16
type SSCMode struct {
}

// Capability5GSM 9.11.4.1
type Capability5GSM struct {
}

// MaximumNumberOfSupportedPacketFilters 9.11.4.9
type MaximumNumberOfSupportedPacketFilters struct {
}

// AlwaysonPDUSessionRequested 9.11.4.4
type AlwaysonPDUSessionRequested struct {
}

// SMPDUDNRequestContainer 9.11.4.15
type SMPDUDNRequestContainer struct {
}

// ExtendedProtocolConfigurationOptions 9.11.4.6
type ExtendedProtocolConfigurationOptions struct {
}

// PDUSESSIONESTABLISHMENTACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTACCEPTMessageIdentity struct {
    Octet uint8
}

// SelectedPDUSessionTypeAndSelectedSSCMode 9.11.4.11 9.11.4.16
type SelectedPDUSessionTypeAndSelectedSSCMode struct {
    Octet uint8
}

// AuthorizedQosRules 9.11.4.13
type AuthorizedQosRules struct {
}

// SessionAMBR 9.11.4.14
type SessionAMBR struct {
}

// Cause5GSM 9.11.4.2
type Cause5GSM struct {
}

// PDUAddress 9.11.4.10
type PDUAddress struct {
}

// RQTimerValue 9.11.2.3
type RQTimerValue struct {
}

// AlwaysonPDUSessionIndication 9.11.4.3
type AlwaysonPDUSessionIndication struct {
}

// MappedEPSBearerContexts 9.11.4.8
type MappedEPSBearerContexts struct {
}

// AuthorizedQosFlowDescriptions 9.11.4.12
type AuthorizedQosFlowDescriptions struct {
}

// PDUSESSIONESTABLISHMENTREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTREJECTMessageIdentity struct {
    Octet uint8
}

// AllowedSSCMode 9.11.4.5
type AllowedSSCMode struct {
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
type RequestedQosRules struct {
}

// RequestedQosFlowDescriptions 9.11.4.12
type RequestedQosFlowDescriptions struct {
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

