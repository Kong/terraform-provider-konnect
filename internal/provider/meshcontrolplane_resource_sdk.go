// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
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
	var features []shared.MeshControlPlaneFeature = []shared.MeshControlPlaneFeature{}
	for _, featuresItem := range r.Features {
		typeVar := shared.Type(featuresItem.Type.ValueString())
		var hostnameGeneratorCreation *shared.MeshControlPlaneFeatureHostnameGenerationCreation
		if featuresItem.HostnameGeneratorCreation != nil {
			var enabled bool
			enabled = featuresItem.HostnameGeneratorCreation.Enabled.ValueBool()

			hostnameGeneratorCreation = &shared.MeshControlPlaneFeatureHostnameGenerationCreation{
				Enabled: enabled,
			}
		}
		var meshCreation *shared.MeshControlPlaneFeatureMeshCreation
		if featuresItem.MeshCreation != nil {
			var enabled1 bool
			enabled1 = featuresItem.MeshCreation.Enabled.ValueBool()

			meshCreation = &shared.MeshControlPlaneFeatureMeshCreation{
				Enabled: enabled1,
			}
		}
		features = append(features, shared.MeshControlPlaneFeature{
			Type:                      typeVar,
			HostnameGeneratorCreation: hostnameGeneratorCreation,
			MeshCreation:              meshCreation,
		})
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
		Features:    features,
		Labels:      labels,
	}
	return &out
}

func (r *MeshControlPlaneResourceModel) RefreshFromSharedMeshControlPlane(resp *shared.MeshControlPlane) {
	if resp != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Description = types.StringPointerValue(resp.Description)
		r.Features = []tfTypes.MeshControlPlaneFeature{}
		if len(r.Features) > len(resp.Features) {
			r.Features = r.Features[:len(resp.Features)]
		}
		for featuresCount, featuresItem := range resp.Features {
			var features1 tfTypes.MeshControlPlaneFeature
			if featuresItem.HostnameGeneratorCreation == nil {
				features1.HostnameGeneratorCreation = nil
			} else {
				features1.HostnameGeneratorCreation = &tfTypes.MeshControlPlaneFeatureHostnameGenerationCreation{}
				features1.HostnameGeneratorCreation.Enabled = types.BoolValue(featuresItem.HostnameGeneratorCreation.Enabled)
			}
			if featuresItem.MeshCreation == nil {
				features1.MeshCreation = nil
			} else {
				features1.MeshCreation = &tfTypes.MeshControlPlaneFeatureHostnameGenerationCreation{}
				features1.MeshCreation.Enabled = types.BoolValue(featuresItem.MeshCreation.Enabled)
			}
			features1.Type = types.StringValue(string(featuresItem.Type))
			if featuresCount+1 > len(r.Features) {
				r.Features = append(r.Features, features1)
			} else {
				r.Features[featuresCount].HostnameGeneratorCreation = features1.HostnameGeneratorCreation
				r.Features[featuresCount].MeshCreation = features1.MeshCreation
				r.Features[featuresCount].Type = features1.Type
			}
		}
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
