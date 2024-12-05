// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AzureTransitGateway struct {
	DNSConfig                      []TransitGatewayDNSConfig        `tfsdk:"dns_config"`
	Name                           types.String                     `tfsdk:"name"`
	TransitGatewayAttachmentConfig AzureVNETPeeringAttachmentConfig `tfsdk:"transit_gateway_attachment_config"`
}