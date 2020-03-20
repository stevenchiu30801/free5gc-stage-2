/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SubscriberDataManagement

import (
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/src/udm/udm_handler"
	"gofree5gc/src/udm/udm_handler/udm_message"

	"github.com/gin-gonic/gin"
)

// UnsubscribeForSharedData - unsubscribe from notifications for shared data
func UnsubscribeForSharedData(c *gin.Context) {

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	handleMsg := udm_message.NewHandlerMessage(udm_message.EventUnsubscribeForSharedData, req)
	udm_handler.SendMessage(handleMsg)

	rsp := <-handleMsg.ResponseChan
	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}