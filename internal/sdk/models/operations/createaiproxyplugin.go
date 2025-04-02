// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateAiproxyPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string               `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	AiProxyPlugin  shared.AiProxyPlugin `request:"mediaType=application/json"`
}

func (o *CreateAiproxyPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateAiproxyPluginRequest) GetAiProxyPlugin() shared.AiProxyPlugin {
	if o == nil {
		return shared.AiProxyPlugin{}
	}
	return o.AiProxyPlugin
}

type CreateAiproxyPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created AiProxy plugin
	AiProxyPlugin *shared.AiProxyPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateAiproxyPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAiproxyPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAiproxyPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateAiproxyPluginResponse) GetAiProxyPlugin() *shared.AiProxyPlugin {
	if o == nil {
		return nil
	}
	return o.AiProxyPlugin
}

func (o *CreateAiproxyPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
