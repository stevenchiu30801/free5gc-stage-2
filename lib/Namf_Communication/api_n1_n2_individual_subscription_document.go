/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_Communication

import (
	"context"
	"fmt"
	"gofree5gc/lib/openapi/common"
	"gofree5gc/lib/openapi/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type N1N2IndividualSubscriptionDocumentApiService service

/*
N1N2IndividualSubscriptionDocumentApiService Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param subscriptionId Subscription Identifier
*/

func (a *N1N2IndividualSubscriptionDocumentApiService) N1N2MessageUnSubscribe(ctx context.Context, ueContextId string, subscriptionId string) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/ue-contexts/{ueContextId}/n1-n2-messages/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueContextId"+"}", fmt.Sprintf("%v", ueContextId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", fmt.Sprintf("%v", subscriptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHttpContentTypes := []string{"application / json"}
	localVarHeaderParams["Content-Type"] = localVarHttpContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := common.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := common.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := common.CallAPI(a.client.cfg, r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	apiError := common.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHttpResponse.Status,
	}
	switch localVarHttpResponse.StatusCode {
	case 204:
		return localVarHttpResponse, err
	case 400:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 411:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 413:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 415:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	default:
		return localVarHttpResponse, common.ReportError("%d is not a valid status code in N1N2MessageUnSubscribe", localVarHttpResponse.StatusCode)
	}
}
