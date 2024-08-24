// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type DeleteAppAuthStrategyRequest struct {
	// Application auth strategy identifier
	AuthStrategyID string `pathParam:"style=simple,explode=false,name=authStrategyId"`
}

func (o *DeleteAppAuthStrategyRequest) GetAuthStrategyID() string {
	if o == nil {
		return ""
	}
	return o.AuthStrategyID
}

type DeleteAppAuthStrategyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Conflict
	ConflictError *shared.ConflictError
}

func (o *DeleteAppAuthStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAppAuthStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAppAuthStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteAppAuthStrategyResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *DeleteAppAuthStrategyResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *DeleteAppAuthStrategyResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *DeleteAppAuthStrategyResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}
