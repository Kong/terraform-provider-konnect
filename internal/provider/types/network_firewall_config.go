// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type NetworkFirewallConfig struct {
	AllowedCidrBlocks []types.String `tfsdk:"allowed_cidr_blocks"`
	DeniedCidrBlocks  []types.String `tfsdk:"denied_cidr_blocks"`
}
