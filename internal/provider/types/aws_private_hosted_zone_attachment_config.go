// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AwsPrivateHostedZoneAttachmentConfig struct {
	HostedZoneID types.String `tfsdk:"hosted_zone_id"`
	Kind         types.String `tfsdk:"kind"`
}
