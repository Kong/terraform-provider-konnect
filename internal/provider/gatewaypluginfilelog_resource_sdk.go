// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginFileLogResourceModel) ToSharedFileLogPluginInput() *shared.FileLogPluginInput {
	customFieldsByLua := make(map[string]interface{})
	for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
		var customFieldsByLuaInst interface{}
		_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
		customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
	}
	path := new(string)
	if !r.Config.Path.IsUnknown() && !r.Config.Path.IsNull() {
		*path = r.Config.Path.ValueString()
	} else {
		path = nil
	}
	reopen := new(bool)
	if !r.Config.Reopen.IsUnknown() && !r.Config.Reopen.IsNull() {
		*reopen = r.Config.Reopen.ValueBool()
	} else {
		reopen = nil
	}
	config := shared.FileLogPluginConfig{
		CustomFieldsByLua: customFieldsByLua,
		Path:              path,
		Reopen:            reopen,
	}
	var consumer *shared.FileLogPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.FileLogPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.FileLogPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.FileLogPluginConsumerGroup{
			ID: id1,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id2 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id2 = r.ID.ValueString()
	} else {
		id2 = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.FileLogPluginOrdering
	if r.Ordering != nil {
		var after *shared.FileLogPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.FileLogPluginAfter{
				Access: access,
			}
		}
		var before *shared.FileLogPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.FileLogPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.FileLogPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.FileLogPluginProtocols = []shared.FileLogPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.FileLogPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.FileLogPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.FileLogPluginRoute{
			ID: id3,
		}
	}
	var service *shared.FileLogPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.FileLogPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.FileLogPluginInput{
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Enabled:       enabled,
		ID:            id2,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
		Tags:          tags,
	}
	return &out
}

func (r *GatewayPluginFileLogResourceModel) RefreshFromSharedFileLogPlugin(resp *shared.FileLogPlugin) {
	if resp != nil {
		if len(resp.Config.CustomFieldsByLua) > 0 {
			r.Config.CustomFieldsByLua = make(map[string]types.String)
			for key, value := range resp.Config.CustomFieldsByLua {
				result, _ := json.Marshal(value)
				r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
			}
		}
		r.Config.Path = types.StringPointerValue(resp.Config.Path)
		r.Config.Reopen = types.BoolPointerValue(resp.Config.Reopen)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLWithoutParentsConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
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
