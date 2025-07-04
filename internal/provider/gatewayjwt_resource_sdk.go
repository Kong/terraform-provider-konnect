// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayJWTResourceModel) ToSharedJWTWithoutParents(ctx context.Context) (*shared.JWTWithoutParents, diag.Diagnostics) {
	var diags diag.Diagnostics

	algorithm := new(shared.Algorithm)
	if !r.Algorithm.IsUnknown() && !r.Algorithm.IsNull() {
		*algorithm = shared.Algorithm(r.Algorithm.ValueString())
	} else {
		algorithm = nil
	}
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
	key := new(string)
	if !r.Key.IsUnknown() && !r.Key.IsNull() {
		*key = r.Key.ValueString()
	} else {
		key = nil
	}
	rsaPublicKey := new(string)
	if !r.RsaPublicKey.IsUnknown() && !r.RsaPublicKey.IsNull() {
		*rsaPublicKey = r.RsaPublicKey.ValueString()
	} else {
		rsaPublicKey = nil
	}
	secret := new(string)
	if !r.Secret.IsUnknown() && !r.Secret.IsNull() {
		*secret = r.Secret.ValueString()
	} else {
		secret = nil
	}
	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.JWTWithoutParents{
		Algorithm:    algorithm,
		CreatedAt:    createdAt,
		ID:           id,
		Key:          key,
		RsaPublicKey: rsaPublicKey,
		Secret:       secret,
		Tags:         tags,
	}

	return &out, diags
}

func (r *GatewayJWTResourceModel) ToOperationsCreateJwtWithConsumerRequest(ctx context.Context) (*operations.CreateJwtWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	jwtWithoutParents, jwtWithoutParentsDiags := r.ToSharedJWTWithoutParents(ctx)
	diags.Append(jwtWithoutParentsDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateJwtWithConsumerRequest{
		ControlPlaneID:    controlPlaneID,
		ConsumerID:        consumerID,
		JWTWithoutParents: jwtWithoutParents,
	}

	return &out, diags
}

func (r *GatewayJWTResourceModel) ToOperationsGetJwtWithConsumerRequest(ctx context.Context) (*operations.GetJwtWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	var jwtID string
	jwtID = r.ID.ValueString()

	out := operations.GetJwtWithConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
		JWTID:          jwtID,
	}

	return &out, diags
}

func (r *GatewayJWTResourceModel) ToOperationsDeleteJwtWithConsumerRequest(ctx context.Context) (*operations.DeleteJwtWithConsumerRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var consumerID string
	consumerID = r.ConsumerID.ValueString()

	var jwtID string
	jwtID = r.ID.ValueString()

	out := operations.DeleteJwtWithConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
		JWTID:          jwtID,
	}

	return &out, diags
}

func (r *GatewayJWTResourceModel) RefreshFromSharedJwt(ctx context.Context, resp *shared.Jwt) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Algorithm != nil {
			r.Algorithm = types.StringValue(string(*resp.Algorithm))
		} else {
			r.Algorithm = types.StringNull()
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringPointerValue(resp.Key)
		r.RsaPublicKey = types.StringPointerValue(resp.RsaPublicKey)
		r.Secret = types.StringPointerValue(resp.Secret)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}

	return diags
}
