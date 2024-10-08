// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteACLWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the ACL to lookup
	ACLID string `pathParam:"style=simple,explode=false,name=ACLId"`
}

func (o *DeleteACLWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *DeleteACLWithConsumerRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *DeleteACLWithConsumerRequest) GetACLID() string {
	if o == nil {
		return ""
	}
	return o.ACLID
}

type DeleteACLWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteACLWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteACLWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteACLWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
