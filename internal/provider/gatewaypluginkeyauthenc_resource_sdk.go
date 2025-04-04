// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginKeyAuthEncResourceModel) ToSharedKeyAuthEncPlugin() *shared.KeyAuthEncPlugin {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
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
	var ordering *shared.KeyAuthEncPluginOrdering
	if r.Ordering != nil {
		var after *shared.KeyAuthEncPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.KeyAuthEncPluginAfter{
				Access: access,
			}
		}
		var before *shared.KeyAuthEncPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.KeyAuthEncPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.KeyAuthEncPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.KeyAuthEncPluginConfig
	if r.Config != nil {
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		hideCredentials := new(bool)
		if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
			*hideCredentials = r.Config.HideCredentials.ValueBool()
		} else {
			hideCredentials = nil
		}
		keyInBody := new(bool)
		if !r.Config.KeyInBody.IsUnknown() && !r.Config.KeyInBody.IsNull() {
			*keyInBody = r.Config.KeyInBody.ValueBool()
		} else {
			keyInBody = nil
		}
		keyInHeader := new(bool)
		if !r.Config.KeyInHeader.IsUnknown() && !r.Config.KeyInHeader.IsNull() {
			*keyInHeader = r.Config.KeyInHeader.ValueBool()
		} else {
			keyInHeader = nil
		}
		keyInQuery := new(bool)
		if !r.Config.KeyInQuery.IsUnknown() && !r.Config.KeyInQuery.IsNull() {
			*keyInQuery = r.Config.KeyInQuery.ValueBool()
		} else {
			keyInQuery = nil
		}
		var keyNames []string = []string{}
		for _, keyNamesItem := range r.Config.KeyNames {
			keyNames = append(keyNames, keyNamesItem.ValueString())
		}
		realm := new(string)
		if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
			*realm = r.Config.Realm.ValueString()
		} else {
			realm = nil
		}
		runOnPreflight := new(bool)
		if !r.Config.RunOnPreflight.IsUnknown() && !r.Config.RunOnPreflight.IsNull() {
			*runOnPreflight = r.Config.RunOnPreflight.ValueBool()
		} else {
			runOnPreflight = nil
		}
		config = &shared.KeyAuthEncPluginConfig{
			Anonymous:       anonymous,
			HideCredentials: hideCredentials,
			KeyInBody:       keyInBody,
			KeyInHeader:     keyInHeader,
			KeyInQuery:      keyInQuery,
			KeyNames:        keyNames,
			Realm:           realm,
			RunOnPreflight:  runOnPreflight,
		}
	}
	var protocols []shared.KeyAuthEncPluginProtocols = []shared.KeyAuthEncPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.KeyAuthEncPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.KeyAuthEncPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.KeyAuthEncPluginRoute{
			ID: id1,
		}
	}
	var service *shared.KeyAuthEncPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.KeyAuthEncPluginService{
			ID: id2,
		}
	}
	out := shared.KeyAuthEncPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginKeyAuthEncResourceModel) RefreshFromSharedKeyAuthEncPlugin(resp *shared.KeyAuthEncPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.KeyAuthPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
			r.Config.KeyInBody = types.BoolPointerValue(resp.Config.KeyInBody)
			r.Config.KeyInHeader = types.BoolPointerValue(resp.Config.KeyInHeader)
			r.Config.KeyInQuery = types.BoolPointerValue(resp.Config.KeyInQuery)
			r.Config.KeyNames = make([]types.String, 0, len(resp.Config.KeyNames))
			for _, v := range resp.Config.KeyNames {
				r.Config.KeyNames = append(r.Config.KeyNames, types.StringValue(v))
			}
			r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
			r.Config.RunOnPreflight = types.BoolPointerValue(resp.Config.RunOnPreflight)
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
				r.Ordering.After.Access = make([]types.String, 0, len(resp.Ordering.After.Access))
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = make([]types.String, 0, len(resp.Ordering.Before.Access))
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
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
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
