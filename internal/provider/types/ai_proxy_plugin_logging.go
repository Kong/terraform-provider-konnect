// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AiProxyPluginLogging struct {
	LogPayloads   types.Bool `tfsdk:"log_payloads"`
	LogStatistics types.Bool `tfsdk:"log_statistics"`
}