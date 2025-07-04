// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayACLDataSourceModel) ToOperationsGetACLWithConsumerRequest(ctx context.Context) (*operations.GetACLWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	var aclID string
	aclID = r.ID.ValueString()

	out := operations.GetACLWithConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
		ACLID:          aclID,
	}

	return &out, diags
}

func (r *GatewayACLDataSourceModel) RefreshFromSharedACL(ctx context.Context, resp *shared.ACL) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Group = types.StringValue(resp.Group)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}

	return diags
}
