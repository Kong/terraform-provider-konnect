// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type PostControlPlanesIDGroupMembershipsRemoveRequest struct {
	// ID of a control plane group
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Request body for removing a list of child control planes from a control plane group membership.
	GroupMembership *shared.GroupMembership `request:"mediaType=application/json"`
}

func (o *PostControlPlanesIDGroupMembershipsRemoveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PostControlPlanesIDGroupMembershipsRemoveRequest) GetGroupMembership() *shared.GroupMembership {
	if o == nil {
		return nil
	}
	return o.GroupMembership
}

type PostControlPlanesIDGroupMembershipsRemoveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Service Unavailable
	ServiceUnavailable *shared.ServiceUnavailable
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *PostControlPlanesIDGroupMembershipsRemoveResponse) GetServiceUnavailable() *shared.ServiceUnavailable {
	if o == nil {
		return nil
	}
	return o.ServiceUnavailable
}