// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginJSONThreatProtectionResourceModel) ToSharedJSONThreatProtectionPluginInput() *shared.JSONThreatProtectionPluginInput {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.JSONThreatProtectionPluginOrdering
	if r.Ordering != nil {
		var after *shared.JSONThreatProtectionPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.JSONThreatProtectionPluginAfter{
				Access: access,
			}
		}
		var before *shared.JSONThreatProtectionPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.JSONThreatProtectionPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.JSONThreatProtectionPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	enforcementMode := new(shared.JSONThreatProtectionPluginEnforcementMode)
	if !r.Config.EnforcementMode.IsUnknown() && !r.Config.EnforcementMode.IsNull() {
		*enforcementMode = shared.JSONThreatProtectionPluginEnforcementMode(r.Config.EnforcementMode.ValueString())
	} else {
		enforcementMode = nil
	}
	errorMessage := new(string)
	if !r.Config.ErrorMessage.IsUnknown() && !r.Config.ErrorMessage.IsNull() {
		*errorMessage = r.Config.ErrorMessage.ValueString()
	} else {
		errorMessage = nil
	}
	errorStatusCode := new(int64)
	if !r.Config.ErrorStatusCode.IsUnknown() && !r.Config.ErrorStatusCode.IsNull() {
		*errorStatusCode = r.Config.ErrorStatusCode.ValueInt64()
	} else {
		errorStatusCode = nil
	}
	maxArrayElementCount := new(int64)
	if !r.Config.MaxArrayElementCount.IsUnknown() && !r.Config.MaxArrayElementCount.IsNull() {
		*maxArrayElementCount = r.Config.MaxArrayElementCount.ValueInt64()
	} else {
		maxArrayElementCount = nil
	}
	maxBodySize := new(int64)
	if !r.Config.MaxBodySize.IsUnknown() && !r.Config.MaxBodySize.IsNull() {
		*maxBodySize = r.Config.MaxBodySize.ValueInt64()
	} else {
		maxBodySize = nil
	}
	maxContainerDepth := new(int64)
	if !r.Config.MaxContainerDepth.IsUnknown() && !r.Config.MaxContainerDepth.IsNull() {
		*maxContainerDepth = r.Config.MaxContainerDepth.ValueInt64()
	} else {
		maxContainerDepth = nil
	}
	maxObjectEntryCount := new(int64)
	if !r.Config.MaxObjectEntryCount.IsUnknown() && !r.Config.MaxObjectEntryCount.IsNull() {
		*maxObjectEntryCount = r.Config.MaxObjectEntryCount.ValueInt64()
	} else {
		maxObjectEntryCount = nil
	}
	maxObjectEntryNameLength := new(int64)
	if !r.Config.MaxObjectEntryNameLength.IsUnknown() && !r.Config.MaxObjectEntryNameLength.IsNull() {
		*maxObjectEntryNameLength = r.Config.MaxObjectEntryNameLength.ValueInt64()
	} else {
		maxObjectEntryNameLength = nil
	}
	maxStringValueLength := new(int64)
	if !r.Config.MaxStringValueLength.IsUnknown() && !r.Config.MaxStringValueLength.IsNull() {
		*maxStringValueLength = r.Config.MaxStringValueLength.ValueInt64()
	} else {
		maxStringValueLength = nil
	}
	config := shared.JSONThreatProtectionPluginConfig{
		EnforcementMode:          enforcementMode,
		ErrorMessage:             errorMessage,
		ErrorStatusCode:          errorStatusCode,
		MaxArrayElementCount:     maxArrayElementCount,
		MaxBodySize:              maxBodySize,
		MaxContainerDepth:        maxContainerDepth,
		MaxObjectEntryCount:      maxObjectEntryCount,
		MaxObjectEntryNameLength: maxObjectEntryNameLength,
		MaxStringValueLength:     maxStringValueLength,
	}
	var protocols []shared.JSONThreatProtectionPluginProtocols = []shared.JSONThreatProtectionPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.JSONThreatProtectionPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.JSONThreatProtectionPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.JSONThreatProtectionPluginRoute{
			ID: id1,
		}
	}
	var service *shared.JSONThreatProtectionPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.JSONThreatProtectionPluginService{
			ID: id2,
		}
	}
	out := shared.JSONThreatProtectionPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginJSONThreatProtectionResourceModel) RefreshFromSharedJSONThreatProtectionPlugin(resp *shared.JSONThreatProtectionPlugin) {
	if resp != nil {
		if resp.Config.EnforcementMode != nil {
			r.Config.EnforcementMode = types.StringValue(string(*resp.Config.EnforcementMode))
		} else {
			r.Config.EnforcementMode = types.StringNull()
		}
		r.Config.ErrorMessage = types.StringPointerValue(resp.Config.ErrorMessage)
		r.Config.ErrorStatusCode = types.Int64PointerValue(resp.Config.ErrorStatusCode)
		r.Config.MaxArrayElementCount = types.Int64PointerValue(resp.Config.MaxArrayElementCount)
		r.Config.MaxBodySize = types.Int64PointerValue(resp.Config.MaxBodySize)
		r.Config.MaxContainerDepth = types.Int64PointerValue(resp.Config.MaxContainerDepth)
		r.Config.MaxObjectEntryCount = types.Int64PointerValue(resp.Config.MaxObjectEntryCount)
		r.Config.MaxObjectEntryNameLength = types.Int64PointerValue(resp.Config.MaxObjectEntryNameLength)
		r.Config.MaxStringValueLength = types.Int64PointerValue(resp.Config.MaxStringValueLength)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.ACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.ACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = []types.String{}
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
