// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateHttplogPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID      string                      `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateHTTPLogPlugin *shared.CreateHTTPLogPlugin `request:"mediaType=application/json"`
}

func (o *UpdateHttplogPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateHttplogPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateHttplogPluginRequest) GetCreateHTTPLogPlugin() *shared.CreateHTTPLogPlugin {
	if o == nil {
		return nil
	}
	return o.CreateHTTPLogPlugin
}

type UpdateHttplogPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// HttpLog plugin
	HTTPLogPlugin *shared.HTTPLogPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateHttplogPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHttplogPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHttplogPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateHttplogPluginResponse) GetHTTPLogPlugin() *shared.HTTPLogPlugin {
	if o == nil {
		return nil
	}
	return o.HTTPLogPlugin
}

func (o *UpdateHttplogPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
