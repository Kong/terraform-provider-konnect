// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// AuditLogWebhookLogFormat - The output format of each log messages.
type AuditLogWebhookLogFormat string

const (
	AuditLogWebhookLogFormatCef  AuditLogWebhookLogFormat = "cef"
	AuditLogWebhookLogFormatJSON AuditLogWebhookLogFormat = "json"
)

func (e AuditLogWebhookLogFormat) ToPointer() *AuditLogWebhookLogFormat {
	return &e
}
func (e *AuditLogWebhookLogFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cef":
		fallthrough
	case "json":
		*e = AuditLogWebhookLogFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuditLogWebhookLogFormat: %v", v)
	}
}

// AuditLogWebhook - Get response for audit log webhook
type AuditLogWebhook struct {
	// The endpoint that will receive audit log messages.
	Endpoint *string `json:"endpoint,omitempty"`
	// Indicates whether audit data should be sent to the webhook.
	Enabled *bool `json:"enabled,omitempty"`
	// The output format of each log messages.
	LogFormat *AuditLogWebhookLogFormat `json:"log_format,omitempty"`
	// Indicates if the SSL certificate verification of the host endpoint should be skipped when delivering payloads.
	SkipSslVerification *bool `json:"skip_ssl_verification,omitempty"`
	// Timestamp when this webhook was last updated. Initial value is 0001-01-01T00:00:0Z.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (a AuditLogWebhook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuditLogWebhook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuditLogWebhook) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *AuditLogWebhook) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AuditLogWebhook) GetLogFormat() *AuditLogWebhookLogFormat {
	if o == nil {
		return nil
	}
	return o.LogFormat
}

func (o *AuditLogWebhook) GetSkipSslVerification() *bool {
	if o == nil {
		return nil
	}
	return o.SkipSslVerification
}

func (o *AuditLogWebhook) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
