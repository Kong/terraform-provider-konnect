// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginGraphqlProxyCacheAdvancedDataSourceModel) RefreshFromSharedGraphqlProxyCacheAdvancedPlugin(resp *shared.GraphqlProxyCacheAdvancedPlugin) {
	if resp != nil {
		r.Config.BypassOnErr = types.BoolPointerValue(resp.Config.BypassOnErr)
		r.Config.CacheTTL = types.Int64PointerValue(resp.Config.CacheTTL)
		if resp.Config.Memory == nil {
			r.Config.Memory = nil
		} else {
			r.Config.Memory = &tfTypes.Memory{}
			r.Config.Memory.DictionaryName = types.StringPointerValue(resp.Config.Memory.DictionaryName)
		}
		if resp.Config.Redis == nil {
			r.Config.Redis = nil
		} else {
			r.Config.Redis = &tfTypes.Redis{}
			r.Config.Redis.ClusterMaxRedirections = types.Int64PointerValue(resp.Config.Redis.ClusterMaxRedirections)
			r.Config.Redis.ClusterNodes = []tfTypes.ClusterNodes{}
			if len(r.Config.Redis.ClusterNodes) > len(resp.Config.Redis.ClusterNodes) {
				r.Config.Redis.ClusterNodes = r.Config.Redis.ClusterNodes[:len(resp.Config.Redis.ClusterNodes)]
			}
			for clusterNodesCount, clusterNodesItem := range resp.Config.Redis.ClusterNodes {
				var clusterNodes1 tfTypes.ClusterNodes
				clusterNodes1.IP = types.StringPointerValue(clusterNodesItem.IP)
				clusterNodes1.Port = types.Int64PointerValue(clusterNodesItem.Port)
				if clusterNodesCount+1 > len(r.Config.Redis.ClusterNodes) {
					r.Config.Redis.ClusterNodes = append(r.Config.Redis.ClusterNodes, clusterNodes1)
				} else {
					r.Config.Redis.ClusterNodes[clusterNodesCount].IP = clusterNodes1.IP
					r.Config.Redis.ClusterNodes[clusterNodesCount].Port = clusterNodes1.Port
				}
			}
			r.Config.Redis.ConnectTimeout = types.Int64PointerValue(resp.Config.Redis.ConnectTimeout)
			r.Config.Redis.ConnectionIsProxied = types.BoolPointerValue(resp.Config.Redis.ConnectionIsProxied)
			r.Config.Redis.Database = types.Int64PointerValue(resp.Config.Redis.Database)
			r.Config.Redis.Host = types.StringPointerValue(resp.Config.Redis.Host)
			r.Config.Redis.KeepaliveBacklog = types.Int64PointerValue(resp.Config.Redis.KeepaliveBacklog)
			r.Config.Redis.KeepalivePoolSize = types.Int64PointerValue(resp.Config.Redis.KeepalivePoolSize)
			r.Config.Redis.Password = types.StringPointerValue(resp.Config.Redis.Password)
			r.Config.Redis.Port = types.Int64PointerValue(resp.Config.Redis.Port)
			r.Config.Redis.ReadTimeout = types.Int64PointerValue(resp.Config.Redis.ReadTimeout)
			r.Config.Redis.SendTimeout = types.Int64PointerValue(resp.Config.Redis.SendTimeout)
			r.Config.Redis.SentinelMaster = types.StringPointerValue(resp.Config.Redis.SentinelMaster)
			r.Config.Redis.SentinelNodes = []tfTypes.SentinelNodes{}
			if len(r.Config.Redis.SentinelNodes) > len(resp.Config.Redis.SentinelNodes) {
				r.Config.Redis.SentinelNodes = r.Config.Redis.SentinelNodes[:len(resp.Config.Redis.SentinelNodes)]
			}
			for sentinelNodesCount, sentinelNodesItem := range resp.Config.Redis.SentinelNodes {
				var sentinelNodes1 tfTypes.SentinelNodes
				sentinelNodes1.Host = types.StringPointerValue(sentinelNodesItem.Host)
				sentinelNodes1.Port = types.Int64PointerValue(sentinelNodesItem.Port)
				if sentinelNodesCount+1 > len(r.Config.Redis.SentinelNodes) {
					r.Config.Redis.SentinelNodes = append(r.Config.Redis.SentinelNodes, sentinelNodes1)
				} else {
					r.Config.Redis.SentinelNodes[sentinelNodesCount].Host = sentinelNodes1.Host
					r.Config.Redis.SentinelNodes[sentinelNodesCount].Port = sentinelNodes1.Port
				}
			}
			r.Config.Redis.SentinelPassword = types.StringPointerValue(resp.Config.Redis.SentinelPassword)
			if resp.Config.Redis.SentinelRole != nil {
				r.Config.Redis.SentinelRole = types.StringValue(string(*resp.Config.Redis.SentinelRole))
			} else {
				r.Config.Redis.SentinelRole = types.StringNull()
			}
			r.Config.Redis.SentinelUsername = types.StringPointerValue(resp.Config.Redis.SentinelUsername)
			r.Config.Redis.ServerName = types.StringPointerValue(resp.Config.Redis.ServerName)
			r.Config.Redis.Ssl = types.BoolPointerValue(resp.Config.Redis.Ssl)
			r.Config.Redis.SslVerify = types.BoolPointerValue(resp.Config.Redis.SslVerify)
			r.Config.Redis.Username = types.StringPointerValue(resp.Config.Redis.Username)
		}
		if resp.Config.Strategy != nil {
			r.Config.Strategy = types.StringValue(string(*resp.Config.Strategy))
		} else {
			r.Config.Strategy = types.StringNull()
		}
		r.Config.VaryHeaders = []types.String{}
		for _, v := range resp.Config.VaryHeaders {
			r.Config.VaryHeaders = append(r.Config.VaryHeaders, types.StringValue(v))
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