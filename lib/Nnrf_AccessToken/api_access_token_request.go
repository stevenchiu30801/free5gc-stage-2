/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_AccessToken

import (
	"gofree5gc/lib/openapi/common"
	"gofree5gc/lib/openapi/models"

	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type AccessTokenRequestApiService service

/*
AccessTokenRequestApiService Access Token Request
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param grantType
 * @param nfInstanceId
 * @param scope
 * @param optional nil or *AccessTokenRequestParamOpts - Optional Parameters:
 * @param "NfType" (optional.Interface of models.NfType) -
 * @param "TargetNfType" (optional.Interface of models.NfType) -
 * @param "TargetNfInstanceId" (optional.Interface of string) -
 * @param "RequesterPlmn" (optional.Interface of models.PlmnId) -
 * @param "TargetPlmn" (optional.Interface of models.PlmnId) -
@return models.AccessTokenRsp
*/

type AccessTokenRequestParamOpts struct {
	NfType             optional.Interface
	TargetNfType       optional.Interface
	TargetNfInstanceId optional.Interface
	RequesterPlmn      optional.Interface
	TargetPlmn         optional.Interface
}

func (a *AccessTokenRequestApiService) AccessTokenRequest(ctx context.Context, grantType string, nfInstanceId string, scope string, localVarOptionals *AccessTokenRequestParamOpts) (models.AccessTokenRsp, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.AccessTokenRsp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	localVarFormParams.Add("grant_type", common.ParameterToString(grantType, ""))
	localVarFormParams.Add("nfInstanceId", common.ParameterToString(nfInstanceId, ""))
	if localVarOptionals != nil && localVarOptionals.NfType.IsSet() {
		localVarFormParams.Add("nfType", common.ParameterToString(localVarOptionals.NfType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetNfType.IsSet() {
		localVarFormParams.Add("targetNfType", common.ParameterToString(localVarOptionals.TargetNfType.Value(), ""))
	}
	localVarFormParams.Add("scope", common.ParameterToString(scope, ""))
	if localVarOptionals != nil && localVarOptionals.TargetNfInstanceId.IsSet() {
		localVarFormParams.Add("targetNfInstanceId", common.ParameterToString(localVarOptionals.TargetNfInstanceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RequesterPlmn.IsSet() {
		localVarFormParams.Add("requesterPlmn", common.ParameterToString(localVarOptionals.RequesterPlmn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetPlmn.IsSet() {
		localVarFormParams.Add("targetPlmn", common.ParameterToString(localVarOptionals.TargetPlmn.Value(), ""))
	}

	r, err := common.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := common.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := common.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = common.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		var v models.AccessTokenErr
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		return localVarReturnValue, localVarHTTPResponse, common.ReportError("%d is not a valid status code in AccessTokenRequest", localVarHTTPResponse.StatusCode)
	}
}