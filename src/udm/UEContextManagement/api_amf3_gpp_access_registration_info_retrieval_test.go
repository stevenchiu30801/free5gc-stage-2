package UEContextManagement_test

import (
	"context"
	"fmt"
	Nudm_UECM_Client "gofree5gc/lib/Nudm_UEContextManagement"
	"gofree5gc/lib/http2_util"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/lib/path_util"
	Nudm_UECM_Server "gofree5gc/src/udm/UEContextManagement"
	"gofree5gc/src/udm/logger"
	"gofree5gc/src/udm/udm_context"
	"gofree5gc/src/udm/udm_handler"
	"net/http"
	"testing"

	"github.com/antihax/optional"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// GetAmf3gppAccess - retrieve the AMF registration for 3GPP access information
func TestGetAmf3gppAccess(t *testing.T) {
	go func() { // udm server
		router := gin.Default()
		Nudm_UECM_Server.AddService(router)

		udmLogPath := path_util.Gofree5gcPath("gofree5gc/udmsslkey.log")
		udmPemPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udm.key")

		server, err := http2_util.NewServer(":29503", udmLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
			assert.True(t, err == nil)
		}
	}()
	// udm_util.testInitUdmConfig()
	udm_context.TestInit()
	go udm_handler.Handle()

	go func() { // fake udr server
		router := gin.Default()

		router.GET("/nudr-dr/v1/subscription-data/:ueId/context-data/amf-3gpp-access", func(c *gin.Context) {
			ueId := c.Param("ueId")
			supportedFeatures := c.Query("supported-features")
			fmt.Println("==========AMF 3Gpp-access Registration Info Retrieval==========")
			fmt.Println("ueId: ", ueId)
			fmt.Println("supportedFeatures: ", supportedFeatures)

			var testAmf3GppAccessRegistration models.Amf3GppAccessRegistration
			testAmf3GppAccessRegistration.AmfInstanceId = "test001"
			c.JSON(http.StatusOK, testAmf3GppAccessRegistration)
		})

		udrLogPath := path_util.Gofree5gcPath("gofree5gc/udrsslkey.log")
		udrPemPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udr.pem")
		udrKeyPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udr.key")

		server, err := http2_util.NewServer(":29504", udrLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
			assert.True(t, err == nil)
		}
	}()

	udm_context.Init()
	cfg := Nudm_UECM_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	clientAPI := Nudm_UECM_Client.NewAPIClient(cfg)

	ueId := "UECM1234"
	var getParamOpts Nudm_UECM_Client.GetParamOpts
	getParamOpts.SupportedFeatures = optional.NewString("test_3gpp_SupportedFeatures")
	amf3GppAccessRegistration, resp, err := clientAPI.AMF3GppAccessRegistrationInfoRetrievalApi.Get(context.Background(), ueId, &getParamOpts)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
		fmt.Println("amf3GppAccessRegistration: ", amf3GppAccessRegistration)
	}
}
