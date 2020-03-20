#ifndef ASN1_COMPILER_LANGUAGE_GO_H
#define ASN1_COMPILER_LANGUAGE_GO_H

#include "asn1c_lang.h"

int asn1c_lang_Go_type_REFERENCE(arg_t *);
int asn1c_lang_Go_type_EXTENSIBLE(arg_t *);

int asn1c_lang_Go_type_SEQUENCE(arg_t *);
int asn1c_lang_Go_type_SET(arg_t *);
int asn1c_lang_Go_type_SEx_OF(arg_t *);	/* SET OF or  SEQUENCE OF */
int asn1c_lang_Go_type_CHOICE(arg_t *);

int asn1c_lang_Go_type_common_INTEGER(arg_t *);
int asn1c_lang_Go_type_BIT_STRING(arg_t *);
int asn1c_lang_Go_type_REAL(arg_t *);
int asn1c_lang_Go_type_SIMPLE_TYPE(arg_t *);
int asn1c_lang_Go_type_REFERENCE_Value(arg_t *);

static asn1_language_map_t asn1_lang_Go[] __attribute__ ((unused)) = {
	{ AMT_VALUE, A1TC_REFERENCE,	asn1c_lang_Go_type_REFERENCE_Value },

	{ AMT_TYPE, A1TC_REFERENCE,	asn1c_lang_Go_type_REFERENCE },
	{ AMT_TYPEREF, A1TC_REFERENCE,	asn1c_lang_Go_type_REFERENCE },
	{ AMT_TYPE, A1TC_EXTENSIBLE,	asn1c_lang_Go_type_EXTENSIBLE },
	/*
	 * Constructed types
	 */
	{ AMT_TYPE, ASN_CONSTR_SEQUENCE,	asn1c_lang_Go_type_SEQUENCE },
	{ AMT_TYPE, ASN_CONSTR_SEQUENCE_OF,	asn1c_lang_Go_type_SEx_OF, },
	{ AMT_TYPEREF, ASN_CONSTR_SEQUENCE_OF,	asn1c_lang_Go_type_SEx_OF },
	{ AMT_TYPE, ASN_CONSTR_SET,		asn1c_lang_Go_type_SET },
	{ AMT_TYPE, ASN_CONSTR_SET_OF,		asn1c_lang_Go_type_SEx_OF },
	{ AMT_TYPEREF, ASN_CONSTR_SET_OF,	asn1c_lang_Go_type_SEx_OF },
	{ AMT_TYPE, ASN_CONSTR_CHOICE,		asn1c_lang_Go_type_CHOICE },
	/*
	 * ANY type (deprecated)
	 */
        { AMT_TYPE, ASN_TYPE_ANY,	asn1c_lang_Go_type_SIMPLE_TYPE },
	/*
	 * Basic types
	 */
	{ AMT_TYPE, ASN_BASIC_BOOLEAN,	asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_BASIC_NULL,	asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_BASIC_INTEGER,	asn1c_lang_Go_type_common_INTEGER },
	{ AMT_TYPE, ASN_BASIC_REAL,	asn1c_lang_Go_type_REAL },
	{ AMT_TYPE, ASN_BASIC_ENUMERATED,  asn1c_lang_Go_type_common_INTEGER },
	{ AMT_TYPE, ASN_BASIC_BIT_STRING,	asn1c_lang_Go_type_BIT_STRING },
	{ AMT_TYPE, ASN_BASIC_OCTET_STRING,	asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_BASIC_OBJECT_IDENTIFIER,asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_BASIC_RELATIVE_OID,	asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_BASIC_CHARACTER_STRING,	asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_BASIC_UTCTime,		asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_BASIC_GeneralizedTime,	asn1c_lang_Go_type_SIMPLE_TYPE },
	/*
	 * String types
	 */
	{ AMT_TYPE, ASN_STRING_BMPString,     asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_GeneralString, asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_GraphicString, asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_IA5String,     asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_ISO646String,  asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_NumericString, asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_PrintableString,asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_TeletexString, asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_T61String,     asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_UniversalString,asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_UTF8String,    asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_VideotexString,asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_VisibleString, asn1c_lang_Go_type_SIMPLE_TYPE },
	{ AMT_TYPE, ASN_STRING_ObjectDescriptor,asn1c_lang_Go_type_SIMPLE_TYPE },
	{ 0, 0, 0 }
};

static const char *IEsToID[] = {
	"IdAllowedNSSAI 0",
	"IdAMFName 1",
	"IdAMFOverloadResponse 2",
	"IdAMFSetID 3",
	"IdAMFTNLAssociationFailedToSetupList 4",
	"IdAMFTNLAssociationSetupList 5",
	"IdAMFTNLAssociationToAddList 6",
	"IdAMFTNLAssociationToRemoveList 7",
	"IdAMFTNLAssociationToUpdateList 8",
	"IdAMFTrafficLoadReductionIndication 9",
	"IdAMFUENGAPID 10",
	"IdAssistanceDataForPaging 11",
	"IdBroadcastCancelledAreaList 12",
	"IdBroadcastCompletedAreaList 13",
	"IdCancelAllWarningMessages 14",
	"IdCause 15",
	"IdCellIDListForRestart 16",
	"IdConcurrentWarningMessageInd 17",
	"IdCoreNetworkAssistanceInformation 18",
	"IdCriticalityDiagnostics 19",
	"IdDataCodingScheme 20",
	"IdDefaultPagingDRX 21",
	"IdDirectForwardingPathAvailability 22",
	"IdEmergencyAreaIDListForRestart 23",
	"IdEmergencyFallbackIndicator 24",
	"IdEUTRACGI 25",
	"IdFiveGSTMSI 26",
	"IdGlobalRANNodeID 27",
	"IdGUAMI 28",
	"IdHandoverType 29",
	"IdIMSVoiceSupportIndicator 30",
	"IdIndexToRFSP 31",
	"IdInfoOnRecommendedCellsAndRANNodesForPaging 32",
	"IdLocationReportingRequestType 33",
	"IdMaskedIMEISV 34",
	"IdMessageIdentifier 35",
	"IdMobilityRestrictionList 36",
	"IdNASC 37",
	"IdNASPDU 38",
	"IdNASSecurityParametersFromNGRAN 39",
	"IdNewAMFUENGAPID 40",
	"IdNewSecurityContextInd 41",
	"IdNGAPMessage 42",
	"IdNGRANCGI 43",
	"IdNGRANTraceID 44",
	"IdNRCGI 45",
	"IdNRPPaPDU 46",
	"IdNumberOfBroadcastsRequested 47",
	"IdOldAMF 48",
	"IdOverloadStartNSSAIList 49",
	"IdPagingDRX 50",
	"IdPagingOrigin 51",
	"IdPagingPriority 52",
	"IdPDUSessionResourceAdmittedList 53",
	"IdPDUSessionResourceFailedToModifyListModRes 54",
	"IdPDUSessionResourceFailedToSetupListCxtRes 55",
	"IdPDUSessionResourceFailedToSetupListHOAck 56",
	"IdPDUSessionResourceFailedToSetupListPSReq 57",
	"IdPDUSessionResourceFailedToSetupListSURes 58",
	"IdPDUSessionResourceHandoverList 59",
	"IdPDUSessionResourceListCxtRelCpl 60",
	"IdPDUSessionResourceListHORqd 61",
	"IdPDUSessionResourceModifyListModCfm 62",
	"IdPDUSessionResourceModifyListModInd 63",
	"IdPDUSessionResourceModifyListModReq 64",
	"IdPDUSessionResourceModifyListModRes 65",
	"IdPDUSessionResourceNotifyList 66",
	"IdPDUSessionResourceReleasedListNot 67",
	"IdPDUSessionResourceReleasedListPSAck 68",
	"IdPDUSessionResourceReleasedListPSFail 69",
	"IdPDUSessionResourceReleasedListRelRes 70",
	"IdPDUSessionResourceSetupListCxtReq 71",
	"IdPDUSessionResourceSetupListCxtRes 72",
	"IdPDUSessionResourceSetupListHOReq 73",
	"IdPDUSessionResourceSetupListSUReq 74",
	"IdPDUSessionResourceSetupListSURes 75",
	"IdPDUSessionResourceToBeSwitchedDLList 76",
	"IdPDUSessionResourceSwitchedList 77",
	"IdPDUSessionResourceToReleaseListHOCmd 78",
	"IdPDUSessionResourceToReleaseListRelCmd 79",
	"IdPLMNSupportList 80",
	"IdPWSFailedCellIDList 81",
	"IdRANNodeName 82",
	"IdRANPagingPriority 83",
	"IdRANStatusTransferTransparentContainer 84",
	"IdRANUENGAPID 85",
	"IdRelativeAMFCapacity 86",
	"IdRepetitionPeriod 87",
	"IdResetType 88",
	"IdRoutingID 89",
	"IdRRCEstablishmentCause 90",
	"IdRRCInactiveTransitionReportRequest 91",
	"IdRRCState 92",
	"IdSecurityContext 93",
	"IdSecurityKey 94",
	"IdSerialNumber 95",
	"IdServedGUAMIList 96",
	"IdSliceSupportList 97",
	"IdSONConfigurationTransferDL 98",
	"IdSONConfigurationTransferUL 99",
	"IdSourceAMFUENGAPID 100",
	"IdSourceToTargetTransparentContainer 101",
	"IdSupportedTAList 102",
	"IdTAIListForPaging 103",
	"IdTAIListForRestart 104",
	"IdTargetID 105",
	"IdTargetToSourceTransparentContainer 106",
	"IdTimeToWait 107",
	"IdTraceActivation 108",
	"IdTraceCollectionEntityIPAddress 109",
	"IdUEAggregateMaximumBitRate 110",
	"IdUEAssociatedLogicalNGConnectionList 111",
	"IdUEContextRequest 112",
	"IdUENGAPIDs 114",
	"IdUEPagingIdentity 115",
	"IdUEPresenceInAreaOfInterestList 116",
	"IdUERadioCapability 117",
	"IdUERadioCapabilityForPaging 118",
	"IdUESecurityCapabilities 119",
	"IdUnavailableGUAMIList 120",
	"IdUserLocationInformation 121",
	"IdWarningAreaList 122",
	"IdWarningMessageContents 123",
	"IdWarningSecurityInfo 124",
	"IdWarningType 125",
	"IdAdditionalULNGUUPTNLInformation 126",
	"IdDataForwardingNotPossible 127",
	"IdDLNGUUPTNLInformation 128",
	"IdNetworkInstance 129",
	"IdPDUSessionAggregateMaximumBitRate 130",
	"IdPDUSessionResourceFailedToModifyListModCfm 131",
	"IdPDUSessionResourceFailedToSetupListCxtFail 132",
	"IdPDUSessionResourceListCxtRelReq 133",
	"IdPDUSessionType 134",
	"IdQosFlowAddOrModifyRequestList 135",
	"IdQosFlowSetupRequestList 136",
	"IdQosFlowToReleaseList 137",
	"IdSecurityIndication 138",
	"IdULNGUUPTNLInformation 139",
	"IdULNGUUPTNLModifyList 140",
	"IdWarningAreaCoordinates 141", 
	"\0"
};

static const char *ProceduresToID[] = {
	"IdAMFConfigurationUpdate 0",
	"IdAMFStatusIndication 1",
	"IdCellTrafficTrace 2",
	"IdDeactivateTrace 3",
	"IdDownlinkNASTransport 4",
	"IdDownlinkNonUEAssociatedNRPPaTransport 5",
	"IdDownlinkRANConfigurationTransfer 6",
	"IdDownlinkRANStatusTransfer 7",
	"IdDownlinkUEAssociatedNRPPaTransport 8",
	"IdErrorIndication 9",
	"IdHandoverCancel 10",
	"IdHandoverNotification 11",
	"IdHandoverPreparation 12",
	"IdHandoverResourceAllocation 13",
	"IdInitialContextSetup 14",
	"IdInitialUEMessage 15",
	"IdLocationReportingControl 16",
	"IdLocationReportingFailureIndication 17",
	"IdLocationReport 18",
	"IdNASNonDeliveryIndication 19",
	"IdNGReset 20",
	"IdNGSetup 21",
	"IdOverloadStart 22",
	"IdOverloadStop 23",
	"IdPaging 24",
	"IdPathSwitchRequest 25",
	"IdPDUSessionResourceModify 26",
	"IdPDUSessionResourceModifyIndication 27",
	"IdPDUSessionResourceRelease 28",
	"IdPDUSessionResourceSetup 29",
	"IdPDUSessionResourceNotify 30",
	"IdPrivateMessage 31",
	"IdPWSCancel 32",
	"IdPWSFailureIndication 33",
	"IdPWSRestartIndication 34",
	"IdRANConfigurationUpdate 35",
	"IdRerouteNASRequest 36",
	"IdRRCInactiveTransitionReport 37",
	"IdTraceFailureIndication 38",
	"IdTraceStart 39",
	"IdUEContextModification 40",
	"IdUEContextRelease 41",
	"IdUEContextReleaseRequest 42",
	"IdUERadioCapabilityCheck 43",
	"IdUERadioCapabilityInfoIndication 44",
	"IdUETNLABindingRelease 45",
	"IdUplinkNASTransport 46",
	"IdUplinkNonUEAssociatedNRPPaTransport 47",
	"IdUplinkRANConfigurationTransfer 48",
	"IdUplinkRANStatusTransfer 49",
	"IdUplinkUEAssociatedNRPPaTransport 50",
	"IdWriteReplaceWarning 51",
	"\0"
};

#endif	/* ASN1_COMPILER_LANGUAGE_GO_H */
