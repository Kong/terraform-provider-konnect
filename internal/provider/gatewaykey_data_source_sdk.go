// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayKeyDataSourceModel) RefreshFromSharedKey(resp *shared.Key) {
	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Jwk = types.StringPointerValue(resp.Jwk)
		r.Kid = types.StringValue(resp.Kid)
		r.Name = types.StringPointerValue(resp.Name)
		if resp.Pem == nil {
			r.Pem = nil
		} else {
			r.Pem = &tfTypes.Pem{}
			r.Pem.PrivateKey = types.StringPointerValue(resp.Pem.PrivateKey)
			r.Pem.PublicKey = types.StringPointerValue(resp.Pem.PublicKey)
		}
		if resp.Set == nil {
			r.Set = nil
		} else {
			r.Set = &tfTypes.ACLConsumer{}
			r.Set.ID = types.StringPointerValue(resp.Set.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
