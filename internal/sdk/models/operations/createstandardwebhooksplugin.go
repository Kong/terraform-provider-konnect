// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateStandardwebhooksPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID         string                        `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	StandardWebhooksPlugin shared.StandardWebhooksPlugin `request:"mediaType=application/json"`
}

func (o *CreateStandardwebhooksPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateStandardwebhooksPluginRequest) GetStandardWebhooksPlugin() shared.StandardWebhooksPlugin {
	if o == nil {
		return shared.StandardWebhooksPlugin{}
	}
	return o.StandardWebhooksPlugin
}

type CreateStandardwebhooksPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created StandardWebhooks plugin
	StandardWebhooksPlugin *shared.StandardWebhooksPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateStandardwebhooksPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateStandardwebhooksPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateStandardwebhooksPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateStandardwebhooksPluginResponse) GetStandardWebhooksPlugin() *shared.StandardWebhooksPlugin {
	if o == nil {
		return nil
	}
	return o.StandardWebhooksPlugin
}

func (o *CreateStandardwebhooksPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
