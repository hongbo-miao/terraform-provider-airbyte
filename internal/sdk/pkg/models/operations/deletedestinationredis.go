// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDestinationRedisRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type DeleteDestinationRedisResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}