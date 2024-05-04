// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"time"
)

func (r *SystemAccountAccessTokenDataSourceModel) RefreshFromSharedSystemAccountAccessToken(resp *shared.SystemAccountAccessToken) {
	if resp != nil {
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		if resp.ExpiresAt != nil {
			r.ExpiresAt = types.StringValue(resp.ExpiresAt.Format(time.RFC3339Nano))
		} else {
			r.ExpiresAt = types.StringNull()
		}
		r.ID = types.StringPointerValue(resp.ID)
		if resp.LastUsedAt != nil {
			r.LastUsedAt = types.StringValue(resp.LastUsedAt.Format(time.RFC3339Nano))
		} else {
			r.LastUsedAt = types.StringNull()
		}
		r.Name = types.StringPointerValue(resp.Name)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}
