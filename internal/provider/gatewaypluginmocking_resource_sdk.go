// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginMockingResourceModel) ToSharedMockingPluginInput() *shared.MockingPluginInput {
	apiSpecification := new(string)
	if !r.Config.APISpecification.IsUnknown() && !r.Config.APISpecification.IsNull() {
		*apiSpecification = r.Config.APISpecification.ValueString()
	} else {
		apiSpecification = nil
	}
	apiSpecificationFilename := new(string)
	if !r.Config.APISpecificationFilename.IsUnknown() && !r.Config.APISpecificationFilename.IsNull() {
		*apiSpecificationFilename = r.Config.APISpecificationFilename.ValueString()
	} else {
		apiSpecificationFilename = nil
	}
	customBasePath := new(string)
	if !r.Config.CustomBasePath.IsUnknown() && !r.Config.CustomBasePath.IsNull() {
		*customBasePath = r.Config.CustomBasePath.ValueString()
	} else {
		customBasePath = nil
	}
	includeBasePath := new(bool)
	if !r.Config.IncludeBasePath.IsUnknown() && !r.Config.IncludeBasePath.IsNull() {
		*includeBasePath = r.Config.IncludeBasePath.ValueBool()
	} else {
		includeBasePath = nil
	}
	var includedStatusCodes []int64 = []int64{}
	for _, includedStatusCodesItem := range r.Config.IncludedStatusCodes {
		includedStatusCodes = append(includedStatusCodes, includedStatusCodesItem.ValueInt64())
	}
	maxDelayTime := new(float64)
	if !r.Config.MaxDelayTime.IsUnknown() && !r.Config.MaxDelayTime.IsNull() {
		*maxDelayTime, _ = r.Config.MaxDelayTime.ValueBigFloat().Float64()
	} else {
		maxDelayTime = nil
	}
	minDelayTime := new(float64)
	if !r.Config.MinDelayTime.IsUnknown() && !r.Config.MinDelayTime.IsNull() {
		*minDelayTime, _ = r.Config.MinDelayTime.ValueBigFloat().Float64()
	} else {
		minDelayTime = nil
	}
	randomDelay := new(bool)
	if !r.Config.RandomDelay.IsUnknown() && !r.Config.RandomDelay.IsNull() {
		*randomDelay = r.Config.RandomDelay.ValueBool()
	} else {
		randomDelay = nil
	}
	randomExamples := new(bool)
	if !r.Config.RandomExamples.IsUnknown() && !r.Config.RandomExamples.IsNull() {
		*randomExamples = r.Config.RandomExamples.ValueBool()
	} else {
		randomExamples = nil
	}
	randomStatusCode := new(bool)
	if !r.Config.RandomStatusCode.IsUnknown() && !r.Config.RandomStatusCode.IsNull() {
		*randomStatusCode = r.Config.RandomStatusCode.ValueBool()
	} else {
		randomStatusCode = nil
	}
	config := shared.MockingPluginConfig{
		APISpecification:         apiSpecification,
		APISpecificationFilename: apiSpecificationFilename,
		CustomBasePath:           customBasePath,
		IncludeBasePath:          includeBasePath,
		IncludedStatusCodes:      includedStatusCodes,
		MaxDelayTime:             maxDelayTime,
		MinDelayTime:             minDelayTime,
		RandomDelay:              randomDelay,
		RandomExamples:           randomExamples,
		RandomStatusCode:         randomStatusCode,
	}
	var consumer *shared.MockingPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.MockingPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.MockingPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.MockingPluginConsumerGroup{
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
	var ordering *shared.MockingPluginOrdering
	if r.Ordering != nil {
		var after *shared.MockingPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.MockingPluginAfter{
				Access: access,
			}
		}
		var before *shared.MockingPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.MockingPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.MockingPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.MockingPluginProtocols = []shared.MockingPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.MockingPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.MockingPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.MockingPluginRoute{
			ID: id3,
		}
	}
	var service *shared.MockingPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.MockingPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.MockingPluginInput{
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

func (r *GatewayPluginMockingResourceModel) RefreshFromSharedMockingPlugin(resp *shared.MockingPlugin) {
	if resp != nil {
		r.Config.APISpecification = types.StringPointerValue(resp.Config.APISpecification)
		r.Config.APISpecificationFilename = types.StringPointerValue(resp.Config.APISpecificationFilename)
		r.Config.CustomBasePath = types.StringPointerValue(resp.Config.CustomBasePath)
		r.Config.IncludeBasePath = types.BoolPointerValue(resp.Config.IncludeBasePath)
		r.Config.IncludedStatusCodes = []types.Int64{}
		for _, v := range resp.Config.IncludedStatusCodes {
			r.Config.IncludedStatusCodes = append(r.Config.IncludedStatusCodes, types.Int64Value(v))
		}
		if resp.Config.MaxDelayTime != nil {
			r.Config.MaxDelayTime = types.NumberValue(big.NewFloat(float64(*resp.Config.MaxDelayTime)))
		} else {
			r.Config.MaxDelayTime = types.NumberNull()
		}
		if resp.Config.MinDelayTime != nil {
			r.Config.MinDelayTime = types.NumberValue(big.NewFloat(float64(*resp.Config.MinDelayTime)))
		} else {
			r.Config.MinDelayTime = types.NumberNull()
		}
		r.Config.RandomDelay = types.BoolPointerValue(resp.Config.RandomDelay)
		r.Config.RandomExamples = types.BoolPointerValue(resp.Config.RandomExamples)
		r.Config.RandomStatusCode = types.BoolPointerValue(resp.Config.RandomStatusCode)
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
