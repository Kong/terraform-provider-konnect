// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type DeletePortalProductVersionRequest struct {
	// API product version identifier
	ProductVersionID string `pathParam:"style=simple,explode=false,name=productVersionId"`
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

func (o *DeletePortalProductVersionRequest) GetProductVersionID() string {
	if o == nil {
		return ""
	}
	return o.ProductVersionID
}

func (o *DeletePortalProductVersionRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

type DeletePortalProductVersionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeletePortalProductVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePortalProductVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePortalProductVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeletePortalProductVersionResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *DeletePortalProductVersionResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *DeletePortalProductVersionResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
