// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *CloudGatewayNetworkDataSourceModel) RefreshFromSharedNetwork(resp *shared.Network) {
	if resp != nil {
		r.AvailabilityZones = make([]types.String, 0, len(resp.AvailabilityZones))
		for _, v := range resp.AvailabilityZones {
			r.AvailabilityZones = append(r.AvailabilityZones, types.StringValue(v))
		}
		r.CidrBlock = types.StringValue(resp.CidrBlock)
		r.CloudGatewayProviderAccountID = types.StringValue(resp.CloudGatewayProviderAccountID)
		r.ConfigurationReferenceCount = types.Int64Value(resp.ConfigurationReferenceCount)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Default = types.BoolValue(resp.Default)
		r.EntityVersion = types.Int64Value(resp.EntityVersion)
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.ProviderMetadata.SubnetIds = make([]types.String, 0, len(resp.ProviderMetadata.SubnetIds))
		for _, v := range resp.ProviderMetadata.SubnetIds {
			r.ProviderMetadata.SubnetIds = append(r.ProviderMetadata.SubnetIds, types.StringValue(v))
		}
		r.ProviderMetadata.VpcID = types.StringPointerValue(resp.ProviderMetadata.VpcID)
		r.Region = types.StringValue(resp.Region)
		r.TransitGatewayCount = types.Int64Value(resp.TransitGatewayCount)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}
