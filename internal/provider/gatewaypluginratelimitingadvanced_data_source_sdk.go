// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginRateLimitingAdvancedDataSourceModel) RefreshFromSharedRateLimitingAdvancedPlugin(resp *shared.RateLimitingAdvancedPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateRateLimitingAdvancedPluginConfig{}
			r.Config.ConsumerGroups = []types.String{}
			for _, v := range resp.Config.ConsumerGroups {
				r.Config.ConsumerGroups = append(r.Config.ConsumerGroups, types.StringValue(v))
			}
			r.Config.DictionaryName = types.StringPointerValue(resp.Config.DictionaryName)
			r.Config.DisablePenalty = types.BoolPointerValue(resp.Config.DisablePenalty)
			r.Config.EnforceConsumerGroups = types.BoolPointerValue(resp.Config.EnforceConsumerGroups)
			if resp.Config.ErrorCode != nil {
				r.Config.ErrorCode = types.NumberValue(big.NewFloat(float64(*resp.Config.ErrorCode)))
			} else {
				r.Config.ErrorCode = types.NumberNull()
			}
			r.Config.ErrorMessage = types.StringPointerValue(resp.Config.ErrorMessage)
			r.Config.HeaderName = types.StringPointerValue(resp.Config.HeaderName)
			r.Config.HideClientHeaders = types.BoolPointerValue(resp.Config.HideClientHeaders)
			if resp.Config.Identifier != nil {
				r.Config.Identifier = types.StringValue(string(*resp.Config.Identifier))
			} else {
				r.Config.Identifier = types.StringNull()
			}
			r.Config.Limit = []types.Number{}
			for _, v := range resp.Config.Limit {
				r.Config.Limit = append(r.Config.Limit, types.NumberValue(big.NewFloat(float64(v))))
			}
			r.Config.Namespace = types.StringPointerValue(resp.Config.Namespace)
			r.Config.Path = types.StringPointerValue(resp.Config.Path)
			if resp.Config.Redis == nil {
				r.Config.Redis = nil
			} else {
				r.Config.Redis = &tfTypes.CreateRateLimitingAdvancedPluginRedis{}
				r.Config.Redis.ClusterAddresses = []types.String{}
				for _, v := range resp.Config.Redis.ClusterAddresses {
					r.Config.Redis.ClusterAddresses = append(r.Config.Redis.ClusterAddresses, types.StringValue(v))
				}
				r.Config.Redis.ConnectTimeout = types.Int64PointerValue(resp.Config.Redis.ConnectTimeout)
				r.Config.Redis.Database = types.Int64PointerValue(resp.Config.Redis.Database)
				r.Config.Redis.Host = types.StringPointerValue(resp.Config.Redis.Host)
				r.Config.Redis.KeepaliveBacklog = types.Int64PointerValue(resp.Config.Redis.KeepaliveBacklog)
				r.Config.Redis.KeepalivePoolSize = types.Int64PointerValue(resp.Config.Redis.KeepalivePoolSize)
				r.Config.Redis.Password = types.StringPointerValue(resp.Config.Redis.Password)
				r.Config.Redis.Port = types.Int64PointerValue(resp.Config.Redis.Port)
				r.Config.Redis.ReadTimeout = types.Int64PointerValue(resp.Config.Redis.ReadTimeout)
				r.Config.Redis.SendTimeout = types.Int64PointerValue(resp.Config.Redis.SendTimeout)
				r.Config.Redis.SentinelAddresses = []types.String{}
				for _, v := range resp.Config.Redis.SentinelAddresses {
					r.Config.Redis.SentinelAddresses = append(r.Config.Redis.SentinelAddresses, types.StringValue(v))
				}
				r.Config.Redis.SentinelMaster = types.StringPointerValue(resp.Config.Redis.SentinelMaster)
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
				r.Config.Redis.Timeout = types.Int64PointerValue(resp.Config.Redis.Timeout)
				r.Config.Redis.Username = types.StringPointerValue(resp.Config.Redis.Username)
			}
			if resp.Config.RetryAfterJitterMax != nil {
				r.Config.RetryAfterJitterMax = types.NumberValue(big.NewFloat(float64(*resp.Config.RetryAfterJitterMax)))
			} else {
				r.Config.RetryAfterJitterMax = types.NumberNull()
			}
			if resp.Config.Strategy != nil {
				r.Config.Strategy = types.StringValue(string(*resp.Config.Strategy))
			} else {
				r.Config.Strategy = types.StringNull()
			}
			if resp.Config.SyncRate != nil {
				r.Config.SyncRate = types.NumberValue(big.NewFloat(float64(*resp.Config.SyncRate)))
			} else {
				r.Config.SyncRate = types.NumberNull()
			}
			r.Config.WindowSize = []types.Number{}
			for _, v := range resp.Config.WindowSize {
				r.Config.WindowSize = append(r.Config.WindowSize, types.NumberValue(big.NewFloat(float64(v))))
			}
			if resp.Config.WindowType != nil {
				r.Config.WindowType = types.StringValue(string(*resp.Config.WindowType))
			} else {
				r.Config.WindowType = types.StringNull()
			}
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
