// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginExitTransformerResourceModel) ToSharedExitTransformerPluginInput() *shared.ExitTransformerPluginInput {
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
	var ordering *shared.ExitTransformerPluginOrdering
	if r.Ordering != nil {
		var after *shared.ExitTransformerPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ExitTransformerPluginAfter{
				Access: access,
			}
		}
		var before *shared.ExitTransformerPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ExitTransformerPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ExitTransformerPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var functions []string = []string{}
	for _, functionsItem := range r.Config.Functions {
		functions = append(functions, functionsItem.ValueString())
	}
	handleUnexpected := new(bool)
	if !r.Config.HandleUnexpected.IsUnknown() && !r.Config.HandleUnexpected.IsNull() {
		*handleUnexpected = r.Config.HandleUnexpected.ValueBool()
	} else {
		handleUnexpected = nil
	}
	handleUnknown := new(bool)
	if !r.Config.HandleUnknown.IsUnknown() && !r.Config.HandleUnknown.IsNull() {
		*handleUnknown = r.Config.HandleUnknown.ValueBool()
	} else {
		handleUnknown = nil
	}
	config := shared.ExitTransformerPluginConfig{
		Functions:        functions,
		HandleUnexpected: handleUnexpected,
		HandleUnknown:    handleUnknown,
	}
	var consumer *shared.ExitTransformerPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.ExitTransformerPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.ExitTransformerPluginProtocols = []shared.ExitTransformerPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ExitTransformerPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ExitTransformerPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.ExitTransformerPluginRoute{
			ID: id2,
		}
	}
	var service *shared.ExitTransformerPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.ExitTransformerPluginService{
			ID: id3,
		}
	}
	out := shared.ExitTransformerPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginExitTransformerResourceModel) RefreshFromSharedExitTransformerPlugin(resp *shared.ExitTransformerPlugin) {
	if resp != nil {
		r.Config.Functions = []types.String{}
		for _, v := range resp.Config.Functions {
			r.Config.Functions = append(r.Config.Functions, types.StringValue(v))
		}
		r.Config.HandleUnexpected = types.BoolPointerValue(resp.Config.HandleUnexpected)
		r.Config.HandleUnknown = types.BoolPointerValue(resp.Config.HandleUnknown)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
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
