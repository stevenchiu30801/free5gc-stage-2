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
	"github.com/antihax/optional"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"gofree5gc/lib/Nudr_DataRepository"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/udr/logger"
	"net/http"
	"testing"
	"time"
)

// CreateAuthenticationStatus - To store the Authentication Status data of a UE
func TestCreateAuthenticationStatus(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.authenticationData.authenticationStatus")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	var timeNow = time.Now()
	testData := models.AuthEvent{
		NfInstanceId:       "NfInstanceId",
		Success:            true,
		TimeStamp:          &(timeNow),
		AuthType:           models.AuthType__5_G_AKA,
		ServingNetworkName: "free5GC",
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Insert test data (Use RESTful PUT)
		var createAuthenticationStatusParamOpts Nudr_DataRepository.CreateAuthenticationStatusParamOpts
		createAuthenticationStatusParamOpts.AuthEvent = optional.NewInterface(testData)
		res, err := client.AuthenticationStatusDocumentApi.CreateAuthenticationStatus(context.TODO(), ueId, &createAuthenticationStatusParamOpts)
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
		var queryAuthenticationStatusParamOpts Nudr_DataRepository.QueryAuthenticationStatusParamOpts
		authEvent, res, err := client.AuthEventDocumentApi.QueryAuthenticationStatus(context.TODO(), ueId, &queryAuthenticationStatusParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, authEvent, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				authEvent, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// QueryAuthenticationStatus - Retrieves the Authentication Status of a UE
func TestQueryAuthenticationStatus(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.authenticationData.authenticationStatus")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	var timeNow = time.Now()
	testData := models.AuthEvent{
		NfInstanceId:       "NfInstanceId",
		Success:            true,
		TimeStamp:          &(timeNow),
		AuthType:           models.AuthType__5_G_AKA,
		ServingNetworkName: "free5GC",
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var queryAuthenticationStatusParamOpts Nudr_DataRepository.QueryAuthenticationStatusParamOpts
		authEvent, res, err := client.AuthEventDocumentApi.QueryAuthenticationStatus(context.TODO(), ueId, &queryAuthenticationStatusParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, authEvent, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				authEvent, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}
