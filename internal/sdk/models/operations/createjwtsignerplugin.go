// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateJwtsignerPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID        string                        `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateJWTSignerPlugin *shared.CreateJWTSignerPlugin `request:"mediaType=application/json"`
}

func (o *CreateJwtsignerPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateJwtsignerPluginRequest) GetCreateJWTSignerPlugin() *shared.CreateJWTSignerPlugin {
	if o == nil {
		return nil
	}
	return o.CreateJWTSignerPlugin
}

type CreateJwtsignerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created JWTSigner plugin
	JWTSignerPlugin *shared.JWTSignerPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateJwtsignerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateJwtsignerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateJwtsignerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateJwtsignerPluginResponse) GetJWTSignerPlugin() *shared.JWTSignerPlugin {
	if o == nil {
		return nil
	}
	return o.JWTSignerPlugin
}

func (o *CreateJwtsignerPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}