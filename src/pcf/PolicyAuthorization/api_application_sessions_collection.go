/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package PolicyAuthorization

import (
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/pcf/logger"
	"gofree5gc/src/pcf/pcf_handler/pcf_message"
	"gofree5gc/src/pcf/pcf_util"

	"github.com/gin-gonic/gin"
)

// PostAppSessions - Creates a new Individual Application Session Context resource
func PostAppSessions(c *gin.Context) {
	var appSessionContext models.AppSessionContext
	err := c.ShouldBindJSON(&appSessionContext)
	if err != nil {
		rsp := pcf_util.GetProblemDetail("Malformed request syntax", pcf_util.ERROR_INITIAL_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}
	reqData := appSessionContext.AscReqData
	if reqData == nil || reqData.SuppFeat == "" || reqData.NotifUri == "" {
		// Check Mandatory IEs
		rsp := pcf_util.GetProblemDetail("Errorneous/Missing Mandotory IE", pcf_util.ERROR_INITIAL_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, appSessionContext)
	channelMsg := pcf_message.NewHttpChannelMessage(pcf_message.EventPostAppSessions, req)

	pcf_message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse

	for key, val := range HTTPResponse.Header {
		c.Header(key, val[0])
	}
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
