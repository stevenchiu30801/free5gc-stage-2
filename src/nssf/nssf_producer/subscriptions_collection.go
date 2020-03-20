/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nssf_producer

import (
	"net/http"

	"gofree5gc/lib/http_wrapper"
	. "gofree5gc/lib/openapi/models"
	"gofree5gc/src/nssf/logger"
	"gofree5gc/src/nssf/nssf_handler/nssf_message"
)

// NSSAIAvailabilityPost - Creates subscriptions for notification about updates to NSSAI availability information
func NSSAIAvailabilityPost(responseChan chan nssf_message.HandlerResponseMessage, n NssfEventSubscriptionCreateData) {

	logger.Nssaiavailability.Infof("Request received - NSSAIAvailabilityPost")

	var (
		isValidRequest bool = true
		status         int
		s              NssfEventSubscriptionCreatedData
		d              ProblemDetails
	)

	// TODO: If NF consumer is not authorized to update NSSAI availability, return ProblemDetails with code 403 Forbidden

	if isValidRequest {
		status = subscriptionPost(n, &s, &d)
	}

	if status == http.StatusCreated {
		responseChan <- nssf_message.HandlerResponseMessage{
			HttpResponse: &http_wrapper.Response{
				Header: nil,
				Status: status,
				Body:   s,
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
