// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AcmePluginVault struct {
	AuthMethod    types.String `tfsdk:"auth_method"`
	AuthPath      types.String `tfsdk:"auth_path"`
	AuthRole      types.String `tfsdk:"auth_role"`
	Host          types.String `tfsdk:"host"`
	HTTPS         types.Bool   `tfsdk:"https"`
	JwtPath       types.String `tfsdk:"jwt_path"`
	KvPath        types.String `tfsdk:"kv_path"`
	Port          types.Int64  `tfsdk:"port"`
	Timeout       types.Number `tfsdk:"timeout"`
	TLSServerName types.String `tfsdk:"tls_server_name"`
	TLSVerify     types.Bool   `tfsdk:"tls_verify"`
	Token         types.String `tfsdk:"token"`
}