// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginPreFunctionResourceModel) ToSharedPreFunctionPluginInput() *shared.PreFunctionPluginInput {
	var access []string = []string{}
	for _, accessItem := range r.Config.Access {
		access = append(access, accessItem.ValueString())
	}
	var bodyFilter []string = []string{}
	for _, bodyFilterItem := range r.Config.BodyFilter {
		bodyFilter = append(bodyFilter, bodyFilterItem.ValueString())
	}
	var certificate []string = []string{}
	for _, certificateItem := range r.Config.Certificate {
		certificate = append(certificate, certificateItem.ValueString())
	}
	var headerFilter []string = []string{}
	for _, headerFilterItem := range r.Config.HeaderFilter {
		headerFilter = append(headerFilter, headerFilterItem.ValueString())
	}
	var log []string = []string{}
	for _, logItem := range r.Config.Log {
		log = append(log, logItem.ValueString())
	}
	var rewrite []string = []string{}
	for _, rewriteItem := range r.Config.Rewrite {
		rewrite = append(rewrite, rewriteItem.ValueString())
	}
	var wsClientFrame []string = []string{}
	for _, wsClientFrameItem := range r.Config.WsClientFrame {
		wsClientFrame = append(wsClientFrame, wsClientFrameItem.ValueString())
	}
	var wsClose []string = []string{}
	for _, wsCloseItem := range r.Config.WsClose {
		wsClose = append(wsClose, wsCloseItem.ValueString())
	}
	var wsHandshake []string = []string{}
	for _, wsHandshakeItem := range r.Config.WsHandshake {
		wsHandshake = append(wsHandshake, wsHandshakeItem.ValueString())
	}
	var wsUpstreamFrame []string = []string{}
	for _, wsUpstreamFrameItem := range r.Config.WsUpstreamFrame {
		wsUpstreamFrame = append(wsUpstreamFrame, wsUpstreamFrameItem.ValueString())
	}
	config := shared.PreFunctionPluginConfig{
		Access:          access,
		BodyFilter:      bodyFilter,
		Certificate:     certificate,
		HeaderFilter:    headerFilter,
		Log:             log,
		Rewrite:         rewrite,
		WsClientFrame:   wsClientFrame,
		WsClose:         wsClose,
		WsHandshake:     wsHandshake,
		WsUpstreamFrame: wsUpstreamFrame,
	}
	var consumer *shared.PreFunctionPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.PreFunctionPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.PreFunctionPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.PreFunctionPluginConsumerGroup{
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
	var ordering *shared.PreFunctionPluginOrdering
	if r.Ordering != nil {
		var after *shared.PreFunctionPluginAfter
		if r.Ordering.After != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.After.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			after = &shared.PreFunctionPluginAfter{
				Access: access1,
			}
		}
		var before *shared.PreFunctionPluginBefore
		if r.Ordering.Before != nil {
			var access2 []string = []string{}
			for _, accessItem2 := range r.Ordering.Before.Access {
				access2 = append(access2, accessItem2.ValueString())
			}
			before = &shared.PreFunctionPluginBefore{
				Access: access2,
			}
		}
		ordering = &shared.PreFunctionPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.PreFunctionPluginProtocols = []shared.PreFunctionPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.PreFunctionPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.PreFunctionPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.PreFunctionPluginRoute{
			ID: id3,
		}
	}
	var service *shared.PreFunctionPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.PreFunctionPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.PreFunctionPluginInput{
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

func (r *GatewayPluginPreFunctionResourceModel) RefreshFromSharedPreFunctionPlugin(resp *shared.PreFunctionPlugin) {
	if resp != nil {
		r.Config.Access = []types.String{}
		for _, v := range resp.Config.Access {
			r.Config.Access = append(r.Config.Access, types.StringValue(v))
		}
		r.Config.BodyFilter = []types.String{}
		for _, v := range resp.Config.BodyFilter {
			r.Config.BodyFilter = append(r.Config.BodyFilter, types.StringValue(v))
		}
		r.Config.Certificate = []types.String{}
		for _, v := range resp.Config.Certificate {
			r.Config.Certificate = append(r.Config.Certificate, types.StringValue(v))
		}
		r.Config.HeaderFilter = []types.String{}
		for _, v := range resp.Config.HeaderFilter {
			r.Config.HeaderFilter = append(r.Config.HeaderFilter, types.StringValue(v))
		}
		r.Config.Log = []types.String{}
		for _, v := range resp.Config.Log {
			r.Config.Log = append(r.Config.Log, types.StringValue(v))
		}
		r.Config.Rewrite = []types.String{}
		for _, v := range resp.Config.Rewrite {
			r.Config.Rewrite = append(r.Config.Rewrite, types.StringValue(v))
		}
		r.Config.WsClientFrame = []types.String{}
		for _, v := range resp.Config.WsClientFrame {
			r.Config.WsClientFrame = append(r.Config.WsClientFrame, types.StringValue(v))
		}
		r.Config.WsClose = []types.String{}
		for _, v := range resp.Config.WsClose {
			r.Config.WsClose = append(r.Config.WsClose, types.StringValue(v))
		}
		r.Config.WsHandshake = []types.String{}
		for _, v := range resp.Config.WsHandshake {
			r.Config.WsHandshake = append(r.Config.WsHandshake, types.StringValue(v))
		}
		r.Config.WsUpstreamFrame = []types.String{}
		for _, v := range resp.Config.WsUpstreamFrame {
			r.Config.WsUpstreamFrame = append(r.Config.WsUpstreamFrame, types.StringValue(v))
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
