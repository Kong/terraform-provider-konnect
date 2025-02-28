// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshControlPlaneDataSourceModel) RefreshFromSharedMeshControlPlane(resp *shared.MeshControlPlane) {
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
				features1.HostnameGeneratorCreation.Enabled = types.BoolPointerValue(featuresItem.HostnameGeneratorCreation.Enabled)
			}
			if featuresItem.MeshCreation == nil {
				features1.MeshCreation = nil
			} else {
				features1.MeshCreation = &tfTypes.MeshControlPlaneFeatureHostnameGenerationCreation{}
				features1.MeshCreation.Enabled = types.BoolPointerValue(featuresItem.MeshCreation.Enabled)
			}
			if featuresItem.Type != nil {
				features1.Type = types.StringValue(string(*featuresItem.Type))
			} else {
				features1.Type = types.StringNull()
			}
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
