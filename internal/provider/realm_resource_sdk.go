// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/provider/typeconvert"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *RealmResourceModel) ToSharedConsumerRealmCreateRequest(ctx context.Context) (*shared.ConsumerRealmCreateRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var name string
	name = r.Name.ValueString()

	allowedControlPlanes := make([]string, 0, len(r.AllowedControlPlanes))
	for _, allowedControlPlanesItem := range r.AllowedControlPlanes {
		allowedControlPlanes = append(allowedControlPlanes, allowedControlPlanesItem.ValueString())
	}
	allowAllControlPlanes := new(bool)
	if !r.AllowAllControlPlanes.IsUnknown() && !r.AllowAllControlPlanes.IsNull() {
		*allowAllControlPlanes = r.AllowAllControlPlanes.ValueBool()
	} else {
		allowAllControlPlanes = nil
	}
	ttl := new(int64)
	if !r.TTL.IsUnknown() && !r.TTL.IsNull() {
		*ttl = r.TTL.ValueInt64()
	} else {
		ttl = nil
	}
	negativeTTL := new(int64)
	if !r.NegativeTTL.IsUnknown() && !r.NegativeTTL.IsNull() {
		*negativeTTL = r.NegativeTTL.ValueInt64()
	} else {
		negativeTTL = nil
	}
	consumerGroups := make([]string, 0, len(r.ConsumerGroups))
	for _, consumerGroupsItem := range r.ConsumerGroups {
		consumerGroups = append(consumerGroups, consumerGroupsItem.ValueString())
	}
	out := shared.ConsumerRealmCreateRequest{
		Name:                  name,
		AllowedControlPlanes:  allowedControlPlanes,
		AllowAllControlPlanes: allowAllControlPlanes,
		TTL:                   ttl,
		NegativeTTL:           negativeTTL,
		ConsumerGroups:        consumerGroups,
	}

	return &out, diags
}

func (r *RealmResourceModel) ToSharedConsumerRealmUpdateRequest(ctx context.Context) (*shared.ConsumerRealmUpdateRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	allowedControlPlanes := make([]string, 0, len(r.AllowedControlPlanes))
	for _, allowedControlPlanesItem := range r.AllowedControlPlanes {
		allowedControlPlanes = append(allowedControlPlanes, allowedControlPlanesItem.ValueString())
	}
	allowAllControlPlanes := new(bool)
	if !r.AllowAllControlPlanes.IsUnknown() && !r.AllowAllControlPlanes.IsNull() {
		*allowAllControlPlanes = r.AllowAllControlPlanes.ValueBool()
	} else {
		allowAllControlPlanes = nil
	}
	consumerGroups := make([]string, 0, len(r.ConsumerGroups))
	for _, consumerGroupsItem := range r.ConsumerGroups {
		consumerGroups = append(consumerGroups, consumerGroupsItem.ValueString())
	}
	ttl := new(int64)
	if !r.TTL.IsUnknown() && !r.TTL.IsNull() {
		*ttl = r.TTL.ValueInt64()
	} else {
		ttl = nil
	}
	negativeTTL := new(int64)
	if !r.NegativeTTL.IsUnknown() && !r.NegativeTTL.IsNull() {
		*negativeTTL = r.NegativeTTL.ValueInt64()
	} else {
		negativeTTL = nil
	}
	out := shared.ConsumerRealmUpdateRequest{
		Name:                  name,
		AllowedControlPlanes:  allowedControlPlanes,
		AllowAllControlPlanes: allowAllControlPlanes,
		ConsumerGroups:        consumerGroups,
		TTL:                   ttl,
		NegativeTTL:           negativeTTL,
	}

	return &out, diags
}

func (r *RealmResourceModel) ToOperationsUpdateRealmRequest(ctx context.Context) (*operations.UpdateRealmRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var realmID string
	realmID = r.ID.ValueString()

	consumerRealmUpdateRequest, consumerRealmUpdateRequestDiags := r.ToSharedConsumerRealmUpdateRequest(ctx)
	diags.Append(consumerRealmUpdateRequestDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateRealmRequest{
		RealmID:                    realmID,
		ConsumerRealmUpdateRequest: *consumerRealmUpdateRequest,
	}

	return &out, diags
}

func (r *RealmResourceModel) ToOperationsGetRealmRequest(ctx context.Context) (*operations.GetRealmRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var realmID string
	realmID = r.ID.ValueString()

	out := operations.GetRealmRequest{
		RealmID: realmID,
	}

	return &out, diags
}

func (r *RealmResourceModel) ToOperationsDeleteRealmRequest(ctx context.Context) (*operations.DeleteRealmRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var realmID string
	realmID = r.ID.ValueString()

	forceDestroy := new(operations.Force)
	if !r.ForceDestroy.IsUnknown() && !r.ForceDestroy.IsNull() {
		*forceDestroy = operations.Force(r.ForceDestroy.ValueString())
	} else {
		forceDestroy = nil
	}
	out := operations.DeleteRealmRequest{
		RealmID:      realmID,
		ForceDestroy: forceDestroy,
	}

	return &out, diags
}

func (r *RealmResourceModel) RefreshFromSharedConsumerRealm(ctx context.Context, resp *shared.ConsumerRealm) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.AllowAllControlPlanes = types.BoolValue(resp.AllowAllControlPlanes)
		r.AllowedControlPlanes = make([]types.String, 0, len(resp.AllowedControlPlanes))
		for _, v := range resp.AllowedControlPlanes {
			r.AllowedControlPlanes = append(r.AllowedControlPlanes, types.StringValue(v))
		}
		r.ConsumerGroups = make([]types.String, 0, len(resp.ConsumerGroups))
		for _, v := range resp.ConsumerGroups {
			r.ConsumerGroups = append(r.ConsumerGroups, types.StringValue(v))
		}
		r.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.CreatedAt))
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.NegativeTTL = types.Int64PointerValue(resp.NegativeTTL)
		r.TTL = types.Int64PointerValue(resp.TTL)
		r.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.UpdatedAt))
	}

	return diags
}
