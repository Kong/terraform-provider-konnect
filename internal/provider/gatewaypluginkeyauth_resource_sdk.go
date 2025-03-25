// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginKeyAuthResourceModel) ToSharedKeyAuthPluginInput() *shared.KeyAuthPluginInput {
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
	ordering := make(map[string]string)
	for orderingKey, orderingValue := range r.Ordering {
		var orderingInst string
		orderingInst = orderingValue.ValueString()

		ordering[orderingKey] = orderingInst
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
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
	config := shared.KeyAuthPluginConfig{
		Anonymous:       anonymous,
		HideCredentials: hideCredentials,
		KeyInBody:       keyInBody,
		KeyInHeader:     keyInHeader,
		KeyInQuery:      keyInQuery,
		KeyNames:        keyNames,
		Realm:           realm,
		RunOnPreflight:  runOnPreflight,
	}
	var protocols []shared.KeyAuthPluginProtocols = []shared.KeyAuthPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.KeyAuthPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.KeyAuthPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.KeyAuthPluginRoute{
			ID: id1,
		}
	}
	var service *shared.KeyAuthPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.KeyAuthPluginService{
			ID: id2,
		}
	}
	out := shared.KeyAuthPluginInput{
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

func (r *GatewayPluginKeyAuthResourceModel) RefreshFromSharedKeyAuthPlugin(resp *shared.KeyAuthPlugin) {
	if resp != nil {
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
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering != nil {
			r.Ordering = make(map[string]types.String, len(resp.Ordering))
			for key, value := range resp.Ordering {
				r.Ordering[key] = types.StringValue(value)
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
