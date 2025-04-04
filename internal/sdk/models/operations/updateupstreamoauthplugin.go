// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateUpstreamoauthPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID      string                     `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	UpstreamOauthPlugin shared.UpstreamOauthPlugin `request:"mediaType=application/json"`
}

func (o *UpdateUpstreamoauthPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateUpstreamoauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateUpstreamoauthPluginRequest) GetUpstreamOauthPlugin() shared.UpstreamOauthPlugin {
	if o == nil {
		return shared.UpstreamOauthPlugin{}
	}
	return o.UpstreamOauthPlugin
}

type UpdateUpstreamoauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// UpstreamOauth plugin
	UpstreamOauthPlugin *shared.UpstreamOauthPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateUpstreamoauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUpstreamoauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUpstreamoauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUpstreamoauthPluginResponse) GetUpstreamOauthPlugin() *shared.UpstreamOauthPlugin {
	if o == nil {
		return nil
	}
	return o.UpstreamOauthPlugin
}

func (o *UpdateUpstreamoauthPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
