/*
 * NSSF Management
 *
 * NSSF Management Service
 */

package nssf_producer

import (
	"net/http"

	"gofree5gc/lib/http_wrapper"
	. "gofree5gc/lib/openapi/models"
	"gofree5gc/src/nssf/factory"
	"gofree5gc/src/nssf/logger"
	"gofree5gc/src/nssf/nssf_handler/nssf_message"
	. "gofree5gc/src/nssf/plugin"
)

// ManagementPost - Creates network slice information for management
func ManagementPost(responseChan chan nssf_message.HandlerResponseMessage, n NetworkSliceManagementDocument) {

	logger.Management.Infof("Request received - ManagementPost")

	var (
		status int
		d      ProblemDetails
	)

	status = networkSliceManagementPost(n, &d)

	if status == http.StatusCreated {
		responseChan <- nssf_message.HandlerResponseMessage{
			HttpResponse: &http_wrapper.Response{
				Header: nil,
				Status: status,
				Body:   factory.NssfConfig.Configuration,
			},
		}
	} else {
		responseChan <- nssf_message.HandlerResponseMessage{
			HttpResponse: &http_wrapper.Response{
				Header: nil,
				Status: status,
				Body:   d,
			},
		}
	}
}
