// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateServiceprotectionPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID          string                              `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	ServiceProtectionPlugin shared.ServiceProtectionPluginInput `request:"mediaType=application/json"`
}

func (o *CreateServiceprotectionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateServiceprotectionPluginRequest) GetServiceProtectionPlugin() shared.ServiceProtectionPluginInput {
	if o == nil {
		return shared.ServiceProtectionPluginInput{}
	}
	return o.ServiceProtectionPlugin
}

type CreateServiceprotectionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created ServiceProtection plugin
	ServiceProtectionPlugin *shared.ServiceProtectionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateServiceprotectionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateServiceprotectionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateServiceprotectionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateServiceprotectionPluginResponse) GetServiceProtectionPlugin() *shared.ServiceProtectionPlugin {
	if o == nil {
		return nil
	}
	return o.ServiceProtectionPlugin
}

func (o *CreateServiceprotectionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
