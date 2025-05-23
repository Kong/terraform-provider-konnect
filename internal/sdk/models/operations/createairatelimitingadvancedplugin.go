// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateAiratelimitingadvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID               string                              `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	AiRateLimitingAdvancedPlugin shared.AiRateLimitingAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *CreateAiratelimitingadvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateAiratelimitingadvancedPluginRequest) GetAiRateLimitingAdvancedPlugin() shared.AiRateLimitingAdvancedPlugin {
	if o == nil {
		return shared.AiRateLimitingAdvancedPlugin{}
	}
	return o.AiRateLimitingAdvancedPlugin
}

type CreateAiratelimitingadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created AiRateLimitingAdvanced plugin
	AiRateLimitingAdvancedPlugin *shared.AiRateLimitingAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateAiratelimitingadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAiratelimitingadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAiratelimitingadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateAiratelimitingadvancedPluginResponse) GetAiRateLimitingAdvancedPlugin() *shared.AiRateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.AiRateLimitingAdvancedPlugin
}

func (o *CreateAiratelimitingadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
