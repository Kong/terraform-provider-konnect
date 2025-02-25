// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetMockingPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetMockingPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *GetMockingPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetMockingPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Mocking plugin
	MockingPlugin *shared.MockingPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetMockingPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMockingPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMockingPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMockingPluginResponse) GetMockingPlugin() *shared.MockingPlugin {
	if o == nil {
		return nil
	}
	return o.MockingPlugin
}

func (o *GetMockingPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
