// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateControlPlaneRequest struct {
	// The control plane ID
	ID                        string                           `pathParam:"style=simple,explode=false,name=id"`
	UpdateControlPlaneRequest shared.UpdateControlPlaneRequest `request:"mediaType=application/json"`
}

func (o *UpdateControlPlaneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateControlPlaneRequest) GetUpdateControlPlaneRequest() shared.UpdateControlPlaneRequest {
	if o == nil {
		return shared.UpdateControlPlaneRequest{}
	}
	return o.UpdateControlPlaneRequest
}

type UpdateControlPlaneResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response to updating a control plane.
	ControlPlane *shared.ControlPlane
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Internal Server Error
	InternalServerError *shared.InternalServerError
	// Service Unavailable
	ServiceUnavailable *shared.ServiceUnavailable
}

func (o *UpdateControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateControlPlaneResponse) GetControlPlane() *shared.ControlPlane {
	if o == nil {
		return nil
	}
	return o.ControlPlane
}

func (o *UpdateControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *UpdateControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdateControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdateControlPlaneResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *UpdateControlPlaneResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateControlPlaneResponse) GetServiceUnavailable() *shared.ServiceUnavailable {
	if o == nil {
		return nil
	}
	return o.ServiceUnavailable
}
