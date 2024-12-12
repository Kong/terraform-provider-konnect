// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateFilelogPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string                    `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	FileLogPlugin  shared.FileLogPluginInput `request:"mediaType=application/json"`
}

func (o *CreateFilelogPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateFilelogPluginRequest) GetFileLogPlugin() shared.FileLogPluginInput {
	if o == nil {
		return shared.FileLogPluginInput{}
	}
	return o.FileLogPlugin
}

type CreateFilelogPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created FileLog plugin
	FileLogPlugin *shared.FileLogPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateFilelogPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFilelogPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFilelogPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateFilelogPluginResponse) GetFileLogPlugin() *shared.FileLogPlugin {
	if o == nil {
		return nil
	}
	return o.FileLogPlugin
}

func (o *CreateFilelogPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
