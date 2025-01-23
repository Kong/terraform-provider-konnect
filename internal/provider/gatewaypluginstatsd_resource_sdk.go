// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginStatsdResourceModel) ToSharedStatsdPluginInput() *shared.StatsdPluginInput {
	var allowStatusCodes []string = []string{}
	for _, allowStatusCodesItem := range r.Config.AllowStatusCodes {
		allowStatusCodes = append(allowStatusCodes, allowStatusCodesItem.ValueString())
	}
	consumerIdentifierDefault := new(shared.ConsumerIdentifierDefault)
	if !r.Config.ConsumerIdentifierDefault.IsUnknown() && !r.Config.ConsumerIdentifierDefault.IsNull() {
		*consumerIdentifierDefault = shared.ConsumerIdentifierDefault(r.Config.ConsumerIdentifierDefault.ValueString())
	} else {
		consumerIdentifierDefault = nil
	}
	flushTimeout := new(float64)
	if !r.Config.FlushTimeout.IsUnknown() && !r.Config.FlushTimeout.IsNull() {
		*flushTimeout, _ = r.Config.FlushTimeout.ValueBigFloat().Float64()
	} else {
		flushTimeout = nil
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
	var metrics []shared.StatsdPluginMetrics = []shared.StatsdPluginMetrics{}
	for _, metricsItem := range r.Config.Metrics {
		consumerIdentifier := new(shared.StatsdPluginConsumerIdentifier)
		if !metricsItem.ConsumerIdentifier.IsUnknown() && !metricsItem.ConsumerIdentifier.IsNull() {
			*consumerIdentifier = shared.StatsdPluginConsumerIdentifier(metricsItem.ConsumerIdentifier.ValueString())
		} else {
			consumerIdentifier = nil
		}
		name := shared.StatsdPluginName(metricsItem.Name.ValueString())
		sampleRate := new(float64)
		if !metricsItem.SampleRate.IsUnknown() && !metricsItem.SampleRate.IsNull() {
			*sampleRate, _ = metricsItem.SampleRate.ValueBigFloat().Float64()
		} else {
			sampleRate = nil
		}
		serviceIdentifier := new(shared.ServiceIdentifier)
		if !metricsItem.ServiceIdentifier.IsUnknown() && !metricsItem.ServiceIdentifier.IsNull() {
			*serviceIdentifier = shared.ServiceIdentifier(metricsItem.ServiceIdentifier.ValueString())
		} else {
			serviceIdentifier = nil
		}
		statType := shared.StatsdPluginStatType(metricsItem.StatType.ValueString())
		workspaceIdentifier := new(shared.WorkspaceIdentifier)
		if !metricsItem.WorkspaceIdentifier.IsUnknown() && !metricsItem.WorkspaceIdentifier.IsNull() {
			*workspaceIdentifier = shared.WorkspaceIdentifier(metricsItem.WorkspaceIdentifier.ValueString())
		} else {
			workspaceIdentifier = nil
		}
		metrics = append(metrics, shared.StatsdPluginMetrics{
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
	var queue *shared.StatsdPluginQueue
	if r.Config.Queue != nil {
		concurrencyLimit := new(shared.StatsdPluginConcurrencyLimit)
		if !r.Config.Queue.ConcurrencyLimit.IsUnknown() && !r.Config.Queue.ConcurrencyLimit.IsNull() {
			*concurrencyLimit = shared.StatsdPluginConcurrencyLimit(r.Config.Queue.ConcurrencyLimit.ValueInt64())
		} else {
			concurrencyLimit = nil
		}
		initialRetryDelay := new(float64)
		if !r.Config.Queue.InitialRetryDelay.IsUnknown() && !r.Config.Queue.InitialRetryDelay.IsNull() {
			*initialRetryDelay, _ = r.Config.Queue.InitialRetryDelay.ValueBigFloat().Float64()
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
			*maxCoalescingDelay, _ = r.Config.Queue.MaxCoalescingDelay.ValueBigFloat().Float64()
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
			*maxRetryDelay, _ = r.Config.Queue.MaxRetryDelay.ValueBigFloat().Float64()
		} else {
			maxRetryDelay = nil
		}
		maxRetryTime := new(float64)
		if !r.Config.Queue.MaxRetryTime.IsUnknown() && !r.Config.Queue.MaxRetryTime.IsNull() {
			*maxRetryTime, _ = r.Config.Queue.MaxRetryTime.ValueBigFloat().Float64()
		} else {
			maxRetryTime = nil
		}
		queue = &shared.StatsdPluginQueue{
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
	queueSize := new(int64)
	if !r.Config.QueueSize.IsUnknown() && !r.Config.QueueSize.IsNull() {
		*queueSize = r.Config.QueueSize.ValueInt64()
	} else {
		queueSize = nil
	}
	retryCount := new(int64)
	if !r.Config.RetryCount.IsUnknown() && !r.Config.RetryCount.IsNull() {
		*retryCount = r.Config.RetryCount.ValueInt64()
	} else {
		retryCount = nil
	}
	serviceIdentifierDefault := new(shared.ServiceIdentifierDefault)
	if !r.Config.ServiceIdentifierDefault.IsUnknown() && !r.Config.ServiceIdentifierDefault.IsNull() {
		*serviceIdentifierDefault = shared.ServiceIdentifierDefault(r.Config.ServiceIdentifierDefault.ValueString())
	} else {
		serviceIdentifierDefault = nil
	}
	tagStyle := new(shared.TagStyle)
	if !r.Config.TagStyle.IsUnknown() && !r.Config.TagStyle.IsNull() {
		*tagStyle = shared.TagStyle(r.Config.TagStyle.ValueString())
	} else {
		tagStyle = nil
	}
	udpPacketSize := new(float64)
	if !r.Config.UDPPacketSize.IsUnknown() && !r.Config.UDPPacketSize.IsNull() {
		*udpPacketSize, _ = r.Config.UDPPacketSize.ValueBigFloat().Float64()
	} else {
		udpPacketSize = nil
	}
	useTCP := new(bool)
	if !r.Config.UseTCP.IsUnknown() && !r.Config.UseTCP.IsNull() {
		*useTCP = r.Config.UseTCP.ValueBool()
	} else {
		useTCP = nil
	}
	workspaceIdentifierDefault := new(shared.WorkspaceIdentifierDefault)
	if !r.Config.WorkspaceIdentifierDefault.IsUnknown() && !r.Config.WorkspaceIdentifierDefault.IsNull() {
		*workspaceIdentifierDefault = shared.WorkspaceIdentifierDefault(r.Config.WorkspaceIdentifierDefault.ValueString())
	} else {
		workspaceIdentifierDefault = nil
	}
	config := shared.StatsdPluginConfig{
		AllowStatusCodes:           allowStatusCodes,
		ConsumerIdentifierDefault:  consumerIdentifierDefault,
		FlushTimeout:               flushTimeout,
		Host:                       host,
		HostnameInPrefix:           hostnameInPrefix,
		Metrics:                    metrics,
		Port:                       port,
		Prefix:                     prefix,
		Queue:                      queue,
		QueueSize:                  queueSize,
		RetryCount:                 retryCount,
		ServiceIdentifierDefault:   serviceIdentifierDefault,
		TagStyle:                   tagStyle,
		UDPPacketSize:              udpPacketSize,
		UseTCP:                     useTCP,
		WorkspaceIdentifierDefault: workspaceIdentifierDefault,
	}
	var consumer *shared.StatsdPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.StatsdPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.StatsdPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.StatsdPluginConsumerGroup{
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
	var ordering *shared.StatsdPluginOrdering
	if r.Ordering != nil {
		var after *shared.StatsdPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.StatsdPluginAfter{
				Access: access,
			}
		}
		var before *shared.StatsdPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.StatsdPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.StatsdPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.StatsdPluginProtocols = []shared.StatsdPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.StatsdPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.StatsdPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.StatsdPluginRoute{
			ID: id3,
		}
	}
	var service *shared.StatsdPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.StatsdPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.StatsdPluginInput{
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

func (r *GatewayPluginStatsdResourceModel) RefreshFromSharedStatsdPlugin(resp *shared.StatsdPlugin) {
	if resp != nil {
		r.Config.AllowStatusCodes = []types.String{}
		for _, v := range resp.Config.AllowStatusCodes {
			r.Config.AllowStatusCodes = append(r.Config.AllowStatusCodes, types.StringValue(v))
		}
		if resp.Config.ConsumerIdentifierDefault != nil {
			r.Config.ConsumerIdentifierDefault = types.StringValue(string(*resp.Config.ConsumerIdentifierDefault))
		} else {
			r.Config.ConsumerIdentifierDefault = types.StringNull()
		}
		if resp.Config.FlushTimeout != nil {
			r.Config.FlushTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.FlushTimeout)))
		} else {
			r.Config.FlushTimeout = types.NumberNull()
		}
		r.Config.Host = types.StringPointerValue(resp.Config.Host)
		r.Config.HostnameInPrefix = types.BoolPointerValue(resp.Config.HostnameInPrefix)
		r.Config.Metrics = []tfTypes.StatsdPluginMetrics{}
		if len(r.Config.Metrics) > len(resp.Config.Metrics) {
			r.Config.Metrics = r.Config.Metrics[:len(resp.Config.Metrics)]
		}
		for metricsCount, metricsItem := range resp.Config.Metrics {
			var metrics1 tfTypes.StatsdPluginMetrics
			if metricsItem.ConsumerIdentifier != nil {
				metrics1.ConsumerIdentifier = types.StringValue(string(*metricsItem.ConsumerIdentifier))
			} else {
				metrics1.ConsumerIdentifier = types.StringNull()
			}
			metrics1.Name = types.StringValue(string(metricsItem.Name))
			if metricsItem.SampleRate != nil {
				metrics1.SampleRate = types.NumberValue(big.NewFloat(float64(*metricsItem.SampleRate)))
			} else {
				metrics1.SampleRate = types.NumberNull()
			}
			if metricsItem.ServiceIdentifier != nil {
				metrics1.ServiceIdentifier = types.StringValue(string(*metricsItem.ServiceIdentifier))
			} else {
				metrics1.ServiceIdentifier = types.StringNull()
			}
			metrics1.StatType = types.StringValue(string(metricsItem.StatType))
			if metricsItem.WorkspaceIdentifier != nil {
				metrics1.WorkspaceIdentifier = types.StringValue(string(*metricsItem.WorkspaceIdentifier))
			} else {
				metrics1.WorkspaceIdentifier = types.StringNull()
			}
			if metricsCount+1 > len(r.Config.Metrics) {
				r.Config.Metrics = append(r.Config.Metrics, metrics1)
			} else {
				r.Config.Metrics[metricsCount].ConsumerIdentifier = metrics1.ConsumerIdentifier
				r.Config.Metrics[metricsCount].Name = metrics1.Name
				r.Config.Metrics[metricsCount].SampleRate = metrics1.SampleRate
				r.Config.Metrics[metricsCount].ServiceIdentifier = metrics1.ServiceIdentifier
				r.Config.Metrics[metricsCount].StatType = metrics1.StatType
				r.Config.Metrics[metricsCount].WorkspaceIdentifier = metrics1.WorkspaceIdentifier
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
			if resp.Config.Queue.InitialRetryDelay != nil {
				r.Config.Queue.InitialRetryDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.InitialRetryDelay)))
			} else {
				r.Config.Queue.InitialRetryDelay = types.NumberNull()
			}
			r.Config.Queue.MaxBatchSize = types.Int64PointerValue(resp.Config.Queue.MaxBatchSize)
			r.Config.Queue.MaxBytes = types.Int64PointerValue(resp.Config.Queue.MaxBytes)
			if resp.Config.Queue.MaxCoalescingDelay != nil {
				r.Config.Queue.MaxCoalescingDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxCoalescingDelay)))
			} else {
				r.Config.Queue.MaxCoalescingDelay = types.NumberNull()
			}
			r.Config.Queue.MaxEntries = types.Int64PointerValue(resp.Config.Queue.MaxEntries)
			if resp.Config.Queue.MaxRetryDelay != nil {
				r.Config.Queue.MaxRetryDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxRetryDelay)))
			} else {
				r.Config.Queue.MaxRetryDelay = types.NumberNull()
			}
			if resp.Config.Queue.MaxRetryTime != nil {
				r.Config.Queue.MaxRetryTime = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxRetryTime)))
			} else {
				r.Config.Queue.MaxRetryTime = types.NumberNull()
			}
		}
		r.Config.QueueSize = types.Int64PointerValue(resp.Config.QueueSize)
		r.Config.RetryCount = types.Int64PointerValue(resp.Config.RetryCount)
		if resp.Config.ServiceIdentifierDefault != nil {
			r.Config.ServiceIdentifierDefault = types.StringValue(string(*resp.Config.ServiceIdentifierDefault))
		} else {
			r.Config.ServiceIdentifierDefault = types.StringNull()
		}
		if resp.Config.TagStyle != nil {
			r.Config.TagStyle = types.StringValue(string(*resp.Config.TagStyle))
		} else {
			r.Config.TagStyle = types.StringNull()
		}
		if resp.Config.UDPPacketSize != nil {
			r.Config.UDPPacketSize = types.NumberValue(big.NewFloat(float64(*resp.Config.UDPPacketSize)))
		} else {
			r.Config.UDPPacketSize = types.NumberNull()
		}
		r.Config.UseTCP = types.BoolPointerValue(resp.Config.UseTCP)
		if resp.Config.WorkspaceIdentifierDefault != nil {
			r.Config.WorkspaceIdentifierDefault = types.StringValue(string(*resp.Config.WorkspaceIdentifierDefault))
		} else {
			r.Config.WorkspaceIdentifierDefault = types.StringNull()
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLWithoutParentsConsumer{}
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
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
