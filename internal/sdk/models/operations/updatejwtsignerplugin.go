// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateJwtsignerPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID        string                        `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateJWTSignerPlugin *shared.CreateJWTSignerPlugin `request:"mediaType=application/json"`
}

func (o *UpdateJwtsignerPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateJwtsignerPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateJwtsignerPluginRequest) GetCreateJWTSignerPlugin() *shared.CreateJWTSignerPlugin {
	if o == nil {
		return nil
	}
	return o.CreateJWTSignerPlugin
}

type UpdateJwtsignerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// JWTSigner plugin
	JWTSignerPlugin *shared.JWTSignerPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateJwtsignerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateJwtsignerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateJwtsignerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateJwtsignerPluginResponse) GetJWTSignerPlugin() *shared.JWTSignerPlugin {
	if o == nil {
		return nil
	}
	return o.JWTSignerPlugin
}

func (o *UpdateJwtsignerPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
