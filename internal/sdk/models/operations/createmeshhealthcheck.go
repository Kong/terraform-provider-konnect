// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateMeshHealthCheckRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshHealthCheck
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshHealthCheckItem shared.MeshHealthCheckItemInput `request:"mediaType=application/json"`
}

func (o *CreateMeshHealthCheckRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *CreateMeshHealthCheckRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *CreateMeshHealthCheckRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeshHealthCheckRequest) GetMeshHealthCheckItem() shared.MeshHealthCheckItemInput {
	if o == nil {
		return shared.MeshHealthCheckItemInput{}
	}
	return o.MeshHealthCheckItem
}

type CreateMeshHealthCheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	MeshHealthCheckCreateOrUpdateSuccessResponse *shared.MeshHealthCheckCreateOrUpdateSuccessResponse
}

func (o *CreateMeshHealthCheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMeshHealthCheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMeshHealthCheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMeshHealthCheckResponse) GetMeshHealthCheckCreateOrUpdateSuccessResponse() *shared.MeshHealthCheckCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshHealthCheckCreateOrUpdateSuccessResponse
}
