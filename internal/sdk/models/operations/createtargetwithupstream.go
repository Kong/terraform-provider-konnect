// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateTargetWithUpstreamRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID or target of the Target to lookup
	UpstreamID string `pathParam:"style=simple,explode=false,name=UpstreamIdForTarget"`
	// Description of new Target for creation
	TargetWithoutParents shared.TargetWithoutParents `request:"mediaType=application/json"`
}

func (o *CreateTargetWithUpstreamRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateTargetWithUpstreamRequest) GetUpstreamID() string {
	if o == nil {
		return ""
	}
	return o.UpstreamID
}

func (o *CreateTargetWithUpstreamRequest) GetTargetWithoutParents() shared.TargetWithoutParents {
	if o == nil {
		return shared.TargetWithoutParents{}
	}
	return o.TargetWithoutParents
}

type CreateTargetWithUpstreamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Target
	Target *shared.Target
}

func (o *CreateTargetWithUpstreamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTargetWithUpstreamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTargetWithUpstreamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTargetWithUpstreamResponse) GetTarget() *shared.Target {
	if o == nil {
		return nil
	}
	return o.Target
}
