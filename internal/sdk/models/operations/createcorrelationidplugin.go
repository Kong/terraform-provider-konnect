// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateCorrelationidPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID            string                            `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateCorrelationIDPlugin *shared.CreateCorrelationIDPlugin `request:"mediaType=application/json"`
}

func (o *CreateCorrelationidPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateCorrelationidPluginRequest) GetCreateCorrelationIDPlugin() *shared.CreateCorrelationIDPlugin {
	if o == nil {
		return nil
	}
	return o.CreateCorrelationIDPlugin
}

type CreateCorrelationidPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// CorrelationId plugin
	CorrelationIDPlugin *shared.CorrelationIDPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreateCorrelationidPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCorrelationidPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCorrelationidPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateCorrelationidPluginResponse) GetCorrelationIDPlugin() *shared.CorrelationIDPlugin {
	if o == nil {
		return nil
	}
	return o.CorrelationIDPlugin
}

func (o *CreateCorrelationidPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
