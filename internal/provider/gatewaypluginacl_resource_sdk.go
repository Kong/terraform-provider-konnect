// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginACLResourceModel) ToSharedACLPluginInput() *shared.ACLPluginInput {
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
	var ordering *shared.ACLPluginOrdering
	if r.Ordering != nil {
		var after *shared.ACLPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ACLPluginAfter{
				Access: access,
			}
		}
		var before *shared.ACLPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ACLPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ACLPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var allow []string = []string{}
	for _, allowItem := range r.Config.Allow {
		allow = append(allow, allowItem.ValueString())
	}
	alwaysUseAuthenticatedGroups := new(bool)
	if !r.Config.AlwaysUseAuthenticatedGroups.IsUnknown() && !r.Config.AlwaysUseAuthenticatedGroups.IsNull() {
		*alwaysUseAuthenticatedGroups = r.Config.AlwaysUseAuthenticatedGroups.ValueBool()
	} else {
		alwaysUseAuthenticatedGroups = nil
	}
	var deny []string = []string{}
	for _, denyItem := range r.Config.Deny {
		deny = append(deny, denyItem.ValueString())
	}
	hideGroupsHeader := new(bool)
	if !r.Config.HideGroupsHeader.IsUnknown() && !r.Config.HideGroupsHeader.IsNull() {
		*hideGroupsHeader = r.Config.HideGroupsHeader.ValueBool()
	} else {
		hideGroupsHeader = nil
	}
	includeConsumerGroups := new(bool)
	if !r.Config.IncludeConsumerGroups.IsUnknown() && !r.Config.IncludeConsumerGroups.IsNull() {
		*includeConsumerGroups = r.Config.IncludeConsumerGroups.ValueBool()
	} else {
		includeConsumerGroups = nil
	}
	config := shared.ACLPluginConfig{
		Allow:                        allow,
		AlwaysUseAuthenticatedGroups: alwaysUseAuthenticatedGroups,
		Deny:                         deny,
		HideGroupsHeader:             hideGroupsHeader,
		IncludeConsumerGroups:        includeConsumerGroups,
	}
	var protocols []shared.ACLPluginProtocols = []shared.ACLPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ACLPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ACLPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.ACLPluginRoute{
			ID: id1,
		}
	}
	var service *shared.ACLPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.ACLPluginService{
			ID: id2,
		}
	}
	out := shared.ACLPluginInput{
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

func (r *GatewayPluginACLResourceModel) RefreshFromSharedACLPlugin(resp *shared.ACLPlugin) {
	if resp != nil {
		r.Config.Allow = []types.String{}
		for _, v := range resp.Config.Allow {
			r.Config.Allow = append(r.Config.Allow, types.StringValue(v))
		}
		r.Config.AlwaysUseAuthenticatedGroups = types.BoolPointerValue(resp.Config.AlwaysUseAuthenticatedGroups)
		r.Config.Deny = []types.String{}
		for _, v := range resp.Config.Deny {
			r.Config.Deny = append(r.Config.Deny, types.StringValue(v))
		}
		r.Config.HideGroupsHeader = types.BoolPointerValue(resp.Config.HideGroupsHeader)
		r.Config.IncludeConsumerGroups = types.BoolPointerValue(resp.Config.IncludeConsumerGroups)
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
