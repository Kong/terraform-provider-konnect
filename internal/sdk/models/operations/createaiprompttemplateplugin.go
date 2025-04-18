// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateAiprompttemplatePluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID         string                        `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	AiPromptTemplatePlugin shared.AiPromptTemplatePlugin `request:"mediaType=application/json"`
}

func (o *CreateAiprompttemplatePluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateAiprompttemplatePluginRequest) GetAiPromptTemplatePlugin() shared.AiPromptTemplatePlugin {
	if o == nil {
		return shared.AiPromptTemplatePlugin{}
	}
	return o.AiPromptTemplatePlugin
}

type CreateAiprompttemplatePluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created AiPromptTemplate plugin
	AiPromptTemplatePlugin *shared.AiPromptTemplatePlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateAiprompttemplatePluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAiprompttemplatePluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAiprompttemplatePluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateAiprompttemplatePluginResponse) GetAiPromptTemplatePlugin() *shared.AiPromptTemplatePlugin {
	if o == nil {
		return nil
	}
	return o.AiPromptTemplatePlugin
}

func (o *CreateAiprompttemplatePluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
