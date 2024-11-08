// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginKafkaLogResourceModel) ToSharedKafkaLogPluginInput() *shared.KafkaLogPluginInput {
	var authentication *shared.Authentication
	if r.Config.Authentication != nil {
		mechanism := new(shared.Mechanism)
		if !r.Config.Authentication.Mechanism.IsUnknown() && !r.Config.Authentication.Mechanism.IsNull() {
			*mechanism = shared.Mechanism(r.Config.Authentication.Mechanism.ValueString())
		} else {
			mechanism = nil
		}
		password := new(string)
		if !r.Config.Authentication.Password.IsUnknown() && !r.Config.Authentication.Password.IsNull() {
			*password = r.Config.Authentication.Password.ValueString()
		} else {
			password = nil
		}
		strategy := new(shared.KafkaLogPluginStrategy)
		if !r.Config.Authentication.Strategy.IsUnknown() && !r.Config.Authentication.Strategy.IsNull() {
			*strategy = shared.KafkaLogPluginStrategy(r.Config.Authentication.Strategy.ValueString())
		} else {
			strategy = nil
		}
		tokenauth := new(bool)
		if !r.Config.Authentication.Tokenauth.IsUnknown() && !r.Config.Authentication.Tokenauth.IsNull() {
			*tokenauth = r.Config.Authentication.Tokenauth.ValueBool()
		} else {
			tokenauth = nil
		}
		user := new(string)
		if !r.Config.Authentication.User.IsUnknown() && !r.Config.Authentication.User.IsNull() {
			*user = r.Config.Authentication.User.ValueString()
		} else {
			user = nil
		}
		authentication = &shared.Authentication{
			Mechanism: mechanism,
			Password:  password,
			Strategy:  strategy,
			Tokenauth: tokenauth,
			User:      user,
		}
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
	clusterName := new(string)
	if !r.Config.ClusterName.IsUnknown() && !r.Config.ClusterName.IsNull() {
		*clusterName = r.Config.ClusterName.ValueString()
	} else {
		clusterName = nil
	}
	customFieldsByLua := make(map[string]interface{})
	for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
		var customFieldsByLuaInst interface{}
		_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
		customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
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
	var security *shared.KafkaLogPluginSecurity
	if r.Config.Security != nil {
		certificateID := new(string)
		if !r.Config.Security.CertificateID.IsUnknown() && !r.Config.Security.CertificateID.IsNull() {
			*certificateID = r.Config.Security.CertificateID.ValueString()
		} else {
			certificateID = nil
		}
		ssl := new(bool)
		if !r.Config.Security.Ssl.IsUnknown() && !r.Config.Security.Ssl.IsNull() {
			*ssl = r.Config.Security.Ssl.ValueBool()
		} else {
			ssl = nil
		}
		security = &shared.KafkaLogPluginSecurity{
			CertificateID: certificateID,
			Ssl:           ssl,
		}
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
	config := shared.KafkaLogPluginConfig{
		Authentication:    authentication,
		BootstrapServers:  bootstrapServers,
		ClusterName:       clusterName,
		CustomFieldsByLua: customFieldsByLua,
		Keepalive:         keepalive,
		KeepaliveEnabled:  keepaliveEnabled,
		ProducerAsync:     producerAsync,
		ProducerAsyncBufferingLimitsMessagesInMemory: producerAsyncBufferingLimitsMessagesInMemory,
		ProducerAsyncFlushTimeout:                    producerAsyncFlushTimeout,
		ProducerRequestAcks:                          producerRequestAcks,
		ProducerRequestLimitsBytesPerRequest:         producerRequestLimitsBytesPerRequest,
		ProducerRequestLimitsMessagesPerRequest:      producerRequestLimitsMessagesPerRequest,
		ProducerRequestRetriesBackoffTimeout:         producerRequestRetriesBackoffTimeout,
		ProducerRequestRetriesMaxAttempts:            producerRequestRetriesMaxAttempts,
		ProducerRequestTimeout:                       producerRequestTimeout,
		Security:                                     security,
		Timeout:                                      timeout,
		Topic:                                        topic,
	}
	var consumer *shared.KafkaLogPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.KafkaLogPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.KafkaLogPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.KafkaLogPluginConsumerGroup{
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
	var ordering *shared.KafkaLogPluginOrdering
	if r.Ordering != nil {
		var after *shared.KafkaLogPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.KafkaLogPluginAfter{
				Access: access,
			}
		}
		var before *shared.KafkaLogPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.KafkaLogPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.KafkaLogPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.KafkaLogPluginProtocols = []shared.KafkaLogPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.KafkaLogPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.KafkaLogPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.KafkaLogPluginRoute{
			ID: id3,
		}
	}
	var service *shared.KafkaLogPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.KafkaLogPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.KafkaLogPluginInput{
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

func (r *GatewayPluginKafkaLogResourceModel) RefreshFromSharedKafkaLogPlugin(resp *shared.KafkaLogPlugin) {
	if resp != nil {
		if resp.Config.Authentication == nil {
			r.Config.Authentication = nil
		} else {
			r.Config.Authentication = &tfTypes.Authentication{}
			if resp.Config.Authentication.Mechanism != nil {
				r.Config.Authentication.Mechanism = types.StringValue(string(*resp.Config.Authentication.Mechanism))
			} else {
				r.Config.Authentication.Mechanism = types.StringNull()
			}
			r.Config.Authentication.Password = types.StringPointerValue(resp.Config.Authentication.Password)
			if resp.Config.Authentication.Strategy != nil {
				r.Config.Authentication.Strategy = types.StringValue(string(*resp.Config.Authentication.Strategy))
			} else {
				r.Config.Authentication.Strategy = types.StringNull()
			}
			r.Config.Authentication.Tokenauth = types.BoolPointerValue(resp.Config.Authentication.Tokenauth)
			r.Config.Authentication.User = types.StringPointerValue(resp.Config.Authentication.User)
		}
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
		r.Config.ClusterName = types.StringPointerValue(resp.Config.ClusterName)
		if len(resp.Config.CustomFieldsByLua) > 0 {
			r.Config.CustomFieldsByLua = make(map[string]types.String)
			for key, value := range resp.Config.CustomFieldsByLua {
				result, _ := json.Marshal(value)
				r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
			}
		}
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
		if resp.Config.Security == nil {
			r.Config.Security = nil
		} else {
			r.Config.Security = &tfTypes.KafkaLogPluginSecurity{}
			r.Config.Security.CertificateID = types.StringPointerValue(resp.Config.Security.CertificateID)
			r.Config.Security.Ssl = types.BoolPointerValue(resp.Config.Security.Ssl)
		}
		r.Config.Timeout = types.Int64PointerValue(resp.Config.Timeout)
		r.Config.Topic = types.StringPointerValue(resp.Config.Topic)
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
