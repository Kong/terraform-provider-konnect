// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayCustomPluginStreamingDataSourceModel) ToOperationsGetCustomPluginRequest(ctx context.Context) (*operations.GetCustomPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var customPluginID string
	customPluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetCustomPluginRequest{
		CustomPluginID: customPluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayCustomPluginStreamingDataSourceModel) RefreshFromSharedCustomPlugin(ctx context.Context, resp *shared.CustomPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Handler = types.StringValue(resp.Handler)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Schema = types.StringValue(resp.Schema)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}

	return diags
}
