// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AwsVpcPeeringGatewayAttachmentConfig struct {
	Kind          types.String `tfsdk:"kind"`
	PeerAccountID types.String `tfsdk:"peer_account_id"`
	PeerVpcID     types.String `tfsdk:"peer_vpc_id"`
	PeerVpcRegion types.String `tfsdk:"peer_vpc_region"`
}
