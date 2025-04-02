// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateHmacauthPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string                `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	HmacAuthPlugin shared.HmacAuthPlugin `request:"mediaType=application/json"`
}

func (o *UpdateHmacauthPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateHmacauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateHmacauthPluginRequest) GetHmacAuthPlugin() shared.HmacAuthPlugin {
	if o == nil {
		return shared.HmacAuthPlugin{}
	}
	return o.HmacAuthPlugin
}

type UpdateHmacauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// HmacAuth plugin
	HmacAuthPlugin *shared.HmacAuthPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateHmacauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHmacauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHmacauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateHmacauthPluginResponse) GetHmacAuthPlugin() *shared.HmacAuthPlugin {
	if o == nil {
		return nil
	}
	return o.HmacAuthPlugin
}

func (o *UpdateHmacauthPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
