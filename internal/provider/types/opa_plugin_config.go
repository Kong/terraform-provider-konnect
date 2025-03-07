// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OpaPluginConfig struct {
	IncludeBodyInOpaInput           types.Bool   `tfsdk:"include_body_in_opa_input"`
	IncludeConsumerInOpaInput       types.Bool   `tfsdk:"include_consumer_in_opa_input"`
	IncludeParsedJSONBodyInOpaInput types.Bool   `tfsdk:"include_parsed_json_body_in_opa_input"`
	IncludeRouteInOpaInput          types.Bool   `tfsdk:"include_route_in_opa_input"`
	IncludeServiceInOpaInput        types.Bool   `tfsdk:"include_service_in_opa_input"`
	IncludeURICapturesInOpaInput    types.Bool   `tfsdk:"include_uri_captures_in_opa_input"`
	OpaHost                         types.String `tfsdk:"opa_host"`
	OpaPath                         types.String `tfsdk:"opa_path"`
	OpaPort                         types.Int64  `tfsdk:"opa_port"`
	OpaProtocol                     types.String `tfsdk:"opa_protocol"`
	SslVerify                       types.Bool   `tfsdk:"ssl_verify"`
}
