// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetHostnameGeneratorRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the HostnameGenerator
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetHostnameGeneratorRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetHostnameGeneratorRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetHostnameGeneratorResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	HostnameGeneratorItem *shared.HostnameGeneratorItem
}

func (o *GetHostnameGeneratorResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHostnameGeneratorResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHostnameGeneratorResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHostnameGeneratorResponse) GetHostnameGeneratorItem() *shared.HostnameGeneratorItem {
	if o == nil {
		return nil
	}
	return o.HostnameGeneratorItem
}
