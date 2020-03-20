package amf_consumer_test

import (
	"go.mongodb.org/mongo-driver/bson"
	"gofree5gc/lib/CommonConsumerTestData/AMF/TestAmf"
	"gofree5gc/lib/MongoDBLibrary"
	"gofree5gc/src/amf/amf_consumer"
	"testing"
	"time"
)

func TestRegisterNFInstance(t *testing.T) {

	nrfInit()

	time.Sleep(200 * time.Millisecond)
	MongoDBLibrary.RestfulAPIDeleteMany("NfProfile", bson.M{})

	// Init AMF
	TestAmf.AmfInit()

	time.Sleep(100 * time.Millisecond)

	nfprofile, err := amf_consumer.BuildNFInstance(TestAmf.TestAmf)
	if err != nil {
		t.Error(err.Error())
	}

	uri, nfId, err1 := amf_consumer.SendRegisterNFInstance(TestAmf.TestAmf.NrfUri, TestAmf.TestAmf.NfId, nfprofile)
	if err1 != nil {
		t.Error(err1.Error())
	} else {
		t.Logf("Retrieve NfInstanceId: %s", nfId)
		TestAmf.Config.Dump(uri)
	}
}
