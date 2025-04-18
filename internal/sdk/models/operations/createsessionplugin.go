// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateSessionPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string               `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	SessionPlugin  shared.SessionPlugin `request:"mediaType=application/json"`
}

func (o *CreateSessionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateSessionPluginRequest) GetSessionPlugin() shared.SessionPlugin {
	if o == nil {
		return shared.SessionPlugin{}
	}
	return o.SessionPlugin
}

type CreateSessionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created Session plugin
	SessionPlugin *shared.SessionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateSessionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSessionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSessionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSessionPluginResponse) GetSessionPlugin() *shared.SessionPlugin {
	if o == nil {
		return nil
	}
	return o.SessionPlugin
}

func (o *CreateSessionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
