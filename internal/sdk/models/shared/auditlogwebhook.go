// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// AuditLogWebhook - Get response for audit log webhook
type AuditLogWebhook struct {
	// The endpoint that will receive audit log messages.
	Endpoint *string `json:"endpoint,omitempty"`
	// Indicates whether audit data should be sent to the webhook.
	Enabled *bool `json:"enabled,omitempty"`
	// The output format of each log messages.
	LogFormat *LogFormat `default:"cef" json:"log_format"`
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

func (o *AuditLogWebhook) GetLogFormat() *LogFormat {
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
