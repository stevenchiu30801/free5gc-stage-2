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

	"github.com/google/go-cmp/cmp"
)

// AmfContext3gpp - To modify operator specific data of a UE
func TestPatchOperSpecData(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.operatorSpecificData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.OperatorSpecificDataContainer{
		StringTypeElements: map[string]string{
			"free5gc": "test",
		},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check patch data (Use RESTful GET)
		var queryOperSpecDataParamOpts Nudr_DataRepository.QueryOperSpecDataParamOpts
		operatorSpecificDataContainer, res, err := client.OperatorSpecificDataContainerDocumentApi.QueryOperSpecData(context.TODO(), ueId, &queryOperSpecDataParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if cmp.Equal(testData, operatorSpecificDataContainer, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				operatorSpecificDataContainer, testData)
		}
	}

	patchItemArray := []models.PatchItem{
		{
			Op:   models.PatchOperation_REPLACE,
			Path: "/StringTypeElements",
			Value: map[string]string{
				"free5gc": "test1",
			},
		},
	}
	patchData := models.OperatorSpecificDataContainer{
		StringTypeElements: map[string]string{
			"free5gc": "test1",
		},
	}

	{
		// Patch data (Use RESTful PATCH)
		res, err := client.OperatorSpecificDataContainerDocumentApi.PatchOperSpecData(context.TODO(), ueId, patchItemArray)
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
		var queryOperSpecDataParamOpts Nudr_DataRepository.QueryOperSpecDataParamOpts
		operatorSpecificDataContainer, res, err := client.OperatorSpecificDataContainerDocumentApi.QueryOperSpecData(context.TODO(), ueId, &queryOperSpecDataParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(patchData, operatorSpecificDataContainer, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				operatorSpecificDataContainer, patchData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})
}

// QueryOperSpecData - Retrieves the operator specific data of a UE
func TestQueryOperSpecData(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.operatorSpecificData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.OperatorSpecificDataContainer{
		StringTypeElements: map[string]string{
			"free5gc": "test",
		},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var queryOperSpecDataParamOpts Nudr_DataRepository.QueryOperSpecDataParamOpts
		operatorSpecificDataContainer, res, err := client.OperatorSpecificDataContainerDocumentApi.QueryOperSpecData(context.TODO(), ueId, &queryOperSpecDataParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, operatorSpecificDataContainer, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				operatorSpecificDataContainer, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}
