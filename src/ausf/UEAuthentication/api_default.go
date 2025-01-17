/*
 * AUSF API
 *
 * OpenAPI specification for AUSF
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package UEAuthentication

import (
	"github.com/gin-gonic/gin"
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/ausf/ausf_handler"
	"gofree5gc/src/ausf/ausf_handler/ausf_message"
	"gofree5gc/src/ausf/logger"
	// "fmt"
)

// EapAuthMethod -
func EapAuthMethod(c *gin.Context) {
	var eapSessionReq models.EapSession

	err := c.ShouldBindJSON(&eapSessionReq)
	if err != nil {
		logger.Auth5gAkaComfirmLog.Errorln(err)
	}

	req := http_wrapper.NewRequest(c.Request, eapSessionReq)
	req.Params["authCtxId"] = c.Param("authCtxId")

	handlerMsg := ausf_message.NewHandlerMessage(ausf_message.EventEapAuthComfirm, req)
	ausf_handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// UeAuthenticationsAuthCtxId5gAkaConfirmationPut -
func UeAuthenticationsAuthCtxId5gAkaConfirmationPut(c *gin.Context) {
	var confirmationData models.ConfirmationData

	err := c.ShouldBindJSON(&confirmationData)
	if err != nil {
		logger.Auth5gAkaComfirmLog.Errorln(err)
	}

	req := http_wrapper.NewRequest(c.Request, confirmationData)
	req.Params["authCtxId"] = c.Param("authCtxId")

	handlerMsg := ausf_message.NewHandlerMessage(ausf_message.EventAuth5gAkaComfirm, req)
	ausf_handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// UeAuthenticationsPost -
func UeAuthenticationsPost(c *gin.Context) {
	var authInfo models.AuthenticationInfo

	err := c.ShouldBindJSON(&authInfo)
	if err != nil {
		logger.UeAuthPostLog.Errorln(err)
	}

	req := http_wrapper.NewRequest(c.Request, authInfo)

	handlerMsg := ausf_message.NewHandlerMessage(ausf_message.EventUeAuthPost, req)
	ausf_handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	HTTPRespHeader := rsp.HTTPResponse.Header
	for k, v := range HTTPRespHeader {
		c.Header(k, v[0])
	}
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
