// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VaultAuthPluginConfig struct {
	AccessTokenName types.String `tfsdk:"access_token_name"`
	Anonymous       types.String `tfsdk:"anonymous"`
	HideCredentials types.Bool   `tfsdk:"hide_credentials"`
	RunOnPreflight  types.Bool   `tfsdk:"run_on_preflight"`
	SecretTokenName types.String `tfsdk:"secret_token_name"`
	TokensInBody    types.Bool   `tfsdk:"tokens_in_body"`
	Vault           types.String `tfsdk:"vault"`
}
