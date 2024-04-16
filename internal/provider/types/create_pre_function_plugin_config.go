// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreatePreFunctionPluginConfig struct {
	Access          []types.String `tfsdk:"access"`
	BodyFilter      []types.String `tfsdk:"body_filter"`
	Certificate     []types.String `tfsdk:"certificate"`
	HeaderFilter    []types.String `tfsdk:"header_filter"`
	Log             []types.String `tfsdk:"log"`
	Rewrite         []types.String `tfsdk:"rewrite"`
	WsClientFrame   []types.String `tfsdk:"ws_client_frame"`
	WsClose         []types.String `tfsdk:"ws_close"`
	WsHandshake     []types.String `tfsdk:"ws_handshake"`
	WsUpstreamFrame []types.String `tfsdk:"ws_upstream_frame"`
}
