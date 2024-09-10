// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginResponseTransformerAdvancedResourceModel) ToSharedCreateResponseTransformerAdvancedPlugin() *shared.CreateResponseTransformerAdvancedPlugin {
	var config *shared.CreateResponseTransformerAdvancedPluginConfig
	if r.Config != nil {
		var add *shared.CreateResponseTransformerAdvancedPluginAdd
		if r.Config.Add != nil {
			var headers []string = []string{}
			for _, headersItem := range r.Config.Add.Headers {
				headers = append(headers, headersItem.ValueString())
			}
			var ifStatus []string = []string{}
			for _, ifStatusItem := range r.Config.Add.IfStatus {
				ifStatus = append(ifStatus, ifStatusItem.ValueString())
			}
			var json []string = []string{}
			for _, jsonItem := range r.Config.Add.JSON {
				json = append(json, jsonItem.ValueString())
			}
			var jsonTypes []shared.CreateResponseTransformerAdvancedPluginJSONTypes = []shared.CreateResponseTransformerAdvancedPluginJSONTypes{}
			for _, jsonTypesItem := range r.Config.Add.JSONTypes {
				jsonTypes = append(jsonTypes, shared.CreateResponseTransformerAdvancedPluginJSONTypes(jsonTypesItem.ValueString()))
			}
			add = &shared.CreateResponseTransformerAdvancedPluginAdd{
				Headers:   headers,
				IfStatus:  ifStatus,
				JSON:      json,
				JSONTypes: jsonTypes,
			}
		}
		var allow *shared.CreateResponseTransformerAdvancedPluginAllow
		if r.Config.Allow != nil {
			var json1 []string = []string{}
			for _, jsonItem1 := range r.Config.Allow.JSON {
				json1 = append(json1, jsonItem1.ValueString())
			}
			allow = &shared.CreateResponseTransformerAdvancedPluginAllow{
				JSON: json1,
			}
		}
		var append1 *shared.CreateResponseTransformerAdvancedPluginAppend
		if r.Config.Append != nil {
			var headers1 []string = []string{}
			for _, headersItem1 := range r.Config.Append.Headers {
				headers1 = append(headers1, headersItem1.ValueString())
			}
			var ifStatus1 []string = []string{}
			for _, ifStatusItem1 := range r.Config.Append.IfStatus {
				ifStatus1 = append(ifStatus1, ifStatusItem1.ValueString())
			}
			var json2 []string = []string{}
			for _, jsonItem2 := range r.Config.Append.JSON {
				json2 = append(json2, jsonItem2.ValueString())
			}
			var jsonTypes1 []shared.CreateResponseTransformerAdvancedPluginConfigJSONTypes = []shared.CreateResponseTransformerAdvancedPluginConfigJSONTypes{}
			for _, jsonTypesItem1 := range r.Config.Append.JSONTypes {
				jsonTypes1 = append(jsonTypes1, shared.CreateResponseTransformerAdvancedPluginConfigJSONTypes(jsonTypesItem1.ValueString()))
			}
			append1 = &shared.CreateResponseTransformerAdvancedPluginAppend{
				Headers:   headers1,
				IfStatus:  ifStatus1,
				JSON:      json2,
				JSONTypes: jsonTypes1,
			}
		}
		dotsInKeys := new(bool)
		if !r.Config.DotsInKeys.IsUnknown() && !r.Config.DotsInKeys.IsNull() {
			*dotsInKeys = r.Config.DotsInKeys.ValueBool()
		} else {
			dotsInKeys = nil
		}
		var remove *shared.CreateResponseTransformerAdvancedPluginRemove
		if r.Config.Remove != nil {
			var headers2 []string = []string{}
			for _, headersItem2 := range r.Config.Remove.Headers {
				headers2 = append(headers2, headersItem2.ValueString())
			}
			var ifStatus2 []string = []string{}
			for _, ifStatusItem2 := range r.Config.Remove.IfStatus {
				ifStatus2 = append(ifStatus2, ifStatusItem2.ValueString())
			}
			var json3 []string = []string{}
			for _, jsonItem3 := range r.Config.Remove.JSON {
				json3 = append(json3, jsonItem3.ValueString())
			}
			remove = &shared.CreateResponseTransformerAdvancedPluginRemove{
				Headers:  headers2,
				IfStatus: ifStatus2,
				JSON:     json3,
			}
		}
		var rename *shared.CreateResponseTransformerAdvancedPluginRename
		if r.Config.Rename != nil {
			var headers3 []string = []string{}
			for _, headersItem3 := range r.Config.Rename.Headers {
				headers3 = append(headers3, headersItem3.ValueString())
			}
			var ifStatus3 []string = []string{}
			for _, ifStatusItem3 := range r.Config.Rename.IfStatus {
				ifStatus3 = append(ifStatus3, ifStatusItem3.ValueString())
			}
			rename = &shared.CreateResponseTransformerAdvancedPluginRename{
				Headers:  headers3,
				IfStatus: ifStatus3,
			}
		}
		var replace *shared.CreateResponseTransformerAdvancedPluginReplace
		if r.Config.Replace != nil {
			body := new(string)
			if !r.Config.Replace.Body.IsUnknown() && !r.Config.Replace.Body.IsNull() {
				*body = r.Config.Replace.Body.ValueString()
			} else {
				body = nil
			}
			var headers4 []string = []string{}
			for _, headersItem4 := range r.Config.Replace.Headers {
				headers4 = append(headers4, headersItem4.ValueString())
			}
			var ifStatus4 []string = []string{}
			for _, ifStatusItem4 := range r.Config.Replace.IfStatus {
				ifStatus4 = append(ifStatus4, ifStatusItem4.ValueString())
			}
			var json4 []string = []string{}
			for _, jsonItem4 := range r.Config.Replace.JSON {
				json4 = append(json4, jsonItem4.ValueString())
			}
			var jsonTypes2 []shared.CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes = []shared.CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes{}
			for _, jsonTypesItem2 := range r.Config.Replace.JSONTypes {
				jsonTypes2 = append(jsonTypes2, shared.CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes(jsonTypesItem2.ValueString()))
			}
			replace = &shared.CreateResponseTransformerAdvancedPluginReplace{
				Body:      body,
				Headers:   headers4,
				IfStatus:  ifStatus4,
				JSON:      json4,
				JSONTypes: jsonTypes2,
			}
		}
		var transform *shared.CreateResponseTransformerAdvancedPluginTransform
		if r.Config.Transform != nil {
			var functions []string = []string{}
			for _, functionsItem := range r.Config.Transform.Functions {
				functions = append(functions, functionsItem.ValueString())
			}
			var ifStatus5 []string = []string{}
			for _, ifStatusItem5 := range r.Config.Transform.IfStatus {
				ifStatus5 = append(ifStatus5, ifStatusItem5.ValueString())
			}
			var json5 []string = []string{}
			for _, jsonItem5 := range r.Config.Transform.JSON {
				json5 = append(json5, jsonItem5.ValueString())
			}
			transform = &shared.CreateResponseTransformerAdvancedPluginTransform{
				Functions: functions,
				IfStatus:  ifStatus5,
				JSON:      json5,
			}
		}
		config = &shared.CreateResponseTransformerAdvancedPluginConfig{
			Add:        add,
			Allow:      allow,
			Append:     append1,
			DotsInKeys: dotsInKeys,
			Remove:     remove,
			Rename:     rename,
			Replace:    replace,
			Transform:  transform,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.CreateResponseTransformerAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateResponseTransformerAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateResponseTransformerAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateResponseTransformerAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateResponseTransformerAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateResponseTransformerAdvancedPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateResponseTransformerAdvancedPluginProtocols = []shared.CreateResponseTransformerAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateResponseTransformerAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateResponseTransformerAdvancedPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateResponseTransformerAdvancedPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateResponseTransformerAdvancedPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateResponseTransformerAdvancedPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateResponseTransformerAdvancedPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateResponseTransformerAdvancedPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateResponseTransformerAdvancedPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateResponseTransformerAdvancedPluginService{
			ID: id3,
		}
	}
	out := shared.CreateResponseTransformerAdvancedPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginResponseTransformerAdvancedResourceModel) RefreshFromSharedResponseTransformerAdvancedPlugin(resp *shared.ResponseTransformerAdvancedPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateResponseTransformerAdvancedPluginConfig{}
			if resp.Config.Add == nil {
				r.Config.Add = nil
			} else {
				r.Config.Add = &tfTypes.CreateResponseTransformerAdvancedPluginAdd{}
				r.Config.Add.Headers = []types.String{}
				for _, v := range resp.Config.Add.Headers {
					r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
				}
				r.Config.Add.IfStatus = []types.String{}
				for _, v := range resp.Config.Add.IfStatus {
					r.Config.Add.IfStatus = append(r.Config.Add.IfStatus, types.StringValue(v))
				}
				r.Config.Add.JSON = []types.String{}
				for _, v := range resp.Config.Add.JSON {
					r.Config.Add.JSON = append(r.Config.Add.JSON, types.StringValue(v))
				}
				r.Config.Add.JSONTypes = []types.String{}
				for _, v := range resp.Config.Add.JSONTypes {
					r.Config.Add.JSONTypes = append(r.Config.Add.JSONTypes, types.StringValue(string(v)))
				}
			}
			if resp.Config.Allow == nil {
				r.Config.Allow = nil
			} else {
				r.Config.Allow = &tfTypes.CreateResponseTransformerAdvancedPluginAllow{}
				r.Config.Allow.JSON = []types.String{}
				for _, v := range resp.Config.Allow.JSON {
					r.Config.Allow.JSON = append(r.Config.Allow.JSON, types.StringValue(v))
				}
			}
			if resp.Config.Append == nil {
				r.Config.Append = nil
			} else {
				r.Config.Append = &tfTypes.CreateResponseTransformerAdvancedPluginAdd{}
				r.Config.Append.Headers = []types.String{}
				for _, v := range resp.Config.Append.Headers {
					r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
				}
				r.Config.Append.IfStatus = []types.String{}
				for _, v := range resp.Config.Append.IfStatus {
					r.Config.Append.IfStatus = append(r.Config.Append.IfStatus, types.StringValue(v))
				}
				r.Config.Append.JSON = []types.String{}
				for _, v := range resp.Config.Append.JSON {
					r.Config.Append.JSON = append(r.Config.Append.JSON, types.StringValue(v))
				}
				r.Config.Append.JSONTypes = []types.String{}
				for _, v := range resp.Config.Append.JSONTypes {
					r.Config.Append.JSONTypes = append(r.Config.Append.JSONTypes, types.StringValue(string(v)))
				}
			}
			r.Config.DotsInKeys = types.BoolPointerValue(resp.Config.DotsInKeys)
			if resp.Config.Remove == nil {
				r.Config.Remove = nil
			} else {
				r.Config.Remove = &tfTypes.CreateResponseTransformerAdvancedPluginRemove{}
				r.Config.Remove.Headers = []types.String{}
				for _, v := range resp.Config.Remove.Headers {
					r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
				}
				r.Config.Remove.IfStatus = []types.String{}
				for _, v := range resp.Config.Remove.IfStatus {
					r.Config.Remove.IfStatus = append(r.Config.Remove.IfStatus, types.StringValue(v))
				}
				r.Config.Remove.JSON = []types.String{}
				for _, v := range resp.Config.Remove.JSON {
					r.Config.Remove.JSON = append(r.Config.Remove.JSON, types.StringValue(v))
				}
			}
			if resp.Config.Rename == nil {
				r.Config.Rename = nil
			} else {
				r.Config.Rename = &tfTypes.CreateResponseTransformerAdvancedPluginRename{}
				r.Config.Rename.Headers = []types.String{}
				for _, v := range resp.Config.Rename.Headers {
					r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
				}
				r.Config.Rename.IfStatus = []types.String{}
				for _, v := range resp.Config.Rename.IfStatus {
					r.Config.Rename.IfStatus = append(r.Config.Rename.IfStatus, types.StringValue(v))
				}
			}
			if resp.Config.Replace == nil {
				r.Config.Replace = nil
			} else {
				r.Config.Replace = &tfTypes.CreateResponseTransformerAdvancedPluginReplace{}
				r.Config.Replace.Body = types.StringPointerValue(resp.Config.Replace.Body)
				r.Config.Replace.Headers = []types.String{}
				for _, v := range resp.Config.Replace.Headers {
					r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
				}
				r.Config.Replace.IfStatus = []types.String{}
				for _, v := range resp.Config.Replace.IfStatus {
					r.Config.Replace.IfStatus = append(r.Config.Replace.IfStatus, types.StringValue(v))
				}
				r.Config.Replace.JSON = []types.String{}
				for _, v := range resp.Config.Replace.JSON {
					r.Config.Replace.JSON = append(r.Config.Replace.JSON, types.StringValue(v))
				}
				r.Config.Replace.JSONTypes = []types.String{}
				for _, v := range resp.Config.Replace.JSONTypes {
					r.Config.Replace.JSONTypes = append(r.Config.Replace.JSONTypes, types.StringValue(string(v)))
				}
			}
			if resp.Config.Transform == nil {
				r.Config.Transform = nil
			} else {
				r.Config.Transform = &tfTypes.CreateResponseTransformerAdvancedPluginTransform{}
				r.Config.Transform.Functions = []types.String{}
				for _, v := range resp.Config.Transform.Functions {
					r.Config.Transform.Functions = append(r.Config.Transform.Functions, types.StringValue(v))
				}
				r.Config.Transform.IfStatus = []types.String{}
				for _, v := range resp.Config.Transform.IfStatus {
					r.Config.Transform.IfStatus = append(r.Config.Transform.IfStatus, types.StringValue(v))
				}
				r.Config.Transform.JSON = []types.String{}
				for _, v := range resp.Config.Transform.JSON {
					r.Config.Transform.JSON = append(r.Config.Transform.JSON, types.StringValue(v))
				}
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.CreateACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.CreateACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.CreateACLPluginAfter{}
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
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
