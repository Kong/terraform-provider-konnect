// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayKeyAuthResourceModel) ToSharedKeyAuthWithoutParents(ctx context.Context) (*shared.KeyAuthWithoutParents, diag.Diagnostics) {
	var diags diag.Diagnostics

	var consumer *shared.KeyAuthWithoutParentsConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.KeyAuthWithoutParentsConsumer{
			ID: id,
		}
	}
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	id1 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueString()
	} else {
		id1 = nil
	}
	var key string
	key = r.Key.ValueString()

	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	ttl := new(int64)
	if !r.TTL.IsUnknown() && !r.TTL.IsNull() {
		*ttl = r.TTL.ValueInt64()
	} else {
		ttl = nil
	}
	out := shared.KeyAuthWithoutParents{
		Consumer:  consumer,
		CreatedAt: createdAt,
		ID:        id1,
		Key:       key,
		Tags:      tags,
		TTL:       ttl,
	}

	return &out, diags
}

func (r *GatewayKeyAuthResourceModel) ToOperationsCreateKeyAuthWithConsumerRequest(ctx context.Context) (*operations.CreateKeyAuthWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	keyAuthWithoutParents, keyAuthWithoutParentsDiags := r.ToSharedKeyAuthWithoutParents(ctx)
	diags.Append(keyAuthWithoutParentsDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateKeyAuthWithConsumerRequest{
		ControlPlaneID:        controlPlaneID,
		ConsumerID:            consumerID,
		KeyAuthWithoutParents: *keyAuthWithoutParents,
	}

	return &out, diags
}

func (r *GatewayKeyAuthResourceModel) ToOperationsGetKeyAuthWithConsumerRequest(ctx context.Context) (*operations.GetKeyAuthWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	var keyAuthID string
	keyAuthID = r.ID.ValueString()

	out := operations.GetKeyAuthWithConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
		KeyAuthID:      keyAuthID,
	}

	return &out, diags
}

func (r *GatewayKeyAuthResourceModel) ToOperationsDeleteKeyAuthWithConsumerRequest(ctx context.Context) (*operations.DeleteKeyAuthWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	var keyAuthID string
	keyAuthID = r.ID.ValueString()

	out := operations.DeleteKeyAuthWithConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
		KeyAuthID:      keyAuthID,
	}

	return &out, diags
}

func (r *GatewayKeyAuthResourceModel) RefreshFromSharedKeyAuth(ctx context.Context, resp *shared.KeyAuth) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringValue(resp.Key)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.TTL = types.Int64PointerValue(resp.TTL)
	}

	return diags
}
