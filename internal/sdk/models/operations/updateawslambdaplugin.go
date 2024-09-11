// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateAwslambdaPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID        string                        `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateAwsLambdaPlugin *shared.CreateAwsLambdaPlugin `request:"mediaType=application/json"`
}

func (o *UpdateAwslambdaPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateAwslambdaPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateAwslambdaPluginRequest) GetCreateAwsLambdaPlugin() *shared.CreateAwsLambdaPlugin {
	if o == nil {
		return nil
	}
	return o.CreateAwsLambdaPlugin
}

type UpdateAwslambdaPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// AwsLambda plugin
	AwsLambdaPlugin *shared.AwsLambdaPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateAwslambdaPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAwslambdaPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAwslambdaPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAwslambdaPluginResponse) GetAwsLambdaPlugin() *shared.AwsLambdaPlugin {
	if o == nil {
		return nil
	}
	return o.AwsLambdaPlugin
}

func (o *UpdateAwslambdaPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
