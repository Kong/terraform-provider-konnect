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

func (r *GatewayKeyResourceModel) ToSharedKey(ctx context.Context) (*shared.Key, diag.Diagnostics) {
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
	jwk := new(string)
	if !r.Jwk.IsUnknown() && !r.Jwk.IsNull() {
		*jwk = r.Jwk.ValueString()
	} else {
		jwk = nil
	}
	var kid string
	kid = r.Kid.ValueString()

	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	var pem *shared.Pem
	if r.Pem != nil {
		privateKey := new(string)
		if !r.Pem.PrivateKey.IsUnknown() && !r.Pem.PrivateKey.IsNull() {
			*privateKey = r.Pem.PrivateKey.ValueString()
		} else {
			privateKey = nil
		}
		publicKey := new(string)
		if !r.Pem.PublicKey.IsUnknown() && !r.Pem.PublicKey.IsNull() {
			*publicKey = r.Pem.PublicKey.ValueString()
		} else {
			publicKey = nil
		}
		pem = &shared.Pem{
			PrivateKey: privateKey,
			PublicKey:  publicKey,
		}
	}
	var set *shared.Set
	if r.Set != nil {
		id1 := new(string)
		if !r.Set.ID.IsUnknown() && !r.Set.ID.IsNull() {
			*id1 = r.Set.ID.ValueString()
		} else {
			id1 = nil
		}
		set = &shared.Set{
			ID: id1,
		}
	}
	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	x5t := new(string)
	if !r.X5t.IsUnknown() && !r.X5t.IsNull() {
		*x5t = r.X5t.ValueString()
	} else {
		x5t = nil
	}
	out := shared.Key{
		CreatedAt: createdAt,
		ID:        id,
		Jwk:       jwk,
		Kid:       kid,
		Name:      name,
		Pem:       pem,
		Set:       set,
		Tags:      tags,
		UpdatedAt: updatedAt,
		X5t:       x5t,
	}

	return &out, diags
}

func (r *GatewayKeyResourceModel) ToOperationsCreateKeyRequest(ctx context.Context) (*operations.CreateKeyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	key, keyDiags := r.ToSharedKey(ctx)
	diags.Append(keyDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateKeyRequest{
		ControlPlaneID: controlPlaneID,
		Key:            *key,
	}

	return &out, diags
}

func (r *GatewayKeyResourceModel) ToOperationsUpsertKeyRequest(ctx context.Context) (*operations.UpsertKeyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var keyID string
	keyID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	key, keyDiags := r.ToSharedKey(ctx)
	diags.Append(keyDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpsertKeyRequest{
		KeyID:          keyID,
		ControlPlaneID: controlPlaneID,
		Key:            *key,
	}

	return &out, diags
}

func (r *GatewayKeyResourceModel) ToOperationsGetKeyRequest(ctx context.Context) (*operations.GetKeyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var keyID string
	keyID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetKeyRequest{
		KeyID:          keyID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayKeyResourceModel) ToOperationsDeleteKeyRequest(ctx context.Context) (*operations.DeleteKeyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	var keyID string
	keyID = r.ID.ValueString()

	out := operations.DeleteKeyRequest{
		ControlPlaneID: controlPlaneID,
		KeyID:          keyID,
	}

	return &out, diags
}

func (r *GatewayKeyResourceModel) RefreshFromSharedKey(ctx context.Context, resp *shared.Key) diag.Diagnostics {
	var diags diag.Diagnostics

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
			r.Set = &tfTypes.Set{}
			r.Set.ID = types.StringPointerValue(resp.Set.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
		r.X5t = types.StringPointerValue(resp.X5t)
	}

	return diags
}
