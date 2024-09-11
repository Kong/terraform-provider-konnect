// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateUpstreamtimeoutPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID              string                              `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateUpstreamTimeoutPlugin *shared.CreateUpstreamTimeoutPlugin `request:"mediaType=application/json"`
}

func (o *UpdateUpstreamtimeoutPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateUpstreamtimeoutPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateUpstreamtimeoutPluginRequest) GetCreateUpstreamTimeoutPlugin() *shared.CreateUpstreamTimeoutPlugin {
	if o == nil {
		return nil
	}
	return o.CreateUpstreamTimeoutPlugin
}

type UpdateUpstreamtimeoutPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// UpstreamTimeout plugin
	UpstreamTimeoutPlugin *shared.UpstreamTimeoutPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateUpstreamtimeoutPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUpstreamtimeoutPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUpstreamtimeoutPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUpstreamtimeoutPluginResponse) GetUpstreamTimeoutPlugin() *shared.UpstreamTimeoutPlugin {
	if o == nil {
		return nil
	}
	return o.UpstreamTimeoutPlugin
}

func (o *UpdateUpstreamtimeoutPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
