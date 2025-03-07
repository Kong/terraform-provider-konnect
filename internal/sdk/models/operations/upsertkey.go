// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpsertKeyRequest struct {
	// ID of the Key to lookup
	KeyID string `pathParam:"style=simple,explode=false,name=KeyId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the Key
	Key shared.KeyInput `request:"mediaType=application/json"`
}

func (o *UpsertKeyRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *UpsertKeyRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertKeyRequest) GetKey() shared.KeyInput {
	if o == nil {
		return shared.KeyInput{}
	}
	return o.Key
}

type UpsertKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Key
	Key *shared.Key
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertKeyResponse) GetKey() *shared.Key {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *UpsertKeyResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
