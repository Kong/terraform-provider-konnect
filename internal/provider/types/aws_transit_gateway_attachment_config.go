// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AwsTransitGatewayAttachmentConfig struct {
	Kind             types.String `tfsdk:"kind"`
	RAMShareArn      types.String `tfsdk:"ram_share_arn"`
	TransitGatewayID types.String `tfsdk:"transit_gateway_id"`
}
