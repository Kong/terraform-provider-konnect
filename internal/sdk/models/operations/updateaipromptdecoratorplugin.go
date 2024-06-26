// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateAipromptdecoratorPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                string                                `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateAIPromptDecoratorPlugin *shared.CreateAIPromptDecoratorPlugin `request:"mediaType=application/json"`
}

func (o *UpdateAipromptdecoratorPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateAipromptdecoratorPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateAipromptdecoratorPluginRequest) GetCreateAIPromptDecoratorPlugin() *shared.CreateAIPromptDecoratorPlugin {
	if o == nil {
		return nil
	}
	return o.CreateAIPromptDecoratorPlugin
}

type UpdateAipromptdecoratorPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// AIPromptDecorator plugin
	AIPromptDecoratorPlugin *shared.AIPromptDecoratorPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateAipromptdecoratorPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAipromptdecoratorPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAipromptdecoratorPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAipromptdecoratorPluginResponse) GetAIPromptDecoratorPlugin() *shared.AIPromptDecoratorPlugin {
	if o == nil {
		return nil
	}
	return o.AIPromptDecoratorPlugin
}

func (o *UpdateAipromptdecoratorPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
