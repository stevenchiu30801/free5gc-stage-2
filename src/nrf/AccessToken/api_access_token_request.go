/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AccessToken

import (
	"gofree5gc/lib/openapi/models"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
)

// AccessTokenRequest - Access Token Request
func AccessTokenRequest(c *gin.Context) {
	var accessTokenReq models.AccessTokenReq

	accessTokenReqBindErr := c.Bind(&accessTokenReq)
	if accessTokenReqBindErr != nil {
		log.Println("accessTokenReqBindErr", accessTokenReqBindErr)
	}

	//log.Printf("%+v", accessTokenReq.RequesterPlmn)
	//log.Printf("%+v", accessTokenReq.TargetPlmn)

	// Param of AccessTokenRsp
	var expiration int32 = 1000
	var scope = accessTokenReq.Scope
	var tokenType = "Bearer"

	// Create AccessToken
	var accessTokenClaims = models.AccessTokenClaims{
		"1234567",                         // TODO: NF instance id of the NRF
		accessTokenReq.NfInstanceId,       // nfInstanceId of service consumer
		accessTokenReq.TargetNfInstanceId, // nfInstanceId of service producer
		accessTokenReq.Scope,              // TODO: the name of the NF services for which the access_token is authorized for use
		expiration,
		jwt.StandardClaims{},
	}

	mySigningKey := []byte("NRF") // AllYourBase
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, SignedStringErr := token.SignedString(mySigningKey)
	if SignedStringErr != nil {
		log.Println(SignedStringErr)
	}

	rep := &models.AccessTokenRsp{
		AccessToken: accessToken,
		TokenType:   tokenType,
		ExpiresIn:   expiration,
		Scope:       scope,
	}

	c.JSON(http.StatusOK, rep)
}
