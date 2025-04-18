// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginStatsdAdvancedResourceModel) ToSharedStatsdAdvancedPlugin() *shared.StatsdAdvancedPlugin {
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
	var ordering *shared.StatsdAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.StatsdAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.StatsdAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.StatsdAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.StatsdAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.StatsdAdvancedPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.StatsdAdvancedPluginConfig
	if r.Config != nil {
		var allowStatusCodes []string = []string{}
		for _, allowStatusCodesItem := range r.Config.AllowStatusCodes {
			allowStatusCodes = append(allowStatusCodes, allowStatusCodesItem.ValueString())
		}
		consumerIdentifierDefault := new(shared.StatsdAdvancedPluginConsumerIdentifierDefault)
		if !r.Config.ConsumerIdentifierDefault.IsUnknown() && !r.Config.ConsumerIdentifierDefault.IsNull() {
			*consumerIdentifierDefault = shared.StatsdAdvancedPluginConsumerIdentifierDefault(r.Config.ConsumerIdentifierDefault.ValueString())
		} else {
			consumerIdentifierDefault = nil
		}
		host := new(string)
		if !r.Config.Host.IsUnknown() && !r.Config.Host.IsNull() {
			*host = r.Config.Host.ValueString()
		} else {
			host = nil
		}
		hostnameInPrefix := new(bool)
		if !r.Config.HostnameInPrefix.IsUnknown() && !r.Config.HostnameInPrefix.IsNull() {
			*hostnameInPrefix = r.Config.HostnameInPrefix.ValueBool()
		} else {
			hostnameInPrefix = nil
		}
		var metrics []shared.StatsdAdvancedPluginMetrics = []shared.StatsdAdvancedPluginMetrics{}
		for _, metricsItem := range r.Config.Metrics {
			consumerIdentifier := new(shared.StatsdAdvancedPluginConsumerIdentifier)
			if !metricsItem.ConsumerIdentifier.IsUnknown() && !metricsItem.ConsumerIdentifier.IsNull() {
				*consumerIdentifier = shared.StatsdAdvancedPluginConsumerIdentifier(metricsItem.ConsumerIdentifier.ValueString())
			} else {
				consumerIdentifier = nil
			}
			name := shared.StatsdAdvancedPluginName(metricsItem.Name.ValueString())
			sampleRate := new(float64)
			if !metricsItem.SampleRate.IsUnknown() && !metricsItem.SampleRate.IsNull() {
				*sampleRate = metricsItem.SampleRate.ValueFloat64()
			} else {
				sampleRate = nil
			}
			serviceIdentifier := new(shared.StatsdAdvancedPluginServiceIdentifier)
			if !metricsItem.ServiceIdentifier.IsUnknown() && !metricsItem.ServiceIdentifier.IsNull() {
				*serviceIdentifier = shared.StatsdAdvancedPluginServiceIdentifier(metricsItem.ServiceIdentifier.ValueString())
			} else {
				serviceIdentifier = nil
			}
			statType := shared.StatsdAdvancedPluginStatType(metricsItem.StatType.ValueString())
			workspaceIdentifier := new(shared.StatsdAdvancedPluginWorkspaceIdentifier)
			if !metricsItem.WorkspaceIdentifier.IsUnknown() && !metricsItem.WorkspaceIdentifier.IsNull() {
				*workspaceIdentifier = shared.StatsdAdvancedPluginWorkspaceIdentifier(metricsItem.WorkspaceIdentifier.ValueString())
			} else {
				workspaceIdentifier = nil
			}
			metrics = append(metrics, shared.StatsdAdvancedPluginMetrics{
				ConsumerIdentifier:  consumerIdentifier,
				Name:                name,
				SampleRate:          sampleRate,
				ServiceIdentifier:   serviceIdentifier,
				StatType:            statType,
				WorkspaceIdentifier: workspaceIdentifier,
			})
		}
		port := new(int64)
		if !r.Config.Port.IsUnknown() && !r.Config.Port.IsNull() {
			*port = r.Config.Port.ValueInt64()
		} else {
			port = nil
		}
		prefix := new(string)
		if !r.Config.Prefix.IsUnknown() && !r.Config.Prefix.IsNull() {
			*prefix = r.Config.Prefix.ValueString()
		} else {
			prefix = nil
		}
		var queue *shared.StatsdAdvancedPluginQueue
		if r.Config.Queue != nil {
			concurrencyLimit := new(shared.StatsdAdvancedPluginConcurrencyLimit)
			if !r.Config.Queue.ConcurrencyLimit.IsUnknown() && !r.Config.Queue.ConcurrencyLimit.IsNull() {
				*concurrencyLimit = shared.StatsdAdvancedPluginConcurrencyLimit(r.Config.Queue.ConcurrencyLimit.ValueInt64())
			} else {
				concurrencyLimit = nil
			}
			initialRetryDelay := new(float64)
			if !r.Config.Queue.InitialRetryDelay.IsUnknown() && !r.Config.Queue.InitialRetryDelay.IsNull() {
				*initialRetryDelay = r.Config.Queue.InitialRetryDelay.ValueFloat64()
			} else {
				initialRetryDelay = nil
			}
			maxBatchSize := new(int64)
			if !r.Config.Queue.MaxBatchSize.IsUnknown() && !r.Config.Queue.MaxBatchSize.IsNull() {
				*maxBatchSize = r.Config.Queue.MaxBatchSize.ValueInt64()
			} else {
				maxBatchSize = nil
			}
			maxBytes := new(int64)
			if !r.Config.Queue.MaxBytes.IsUnknown() && !r.Config.Queue.MaxBytes.IsNull() {
				*maxBytes = r.Config.Queue.MaxBytes.ValueInt64()
			} else {
				maxBytes = nil
			}
			maxCoalescingDelay := new(float64)
			if !r.Config.Queue.MaxCoalescingDelay.IsUnknown() && !r.Config.Queue.MaxCoalescingDelay.IsNull() {
				*maxCoalescingDelay = r.Config.Queue.MaxCoalescingDelay.ValueFloat64()
			} else {
				maxCoalescingDelay = nil
			}
			maxEntries := new(int64)
			if !r.Config.Queue.MaxEntries.IsUnknown() && !r.Config.Queue.MaxEntries.IsNull() {
				*maxEntries = r.Config.Queue.MaxEntries.ValueInt64()
			} else {
				maxEntries = nil
			}
			maxRetryDelay := new(float64)
			if !r.Config.Queue.MaxRetryDelay.IsUnknown() && !r.Config.Queue.MaxRetryDelay.IsNull() {
				*maxRetryDelay = r.Config.Queue.MaxRetryDelay.ValueFloat64()
			} else {
				maxRetryDelay = nil
			}
			maxRetryTime := new(float64)
			if !r.Config.Queue.MaxRetryTime.IsUnknown() && !r.Config.Queue.MaxRetryTime.IsNull() {
				*maxRetryTime = r.Config.Queue.MaxRetryTime.ValueFloat64()
			} else {
				maxRetryTime = nil
			}
			queue = &shared.StatsdAdvancedPluginQueue{
				ConcurrencyLimit:   concurrencyLimit,
				InitialRetryDelay:  initialRetryDelay,
				MaxBatchSize:       maxBatchSize,
				MaxBytes:           maxBytes,
				MaxCoalescingDelay: maxCoalescingDelay,
				MaxEntries:         maxEntries,
				MaxRetryDelay:      maxRetryDelay,
				MaxRetryTime:       maxRetryTime,
			}
		}
		serviceIdentifierDefault := new(shared.StatsdAdvancedPluginServiceIdentifierDefault)
		if !r.Config.ServiceIdentifierDefault.IsUnknown() && !r.Config.ServiceIdentifierDefault.IsNull() {
			*serviceIdentifierDefault = shared.StatsdAdvancedPluginServiceIdentifierDefault(r.Config.ServiceIdentifierDefault.ValueString())
		} else {
			serviceIdentifierDefault = nil
		}
		udpPacketSize := new(float64)
		if !r.Config.UDPPacketSize.IsUnknown() && !r.Config.UDPPacketSize.IsNull() {
			*udpPacketSize = r.Config.UDPPacketSize.ValueFloat64()
		} else {
			udpPacketSize = nil
		}
		useTCP := new(bool)
		if !r.Config.UseTCP.IsUnknown() && !r.Config.UseTCP.IsNull() {
			*useTCP = r.Config.UseTCP.ValueBool()
		} else {
			useTCP = nil
		}
		workspaceIdentifierDefault := new(shared.StatsdAdvancedPluginWorkspaceIdentifierDefault)
		if !r.Config.WorkspaceIdentifierDefault.IsUnknown() && !r.Config.WorkspaceIdentifierDefault.IsNull() {
			*workspaceIdentifierDefault = shared.StatsdAdvancedPluginWorkspaceIdentifierDefault(r.Config.WorkspaceIdentifierDefault.ValueString())
		} else {
			workspaceIdentifierDefault = nil
		}
		config = &shared.StatsdAdvancedPluginConfig{
			AllowStatusCodes:           allowStatusCodes,
			ConsumerIdentifierDefault:  consumerIdentifierDefault,
			Host:                       host,
			HostnameInPrefix:           hostnameInPrefix,
			Metrics:                    metrics,
			Port:                       port,
			Prefix:                     prefix,
			Queue:                      queue,
			ServiceIdentifierDefault:   serviceIdentifierDefault,
			UDPPacketSize:              udpPacketSize,
			UseTCP:                     useTCP,
			WorkspaceIdentifierDefault: workspaceIdentifierDefault,
		}
	}
	var consumer *shared.StatsdAdvancedPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.StatsdAdvancedPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.StatsdAdvancedPluginProtocols = []shared.StatsdAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.StatsdAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.StatsdAdvancedPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.StatsdAdvancedPluginRoute{
			ID: id2,
		}
	}
	var service *shared.StatsdAdvancedPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.StatsdAdvancedPluginService{
			ID: id3,
		}
	}
	out := shared.StatsdAdvancedPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginStatsdAdvancedResourceModel) RefreshFromSharedStatsdAdvancedPlugin(ctx context.Context, resp *shared.StatsdAdvancedPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.StatsdAdvancedPluginConfig{}
			r.Config.AllowStatusCodes = make([]types.String, 0, len(resp.Config.AllowStatusCodes))
			for _, v := range resp.Config.AllowStatusCodes {
				r.Config.AllowStatusCodes = append(r.Config.AllowStatusCodes, types.StringValue(v))
			}
			if resp.Config.ConsumerIdentifierDefault != nil {
				r.Config.ConsumerIdentifierDefault = types.StringValue(string(*resp.Config.ConsumerIdentifierDefault))
			} else {
				r.Config.ConsumerIdentifierDefault = types.StringNull()
			}
			r.Config.Host = types.StringPointerValue(resp.Config.Host)
			r.Config.HostnameInPrefix = types.BoolPointerValue(resp.Config.HostnameInPrefix)
			r.Config.Metrics = []tfTypes.StatsdPluginMetrics{}
			if len(r.Config.Metrics) > len(resp.Config.Metrics) {
				r.Config.Metrics = r.Config.Metrics[:len(resp.Config.Metrics)]
			}
			for metricsCount, metricsItem := range resp.Config.Metrics {
				var metrics tfTypes.StatsdPluginMetrics
				if metricsItem.ConsumerIdentifier != nil {
					metrics.ConsumerIdentifier = types.StringValue(string(*metricsItem.ConsumerIdentifier))
				} else {
					metrics.ConsumerIdentifier = types.StringNull()
				}
				metrics.Name = types.StringValue(string(metricsItem.Name))
				metrics.SampleRate = types.Float64PointerValue(metricsItem.SampleRate)
				if metricsItem.ServiceIdentifier != nil {
					metrics.ServiceIdentifier = types.StringValue(string(*metricsItem.ServiceIdentifier))
				} else {
					metrics.ServiceIdentifier = types.StringNull()
				}
				metrics.StatType = types.StringValue(string(metricsItem.StatType))
				if metricsItem.WorkspaceIdentifier != nil {
					metrics.WorkspaceIdentifier = types.StringValue(string(*metricsItem.WorkspaceIdentifier))
				} else {
					metrics.WorkspaceIdentifier = types.StringNull()
				}
				if metricsCount+1 > len(r.Config.Metrics) {
					r.Config.Metrics = append(r.Config.Metrics, metrics)
				} else {
					r.Config.Metrics[metricsCount].ConsumerIdentifier = metrics.ConsumerIdentifier
					r.Config.Metrics[metricsCount].Name = metrics.Name
					r.Config.Metrics[metricsCount].SampleRate = metrics.SampleRate
					r.Config.Metrics[metricsCount].ServiceIdentifier = metrics.ServiceIdentifier
					r.Config.Metrics[metricsCount].StatType = metrics.StatType
					r.Config.Metrics[metricsCount].WorkspaceIdentifier = metrics.WorkspaceIdentifier
				}
			}
			r.Config.Port = types.Int64PointerValue(resp.Config.Port)
			r.Config.Prefix = types.StringPointerValue(resp.Config.Prefix)
			if resp.Config.Queue == nil {
				r.Config.Queue = nil
			} else {
				r.Config.Queue = &tfTypes.Queue{}
				if resp.Config.Queue.ConcurrencyLimit != nil {
					r.Config.Queue.ConcurrencyLimit = types.Int64Value(int64(*resp.Config.Queue.ConcurrencyLimit))
				} else {
					r.Config.Queue.ConcurrencyLimit = types.Int64Null()
				}
				r.Config.Queue.InitialRetryDelay = types.Float64PointerValue(resp.Config.Queue.InitialRetryDelay)
				r.Config.Queue.MaxBatchSize = types.Int64PointerValue(resp.Config.Queue.MaxBatchSize)
				r.Config.Queue.MaxBytes = types.Int64PointerValue(resp.Config.Queue.MaxBytes)
				r.Config.Queue.MaxCoalescingDelay = types.Float64PointerValue(resp.Config.Queue.MaxCoalescingDelay)
				r.Config.Queue.MaxEntries = types.Int64PointerValue(resp.Config.Queue.MaxEntries)
				r.Config.Queue.MaxRetryDelay = types.Float64PointerValue(resp.Config.Queue.MaxRetryDelay)
				r.Config.Queue.MaxRetryTime = types.Float64PointerValue(resp.Config.Queue.MaxRetryTime)
			}
			if resp.Config.ServiceIdentifierDefault != nil {
				r.Config.ServiceIdentifierDefault = types.StringValue(string(*resp.Config.ServiceIdentifierDefault))
			} else {
				r.Config.ServiceIdentifierDefault = types.StringNull()
			}
			r.Config.UDPPacketSize = types.Float64PointerValue(resp.Config.UDPPacketSize)
			r.Config.UseTCP = types.BoolPointerValue(resp.Config.UseTCP)
			if resp.Config.WorkspaceIdentifierDefault != nil {
				r.Config.WorkspaceIdentifierDefault = types.StringValue(string(*resp.Config.WorkspaceIdentifierDefault))
			} else {
				r.Config.WorkspaceIdentifierDefault = types.StringNull()
			}
		}
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

	return diags
}
