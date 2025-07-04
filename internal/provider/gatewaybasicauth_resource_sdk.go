// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayBasicAuthResourceModel) ToSharedBasicAuthWithoutParents(ctx context.Context) (*shared.BasicAuthWithoutParents, diag.Diagnostics) {
	var diags diag.Diagnostics

	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var password string
	password = r.Password.ValueString()

	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var username string
	username = r.Username.ValueString()

	out := shared.BasicAuthWithoutParents{
		CreatedAt: createdAt,
		ID:        id,
		Password:  password,
		Tags:      tags,
		Username:  username,
	}

	return &out, diags
}

func (r *GatewayBasicAuthResourceModel) ToOperationsCreateBasicAuthWithConsumerRequest(ctx context.Context) (*operations.CreateBasicAuthWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	basicAuthWithoutParents, basicAuthWithoutParentsDiags := r.ToSharedBasicAuthWithoutParents(ctx)
	diags.Append(basicAuthWithoutParentsDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateBasicAuthWithConsumerRequest{
		ControlPlaneID:          controlPlaneID,
		ConsumerID:              consumerID,
		BasicAuthWithoutParents: *basicAuthWithoutParents,
	}

	return &out, diags
}

func (r *GatewayBasicAuthResourceModel) ToOperationsGetBasicAuthWithConsumerRequest(ctx context.Context) (*operations.GetBasicAuthWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	var basicAuthID string
	basicAuthID = r.ID.ValueString()

	out := operations.GetBasicAuthWithConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
		BasicAuthID:    basicAuthID,
	}

	return &out, diags
}

func (r *GatewayBasicAuthResourceModel) ToOperationsDeleteBasicAuthWithConsumerRequest(ctx context.Context) (*operations.DeleteBasicAuthWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	var basicAuthID string
	basicAuthID = r.ID.ValueString()

	out := operations.DeleteBasicAuthWithConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
		BasicAuthID:    basicAuthID,
	}

	return &out, diags
}

func (r *GatewayBasicAuthResourceModel) RefreshFromSharedBasicAuth(ctx context.Context, resp *shared.BasicAuth) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Password = types.StringValue(resp.Password)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Username = types.StringValue(resp.Username)
	}

	return diags
}
