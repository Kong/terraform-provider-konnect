// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginUpstreamOauthDataSourceModel) RefreshFromSharedUpstreamOauthPlugin(resp *shared.UpstreamOauthPlugin) {
	if resp != nil {
		if resp.Config.Behavior == nil {
			r.Config.Behavior = nil
		} else {
			r.Config.Behavior = &tfTypes.Behavior{}
			r.Config.Behavior.IdpErrorResponseBodyTemplate = types.StringPointerValue(resp.Config.Behavior.IdpErrorResponseBodyTemplate)
			r.Config.Behavior.IdpErrorResponseContentType = types.StringPointerValue(resp.Config.Behavior.IdpErrorResponseContentType)
			r.Config.Behavior.IdpErrorResponseMessage = types.StringPointerValue(resp.Config.Behavior.IdpErrorResponseMessage)
			r.Config.Behavior.IdpErrorResponseStatusCode = types.Int64PointerValue(resp.Config.Behavior.IdpErrorResponseStatusCode)
			r.Config.Behavior.PurgeTokenOnUpstreamStatusCodes = make([]types.Int64, 0, len(resp.Config.Behavior.PurgeTokenOnUpstreamStatusCodes))
			for _, v := range resp.Config.Behavior.PurgeTokenOnUpstreamStatusCodes {
				r.Config.Behavior.PurgeTokenOnUpstreamStatusCodes = append(r.Config.Behavior.PurgeTokenOnUpstreamStatusCodes, types.Int64Value(v))
			}
			r.Config.Behavior.UpstreamAccessTokenHeaderName = types.StringPointerValue(resp.Config.Behavior.UpstreamAccessTokenHeaderName)
		}
		if resp.Config.Cache == nil {
			r.Config.Cache = nil
		} else {
			r.Config.Cache = &tfTypes.Cache{}
			if resp.Config.Cache.DefaultTTL != nil {
				r.Config.Cache.DefaultTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.Cache.DefaultTTL)))
			} else {
				r.Config.Cache.DefaultTTL = types.NumberNull()
			}
			r.Config.Cache.EagerlyExpire = types.Int64PointerValue(resp.Config.Cache.EagerlyExpire)
			if resp.Config.Cache.Memory == nil {
				r.Config.Cache.Memory = nil
			} else {
				r.Config.Cache.Memory = &tfTypes.Memory{}
				r.Config.Cache.Memory.DictionaryName = types.StringPointerValue(resp.Config.Cache.Memory.DictionaryName)
			}
			if resp.Config.Cache.Redis == nil {
				r.Config.Cache.Redis = nil
			} else {
				r.Config.Cache.Redis = &tfTypes.AiProxyAdvancedPluginRedis{}
				r.Config.Cache.Redis.ClusterMaxRedirections = types.Int64PointerValue(resp.Config.Cache.Redis.ClusterMaxRedirections)
				r.Config.Cache.Redis.ClusterNodes = []tfTypes.AiProxyAdvancedPluginClusterNodes{}
				if len(r.Config.Cache.Redis.ClusterNodes) > len(resp.Config.Cache.Redis.ClusterNodes) {
					r.Config.Cache.Redis.ClusterNodes = r.Config.Cache.Redis.ClusterNodes[:len(resp.Config.Cache.Redis.ClusterNodes)]
				}
				for clusterNodesCount, clusterNodesItem := range resp.Config.Cache.Redis.ClusterNodes {
					var clusterNodes1 tfTypes.AiProxyAdvancedPluginClusterNodes
					clusterNodes1.IP = types.StringPointerValue(clusterNodesItem.IP)
					clusterNodes1.Port = types.Int64PointerValue(clusterNodesItem.Port)
					if clusterNodesCount+1 > len(r.Config.Cache.Redis.ClusterNodes) {
						r.Config.Cache.Redis.ClusterNodes = append(r.Config.Cache.Redis.ClusterNodes, clusterNodes1)
					} else {
						r.Config.Cache.Redis.ClusterNodes[clusterNodesCount].IP = clusterNodes1.IP
						r.Config.Cache.Redis.ClusterNodes[clusterNodesCount].Port = clusterNodes1.Port
					}
				}
				r.Config.Cache.Redis.ConnectTimeout = types.Int64PointerValue(resp.Config.Cache.Redis.ConnectTimeout)
				r.Config.Cache.Redis.ConnectionIsProxied = types.BoolPointerValue(resp.Config.Cache.Redis.ConnectionIsProxied)
				r.Config.Cache.Redis.Database = types.Int64PointerValue(resp.Config.Cache.Redis.Database)
				r.Config.Cache.Redis.Host = types.StringPointerValue(resp.Config.Cache.Redis.Host)
				r.Config.Cache.Redis.KeepaliveBacklog = types.Int64PointerValue(resp.Config.Cache.Redis.KeepaliveBacklog)
				r.Config.Cache.Redis.KeepalivePoolSize = types.Int64PointerValue(resp.Config.Cache.Redis.KeepalivePoolSize)
				r.Config.Cache.Redis.Password = types.StringPointerValue(resp.Config.Cache.Redis.Password)
				r.Config.Cache.Redis.Port = types.Int64PointerValue(resp.Config.Cache.Redis.Port)
				r.Config.Cache.Redis.ReadTimeout = types.Int64PointerValue(resp.Config.Cache.Redis.ReadTimeout)
				r.Config.Cache.Redis.SendTimeout = types.Int64PointerValue(resp.Config.Cache.Redis.SendTimeout)
				r.Config.Cache.Redis.SentinelMaster = types.StringPointerValue(resp.Config.Cache.Redis.SentinelMaster)
				r.Config.Cache.Redis.SentinelNodes = []tfTypes.AiProxyAdvancedPluginSentinelNodes{}
				if len(r.Config.Cache.Redis.SentinelNodes) > len(resp.Config.Cache.Redis.SentinelNodes) {
					r.Config.Cache.Redis.SentinelNodes = r.Config.Cache.Redis.SentinelNodes[:len(resp.Config.Cache.Redis.SentinelNodes)]
				}
				for sentinelNodesCount, sentinelNodesItem := range resp.Config.Cache.Redis.SentinelNodes {
					var sentinelNodes1 tfTypes.AiProxyAdvancedPluginSentinelNodes
					sentinelNodes1.Host = types.StringPointerValue(sentinelNodesItem.Host)
					sentinelNodes1.Port = types.Int64PointerValue(sentinelNodesItem.Port)
					if sentinelNodesCount+1 > len(r.Config.Cache.Redis.SentinelNodes) {
						r.Config.Cache.Redis.SentinelNodes = append(r.Config.Cache.Redis.SentinelNodes, sentinelNodes1)
					} else {
						r.Config.Cache.Redis.SentinelNodes[sentinelNodesCount].Host = sentinelNodes1.Host
						r.Config.Cache.Redis.SentinelNodes[sentinelNodesCount].Port = sentinelNodes1.Port
					}
				}
				r.Config.Cache.Redis.SentinelPassword = types.StringPointerValue(resp.Config.Cache.Redis.SentinelPassword)
				if resp.Config.Cache.Redis.SentinelRole != nil {
					r.Config.Cache.Redis.SentinelRole = types.StringValue(string(*resp.Config.Cache.Redis.SentinelRole))
				} else {
					r.Config.Cache.Redis.SentinelRole = types.StringNull()
				}
				r.Config.Cache.Redis.SentinelUsername = types.StringPointerValue(resp.Config.Cache.Redis.SentinelUsername)
				r.Config.Cache.Redis.ServerName = types.StringPointerValue(resp.Config.Cache.Redis.ServerName)
				r.Config.Cache.Redis.Ssl = types.BoolPointerValue(resp.Config.Cache.Redis.Ssl)
				r.Config.Cache.Redis.SslVerify = types.BoolPointerValue(resp.Config.Cache.Redis.SslVerify)
				r.Config.Cache.Redis.Username = types.StringPointerValue(resp.Config.Cache.Redis.Username)
			}
			if resp.Config.Cache.Strategy != nil {
				r.Config.Cache.Strategy = types.StringValue(string(*resp.Config.Cache.Strategy))
			} else {
				r.Config.Cache.Strategy = types.StringNull()
			}
		}
		if resp.Config.Client == nil {
			r.Config.Client = nil
		} else {
			r.Config.Client = &tfTypes.Client{}
			if resp.Config.Client.AuthMethod != nil {
				r.Config.Client.AuthMethod = types.StringValue(string(*resp.Config.Client.AuthMethod))
			} else {
				r.Config.Client.AuthMethod = types.StringNull()
			}
			if resp.Config.Client.ClientSecretJwtAlg != nil {
				r.Config.Client.ClientSecretJwtAlg = types.StringValue(string(*resp.Config.Client.ClientSecretJwtAlg))
			} else {
				r.Config.Client.ClientSecretJwtAlg = types.StringNull()
			}
			r.Config.Client.HTTPProxy = types.StringPointerValue(resp.Config.Client.HTTPProxy)
			r.Config.Client.HTTPProxyAuthorization = types.StringPointerValue(resp.Config.Client.HTTPProxyAuthorization)
			if resp.Config.Client.HTTPVersion != nil {
				r.Config.Client.HTTPVersion = types.NumberValue(big.NewFloat(float64(*resp.Config.Client.HTTPVersion)))
			} else {
				r.Config.Client.HTTPVersion = types.NumberNull()
			}
			r.Config.Client.HTTPSProxy = types.StringPointerValue(resp.Config.Client.HTTPSProxy)
			r.Config.Client.HTTPSProxyAuthorization = types.StringPointerValue(resp.Config.Client.HTTPSProxyAuthorization)
			r.Config.Client.KeepAlive = types.BoolPointerValue(resp.Config.Client.KeepAlive)
			r.Config.Client.NoProxy = types.StringPointerValue(resp.Config.Client.NoProxy)
			r.Config.Client.SslVerify = types.BoolPointerValue(resp.Config.Client.SslVerify)
			r.Config.Client.Timeout = types.Int64PointerValue(resp.Config.Client.Timeout)
		}
		if resp.Config.Oauth == nil {
			r.Config.Oauth = nil
		} else {
			r.Config.Oauth = &tfTypes.Oauth{}
			r.Config.Oauth.Audience = make([]types.String, 0, len(resp.Config.Oauth.Audience))
			for _, v := range resp.Config.Oauth.Audience {
				r.Config.Oauth.Audience = append(r.Config.Oauth.Audience, types.StringValue(v))
			}
			r.Config.Oauth.ClientID = types.StringPointerValue(resp.Config.Oauth.ClientID)
			r.Config.Oauth.ClientSecret = types.StringPointerValue(resp.Config.Oauth.ClientSecret)
			if resp.Config.Oauth.GrantType != nil {
				r.Config.Oauth.GrantType = types.StringValue(string(*resp.Config.Oauth.GrantType))
			} else {
				r.Config.Oauth.GrantType = types.StringNull()
			}
			r.Config.Oauth.Password = types.StringPointerValue(resp.Config.Oauth.Password)
			r.Config.Oauth.Scopes = make([]types.String, 0, len(resp.Config.Oauth.Scopes))
			for _, v := range resp.Config.Oauth.Scopes {
				r.Config.Oauth.Scopes = append(r.Config.Oauth.Scopes, types.StringValue(v))
			}
			r.Config.Oauth.TokenEndpoint = types.StringPointerValue(resp.Config.Oauth.TokenEndpoint)
			if len(resp.Config.Oauth.TokenHeaders) > 0 {
				r.Config.Oauth.TokenHeaders = make(map[string]types.String)
				for key, value := range resp.Config.Oauth.TokenHeaders {
					result, _ := json.Marshal(value)
					r.Config.Oauth.TokenHeaders[key] = types.StringValue(string(result))
				}
			}
			if len(resp.Config.Oauth.TokenPostArgs) > 0 {
				r.Config.Oauth.TokenPostArgs = make(map[string]types.String)
				for key1, value1 := range resp.Config.Oauth.TokenPostArgs {
					result1, _ := json.Marshal(value1)
					r.Config.Oauth.TokenPostArgs[key1] = types.StringValue(string(result1))
				}
			}
			r.Config.Oauth.Username = types.StringPointerValue(resp.Config.Oauth.Username)
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
