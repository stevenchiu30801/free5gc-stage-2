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
	"github.com/gin-gonic/gin"
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/src/udm/udm_handler"
	"gofree5gc/src/udm/udm_handler/udm_message"
)

// GetSharedData - retrieve shared data
func GetSharedData(c *gin.Context) {

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Query["sharedDataIds"] = c.QueryArray("shared-data-ids")

	handleMsg := udm_message.NewHandlerMessage(udm_message.EventGetSharedData, req)
	udm_handler.SendMessage(handleMsg)

	rsp := <-handleMsg.ResponseChan
	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
