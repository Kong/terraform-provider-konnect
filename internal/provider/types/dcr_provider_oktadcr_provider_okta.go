// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DCRProviderOKTADCRProviderOKTA struct {
	Active       types.Bool              `tfsdk:"active"`
	CreatedAt    types.String            `tfsdk:"created_at"`
	DcrConfig    Metadata                `tfsdk:"dcr_config"`
	DisplayName  types.String            `tfsdk:"display_name"`
	ID           types.String            `tfsdk:"id"`
	Issuer       types.String            `tfsdk:"issuer"`
	Labels       map[string]types.String `tfsdk:"labels"`
	Name         types.String            `tfsdk:"name"`
	ProviderType types.String            `tfsdk:"provider_type"`
	UpdatedAt    types.String            `tfsdk:"updated_at"`
}
