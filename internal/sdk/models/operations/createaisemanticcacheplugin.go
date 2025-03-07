// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateAisemanticcachePluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID        string                            `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	AiSemanticCachePlugin shared.AiSemanticCachePluginInput `request:"mediaType=application/json"`
}

func (o *CreateAisemanticcachePluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateAisemanticcachePluginRequest) GetAiSemanticCachePlugin() shared.AiSemanticCachePluginInput {
	if o == nil {
		return shared.AiSemanticCachePluginInput{}
	}
	return o.AiSemanticCachePlugin
}

type CreateAisemanticcachePluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created AiSemanticCache plugin
	AiSemanticCachePlugin *shared.AiSemanticCachePlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateAisemanticcachePluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAisemanticcachePluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAisemanticcachePluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateAisemanticcachePluginResponse) GetAiSemanticCachePlugin() *shared.AiSemanticCachePlugin {
	if o == nil {
		return nil
	}
	return o.AiSemanticCachePlugin
}

func (o *CreateAisemanticcachePluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
