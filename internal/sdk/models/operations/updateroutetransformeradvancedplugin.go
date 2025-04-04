// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateRoutetransformeradvancedPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                 string                                `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	RouteTransformerAdvancedPlugin shared.RouteTransformerAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateRoutetransformeradvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRoutetransformeradvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateRoutetransformeradvancedPluginRequest) GetRouteTransformerAdvancedPlugin() shared.RouteTransformerAdvancedPlugin {
	if o == nil {
		return shared.RouteTransformerAdvancedPlugin{}
	}
	return o.RouteTransformerAdvancedPlugin
}

type UpdateRoutetransformeradvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// RouteTransformerAdvanced plugin
	RouteTransformerAdvancedPlugin *shared.RouteTransformerAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateRoutetransformeradvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRoutetransformeradvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRoutetransformeradvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRoutetransformeradvancedPluginResponse) GetRouteTransformerAdvancedPlugin() *shared.RouteTransformerAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.RouteTransformerAdvancedPlugin
}

func (o *UpdateRoutetransformeradvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
