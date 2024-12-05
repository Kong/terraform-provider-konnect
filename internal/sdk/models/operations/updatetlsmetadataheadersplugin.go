// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateTlsmetadataheadersPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID           string                                `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	TLSMetadataHeadersPlugin *shared.TLSMetadataHeadersPluginInput `request:"mediaType=application/json"`
}

func (o *UpdateTlsmetadataheadersPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateTlsmetadataheadersPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateTlsmetadataheadersPluginRequest) GetTLSMetadataHeadersPlugin() *shared.TLSMetadataHeadersPluginInput {
	if o == nil {
		return nil
	}
	return o.TLSMetadataHeadersPlugin
}

type UpdateTlsmetadataheadersPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// TlsMetadataHeaders plugin
	TLSMetadataHeadersPlugin *shared.TLSMetadataHeadersPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateTlsmetadataheadersPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTlsmetadataheadersPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTlsmetadataheadersPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTlsmetadataheadersPluginResponse) GetTLSMetadataHeadersPlugin() *shared.TLSMetadataHeadersPlugin {
	if o == nil {
		return nil
	}
	return o.TLSMetadataHeadersPlugin
}

func (o *UpdateTlsmetadataheadersPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}