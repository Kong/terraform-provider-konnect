// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

// GetMeshMetricListQueryParamFilter - filter by labels when multiple filters are present, they are ANDed
type GetMeshMetricListQueryParamFilter struct {
	Key   *string `queryParam:"name=key"`
	Value *string `queryParam:"name=value"`
}

func (o *GetMeshMetricListQueryParamFilter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMeshMetricListQueryParamFilter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetMeshMetricListRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// filter by labels when multiple filters are present, they are ANDed
	Filter *GetMeshMetricListQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
}

func (o *GetMeshMetricListRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshMetricListRequest) GetFilter() *GetMeshMetricListQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetMeshMetricListRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

type GetMeshMetricListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List
	MeshMetricList *shared.MeshMetricList
}

func (o *GetMeshMetricListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshMetricListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshMetricListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshMetricListResponse) GetMeshMetricList() *shared.MeshMetricList {
	if o == nil {
		return nil
	}
	return o.MeshMetricList
}
