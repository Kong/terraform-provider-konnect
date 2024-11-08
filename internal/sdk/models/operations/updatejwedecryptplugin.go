// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateJwedecryptPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID         string                         `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateJweDecryptPlugin *shared.CreateJweDecryptPlugin `request:"mediaType=application/json"`
}

func (o *UpdateJwedecryptPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateJwedecryptPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateJwedecryptPluginRequest) GetCreateJweDecryptPlugin() *shared.CreateJweDecryptPlugin {
	if o == nil {
		return nil
	}
	return o.CreateJweDecryptPlugin
}

type UpdateJwedecryptPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// JweDecrypt plugin
	JweDecryptPlugin *shared.JweDecryptPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateJwedecryptPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateJwedecryptPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateJwedecryptPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateJwedecryptPluginResponse) GetJweDecryptPlugin() *shared.JweDecryptPlugin {
	if o == nil {
		return nil
	}
	return o.JweDecryptPlugin
}

func (o *UpdateJwedecryptPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}