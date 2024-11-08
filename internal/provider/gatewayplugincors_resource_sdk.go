// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginCorsResourceModel) ToSharedCreateCorsPlugin() *shared.CreateCorsPlugin {
	var config *shared.CreateCorsPluginConfig
	if r.Config != nil {
		credentials := new(bool)
		if !r.Config.Credentials.IsUnknown() && !r.Config.Credentials.IsNull() {
			*credentials = r.Config.Credentials.ValueBool()
		} else {
			credentials = nil
		}
		var exposedHeaders []string = []string{}
		for _, exposedHeadersItem := range r.Config.ExposedHeaders {
			exposedHeaders = append(exposedHeaders, exposedHeadersItem.ValueString())
		}
		var headers []string = []string{}
		for _, headersItem := range r.Config.Headers {
			headers = append(headers, headersItem.ValueString())
		}
		maxAge := new(float64)
		if !r.Config.MaxAge.IsUnknown() && !r.Config.MaxAge.IsNull() {
			*maxAge, _ = r.Config.MaxAge.ValueBigFloat().Float64()
		} else {
			maxAge = nil
		}
		var methods []shared.CreateCorsPluginMethods = []shared.CreateCorsPluginMethods{}
		for _, methodsItem := range r.Config.Methods {
			methods = append(methods, shared.CreateCorsPluginMethods(methodsItem.ValueString()))
		}
		var origins []string = []string{}
		for _, originsItem := range r.Config.Origins {
			origins = append(origins, originsItem.ValueString())
		}
		preflightContinue := new(bool)
		if !r.Config.PreflightContinue.IsUnknown() && !r.Config.PreflightContinue.IsNull() {
			*preflightContinue = r.Config.PreflightContinue.ValueBool()
		} else {
			preflightContinue = nil
		}
		privateNetwork := new(bool)
		if !r.Config.PrivateNetwork.IsUnknown() && !r.Config.PrivateNetwork.IsNull() {
			*privateNetwork = r.Config.PrivateNetwork.ValueBool()
		} else {
			privateNetwork = nil
		}
		config = &shared.CreateCorsPluginConfig{
			Credentials:       credentials,
			ExposedHeaders:    exposedHeaders,
			Headers:           headers,
			MaxAge:            maxAge,
			Methods:           methods,
			Origins:           origins,
			PreflightContinue: preflightContinue,
			PrivateNetwork:    privateNetwork,
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
	var ordering *shared.CreateCorsPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateCorsPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateCorsPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateCorsPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateCorsPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateCorsPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateCorsPluginProtocols = []shared.CreateCorsPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateCorsPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateCorsPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateCorsPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateCorsPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateCorsPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateCorsPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateCorsPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateCorsPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateCorsPluginService{
			ID: id3,
		}
	}
	out := shared.CreateCorsPlugin{
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

func (r *GatewayPluginCorsResourceModel) RefreshFromSharedCorsPlugin(resp *shared.CorsPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateCorsPluginConfig{}
			r.Config.Credentials = types.BoolPointerValue(resp.Config.Credentials)
			r.Config.ExposedHeaders = []types.String{}
			for _, v := range resp.Config.ExposedHeaders {
				r.Config.ExposedHeaders = append(r.Config.ExposedHeaders, types.StringValue(v))
			}
			r.Config.Headers = []types.String{}
			for _, v := range resp.Config.Headers {
				r.Config.Headers = append(r.Config.Headers, types.StringValue(v))
			}
			if resp.Config.MaxAge != nil {
				r.Config.MaxAge = types.NumberValue(big.NewFloat(float64(*resp.Config.MaxAge)))
			} else {
				r.Config.MaxAge = types.NumberNull()
			}
			r.Config.Methods = []types.String{}
			for _, v := range resp.Config.Methods {
				r.Config.Methods = append(r.Config.Methods, types.StringValue(string(v)))
			}
			r.Config.Origins = []types.String{}
			for _, v := range resp.Config.Origins {
				r.Config.Origins = append(r.Config.Origins, types.StringValue(v))
			}
			r.Config.PreflightContinue = types.BoolPointerValue(resp.Config.PreflightContinue)
			r.Config.PrivateNetwork = types.BoolPointerValue(resp.Config.PrivateNetwork)
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