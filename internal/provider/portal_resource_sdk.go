// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *PortalResourceModel) ToSharedCreatePortalRequest() *shared.CreatePortalRequest {
	name := r.Name.ValueString()
	displayName := new(string)
	if !r.DisplayName.IsUnknown() && !r.DisplayName.IsNull() {
		*displayName = r.DisplayName.ValueString()
	} else {
		displayName = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	isPublic := new(bool)
	if !r.IsPublic.IsUnknown() && !r.IsPublic.IsNull() {
		*isPublic = r.IsPublic.ValueBool()
	} else {
		isPublic = nil
	}
	rbacEnabled := new(bool)
	if !r.RbacEnabled.IsUnknown() && !r.RbacEnabled.IsNull() {
		*rbacEnabled = r.RbacEnabled.ValueBool()
	} else {
		rbacEnabled = nil
	}
	autoApproveApplications := new(bool)
	if !r.AutoApproveApplications.IsUnknown() && !r.AutoApproveApplications.IsNull() {
		*autoApproveApplications = r.AutoApproveApplications.ValueBool()
	} else {
		autoApproveApplications = nil
	}
	autoApproveDevelopers := new(bool)
	if !r.AutoApproveDevelopers.IsUnknown() && !r.AutoApproveDevelopers.IsNull() {
		*autoApproveDevelopers = r.AutoApproveDevelopers.ValueBool()
	} else {
		autoApproveDevelopers = nil
	}
	customDomain := new(string)
	if !r.CustomDomain.IsUnknown() && !r.CustomDomain.IsNull() {
		*customDomain = r.CustomDomain.ValueString()
	} else {
		customDomain = nil
	}
	customClientDomain := new(string)
	if !r.CustomClientDomain.IsUnknown() && !r.CustomClientDomain.IsNull() {
		*customClientDomain = r.CustomClientDomain.ValueString()
	} else {
		customClientDomain = nil
	}
	defaultApplicationAuthStrategyID := new(string)
	if !r.DefaultApplicationAuthStrategyID.IsUnknown() && !r.DefaultApplicationAuthStrategyID.IsNull() {
		*defaultApplicationAuthStrategyID = r.DefaultApplicationAuthStrategyID.ValueString()
	} else {
		defaultApplicationAuthStrategyID = nil
	}
	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		labelsInst := labelsValue.ValueString()
		labels[labelsKey] = labelsInst
	}
	out := shared.CreatePortalRequest{
		Name:                             name,
		DisplayName:                      displayName,
		Description:                      description,
		IsPublic:                         isPublic,
		RbacEnabled:                      rbacEnabled,
		AutoApproveApplications:          autoApproveApplications,
		AutoApproveDevelopers:            autoApproveDevelopers,
		CustomDomain:                     customDomain,
		CustomClientDomain:               customClientDomain,
		DefaultApplicationAuthStrategyID: defaultApplicationAuthStrategyID,
		Labels:                           labels,
	}
	return &out
}

func (r *PortalResourceModel) RefreshFromSharedCreatePortalResponse(resp *shared.CreatePortalResponse) {
	if resp != nil {
		r.ApplicationCount = types.NumberValue(big.NewFloat(float64(resp.ApplicationCount)))
		r.AutoApproveApplications = types.BoolValue(resp.AutoApproveApplications)
		r.AutoApproveDevelopers = types.BoolValue(resp.AutoApproveDevelopers)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.CustomClientDomain = types.StringPointerValue(resp.CustomClientDomain)
		r.CustomDomain = types.StringPointerValue(resp.CustomDomain)
		r.DefaultApplicationAuthStrategyID = types.StringPointerValue(resp.DefaultApplicationAuthStrategyID)
		r.DefaultDomain = types.StringValue(resp.DefaultDomain)
		r.Description = types.StringPointerValue(resp.Description)
		r.DeveloperCount = types.NumberValue(big.NewFloat(float64(resp.DeveloperCount)))
		r.DisplayName = types.StringValue(resp.DisplayName)
		r.ID = types.StringValue(resp.ID)
		r.IsPublic = types.BoolValue(resp.IsPublic)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.PublishedProductCount = types.NumberValue(big.NewFloat(float64(resp.PublishedProductCount)))
		r.RbacEnabled = types.BoolValue(resp.RbacEnabled)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}

func (r *PortalResourceModel) RefreshFromSharedGetPortalResponse(resp *shared.GetPortalResponse) {
	if resp != nil {
		r.ApplicationCount = types.NumberValue(big.NewFloat(float64(resp.ApplicationCount)))
		r.AutoApproveApplications = types.BoolValue(resp.AutoApproveApplications)
		r.AutoApproveDevelopers = types.BoolValue(resp.AutoApproveDevelopers)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.CustomClientDomain = types.StringPointerValue(resp.CustomClientDomain)
		r.CustomDomain = types.StringPointerValue(resp.CustomDomain)
		r.DefaultApplicationAuthStrategyID = types.StringPointerValue(resp.DefaultApplicationAuthStrategyID)
		r.DefaultDomain = types.StringValue(resp.DefaultDomain)
		r.Description = types.StringPointerValue(resp.Description)
		r.DeveloperCount = types.NumberValue(big.NewFloat(float64(resp.DeveloperCount)))
		r.DisplayName = types.StringValue(resp.DisplayName)
		r.ID = types.StringValue(resp.ID)
		r.IsPublic = types.BoolValue(resp.IsPublic)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.PublishedProductCount = types.NumberValue(big.NewFloat(float64(resp.PublishedProductCount)))
		r.RbacEnabled = types.BoolValue(resp.RbacEnabled)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}

func (r *PortalResourceModel) ToSharedUpdatePortalRequest() *shared.UpdatePortalRequest {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	displayName := new(string)
	if !r.DisplayName.IsUnknown() && !r.DisplayName.IsNull() {
		*displayName = r.DisplayName.ValueString()
	} else {
		displayName = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	isPublic := new(bool)
	if !r.IsPublic.IsUnknown() && !r.IsPublic.IsNull() {
		*isPublic = r.IsPublic.ValueBool()
	} else {
		isPublic = nil
	}
	rbacEnabled := new(bool)
	if !r.RbacEnabled.IsUnknown() && !r.RbacEnabled.IsNull() {
		*rbacEnabled = r.RbacEnabled.ValueBool()
	} else {
		rbacEnabled = nil
	}
	autoApproveApplications := new(bool)
	if !r.AutoApproveApplications.IsUnknown() && !r.AutoApproveApplications.IsNull() {
		*autoApproveApplications = r.AutoApproveApplications.ValueBool()
	} else {
		autoApproveApplications = nil
	}
	autoApproveDevelopers := new(bool)
	if !r.AutoApproveDevelopers.IsUnknown() && !r.AutoApproveDevelopers.IsNull() {
		*autoApproveDevelopers = r.AutoApproveDevelopers.ValueBool()
	} else {
		autoApproveDevelopers = nil
	}
	customDomain := new(string)
	if !r.CustomDomain.IsUnknown() && !r.CustomDomain.IsNull() {
		*customDomain = r.CustomDomain.ValueString()
	} else {
		customDomain = nil
	}
	customClientDomain := new(string)
	if !r.CustomClientDomain.IsUnknown() && !r.CustomClientDomain.IsNull() {
		*customClientDomain = r.CustomClientDomain.ValueString()
	} else {
		customClientDomain = nil
	}
	defaultApplicationAuthStrategyID := new(string)
	if !r.DefaultApplicationAuthStrategyID.IsUnknown() && !r.DefaultApplicationAuthStrategyID.IsNull() {
		*defaultApplicationAuthStrategyID = r.DefaultApplicationAuthStrategyID.ValueString()
	} else {
		defaultApplicationAuthStrategyID = nil
	}
	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		labelsInst := labelsValue.ValueString()
		labels[labelsKey] = labelsInst
	}
	out := shared.UpdatePortalRequest{
		Name:                             name,
		DisplayName:                      displayName,
		Description:                      description,
		IsPublic:                         isPublic,
		RbacEnabled:                      rbacEnabled,
		AutoApproveApplications:          autoApproveApplications,
		AutoApproveDevelopers:            autoApproveDevelopers,
		CustomDomain:                     customDomain,
		CustomClientDomain:               customClientDomain,
		DefaultApplicationAuthStrategyID: defaultApplicationAuthStrategyID,
		Labels:                           labels,
	}
	return &out
}

func (r *PortalResourceModel) RefreshFromSharedUpdatePortalResponse(resp *shared.UpdatePortalResponse) {
	if resp != nil {
		r.ApplicationCount = types.NumberValue(big.NewFloat(float64(resp.ApplicationCount)))
		r.AutoApproveApplications = types.BoolValue(resp.AutoApproveApplications)
		r.AutoApproveDevelopers = types.BoolValue(resp.AutoApproveDevelopers)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.CustomClientDomain = types.StringPointerValue(resp.CustomClientDomain)
		r.CustomDomain = types.StringPointerValue(resp.CustomDomain)
		r.DefaultApplicationAuthStrategyID = types.StringPointerValue(resp.DefaultApplicationAuthStrategyID)
		r.DefaultDomain = types.StringValue(resp.DefaultDomain)
		r.Description = types.StringPointerValue(resp.Description)
		r.DeveloperCount = types.NumberValue(big.NewFloat(float64(resp.DeveloperCount)))
		r.DisplayName = types.StringValue(resp.DisplayName)
		r.ID = types.StringValue(resp.ID)
		r.IsPublic = types.BoolValue(resp.IsPublic)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.PublishedProductCount = types.NumberValue(big.NewFloat(float64(resp.PublishedProductCount)))
		r.RbacEnabled = types.BoolValue(resp.RbacEnabled)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}
