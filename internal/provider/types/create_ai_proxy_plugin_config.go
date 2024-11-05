// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateAiProxyPluginConfig struct {
	Auth               *CreateAiProxyPluginAuth    `tfsdk:"auth"`
	Logging            *CreateAiProxyPluginLogging `tfsdk:"logging"`
	MaxRequestBodySize types.Int64                 `tfsdk:"max_request_body_size"`
	Model              *CreateAiProxyPluginModel   `tfsdk:"model"`
	ModelNameHeader    types.Bool                  `tfsdk:"model_name_header"`
	ResponseStreaming  types.String                `tfsdk:"response_streaming"`
	RouteType          types.String                `tfsdk:"route_type"`
}