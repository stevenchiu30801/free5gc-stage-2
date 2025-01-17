package UEAuthentication_test

import (
	"context"
	"fmt"
	Nudm_UEAU_Client "gofree5gc/lib/Nudm_UEAuthentication"
	"gofree5gc/lib/http2_util"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/lib/path_util"
	Nudm_UEAU_Server "gofree5gc/src/udm/UEAuthentication"
	"gofree5gc/src/udm/logger"
	"gofree5gc/src/udm/udm_context"
	"gofree5gc/src/udm/udm_handler"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConfirmAuth(t *testing.T) {
	go func() { // udm server
		router := gin.Default()
		Nudm_UEAU_Server.AddService(router)

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

		router.PUT("/nudr-dr/v1/subscription-data/:ueId/authentication-data/authentication-status", func(c *gin.Context) {
			ueId := c.Param("ueId")
			fmt.Println("===================================")
			fmt.Println("ueId: ", ueId)
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

	var authEvent models.AuthEvent
	authEvent.AuthType = models.AuthType__5_G_AKA
	authEvent.Success = true
	var now = time.Now()
	authEvent.TimeStamp = &now

	cfg := Nudm_UEAU_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	client := Nudm_UEAU_Client.NewAPIClient(cfg)

	supi := "11223344"
	_, resp, err := client.ConfirmAuthApi.ConfirmAuth(context.TODO(), supi, authEvent)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("status code: ", resp.StatusCode)
	}
}
