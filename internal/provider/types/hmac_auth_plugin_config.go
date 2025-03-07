// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type HmacAuthPluginConfig struct {
	Algorithms          []types.String `tfsdk:"algorithms"`
	Anonymous           types.String   `tfsdk:"anonymous"`
	ClockSkew           types.Number   `tfsdk:"clock_skew"`
	EnforceHeaders      []types.String `tfsdk:"enforce_headers"`
	HideCredentials     types.Bool     `tfsdk:"hide_credentials"`
	Realm               types.String   `tfsdk:"realm"`
	ValidateRequestBody types.Bool     `tfsdk:"validate_request_body"`
}
