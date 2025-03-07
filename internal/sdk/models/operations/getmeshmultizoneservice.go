// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetMeshMultiZoneServiceRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshMultiZoneService
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshMultiZoneServiceRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshMultiZoneServiceRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshMultiZoneServiceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshMultiZoneServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshMultiZoneServiceItem *shared.MeshMultiZoneServiceItem
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetMeshMultiZoneServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshMultiZoneServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshMultiZoneServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshMultiZoneServiceResponse) GetMeshMultiZoneServiceItem() *shared.MeshMultiZoneServiceItem {
	if o == nil {
		return nil
	}
	return o.MeshMultiZoneServiceItem
}

func (o *GetMeshMultiZoneServiceResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
