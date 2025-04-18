// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/provider/typeconvert"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
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
	labels := make(map[string]*string)
	for labelsKey, labelsValue := range r.Labels {
		labelsInst := new(string)
		if !labelsValue.IsUnknown() && !labelsValue.IsNull() {
			*labelsInst = labelsValue.ValueString()
		} else {
			labelsInst = nil
		}
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

func (r *MeshControlPlaneResourceModel) RefreshFromSharedMeshControlPlane(ctx context.Context, resp *shared.MeshControlPlane) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.CreatedAt))
		r.Description = types.StringPointerValue(resp.Description)
		r.Features = []tfTypes.MeshControlPlaneFeature{}
		if len(r.Features) > len(resp.Features) {
			r.Features = r.Features[:len(resp.Features)]
		}
		for featuresCount, featuresItem := range resp.Features {
			var features tfTypes.MeshControlPlaneFeature
			if featuresItem.HostnameGeneratorCreation == nil {
				features.HostnameGeneratorCreation = nil
			} else {
				features.HostnameGeneratorCreation = &tfTypes.MeshControlPlaneFeatureHostnameGenerationCreation{}
				features.HostnameGeneratorCreation.Enabled = types.BoolValue(featuresItem.HostnameGeneratorCreation.Enabled)
			}
			if featuresItem.MeshCreation == nil {
				features.MeshCreation = nil
			} else {
				features.MeshCreation = &tfTypes.MeshControlPlaneFeatureHostnameGenerationCreation{}
				features.MeshCreation.Enabled = types.BoolValue(featuresItem.MeshCreation.Enabled)
			}
			features.Type = types.StringValue(string(featuresItem.Type))
			if featuresCount+1 > len(r.Features) {
				r.Features = append(r.Features, features)
			} else {
				r.Features[featuresCount].HostnameGeneratorCreation = features.HostnameGeneratorCreation
				r.Features[featuresCount].MeshCreation = features.MeshCreation
				r.Features[featuresCount].Type = features.Type
			}
		}
		r.ID = types.StringValue(resp.ID)
		if resp.Labels != nil {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringPointerValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.UpdatedAt))
	}

	return diags
}

func (r *MeshControlPlaneResourceModel) ToSharedPutMeshControlPlaneRequest() *shared.PutMeshControlPlaneRequest {
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
	out := shared.PutMeshControlPlaneRequest{
		Name:        name,
		Description: description,
		Labels:      labels,
	}
	return &out
}
