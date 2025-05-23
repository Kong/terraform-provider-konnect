// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/provider/typeconvert"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *AuditLogDataSourceModel) RefreshFromSharedAuditLogWebhook(ctx context.Context, resp *shared.AuditLogWebhook) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.Endpoint = types.StringPointerValue(resp.Endpoint)
		if resp.LogFormat != nil {
			r.LogFormat = types.StringValue(string(*resp.LogFormat))
		} else {
			r.LogFormat = types.StringNull()
		}
		r.SkipSslVerification = types.BoolPointerValue(resp.SkipSslVerification)
		r.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.UpdatedAt))
	}

	return diags
}
