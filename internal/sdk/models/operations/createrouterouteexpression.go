// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateRouteRouteExpressionRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the new Route for creation
	RouteExpression shared.RouteExpressionInput `request:"mediaType=application/json"`
}

func (o *CreateRouteRouteExpressionRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateRouteRouteExpressionRequest) GetRouteExpression() shared.RouteExpressionInput {
	if o == nil {
		return shared.RouteExpressionInput{}
	}
	return o.RouteExpression
}

type CreateRouteRouteExpressionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Route
	RouteExpression *shared.RouteExpression
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateRouteRouteExpressionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRouteRouteExpressionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRouteRouteExpressionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRouteRouteExpressionResponse) GetRouteExpression() *shared.RouteExpression {
	if o == nil {
		return nil
	}
	return o.RouteExpression
}

func (o *CreateRouteRouteExpressionResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
