// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type PutMeshProxyPatchRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshProxyPatch
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshProxyPatchItem shared.MeshProxyPatchItem `request:"mediaType=application/json"`
}

func (o *PutMeshProxyPatchRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *PutMeshProxyPatchRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *PutMeshProxyPatchRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PutMeshProxyPatchRequest) GetMeshProxyPatchItem() shared.MeshProxyPatchItem {
	if o == nil {
		return shared.MeshProxyPatchItem{}
	}
	return o.MeshProxyPatchItem
}

type PutMeshProxyPatchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	MeshProxyPatchCreateOrUpdateSuccessResponse *shared.MeshProxyPatchCreateOrUpdateSuccessResponse
}

func (o *PutMeshProxyPatchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutMeshProxyPatchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutMeshProxyPatchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutMeshProxyPatchResponse) GetMeshProxyPatchCreateOrUpdateSuccessResponse() *shared.MeshProxyPatchCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshProxyPatchCreateOrUpdateSuccessResponse
}
