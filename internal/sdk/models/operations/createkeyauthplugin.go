// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateKeyauthPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string               `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	KeyAuthPlugin  shared.KeyAuthPlugin `request:"mediaType=application/json"`
}

func (o *CreateKeyauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateKeyauthPluginRequest) GetKeyAuthPlugin() shared.KeyAuthPlugin {
	if o == nil {
		return shared.KeyAuthPlugin{}
	}
	return o.KeyAuthPlugin
}

type CreateKeyauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created KeyAuth plugin
	KeyAuthPlugin *shared.KeyAuthPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateKeyauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateKeyauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateKeyauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateKeyauthPluginResponse) GetKeyAuthPlugin() *shared.KeyAuthPlugin {
	if o == nil {
		return nil
	}
	return o.KeyAuthPlugin
}

func (o *CreateKeyauthPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
