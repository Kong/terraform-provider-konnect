// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateBotdetectionPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID           string                           `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateBotDetectionPlugin *shared.CreateBotDetectionPlugin `request:"mediaType=application/json"`
}

func (o *UpdateBotdetectionPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateBotdetectionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateBotdetectionPluginRequest) GetCreateBotDetectionPlugin() *shared.CreateBotDetectionPlugin {
	if o == nil {
		return nil
	}
	return o.CreateBotDetectionPlugin
}

type UpdateBotdetectionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// BotDetection plugin
	BotDetectionPlugin *shared.BotDetectionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateBotdetectionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateBotdetectionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBotdetectionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateBotdetectionPluginResponse) GetBotDetectionPlugin() *shared.BotDetectionPlugin {
	if o == nil {
		return nil
	}
	return o.BotDetectionPlugin
}

func (o *UpdateBotdetectionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
