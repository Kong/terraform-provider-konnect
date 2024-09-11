// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateKeyauthencPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID         string                         `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateKeyAuthEncPlugin *shared.CreateKeyAuthEncPlugin `request:"mediaType=application/json"`
}

func (o *UpdateKeyauthencPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateKeyauthencPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateKeyauthencPluginRequest) GetCreateKeyAuthEncPlugin() *shared.CreateKeyAuthEncPlugin {
	if o == nil {
		return nil
	}
	return o.CreateKeyAuthEncPlugin
}

type UpdateKeyauthencPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// KeyAuthEnc plugin
	KeyAuthEncPlugin *shared.KeyAuthEncPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateKeyauthencPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateKeyauthencPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateKeyauthencPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateKeyauthencPluginResponse) GetKeyAuthEncPlugin() *shared.KeyAuthEncPlugin {
	if o == nil {
		return nil
	}
	return o.KeyAuthEncPlugin
}

func (o *UpdateKeyauthencPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
