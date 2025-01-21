// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetMeshTraceRequest struct {
	// Id of the Konnect resource
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshTrace
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshTraceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetMeshTraceRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshTraceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshTraceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshTraceItem *shared.MeshTraceItem
}

func (o *GetMeshTraceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshTraceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshTraceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshTraceResponse) GetMeshTraceItem() *shared.MeshTraceItem {
	if o == nil {
		return nil
	}
	return o.MeshTraceItem
}
