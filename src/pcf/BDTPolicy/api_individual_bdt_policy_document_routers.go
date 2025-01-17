/*
 * Npcf_BDTPolicyControl Service API
 *
 * The Npcf_BDTPolicyControl Service is used by an NF service consumer to retrieve background data transfer policies from the PCF and to update the PCF with the background data transfer policy selected by the NF service consumer.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package BDTPolicy

import (
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/pcf/pcf_handler/pcf_message"

	"github.com/gin-gonic/gin"
)

// GetBDTPolicy - Read an Individual BDT policy
func GetBDTPolicy(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["bdtPolicyId"] = c.Params.ByName("bdtPolicyId")
	channelMsg := pcf_message.NewHttpChannelMessage(pcf_message.EventBDTPolicyGet, req)

	pcf_message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// UpdateBDTPolicy - Update an Individual BDT policy
func UpdateBDTPolicy(c *gin.Context) {
	var bdtPolicyDataPatch models.BdtPolicyDataPatch
	c.ShouldBindJSON(&bdtPolicyDataPatch)

	req := http_wrapper.NewRequest(c.Request, bdtPolicyDataPatch)
	req.Params["bdtPolicyId"] = c.Params.ByName("bdtPolicyId")
	channelMsg := pcf_message.NewHttpChannelMessage(pcf_message.EventBDTPolicyUpdate, req)

	pcf_message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
