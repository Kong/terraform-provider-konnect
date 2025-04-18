// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/provider/typeconvert"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *PortalTeamDataSourceModel) RefreshFromSharedPortalTeamResponse(ctx context.Context, resp *shared.PortalTeamResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreatedAt))
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.UpdatedAt))
	}

	return diags
}
