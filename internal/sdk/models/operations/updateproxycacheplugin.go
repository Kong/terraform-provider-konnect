// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateProxycachePluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID   string                        `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	ProxyCachePlugin *shared.ProxyCachePluginInput `request:"mediaType=application/json"`
}

func (o *UpdateProxycachePluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateProxycachePluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateProxycachePluginRequest) GetProxyCachePlugin() *shared.ProxyCachePluginInput {
	if o == nil {
		return nil
	}
	return o.ProxyCachePlugin
}

type UpdateProxycachePluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ProxyCache plugin
	ProxyCachePlugin *shared.ProxyCachePlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateProxycachePluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateProxycachePluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateProxycachePluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateProxycachePluginResponse) GetProxyCachePlugin() *shared.ProxyCachePlugin {
	if o == nil {
		return nil
	}
	return o.ProxyCachePlugin
}

func (o *UpdateProxycachePluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
