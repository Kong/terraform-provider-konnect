// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateResponseratelimitingPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID             string                                 `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	ResponseRatelimitingPlugin shared.ResponseRatelimitingPluginInput `request:"mediaType=application/json"`
}

func (o *CreateResponseratelimitingPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateResponseratelimitingPluginRequest) GetResponseRatelimitingPlugin() shared.ResponseRatelimitingPluginInput {
	if o == nil {
		return shared.ResponseRatelimitingPluginInput{}
	}
	return o.ResponseRatelimitingPlugin
}

type CreateResponseratelimitingPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created ResponseRatelimiting plugin
	ResponseRatelimitingPlugin *shared.ResponseRatelimitingPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateResponseratelimitingPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateResponseratelimitingPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateResponseratelimitingPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateResponseratelimitingPluginResponse) GetResponseRatelimitingPlugin() *shared.ResponseRatelimitingPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseRatelimitingPlugin
}

func (o *CreateResponseratelimitingPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
