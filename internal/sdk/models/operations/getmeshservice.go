// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetMeshServiceRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshService
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshServiceRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshServiceRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshServiceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshServiceItem *shared.MeshServiceItem
}

func (o *GetMeshServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshServiceResponse) GetMeshServiceItem() *shared.MeshServiceItem {
	if o == nil {
		return nil
	}
	return o.MeshServiceItem
}
