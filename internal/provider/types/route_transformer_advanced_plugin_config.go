// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type RouteTransformerAdvancedPluginConfig struct {
	EscapePath types.Bool   `tfsdk:"escape_path"`
	Host       types.String `tfsdk:"host"`
	Path       types.String `tfsdk:"path"`
	Port       types.String `tfsdk:"port"`
}
