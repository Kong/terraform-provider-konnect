// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginKafkaLogDataSourceModel) RefreshFromSharedKafkaLogPlugin(resp *shared.KafkaLogPlugin) {
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
