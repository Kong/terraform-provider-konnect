// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginAIPromptGuardDataSourceModel) RefreshFromSharedAIPromptGuardPlugin(resp *shared.AIPromptGuardPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAIPromptGuardPluginConfig{}
			r.Config.AllowAllConversationHistory = types.BoolPointerValue(resp.Config.AllowAllConversationHistory)
			r.Config.AllowPatterns = []types.String{}
			for _, v := range resp.Config.AllowPatterns {
				r.Config.AllowPatterns = append(r.Config.AllowPatterns, types.StringValue(v))
			}
			r.Config.DenyPatterns = []types.String{}
			for _, v := range resp.Config.DenyPatterns {
				r.Config.DenyPatterns = append(r.Config.DenyPatterns, types.StringValue(v))
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
