// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateGrpcgatewayPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID          string                          `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateGrpcGatewayPlugin *shared.CreateGrpcGatewayPlugin `request:"mediaType=application/json"`
}

func (o *UpdateGrpcgatewayPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateGrpcgatewayPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateGrpcgatewayPluginRequest) GetCreateGrpcGatewayPlugin() *shared.CreateGrpcGatewayPlugin {
	if o == nil {
		return nil
	}
	return o.CreateGrpcGatewayPlugin
}

type UpdateGrpcgatewayPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// GrpcGateway plugin
	GrpcGatewayPlugin *shared.GrpcGatewayPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateGrpcgatewayPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateGrpcgatewayPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateGrpcgatewayPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateGrpcgatewayPluginResponse) GetGrpcGatewayPlugin() *shared.GrpcGatewayPlugin {
	if o == nil {
		return nil
	}
	return o.GrpcGatewayPlugin
}

func (o *UpdateGrpcgatewayPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}