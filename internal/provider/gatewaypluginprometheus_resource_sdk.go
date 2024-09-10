// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginPrometheusResourceModel) ToSharedCreatePrometheusPlugin() *shared.CreatePrometheusPlugin {
	var config *shared.CreatePrometheusPluginConfig
	if r.Config != nil {
		bandwidthMetrics := new(bool)
		if !r.Config.BandwidthMetrics.IsUnknown() && !r.Config.BandwidthMetrics.IsNull() {
			*bandwidthMetrics = r.Config.BandwidthMetrics.ValueBool()
		} else {
			bandwidthMetrics = nil
		}
		latencyMetrics := new(bool)
		if !r.Config.LatencyMetrics.IsUnknown() && !r.Config.LatencyMetrics.IsNull() {
			*latencyMetrics = r.Config.LatencyMetrics.ValueBool()
		} else {
			latencyMetrics = nil
		}
		perConsumer := new(bool)
		if !r.Config.PerConsumer.IsUnknown() && !r.Config.PerConsumer.IsNull() {
			*perConsumer = r.Config.PerConsumer.ValueBool()
		} else {
			perConsumer = nil
		}
		statusCodeMetrics := new(bool)
		if !r.Config.StatusCodeMetrics.IsUnknown() && !r.Config.StatusCodeMetrics.IsNull() {
			*statusCodeMetrics = r.Config.StatusCodeMetrics.ValueBool()
		} else {
			statusCodeMetrics = nil
		}
		upstreamHealthMetrics := new(bool)
		if !r.Config.UpstreamHealthMetrics.IsUnknown() && !r.Config.UpstreamHealthMetrics.IsNull() {
			*upstreamHealthMetrics = r.Config.UpstreamHealthMetrics.ValueBool()
		} else {
			upstreamHealthMetrics = nil
		}
		config = &shared.CreatePrometheusPluginConfig{
			BandwidthMetrics:      bandwidthMetrics,
			LatencyMetrics:        latencyMetrics,
			PerConsumer:           perConsumer,
			StatusCodeMetrics:     statusCodeMetrics,
			UpstreamHealthMetrics: upstreamHealthMetrics,
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
	var ordering *shared.CreatePrometheusPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreatePrometheusPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreatePrometheusPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreatePrometheusPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreatePrometheusPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreatePrometheusPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreatePrometheusPluginProtocols = []shared.CreatePrometheusPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreatePrometheusPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreatePrometheusPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreatePrometheusPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreatePrometheusPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreatePrometheusPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreatePrometheusPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreatePrometheusPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreatePrometheusPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreatePrometheusPluginService{
			ID: id3,
		}
	}
	out := shared.CreatePrometheusPlugin{
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

func (r *GatewayPluginPrometheusResourceModel) RefreshFromSharedPrometheusPlugin(resp *shared.PrometheusPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreatePrometheusPluginConfig{}
			r.Config.BandwidthMetrics = types.BoolPointerValue(resp.Config.BandwidthMetrics)
			r.Config.LatencyMetrics = types.BoolPointerValue(resp.Config.LatencyMetrics)
			r.Config.PerConsumer = types.BoolPointerValue(resp.Config.PerConsumer)
			r.Config.StatusCodeMetrics = types.BoolPointerValue(resp.Config.StatusCodeMetrics)
			r.Config.UpstreamHealthMetrics = types.BoolPointerValue(resp.Config.UpstreamHealthMetrics)
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
