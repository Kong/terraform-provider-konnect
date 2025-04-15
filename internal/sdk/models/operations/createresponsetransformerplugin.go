// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateResponsetransformerPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID            string                           `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	ResponseTransformerPlugin shared.ResponseTransformerPlugin `request:"mediaType=application/json"`
}

func (o *CreateResponsetransformerPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateResponsetransformerPluginRequest) GetResponseTransformerPlugin() shared.ResponseTransformerPlugin {
	if o == nil {
		return shared.ResponseTransformerPlugin{}
	}
	return o.ResponseTransformerPlugin
}

type CreateResponsetransformerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created ResponseTransformer plugin
	ResponseTransformerPlugin *shared.ResponseTransformerPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateResponsetransformerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateResponsetransformerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateResponsetransformerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateResponsetransformerPluginResponse) GetResponseTransformerPlugin() *shared.ResponseTransformerPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseTransformerPlugin
}

func (o *CreateResponsetransformerPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
