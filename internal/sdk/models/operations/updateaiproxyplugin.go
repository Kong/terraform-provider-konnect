// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateAiproxyPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID      string                      `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateAIProxyPlugin *shared.CreateAIProxyPlugin `request:"mediaType=application/json"`
}

func (o *UpdateAiproxyPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateAiproxyPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateAiproxyPluginRequest) GetCreateAIProxyPlugin() *shared.CreateAIProxyPlugin {
	if o == nil {
		return nil
	}
	return o.CreateAIProxyPlugin
}

type UpdateAiproxyPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// AIProxy plugin
	AIProxyPlugin *shared.AIProxyPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateAiproxyPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAiproxyPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAiproxyPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAiproxyPluginResponse) GetAIProxyPlugin() *shared.AIProxyPlugin {
	if o == nil {
		return nil
	}
	return o.AIProxyPlugin
}

func (o *UpdateAiproxyPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
