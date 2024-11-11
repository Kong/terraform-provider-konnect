// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateSyslogPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string                    `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	SyslogPlugin   *shared.SyslogPluginInput `request:"mediaType=application/json"`
}

func (o *CreateSyslogPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateSyslogPluginRequest) GetSyslogPlugin() *shared.SyslogPluginInput {
	if o == nil {
		return nil
	}
	return o.SyslogPlugin
}

type CreateSyslogPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created Syslog plugin
	SyslogPlugin *shared.SyslogPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateSyslogPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSyslogPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSyslogPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSyslogPluginResponse) GetSyslogPlugin() *shared.SyslogPlugin {
	if o == nil {
		return nil
	}
	return o.SyslogPlugin
}

func (o *CreateSyslogPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
