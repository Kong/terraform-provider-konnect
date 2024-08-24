// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateAiproxyPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID      string                      `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateAIProxyPlugin *shared.CreateAIProxyPlugin `request:"mediaType=application/json"`
}

func (o *CreateAiproxyPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateAiproxyPluginRequest) GetCreateAIProxyPlugin() *shared.CreateAIProxyPlugin {
	if o == nil {
		return nil
	}
	return o.CreateAIProxyPlugin
}

type CreateAiproxyPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created AIProxy plugin
	AIProxyPlugin *shared.AIProxyPlugin
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

func (o *CreateAiproxyPluginResponse) GetAIProxyPlugin() *shared.AIProxyPlugin {
	if o == nil {
		return nil
	}
	return o.AIProxyPlugin
}

func (o *CreateAiproxyPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
