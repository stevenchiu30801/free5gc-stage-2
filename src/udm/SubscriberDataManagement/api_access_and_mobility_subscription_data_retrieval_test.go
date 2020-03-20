/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SubscriberDataManagement_test

import (
	"context"
	"encoding/json"
	"flag"
	"gofree5gc/lib/MongoDBLibrary"
	"gofree5gc/lib/Nudm_SubscriberDataManagement"
	"gofree5gc/lib/http2_util"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/lib/path_util"
	"gofree5gc/src/udm/SubscriberDataManagement"
	"gofree5gc/src/udm/udm_context"
	"gofree5gc/src/udm/udm_handler"
	"gofree5gc/src/udm/udm_service"
	"gofree5gc/src/udr/DataRepository"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/antihax/optional"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/bson"
)

var UDM = &udm_service.UDM{}

var udmContext udm_context.UDMContext

func testtestInitUdmConfig() {
	flag := flag.FlagSet{}
	cli := cli.NewContext(nil, &flag, nil)
	UDM.Initialize(cli)
}

// GetAmData - retrieve a UE's Access and Mobility Subscription Data
func TestGetAmData(t *testing.T) {

	go udm_handler.Handle()

	go func() {
		router := gin.Default()
		SubscriberDataManagement.AddService(router)

		udmLogPath := path_util.Gofree5gcPath("gofree5gc/udrmslkey.log")
		udmPemPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udm.key")

		server, err := http2_util.NewServer(":29503", udmLogPath, router)
		if err == nil && server != nil {
			err := server.ListenAndServeTLS(udmPemPath, udmKeyPath)
			assert.True(t, err == nil)
		}
	}()

	go func() {
		router := gin.Default()
		DataRepository.AddService(router)

		udrLogPath := path_util.Gofree5gcPath("gofree5gc/udrsslkey.log")
		udrPemPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udr.pem")
		udrKeyPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udr.key")

		server, err := http2_util.NewServer(":29504", udrLogPath, router)
		if err == nil && server != nil {
			err := server.ListenAndServeTLS(udrPemPath, udrKeyPath)
			assert.True(t, err == nil)
		}
	}()

	MongoDBLibrary.SetMongoDB("free5gc", "mongodb://localhost:27017")
	Client := MongoDBLibrary.Client

	alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })

	// This option handles slices and maps of any type.
	Opt := cmp.FilterValues(func(x, y interface{}) bool {
		vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
		// fmt.Println(vx.Kind(), "and", vy.Kind())
		return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
			(vx.Kind() == reflect.Map || vx.Kind() == reflect.Slice)
	}, alwaysEqual)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.provisionedData.amData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	configuration := Nudm_SubscriberDataManagement.NewConfiguration()
	configuration.SetBasePath("https://127.0.0.1:29503") // UDM Client
	clientAPI := Nudm_SubscriberDataManagement.NewAPIClient(configuration)

	// Set test data
	ueId := "imsi-0123456789"
	servingPlmnId := "20893"
	testData := models.AccessAndMobilitySubscriptionData{
		UeUsageType: 1,
	}
	tmp, _ := json.Marshal(testData)
	var insertTestData = bson.M{}
	json.Unmarshal(tmp, &insertTestData)
	insertTestData["ueId"] = ueId
	insertTestData["servingPlmnId"] = servingPlmnId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		var getAmDataParamOpts Nudm_SubscriberDataManagement.GetAmDataParamOpts
		getAmDataParamOpts.PlmnId = optional.NewInterface(servingPlmnId)
		supi := ueId

		// Check test data (Use RESTful GET)
		accessAndMobilitySubscriptionData, res, err := clientAPI.AccessAndMobilitySubscriptionDataRetrievalApi.GetAmData(context.TODO(), supi, &getAmDataParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, accessAndMobilitySubscriptionData, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				accessAndMobilitySubscriptionData, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END*/

}