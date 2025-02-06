// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginPrometheusResourceModel) ToSharedPrometheusPluginInput() *shared.PrometheusPluginInput {
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
	var ordering *shared.PrometheusPluginOrdering
	if r.Ordering != nil {
		var after *shared.PrometheusPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.PrometheusPluginAfter{
				Access: access,
			}
		}
		var before *shared.PrometheusPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.PrometheusPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.PrometheusPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	aiMetrics := new(bool)
	if !r.Config.AiMetrics.IsUnknown() && !r.Config.AiMetrics.IsNull() {
		*aiMetrics = r.Config.AiMetrics.ValueBool()
	} else {
		aiMetrics = nil
	}
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
	config := shared.PrometheusPluginConfig{
		AiMetrics:             aiMetrics,
		BandwidthMetrics:      bandwidthMetrics,
		LatencyMetrics:        latencyMetrics,
		PerConsumer:           perConsumer,
		StatusCodeMetrics:     statusCodeMetrics,
		UpstreamHealthMetrics: upstreamHealthMetrics,
	}
	var consumer *shared.PrometheusPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.PrometheusPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.PrometheusPluginProtocols = []shared.PrometheusPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.PrometheusPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.PrometheusPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.PrometheusPluginRoute{
			ID: id2,
		}
	}
	var service *shared.PrometheusPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.PrometheusPluginService{
			ID: id3,
		}
	}
	out := shared.PrometheusPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginPrometheusResourceModel) RefreshFromSharedPrometheusPlugin(resp *shared.PrometheusPlugin) {
	if resp != nil {
		r.Config.AiMetrics = types.BoolPointerValue(resp.Config.AiMetrics)
		r.Config.BandwidthMetrics = types.BoolPointerValue(resp.Config.BandwidthMetrics)
		r.Config.LatencyMetrics = types.BoolPointerValue(resp.Config.LatencyMetrics)
		r.Config.PerConsumer = types.BoolPointerValue(resp.Config.PerConsumer)
		r.Config.StatusCodeMetrics = types.BoolPointerValue(resp.Config.StatusCodeMetrics)
		r.Config.UpstreamHealthMetrics = types.BoolPointerValue(resp.Config.UpstreamHealthMetrics)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
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
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
