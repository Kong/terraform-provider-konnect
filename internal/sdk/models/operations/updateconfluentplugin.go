// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateConfluentPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID  string                 `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	ConfluentPlugin shared.ConfluentPlugin `request:"mediaType=application/json"`
}

func (o *UpdateConfluentPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateConfluentPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateConfluentPluginRequest) GetConfluentPlugin() shared.ConfluentPlugin {
	if o == nil {
		return shared.ConfluentPlugin{}
	}
	return o.ConfluentPlugin
}

type UpdateConfluentPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Confluent plugin
	ConfluentPlugin *shared.ConfluentPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateConfluentPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConfluentPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConfluentPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateConfluentPluginResponse) GetConfluentPlugin() *shared.ConfluentPlugin {
	if o == nil {
		return nil
	}
	return o.ConfluentPlugin
}

func (o *UpdateConfluentPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
