// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Healthy struct {
	HTTPStatuses []types.Int64 `tfsdk:"http_statuses"`
	Interval     types.Float64 `tfsdk:"interval"`
	Successes    types.Int64   `tfsdk:"successes"`
}
