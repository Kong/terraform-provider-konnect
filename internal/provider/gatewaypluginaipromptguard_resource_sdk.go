// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginAIPromptGuardResourceModel) ToSharedCreateAIPromptGuardPlugin() *shared.CreateAIPromptGuardPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateAIPromptGuardPluginProtocols = []shared.CreateAIPromptGuardPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateAIPromptGuardPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateAIPromptGuardPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateAIPromptGuardPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateAIPromptGuardPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateAIPromptGuardPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateAIPromptGuardPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateAIPromptGuardPluginService{
			ID: id2,
		}
	}
	var allowPatterns []string = []string{}
	for _, allowPatternsItem := range r.Config.AllowPatterns {
		allowPatterns = append(allowPatterns, allowPatternsItem.ValueString())
	}
	var denyPatterns []string = []string{}
	for _, denyPatternsItem := range r.Config.DenyPatterns {
		denyPatterns = append(denyPatterns, denyPatternsItem.ValueString())
	}
	allowAllConversationHistory := new(bool)
	if !r.Config.AllowAllConversationHistory.IsUnknown() && !r.Config.AllowAllConversationHistory.IsNull() {
		*allowAllConversationHistory = r.Config.AllowAllConversationHistory.ValueBool()
	} else {
		allowAllConversationHistory = nil
	}
	config := shared.CreateAIPromptGuardPluginConfig{
		AllowPatterns:               allowPatterns,
		DenyPatterns:                denyPatterns,
		AllowAllConversationHistory: allowAllConversationHistory,
	}
	out := shared.CreateAIPromptGuardPlugin{
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

func (r *GatewayPluginAIPromptGuardResourceModel) RefreshFromSharedAIPromptGuardPlugin(resp *shared.AIPromptGuardPlugin) {
	if resp != nil {
		r.Config.AllowAllConversationHistory = types.BoolPointerValue(resp.Config.AllowAllConversationHistory)
		r.Config.AllowPatterns = []types.String{}
		for _, v := range resp.Config.AllowPatterns {
			r.Config.AllowPatterns = append(r.Config.AllowPatterns, types.StringValue(v))
		}
		r.Config.DenyPatterns = []types.String{}
		for _, v := range resp.Config.DenyPatterns {
			r.Config.DenyPatterns = append(r.Config.DenyPatterns, types.StringValue(v))
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
