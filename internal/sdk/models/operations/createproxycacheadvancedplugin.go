// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateProxycacheadvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID           string                          `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	ProxyCacheAdvancedPlugin shared.ProxyCacheAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *CreateProxycacheadvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateProxycacheadvancedPluginRequest) GetProxyCacheAdvancedPlugin() shared.ProxyCacheAdvancedPlugin {
	if o == nil {
		return shared.ProxyCacheAdvancedPlugin{}
	}
	return o.ProxyCacheAdvancedPlugin
}

type CreateProxycacheadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created ProxyCacheAdvanced plugin
	ProxyCacheAdvancedPlugin *shared.ProxyCacheAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateProxycacheadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateProxycacheadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateProxycacheadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateProxycacheadvancedPluginResponse) GetProxyCacheAdvancedPlugin() *shared.ProxyCacheAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.ProxyCacheAdvancedPlugin
}

func (o *CreateProxycacheadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
