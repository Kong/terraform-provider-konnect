// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpsertCustomPluginRequest struct {
	// ID of the CustomPlugin to lookup
	CustomPluginID string `pathParam:"style=simple,explode=false,name=CustomPluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the CustomPlugin
	CustomPlugin shared.CustomPlugin `request:"mediaType=application/json"`
}

func (o *UpsertCustomPluginRequest) GetCustomPluginID() string {
	if o == nil {
		return ""
	}
	return o.CustomPluginID
}

func (o *UpsertCustomPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertCustomPluginRequest) GetCustomPlugin() shared.CustomPlugin {
	if o == nil {
		return shared.CustomPlugin{}
	}
	return o.CustomPlugin
}

type UpsertCustomPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted CustomPlugin
	CustomPlugin *shared.CustomPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertCustomPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertCustomPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertCustomPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertCustomPluginResponse) GetCustomPlugin() *shared.CustomPlugin {
	if o == nil {
		return nil
	}
	return o.CustomPlugin
}

func (o *UpsertCustomPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
