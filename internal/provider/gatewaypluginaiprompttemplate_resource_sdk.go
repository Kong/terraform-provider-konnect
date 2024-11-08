// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginAiPromptTemplateResourceModel) ToSharedAiPromptTemplatePluginInput() *shared.AiPromptTemplatePluginInput {
	allowUntemplatedRequests := new(bool)
	if !r.Config.AllowUntemplatedRequests.IsUnknown() && !r.Config.AllowUntemplatedRequests.IsNull() {
		*allowUntemplatedRequests = r.Config.AllowUntemplatedRequests.ValueBool()
	} else {
		allowUntemplatedRequests = nil
	}
	logOriginalRequest := new(bool)
	if !r.Config.LogOriginalRequest.IsUnknown() && !r.Config.LogOriginalRequest.IsNull() {
		*logOriginalRequest = r.Config.LogOriginalRequest.ValueBool()
	} else {
		logOriginalRequest = nil
	}
	maxRequestBodySize := new(int64)
	if !r.Config.MaxRequestBodySize.IsUnknown() && !r.Config.MaxRequestBodySize.IsNull() {
		*maxRequestBodySize = r.Config.MaxRequestBodySize.ValueInt64()
	} else {
		maxRequestBodySize = nil
	}
	var templates []shared.Templates = []shared.Templates{}
	for _, templatesItem := range r.Config.Templates {
		var name string
		name = templatesItem.Name.ValueString()

		var template string
		template = templatesItem.Template.ValueString()

		templates = append(templates, shared.Templates{
			Name:     name,
			Template: template,
		})
	}
	config := shared.AiPromptTemplatePluginConfig{
		AllowUntemplatedRequests: allowUntemplatedRequests,
		LogOriginalRequest:       logOriginalRequest,
		MaxRequestBodySize:       maxRequestBodySize,
		Templates:                templates,
	}
	var consumer *shared.AiPromptTemplatePluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.AiPromptTemplatePluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.AiPromptTemplatePluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.AiPromptTemplatePluginConsumerGroup{
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
	var ordering *shared.AiPromptTemplatePluginOrdering
	if r.Ordering != nil {
		var after *shared.AiPromptTemplatePluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AiPromptTemplatePluginAfter{
				Access: access,
			}
		}
		var before *shared.AiPromptTemplatePluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AiPromptTemplatePluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AiPromptTemplatePluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.AiPromptTemplatePluginProtocols = []shared.AiPromptTemplatePluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AiPromptTemplatePluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AiPromptTemplatePluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.AiPromptTemplatePluginRoute{
			ID: id3,
		}
	}
	var service *shared.AiPromptTemplatePluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.AiPromptTemplatePluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.AiPromptTemplatePluginInput{
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

func (r *GatewayPluginAiPromptTemplateResourceModel) RefreshFromSharedAiPromptTemplatePlugin(resp *shared.AiPromptTemplatePlugin) {
	if resp != nil {
		r.Config.AllowUntemplatedRequests = types.BoolPointerValue(resp.Config.AllowUntemplatedRequests)
		r.Config.LogOriginalRequest = types.BoolPointerValue(resp.Config.LogOriginalRequest)
		r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
		r.Config.Templates = []tfTypes.Templates{}
		if len(r.Config.Templates) > len(resp.Config.Templates) {
			r.Config.Templates = r.Config.Templates[:len(resp.Config.Templates)]
		}
		for templatesCount, templatesItem := range resp.Config.Templates {
			var templates1 tfTypes.Templates
			templates1.Name = types.StringValue(templatesItem.Name)
			templates1.Template = types.StringValue(templatesItem.Template)
			if templatesCount+1 > len(r.Config.Templates) {
				r.Config.Templates = append(r.Config.Templates, templates1)
			} else {
				r.Config.Templates[templatesCount].Name = templates1.Name
				r.Config.Templates[templatesCount].Template = templates1.Template
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
