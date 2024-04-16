// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayConsumerDataSourceModel) RefreshFromSharedConsumer(resp *shared.Consumer) {
	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.CustomID = types.StringPointerValue(resp.CustomID)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Username = types.StringPointerValue(resp.Username)
	}
}
