// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginRequestTransformerResourceModel) ToSharedCreateRequestTransformerPlugin() *shared.CreateRequestTransformerPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateRequestTransformerPluginProtocols = []shared.CreateRequestTransformerPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateRequestTransformerPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateRequestTransformerPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateRequestTransformerPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateRequestTransformerPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateRequestTransformerPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateRequestTransformerPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateRequestTransformerPluginService{
			ID: id2,
		}
	}
	httpMethod := new(string)
	if !r.Config.HTTPMethod.IsUnknown() && !r.Config.HTTPMethod.IsNull() {
		*httpMethod = r.Config.HTTPMethod.ValueString()
	} else {
		httpMethod = nil
	}
	var remove *shared.Remove
	if r.Config.Remove != nil {
		var body []string = []string{}
		for _, bodyItem := range r.Config.Remove.Body {
			body = append(body, bodyItem.ValueString())
		}
		var headers []string = []string{}
		for _, headersItem := range r.Config.Remove.Headers {
			headers = append(headers, headersItem.ValueString())
		}
		var querystring []string = []string{}
		for _, querystringItem := range r.Config.Remove.Querystring {
			querystring = append(querystring, querystringItem.ValueString())
		}
		remove = &shared.Remove{
			Body:        body,
			Headers:     headers,
			Querystring: querystring,
		}
	}
	var rename *shared.Rename
	if r.Config.Rename != nil {
		var body1 []string = []string{}
		for _, bodyItem1 := range r.Config.Rename.Body {
			body1 = append(body1, bodyItem1.ValueString())
		}
		var headers1 []string = []string{}
		for _, headersItem1 := range r.Config.Rename.Headers {
			headers1 = append(headers1, headersItem1.ValueString())
		}
		var querystring1 []string = []string{}
		for _, querystringItem1 := range r.Config.Rename.Querystring {
			querystring1 = append(querystring1, querystringItem1.ValueString())
		}
		rename = &shared.Rename{
			Body:        body1,
			Headers:     headers1,
			Querystring: querystring1,
		}
	}
	var replace *shared.Replace
	if r.Config.Replace != nil {
		var body2 []string = []string{}
		for _, bodyItem2 := range r.Config.Replace.Body {
			body2 = append(body2, bodyItem2.ValueString())
		}
		var headers2 []string = []string{}
		for _, headersItem2 := range r.Config.Replace.Headers {
			headers2 = append(headers2, headersItem2.ValueString())
		}
		var querystring2 []string = []string{}
		for _, querystringItem2 := range r.Config.Replace.Querystring {
			querystring2 = append(querystring2, querystringItem2.ValueString())
		}
		uri := new(string)
		if !r.Config.Replace.URI.IsUnknown() && !r.Config.Replace.URI.IsNull() {
			*uri = r.Config.Replace.URI.ValueString()
		} else {
			uri = nil
		}
		replace = &shared.Replace{
			Body:        body2,
			Headers:     headers2,
			Querystring: querystring2,
			URI:         uri,
		}
	}
	var add *shared.Add
	if r.Config.Add != nil {
		var body3 []string = []string{}
		for _, bodyItem3 := range r.Config.Add.Body {
			body3 = append(body3, bodyItem3.ValueString())
		}
		var headers3 []string = []string{}
		for _, headersItem3 := range r.Config.Add.Headers {
			headers3 = append(headers3, headersItem3.ValueString())
		}
		var querystring3 []string = []string{}
		for _, querystringItem3 := range r.Config.Add.Querystring {
			querystring3 = append(querystring3, querystringItem3.ValueString())
		}
		add = &shared.Add{
			Body:        body3,
			Headers:     headers3,
			Querystring: querystring3,
		}
	}
	var append1 *shared.Append
	if r.Config.Append != nil {
		var body4 []string = []string{}
		for _, bodyItem4 := range r.Config.Append.Body {
			body4 = append(body4, bodyItem4.ValueString())
		}
		var headers4 []string = []string{}
		for _, headersItem4 := range r.Config.Append.Headers {
			headers4 = append(headers4, headersItem4.ValueString())
		}
		var querystring4 []string = []string{}
		for _, querystringItem4 := range r.Config.Append.Querystring {
			querystring4 = append(querystring4, querystringItem4.ValueString())
		}
		append1 = &shared.Append{
			Body:        body4,
			Headers:     headers4,
			Querystring: querystring4,
		}
	}
	config := shared.CreateRequestTransformerPluginConfig{
		HTTPMethod: httpMethod,
		Remove:     remove,
		Rename:     rename,
		Replace:    replace,
		Add:        add,
		Append:     append1,
	}
	out := shared.CreateRequestTransformerPlugin{
		Enabled:   enabled,
		Protocols: protocols,
		Tags:      tags,
		Consumer:  consumer,
		Route:     route,
		Service:   service,
		Config:    config,
	}
	return &out
}

func (r *GatewayPluginRequestTransformerResourceModel) RefreshFromSharedRequestTransformerPlugin(resp *shared.RequestTransformerPlugin) {
	if resp != nil {
		if resp.Config.Add == nil {
			r.Config.Add = nil
		} else {
			r.Config.Add = &tfTypes.Add{}
			r.Config.Add.Body = []types.String{}
			for _, v := range resp.Config.Add.Body {
				r.Config.Add.Body = append(r.Config.Add.Body, types.StringValue(v))
			}
			r.Config.Add.Headers = []types.String{}
			for _, v := range resp.Config.Add.Headers {
				r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
			}
			r.Config.Add.Querystring = []types.String{}
			for _, v := range resp.Config.Add.Querystring {
				r.Config.Add.Querystring = append(r.Config.Add.Querystring, types.StringValue(v))
			}
		}
		if resp.Config.Append == nil {
			r.Config.Append = nil
		} else {
			r.Config.Append = &tfTypes.Add{}
			r.Config.Append.Body = []types.String{}
			for _, v := range resp.Config.Append.Body {
				r.Config.Append.Body = append(r.Config.Append.Body, types.StringValue(v))
			}
			r.Config.Append.Headers = []types.String{}
			for _, v := range resp.Config.Append.Headers {
				r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
			}
			r.Config.Append.Querystring = []types.String{}
			for _, v := range resp.Config.Append.Querystring {
				r.Config.Append.Querystring = append(r.Config.Append.Querystring, types.StringValue(v))
			}
		}
		r.Config.HTTPMethod = types.StringPointerValue(resp.Config.HTTPMethod)
		if resp.Config.Remove == nil {
			r.Config.Remove = nil
		} else {
			r.Config.Remove = &tfTypes.Add{}
			r.Config.Remove.Body = []types.String{}
			for _, v := range resp.Config.Remove.Body {
				r.Config.Remove.Body = append(r.Config.Remove.Body, types.StringValue(v))
			}
			r.Config.Remove.Headers = []types.String{}
			for _, v := range resp.Config.Remove.Headers {
				r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
			}
			r.Config.Remove.Querystring = []types.String{}
			for _, v := range resp.Config.Remove.Querystring {
				r.Config.Remove.Querystring = append(r.Config.Remove.Querystring, types.StringValue(v))
			}
		}
		if resp.Config.Rename == nil {
			r.Config.Rename = nil
		} else {
			r.Config.Rename = &tfTypes.Add{}
			r.Config.Rename.Body = []types.String{}
			for _, v := range resp.Config.Rename.Body {
				r.Config.Rename.Body = append(r.Config.Rename.Body, types.StringValue(v))
			}
			r.Config.Rename.Headers = []types.String{}
			for _, v := range resp.Config.Rename.Headers {
				r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
			}
			r.Config.Rename.Querystring = []types.String{}
			for _, v := range resp.Config.Rename.Querystring {
				r.Config.Rename.Querystring = append(r.Config.Rename.Querystring, types.StringValue(v))
			}
		}
		if resp.Config.Replace == nil {
			r.Config.Replace = nil
		} else {
			r.Config.Replace = &tfTypes.Replace{}
			r.Config.Replace.Body = []types.String{}
			for _, v := range resp.Config.Replace.Body {
				r.Config.Replace.Body = append(r.Config.Replace.Body, types.StringValue(v))
			}
			r.Config.Replace.Headers = []types.String{}
			for _, v := range resp.Config.Replace.Headers {
				r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
			}
			r.Config.Replace.Querystring = []types.String{}
			for _, v := range resp.Config.Replace.Querystring {
				r.Config.Replace.Querystring = append(r.Config.Replace.Querystring, types.StringValue(v))
			}
			r.Config.Replace.URI = types.StringPointerValue(resp.Config.Replace.URI)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
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
	}
}
