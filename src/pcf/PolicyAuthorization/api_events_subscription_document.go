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

// DeleteEventsSubsc - deletes the Events Subscription subresource
func DeleteEventsSubsc(c *gin.Context) {

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["appSessionId"], _ = c.Params.Get("appSessionId")
	channelMsg := pcf_message.NewHttpChannelMessage(pcf_message.EventDeleteEventsSubsc, req)

	pcf_message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel

	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// UpdateEventsSubsc - creates or modifies an Events Subscription subresource
func UpdateEventsSubsc(c *gin.Context) {
	var eventsSubscReqData models.EventsSubscReqData
	err := c.ShouldBindJSON(&eventsSubscReqData)
	if err != nil {
		rsp := pcf_util.GetProblemDetail("Malformed request syntax", pcf_util.ERROR_REQUEST_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}
	if eventsSubscReqData.Events == nil || eventsSubscReqData.NotifUri == "" {
		rsp := pcf_util.GetProblemDetail("Errorneous/Missing Mandotory IE", pcf_util.ERROR_REQUEST_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, eventsSubscReqData)
	req.Params["appSessionId"], _ = c.Params.Get("appSessionId")
	channelMsg := pcf_message.NewHttpChannelMessage(pcf_message.EventUpdateEventsSubsc, req)

	pcf_message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel

	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
