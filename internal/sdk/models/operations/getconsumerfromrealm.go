// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetConsumerFromRealmRequest struct {
	// ID of the realm
	RealmID string `pathParam:"style=simple,explode=false,name=realmId"`
	// ID of the of the consumer
	ConsumerID string `pathParam:"style=simple,explode=false,name=consumerId"`
}

func (o *GetConsumerFromRealmRequest) GetRealmID() string {
	if o == nil {
		return ""
	}
	return o.RealmID
}

func (o *GetConsumerFromRealmRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

type GetConsumerFromRealmResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// success
	CentralizedConsumer *shared.CentralizedConsumer
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetConsumerFromRealmResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConsumerFromRealmResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConsumerFromRealmResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConsumerFromRealmResponse) GetCentralizedConsumer() *shared.CentralizedConsumer {
	if o == nil {
		return nil
	}
	return o.CentralizedConsumer
}

func (o *GetConsumerFromRealmResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
