package ngap_message_test

import (
	"encoding/hex"
	"git.cs.nctu.edu.tw/calee/sctp"
	"gofree5gc/lib/CommonConsumerTestData/AMF/TestAmf"
	"gofree5gc/lib/CommonConsumerTestData/AMF/TestComm"
	"gofree5gc/lib/aper"
	"gofree5gc/lib/http2_util"
	"gofree5gc/lib/nas/nasMessage"
	"gofree5gc/lib/nas/nasTestpacket"
	"gofree5gc/lib/nas/nasType"
	"gofree5gc/lib/ngap"
	"gofree5gc/lib/ngap/ngapConvert"
	"gofree5gc/lib/ngap/ngapType"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/amf/amf_context"
	"gofree5gc/src/amf/amf_handler"
	"gofree5gc/src/amf/amf_ngap/ngap_message"
	"gofree5gc/src/amf/amf_util"
	"gofree5gc/src/amf/gmm/gmm_message"
	"gofree5gc/src/amf/logger"
	"gofree5gc/src/smf/PDUSession"
	"gofree5gc/src/smf/smf_handler"
	"gofree5gc/src/test/ngapTestpacket"
	"testing"
	"time"
)

var ran *amf_context.AmfRan

func init() {
	go amf_handler.Handle()
	go smf_handler.Handle()
	go func() {
		router := PDUSession.NewRouter()
		server, err := http2_util.NewServer(":29502", TestAmf.AmfLogPath, router)
		if err == nil && server != nil {
			err = server.ListenAndServeTLS(TestAmf.AmfPemPath, TestAmf.AmfKeyPath)
		}
		if err != nil {
			logger.NgapLog.Error(err.Error())
		}
	}()
	TestAmf.SctpSever()
	TestAmf.AmfInit()
	TestAmf.SctpConnectToServer(models.AccessType__3_GPP_ACCESS)
	ran = TestAmf.TestAmf.AmfRanPool[TestAmf.Laddr.String()]
}

func TestSendNGSetupResponse(t *testing.T) {
	time.Sleep(200 * time.Millisecond)

	ngap_message.SendNGSetupResponse(ran)
}

func TestSendNGSetupFailure(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	cause := ngapType.Cause{
		Present: ngapType.CausePresentMisc,
		Misc: &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		},
	}

	ngap_message.SendNGSetupFailure(ran, cause)
}

func TestSendNGReset(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	cause := ngapType.Cause{
		Present: ngapType.CausePresentMisc,
		Misc: &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		},
	}

	ngap_message.SendNGReset(ran, cause, nil)
}

func TestSendNGResetAcknowledge(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	AmfUeNgapID := ngapType.AMFUENGAPID{Value: 13}
	RanUeNgapID := ngapType.RANUENGAPID{Value: 14}
	item := ngapType.UEAssociatedLogicalNGConnectionItem{
		AMFUENGAPID: &AmfUeNgapID,
		RANUENGAPID: &RanUeNgapID,
	}

	partOfNGInterface := ngapType.UEAssociatedLogicalNGConnectionList{
		List: []ngapType.UEAssociatedLogicalNGConnectionItem{
			item,
		},
	}

	ngap_message.SendNGResetAcknowledge(ran, &partOfNGInterface, nil)
}

func TestSendDownlinkNasTransport(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	nasPdu, err := gmm_message.BuildIdentityRequest(nasMessage.MobileIdentity5GSTypeSuci)
	if err != nil {
		t.Error(err.Error())
	}
	ngap_message.SendDownlinkNasTransport(ue.RanUe[models.AccessType__3_GPP_ACCESS], nasPdu, nil)
}

func TestSendPDUSessionResourceReleaseCommand(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	transfer := TestComm.GetPDUSessionResourceReleaseCommandTransfer()
	pduSessionResourceToReleaseListRelCmd := ngapType.PDUSessionResourceToReleaseListRelCmd{}

	item := ngapType.PDUSessionResourceToReleaseItemRelCmd{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 10,
		},
		PDUSessionResourceReleaseCommandTransfer: transfer,
		// TODO: use real PDUSessionResourceSetupRequestTransfer
		// PDUSessionResourceSetupRequestTransfer: aper.OctetString("\x01\x02\x03"),
	}
	pduSessionResourceToReleaseListRelCmd.List = append(pduSessionResourceToReleaseListRelCmd.List, item)

	ngap_message.SendPDUSessionResourceReleaseCommand(ue.RanUe[models.AccessType__3_GPP_ACCESS], []byte{12}, pduSessionResourceToReleaseListRelCmd)
}

func TestSendUEContextReleaseCommand(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	ngap_message.SendUEContextReleaseCommand(ue.RanUe[models.AccessType__3_GPP_ACCESS], amf_context.UeContextN2NormalRelease, ngapType.CausePresentMisc, ngapType.CauseMiscPresentUnspecified)
}

func TestSendErrorIndication(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	amfUeNgapID := int64(123)
	ranUeNgapID := int64(456)

	cause := ngapType.Cause{
		Present: ngapType.CausePresentMisc,
		Misc: &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		},
	}

	ngap_message.SendErrorIndication(ran, &amfUeNgapID, &ranUeNgapID, &cause, nil)
}

func TestSendUERadioCapabilityCheckRequest(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	ngap_message.SendUERadioCapabilityCheckRequest(ue.RanUe[models.AccessType__3_GPP_ACCESS])
}

func TestSendHandoverCancelAcknowledge(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	ngap_message.SendHandoverCancelAcknowledge(ue.RanUe[models.AccessType__3_GPP_ACCESS], nil)
}

func TestSendPDUSessionResourceSetupRequest(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	// nasPdu := []byte{0x01, 0x02}
	pduSessionResourceSetupListSUReq := ngapType.PDUSessionResourceSetupListSUReq{}

	item := ngapType.PDUSessionResourceSetupItemSUReq{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 1,
		},
		SNSSAI: ngapType.SNSSAI{
			SST: ngapType.SST{
				Value: aper.OctetString("\x01"),
			},
		},
		// TODO: use real PDUSessionResourceSetupRequestTransfer
		// PDUSessionResourceSetupRequestTransfer: aper.OctetString("\x01\x02\x03"),
	}

	pduSessionResourceSetupListSUReq.List = append(pduSessionResourceSetupListSUReq.List, item)

	ngap_message.SendPDUSessionResourceSetupRequest(ue.RanUe[models.AccessType__3_GPP_ACCESS], nil, pduSessionResourceSetupListSUReq)
}

func TestSendPDUSessionResourceModifyConfirm(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ue.RanUe[models.AccessType__3_GPP_ACCESS].AmfUeNgapId = 1
	ue.RanUe[models.AccessType__3_GPP_ACCESS].RanUeNgapId = 2
	pduSessionResourceModifyConfirmList := ngapType.PDUSessionResourceModifyListModCfm{}
	item := ngapType.PDUSessionResourceModifyItemModCfm{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 10,
		},
		// TODO: use real transfer
		PDUSessionResourceModifyConfirmTransfer: aper.OctetString("\x01\x02"),
	}
	pduSessionResourceModifyConfirmList.List = append(pduSessionResourceModifyConfirmList.List, item)

	pduSessionResourceFailedToModifyListModCfm := ngapType.PDUSessionResourceFailedToModifyListModCfm{}
	item2 := ngapType.PDUSessionResourceFailedToModifyItemModCfm{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 5,
		},
		// TODO: use real transfer
		PDUSessionResourceModifyIndicationUnsuccessfulTransfer: aper.OctetString("\x01\x02"),
	}
	pduSessionResourceFailedToModifyListModCfm.List = append(pduSessionResourceFailedToModifyListModCfm.List, item2)

	ngap_message.SendPDUSessionResourceModifyConfirm(ue.RanUe[models.AccessType__3_GPP_ACCESS], pduSessionResourceModifyConfirmList, pduSessionResourceFailedToModifyListModCfm, nil)
}

func TestSendPDUSessionResourceModifyRequest(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	pduModifyRequestList := ngapType.PDUSessionResourceModifyListModReq{}
	item := ngapType.PDUSessionResourceModifyItemModReq{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 1,
		},
		NASPDU: &ngapType.NASPDU{
			// TODO: use real nas pdu
			Value: aper.OctetString("\x01\x02"),
		},
		// TODO: use real transfer
		PDUSessionResourceModifyRequestTransfer: aper.OctetString("\x01\x02\x03"),
	}

	pduModifyRequestList.List = append(pduModifyRequestList.List, item)

	ngap_message.SendPDUSessionResourceModifyRequest(ue.RanUe[models.AccessType__3_GPP_ACCESS], pduModifyRequestList)
}

func TestSendInitialContextSetupRequest(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ue.DerivateAnKey(models.AccessType__3_GPP_ACCESS)
	ue.PlmnId = models.PlmnId{
		Mcc: "208",
		Mnc: "93",
	}

	ue.AmPolicyAssociation = &models.PolicyAssociation{}
	ue.AmPolicyAssociation.ServAreaRes = &models.ServiceAreaRestriction{}
	ue.AmPolicyAssociation.ServAreaRes.Areas = append(ue.AmPolicyAssociation.ServAreaRes.Areas, models.Area{
		Tacs: []string{
			"000102",
		},
	})

	ngap_message.SendInitialContextSetupRequest(ue, ue.GetAnType(), nil, nil, nil, nil, nil, nil)
}

func TestSendUEContextModificationRequest(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	oldAmfUeNgapID := int64(1234)

	emergencyFallbackIndicator := ngapType.EmergencyFallbackIndicator{}
	emergencyFallbackIndicator.EmergencyFallbackRequestIndicator.Value = ngapType.EmergencyFallbackRequestIndicatorPresentEmergencyFallbackRequested
	emergencyFallbackIndicator.EmergencyServiceTargetCN = new(ngapType.EmergencyServiceTargetCN)
	emergencyFallbackIndicator.EmergencyServiceTargetCN.Value = ngapType.EmergencyServiceTargetCNPresentFiveGC

	ngap_message.SendUEContextModificationRequest(ue, models.AccessType__3_GPP_ACCESS, &oldAmfUeNgapID, nil, nil, nil, &emergencyFallbackIndicator)
}

func TestSendHandoverCommand(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	pduSessionResourceHandoverList := ngapType.PDUSessionResourceHandoverList{}
	pduSessionResourceToReleaseList := ngapType.PDUSessionResourceToReleaseListHOCmd{}

	item1 := ngapType.PDUSessionResourceHandoverItem{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 10,
		},
		// TODO: use real transfer
		HandoverCommandTransfer: aper.OctetString("\x01\x02\x03"),
	}
	pduSessionResourceHandoverList.List = append(pduSessionResourceHandoverList.List, item1)

	item2 := ngapType.PDUSessionResourceToReleaseItemHOCmd{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 5,
		},
		// TODO: use real transfer
		HandoverPreparationUnsuccessfulTransfer: aper.OctetString("\x01\x02"),
	}
	pduSessionResourceToReleaseList.List = append(pduSessionResourceToReleaseList.List, item2)

	container := ngapType.TargetToSourceTransparentContainer{
		Value: []byte{0x00, 0x01, 0x00, 0x00},
	}

	ngap_message.SendHandoverCommand(ue.RanUe[models.AccessType__3_GPP_ACCESS], pduSessionResourceHandoverList, pduSessionResourceToReleaseList, container, nil)
}

func TestSendHandoverPreparationFailure(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	cause := ngapType.Cause{
		Present: ngapType.CausePresentMisc,
		Misc: &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		},
	}

	ngap_message.SendHandoverPreparationFailure(ue.RanUe[models.AccessType__3_GPP_ACCESS], cause, nil)
}

func TestSendHandoverRequest(t *testing.T) {

	time.Sleep(100 * time.Millisecond)
	// Target Ran
	_, err := sctp.DialSCTP("sctp", TestAmf.Laddr2, TestAmf.ServerAddr)
	if err != nil {
		t.Errorf("Dial Sctp Err : %s", err.Error())
	}
	time.Sleep(100 * time.Millisecond)
	ran := TestAmf.TestAmf.AmfRanPool[TestAmf.Laddr2.String()]
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ue.NCC = 5
	ue.NH, _ = hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

	ue.SecurityCapabilities.NREncryptionAlgorithms = [2]byte{0xe0, 0x00}
	ue.SecurityCapabilities.NRIntegrityProtectionAlgorithms = [2]byte{0xe0, 0x00}
	ue.SecurityCapabilities.EUTRAEncryptionAlgorithms = [2]byte{0xe0, 0x00}
	ue.SecurityCapabilities.EUTRAIntegrityProtectionAlgorithms = [2]byte{0xe0, 0x00}

	cause := ngapType.Cause{
		Present: ngapType.CausePresentMisc,
		Misc: &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		},
	}

	pduSessionResourceSetupListHOReq := ngapType.PDUSessionResourceSetupListHOReq{}
	setupItem := ngapType.PDUSessionResourceSetupItemHOReq{

		PDUSessionID: ngapType.PDUSessionID{
			Value: 10,
		},
		SNSSAI: ngapType.SNSSAI{
			SST: ngapType.SST{
				Value: aper.OctetString("\x11"),
			},
		},
		// HandoverRequestTransfer: aper.OctetString("\x11\x22\x33"),
	}

	pduSessionResourceSetupListHOReq.List = append(pduSessionResourceSetupListHOReq.List, setupItem)

	sourceToTargetTransparentContainer := ngapType.SourceToTargetTransparentContainer{
		Value: aper.OctetString("\x30\x33\x99"),
	}

	ngap_message.SendHandoverRequest(ue.RanUe[models.AccessType__3_GPP_ACCESS], ran, cause, pduSessionResourceSetupListHOReq, sourceToTargetTransparentContainer, false)
}

func TestSendPathSwitchRequestAcknowledge(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	pduSessionResourceSwitchedList := ngapType.PDUSessionResourceSwitchedList{}
	pduSessionResourceReleasedList := ngapType.PDUSessionResourceReleasedListPSAck{}

	switchedItem := ngapType.PDUSessionResourceSwitchedItem{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 10,
		},
		// PathSwitchRequestAcknowledgeTransfer: aper.OctetString("\x11\x22\x33"),
	}
	pduSessionResourceSwitchedList.List = append(pduSessionResourceSwitchedList.List, switchedItem)

	releasedItem := ngapType.PDUSessionResourceReleasedItemPSAck{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 5,
		},
		// PathSwitchRequestUnsuccessfulTransfer: aper.OctetString("\x22\x33\x44"),
	}
	pduSessionResourceReleasedList.List = append(pduSessionResourceReleasedList.List, releasedItem)

	ue.NCC = 5
	ue.NH, _ = hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

	ue.SecurityCapabilities.NREncryptionAlgorithms = [2]byte{0xe0, 0x00}
	ue.SecurityCapabilities.NRIntegrityProtectionAlgorithms = [2]byte{0xe0, 0x00}
	ue.SecurityCapabilities.EUTRAEncryptionAlgorithms = [2]byte{0xe0, 0x00}
	ue.SecurityCapabilities.EUTRAIntegrityProtectionAlgorithms = [2]byte{0xe0, 0x00}

	ngap_message.SendPathSwitchRequestAcknowledge(ue.RanUe[models.AccessType__3_GPP_ACCESS], pduSessionResourceSwitchedList, pduSessionResourceReleasedList, false, nil, nil, nil)
}

func TestSendPathSwitchRequestFailure(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	pduSessionResourceReleasedList := ngapType.PDUSessionResourceReleasedListPSFail{}

	releasedItem := ngapType.PDUSessionResourceReleasedItemPSFail{
		PDUSessionID: ngapType.PDUSessionID{
			Value: 5,
		},
		PathSwitchRequestUnsuccessfulTransfer: aper.OctetString("\x22\x33\x44"),
	}
	pduSessionResourceReleasedList.List = append(pduSessionResourceReleasedList.List, releasedItem)
	ranUe := ue.RanUe[models.AccessType__3_GPP_ACCESS]

	ngap_message.SendPathSwitchRequestFailure(ranUe.Ran, ranUe.AmfUeNgapId, ranUe.RanUeNgapId, &pduSessionResourceReleasedList, nil)
}

func TestSendDownlinkRanStatusTransfer(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	item := ngapType.DRBsSubjectToStatusTransferItem{
		DRBID: ngapType.DRBID{Value: int64(8)},
		DRBStatusUL: ngapType.DRBStatusUL{
			Present: ngapType.DRBStatusULPresentDRBStatusUL18,
			DRBStatusUL12: &ngapType.DRBStatusUL12{
				ULCOUNTValue: ngapType.COUNTValueForPDCPSN12{
					PDCPSN12:    int64(2),
					HFNPDCPSN12: int64(1),
				},
			},
			DRBStatusUL18: &ngapType.DRBStatusUL18{
				ULCOUNTValue: ngapType.COUNTValueForPDCPSN18{
					PDCPSN18:    int64(4),
					HFNPDCPSN18: int64(3),
				},
			},
		},
		DRBStatusDL: ngapType.DRBStatusDL{
			Present: ngapType.DRBStatusDLPresentDRBStatusDL12,
			DRBStatusDL12: &ngapType.DRBStatusDL12{
				DLCOUNTValue: ngapType.COUNTValueForPDCPSN12{
					PDCPSN12:    int64(2),
					HFNPDCPSN12: int64(1),
				},
			},
			DRBStatusDL18: &ngapType.DRBStatusDL18{
				DLCOUNTValue: ngapType.COUNTValueForPDCPSN18{
					PDCPSN18:    int64(4),
					HFNPDCPSN18: int64(3),
				},
			},
		},
	}
	ranStatusTransferTransparentContainer := ngapType.RANStatusTransferTransparentContainer{
		DRBsSubjectToStatusTransferList: ngapType.DRBsSubjectToStatusTransferList{
			List: []ngapType.DRBsSubjectToStatusTransferItem{
				item,
			},
		},
	}

	ngap_message.SendDownlinkRanStatusTransfer(ue.RanUe[models.AccessType__3_GPP_ACCESS], ranStatusTransferTransparentContainer)
}

func TestSendPaging(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ue.Tai.PlmnId = &models.PlmnId{
		Mcc: "208",
		Mnc: "93",
	}
	ue.Tai.Tac = "000001"
	ue.RegistrationArea[models.AccessType__3_GPP_ACCESS] = append(ue.RegistrationArea[models.AccessType__3_GPP_ACCESS], ue.Tai)

	tai2 := models.Tai{
		PlmnId: &models.PlmnId{
			Mcc: "208",
			Mnc: "93",
		},
		Tac: "000001",
	}
	ue.RegistrationArea[models.AccessType__3_GPP_ACCESS] = append(ue.RegistrationArea[models.AccessType__3_GPP_ACCESS], tai2)

	ue.Guti = "20893cafe0000000012"

	var ppi int32 = 0

	recommendedCell := amf_context.RecommendedCell{
		NgRanCGI: amf_context.NGRANCGI{
			Present: amf_context.NgRanCgiPresentNRCGI,
			NRCGI: &models.Ncgi{
				PlmnId: &models.PlmnId{
					Mcc: "208",
					Mnc: "93",
				},
				NrCellId: "000000001",
			},
		},
	}
	ue.InfoOnRecommendedCellsAndRanNodesForPaging = new(amf_context.InfoOnRecommendedCellsAndRanNodesForPaging)
	ue.InfoOnRecommendedCellsAndRanNodesForPaging.RecommendedCells = append(ue.InfoOnRecommendedCellsAndRanNodesForPaging.RecommendedCells, recommendedCell)
	pagingPriority := &ngapType.PagingPriority{
		Value: aper.Enumerated(ppi),
	}
	pkg, err := ngap_message.BuildPaging(ue, pagingPriority, false)
	if err != nil {
		t.Errorf("Build Paging failed : %s", err.Error())
	}
	ngap_message.SendPaging(ue, pkg)
	time.Sleep(1 * time.Second)
	amf_util.ClearT3513(ue)
}

func TestSendRerouteNasRequest(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	mobileIdentity5GS := nasType.MobileIdentity5GS{
		Len:    12, // suci
		Buffer: []uint8{0x01, 0x02, 0xf8, 0x39, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x47, 0x78},
	}
	nasPdu := nasTestpacket.GetRegistrationRequest(nasMessage.RegistrationType5GSInitialRegistration, mobileIdentity5GS, nil, nil)
	initialUeMessage := ngapTestpacket.BuildInitialUEMessage(1, nasPdu, "")
	initialUeMessagePkg, _ := ngap.Encoder(initialUeMessage)

	ue.Guti = "20893cafe0000000001"
	ngap_message.SendRerouteNasRequest(ue, ue.GetAnType(), nil, initialUeMessagePkg, nil)
}

func TestSendRanConfigurationUpdateAcknowledge(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ngap_message.SendRanConfigurationUpdateAcknowledge(ran, nil)
}

func TestSendRanConfigurationUpdateFailure(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	cause := ngapType.Cause{
		Present: ngapType.CausePresentMisc,
		Misc: &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		},
	}

	ngap_message.SendRanConfigurationUpdateFailure(ran, cause, nil)
}

func TestSendAMFStatusIndication(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	item := ngapType.UnavailableGUAMIItem{
		GUAMI: ngapType.GUAMI{
			PLMNIdentity: ngapType.PLMNIdentity{
				Value: aper.OctetString("\x02\xf8\x39"),
			},
			AMFRegionID: ngapType.AMFRegionID{
				Value: aper.BitString{
					Bytes:     []byte{0x45, 0x46},
					BitLength: 8,
				},
			},
			AMFSetID: ngapType.AMFSetID{
				Value: aper.BitString{
					Bytes:     []byte{0x45, 0x46},
					BitLength: 10,
				},
			},
			AMFPointer: ngapType.AMFPointer{
				Value: aper.BitString{
					Bytes:     []byte{0x45},
					BitLength: 6,
				},
			},
		},
	}
	unavailableGUAMIList := ngapType.UnavailableGUAMIList{
		List: []ngapType.UnavailableGUAMIItem{
			item,
		},
	}

	ngap_message.SendAMFStatusIndication(ran, unavailableGUAMIList)
}

func TestSendOverloadStart(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	overloadResponse := ngapType.OverloadResponse{
		Present: ngapType.OverloadResponsePresentOverloadAction,
		OverloadAction: &ngapType.OverloadAction{
			Value: ngapType.OverloadActionPresentRejectNonEmergencyMoDt,
		},
	}

	snssai := ngapConvert.SNssaiToNgap(TestAmf.TestAmf.PlmnSupportList[0].SNssaiList[0])
	sliceOverloadItem := ngapType.SliceOverloadItem{
		SNSSAI: snssai,
	}
	overloadStartNSSAIItem := ngapType.OverloadStartNSSAIItem{}
	overloadStartNSSAIItem.SliceOverloadList.List = append(overloadStartNSSAIItem.SliceOverloadList.List, sliceOverloadItem)

	overloadStartNSSAIList := ngapType.OverloadStartNSSAIList{}
	overloadStartNSSAIList.List = append(overloadStartNSSAIList.List, overloadStartNSSAIItem)

	ngap_message.SendOverloadStart(ran, &overloadResponse, 80, &overloadStartNSSAIList)
}

func TestSendOverloadStop(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ngap_message.SendOverloadStop(ran)
}

func TestSendDownlinkRanConfigurationTransfer(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	TLA := ngapType.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0x12, 0x34, 0x56, 0x78},
			BitLength: 32,
		},
	}

	var xnTLAs ngapType.XnTLAs

	xnTLAs.List = append(xnTLAs.List, TLA)

	sONConfigurationTransfer := ngapType.SONConfigurationTransfer{
		TargetRANNodeID: ngapType.TargetRANNodeID{
			GlobalRANNodeID: ngapType.GlobalRANNodeID{
				Present: ngapType.GlobalRANNodeIDPresentGlobalGNBID,
				GlobalGNBID: &ngapType.GlobalGNBID{
					PLMNIdentity: ngapType.PLMNIdentity{
						Value: aper.OctetString("\x02\xf8\x39"),
					},
					GNBID: ngapType.GNBID{
						Present: ngapType.GNBIDPresentGNBID,
						GNBID: &aper.BitString{
							Bytes:     []byte{0x45, 0x46, 0x47},
							BitLength: 24,
						},
					},
				},
			},
			SelectedTAI: ngapType.TAI{
				PLMNIdentity: ngapType.PLMNIdentity{
					Value: aper.OctetString("\x02\xf8\x39"),
				},
				TAC: ngapType.TAC{
					Value: aper.OctetString("\x00\x00\x01"),
				},
			},
		},
		SourceRANNodeID: ngapType.SourceRANNodeID{
			GlobalRANNodeID: ngapType.GlobalRANNodeID{
				Present: ngapType.GlobalRANNodeIDPresentGlobalGNBID,
				GlobalGNBID: &ngapType.GlobalGNBID{
					PLMNIdentity: ngapType.PLMNIdentity{
						Value: aper.OctetString("\x02\xf8\x39"),
					},
					GNBID: ngapType.GNBID{
						Present: ngapType.GNBIDPresentGNBID,
						GNBID: &aper.BitString{
							Bytes:     []byte{0x41, 0x42, 0x43},
							BitLength: 24,
						},
					},
				},
			},
			SelectedTAI: ngapType.TAI{
				PLMNIdentity: ngapType.PLMNIdentity{
					Value: aper.OctetString("\x02\xf8\x39"),
				},
				TAC: ngapType.TAC{
					Value: aper.OctetString("\x00\x00\x01"),
				},
			},
		},
		SONInformation: ngapType.SONInformation{
			Present: ngapType.SONInformationPresentSONInformationRequest,
			SONInformationRequest: &ngapType.SONInformationRequest{
				Value: ngapType.SONInformationRequestPresentXnTNLConfigurationInfo,
			},
		},
		XnTNLConfigurationInfo: ngapType.XnTNLConfigurationInfo{
			XnTransportLayerAddresses: ngapType.XnTLAs{
				List: xnTLAs.List,
			},
		},
	}

	ngap_message.SendDownlinkRanConfigurationTransfer(ran, &sONConfigurationTransfer)
}

func TestSendDownlinkNonUEAssociatedNRPPATransport(t *testing.T) {

	time.Sleep(200 * time.Millisecond)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ue.RanUe[models.AccessType__3_GPP_ACCESS].RoutingID = "ff"

	nRPPaPDU := ngapType.NRPPaPDU{
		Value: aper.OctetString("\x03\x02"),
	}

	ngap_message.SendDownlinkNonUEAssociatedNRPPATransport(ue.RanUe[models.AccessType__3_GPP_ACCESS], nRPPaPDU)
}

func TestSendDeactivateTrace(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	ngap_message.SendDeactivateTrace(ue, models.AccessType__3_GPP_ACCESS)
}

func TestSendLocationReportingControl(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ue.RanUe[models.AccessType__3_GPP_ACCESS].AmfUeNgapId = 123
	ue.RanUe[models.AccessType__3_GPP_ACCESS].RanUeNgapId = 456

	eventType := ngapType.EventType{
		Value: ngapType.EventTypePresentStopChangeOfServeCell,
	}

	ngap_message.SendLocationReportingControl(ue.RanUe[models.AccessType__3_GPP_ACCESS], nil, 0, eventType)
}

func TestSendUETNLABindingReleaseRequest(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ngap_message.SendUETNLABindingReleaseRequest(ue.RanUe[models.AccessType__3_GPP_ACCESS])
}

func TestSendAMFConfigurationUpdate(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	amfSelf := amf_context.AMF_Self()
	amfSelf.HttpIPv4Address = "127.0.0.1"
	amfSelf.HttpIPv6Address = "2001:0db8:85a3:08d3:1319:8a2e:0370:7344"
	amfSelf.TNLWeightFactor = 123

	tNLassociationUsage := ngapType.TNLAssociationUsage{
		Value: ngapType.TNLAssociationUsagePresentBoth,
	}
	tNLAddressWeightFactor := ngapType.TNLAddressWeightFactor{
		Value: amfSelf.TNLWeightFactor,
	}

	ngap_message.SendAMFConfigurationUpdate(ran, tNLassociationUsage, tNLAddressWeightFactor)
}

func TestSendDownlinkUEAssociatedNRPPaTransport(t *testing.T) {

	time.Sleep(200 * time.Millisecond)

	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]
	ue.RanUe[models.AccessType__3_GPP_ACCESS].RoutingID = "ff"

	nRPPaPDU := ngapType.NRPPaPDU{
		Value: aper.OctetString("\x03\x02"),
	}

	ngap_message.SendDownlinkUEAssociatedNRPPaTransport(ue.RanUe[models.AccessType__3_GPP_ACCESS], nRPPaPDU)
}
