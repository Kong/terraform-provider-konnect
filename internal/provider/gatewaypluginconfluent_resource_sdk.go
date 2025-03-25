// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginConfluentResourceModel) ToSharedConfluentPluginInput() *shared.ConfluentPluginInput {
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
	var bootstrapServers []shared.BootstrapServers = []shared.BootstrapServers{}
	for _, bootstrapServersItem := range r.Config.BootstrapServers {
		var host string
		host = bootstrapServersItem.Host.ValueString()

		var port int64
		port = bootstrapServersItem.Port.ValueInt64()

		bootstrapServers = append(bootstrapServers, shared.BootstrapServers{
			Host: host,
			Port: port,
		})
	}
	clusterAPIKey := new(string)
	if !r.Config.ClusterAPIKey.IsUnknown() && !r.Config.ClusterAPIKey.IsNull() {
		*clusterAPIKey = r.Config.ClusterAPIKey.ValueString()
	} else {
		clusterAPIKey = nil
	}
	clusterAPISecret := new(string)
	if !r.Config.ClusterAPISecret.IsUnknown() && !r.Config.ClusterAPISecret.IsNull() {
		*clusterAPISecret = r.Config.ClusterAPISecret.ValueString()
	} else {
		clusterAPISecret = nil
	}
	clusterName := new(string)
	if !r.Config.ClusterName.IsUnknown() && !r.Config.ClusterName.IsNull() {
		*clusterName = r.Config.ClusterName.ValueString()
	} else {
		clusterName = nil
	}
	confluentCloudAPIKey := new(string)
	if !r.Config.ConfluentCloudAPIKey.IsUnknown() && !r.Config.ConfluentCloudAPIKey.IsNull() {
		*confluentCloudAPIKey = r.Config.ConfluentCloudAPIKey.ValueString()
	} else {
		confluentCloudAPIKey = nil
	}
	confluentCloudAPISecret := new(string)
	if !r.Config.ConfluentCloudAPISecret.IsUnknown() && !r.Config.ConfluentCloudAPISecret.IsNull() {
		*confluentCloudAPISecret = r.Config.ConfluentCloudAPISecret.ValueString()
	} else {
		confluentCloudAPISecret = nil
	}
	forwardBody := new(bool)
	if !r.Config.ForwardBody.IsUnknown() && !r.Config.ForwardBody.IsNull() {
		*forwardBody = r.Config.ForwardBody.ValueBool()
	} else {
		forwardBody = nil
	}
	forwardHeaders := new(bool)
	if !r.Config.ForwardHeaders.IsUnknown() && !r.Config.ForwardHeaders.IsNull() {
		*forwardHeaders = r.Config.ForwardHeaders.ValueBool()
	} else {
		forwardHeaders = nil
	}
	forwardMethod := new(bool)
	if !r.Config.ForwardMethod.IsUnknown() && !r.Config.ForwardMethod.IsNull() {
		*forwardMethod = r.Config.ForwardMethod.ValueBool()
	} else {
		forwardMethod = nil
	}
	forwardURI := new(bool)
	if !r.Config.ForwardURI.IsUnknown() && !r.Config.ForwardURI.IsNull() {
		*forwardURI = r.Config.ForwardURI.ValueBool()
	} else {
		forwardURI = nil
	}
	keepalive := new(int64)
	if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
		*keepalive = r.Config.Keepalive.ValueInt64()
	} else {
		keepalive = nil
	}
	keepaliveEnabled := new(bool)
	if !r.Config.KeepaliveEnabled.IsUnknown() && !r.Config.KeepaliveEnabled.IsNull() {
		*keepaliveEnabled = r.Config.KeepaliveEnabled.ValueBool()
	} else {
		keepaliveEnabled = nil
	}
	producerAsync := new(bool)
	if !r.Config.ProducerAsync.IsUnknown() && !r.Config.ProducerAsync.IsNull() {
		*producerAsync = r.Config.ProducerAsync.ValueBool()
	} else {
		producerAsync = nil
	}
	producerAsyncBufferingLimitsMessagesInMemory := new(int64)
	if !r.Config.ProducerAsyncBufferingLimitsMessagesInMemory.IsUnknown() && !r.Config.ProducerAsyncBufferingLimitsMessagesInMemory.IsNull() {
		*producerAsyncBufferingLimitsMessagesInMemory = r.Config.ProducerAsyncBufferingLimitsMessagesInMemory.ValueInt64()
	} else {
		producerAsyncBufferingLimitsMessagesInMemory = nil
	}
	producerAsyncFlushTimeout := new(int64)
	if !r.Config.ProducerAsyncFlushTimeout.IsUnknown() && !r.Config.ProducerAsyncFlushTimeout.IsNull() {
		*producerAsyncFlushTimeout = r.Config.ProducerAsyncFlushTimeout.ValueInt64()
	} else {
		producerAsyncFlushTimeout = nil
	}
	producerRequestAcks := new(shared.ProducerRequestAcks)
	if !r.Config.ProducerRequestAcks.IsUnknown() && !r.Config.ProducerRequestAcks.IsNull() {
		*producerRequestAcks = shared.ProducerRequestAcks(r.Config.ProducerRequestAcks.ValueInt64())
	} else {
		producerRequestAcks = nil
	}
	producerRequestLimitsBytesPerRequest := new(int64)
	if !r.Config.ProducerRequestLimitsBytesPerRequest.IsUnknown() && !r.Config.ProducerRequestLimitsBytesPerRequest.IsNull() {
		*producerRequestLimitsBytesPerRequest = r.Config.ProducerRequestLimitsBytesPerRequest.ValueInt64()
	} else {
		producerRequestLimitsBytesPerRequest = nil
	}
	producerRequestLimitsMessagesPerRequest := new(int64)
	if !r.Config.ProducerRequestLimitsMessagesPerRequest.IsUnknown() && !r.Config.ProducerRequestLimitsMessagesPerRequest.IsNull() {
		*producerRequestLimitsMessagesPerRequest = r.Config.ProducerRequestLimitsMessagesPerRequest.ValueInt64()
	} else {
		producerRequestLimitsMessagesPerRequest = nil
	}
	producerRequestRetriesBackoffTimeout := new(int64)
	if !r.Config.ProducerRequestRetriesBackoffTimeout.IsUnknown() && !r.Config.ProducerRequestRetriesBackoffTimeout.IsNull() {
		*producerRequestRetriesBackoffTimeout = r.Config.ProducerRequestRetriesBackoffTimeout.ValueInt64()
	} else {
		producerRequestRetriesBackoffTimeout = nil
	}
	producerRequestRetriesMaxAttempts := new(int64)
	if !r.Config.ProducerRequestRetriesMaxAttempts.IsUnknown() && !r.Config.ProducerRequestRetriesMaxAttempts.IsNull() {
		*producerRequestRetriesMaxAttempts = r.Config.ProducerRequestRetriesMaxAttempts.ValueInt64()
	} else {
		producerRequestRetriesMaxAttempts = nil
	}
	producerRequestTimeout := new(int64)
	if !r.Config.ProducerRequestTimeout.IsUnknown() && !r.Config.ProducerRequestTimeout.IsNull() {
		*producerRequestTimeout = r.Config.ProducerRequestTimeout.ValueInt64()
	} else {
		producerRequestTimeout = nil
	}
	timeout := new(int64)
	if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
		*timeout = r.Config.Timeout.ValueInt64()
	} else {
		timeout = nil
	}
	topic := new(string)
	if !r.Config.Topic.IsUnknown() && !r.Config.Topic.IsNull() {
		*topic = r.Config.Topic.ValueString()
	} else {
		topic = nil
	}
	config := shared.ConfluentPluginConfig{
		BootstrapServers:        bootstrapServers,
		ClusterAPIKey:           clusterAPIKey,
		ClusterAPISecret:        clusterAPISecret,
		ClusterName:             clusterName,
		ConfluentCloudAPIKey:    confluentCloudAPIKey,
		ConfluentCloudAPISecret: confluentCloudAPISecret,
		ForwardBody:             forwardBody,
		ForwardHeaders:          forwardHeaders,
		ForwardMethod:           forwardMethod,
		ForwardURI:              forwardURI,
		Keepalive:               keepalive,
		KeepaliveEnabled:        keepaliveEnabled,
		ProducerAsync:           producerAsync,
		ProducerAsyncBufferingLimitsMessagesInMemory: producerAsyncBufferingLimitsMessagesInMemory,
		ProducerAsyncFlushTimeout:                    producerAsyncFlushTimeout,
		ProducerRequestAcks:                          producerRequestAcks,
		ProducerRequestLimitsBytesPerRequest:         producerRequestLimitsBytesPerRequest,
		ProducerRequestLimitsMessagesPerRequest:      producerRequestLimitsMessagesPerRequest,
		ProducerRequestRetriesBackoffTimeout:         producerRequestRetriesBackoffTimeout,
		ProducerRequestRetriesMaxAttempts:            producerRequestRetriesMaxAttempts,
		ProducerRequestTimeout:                       producerRequestTimeout,
		Timeout:                                      timeout,
		Topic:                                        topic,
	}
	var consumer *shared.ConfluentPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.ConfluentPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.ConfluentPluginProtocols = []shared.ConfluentPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ConfluentPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ConfluentPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.ConfluentPluginRoute{
			ID: id2,
		}
	}
	var service *shared.ConfluentPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.ConfluentPluginService{
			ID: id3,
		}
	}
	out := shared.ConfluentPluginInput{
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

func (r *GatewayPluginConfluentResourceModel) RefreshFromSharedConfluentPlugin(resp *shared.ConfluentPlugin) {
	if resp != nil {
		r.Config.BootstrapServers = []tfTypes.BootstrapServers{}
		if len(r.Config.BootstrapServers) > len(resp.Config.BootstrapServers) {
			r.Config.BootstrapServers = r.Config.BootstrapServers[:len(resp.Config.BootstrapServers)]
		}
		for bootstrapServersCount, bootstrapServersItem := range resp.Config.BootstrapServers {
			var bootstrapServers1 tfTypes.BootstrapServers
			bootstrapServers1.Host = types.StringValue(bootstrapServersItem.Host)
			bootstrapServers1.Port = types.Int64Value(bootstrapServersItem.Port)
			if bootstrapServersCount+1 > len(r.Config.BootstrapServers) {
				r.Config.BootstrapServers = append(r.Config.BootstrapServers, bootstrapServers1)
			} else {
				r.Config.BootstrapServers[bootstrapServersCount].Host = bootstrapServers1.Host
				r.Config.BootstrapServers[bootstrapServersCount].Port = bootstrapServers1.Port
			}
		}
		r.Config.ClusterAPIKey = types.StringPointerValue(resp.Config.ClusterAPIKey)
		r.Config.ClusterAPISecret = types.StringPointerValue(resp.Config.ClusterAPISecret)
		r.Config.ClusterName = types.StringPointerValue(resp.Config.ClusterName)
		r.Config.ConfluentCloudAPIKey = types.StringPointerValue(resp.Config.ConfluentCloudAPIKey)
		r.Config.ConfluentCloudAPISecret = types.StringPointerValue(resp.Config.ConfluentCloudAPISecret)
		r.Config.ForwardBody = types.BoolPointerValue(resp.Config.ForwardBody)
		r.Config.ForwardHeaders = types.BoolPointerValue(resp.Config.ForwardHeaders)
		r.Config.ForwardMethod = types.BoolPointerValue(resp.Config.ForwardMethod)
		r.Config.ForwardURI = types.BoolPointerValue(resp.Config.ForwardURI)
		r.Config.Keepalive = types.Int64PointerValue(resp.Config.Keepalive)
		r.Config.KeepaliveEnabled = types.BoolPointerValue(resp.Config.KeepaliveEnabled)
		r.Config.ProducerAsync = types.BoolPointerValue(resp.Config.ProducerAsync)
		r.Config.ProducerAsyncBufferingLimitsMessagesInMemory = types.Int64PointerValue(resp.Config.ProducerAsyncBufferingLimitsMessagesInMemory)
		r.Config.ProducerAsyncFlushTimeout = types.Int64PointerValue(resp.Config.ProducerAsyncFlushTimeout)
		if resp.Config.ProducerRequestAcks != nil {
			r.Config.ProducerRequestAcks = types.Int64Value(int64(*resp.Config.ProducerRequestAcks))
		} else {
			r.Config.ProducerRequestAcks = types.Int64Null()
		}
		r.Config.ProducerRequestLimitsBytesPerRequest = types.Int64PointerValue(resp.Config.ProducerRequestLimitsBytesPerRequest)
		r.Config.ProducerRequestLimitsMessagesPerRequest = types.Int64PointerValue(resp.Config.ProducerRequestLimitsMessagesPerRequest)
		r.Config.ProducerRequestRetriesBackoffTimeout = types.Int64PointerValue(resp.Config.ProducerRequestRetriesBackoffTimeout)
		r.Config.ProducerRequestRetriesMaxAttempts = types.Int64PointerValue(resp.Config.ProducerRequestRetriesMaxAttempts)
		r.Config.ProducerRequestTimeout = types.Int64PointerValue(resp.Config.ProducerRequestTimeout)
		r.Config.Timeout = types.Int64PointerValue(resp.Config.Timeout)
		r.Config.Topic = types.StringPointerValue(resp.Config.Topic)
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
