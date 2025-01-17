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
	"fmt"
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/udm/logger"
	"gofree5gc/src/udm/udm_handler"
	"gofree5gc/src/udm/udm_handler/udm_message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Modify - modify the subscription
func Modify(c *gin.Context) {
	fmt.Println("asdasd")

	var sdmSubsModificationReq models.SdmSubsModification
	err := c.ShouldBindJSON(&sdmSubsModificationReq)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.Handlelog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, sdmSubsModificationReq)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	handleMsg := udm_message.NewHandlerMessage(udm_message.EventModify, req)
	udm_handler.SendMessage(handleMsg)

	rsp := <-handleMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}

// ModifyForSharedData - modify the subscription
func ModifyForSharedData(c *gin.Context) {

	var sharedDataSubscriptions models.SdmSubsModification
	err := c.ShouldBindJSON(&sharedDataSubscriptions)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.SdmLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, sharedDataSubscriptions)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	handleMsg := udm_message.NewHandlerMessage(udm_message.EventModifyForSharedData, req)
	udm_handler.SendMessage(handleMsg)

	rsp := <-handleMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
