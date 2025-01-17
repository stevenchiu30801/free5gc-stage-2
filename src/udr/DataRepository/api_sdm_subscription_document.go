/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/gin-gonic/gin"
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/udr/logger"
	"gofree5gc/src/udr/udr_handler/udr_message"
)

// RemovesdmSubscriptions - Deletes a sdmsubscriptions
func RemovesdmSubscriptions(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["subsId"] = c.Params.ByName("subsId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventRemovesdmSubscriptions, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// Updatesdmsubscriptions - Stores an individual sdm subscriptions of a UE
func Updatesdmsubscriptions(c *gin.Context) {
	var sdmSubscription models.SdmSubscription
	if err := c.ShouldBindJSON(&sdmSubscription); err != nil {
		logger.DataRepoLog.Panic(err.Error())
	}

	req := http_wrapper.NewRequest(c.Request, sdmSubscription)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["subsId"] = c.Params.ByName("subsId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventUpdatesdmsubscriptions, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
