// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateBasicAuthWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerIDForNestedEntities string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// Description of new Basic-auth credential for creation
	CreateBasicAuthWithoutParents shared.CreateBasicAuthWithoutParents `request:"mediaType=application/json"`
}

func (o *CreateBasicAuthWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateBasicAuthWithConsumerRequest) GetConsumerIDForNestedEntities() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDForNestedEntities
}

func (o *CreateBasicAuthWithConsumerRequest) GetCreateBasicAuthWithoutParents() shared.CreateBasicAuthWithoutParents {
	if o == nil {
		return shared.CreateBasicAuthWithoutParents{}
	}
	return o.CreateBasicAuthWithoutParents
}

type CreateBasicAuthWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Basic-auth credential
	BasicAuth *shared.BasicAuth
}

func (o *CreateBasicAuthWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBasicAuthWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBasicAuthWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateBasicAuthWithConsumerResponse) GetBasicAuth() *shared.BasicAuth {
	if o == nil {
		return nil
	}
	return o.BasicAuth
}
