// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginCanaryResourceModel) ToSharedCanaryPluginInput() *shared.CanaryPluginInput {
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
	ordering := make(map[string]string)
	for orderingKey, orderingValue := range r.Ordering {
		var orderingInst string
		orderingInst = orderingValue.ValueString()

		ordering[orderingKey] = orderingInst
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	canaryByHeaderName := new(string)
	if !r.Config.CanaryByHeaderName.IsUnknown() && !r.Config.CanaryByHeaderName.IsNull() {
		*canaryByHeaderName = r.Config.CanaryByHeaderName.ValueString()
	} else {
		canaryByHeaderName = nil
	}
	duration := new(float64)
	if !r.Config.Duration.IsUnknown() && !r.Config.Duration.IsNull() {
		*duration, _ = r.Config.Duration.ValueBigFloat().Float64()
	} else {
		duration = nil
	}
	var groups []string = []string{}
	for _, groupsItem := range r.Config.Groups {
		groups = append(groups, groupsItem.ValueString())
	}
	hash := new(shared.Hash)
	if !r.Config.Hash.IsUnknown() && !r.Config.Hash.IsNull() {
		*hash = shared.Hash(r.Config.Hash.ValueString())
	} else {
		hash = nil
	}
	hashHeader := new(string)
	if !r.Config.HashHeader.IsUnknown() && !r.Config.HashHeader.IsNull() {
		*hashHeader = r.Config.HashHeader.ValueString()
	} else {
		hashHeader = nil
	}
	percentage := new(float64)
	if !r.Config.Percentage.IsUnknown() && !r.Config.Percentage.IsNull() {
		*percentage, _ = r.Config.Percentage.ValueBigFloat().Float64()
	} else {
		percentage = nil
	}
	start := new(float64)
	if !r.Config.Start.IsUnknown() && !r.Config.Start.IsNull() {
		*start, _ = r.Config.Start.ValueBigFloat().Float64()
	} else {
		start = nil
	}
	steps := new(float64)
	if !r.Config.Steps.IsUnknown() && !r.Config.Steps.IsNull() {
		*steps, _ = r.Config.Steps.ValueBigFloat().Float64()
	} else {
		steps = nil
	}
	upstreamFallback := new(bool)
	if !r.Config.UpstreamFallback.IsUnknown() && !r.Config.UpstreamFallback.IsNull() {
		*upstreamFallback = r.Config.UpstreamFallback.ValueBool()
	} else {
		upstreamFallback = nil
	}
	upstreamHost := new(string)
	if !r.Config.UpstreamHost.IsUnknown() && !r.Config.UpstreamHost.IsNull() {
		*upstreamHost = r.Config.UpstreamHost.ValueString()
	} else {
		upstreamHost = nil
	}
	upstreamPort := new(int64)
	if !r.Config.UpstreamPort.IsUnknown() && !r.Config.UpstreamPort.IsNull() {
		*upstreamPort = r.Config.UpstreamPort.ValueInt64()
	} else {
		upstreamPort = nil
	}
	upstreamURI := new(string)
	if !r.Config.UpstreamURI.IsUnknown() && !r.Config.UpstreamURI.IsNull() {
		*upstreamURI = r.Config.UpstreamURI.ValueString()
	} else {
		upstreamURI = nil
	}
	config := shared.CanaryPluginConfig{
		CanaryByHeaderName: canaryByHeaderName,
		Duration:           duration,
		Groups:             groups,
		Hash:               hash,
		HashHeader:         hashHeader,
		Percentage:         percentage,
		Start:              start,
		Steps:              steps,
		UpstreamFallback:   upstreamFallback,
		UpstreamHost:       upstreamHost,
		UpstreamPort:       upstreamPort,
		UpstreamURI:        upstreamURI,
	}
	var protocols []shared.CanaryPluginProtocols = []shared.CanaryPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CanaryPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.CanaryPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CanaryPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CanaryPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CanaryPluginService{
			ID: id2,
		}
	}
	out := shared.CanaryPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginCanaryResourceModel) RefreshFromSharedCanaryPlugin(resp *shared.CanaryPlugin) {
	if resp != nil {
		r.Config.CanaryByHeaderName = types.StringPointerValue(resp.Config.CanaryByHeaderName)
		if resp.Config.Duration != nil {
			r.Config.Duration = types.NumberValue(big.NewFloat(float64(*resp.Config.Duration)))
		} else {
			r.Config.Duration = types.NumberNull()
		}
		r.Config.Groups = make([]types.String, 0, len(resp.Config.Groups))
		for _, v := range resp.Config.Groups {
			r.Config.Groups = append(r.Config.Groups, types.StringValue(v))
		}
		if resp.Config.Hash != nil {
			r.Config.Hash = types.StringValue(string(*resp.Config.Hash))
		} else {
			r.Config.Hash = types.StringNull()
		}
		r.Config.HashHeader = types.StringPointerValue(resp.Config.HashHeader)
		if resp.Config.Percentage != nil {
			r.Config.Percentage = types.NumberValue(big.NewFloat(float64(*resp.Config.Percentage)))
		} else {
			r.Config.Percentage = types.NumberNull()
		}
		if resp.Config.Start != nil {
			r.Config.Start = types.NumberValue(big.NewFloat(float64(*resp.Config.Start)))
		} else {
			r.Config.Start = types.NumberNull()
		}
		if resp.Config.Steps != nil {
			r.Config.Steps = types.NumberValue(big.NewFloat(float64(*resp.Config.Steps)))
		} else {
			r.Config.Steps = types.NumberNull()
		}
		r.Config.UpstreamFallback = types.BoolPointerValue(resp.Config.UpstreamFallback)
		r.Config.UpstreamHost = types.StringPointerValue(resp.Config.UpstreamHost)
		r.Config.UpstreamPort = types.Int64PointerValue(resp.Config.UpstreamPort)
		r.Config.UpstreamURI = types.StringPointerValue(resp.Config.UpstreamURI)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering != nil {
			r.Ordering = make(map[string]types.String, len(resp.Ordering))
			for key, value := range resp.Ordering {
				r.Ordering[key] = types.StringValue(value)
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
