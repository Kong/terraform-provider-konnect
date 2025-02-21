// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayConsumerResourceModel) ToSharedConsumerInput() *shared.ConsumerInput {
	customID := new(string)
	if !r.CustomID.IsUnknown() && !r.CustomID.IsNull() {
		*customID = r.CustomID.ValueString()
	} else {
		customID = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	username := new(string)
	if !r.Username.IsUnknown() && !r.Username.IsNull() {
		*username = r.Username.ValueString()
	} else {
		username = nil
	}
	out := shared.ConsumerInput{
		CustomID: customID,
		ID:       id,
		Tags:     tags,
		Username: username,
	}
	return &out
}

func (r *GatewayConsumerResourceModel) RefreshFromSharedConsumer(resp *shared.Consumer) {
	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.CustomID = types.StringPointerValue(resp.CustomID)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
		r.Username = types.StringPointerValue(resp.Username)
	}
}
