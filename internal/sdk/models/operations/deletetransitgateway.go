// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

var DeleteTransitGatewayServerList = []string{
	"https://global.api.konghq.com/",
}

type DeleteTransitGatewayRequest struct {
	// The network to operate on.
	NetworkID string `pathParam:"style=simple,explode=false,name=networkId"`
	// The ID of the transit gateway to operate on.
	TransitGatewayID string `pathParam:"style=simple,explode=false,name=transitGatewayId"`
}

func (o *DeleteTransitGatewayRequest) GetNetworkID() string {
	if o == nil {
		return ""
	}
	return o.NetworkID
}

func (o *DeleteTransitGatewayRequest) GetTransitGatewayID() string {
	if o == nil {
		return ""
	}
	return o.TransitGatewayID
}

type DeleteTransitGatewayResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteTransitGatewayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTransitGatewayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTransitGatewayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteTransitGatewayResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *DeleteTransitGatewayResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *DeleteTransitGatewayResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *DeleteTransitGatewayResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
