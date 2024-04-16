// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateOauth2PluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// Description of the Plugin
	CreateOauth2Plugin shared.CreateOauth2Plugin `request:"mediaType=application/json"`
}

func (o *UpdateOauth2PluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateOauth2PluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateOauth2PluginRequest) GetCreateOauth2Plugin() shared.CreateOauth2Plugin {
	if o == nil {
		return shared.CreateOauth2Plugin{}
	}
	return o.CreateOauth2Plugin
}

type UpdateOauth2PluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Plugin
	Oauth2Plugin *shared.Oauth2Plugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateOauth2PluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateOauth2PluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateOauth2PluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateOauth2PluginResponse) GetOauth2Plugin() *shared.Oauth2Plugin {
	if o == nil {
		return nil
	}
	return o.Oauth2Plugin
}

func (o *UpdateOauth2PluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
