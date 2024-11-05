// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginKafkaLogResourceModel) ToSharedCreateKafkaLogPlugin() *shared.CreateKafkaLogPlugin {
	var config *shared.CreateKafkaLogPluginConfig
	if r.Config != nil {
		var authentication *shared.CreateKafkaLogPluginAuthentication
		if r.Config.Authentication != nil {
			mechanism := new(shared.CreateKafkaLogPluginMechanism)
			if !r.Config.Authentication.Mechanism.IsUnknown() && !r.Config.Authentication.Mechanism.IsNull() {
				*mechanism = shared.CreateKafkaLogPluginMechanism(r.Config.Authentication.Mechanism.ValueString())
			} else {
				mechanism = nil
			}
			password := new(string)
			if !r.Config.Authentication.Password.IsUnknown() && !r.Config.Authentication.Password.IsNull() {
				*password = r.Config.Authentication.Password.ValueString()
			} else {
				password = nil
			}
			strategy := new(shared.CreateKafkaLogPluginStrategy)
			if !r.Config.Authentication.Strategy.IsUnknown() && !r.Config.Authentication.Strategy.IsNull() {
				*strategy = shared.CreateKafkaLogPluginStrategy(r.Config.Authentication.Strategy.ValueString())
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
			authentication = &shared.CreateKafkaLogPluginAuthentication{
				Mechanism: mechanism,
				Password:  password,
				Strategy:  strategy,
				Tokenauth: tokenauth,
				User:      user,
			}
		}
		var bootstrapServers []shared.CreateKafkaLogPluginBootstrapServers = []shared.CreateKafkaLogPluginBootstrapServers{}
		for _, bootstrapServersItem := range r.Config.BootstrapServers {
			var host string
			host = bootstrapServersItem.Host.ValueString()

			var port int64
			port = bootstrapServersItem.Port.ValueInt64()

			bootstrapServers = append(bootstrapServers, shared.CreateKafkaLogPluginBootstrapServers{
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
		producerRequestAcks := new(shared.CreateKafkaLogPluginProducerRequestAcks)
		if !r.Config.ProducerRequestAcks.IsUnknown() && !r.Config.ProducerRequestAcks.IsNull() {
			*producerRequestAcks = shared.CreateKafkaLogPluginProducerRequestAcks(r.Config.ProducerRequestAcks.ValueInt64())
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
		var security *shared.CreateKafkaLogPluginSecurity
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
			security = &shared.CreateKafkaLogPluginSecurity{
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
		config = &shared.CreateKafkaLogPluginConfig{
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
	var ordering *shared.CreateKafkaLogPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateKafkaLogPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateKafkaLogPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateKafkaLogPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateKafkaLogPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateKafkaLogPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateKafkaLogPluginProtocols = []shared.CreateKafkaLogPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateKafkaLogPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateKafkaLogPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateKafkaLogPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateKafkaLogPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateKafkaLogPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateKafkaLogPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateKafkaLogPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateKafkaLogPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateKafkaLogPluginService{
			ID: id3,
		}
	}
	out := shared.CreateKafkaLogPlugin{
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

func (r *GatewayPluginKafkaLogResourceModel) RefreshFromSharedKafkaLogPlugin(resp *shared.KafkaLogPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateKafkaLogPluginConfig{}
			if resp.Config.Authentication == nil {
				r.Config.Authentication = nil
			} else {
				r.Config.Authentication = &tfTypes.CreateKafkaLogPluginAuthentication{}
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
				r.Config.Security = &tfTypes.CreateKafkaLogPluginSecurity{}
				r.Config.Security.CertificateID = types.StringPointerValue(resp.Config.Security.CertificateID)
				r.Config.Security.Ssl = types.BoolPointerValue(resp.Config.Security.Ssl)
			}
			r.Config.Timeout = types.Int64PointerValue(resp.Config.Timeout)
			r.Config.Topic = types.StringPointerValue(resp.Config.Topic)
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