// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshRetryItemHTTP struct {
	BackOff                  *BackOff            `tfsdk:"back_off"`
	HostSelection            []HostSelection     `tfsdk:"host_selection"`
	HostSelectionMaxAttempts types.Int64         `tfsdk:"host_selection_max_attempts"`
	NumRetries               types.Int64         `tfsdk:"num_retries"`
	PerTryTimeout            types.String        `tfsdk:"per_try_timeout"`
	RateLimitedBackOff       *RateLimitedBackOff `tfsdk:"rate_limited_back_off"`
	RetriableRequestHeaders  []Headers           `tfsdk:"retriable_request_headers"`
	RetriableResponseHeaders []Headers           `tfsdk:"retriable_response_headers"`
	RetryOn                  []types.String      `tfsdk:"retry_on"`
}