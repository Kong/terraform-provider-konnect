// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginAiPromptDecoratorDataSourceModel) RefreshFromSharedAiPromptDecoratorPlugin(resp *shared.AiPromptDecoratorPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAiPromptDecoratorPluginConfig{}
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			if resp.Config.Prompts == nil {
				r.Config.Prompts = nil
			} else {
				r.Config.Prompts = &tfTypes.CreateAiPromptDecoratorPluginPrompts{}
				r.Config.Prompts.Append = []tfTypes.AiPromptDecoratorPluginAppend{}
				if len(r.Config.Prompts.Append) > len(resp.Config.Prompts.Append) {
					r.Config.Prompts.Append = r.Config.Prompts.Append[:len(resp.Config.Prompts.Append)]
				}
				for appendCount, appendItem := range resp.Config.Prompts.Append {
					var append2 tfTypes.AiPromptDecoratorPluginAppend
					append2.Content = types.StringValue(appendItem.Content)
					if appendItem.Role != nil {
						append2.Role = types.StringValue(string(*appendItem.Role))
					} else {
						append2.Role = types.StringNull()
					}
					if appendCount+1 > len(r.Config.Prompts.Append) {
						r.Config.Prompts.Append = append(r.Config.Prompts.Append, append2)
					} else {
						r.Config.Prompts.Append[appendCount].Content = append2.Content
						r.Config.Prompts.Append[appendCount].Role = append2.Role
					}
				}
				r.Config.Prompts.Prepend = []tfTypes.AiPromptDecoratorPluginAppend{}
				if len(r.Config.Prompts.Prepend) > len(resp.Config.Prompts.Prepend) {
					r.Config.Prompts.Prepend = r.Config.Prompts.Prepend[:len(resp.Config.Prompts.Prepend)]
				}
				for prependCount, prependItem := range resp.Config.Prompts.Prepend {
					var prepend1 tfTypes.AiPromptDecoratorPluginAppend
					prepend1.Content = types.StringValue(prependItem.Content)
					if prependItem.Role != nil {
						prepend1.Role = types.StringValue(string(*prependItem.Role))
					} else {
						prepend1.Role = types.StringNull()
					}
					if prependCount+1 > len(r.Config.Prompts.Prepend) {
						r.Config.Prompts.Prepend = append(r.Config.Prompts.Prepend, prepend1)
					} else {
						r.Config.Prompts.Prepend[prependCount].Content = prepend1.Content
						r.Config.Prompts.Prepend[prependCount].Role = prepend1.Role
					}
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
