// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateRequestTerminationPluginConfig struct {
	Body        types.String `tfsdk:"body"`
	ContentType types.String `tfsdk:"content_type"`
	Echo        types.Bool   `tfsdk:"echo"`
	Message     types.String `tfsdk:"message"`
	StatusCode  types.Int64  `tfsdk:"status_code"`
	Trigger     types.String `tfsdk:"trigger"`
}
