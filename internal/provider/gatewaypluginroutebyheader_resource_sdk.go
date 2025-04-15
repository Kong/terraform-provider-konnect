// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginRouteByHeaderResourceModel) ToSharedRouteByHeaderPlugin() *shared.RouteByHeaderPlugin {
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
	var ordering *shared.RouteByHeaderPluginOrdering
	if r.Ordering != nil {
		var after *shared.RouteByHeaderPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.RouteByHeaderPluginAfter{
				Access: access,
			}
		}
		var before *shared.RouteByHeaderPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.RouteByHeaderPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.RouteByHeaderPluginOrdering{
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
	var config *shared.RouteByHeaderPluginConfig
	if r.Config != nil {
		var rules []shared.RouteByHeaderPluginRules = []shared.RouteByHeaderPluginRules{}
		for _, rulesItem := range r.Config.Rules {
			condition := make(map[string]interface{})
			for conditionKey, conditionValue := range rulesItem.Condition {
				var conditionInst interface{}
				_ = json.Unmarshal([]byte(conditionValue.ValueString()), &conditionInst)
				condition[conditionKey] = conditionInst
			}
			var upstreamName string
			upstreamName = rulesItem.UpstreamName.ValueString()

			rules = append(rules, shared.RouteByHeaderPluginRules{
				Condition:    condition,
				UpstreamName: upstreamName,
			})
		}
		config = &shared.RouteByHeaderPluginConfig{
			Rules: rules,
		}
	}
	var consumer *shared.RouteByHeaderPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.RouteByHeaderPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.RouteByHeaderPluginProtocols = []shared.RouteByHeaderPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.RouteByHeaderPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.RouteByHeaderPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.RouteByHeaderPluginRoute{
			ID: id2,
		}
	}
	var service *shared.RouteByHeaderPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.RouteByHeaderPluginService{
			ID: id3,
		}
	}
	out := shared.RouteByHeaderPlugin{
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

func (r *GatewayPluginRouteByHeaderResourceModel) RefreshFromSharedRouteByHeaderPlugin(ctx context.Context, resp *shared.RouteByHeaderPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.RouteByHeaderPluginConfig{}
			r.Config.Rules = []tfTypes.RouteByHeaderPluginRules{}
			if len(r.Config.Rules) > len(resp.Config.Rules) {
				r.Config.Rules = r.Config.Rules[:len(resp.Config.Rules)]
			}
			for rulesCount, rulesItem := range resp.Config.Rules {
				var rules tfTypes.RouteByHeaderPluginRules
				if len(rulesItem.Condition) > 0 {
					rules.Condition = make(map[string]types.String, len(rulesItem.Condition))
					for key, value := range rulesItem.Condition {
						result, _ := json.Marshal(value)
						rules.Condition[key] = types.StringValue(string(result))
					}
				}
				rules.UpstreamName = types.StringValue(rulesItem.UpstreamName)
				if rulesCount+1 > len(r.Config.Rules) {
					r.Config.Rules = append(r.Config.Rules, rules)
				} else {
					r.Config.Rules[rulesCount].Condition = rules.Condition
					r.Config.Rules[rulesCount].UpstreamName = rules.UpstreamName
				}
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

	return diags
}
