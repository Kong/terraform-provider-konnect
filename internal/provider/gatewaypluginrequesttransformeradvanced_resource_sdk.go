// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginRequestTransformerAdvancedResourceModel) ToSharedRequestTransformerAdvancedPlugin() *shared.RequestTransformerAdvancedPlugin {
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
	var ordering *shared.RequestTransformerAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.RequestTransformerAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.RequestTransformerAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.RequestTransformerAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.RequestTransformerAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.RequestTransformerAdvancedPluginOrdering{
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
	var config *shared.RequestTransformerAdvancedPluginConfig
	if r.Config != nil {
		var add *shared.RequestTransformerAdvancedPluginAdd
		if r.Config.Add != nil {
			var body []string = []string{}
			for _, bodyItem := range r.Config.Add.Body {
				body = append(body, bodyItem.ValueString())
			}
			var headers []string = []string{}
			for _, headersItem := range r.Config.Add.Headers {
				headers = append(headers, headersItem.ValueString())
			}
			var jsonTypes []shared.JSONTypes = []shared.JSONTypes{}
			for _, jsonTypesItem := range r.Config.Add.JSONTypes {
				jsonTypes = append(jsonTypes, shared.JSONTypes(jsonTypesItem.ValueString()))
			}
			var querystring []string = []string{}
			for _, querystringItem := range r.Config.Add.Querystring {
				querystring = append(querystring, querystringItem.ValueString())
			}
			add = &shared.RequestTransformerAdvancedPluginAdd{
				Body:        body,
				Headers:     headers,
				JSONTypes:   jsonTypes,
				Querystring: querystring,
			}
		}
		var allow *shared.Allow
		if r.Config.Allow != nil {
			var body1 []string = []string{}
			for _, bodyItem1 := range r.Config.Allow.Body {
				body1 = append(body1, bodyItem1.ValueString())
			}
			allow = &shared.Allow{
				Body: body1,
			}
		}
		var append1 *shared.RequestTransformerAdvancedPluginAppend
		if r.Config.Append != nil {
			var body2 []string = []string{}
			for _, bodyItem2 := range r.Config.Append.Body {
				body2 = append(body2, bodyItem2.ValueString())
			}
			var headers1 []string = []string{}
			for _, headersItem1 := range r.Config.Append.Headers {
				headers1 = append(headers1, headersItem1.ValueString())
			}
			var jsonTypes1 []shared.RequestTransformerAdvancedPluginJSONTypes = []shared.RequestTransformerAdvancedPluginJSONTypes{}
			for _, jsonTypesItem1 := range r.Config.Append.JSONTypes {
				jsonTypes1 = append(jsonTypes1, shared.RequestTransformerAdvancedPluginJSONTypes(jsonTypesItem1.ValueString()))
			}
			var querystring1 []string = []string{}
			for _, querystringItem1 := range r.Config.Append.Querystring {
				querystring1 = append(querystring1, querystringItem1.ValueString())
			}
			append1 = &shared.RequestTransformerAdvancedPluginAppend{
				Body:        body2,
				Headers:     headers1,
				JSONTypes:   jsonTypes1,
				Querystring: querystring1,
			}
		}
		dotsInKeys := new(bool)
		if !r.Config.DotsInKeys.IsUnknown() && !r.Config.DotsInKeys.IsNull() {
			*dotsInKeys = r.Config.DotsInKeys.ValueBool()
		} else {
			dotsInKeys = nil
		}
		httpMethod := new(string)
		if !r.Config.HTTPMethod.IsUnknown() && !r.Config.HTTPMethod.IsNull() {
			*httpMethod = r.Config.HTTPMethod.ValueString()
		} else {
			httpMethod = nil
		}
		var remove *shared.RequestTransformerAdvancedPluginRemove
		if r.Config.Remove != nil {
			var body3 []string = []string{}
			for _, bodyItem3 := range r.Config.Remove.Body {
				body3 = append(body3, bodyItem3.ValueString())
			}
			var headers2 []string = []string{}
			for _, headersItem2 := range r.Config.Remove.Headers {
				headers2 = append(headers2, headersItem2.ValueString())
			}
			var querystring2 []string = []string{}
			for _, querystringItem2 := range r.Config.Remove.Querystring {
				querystring2 = append(querystring2, querystringItem2.ValueString())
			}
			remove = &shared.RequestTransformerAdvancedPluginRemove{
				Body:        body3,
				Headers:     headers2,
				Querystring: querystring2,
			}
		}
		var rename *shared.RequestTransformerAdvancedPluginRename
		if r.Config.Rename != nil {
			var body4 []string = []string{}
			for _, bodyItem4 := range r.Config.Rename.Body {
				body4 = append(body4, bodyItem4.ValueString())
			}
			var headers3 []string = []string{}
			for _, headersItem3 := range r.Config.Rename.Headers {
				headers3 = append(headers3, headersItem3.ValueString())
			}
			var querystring3 []string = []string{}
			for _, querystringItem3 := range r.Config.Rename.Querystring {
				querystring3 = append(querystring3, querystringItem3.ValueString())
			}
			rename = &shared.RequestTransformerAdvancedPluginRename{
				Body:        body4,
				Headers:     headers3,
				Querystring: querystring3,
			}
		}
		var replace *shared.RequestTransformerAdvancedPluginReplace
		if r.Config.Replace != nil {
			var body5 []string = []string{}
			for _, bodyItem5 := range r.Config.Replace.Body {
				body5 = append(body5, bodyItem5.ValueString())
			}
			var headers4 []string = []string{}
			for _, headersItem4 := range r.Config.Replace.Headers {
				headers4 = append(headers4, headersItem4.ValueString())
			}
			var jsonTypes2 []shared.RequestTransformerAdvancedPluginConfigJSONTypes = []shared.RequestTransformerAdvancedPluginConfigJSONTypes{}
			for _, jsonTypesItem2 := range r.Config.Replace.JSONTypes {
				jsonTypes2 = append(jsonTypes2, shared.RequestTransformerAdvancedPluginConfigJSONTypes(jsonTypesItem2.ValueString()))
			}
			var querystring4 []string = []string{}
			for _, querystringItem4 := range r.Config.Replace.Querystring {
				querystring4 = append(querystring4, querystringItem4.ValueString())
			}
			uri := new(string)
			if !r.Config.Replace.URI.IsUnknown() && !r.Config.Replace.URI.IsNull() {
				*uri = r.Config.Replace.URI.ValueString()
			} else {
				uri = nil
			}
			replace = &shared.RequestTransformerAdvancedPluginReplace{
				Body:        body5,
				Headers:     headers4,
				JSONTypes:   jsonTypes2,
				Querystring: querystring4,
				URI:         uri,
			}
		}
		config = &shared.RequestTransformerAdvancedPluginConfig{
			Add:        add,
			Allow:      allow,
			Append:     append1,
			DotsInKeys: dotsInKeys,
			HTTPMethod: httpMethod,
			Remove:     remove,
			Rename:     rename,
			Replace:    replace,
		}
	}
	var consumer *shared.RequestTransformerAdvancedPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.RequestTransformerAdvancedPluginConsumer{
			ID: id1,
		}
	}
	var consumerGroup *shared.RequestTransformerAdvancedPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id2 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id2 = r.ConsumerGroup.ID.ValueString()
		} else {
			id2 = nil
		}
		consumerGroup = &shared.RequestTransformerAdvancedPluginConsumerGroup{
			ID: id2,
		}
	}
	var protocols []shared.RequestTransformerAdvancedPluginProtocols = []shared.RequestTransformerAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.RequestTransformerAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.RequestTransformerAdvancedPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.RequestTransformerAdvancedPluginRoute{
			ID: id3,
		}
	}
	var service *shared.RequestTransformerAdvancedPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.RequestTransformerAdvancedPluginService{
			ID: id4,
		}
	}
	out := shared.RequestTransformerAdvancedPlugin{
		CreatedAt:     createdAt,
		Enabled:       enabled,
		ID:            id,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Tags:          tags,
		UpdatedAt:     updatedAt,
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginRequestTransformerAdvancedResourceModel) RefreshFromSharedRequestTransformerAdvancedPlugin(resp *shared.RequestTransformerAdvancedPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.RequestTransformerAdvancedPluginConfig{}
			if resp.Config.Add == nil {
				r.Config.Add = nil
			} else {
				r.Config.Add = &tfTypes.RequestTransformerAdvancedPluginAdd{}
				r.Config.Add.Body = make([]types.String, 0, len(resp.Config.Add.Body))
				for _, v := range resp.Config.Add.Body {
					r.Config.Add.Body = append(r.Config.Add.Body, types.StringValue(v))
				}
				r.Config.Add.Headers = make([]types.String, 0, len(resp.Config.Add.Headers))
				for _, v := range resp.Config.Add.Headers {
					r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
				}
				r.Config.Add.JSONTypes = make([]types.String, 0, len(resp.Config.Add.JSONTypes))
				for _, v := range resp.Config.Add.JSONTypes {
					r.Config.Add.JSONTypes = append(r.Config.Add.JSONTypes, types.StringValue(string(v)))
				}
				r.Config.Add.Querystring = make([]types.String, 0, len(resp.Config.Add.Querystring))
				for _, v := range resp.Config.Add.Querystring {
					r.Config.Add.Querystring = append(r.Config.Add.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Allow == nil {
				r.Config.Allow = nil
			} else {
				r.Config.Allow = &tfTypes.Allow{}
				r.Config.Allow.Body = make([]types.String, 0, len(resp.Config.Allow.Body))
				for _, v := range resp.Config.Allow.Body {
					r.Config.Allow.Body = append(r.Config.Allow.Body, types.StringValue(v))
				}
			}
			if resp.Config.Append == nil {
				r.Config.Append = nil
			} else {
				r.Config.Append = &tfTypes.RequestTransformerAdvancedPluginAdd{}
				r.Config.Append.Body = make([]types.String, 0, len(resp.Config.Append.Body))
				for _, v := range resp.Config.Append.Body {
					r.Config.Append.Body = append(r.Config.Append.Body, types.StringValue(v))
				}
				r.Config.Append.Headers = make([]types.String, 0, len(resp.Config.Append.Headers))
				for _, v := range resp.Config.Append.Headers {
					r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
				}
				r.Config.Append.JSONTypes = make([]types.String, 0, len(resp.Config.Append.JSONTypes))
				for _, v := range resp.Config.Append.JSONTypes {
					r.Config.Append.JSONTypes = append(r.Config.Append.JSONTypes, types.StringValue(string(v)))
				}
				r.Config.Append.Querystring = make([]types.String, 0, len(resp.Config.Append.Querystring))
				for _, v := range resp.Config.Append.Querystring {
					r.Config.Append.Querystring = append(r.Config.Append.Querystring, types.StringValue(v))
				}
			}
			r.Config.DotsInKeys = types.BoolPointerValue(resp.Config.DotsInKeys)
			r.Config.HTTPMethod = types.StringPointerValue(resp.Config.HTTPMethod)
			if resp.Config.Remove == nil {
				r.Config.Remove = nil
			} else {
				r.Config.Remove = &tfTypes.Add{}
				r.Config.Remove.Body = make([]types.String, 0, len(resp.Config.Remove.Body))
				for _, v := range resp.Config.Remove.Body {
					r.Config.Remove.Body = append(r.Config.Remove.Body, types.StringValue(v))
				}
				r.Config.Remove.Headers = make([]types.String, 0, len(resp.Config.Remove.Headers))
				for _, v := range resp.Config.Remove.Headers {
					r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
				}
				r.Config.Remove.Querystring = make([]types.String, 0, len(resp.Config.Remove.Querystring))
				for _, v := range resp.Config.Remove.Querystring {
					r.Config.Remove.Querystring = append(r.Config.Remove.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Rename == nil {
				r.Config.Rename = nil
			} else {
				r.Config.Rename = &tfTypes.Add{}
				r.Config.Rename.Body = make([]types.String, 0, len(resp.Config.Rename.Body))
				for _, v := range resp.Config.Rename.Body {
					r.Config.Rename.Body = append(r.Config.Rename.Body, types.StringValue(v))
				}
				r.Config.Rename.Headers = make([]types.String, 0, len(resp.Config.Rename.Headers))
				for _, v := range resp.Config.Rename.Headers {
					r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
				}
				r.Config.Rename.Querystring = make([]types.String, 0, len(resp.Config.Rename.Querystring))
				for _, v := range resp.Config.Rename.Querystring {
					r.Config.Rename.Querystring = append(r.Config.Rename.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Replace == nil {
				r.Config.Replace = nil
			} else {
				r.Config.Replace = &tfTypes.RequestTransformerAdvancedPluginReplace{}
				r.Config.Replace.Body = make([]types.String, 0, len(resp.Config.Replace.Body))
				for _, v := range resp.Config.Replace.Body {
					r.Config.Replace.Body = append(r.Config.Replace.Body, types.StringValue(v))
				}
				r.Config.Replace.Headers = make([]types.String, 0, len(resp.Config.Replace.Headers))
				for _, v := range resp.Config.Replace.Headers {
					r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
				}
				r.Config.Replace.JSONTypes = make([]types.String, 0, len(resp.Config.Replace.JSONTypes))
				for _, v := range resp.Config.Replace.JSONTypes {
					r.Config.Replace.JSONTypes = append(r.Config.Replace.JSONTypes, types.StringValue(string(v)))
				}
				r.Config.Replace.Querystring = make([]types.String, 0, len(resp.Config.Replace.Querystring))
				for _, v := range resp.Config.Replace.Querystring {
					r.Config.Replace.Querystring = append(r.Config.Replace.Querystring, types.StringValue(v))
				}
				r.Config.Replace.URI = types.StringPointerValue(resp.Config.Replace.URI)
			}
		}
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
