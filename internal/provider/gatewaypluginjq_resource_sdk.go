// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginJQResourceModel) ToSharedCreateJQPlugin() *shared.CreateJQPlugin {
	var config *shared.CreateJQPluginConfig
	if r.Config != nil {
		var requestIfMediaType []string = []string{}
		for _, requestIfMediaTypeItem := range r.Config.RequestIfMediaType {
			requestIfMediaType = append(requestIfMediaType, requestIfMediaTypeItem.ValueString())
		}
		requestJqProgram := new(string)
		if !r.Config.RequestJqProgram.IsUnknown() && !r.Config.RequestJqProgram.IsNull() {
			*requestJqProgram = r.Config.RequestJqProgram.ValueString()
		} else {
			requestJqProgram = nil
		}
		var requestJqProgramOptions *shared.CreateJQPluginRequestJQProgramOptions
		if r.Config.RequestJqProgramOptions != nil {
			asciiOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.ASCIIOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.ASCIIOutput.IsNull() {
				*asciiOutput = r.Config.RequestJqProgramOptions.ASCIIOutput.ValueBool()
			} else {
				asciiOutput = nil
			}
			compactOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.CompactOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.CompactOutput.IsNull() {
				*compactOutput = r.Config.RequestJqProgramOptions.CompactOutput.ValueBool()
			} else {
				compactOutput = nil
			}
			joinOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.JoinOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.JoinOutput.IsNull() {
				*joinOutput = r.Config.RequestJqProgramOptions.JoinOutput.ValueBool()
			} else {
				joinOutput = nil
			}
			rawOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.RawOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.RawOutput.IsNull() {
				*rawOutput = r.Config.RequestJqProgramOptions.RawOutput.ValueBool()
			} else {
				rawOutput = nil
			}
			sortKeys := new(bool)
			if !r.Config.RequestJqProgramOptions.SortKeys.IsUnknown() && !r.Config.RequestJqProgramOptions.SortKeys.IsNull() {
				*sortKeys = r.Config.RequestJqProgramOptions.SortKeys.ValueBool()
			} else {
				sortKeys = nil
			}
			requestJqProgramOptions = &shared.CreateJQPluginRequestJQProgramOptions{
				ASCIIOutput:   asciiOutput,
				CompactOutput: compactOutput,
				JoinOutput:    joinOutput,
				RawOutput:     rawOutput,
				SortKeys:      sortKeys,
			}
		}
		var responseIfMediaType []string = []string{}
		for _, responseIfMediaTypeItem := range r.Config.ResponseIfMediaType {
			responseIfMediaType = append(responseIfMediaType, responseIfMediaTypeItem.ValueString())
		}
		var responseIfStatusCode []int64 = []int64{}
		for _, responseIfStatusCodeItem := range r.Config.ResponseIfStatusCode {
			responseIfStatusCode = append(responseIfStatusCode, responseIfStatusCodeItem.ValueInt64())
		}
		responseJqProgram := new(string)
		if !r.Config.ResponseJqProgram.IsUnknown() && !r.Config.ResponseJqProgram.IsNull() {
			*responseJqProgram = r.Config.ResponseJqProgram.ValueString()
		} else {
			responseJqProgram = nil
		}
		var responseJqProgramOptions *shared.CreateJQPluginResponseJQProgramOptions
		if r.Config.ResponseJqProgramOptions != nil {
			asciiOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.ASCIIOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.ASCIIOutput.IsNull() {
				*asciiOutput1 = r.Config.ResponseJqProgramOptions.ASCIIOutput.ValueBool()
			} else {
				asciiOutput1 = nil
			}
			compactOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.CompactOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.CompactOutput.IsNull() {
				*compactOutput1 = r.Config.ResponseJqProgramOptions.CompactOutput.ValueBool()
			} else {
				compactOutput1 = nil
			}
			joinOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.JoinOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.JoinOutput.IsNull() {
				*joinOutput1 = r.Config.ResponseJqProgramOptions.JoinOutput.ValueBool()
			} else {
				joinOutput1 = nil
			}
			rawOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.RawOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.RawOutput.IsNull() {
				*rawOutput1 = r.Config.ResponseJqProgramOptions.RawOutput.ValueBool()
			} else {
				rawOutput1 = nil
			}
			sortKeys1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.SortKeys.IsUnknown() && !r.Config.ResponseJqProgramOptions.SortKeys.IsNull() {
				*sortKeys1 = r.Config.ResponseJqProgramOptions.SortKeys.ValueBool()
			} else {
				sortKeys1 = nil
			}
			responseJqProgramOptions = &shared.CreateJQPluginResponseJQProgramOptions{
				ASCIIOutput:   asciiOutput1,
				CompactOutput: compactOutput1,
				JoinOutput:    joinOutput1,
				RawOutput:     rawOutput1,
				SortKeys:      sortKeys1,
			}
		}
		config = &shared.CreateJQPluginConfig{
			RequestIfMediaType:       requestIfMediaType,
			RequestJqProgram:         requestJqProgram,
			RequestJqProgramOptions:  requestJqProgramOptions,
			ResponseIfMediaType:      responseIfMediaType,
			ResponseIfStatusCode:     responseIfStatusCode,
			ResponseJqProgram:        responseJqProgram,
			ResponseJqProgramOptions: responseJqProgramOptions,
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
	var ordering *shared.CreateJQPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateJQPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateJQPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateJQPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateJQPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateJQPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateJQPluginProtocols = []shared.CreateJQPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateJQPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateJQPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateJQPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateJQPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateJQPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateJQPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateJQPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateJQPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateJQPluginService{
			ID: id3,
		}
	}
	out := shared.CreateJQPlugin{
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

func (r *GatewayPluginJQResourceModel) RefreshFromSharedJQPlugin(resp *shared.JQPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateJQPluginConfig{}
			r.Config.RequestIfMediaType = []types.String{}
			for _, v := range resp.Config.RequestIfMediaType {
				r.Config.RequestIfMediaType = append(r.Config.RequestIfMediaType, types.StringValue(v))
			}
			r.Config.RequestJqProgram = types.StringPointerValue(resp.Config.RequestJqProgram)
			if resp.Config.RequestJqProgramOptions == nil {
				r.Config.RequestJqProgramOptions = nil
			} else {
				r.Config.RequestJqProgramOptions = &tfTypes.CreateJQPluginRequestJQProgramOptions{}
				r.Config.RequestJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.ASCIIOutput)
				r.Config.RequestJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.CompactOutput)
				r.Config.RequestJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.JoinOutput)
				r.Config.RequestJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.RawOutput)
				r.Config.RequestJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.SortKeys)
			}
			r.Config.ResponseIfMediaType = []types.String{}
			for _, v := range resp.Config.ResponseIfMediaType {
				r.Config.ResponseIfMediaType = append(r.Config.ResponseIfMediaType, types.StringValue(v))
			}
			r.Config.ResponseIfStatusCode = []types.Int64{}
			for _, v := range resp.Config.ResponseIfStatusCode {
				r.Config.ResponseIfStatusCode = append(r.Config.ResponseIfStatusCode, types.Int64Value(v))
			}
			r.Config.ResponseJqProgram = types.StringPointerValue(resp.Config.ResponseJqProgram)
			if resp.Config.ResponseJqProgramOptions == nil {
				r.Config.ResponseJqProgramOptions = nil
			} else {
				r.Config.ResponseJqProgramOptions = &tfTypes.CreateJQPluginRequestJQProgramOptions{}
				r.Config.ResponseJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.ASCIIOutput)
				r.Config.ResponseJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.CompactOutput)
				r.Config.ResponseJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.JoinOutput)
				r.Config.ResponseJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.RawOutput)
				r.Config.ResponseJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.SortKeys)
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
