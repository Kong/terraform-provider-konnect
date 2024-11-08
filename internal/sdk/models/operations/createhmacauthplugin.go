// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateHmacauthPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string                      `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	HmacAuthPlugin *shared.HmacAuthPluginInput `request:"mediaType=application/json"`
}

func (o *CreateHmacauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateHmacauthPluginRequest) GetHmacAuthPlugin() *shared.HmacAuthPluginInput {
	if o == nil {
		return nil
	}
	return o.HmacAuthPlugin
}

type CreateHmacauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created HmacAuth plugin
	HmacAuthPlugin *shared.HmacAuthPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateHmacauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateHmacauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateHmacauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateHmacauthPluginResponse) GetHmacAuthPlugin() *shared.HmacAuthPlugin {
	if o == nil {
		return nil
	}
	return o.HmacAuthPlugin
}

func (o *CreateHmacauthPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
