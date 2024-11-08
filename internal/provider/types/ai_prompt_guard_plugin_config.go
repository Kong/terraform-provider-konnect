// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AiPromptGuardPluginConfig struct {
	AllowAllConversationHistory types.Bool     `tfsdk:"allow_all_conversation_history"`
	AllowPatterns               []types.String `tfsdk:"allow_patterns"`
	DenyPatterns                []types.String `tfsdk:"deny_patterns"`
	MatchAllRoles               types.Bool     `tfsdk:"match_all_roles"`
	MaxRequestBodySize          types.Int64    `tfsdk:"max_request_body_size"`
}