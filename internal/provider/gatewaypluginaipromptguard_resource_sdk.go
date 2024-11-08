// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginAiPromptGuardResourceModel) ToSharedCreateAiPromptGuardPlugin() *shared.CreateAiPromptGuardPlugin {
	var config *shared.CreateAiPromptGuardPluginConfig
	if r.Config != nil {
		allowAllConversationHistory := new(bool)
		if !r.Config.AllowAllConversationHistory.IsUnknown() && !r.Config.AllowAllConversationHistory.IsNull() {
			*allowAllConversationHistory = r.Config.AllowAllConversationHistory.ValueBool()
		} else {
			allowAllConversationHistory = nil
		}
		var allowPatterns []string = []string{}
		for _, allowPatternsItem := range r.Config.AllowPatterns {
			allowPatterns = append(allowPatterns, allowPatternsItem.ValueString())
		}
		var denyPatterns []string = []string{}
		for _, denyPatternsItem := range r.Config.DenyPatterns {
			denyPatterns = append(denyPatterns, denyPatternsItem.ValueString())
		}
		matchAllRoles := new(bool)
		if !r.Config.MatchAllRoles.IsUnknown() && !r.Config.MatchAllRoles.IsNull() {
			*matchAllRoles = r.Config.MatchAllRoles.ValueBool()
		} else {
			matchAllRoles = nil
		}
		maxRequestBodySize := new(int64)
		if !r.Config.MaxRequestBodySize.IsUnknown() && !r.Config.MaxRequestBodySize.IsNull() {
			*maxRequestBodySize = r.Config.MaxRequestBodySize.ValueInt64()
		} else {
			maxRequestBodySize = nil
		}
		config = &shared.CreateAiPromptGuardPluginConfig{
			AllowAllConversationHistory: allowAllConversationHistory,
			AllowPatterns:               allowPatterns,
			DenyPatterns:                denyPatterns,
			MatchAllRoles:               matchAllRoles,
			MaxRequestBodySize:          maxRequestBodySize,
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
	var ordering *shared.CreateAiPromptGuardPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateAiPromptGuardPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateAiPromptGuardPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateAiPromptGuardPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateAiPromptGuardPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateAiPromptGuardPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateAiPromptGuardPluginProtocols = []shared.CreateAiPromptGuardPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateAiPromptGuardPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateAiPromptGuardPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateAiPromptGuardPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateAiPromptGuardPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateAiPromptGuardPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateAiPromptGuardPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateAiPromptGuardPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateAiPromptGuardPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateAiPromptGuardPluginService{
			ID: id3,
		}
	}
	out := shared.CreateAiPromptGuardPlugin{
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

func (r *GatewayPluginAiPromptGuardResourceModel) RefreshFromSharedAiPromptGuardPlugin(resp *shared.AiPromptGuardPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAiPromptGuardPluginConfig{}
			r.Config.AllowAllConversationHistory = types.BoolPointerValue(resp.Config.AllowAllConversationHistory)
			r.Config.AllowPatterns = []types.String{}
			for _, v := range resp.Config.AllowPatterns {
				r.Config.AllowPatterns = append(r.Config.AllowPatterns, types.StringValue(v))
			}
			r.Config.DenyPatterns = []types.String{}
			for _, v := range resp.Config.DenyPatterns {
				r.Config.DenyPatterns = append(r.Config.DenyPatterns, types.StringValue(v))
			}
			r.Config.MatchAllRoles = types.BoolPointerValue(resp.Config.MatchAllRoles)
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
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