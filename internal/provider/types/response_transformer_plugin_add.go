// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ResponseTransformerPluginAdd struct {
	Headers   []types.String `tfsdk:"headers"`
	JSON      []types.String `tfsdk:"json"`
	JSONTypes []types.String `tfsdk:"json_types"`
}