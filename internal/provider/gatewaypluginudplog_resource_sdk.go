// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginUDPLogResourceModel) ToSharedUDPLogPlugin() *shared.UDPLogPlugin {
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
	var ordering *shared.UDPLogPluginOrdering
	if r.Ordering != nil {
		var after *shared.UDPLogPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.UDPLogPluginAfter{
				Access: access,
			}
		}
		var before *shared.UDPLogPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.UDPLogPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.UDPLogPluginOrdering{
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
	var config *shared.UDPLogPluginConfig
	if r.Config != nil {
		customFieldsByLua := make(map[string]interface{})
		for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
			var customFieldsByLuaInst interface{}
			_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
			customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
		}
		host := new(string)
		if !r.Config.Host.IsUnknown() && !r.Config.Host.IsNull() {
			*host = r.Config.Host.ValueString()
		} else {
			host = nil
		}
		port := new(int64)
		if !r.Config.Port.IsUnknown() && !r.Config.Port.IsNull() {
			*port = r.Config.Port.ValueInt64()
		} else {
			port = nil
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
		} else {
			timeout = nil
		}
		config = &shared.UDPLogPluginConfig{
			CustomFieldsByLua: customFieldsByLua,
			Host:              host,
			Port:              port,
			Timeout:           timeout,
		}
	}
	var consumer *shared.UDPLogPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.UDPLogPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.UDPLogPluginProtocols = []shared.UDPLogPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.UDPLogPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.UDPLogPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.UDPLogPluginRoute{
			ID: id2,
		}
	}
	var service *shared.UDPLogPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.UDPLogPluginService{
			ID: id3,
		}
	}
	out := shared.UDPLogPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginUDPLogResourceModel) RefreshFromSharedUDPLogPlugin(resp *shared.UDPLogPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.UDPLogPluginConfig{}
			if len(resp.Config.CustomFieldsByLua) > 0 {
				r.Config.CustomFieldsByLua = make(map[string]types.String, len(resp.Config.CustomFieldsByLua))
				for key, value := range resp.Config.CustomFieldsByLua {
					result, _ := json.Marshal(value)
					r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
				}
			}
			r.Config.Host = types.StringPointerValue(resp.Config.Host)
			r.Config.Port = types.Int64PointerValue(resp.Config.Port)
			if resp.Config.Timeout != nil {
				r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
			} else {
				r.Config.Timeout = types.NumberNull()
			}
		}
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
