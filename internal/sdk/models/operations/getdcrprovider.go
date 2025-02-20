// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetDcrProviderRequest struct {
	// DCR provider identifier
	DcrProviderID string `pathParam:"style=simple,explode=false,name=dcrProviderId"`
}

func (o *GetDcrProviderRequest) GetDcrProviderID() string {
	if o == nil {
		return ""
	}
	return o.DcrProviderID
}

type GetDcrProviderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response containing a single DCR provider object. Sensitive fields will be removed from the response.
	DcrProviderResponse *shared.DcrProviderResponse
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetDcrProviderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDcrProviderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDcrProviderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDcrProviderResponse) GetDcrProviderResponse() *shared.DcrProviderResponse {
	if o == nil {
		return nil
	}
	return o.DcrProviderResponse
}

func (o *GetDcrProviderResponse) GetDcrProviderResponseAuth0() *shared.DCRProviderAuth0DCRProviderAuth0 {
	if v := o.GetDcrProviderResponse(); v != nil {
		return v.DCRProviderAuth0DCRProviderAuth0
	}
	return nil
}

func (o *GetDcrProviderResponse) GetDcrProviderResponseAzureAd() *shared.DCRProviderAzureADDCRProviderAzureAD {
	if v := o.GetDcrProviderResponse(); v != nil {
		return v.DCRProviderAzureADDCRProviderAzureAD
	}
	return nil
}

func (o *GetDcrProviderResponse) GetDcrProviderResponseCurity() *shared.DCRProviderCurityDCRProviderCurity {
	if v := o.GetDcrProviderResponse(); v != nil {
		return v.DCRProviderCurityDCRProviderCurity
	}
	return nil
}

func (o *GetDcrProviderResponse) GetDcrProviderResponseOkta() *shared.DCRProviderOKTADCRProviderOKTA {
	if v := o.GetDcrProviderResponse(); v != nil {
		return v.DCRProviderOKTADCRProviderOKTA
	}
	return nil
}

func (o *GetDcrProviderResponse) GetDcrProviderResponseHTTP() *shared.DCRProviderHTTPDCRProviderHTTP {
	if v := o.GetDcrProviderResponse(); v != nil {
		return v.DCRProviderHTTPDCRProviderHTTP
	}
	return nil
}

func (o *GetDcrProviderResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetDcrProviderResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetDcrProviderResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
