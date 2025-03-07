// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateRoutebyheaderPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID      string                          `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	RouteByHeaderPlugin shared.RouteByHeaderPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateRoutebyheaderPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRoutebyheaderPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateRoutebyheaderPluginRequest) GetRouteByHeaderPlugin() shared.RouteByHeaderPluginInput {
	if o == nil {
		return shared.RouteByHeaderPluginInput{}
	}
	return o.RouteByHeaderPlugin
}

type UpdateRoutebyheaderPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// RouteByHeader plugin
	RouteByHeaderPlugin *shared.RouteByHeaderPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateRoutebyheaderPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRoutebyheaderPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRoutebyheaderPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRoutebyheaderPluginResponse) GetRouteByHeaderPlugin() *shared.RouteByHeaderPlugin {
	if o == nil {
		return nil
	}
	return o.RouteByHeaderPlugin
}

func (o *UpdateRoutebyheaderPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
