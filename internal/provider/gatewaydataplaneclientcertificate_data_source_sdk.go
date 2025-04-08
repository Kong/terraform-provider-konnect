// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayDataPlaneClientCertificateDataSourceModel) RefreshFromSharedDataPlaneClientCertificate(ctx context.Context, resp *shared.DataPlaneClientCertificate) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Cert = types.StringPointerValue(resp.Cert)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}

	return diags
}
