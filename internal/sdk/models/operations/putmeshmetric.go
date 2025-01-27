// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type PutMeshMetricRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshMetric
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshMetricItem shared.MeshMetricItem `request:"mediaType=application/json"`
}

func (o *PutMeshMetricRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *PutMeshMetricRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *PutMeshMetricRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PutMeshMetricRequest) GetMeshMetricItem() shared.MeshMetricItem {
	if o == nil {
		return shared.MeshMetricItem{}
	}
	return o.MeshMetricItem
}

type PutMeshMetricResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	MeshMetricCreateOrUpdateSuccessResponse *shared.MeshMetricCreateOrUpdateSuccessResponse
}

func (o *PutMeshMetricResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutMeshMetricResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutMeshMetricResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutMeshMetricResponse) GetMeshMetricCreateOrUpdateSuccessResponse() *shared.MeshMetricCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshMetricCreateOrUpdateSuccessResponse
}
