// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"time"
)

func (r *APIProductVersionResourceModel) ToSharedCreateAPIProductVersionDTO() *shared.CreateAPIProductVersionDTO {
	name := r.Name.ValueString()
	deprecated := new(bool)
	if !r.Deprecated.IsUnknown() && !r.Deprecated.IsNull() {
		*deprecated = r.Deprecated.ValueBool()
	} else {
		deprecated = nil
	}
	var gatewayService *shared.GatewayServicePayload
	if r.GatewayService != nil {
		id := r.GatewayService.ID.ValueString()
		controlPlaneID := r.GatewayService.ControlPlaneID.ValueString()
		gatewayService = &shared.GatewayServicePayload{
			ID:             id,
			ControlPlaneID: controlPlaneID,
		}
	}
	out := shared.CreateAPIProductVersionDTO{
		Name:           name,
		Deprecated:     deprecated,
		GatewayService: gatewayService,
	}
	return &out
}

func (r *APIProductVersionResourceModel) RefreshFromSharedAPIProductVersion(resp *shared.APIProductVersion) {
	if resp != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Deprecated = types.BoolValue(resp.Deprecated)
		if resp.GatewayService == nil {
			r.GatewayService = nil
		} else {
			r.GatewayService = &tfTypes.GatewayServicePayload{}
			r.GatewayService.ControlPlaneID = types.StringValue(resp.GatewayService.ControlPlaneID)
			r.GatewayService.ID = types.StringPointerValue(resp.GatewayService.ID)
			r.GatewayService.RuntimeGroupID = types.StringPointerValue(resp.GatewayService.RuntimeGroupID)
		}
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}

func (r *APIProductVersionResourceModel) ToSharedUpdateAPIProductVersionDTO() *shared.UpdateAPIProductVersionDTO {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	deprecated := new(bool)
	if !r.Deprecated.IsUnknown() && !r.Deprecated.IsNull() {
		*deprecated = r.Deprecated.ValueBool()
	} else {
		deprecated = nil
	}
	var gatewayService *shared.GatewayServicePayload
	if r.GatewayService != nil {
		id := r.GatewayService.ID.ValueString()
		controlPlaneID := r.GatewayService.ControlPlaneID.ValueString()
		gatewayService = &shared.GatewayServicePayload{
			ID:             id,
			ControlPlaneID: controlPlaneID,
		}
	}
	out := shared.UpdateAPIProductVersionDTO{
		Name:           name,
		Deprecated:     deprecated,
		GatewayService: gatewayService,
	}
	return &out
}
