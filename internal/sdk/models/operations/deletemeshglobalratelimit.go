// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type DeleteMeshGlobalRateLimitRequest struct {
	// Id of the Konnect resource
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshGlobalRateLimit
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *DeleteMeshGlobalRateLimitRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteMeshGlobalRateLimitRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *DeleteMeshGlobalRateLimitRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type DeleteMeshGlobalRateLimitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteMeshGlobalRateLimitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMeshGlobalRateLimitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMeshGlobalRateLimitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMeshGlobalRateLimitResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
