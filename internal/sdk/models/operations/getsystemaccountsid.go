// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

var GetSystemAccountsIDServerList = []string{
	"https://global.api.konghq.com/",
}

type GetSystemAccountsIDRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
}

func (o *GetSystemAccountsIDRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type GetSystemAccountsIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response including a single system account.
	SystemAccount *shared.SystemAccount
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetSystemAccountsIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSystemAccountsIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSystemAccountsIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSystemAccountsIDResponse) GetSystemAccount() *shared.SystemAccount {
	if o == nil {
		return nil
	}
	return o.SystemAccount
}

func (o *GetSystemAccountsIDResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetSystemAccountsIDResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
