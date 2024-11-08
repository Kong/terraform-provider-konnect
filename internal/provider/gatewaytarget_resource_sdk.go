// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayTargetResourceModel) ToSharedTargetWithoutParents() *shared.TargetWithoutParents {
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
	target := new(string)
	if !r.Target.IsUnknown() && !r.Target.IsNull() {
		*target = r.Target.ValueString()
	} else {
		target = nil
	}
	weight := new(int64)
	if !r.Weight.IsUnknown() && !r.Weight.IsNull() {
		*weight = r.Weight.ValueInt64()
	} else {
		weight = nil
	}
	out := shared.TargetWithoutParents{
		ID:     id,
		Tags:   tags,
		Target: target,
		Weight: weight,
	}
	return &out
}

func (r *GatewayTargetResourceModel) RefreshFromSharedTarget(resp *shared.Target) {
	if resp != nil {
		if resp.CreatedAt != nil {
			r.CreatedAt = types.NumberValue(big.NewFloat(float64(*resp.CreatedAt)))
		} else {
			r.CreatedAt = types.NumberNull()
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Target = types.StringPointerValue(resp.Target)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.NumberValue(big.NewFloat(float64(*resp.UpdatedAt)))
		} else {
			r.UpdatedAt = types.NumberNull()
		}
		if resp.Upstream == nil {
			r.Upstream = nil
		} else {
			r.Upstream = &tfTypes.ACLConsumer{}
			r.Upstream.ID = types.StringPointerValue(resp.Upstream.ID)
		}
		r.Weight = types.Int64PointerValue(resp.Weight)
	}
}
