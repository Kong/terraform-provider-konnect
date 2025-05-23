// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/provider/typeconvert"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *APIProductVersionDataSourceModel) ToOperationsGetAPIProductVersionRequest(ctx context.Context) (*operations.GetAPIProductVersionRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var apiProductID string
	apiProductID = r.APIProductID.ValueString()

	var id string
	id = r.ID.ValueString()

	out := operations.GetAPIProductVersionRequest{
		APIProductID: apiProductID,
		ID:           id,
	}

	return &out, diags
}

func (r *APIProductVersionDataSourceModel) RefreshFromSharedAPIProductVersion(ctx context.Context, resp *shared.APIProductVersion) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.CreatedAt))
		r.Deprecated = types.BoolValue(resp.Deprecated)
		if resp.GatewayService == nil {
			r.GatewayService = nil
		} else {
			r.GatewayService = &tfTypes.GatewayService{}
			r.GatewayService.ControlPlaneID = types.StringValue(resp.GatewayService.ControlPlaneID)
			r.GatewayService.ID = types.StringPointerValue(resp.GatewayService.ID)
			r.GatewayService.RuntimeGroupID = types.StringPointerValue(resp.GatewayService.RuntimeGroupID)
		}
		r.ID = types.StringValue(resp.ID)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringPointerValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.Portals = []tfTypes.APIProductVersionPortal{}
		if len(r.Portals) > len(resp.Portals) {
			r.Portals = r.Portals[:len(resp.Portals)]
		}
		for portalsCount, portalsItem := range resp.Portals {
			var portals tfTypes.APIProductVersionPortal
			portals.ApplicationRegistrationEnabled = types.BoolValue(portalsItem.ApplicationRegistrationEnabled)
			portals.AuthStrategies = []tfTypes.APIProductVersionAuthStrategy{}
			for authStrategiesCount, authStrategiesItem := range portalsItem.AuthStrategies {
				var authStrategies tfTypes.APIProductVersionAuthStrategy
				authStrategies.ID = types.StringValue(authStrategiesItem.ID)
				authStrategies.Name = types.StringValue(authStrategiesItem.Name)
				if authStrategiesCount+1 > len(portals.AuthStrategies) {
					portals.AuthStrategies = append(portals.AuthStrategies, authStrategies)
				} else {
					portals.AuthStrategies[authStrategiesCount].ID = authStrategies.ID
					portals.AuthStrategies[authStrategiesCount].Name = authStrategies.Name
				}
			}
			portals.AutoApproveRegistration = types.BoolValue(portalsItem.AutoApproveRegistration)
			portals.Deprecated = types.BoolValue(portalsItem.Deprecated)
			portals.PortalID = types.StringValue(portalsItem.PortalID)
			portals.PortalName = types.StringValue(portalsItem.PortalName)
			portals.PortalProductVersionID = types.StringValue(portalsItem.PortalProductVersionID)
			portals.PublishStatus = types.StringValue(string(portalsItem.PublishStatus))
			if portalsCount+1 > len(r.Portals) {
				r.Portals = append(r.Portals, portals)
			} else {
				r.Portals[portalsCount].ApplicationRegistrationEnabled = portals.ApplicationRegistrationEnabled
				r.Portals[portalsCount].AuthStrategies = portals.AuthStrategies
				r.Portals[portalsCount].AutoApproveRegistration = portals.AutoApproveRegistration
				r.Portals[portalsCount].Deprecated = portals.Deprecated
				r.Portals[portalsCount].PortalID = portals.PortalID
				r.Portals[portalsCount].PortalName = portals.PortalName
				r.Portals[portalsCount].PortalProductVersionID = portals.PortalProductVersionID
				r.Portals[portalsCount].PublishStatus = portals.PublishStatus
			}
		}
		r.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.UpdatedAt))
	}

	return diags
}
