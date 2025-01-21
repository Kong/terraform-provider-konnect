// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetMeshOPARequest struct {
	// Id of the Konnect resource
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshOPA
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshOPARequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetMeshOPARequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshOPARequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshOPAResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshOPAItem *shared.MeshOPAItem
}

func (o *GetMeshOPAResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshOPAResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshOPAResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshOPAResponse) GetMeshOPAItem() *shared.MeshOPAItem {
	if o == nil {
		return nil
	}
	return o.MeshOPAItem
}
