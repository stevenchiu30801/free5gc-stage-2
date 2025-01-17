/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Management

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"gofree5gc/lib/MongoDBLibrary"
	"gofree5gc/src/nrf/UriList"
	"net/http"
	"strconv"
)

// GetNFInstances - Retrieves a collection of NF Instances
func GetNFInstances(c *gin.Context) {

	nfType := c.Query("nf-type")
	//limit := c.Query("limit")
	limit, err := strconv.Atoi(c.Query("limit"))
	fmt.Println("limit.........", limit)
	collName := "UriList"
	filter := bson.M{"nfType": nfType}

	UL := MongoDBLibrary.RestfulAPIGetOne(collName, filter)

	var originalUL UriList.UriList
	err2 := mapstructure.Decode(UL, &originalUL)
	if err2 != nil {
		panic(err)
	}
	nnrfUriListLimit(&originalUL, limit)
	c.JSON(http.StatusOK, originalUL)
}
