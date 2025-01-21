// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdateMeshControlPlaneRequest struct {
	// Id of the Konnect resource
	ID                            string                               `pathParam:"style=simple,explode=false,name=id"`
	UpdateMeshControlPlaneRequest shared.UpdateMeshControlPlaneRequest `request:"mediaType=application/json"`
}

func (o *UpdateMeshControlPlaneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateMeshControlPlaneRequest) GetUpdateMeshControlPlaneRequest() shared.UpdateMeshControlPlaneRequest {
	if o == nil {
		return shared.UpdateMeshControlPlaneRequest{}
	}
	return o.UpdateMeshControlPlaneRequest
}

type UpdateMeshControlPlaneResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response to updating a control plane.
	MeshControlPlane *shared.MeshControlPlane
	// Validation Error
	BadRequestError *shared.BadRequestError
	// Unauthorized Error
	UnauthorizedError *shared.UnauthorizedError
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// Not found
	NotFoundError *shared.NotFoundError
}

func (o *UpdateMeshControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMeshControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMeshControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateMeshControlPlaneResponse) GetMeshControlPlane() *shared.MeshControlPlane {
	if o == nil {
		return nil
	}
	return o.MeshControlPlane
}

func (o *UpdateMeshControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *UpdateMeshControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdateMeshControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdateMeshControlPlaneResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
