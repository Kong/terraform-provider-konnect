// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateMockingPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string               `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	MockingPlugin  shared.MockingPlugin `request:"mediaType=application/json"`
}

func (o *CreateMockingPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateMockingPluginRequest) GetMockingPlugin() shared.MockingPlugin {
	if o == nil {
		return shared.MockingPlugin{}
	}
	return o.MockingPlugin
}

type CreateMockingPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created Mocking plugin
	MockingPlugin *shared.MockingPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateMockingPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMockingPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMockingPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMockingPluginResponse) GetMockingPlugin() *shared.MockingPlugin {
	if o == nil {
		return nil
	}
	return o.MockingPlugin
}

func (o *CreateMockingPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
