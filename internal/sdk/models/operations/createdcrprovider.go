// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateDcrProviderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response containing the newly created DCR provider object.
	CreateDcrProviderResponse *shared.CreateDcrProviderResponse
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
}

func (o *CreateDcrProviderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDcrProviderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDcrProviderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDcrProviderResponse) GetCreateDcrProviderResponse() *shared.CreateDcrProviderResponse {
	if o == nil {
		return nil
	}
	return o.CreateDcrProviderResponse
}

func (o *CreateDcrProviderResponse) GetCreateDcrProviderResponseAuth0() *shared.DcrProviderAuth0 {
	if v := o.GetCreateDcrProviderResponse(); v != nil {
		return v.DcrProviderAuth0
	}
	return nil
}

func (o *CreateDcrProviderResponse) GetCreateDcrProviderResponseAzureAd() *shared.DcrProviderAzureAd {
	if v := o.GetCreateDcrProviderResponse(); v != nil {
		return v.DcrProviderAzureAd
	}
	return nil
}

func (o *CreateDcrProviderResponse) GetCreateDcrProviderResponseCurity() *shared.DcrProviderCurity {
	if v := o.GetCreateDcrProviderResponse(); v != nil {
		return v.DcrProviderCurity
	}
	return nil
}

func (o *CreateDcrProviderResponse) GetCreateDcrProviderResponseOkta() *shared.DcrProviderOkta {
	if v := o.GetCreateDcrProviderResponse(); v != nil {
		return v.DcrProviderOkta
	}
	return nil
}

func (o *CreateDcrProviderResponse) GetCreateDcrProviderResponseHTTP() *shared.DcrProviderHTTP {
	if v := o.GetCreateDcrProviderResponse(); v != nil {
		return v.DcrProviderHTTP
	}
	return nil
}

func (o *CreateDcrProviderResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *CreateDcrProviderResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateDcrProviderResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}
