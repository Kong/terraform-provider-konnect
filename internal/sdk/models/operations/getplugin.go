// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *GetPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched Plugin
	Plugin *shared.Plugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPluginResponse) GetPlugin() *shared.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}

func (o *GetPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
