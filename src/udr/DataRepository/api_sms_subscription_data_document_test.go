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
	"gofree5gc/src/udr/logger"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"gofree5gc/lib/Nudr_DataRepository"
	"gofree5gc/lib/openapi/models"
)

// QuerySmsData - Retrieves the SMS subscription data of a UE
func TestQuerySmsData(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.provisionedData.smsData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789", "servingPlmnId": "20893"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	servingPlmnId := "20893"
	testData := models.SmsSubscriptionData{
		SharedSmsSubsDataId: []string{"a", "b"},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	insertTestData["servingPlmnId"] = servingPlmnId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var querySmsDataParamOpts Nudr_DataRepository.QuerySmsDataParamOpts
		smsSubscriptionData, res, err := client.SMSSubscriptionDataDocumentApi.QuerySmsData(context.TODO(), ueId, servingPlmnId, &querySmsDataParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, smsSubscriptionData, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				smsSubscriptionData, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}
