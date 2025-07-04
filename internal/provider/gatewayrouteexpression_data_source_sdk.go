// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayRouteExpressionDataSourceModel) ToOperationsGetRouteRouteExpressionRequest(ctx context.Context) (*operations.GetRouteRouteExpressionRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var routeID string
	routeID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetRouteRouteExpressionRequest{
		RouteID:        routeID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayRouteExpressionDataSourceModel) RefreshFromSharedRouteExpression(ctx context.Context, resp *shared.RouteExpression) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Expression = types.StringPointerValue(resp.Expression)
		if resp.HTTPSRedirectStatusCode != nil {
			r.HTTPSRedirectStatusCode = types.Int64Value(int64(*resp.HTTPSRedirectStatusCode))
		} else {
			r.HTTPSRedirectStatusCode = types.Int64Null()
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		if resp.PathHandling != nil {
			r.PathHandling = types.StringValue(string(*resp.PathHandling))
		} else {
			r.PathHandling = types.StringNull()
		}
		r.PreserveHost = types.BoolPointerValue(resp.PreserveHost)
		r.Priority = types.Int64PointerValue(resp.Priority)
		if resp.Protocols != nil {
			r.Protocols = make([]types.String, 0, len(resp.Protocols))
			for _, v := range resp.Protocols {
				r.Protocols = append(r.Protocols, types.StringValue(string(v)))
			}
		}
		r.RequestBuffering = types.BoolPointerValue(resp.RequestBuffering)
		r.ResponseBuffering = types.BoolPointerValue(resp.ResponseBuffering)
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.Set{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.StripPath = types.BoolPointerValue(resp.StripPath)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}

	return diags
}
