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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationAmfNon3gppAccess(t *testing.T) {
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

	udm_context.TestInit()
	go udm_handler.Handle()

	go func() { // fake udr server
		router := gin.Default()

		router.PUT("/nudr-dr/v1/subscription-data/:ueId/context-data/amf-non-3gpp-access", func(c *gin.Context) {
			ueId := c.Param("ueId")
			fmt.Println("==========AMF registration for non-3GPP access==========")
			fmt.Println("ueId: ", ueId)

			var amfNon3GppAccessRegistration models.AmfNon3GppAccessRegistration
			if err := c.ShouldBindJSON(&amfNon3GppAccessRegistration); err != nil {
				fmt.Println("fake udr server error")
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			fmt.Println("amfNon3GppAccessRegistration - ", amfNon3GppAccessRegistration.AmfInstanceId)
			c.JSON(http.StatusNoContent, gin.H{})
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
	var amfNon3GppAccessRegistration models.AmfNon3GppAccessRegistration
	amfNon3GppAccessRegistration.AmfInstanceId = "NON_3GPP_PUT_TEST_001"
	_, resp, err := clientAPI.AMFRegistrationForNon3GPPAccessApi.Register(context.Background(), ueId, amfNon3GppAccessRegistration)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
	}
}
