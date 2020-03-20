package Npcf_Callback

import (
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/pcf/logger"
	"gofree5gc/src/pcf/pcf_handler/pcf_message"

	"github.com/gin-gonic/gin"
)

// Nudr-Notify-smpolicy
func NudrNotify(c *gin.Context) {
	var policyDataChangeNotification models.PolicyDataChangeNotification
	if err := c.ShouldBindJSON(&policyDataChangeNotification); err != nil {
		logger.SMpolicylog.Warnln("Nudr Notify fail error message is : ", err)
	}

	req := http_wrapper.NewRequest(c.Request, policyDataChangeNotification)
	req.Params["ReqURI"] = c.Params.ByName("supi")
	channelMsg := pcf_message.NewHttpChannelMessage(pcf_message.EventSMPolicyNotify, req)

	pcf_message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
