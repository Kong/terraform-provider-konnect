// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetMeshOPAListRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
}

func (o *GetMeshOPAListRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshOPAListRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

type GetMeshOPAListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List
	MeshOPAList *shared.MeshOPAList
}

func (o *GetMeshOPAListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshOPAListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshOPAListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshOPAListResponse) GetMeshOPAList() *shared.MeshOPAList {
	if o == nil {
		return nil
	}
	return o.MeshOPAList
}