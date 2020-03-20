/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository_test

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"gofree5gc/src/udr/logger"
	"net/http"
	"testing"

	"gofree5gc/lib/Nudr_DataRepository"
	"gofree5gc/lib/openapi/models"

	"github.com/antihax/optional"

	"github.com/google/go-cmp/cmp"
)

// AmfContext3gpp - To modify the AMF context data of a UE using 3gpp access in the UDR
func TestAmfContext3gpp(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.contextData.amf3gppAccess")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.Amf3GppAccessRegistration{
		AmfInstanceId:    "AmfInstanceId_test",
		ImsVoPs:          models.ImsVoPs_HOMOGENEOUS_SUPPORT,
		DeregCallbackUri: "DeregCallbackUri_test",
		Guami: &models.Guami{
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			},
			AmfId: "1",
		},
		RatType: models.RatType_NR,
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check patch data (Use RESTful GET)
		var queryAmfContext3gppParamOpts Nudr_DataRepository.QueryAmfContext3gppParamOpts
		amfNon3GppAccessRegistration, res, err := client.AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp(context.TODO(), ueId, &queryAmfContext3gppParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if cmp.Equal(testData, amfNon3GppAccessRegistration, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				amfNon3GppAccessRegistration, testData)
		}
	}

	patchItemArray := []models.PatchItem{
		{
			Op:    models.PatchOperation_REPLACE,
			Path:  "/amfInstanceId",
			Value: "AAA",
		},
	}
	patchData := models.Amf3GppAccessRegistration{
		AmfInstanceId:    "AAA",
		ImsVoPs:          models.ImsVoPs_HOMOGENEOUS_SUPPORT,
		DeregCallbackUri: "DeregCallbackUri_test",
		Guami: &models.Guami{
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			},
			AmfId: "1",
		},
		RatType: models.RatType_NR,
	}

	{
		// Patch data (Use RESTful PATCH)
		res, err := client.AMF3GPPAccessRegistrationDocumentApi.AmfContext3gpp(context.TODO(), ueId, patchItemArray)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	}

	{
		// Check patch data (Use RESTful GET)
		var queryAmfContext3gppParamOpts Nudr_DataRepository.QueryAmfContext3gppParamOpts
		amf3GppAccessRegistration, res, err := client.AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp(context.TODO(), ueId, &queryAmfContext3gppParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(patchData, amf3GppAccessRegistration, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				amf3GppAccessRegistration, patchData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// CreateAmfContext3gpp - To store the AMF context data of a UE using 3gpp access in the UDR
func TestCreateAmfContext3gpp(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.contextData.amf3gppAccess")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.Amf3GppAccessRegistration{
		AmfInstanceId:    "AmfInstanceId_test",
		ImsVoPs:          models.ImsVoPs_HOMOGENEOUS_SUPPORT,
		DeregCallbackUri: "DeregCallbackUri_test",
		Guami: &models.Guami{
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			},
			AmfId: "1",
		},
		RatType: models.RatType_NR,
	}

	{
		// Insert test data (Use RESTful PUT)
		var createAmfContext3gppParamOpts Nudr_DataRepository.CreateAmfContext3gppParamOpts
		createAmfContext3gppParamOpts.Amf3GppAccessRegistration = optional.NewInterface(testData)
		res, err := client.AMF3GPPAccessRegistrationDocumentApi.CreateAmfContext3gpp(context.TODO(), ueId, &createAmfContext3gppParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	}

	{
		// Check test data (Use RESTful GET)
		var queryAmfContext3gppParamOpts Nudr_DataRepository.QueryAmfContext3gppParamOpts
		amf3GppAccessRegistration, res, err := client.AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp(context.TODO(), ueId, &queryAmfContext3gppParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, amf3GppAccessRegistration, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				amf3GppAccessRegistration, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// QueryAmfContext3gpp - Retrieves the AMF context data of a UE using 3gpp access
func TestQueryAmfContext3gpp(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.contextData.amf3gppAccess")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.Amf3GppAccessRegistration{
		AmfInstanceId:    "AmfInstanceId_test",
		ImsVoPs:          models.ImsVoPs_HOMOGENEOUS_SUPPORT,
		DeregCallbackUri: "DeregCallbackUri_test",
		Guami: &models.Guami{
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			},
			AmfId: "1",
		},
		RatType: models.RatType_NR,
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var queryAmfContext3gppParamOpts Nudr_DataRepository.QueryAmfContext3gppParamOpts
		amf3GppAccessRegistration, res, err := client.AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp(context.TODO(), ueId, &queryAmfContext3gppParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, amf3GppAccessRegistration, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				amf3GppAccessRegistration, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}
