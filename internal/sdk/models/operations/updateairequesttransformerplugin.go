// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateAirequesttransformerPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                   string                                   `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateAiRequestTransformerPlugin *shared.CreateAiRequestTransformerPlugin `request:"mediaType=application/json"`
}

func (o *UpdateAirequesttransformerPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateAirequesttransformerPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateAirequesttransformerPluginRequest) GetCreateAiRequestTransformerPlugin() *shared.CreateAiRequestTransformerPlugin {
	if o == nil {
		return nil
	}
	return o.CreateAiRequestTransformerPlugin
}

type UpdateAirequesttransformerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// AiRequestTransformer plugin
	AiRequestTransformerPlugin *shared.AiRequestTransformerPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateAirequesttransformerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAirequesttransformerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAirequesttransformerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAirequesttransformerPluginResponse) GetAiRequestTransformerPlugin() *shared.AiRequestTransformerPlugin {
	if o == nil {
		return nil
	}
	return o.AiRequestTransformerPlugin
}

func (o *UpdateAirequesttransformerPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
