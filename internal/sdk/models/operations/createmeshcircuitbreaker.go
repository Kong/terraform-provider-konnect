// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateMeshCircuitBreakerRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshCircuitBreaker
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshCircuitBreakerItem shared.MeshCircuitBreakerItemInput `request:"mediaType=application/json"`
}

func (o *CreateMeshCircuitBreakerRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *CreateMeshCircuitBreakerRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *CreateMeshCircuitBreakerRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeshCircuitBreakerRequest) GetMeshCircuitBreakerItem() shared.MeshCircuitBreakerItemInput {
	if o == nil {
		return shared.MeshCircuitBreakerItemInput{}
	}
	return o.MeshCircuitBreakerItem
}

type CreateMeshCircuitBreakerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	MeshCircuitBreakerCreateOrUpdateSuccessResponse *shared.MeshCircuitBreakerCreateOrUpdateSuccessResponse
}

func (o *CreateMeshCircuitBreakerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMeshCircuitBreakerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMeshCircuitBreakerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMeshCircuitBreakerResponse) GetMeshCircuitBreakerCreateOrUpdateSuccessResponse() *shared.MeshCircuitBreakerCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshCircuitBreakerCreateOrUpdateSuccessResponse
}
