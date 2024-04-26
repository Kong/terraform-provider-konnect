// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateJqPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string                 `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateJQPlugin *shared.CreateJQPlugin `request:"mediaType=application/json"`
}

func (o *UpdateJqPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateJqPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateJqPluginRequest) GetCreateJQPlugin() *shared.CreateJQPlugin {
	if o == nil {
		return nil
	}
	return o.CreateJQPlugin
}

type UpdateJqPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// JQ plugin
	JQPlugin *shared.JQPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateJqPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateJqPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateJqPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateJqPluginResponse) GetJQPlugin() *shared.JQPlugin {
	if o == nil {
		return nil
	}
	return o.JQPlugin
}

func (o *UpdateJqPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
