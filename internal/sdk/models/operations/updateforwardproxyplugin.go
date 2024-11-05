// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateForwardproxyPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID           string                           `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateForwardProxyPlugin *shared.CreateForwardProxyPlugin `request:"mediaType=application/json"`
}

func (o *UpdateForwardproxyPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateForwardproxyPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateForwardproxyPluginRequest) GetCreateForwardProxyPlugin() *shared.CreateForwardProxyPlugin {
	if o == nil {
		return nil
	}
	return o.CreateForwardProxyPlugin
}

type UpdateForwardproxyPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ForwardProxy plugin
	ForwardProxyPlugin *shared.ForwardProxyPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateForwardproxyPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateForwardproxyPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateForwardproxyPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateForwardproxyPluginResponse) GetForwardProxyPlugin() *shared.ForwardProxyPlugin {
	if o == nil {
		return nil
	}
	return o.ForwardProxyPlugin
}

func (o *UpdateForwardproxyPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}