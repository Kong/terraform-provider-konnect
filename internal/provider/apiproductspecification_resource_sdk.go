// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"time"
)

func (r *APIProductSpecificationResourceModel) ToSharedCreateAPIProductVersionSpecDTO() *shared.CreateAPIProductVersionSpecDTO {
	content := r.Content.ValueString()
	name := r.Name.ValueString()
	out := shared.CreateAPIProductVersionSpecDTO{
		Content: content,
		Name:    name,
	}
	return &out
}

func (r *APIProductSpecificationResourceModel) RefreshFromSharedAPIProductVersionSpec(resp *shared.APIProductVersionSpec) {
	if resp != nil {
		r.Content = types.StringValue(resp.Content)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}

func (r *APIProductSpecificationResourceModel) ToSharedUpdateAPIProductVersionSpecDTO() *shared.UpdateAPIProductVersionSpecDTO {
	content := new(string)
	if !r.Content.IsUnknown() && !r.Content.IsNull() {
		*content = r.Content.ValueString()
	} else {
		content = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	out := shared.UpdateAPIProductVersionSpecDTO{
		Content: content,
		Name:    name,
	}
	return &out
}
