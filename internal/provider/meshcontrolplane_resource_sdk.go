// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshControlPlaneResourceModel) ToSharedCreateMeshControlPlaneRequest() *shared.CreateMeshControlPlaneRequest {
	var name string
	name = r.Name.ValueString()

	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		var labelsInst string
		labelsInst = labelsValue.ValueString()

		labels[labelsKey] = labelsInst
	}
	out := shared.CreateMeshControlPlaneRequest{
		Name:        name,
		Description: description,
		Labels:      labels,
	}
	return &out
}

func (r *MeshControlPlaneResourceModel) RefreshFromSharedMeshControlPlane(resp *shared.MeshControlPlane) {
	if resp != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringValue(resp.ID)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}

func (r *MeshControlPlaneResourceModel) ToSharedUpdateMeshControlPlaneRequest() *shared.UpdateMeshControlPlaneRequest {
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
	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		var labelsInst string
		labelsInst = labelsValue.ValueString()

		labels[labelsKey] = labelsInst
	}
	out := shared.UpdateMeshControlPlaneRequest{
		Name:        name,
		Description: description,
		Labels:      labels,
	}
	return &out
}
