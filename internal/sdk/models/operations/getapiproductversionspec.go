// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetAPIProductVersionSpecRequest struct {
	// The API product identifier
	APIProductID string `pathParam:"style=simple,explode=false,name=apiProductId"`
	// The API product version identifier
	APIProductVersionID string `pathParam:"style=simple,explode=false,name=apiProductVersionId"`
	// The API product version specification identifier
	SpecificationID string `pathParam:"style=simple,explode=false,name=specificationId"`
}

func (o *GetAPIProductVersionSpecRequest) GetAPIProductID() string {
	if o == nil {
		return ""
	}
	return o.APIProductID
}

func (o *GetAPIProductVersionSpecRequest) GetAPIProductVersionID() string {
	if o == nil {
		return ""
	}
	return o.APIProductVersionID
}

func (o *GetAPIProductVersionSpecRequest) GetSpecificationID() string {
	if o == nil {
		return ""
	}
	return o.SpecificationID
}

type GetAPIProductVersionSpecResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// API product version specification
	APIProductVersionSpec *shared.APIProductVersionSpec
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetAPIProductVersionSpecResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIProductVersionSpecResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIProductVersionSpecResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPIProductVersionSpecResponse) GetAPIProductVersionSpec() *shared.APIProductVersionSpec {
	if o == nil {
		return nil
	}
	return o.APIProductVersionSpec
}

func (o *GetAPIProductVersionSpecResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetAPIProductVersionSpecResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetAPIProductVersionSpecResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
