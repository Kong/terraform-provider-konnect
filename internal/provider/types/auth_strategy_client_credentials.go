// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AuthStrategyClientCredentials struct {
	AuthMethods    []types.String `tfsdk:"auth_methods"`
	CredentialType types.String   `tfsdk:"credential_type"`
	ID             types.String   `tfsdk:"id"`
	Name           types.String   `tfsdk:"name"`
}
