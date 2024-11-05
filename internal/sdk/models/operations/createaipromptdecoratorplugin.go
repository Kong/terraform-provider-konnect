// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateAipromptdecoratorPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                string                                `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateAiPromptDecoratorPlugin *shared.CreateAiPromptDecoratorPlugin `request:"mediaType=application/json"`
}

func (o *CreateAipromptdecoratorPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateAipromptdecoratorPluginRequest) GetCreateAiPromptDecoratorPlugin() *shared.CreateAiPromptDecoratorPlugin {
	if o == nil {
		return nil
	}
	return o.CreateAiPromptDecoratorPlugin
}

type CreateAipromptdecoratorPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created AiPromptDecorator plugin
	AiPromptDecoratorPlugin *shared.AiPromptDecoratorPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateAipromptdecoratorPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAipromptdecoratorPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAipromptdecoratorPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateAipromptdecoratorPluginResponse) GetAiPromptDecoratorPlugin() *shared.AiPromptDecoratorPlugin {
	if o == nil {
		return nil
	}
	return o.AiPromptDecoratorPlugin
}

func (o *CreateAipromptdecoratorPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}