// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type PutMeshGlobalRateLimitRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshGlobalRateLimit
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshGlobalRateLimitItem shared.MeshGlobalRateLimitItem `request:"mediaType=application/json"`
}

func (o *PutMeshGlobalRateLimitRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *PutMeshGlobalRateLimitRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *PutMeshGlobalRateLimitRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PutMeshGlobalRateLimitRequest) GetMeshGlobalRateLimitItem() shared.MeshGlobalRateLimitItem {
	if o == nil {
		return shared.MeshGlobalRateLimitItem{}
	}
	return o.MeshGlobalRateLimitItem
}

type PutMeshGlobalRateLimitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	MeshGlobalRateLimitCreateOrUpdateSuccessResponse *shared.MeshGlobalRateLimitCreateOrUpdateSuccessResponse
}

func (o *PutMeshGlobalRateLimitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutMeshGlobalRateLimitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutMeshGlobalRateLimitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutMeshGlobalRateLimitResponse) GetMeshGlobalRateLimitCreateOrUpdateSuccessResponse() *shared.MeshGlobalRateLimitCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshGlobalRateLimitCreateOrUpdateSuccessResponse
}
