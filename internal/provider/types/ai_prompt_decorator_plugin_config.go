// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AiPromptDecoratorPluginConfig struct {
	MaxRequestBodySize types.Int64 `tfsdk:"max_request_body_size"`
	Prompts            *Prompts    `tfsdk:"prompts"`
}