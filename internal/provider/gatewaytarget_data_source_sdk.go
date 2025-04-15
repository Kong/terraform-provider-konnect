// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayTargetDataSourceModel) RefreshFromSharedTarget(ctx context.Context, resp *shared.Target) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.Float64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Target = types.StringPointerValue(resp.Target)
		r.UpdatedAt = types.Float64PointerValue(resp.UpdatedAt)
		if resp.Upstream == nil {
			r.Upstream = nil
		} else {
			r.Upstream = &tfTypes.ACLWithoutParentsConsumer{}
			r.Upstream.ID = types.StringPointerValue(resp.Upstream.ID)
		}
		r.Weight = types.Int64PointerValue(resp.Weight)
	}

	return diags
}
