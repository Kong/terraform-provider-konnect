// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type GetAuditLogWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Get response for audit log webhook
	AuditLogWebhook *shared.AuditLogWebhook
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
}

func (o *GetAuditLogWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAuditLogWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAuditLogWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAuditLogWebhookResponse) GetAuditLogWebhook() *shared.AuditLogWebhook {
	if o == nil {
		return nil
	}
	return o.AuditLogWebhook
}

func (o *GetAuditLogWebhookResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetAuditLogWebhookResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}
