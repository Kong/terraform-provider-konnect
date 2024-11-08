// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type TCPLogPluginConfig struct {
	CustomFieldsByLua map[string]types.String `tfsdk:"custom_fields_by_lua"`
	Host              types.String            `tfsdk:"host"`
	Keepalive         types.Number            `tfsdk:"keepalive"`
	Port              types.Int64             `tfsdk:"port"`
	Timeout           types.Number            `tfsdk:"timeout"`
	TLS               types.Bool              `tfsdk:"tls"`
	TLSSni            types.String            `tfsdk:"tls_sni"`
}