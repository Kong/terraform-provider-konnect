// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type NetworkProviderMetadata struct {
	SubnetIds []types.String `tfsdk:"subnet_ids"`
	VpcID     types.String   `tfsdk:"vpc_id"`
}