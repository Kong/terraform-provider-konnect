// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateDataplaneCertificateRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Request body for creating a dp-client-certificate.
	DataPlaneClientCertificateRequest *shared.DataPlaneClientCertificateRequest `request:"mediaType=application/json"`
}

func (o *CreateDataplaneCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateDataplaneCertificateRequest) GetDataPlaneClientCertificateRequest() *shared.DataPlaneClientCertificateRequest {
	if o == nil {
		return nil
	}
	return o.DataPlaneClientCertificateRequest
}

type CreateDataplaneCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Response body for retrieving a dp-client-certificate.
	DataPlaneClientCertificate *shared.DataPlaneClientCertificate
}

func (o *CreateDataplaneCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDataplaneCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDataplaneCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDataplaneCertificateResponse) GetDataPlaneClientCertificate() *shared.DataPlaneClientCertificate {
	if o == nil {
		return nil
	}
	return o.DataPlaneClientCertificate
}
