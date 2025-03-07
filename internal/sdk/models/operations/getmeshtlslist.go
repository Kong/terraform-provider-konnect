// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

// GetMeshTLSListQueryParamFilter - filter by labels when multiple filters are present, they are ANDed
type GetMeshTLSListQueryParamFilter struct {
	Key   *string `queryParam:"name=key"`
	Value *string `queryParam:"name=value"`
}

func (o *GetMeshTLSListQueryParamFilter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMeshTLSListQueryParamFilter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetMeshTLSListRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// offset in the list of entities
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// the number of items per page
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// filter by labels when multiple filters are present, they are ANDed
	Filter *GetMeshTLSListQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
}

func (g GetMeshTLSListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMeshTLSListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMeshTLSListRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshTLSListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetMeshTLSListRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetMeshTLSListRequest) GetFilter() *GetMeshTLSListQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetMeshTLSListRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

type GetMeshTLSListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List
	MeshTLSList *shared.MeshTLSList
}

func (o *GetMeshTLSListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshTLSListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshTLSListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshTLSListResponse) GetMeshTLSList() *shared.MeshTLSList {
	if o == nil {
		return nil
	}
	return o.MeshTLSList
}
