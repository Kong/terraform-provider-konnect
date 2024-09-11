// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateOasvalidationPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID            string                            `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateOasValidationPlugin *shared.CreateOasValidationPlugin `request:"mediaType=application/json"`
}

func (o *UpdateOasvalidationPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateOasvalidationPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateOasvalidationPluginRequest) GetCreateOasValidationPlugin() *shared.CreateOasValidationPlugin {
	if o == nil {
		return nil
	}
	return o.CreateOasValidationPlugin
}

type UpdateOasvalidationPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OasValidation plugin
	OasValidationPlugin *shared.OasValidationPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateOasvalidationPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateOasvalidationPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateOasvalidationPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateOasvalidationPluginResponse) GetOasValidationPlugin() *shared.OasValidationPlugin {
	if o == nil {
		return nil
	}
	return o.OasValidationPlugin
}

func (o *UpdateOasvalidationPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
