// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreatePrefunctionPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Create a new PreFunction plugin
	CreatePreFunctionPlugin shared.CreatePreFunctionPlugin `request:"mediaType=application/json"`
}

func (o *CreatePrefunctionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreatePrefunctionPluginRequest) GetCreatePreFunctionPlugin() shared.CreatePreFunctionPlugin {
	if o == nil {
		return shared.CreatePreFunctionPlugin{}
	}
	return o.CreatePreFunctionPlugin
}

type CreatePrefunctionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Plugin
	PreFunctionPlugin *shared.PreFunctionPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreatePrefunctionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePrefunctionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePrefunctionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePrefunctionPluginResponse) GetPreFunctionPlugin() *shared.PreFunctionPlugin {
	if o == nil {
		return nil
	}
	return o.PreFunctionPlugin
}

func (o *CreatePrefunctionPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
