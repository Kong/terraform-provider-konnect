// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetPluginSchemaRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// The custom plugin name
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetPluginSchemaRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetPluginSchemaRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetPluginSchemaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response for a single custom plugin schema.
	PluginSchemas *shared.PluginSchemas
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetPluginSchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPluginSchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPluginSchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPluginSchemaResponse) GetPluginSchemas() *shared.PluginSchemas {
	if o == nil {
		return nil
	}
	return o.PluginSchemas
}

func (o *GetPluginSchemaResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetPluginSchemaResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetPluginSchemaResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
