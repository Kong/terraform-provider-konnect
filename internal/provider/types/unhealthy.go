// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Unhealthy struct {
	HTTPFailures types.Int64   `tfsdk:"http_failures"`
	HTTPStatuses []types.Int64 `tfsdk:"http_statuses"`
	Interval     types.Number  `tfsdk:"interval"`
	TCPFailures  types.Int64   `tfsdk:"tcp_failures"`
	Timeouts     types.Int64   `tfsdk:"timeouts"`
}
