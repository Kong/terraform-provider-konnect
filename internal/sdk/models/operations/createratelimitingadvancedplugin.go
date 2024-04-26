// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateRatelimitingadvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                   string                                   `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateRateLimitingAdvancedPlugin *shared.CreateRateLimitingAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *CreateRatelimitingadvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateRatelimitingadvancedPluginRequest) GetCreateRateLimitingAdvancedPlugin() *shared.CreateRateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.CreateRateLimitingAdvancedPlugin
}

type CreateRatelimitingadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// RateLimitingAdvanced plugin
	RateLimitingAdvancedPlugin *shared.RateLimitingAdvancedPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreateRatelimitingadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRatelimitingadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRatelimitingadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRatelimitingadvancedPluginResponse) GetRateLimitingAdvancedPlugin() *shared.RateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.RateLimitingAdvancedPlugin
}

func (o *CreateRatelimitingadvancedPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
