// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"time"
)

func (r *SystemAccountResourceModel) ToSharedCreateSystemAccount() *shared.CreateSystemAccount {
	var name string
	name = r.Name.ValueString()

	var description string
	description = r.Description.ValueString()

	konnectManaged := new(bool)
	if !r.KonnectManaged.IsUnknown() && !r.KonnectManaged.IsNull() {
		*konnectManaged = r.KonnectManaged.ValueBool()
	} else {
		konnectManaged = nil
	}
	out := shared.CreateSystemAccount{
		Name:           name,
		Description:    description,
		KonnectManaged: konnectManaged,
	}
	return &out
}

func (r *SystemAccountResourceModel) RefreshFromSharedSystemAccount(resp *shared.SystemAccount) {
	if resp != nil {
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.KonnectManaged = types.BoolPointerValue(resp.KonnectManaged)
		r.Name = types.StringPointerValue(resp.Name)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}

func (r *SystemAccountResourceModel) ToSharedUpdateSystemAccount() *shared.UpdateSystemAccount {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	out := shared.UpdateSystemAccount{
		Name:        name,
		Description: description,
	}
	return &out
}
