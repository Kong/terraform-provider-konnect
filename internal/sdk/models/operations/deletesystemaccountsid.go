// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

var DeleteSystemAccountsIDServerList = []string{
	"https://global.api.konghq.com/",
}

type DeleteSystemAccountsIDRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
}

func (o *DeleteSystemAccountsIDRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type DeleteSystemAccountsIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteSystemAccountsIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSystemAccountsIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSystemAccountsIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteSystemAccountsIDResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *DeleteSystemAccountsIDResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}