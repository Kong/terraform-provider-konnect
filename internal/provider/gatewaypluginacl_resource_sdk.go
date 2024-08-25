// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginACLResourceModel) ToSharedCreateACLPlugin() *shared.CreateACLPlugin {
	var config *shared.CreateACLPluginConfig
	if r.Config != nil {
		var allow []string = []string{}
		for _, allowItem := range r.Config.Allow {
			allow = append(allow, allowItem.ValueString())
		}
		var deny []string = []string{}
		for _, denyItem := range r.Config.Deny {
			deny = append(deny, denyItem.ValueString())
		}
		hideGroupsHeader := new(bool)
		if !r.Config.HideGroupsHeader.IsUnknown() && !r.Config.HideGroupsHeader.IsNull() {
			*hideGroupsHeader = r.Config.HideGroupsHeader.ValueBool()
		} else {
			hideGroupsHeader = nil
		}
		includeConsumerGroups := new(bool)
		if !r.Config.IncludeConsumerGroups.IsUnknown() && !r.Config.IncludeConsumerGroups.IsNull() {
			*includeConsumerGroups = r.Config.IncludeConsumerGroups.ValueBool()
		} else {
			includeConsumerGroups = nil
		}
		config = &shared.CreateACLPluginConfig{
			Allow:                 allow,
			Deny:                  deny,
			HideGroupsHeader:      hideGroupsHeader,
			IncludeConsumerGroups: includeConsumerGroups,
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
	var protocols []shared.CreateACLPluginProtocols = []shared.CreateACLPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateACLPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateACLPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateACLPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateACLPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateACLPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateACLPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateACLPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateACLPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateACLPluginService{
			ID: id3,
		}
	}
	out := shared.CreateACLPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginACLResourceModel) RefreshFromSharedACLPlugin(resp *shared.ACLPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateACLPluginConfig{}
			r.Config.Allow = []types.String{}
			for _, v := range resp.Config.Allow {
				r.Config.Allow = append(r.Config.Allow, types.StringValue(v))
			}
			r.Config.Deny = []types.String{}
			for _, v := range resp.Config.Deny {
				r.Config.Deny = append(r.Config.Deny, types.StringValue(v))
			}
			r.Config.HideGroupsHeader = types.BoolPointerValue(resp.Config.HideGroupsHeader)
			r.Config.IncludeConsumerGroups = types.BoolPointerValue(resp.Config.IncludeConsumerGroups)
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