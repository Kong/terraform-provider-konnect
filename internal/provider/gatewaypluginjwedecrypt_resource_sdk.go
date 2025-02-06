// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginJweDecryptResourceModel) ToSharedJweDecryptPluginInput() *shared.JweDecryptPluginInput {
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
	var ordering *shared.JweDecryptPluginOrdering
	if r.Ordering != nil {
		var after *shared.JweDecryptPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.JweDecryptPluginAfter{
				Access: access,
			}
		}
		var before *shared.JweDecryptPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.JweDecryptPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.JweDecryptPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	forwardHeaderName := new(string)
	if !r.Config.ForwardHeaderName.IsUnknown() && !r.Config.ForwardHeaderName.IsNull() {
		*forwardHeaderName = r.Config.ForwardHeaderName.ValueString()
	} else {
		forwardHeaderName = nil
	}
	var keySets []string = []string{}
	for _, keySetsItem := range r.Config.KeySets {
		keySets = append(keySets, keySetsItem.ValueString())
	}
	lookupHeaderName := new(string)
	if !r.Config.LookupHeaderName.IsUnknown() && !r.Config.LookupHeaderName.IsNull() {
		*lookupHeaderName = r.Config.LookupHeaderName.ValueString()
	} else {
		lookupHeaderName = nil
	}
	strict := new(bool)
	if !r.Config.Strict.IsUnknown() && !r.Config.Strict.IsNull() {
		*strict = r.Config.Strict.ValueBool()
	} else {
		strict = nil
	}
	config := shared.JweDecryptPluginConfig{
		ForwardHeaderName: forwardHeaderName,
		KeySets:           keySets,
		LookupHeaderName:  lookupHeaderName,
		Strict:            strict,
	}
	var protocols []shared.JweDecryptPluginProtocols = []shared.JweDecryptPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.JweDecryptPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.JweDecryptPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.JweDecryptPluginRoute{
			ID: id1,
		}
	}
	var service *shared.JweDecryptPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.JweDecryptPluginService{
			ID: id2,
		}
	}
	out := shared.JweDecryptPluginInput{
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

func (r *GatewayPluginJweDecryptResourceModel) RefreshFromSharedJweDecryptPlugin(resp *shared.JweDecryptPlugin) {
	if resp != nil {
		r.Config.ForwardHeaderName = types.StringPointerValue(resp.Config.ForwardHeaderName)
		r.Config.KeySets = []types.String{}
		for _, v := range resp.Config.KeySets {
			r.Config.KeySets = append(r.Config.KeySets, types.StringValue(v))
		}
		r.Config.LookupHeaderName = types.StringPointerValue(resp.Config.LookupHeaderName)
		r.Config.Strict = types.BoolPointerValue(resp.Config.Strict)
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
