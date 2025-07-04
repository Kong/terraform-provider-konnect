// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdatePluginSchemasRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// The custom plugin name
	Name                string                      `pathParam:"style=simple,explode=false,name=name"`
	CreatePluginSchemas *shared.CreatePluginSchemas `request:"mediaType=application/json"`
}

func (o *UpdatePluginSchemasRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdatePluginSchemasRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdatePluginSchemasRequest) GetCreatePluginSchemas() *shared.CreatePluginSchemas {
	if o == nil {
		return nil
	}
	return o.CreatePluginSchemas
}

type UpdatePluginSchemasResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response for a single custom plugin schema.
	PluginSchemas *shared.PluginSchemas
	// Forbidden
	KonnectCPLegacyBadRequestError *shared.KonnectCPLegacyBadRequestError
	// Unauthorized
	KonnectCPLegacyUnauthorizedError *shared.KonnectCPLegacyUnauthorizedError
	// Forbidden
	KonnectCPLegacyForbiddenError *shared.KonnectCPLegacyForbiddenError
	// Forbidden
	KonnectCPLegacyNotFoundError *shared.KonnectCPLegacyNotFoundError
}

func (o *UpdatePluginSchemasResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePluginSchemasResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePluginSchemasResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePluginSchemasResponse) GetPluginSchemas() *shared.PluginSchemas {
	if o == nil {
		return nil
	}
	return o.PluginSchemas
}

func (o *UpdatePluginSchemasResponse) GetKonnectCPLegacyBadRequestError() *shared.KonnectCPLegacyBadRequestError {
	if o == nil {
		return nil
	}
	return o.KonnectCPLegacyBadRequestError
}

func (o *UpdatePluginSchemasResponse) GetKonnectCPLegacyUnauthorizedError() *shared.KonnectCPLegacyUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.KonnectCPLegacyUnauthorizedError
}

func (o *UpdatePluginSchemasResponse) GetKonnectCPLegacyForbiddenError() *shared.KonnectCPLegacyForbiddenError {
	if o == nil {
		return nil
	}
	return o.KonnectCPLegacyForbiddenError
}

func (o *UpdatePluginSchemasResponse) GetKonnectCPLegacyNotFoundError() *shared.KonnectCPLegacyNotFoundError {
	if o == nil {
		return nil
	}
	return o.KonnectCPLegacyNotFoundError
}
