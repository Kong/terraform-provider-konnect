// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayConsumerGroupResourceModel) ToSharedConsumerGroupInput() *shared.ConsumerGroupInput {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.ConsumerGroupInput{
		Name: name,
		Tags: tags,
	}
	return &out
}

func (r *GatewayConsumerGroupResourceModel) RefreshFromSharedConsumerGroup(resp *shared.ConsumerGroup) {
	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
