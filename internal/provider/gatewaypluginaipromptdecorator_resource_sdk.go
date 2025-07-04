// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginAiPromptDecoratorResourceModel) ToSharedAiPromptDecoratorPlugin(ctx context.Context) (*shared.AiPromptDecoratorPlugin, diag.Diagnostics) {
	var diags diag.Diagnostics

	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.AiPromptDecoratorPluginOrdering
	if r.Ordering != nil {
		var after *shared.AiPromptDecoratorPluginAfter
		if r.Ordering.After != nil {
			access := make([]string, 0, len(r.Ordering.After.Access))
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AiPromptDecoratorPluginAfter{
				Access: access,
			}
		}
		var before *shared.AiPromptDecoratorPluginBefore
		if r.Ordering.Before != nil {
			access1 := make([]string, 0, len(r.Ordering.Before.Access))
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AiPromptDecoratorPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AiPromptDecoratorPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.AiPromptDecoratorPluginPartials
	if r.Partials != nil {
		partials = make([]shared.AiPromptDecoratorPluginPartials, 0, len(r.Partials))
		for _, partialsItem := range r.Partials {
			id1 := new(string)
			if !partialsItem.ID.IsUnknown() && !partialsItem.ID.IsNull() {
				*id1 = partialsItem.ID.ValueString()
			} else {
				id1 = nil
			}
			name := new(string)
			if !partialsItem.Name.IsUnknown() && !partialsItem.Name.IsNull() {
				*name = partialsItem.Name.ValueString()
			} else {
				name = nil
			}
			path := new(string)
			if !partialsItem.Path.IsUnknown() && !partialsItem.Path.IsNull() {
				*path = partialsItem.Path.ValueString()
			} else {
				path = nil
			}
			partials = append(partials, shared.AiPromptDecoratorPluginPartials{
				ID:   id1,
				Name: name,
				Path: path,
			})
		}
	}
	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.AiPromptDecoratorPluginConfig
	if r.Config != nil {
		llmFormat := new(shared.LlmFormat)
		if !r.Config.LlmFormat.IsUnknown() && !r.Config.LlmFormat.IsNull() {
			*llmFormat = shared.LlmFormat(r.Config.LlmFormat.ValueString())
		} else {
			llmFormat = nil
		}
		maxRequestBodySize := new(int64)
		if !r.Config.MaxRequestBodySize.IsUnknown() && !r.Config.MaxRequestBodySize.IsNull() {
			*maxRequestBodySize = r.Config.MaxRequestBodySize.ValueInt64()
		} else {
			maxRequestBodySize = nil
		}
		var prompts *shared.Prompts
		if r.Config.Prompts != nil {
			append1 := make([]shared.AiPromptDecoratorPluginAppend, 0, len(r.Config.Prompts.Append))
			for _, appendItem := range r.Config.Prompts.Append {
				var content string
				content = appendItem.Content.ValueString()

				role := new(shared.Role)
				if !appendItem.Role.IsUnknown() && !appendItem.Role.IsNull() {
					*role = shared.Role(appendItem.Role.ValueString())
				} else {
					role = nil
				}
				append1 = append(append1, shared.AiPromptDecoratorPluginAppend{
					Content: content,
					Role:    role,
				})
			}
			prepend := make([]shared.Prepend, 0, len(r.Config.Prompts.Prepend))
			for _, prependItem := range r.Config.Prompts.Prepend {
				var content1 string
				content1 = prependItem.Content.ValueString()

				role1 := new(shared.AiPromptDecoratorPluginRole)
				if !prependItem.Role.IsUnknown() && !prependItem.Role.IsNull() {
					*role1 = shared.AiPromptDecoratorPluginRole(prependItem.Role.ValueString())
				} else {
					role1 = nil
				}
				prepend = append(prepend, shared.Prepend{
					Content: content1,
					Role:    role1,
				})
			}
			prompts = &shared.Prompts{
				Append:  append1,
				Prepend: prepend,
			}
		}
		config = &shared.AiPromptDecoratorPluginConfig{
			LlmFormat:          llmFormat,
			MaxRequestBodySize: maxRequestBodySize,
			Prompts:            prompts,
		}
	}
	var consumer *shared.AiPromptDecoratorPluginConsumer
	if r.Consumer != nil {
		id2 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id2 = r.Consumer.ID.ValueString()
		} else {
			id2 = nil
		}
		consumer = &shared.AiPromptDecoratorPluginConsumer{
			ID: id2,
		}
	}
	var consumerGroup *shared.AiPromptDecoratorPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id3 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id3 = r.ConsumerGroup.ID.ValueString()
		} else {
			id3 = nil
		}
		consumerGroup = &shared.AiPromptDecoratorPluginConsumerGroup{
			ID: id3,
		}
	}
	protocols := make([]shared.AiPromptDecoratorPluginProtocols, 0, len(r.Protocols))
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AiPromptDecoratorPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AiPromptDecoratorPluginRoute
	if r.Route != nil {
		id4 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id4 = r.Route.ID.ValueString()
		} else {
			id4 = nil
		}
		route = &shared.AiPromptDecoratorPluginRoute{
			ID: id4,
		}
	}
	var service *shared.AiPromptDecoratorPluginService
	if r.Service != nil {
		id5 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id5 = r.Service.ID.ValueString()
		} else {
			id5 = nil
		}
		service = &shared.AiPromptDecoratorPluginService{
			ID: id5,
		}
	}
	out := shared.AiPromptDecoratorPlugin{
		CreatedAt:     createdAt,
		Enabled:       enabled,
		ID:            id,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Partials:      partials,
		Tags:          tags,
		UpdatedAt:     updatedAt,
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
	}

	return &out, diags
}

func (r *GatewayPluginAiPromptDecoratorResourceModel) ToOperationsCreateAipromptdecoratorPluginRequest(ctx context.Context) (*operations.CreateAipromptdecoratorPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	aiPromptDecoratorPlugin, aiPromptDecoratorPluginDiags := r.ToSharedAiPromptDecoratorPlugin(ctx)
	diags.Append(aiPromptDecoratorPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateAipromptdecoratorPluginRequest{
		ControlPlaneID:          controlPlaneID,
		AiPromptDecoratorPlugin: *aiPromptDecoratorPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginAiPromptDecoratorResourceModel) ToOperationsUpdateAipromptdecoratorPluginRequest(ctx context.Context) (*operations.UpdateAipromptdecoratorPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	aiPromptDecoratorPlugin, aiPromptDecoratorPluginDiags := r.ToSharedAiPromptDecoratorPlugin(ctx)
	diags.Append(aiPromptDecoratorPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateAipromptdecoratorPluginRequest{
		PluginID:                pluginID,
		ControlPlaneID:          controlPlaneID,
		AiPromptDecoratorPlugin: *aiPromptDecoratorPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginAiPromptDecoratorResourceModel) ToOperationsGetAipromptdecoratorPluginRequest(ctx context.Context) (*operations.GetAipromptdecoratorPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetAipromptdecoratorPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginAiPromptDecoratorResourceModel) ToOperationsDeleteAipromptdecoratorPluginRequest(ctx context.Context) (*operations.DeleteAipromptdecoratorPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.DeleteAipromptdecoratorPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginAiPromptDecoratorResourceModel) RefreshFromSharedAiPromptDecoratorPlugin(ctx context.Context, resp *shared.AiPromptDecoratorPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.AiPromptDecoratorPluginConfig{}
			if resp.Config.LlmFormat != nil {
				r.Config.LlmFormat = types.StringValue(string(*resp.Config.LlmFormat))
			} else {
				r.Config.LlmFormat = types.StringNull()
			}
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			if resp.Config.Prompts == nil {
				r.Config.Prompts = nil
			} else {
				r.Config.Prompts = &tfTypes.Prompts{}
				r.Config.Prompts.Append = []tfTypes.AiPromptDecoratorPluginAppend{}
				if len(r.Config.Prompts.Append) > len(resp.Config.Prompts.Append) {
					r.Config.Prompts.Append = r.Config.Prompts.Append[:len(resp.Config.Prompts.Append)]
				}
				for appendCount, appendItem := range resp.Config.Prompts.Append {
					var append1 tfTypes.AiPromptDecoratorPluginAppend
					append1.Content = types.StringValue(appendItem.Content)
					if appendItem.Role != nil {
						append1.Role = types.StringValue(string(*appendItem.Role))
					} else {
						append1.Role = types.StringNull()
					}
					if appendCount+1 > len(r.Config.Prompts.Append) {
						r.Config.Prompts.Append = append(r.Config.Prompts.Append, append1)
					} else {
						r.Config.Prompts.Append[appendCount].Content = append1.Content
						r.Config.Prompts.Append[appendCount].Role = append1.Role
					}
				}
				r.Config.Prompts.Prepend = []tfTypes.AiPromptDecoratorPluginAppend{}
				if len(r.Config.Prompts.Prepend) > len(resp.Config.Prompts.Prepend) {
					r.Config.Prompts.Prepend = r.Config.Prompts.Prepend[:len(resp.Config.Prompts.Prepend)]
				}
				for prependCount, prependItem := range resp.Config.Prompts.Prepend {
					var prepend tfTypes.AiPromptDecoratorPluginAppend
					prepend.Content = types.StringValue(prependItem.Content)
					if prependItem.Role != nil {
						prepend.Role = types.StringValue(string(*prependItem.Role))
					} else {
						prepend.Role = types.StringNull()
					}
					if prependCount+1 > len(r.Config.Prompts.Prepend) {
						r.Config.Prompts.Prepend = append(r.Config.Prompts.Prepend, prepend)
					} else {
						r.Config.Prompts.Prepend[prependCount].Content = prepend.Content
						r.Config.Prompts.Prepend[prependCount].Role = prepend.Role
					}
				}
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.Set{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.Set{}
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
				r.Ordering.After.Access = make([]types.String, 0, len(resp.Ordering.After.Access))
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = make([]types.String, 0, len(resp.Ordering.Before.Access))
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		if resp.Partials != nil {
			r.Partials = []tfTypes.Partials{}
			if len(r.Partials) > len(resp.Partials) {
				r.Partials = r.Partials[:len(resp.Partials)]
			}
			for partialsCount, partialsItem := range resp.Partials {
				var partials tfTypes.Partials
				partials.ID = types.StringPointerValue(partialsItem.ID)
				partials.Name = types.StringPointerValue(partialsItem.Name)
				partials.Path = types.StringPointerValue(partialsItem.Path)
				if partialsCount+1 > len(r.Partials) {
					r.Partials = append(r.Partials, partials)
				} else {
					r.Partials[partialsCount].ID = partials.ID
					r.Partials[partialsCount].Name = partials.Name
					r.Partials[partialsCount].Path = partials.Path
				}
			}
		}
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.Set{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.Set{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}

	return diags
}
