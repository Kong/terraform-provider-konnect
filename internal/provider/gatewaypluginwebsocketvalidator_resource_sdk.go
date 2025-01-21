// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginWebsocketValidatorResourceModel) ToSharedWebsocketValidatorPluginInput() *shared.WebsocketValidatorPluginInput {
	var client *shared.Client
	if r.Config.Client != nil {
		var binary *shared.Binary
		if r.Config.Client.Binary != nil {
			var schema string
			schema = r.Config.Client.Binary.Schema.ValueString()

			typeVar := shared.WebsocketValidatorPluginType(r.Config.Client.Binary.Type.ValueString())
			binary = &shared.Binary{
				Schema: schema,
				Type:   typeVar,
			}
		}
		var text *shared.WebsocketValidatorPluginConfigText
		if r.Config.Client.Text != nil {
			var schema1 string
			schema1 = r.Config.Client.Text.Schema.ValueString()

			typeVar1 := shared.WebsocketValidatorPluginConfigType(r.Config.Client.Text.Type.ValueString())
			text = &shared.WebsocketValidatorPluginConfigText{
				Schema: schema1,
				Type:   typeVar1,
			}
		}
		client = &shared.Client{
			Binary: binary,
			Text:   text,
		}
	}
	var upstream *shared.WebsocketValidatorPluginUpstream
	if r.Config.Upstream != nil {
		var binary1 *shared.WebsocketValidatorPluginBinary
		if r.Config.Upstream.Binary != nil {
			var schema2 string
			schema2 = r.Config.Upstream.Binary.Schema.ValueString()

			typeVar2 := shared.WebsocketValidatorPluginConfigUpstreamType(r.Config.Upstream.Binary.Type.ValueString())
			binary1 = &shared.WebsocketValidatorPluginBinary{
				Schema: schema2,
				Type:   typeVar2,
			}
		}
		var text1 *shared.WebsocketValidatorPluginText
		if r.Config.Upstream.Text != nil {
			var schema3 string
			schema3 = r.Config.Upstream.Text.Schema.ValueString()

			typeVar3 := shared.WebsocketValidatorPluginConfigUpstreamTextType(r.Config.Upstream.Text.Type.ValueString())
			text1 = &shared.WebsocketValidatorPluginText{
				Schema: schema3,
				Type:   typeVar3,
			}
		}
		upstream = &shared.WebsocketValidatorPluginUpstream{
			Binary: binary1,
			Text:   text1,
		}
	}
	config := shared.WebsocketValidatorPluginConfig{
		Client:   client,
		Upstream: upstream,
	}
	var consumer *shared.WebsocketValidatorPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.WebsocketValidatorPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.WebsocketValidatorPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.WebsocketValidatorPluginConsumerGroup{
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
	var ordering *shared.WebsocketValidatorPluginOrdering
	if r.Ordering != nil {
		var after *shared.WebsocketValidatorPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.WebsocketValidatorPluginAfter{
				Access: access,
			}
		}
		var before *shared.WebsocketValidatorPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.WebsocketValidatorPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.WebsocketValidatorPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.WebsocketValidatorPluginProtocols = []shared.WebsocketValidatorPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.WebsocketValidatorPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.WebsocketValidatorPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.WebsocketValidatorPluginRoute{
			ID: id3,
		}
	}
	var service *shared.WebsocketValidatorPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.WebsocketValidatorPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.WebsocketValidatorPluginInput{
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

func (r *GatewayPluginWebsocketValidatorResourceModel) RefreshFromSharedWebsocketValidatorPlugin(resp *shared.WebsocketValidatorPlugin) {
	if resp != nil {
		if resp.Config.Client == nil {
			r.Config.Client = nil
		} else {
			r.Config.Client = &tfTypes.Client{}
			if resp.Config.Client.Binary == nil {
				r.Config.Client.Binary = nil
			} else {
				r.Config.Client.Binary = &tfTypes.Binary{}
				r.Config.Client.Binary.Schema = types.StringValue(resp.Config.Client.Binary.Schema)
				r.Config.Client.Binary.Type = types.StringValue(string(resp.Config.Client.Binary.Type))
			}
			if resp.Config.Client.Text == nil {
				r.Config.Client.Text = nil
			} else {
				r.Config.Client.Text = &tfTypes.Binary{}
				r.Config.Client.Text.Schema = types.StringValue(resp.Config.Client.Text.Schema)
				r.Config.Client.Text.Type = types.StringValue(string(resp.Config.Client.Text.Type))
			}
		}
		if resp.Config.Upstream == nil {
			r.Config.Upstream = nil
		} else {
			r.Config.Upstream = &tfTypes.Client{}
			if resp.Config.Upstream.Binary == nil {
				r.Config.Upstream.Binary = nil
			} else {
				r.Config.Upstream.Binary = &tfTypes.Binary{}
				r.Config.Upstream.Binary.Schema = types.StringValue(resp.Config.Upstream.Binary.Schema)
				r.Config.Upstream.Binary.Type = types.StringValue(string(resp.Config.Upstream.Binary.Type))
			}
			if resp.Config.Upstream.Text == nil {
				r.Config.Upstream.Text = nil
			} else {
				r.Config.Upstream.Text = &tfTypes.Binary{}
				r.Config.Upstream.Text.Schema = types.StringValue(resp.Config.Upstream.Text.Schema)
				r.Config.Upstream.Text.Type = types.StringValue(string(resp.Config.Upstream.Text.Type))
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
