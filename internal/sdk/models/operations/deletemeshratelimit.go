// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type DeleteMeshRateLimitRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshRateLimit
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *DeleteMeshRateLimitRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *DeleteMeshRateLimitRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *DeleteMeshRateLimitRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type DeleteMeshRateLimitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteMeshRateLimitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMeshRateLimitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMeshRateLimitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMeshRateLimitResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}