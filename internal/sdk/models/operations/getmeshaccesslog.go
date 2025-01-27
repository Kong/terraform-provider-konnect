// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetMeshAccessLogRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshAccessLog
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshAccessLogRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshAccessLogRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshAccessLogRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshAccessLogResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshAccessLogItem *shared.MeshAccessLogItem
}

func (o *GetMeshAccessLogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshAccessLogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshAccessLogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshAccessLogResponse) GetMeshAccessLogItem() *shared.MeshAccessLogItem {
	if o == nil {
		return nil
	}
	return o.MeshAccessLogItem
}
