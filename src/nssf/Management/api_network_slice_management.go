/*
 * NSSF Management
 *
 * NSSF Management Service
 */

package Management

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gofree5gc/lib/http_wrapper"
	. "gofree5gc/lib/openapi/models"
	"gofree5gc/src/nssf/nssf_handler"
	"gofree5gc/src/nssf/nssf_handler/nssf_message"
	. "gofree5gc/src/nssf/plugin"
	"gofree5gc/src/nssf/util"
)

func ApiNetworkSliceManagementPost(c *gin.Context) {
	var request NetworkSliceManagementDocument
	err := c.ShouldBindJSON(&request)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		d := ProblemDetails{
			Title:  util.MALFORMED_REQUEST,
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		c.JSON(http.StatusBadRequest, d)
		return
	}
	req := http_wrapper.NewRequest(c.Request, request)

	message := nssf_message.NewMessage(nssf_message.ManagementPost, req)

	nssf_handler.SendMessage(message)
	rsp := <-message.ResponseChan

	httpResponse := rsp.HttpResponse
	c.JSON(httpResponse.Status, httpResponse.Body)
}
