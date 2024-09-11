// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateResponseratelimitingPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                   string                                   `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateResponseRatelimitingPlugin *shared.CreateResponseRatelimitingPlugin `request:"mediaType=application/json"`
}

func (o *UpdateResponseratelimitingPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateResponseratelimitingPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateResponseratelimitingPluginRequest) GetCreateResponseRatelimitingPlugin() *shared.CreateResponseRatelimitingPlugin {
	if o == nil {
		return nil
	}
	return o.CreateResponseRatelimitingPlugin
}

type UpdateResponseratelimitingPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ResponseRatelimiting plugin
	ResponseRatelimitingPlugin *shared.ResponseRatelimitingPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateResponseratelimitingPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateResponseratelimitingPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateResponseratelimitingPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateResponseratelimitingPluginResponse) GetResponseRatelimitingPlugin() *shared.ResponseRatelimitingPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseRatelimitingPlugin
}

func (o *UpdateResponseratelimitingPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
