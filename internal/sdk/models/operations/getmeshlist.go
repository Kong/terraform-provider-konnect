// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

// QueryParamFilter - filter by labels when multiple filters are present, they are ANDed
type QueryParamFilter struct {
	Key   *string `queryParam:"name=key"`
	Value *string `queryParam:"name=value"`
}

func (o *QueryParamFilter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *QueryParamFilter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetMeshListRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// filter by labels when multiple filters are present, they are ANDed
	Filter *QueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
}

func (o *GetMeshListRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshListRequest) GetFilter() *QueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetMeshListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List
	MeshList *shared.MeshList
}

func (o *GetMeshListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshListResponse) GetMeshList() *shared.MeshList {
	if o == nil {
		return nil
	}
	return o.MeshList
}
