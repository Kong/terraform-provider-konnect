// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayVaultDataSourceModel) RefreshFromSharedVault(resp *shared.Vault) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = types.StringNull()
		} else {
			configResult, _ := json.Marshal(resp.Config)
			r.Config = types.StringValue(string(configResult))
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Prefix = types.StringValue(resp.Prefix)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
