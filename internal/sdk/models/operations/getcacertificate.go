// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type GetCaCertificateRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the CA Certificate to lookup
	CACertificateID string `pathParam:"style=simple,explode=false,name=CACertificateId"`
}

func (o *GetCaCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetCaCertificateRequest) GetCACertificateID() string {
	if o == nil {
		return ""
	}
	return o.CACertificateID
}

type GetCaCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched CA Certificate
	CACertificate *shared.CACertificate
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetCaCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCaCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCaCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCaCertificateResponse) GetCACertificate() *shared.CACertificate {
	if o == nil {
		return nil
	}
	return o.CACertificate
}

func (o *GetCaCertificateResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
