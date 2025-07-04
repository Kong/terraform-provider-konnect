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

func (r *GatewaySNIDataSourceModel) ToOperationsGetSniRequest(ctx context.Context) (*operations.GetSniRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var sniID string
	sniID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetSniRequest{
		SNIID:          sniID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewaySNIDataSourceModel) RefreshFromSharedSni(ctx context.Context, resp *shared.Sni) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Certificate == nil {
			r.Certificate = nil
		} else {
			r.Certificate = &tfTypes.Set{}
			r.Certificate.ID = types.StringPointerValue(resp.Certificate.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}

	return diags
}
