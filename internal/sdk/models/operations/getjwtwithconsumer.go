// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type GetJwtWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the JWT to lookup
	JWTID string `pathParam:"style=simple,explode=false,name=JWTId"`
}

func (o *GetJwtWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetJwtWithConsumerRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *GetJwtWithConsumerRequest) GetJWTID() string {
	if o == nil {
		return ""
	}
	return o.JWTID
}

type GetJwtWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched JWT
	Jwt *shared.Jwt
}

func (o *GetJwtWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetJwtWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetJwtWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetJwtWithConsumerResponse) GetJwt() *shared.Jwt {
	if o == nil {
		return nil
	}
	return o.Jwt
}