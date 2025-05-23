// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpsertUpstreamRequest struct {
	// ID of the Upstream to lookup
	UpstreamID string `pathParam:"style=simple,explode=false,name=UpstreamId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the Upstream
	Upstream shared.Upstream `request:"mediaType=application/json"`
}

func (o *UpsertUpstreamRequest) GetUpstreamID() string {
	if o == nil {
		return ""
	}
	return o.UpstreamID
}

func (o *UpsertUpstreamRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertUpstreamRequest) GetUpstream() shared.Upstream {
	if o == nil {
		return shared.Upstream{}
	}
	return o.Upstream
}

type UpsertUpstreamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Upstream
	Upstream *shared.Upstream
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertUpstreamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertUpstreamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertUpstreamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertUpstreamResponse) GetUpstream() *shared.Upstream {
	if o == nil {
		return nil
	}
	return o.Upstream
}

func (o *UpsertUpstreamResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
