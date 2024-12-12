// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateIprestrictionPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID      string                          `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	IPRestrictionPlugin shared.IPRestrictionPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateIprestrictionPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateIprestrictionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateIprestrictionPluginRequest) GetIPRestrictionPlugin() shared.IPRestrictionPluginInput {
	if o == nil {
		return shared.IPRestrictionPluginInput{}
	}
	return o.IPRestrictionPlugin
}

type UpdateIprestrictionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// IpRestriction plugin
	IPRestrictionPlugin *shared.IPRestrictionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateIprestrictionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIprestrictionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIprestrictionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateIprestrictionPluginResponse) GetIPRestrictionPlugin() *shared.IPRestrictionPlugin {
	if o == nil {
		return nil
	}
	return o.IPRestrictionPlugin
}

func (o *UpdateIprestrictionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
