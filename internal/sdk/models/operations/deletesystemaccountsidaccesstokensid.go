// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

var DeleteSystemAccountsIDAccessTokensIDServerList = []string{
	"https://global.api.konghq.com/",
}

type DeleteSystemAccountsIDAccessTokensIDRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// ID of the system account access token.
	TokenID string `pathParam:"style=simple,explode=false,name=tokenId"`
}

func (o *DeleteSystemAccountsIDAccessTokensIDRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *DeleteSystemAccountsIDAccessTokensIDRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

type DeleteSystemAccountsIDAccessTokensIDResponse struct {
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

func (o *DeleteSystemAccountsIDAccessTokensIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSystemAccountsIDAccessTokensIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSystemAccountsIDAccessTokensIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteSystemAccountsIDAccessTokensIDResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *DeleteSystemAccountsIDAccessTokensIDResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
