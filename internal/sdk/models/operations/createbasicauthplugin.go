// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateBasicauthPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID  string                       `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	BasicAuthPlugin *shared.BasicAuthPluginInput `request:"mediaType=application/json"`
}

func (o *CreateBasicauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateBasicauthPluginRequest) GetBasicAuthPlugin() *shared.BasicAuthPluginInput {
	if o == nil {
		return nil
	}
	return o.BasicAuthPlugin
}

type CreateBasicauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created BasicAuth plugin
	BasicAuthPlugin *shared.BasicAuthPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateBasicauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBasicauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBasicauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateBasicauthPluginResponse) GetBasicAuthPlugin() *shared.BasicAuthPlugin {
	if o == nil {
		return nil
	}
	return o.BasicAuthPlugin
}

func (o *CreateBasicauthPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
