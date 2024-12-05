// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateUdplogPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string                    `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	UDPLogPlugin   *shared.UDPLogPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateUdplogPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateUdplogPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateUdplogPluginRequest) GetUDPLogPlugin() *shared.UDPLogPluginInput {
	if o == nil {
		return nil
	}
	return o.UDPLogPlugin
}

type UpdateUdplogPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// UdpLog plugin
	UDPLogPlugin *shared.UDPLogPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateUdplogPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUdplogPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUdplogPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUdplogPluginResponse) GetUDPLogPlugin() *shared.UDPLogPlugin {
	if o == nil {
		return nil
	}
	return o.UDPLogPlugin
}

func (o *UpdateUdplogPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}