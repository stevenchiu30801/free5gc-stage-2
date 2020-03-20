package amf_consumer_test

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/bson"
	"gofree5gc/lib/CommonConsumerTestData/AMF/TestAmf"
	"gofree5gc/lib/MongoDBLibrary"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/amf/amf_consumer"
	"gofree5gc/src/amf/logger"
	"gofree5gc/src/pcf/pcf_service"
	"testing"
	"time"
)

const amPolicyDataColl = "policyData.ues.amData"

var filterUeIdOnly bson.M

func pcfInit() {
	flags := flag.FlagSet{}
	c := cli.NewContext(nil, &flags, nil)
	pcf := &pcf_service.PCF{}
	pcf.Initialize(c)
	go pcf.Start()
	time.Sleep(100 * time.Millisecond)
}

func insertDefaultAmPolicyToDb(ueId string) {
	amPolicyData := models.AmPolicyData{
		SubscCats: []string{
			"free5gc",
		},
	}
	filterUeIdOnly = bson.M{"ueId": ueId}
	amPolicyDataBsonM := toBsonM(amPolicyData)
	amPolicyDataBsonM["ueId"] = ueId
	MongoDBLibrary.RestfulAPIPutOne(amPolicyDataColl, filterUeIdOnly, amPolicyDataBsonM)
}

func TestAmPolicyControlCreate(t *testing.T) {

	logger.SetLogLevel(logrus.DebugLevel)

	defer MongoDBLibrary.RestfulAPIDeleteMany(amPolicyDataColl, filterUeIdOnly)

	nrfInit()
	pcfInit()
	udrinit()
	insertDefaultAmPolicyToDb("imsi-2089300007487")

	TestAmf.AmfInit()
	TestAmf.UeAttach(models.AccessType__3_GPP_ACCESS)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	ue.PcfUri = "https://localhost:29507"
	ue.AccessAndMobilitySubscriptionData = &models.AccessAndMobilitySubscriptionData{
		RfspIndex: 1,
	}
	problemDetails, err := amf_consumer.AMPolicyControlCreate(ue)
	if err != nil {
		t.Error(err)
	} else if problemDetails != nil {
		t.Errorf("problemDetail: %+v", problemDetails)
	} else {
		t.Logf("Policy Association ID: %+v", ue.PolicyAssociationId)
		t.Logf("AM Policy Association: %+v", ue.AmPolicyAssociation)
	}
}

func TestAmPolicyControlUpdate(t *testing.T) {

	defer MongoDBLibrary.RestfulAPIDeleteMany(amPolicyDataColl, filterUeIdOnly)

	nrfInit()
	pcfInit()
	udrinit()
	insertDefaultAmPolicyToDb("imsi-2089300007487")

	TestAmf.AmfInit()
	TestAmf.UeAttach(models.AccessType__3_GPP_ACCESS)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	ue.PcfUri = "https://localhost:29507"
	ue.AccessAndMobilitySubscriptionData = &models.AccessAndMobilitySubscriptionData{
		RfspIndex: 1,
	}

	// Create an AM Policy Association
	problemDetails, err := amf_consumer.AMPolicyControlCreate(ue)
	if err != nil {
		t.Error(err)
	} else if problemDetails != nil {
		t.Errorf("problemDetail: %+v", problemDetails)
	} else {
		t.Logf("Policy Association ID: %+v", ue.PolicyAssociationId)
		t.Logf("AM Policy Association: %+v", ue.AmPolicyAssociation)
	}

	updateRequest := models.PolicyAssociationUpdateRequest{
		Triggers: []models.RequestTrigger{
			models.RequestTrigger_RFSP_CH,
		},
		Rfsp: 2,
	}

	// Update Rfsp of AM Policy
	problemDetails, err = amf_consumer.AMPolicyControlUpdate(ue, updateRequest)
	if err != nil {
		t.Error(err)
	} else if problemDetails != nil {
		t.Errorf("problemDetail: %+v", problemDetails)
	} else {
		t.Logf("AM Policy Association: %+v", ue.AmPolicyAssociation)
	}
}

func TestAmPolicyControlDelete(t *testing.T) {

	defer MongoDBLibrary.RestfulAPIDeleteMany(amPolicyDataColl, filterUeIdOnly)

	nrfInit()
	pcfInit()
	udrinit()
	insertDefaultAmPolicyToDb("imsi-2089300007487")

	TestAmf.AmfInit()
	TestAmf.UeAttach(models.AccessType__3_GPP_ACCESS)
	ue := TestAmf.TestAmf.UePool["imsi-2089300007487"]

	ue.PcfUri = "https://localhost:29507"
	ue.AccessAndMobilitySubscriptionData = &models.AccessAndMobilitySubscriptionData{
		RfspIndex: 1,
	}

	// Create an AM Policy Association
	problemDetails, err := amf_consumer.AMPolicyControlCreate(ue)
	if err != nil {
		t.Error(err)
	} else if problemDetails != nil {
		t.Logf("problemDetail: %+v", problemDetails)
	} else {
		t.Logf("Policy Association ID: %+v", ue.PolicyAssociationId)
		t.Logf("AM Policy Association: %+v", ue.AmPolicyAssociation)
	}

	// Delete AM Policy Association
	problemDetails, err = amf_consumer.AMPolicyControlDelete(ue)
	if err != nil {
		t.Error(err)
	} else if problemDetails != nil {
		t.Errorf("problemDetail: %+v", problemDetails)
	} else {
		t.Logf("AM Policy Control delete success")
	}
}
