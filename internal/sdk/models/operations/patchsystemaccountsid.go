// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

var PatchSystemAccountsIDServerList = []string{
	"https://global.api.konghq.com/",
}

type PatchSystemAccountsIDRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// The request schema for the update system account request.
	UpdateSystemAccount *shared.UpdateSystemAccount `request:"mediaType=application/json"`
}

func (o *PatchSystemAccountsIDRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *PatchSystemAccountsIDRequest) GetUpdateSystemAccount() *shared.UpdateSystemAccount {
	if o == nil {
		return nil
	}
	return o.UpdateSystemAccount
}

type PatchSystemAccountsIDResponse struct {
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
	// Conflict
	ConflictError *shared.ConflictError
}

func (o *PatchSystemAccountsIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchSystemAccountsIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchSystemAccountsIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchSystemAccountsIDResponse) GetSystemAccount() *shared.SystemAccount {
	if o == nil {
		return nil
	}
	return o.SystemAccount
}

func (o *PatchSystemAccountsIDResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *PatchSystemAccountsIDResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *PatchSystemAccountsIDResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}
