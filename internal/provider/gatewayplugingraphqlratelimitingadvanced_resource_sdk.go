// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginGraphqlRateLimitingAdvancedResourceModel) ToSharedGraphqlRateLimitingAdvancedPlugin() *shared.GraphqlRateLimitingAdvancedPlugin {
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
	var ordering *shared.GraphqlRateLimitingAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.GraphqlRateLimitingAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.GraphqlRateLimitingAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.GraphqlRateLimitingAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.GraphqlRateLimitingAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.GraphqlRateLimitingAdvancedPluginOrdering{
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
	var config *shared.GraphqlRateLimitingAdvancedPluginConfig
	if r.Config != nil {
		costStrategy := new(shared.CostStrategy)
		if !r.Config.CostStrategy.IsUnknown() && !r.Config.CostStrategy.IsNull() {
			*costStrategy = shared.CostStrategy(r.Config.CostStrategy.ValueString())
		} else {
			costStrategy = nil
		}
		dictionaryName := new(string)
		if !r.Config.DictionaryName.IsUnknown() && !r.Config.DictionaryName.IsNull() {
			*dictionaryName = r.Config.DictionaryName.ValueString()
		} else {
			dictionaryName = nil
		}
		hideClientHeaders := new(bool)
		if !r.Config.HideClientHeaders.IsUnknown() && !r.Config.HideClientHeaders.IsNull() {
			*hideClientHeaders = r.Config.HideClientHeaders.ValueBool()
		} else {
			hideClientHeaders = nil
		}
		identifier := new(shared.GraphqlRateLimitingAdvancedPluginIdentifier)
		if !r.Config.Identifier.IsUnknown() && !r.Config.Identifier.IsNull() {
			*identifier = shared.GraphqlRateLimitingAdvancedPluginIdentifier(r.Config.Identifier.ValueString())
		} else {
			identifier = nil
		}
		var limit []float64 = []float64{}
		for _, limitItem := range r.Config.Limit {
			limit = append(limit, limitItem.ValueFloat64())
		}
		maxCost := new(float64)
		if !r.Config.MaxCost.IsUnknown() && !r.Config.MaxCost.IsNull() {
			*maxCost = r.Config.MaxCost.ValueFloat64()
		} else {
			maxCost = nil
		}
		namespace := new(string)
		if !r.Config.Namespace.IsUnknown() && !r.Config.Namespace.IsNull() {
			*namespace = r.Config.Namespace.ValueString()
		} else {
			namespace = nil
		}
		var redis *shared.GraphqlRateLimitingAdvancedPluginRedis
		if r.Config.Redis != nil {
			clusterMaxRedirections := new(int64)
			if !r.Config.Redis.ClusterMaxRedirections.IsUnknown() && !r.Config.Redis.ClusterMaxRedirections.IsNull() {
				*clusterMaxRedirections = r.Config.Redis.ClusterMaxRedirections.ValueInt64()
			} else {
				clusterMaxRedirections = nil
			}
			var clusterNodes []shared.GraphqlRateLimitingAdvancedPluginClusterNodes = []shared.GraphqlRateLimitingAdvancedPluginClusterNodes{}
			for _, clusterNodesItem := range r.Config.Redis.ClusterNodes {
				ip := new(string)
				if !clusterNodesItem.IP.IsUnknown() && !clusterNodesItem.IP.IsNull() {
					*ip = clusterNodesItem.IP.ValueString()
				} else {
					ip = nil
				}
				port := new(int64)
				if !clusterNodesItem.Port.IsUnknown() && !clusterNodesItem.Port.IsNull() {
					*port = clusterNodesItem.Port.ValueInt64()
				} else {
					port = nil
				}
				clusterNodes = append(clusterNodes, shared.GraphqlRateLimitingAdvancedPluginClusterNodes{
					IP:   ip,
					Port: port,
				})
			}
			connectTimeout := new(int64)
			if !r.Config.Redis.ConnectTimeout.IsUnknown() && !r.Config.Redis.ConnectTimeout.IsNull() {
				*connectTimeout = r.Config.Redis.ConnectTimeout.ValueInt64()
			} else {
				connectTimeout = nil
			}
			connectionIsProxied := new(bool)
			if !r.Config.Redis.ConnectionIsProxied.IsUnknown() && !r.Config.Redis.ConnectionIsProxied.IsNull() {
				*connectionIsProxied = r.Config.Redis.ConnectionIsProxied.ValueBool()
			} else {
				connectionIsProxied = nil
			}
			database := new(int64)
			if !r.Config.Redis.Database.IsUnknown() && !r.Config.Redis.Database.IsNull() {
				*database = r.Config.Redis.Database.ValueInt64()
			} else {
				database = nil
			}
			host := new(string)
			if !r.Config.Redis.Host.IsUnknown() && !r.Config.Redis.Host.IsNull() {
				*host = r.Config.Redis.Host.ValueString()
			} else {
				host = nil
			}
			keepaliveBacklog := new(int64)
			if !r.Config.Redis.KeepaliveBacklog.IsUnknown() && !r.Config.Redis.KeepaliveBacklog.IsNull() {
				*keepaliveBacklog = r.Config.Redis.KeepaliveBacklog.ValueInt64()
			} else {
				keepaliveBacklog = nil
			}
			keepalivePoolSize := new(int64)
			if !r.Config.Redis.KeepalivePoolSize.IsUnknown() && !r.Config.Redis.KeepalivePoolSize.IsNull() {
				*keepalivePoolSize = r.Config.Redis.KeepalivePoolSize.ValueInt64()
			} else {
				keepalivePoolSize = nil
			}
			password := new(string)
			if !r.Config.Redis.Password.IsUnknown() && !r.Config.Redis.Password.IsNull() {
				*password = r.Config.Redis.Password.ValueString()
			} else {
				password = nil
			}
			port1 := new(int64)
			if !r.Config.Redis.Port.IsUnknown() && !r.Config.Redis.Port.IsNull() {
				*port1 = r.Config.Redis.Port.ValueInt64()
			} else {
				port1 = nil
			}
			readTimeout := new(int64)
			if !r.Config.Redis.ReadTimeout.IsUnknown() && !r.Config.Redis.ReadTimeout.IsNull() {
				*readTimeout = r.Config.Redis.ReadTimeout.ValueInt64()
			} else {
				readTimeout = nil
			}
			sendTimeout := new(int64)
			if !r.Config.Redis.SendTimeout.IsUnknown() && !r.Config.Redis.SendTimeout.IsNull() {
				*sendTimeout = r.Config.Redis.SendTimeout.ValueInt64()
			} else {
				sendTimeout = nil
			}
			sentinelMaster := new(string)
			if !r.Config.Redis.SentinelMaster.IsUnknown() && !r.Config.Redis.SentinelMaster.IsNull() {
				*sentinelMaster = r.Config.Redis.SentinelMaster.ValueString()
			} else {
				sentinelMaster = nil
			}
			var sentinelNodes []shared.GraphqlRateLimitingAdvancedPluginSentinelNodes = []shared.GraphqlRateLimitingAdvancedPluginSentinelNodes{}
			for _, sentinelNodesItem := range r.Config.Redis.SentinelNodes {
				host1 := new(string)
				if !sentinelNodesItem.Host.IsUnknown() && !sentinelNodesItem.Host.IsNull() {
					*host1 = sentinelNodesItem.Host.ValueString()
				} else {
					host1 = nil
				}
				port2 := new(int64)
				if !sentinelNodesItem.Port.IsUnknown() && !sentinelNodesItem.Port.IsNull() {
					*port2 = sentinelNodesItem.Port.ValueInt64()
				} else {
					port2 = nil
				}
				sentinelNodes = append(sentinelNodes, shared.GraphqlRateLimitingAdvancedPluginSentinelNodes{
					Host: host1,
					Port: port2,
				})
			}
			sentinelPassword := new(string)
			if !r.Config.Redis.SentinelPassword.IsUnknown() && !r.Config.Redis.SentinelPassword.IsNull() {
				*sentinelPassword = r.Config.Redis.SentinelPassword.ValueString()
			} else {
				sentinelPassword = nil
			}
			sentinelRole := new(shared.GraphqlRateLimitingAdvancedPluginSentinelRole)
			if !r.Config.Redis.SentinelRole.IsUnknown() && !r.Config.Redis.SentinelRole.IsNull() {
				*sentinelRole = shared.GraphqlRateLimitingAdvancedPluginSentinelRole(r.Config.Redis.SentinelRole.ValueString())
			} else {
				sentinelRole = nil
			}
			sentinelUsername := new(string)
			if !r.Config.Redis.SentinelUsername.IsUnknown() && !r.Config.Redis.SentinelUsername.IsNull() {
				*sentinelUsername = r.Config.Redis.SentinelUsername.ValueString()
			} else {
				sentinelUsername = nil
			}
			serverName := new(string)
			if !r.Config.Redis.ServerName.IsUnknown() && !r.Config.Redis.ServerName.IsNull() {
				*serverName = r.Config.Redis.ServerName.ValueString()
			} else {
				serverName = nil
			}
			ssl := new(bool)
			if !r.Config.Redis.Ssl.IsUnknown() && !r.Config.Redis.Ssl.IsNull() {
				*ssl = r.Config.Redis.Ssl.ValueBool()
			} else {
				ssl = nil
			}
			sslVerify := new(bool)
			if !r.Config.Redis.SslVerify.IsUnknown() && !r.Config.Redis.SslVerify.IsNull() {
				*sslVerify = r.Config.Redis.SslVerify.ValueBool()
			} else {
				sslVerify = nil
			}
			username := new(string)
			if !r.Config.Redis.Username.IsUnknown() && !r.Config.Redis.Username.IsNull() {
				*username = r.Config.Redis.Username.ValueString()
			} else {
				username = nil
			}
			redis = &shared.GraphqlRateLimitingAdvancedPluginRedis{
				ClusterMaxRedirections: clusterMaxRedirections,
				ClusterNodes:           clusterNodes,
				ConnectTimeout:         connectTimeout,
				ConnectionIsProxied:    connectionIsProxied,
				Database:               database,
				Host:                   host,
				KeepaliveBacklog:       keepaliveBacklog,
				KeepalivePoolSize:      keepalivePoolSize,
				Password:               password,
				Port:                   port1,
				ReadTimeout:            readTimeout,
				SendTimeout:            sendTimeout,
				SentinelMaster:         sentinelMaster,
				SentinelNodes:          sentinelNodes,
				SentinelPassword:       sentinelPassword,
				SentinelRole:           sentinelRole,
				SentinelUsername:       sentinelUsername,
				ServerName:             serverName,
				Ssl:                    ssl,
				SslVerify:              sslVerify,
				Username:               username,
			}
		}
		scoreFactor := new(float64)
		if !r.Config.ScoreFactor.IsUnknown() && !r.Config.ScoreFactor.IsNull() {
			*scoreFactor = r.Config.ScoreFactor.ValueFloat64()
		} else {
			scoreFactor = nil
		}
		strategy := new(shared.GraphqlRateLimitingAdvancedPluginStrategy)
		if !r.Config.Strategy.IsUnknown() && !r.Config.Strategy.IsNull() {
			*strategy = shared.GraphqlRateLimitingAdvancedPluginStrategy(r.Config.Strategy.ValueString())
		} else {
			strategy = nil
		}
		syncRate := new(float64)
		if !r.Config.SyncRate.IsUnknown() && !r.Config.SyncRate.IsNull() {
			*syncRate = r.Config.SyncRate.ValueFloat64()
		} else {
			syncRate = nil
		}
		var windowSize []float64 = []float64{}
		for _, windowSizeItem := range r.Config.WindowSize {
			windowSize = append(windowSize, windowSizeItem.ValueFloat64())
		}
		windowType := new(shared.GraphqlRateLimitingAdvancedPluginWindowType)
		if !r.Config.WindowType.IsUnknown() && !r.Config.WindowType.IsNull() {
			*windowType = shared.GraphqlRateLimitingAdvancedPluginWindowType(r.Config.WindowType.ValueString())
		} else {
			windowType = nil
		}
		config = &shared.GraphqlRateLimitingAdvancedPluginConfig{
			CostStrategy:      costStrategy,
			DictionaryName:    dictionaryName,
			HideClientHeaders: hideClientHeaders,
			Identifier:        identifier,
			Limit:             limit,
			MaxCost:           maxCost,
			Namespace:         namespace,
			Redis:             redis,
			ScoreFactor:       scoreFactor,
			Strategy:          strategy,
			SyncRate:          syncRate,
			WindowSize:        windowSize,
			WindowType:        windowType,
		}
	}
	var consumer *shared.GraphqlRateLimitingAdvancedPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.GraphqlRateLimitingAdvancedPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.GraphqlRateLimitingAdvancedPluginProtocols = []shared.GraphqlRateLimitingAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.GraphqlRateLimitingAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.GraphqlRateLimitingAdvancedPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.GraphqlRateLimitingAdvancedPluginRoute{
			ID: id2,
		}
	}
	var service *shared.GraphqlRateLimitingAdvancedPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.GraphqlRateLimitingAdvancedPluginService{
			ID: id3,
		}
	}
	out := shared.GraphqlRateLimitingAdvancedPlugin{
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

func (r *GatewayPluginGraphqlRateLimitingAdvancedResourceModel) RefreshFromSharedGraphqlRateLimitingAdvancedPlugin(ctx context.Context, resp *shared.GraphqlRateLimitingAdvancedPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.GraphqlRateLimitingAdvancedPluginConfig{}
			if resp.Config.CostStrategy != nil {
				r.Config.CostStrategy = types.StringValue(string(*resp.Config.CostStrategy))
			} else {
				r.Config.CostStrategy = types.StringNull()
			}
			r.Config.DictionaryName = types.StringPointerValue(resp.Config.DictionaryName)
			r.Config.HideClientHeaders = types.BoolPointerValue(resp.Config.HideClientHeaders)
			if resp.Config.Identifier != nil {
				r.Config.Identifier = types.StringValue(string(*resp.Config.Identifier))
			} else {
				r.Config.Identifier = types.StringNull()
			}
			r.Config.Limit = make([]types.Float64, 0, len(resp.Config.Limit))
			for _, v := range resp.Config.Limit {
				r.Config.Limit = append(r.Config.Limit, types.Float64Value(v))
			}
			r.Config.MaxCost = types.Float64PointerValue(resp.Config.MaxCost)
			r.Config.Namespace = types.StringPointerValue(resp.Config.Namespace)
			if resp.Config.Redis == nil {
				r.Config.Redis = nil
			} else {
				r.Config.Redis = &tfTypes.AiProxyAdvancedPluginRedis{}
				r.Config.Redis.ClusterMaxRedirections = types.Int64PointerValue(resp.Config.Redis.ClusterMaxRedirections)
				r.Config.Redis.ClusterNodes = []tfTypes.AiProxyAdvancedPluginClusterNodes{}
				if len(r.Config.Redis.ClusterNodes) > len(resp.Config.Redis.ClusterNodes) {
					r.Config.Redis.ClusterNodes = r.Config.Redis.ClusterNodes[:len(resp.Config.Redis.ClusterNodes)]
				}
				for clusterNodesCount, clusterNodesItem := range resp.Config.Redis.ClusterNodes {
					var clusterNodes tfTypes.AiProxyAdvancedPluginClusterNodes
					clusterNodes.IP = types.StringPointerValue(clusterNodesItem.IP)
					clusterNodes.Port = types.Int64PointerValue(clusterNodesItem.Port)
					if clusterNodesCount+1 > len(r.Config.Redis.ClusterNodes) {
						r.Config.Redis.ClusterNodes = append(r.Config.Redis.ClusterNodes, clusterNodes)
					} else {
						r.Config.Redis.ClusterNodes[clusterNodesCount].IP = clusterNodes.IP
						r.Config.Redis.ClusterNodes[clusterNodesCount].Port = clusterNodes.Port
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
				r.Config.Redis.SentinelNodes = []tfTypes.AiProxyAdvancedPluginSentinelNodes{}
				if len(r.Config.Redis.SentinelNodes) > len(resp.Config.Redis.SentinelNodes) {
					r.Config.Redis.SentinelNodes = r.Config.Redis.SentinelNodes[:len(resp.Config.Redis.SentinelNodes)]
				}
				for sentinelNodesCount, sentinelNodesItem := range resp.Config.Redis.SentinelNodes {
					var sentinelNodes tfTypes.AiProxyAdvancedPluginSentinelNodes
					sentinelNodes.Host = types.StringPointerValue(sentinelNodesItem.Host)
					sentinelNodes.Port = types.Int64PointerValue(sentinelNodesItem.Port)
					if sentinelNodesCount+1 > len(r.Config.Redis.SentinelNodes) {
						r.Config.Redis.SentinelNodes = append(r.Config.Redis.SentinelNodes, sentinelNodes)
					} else {
						r.Config.Redis.SentinelNodes[sentinelNodesCount].Host = sentinelNodes.Host
						r.Config.Redis.SentinelNodes[sentinelNodesCount].Port = sentinelNodes.Port
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
			r.Config.ScoreFactor = types.Float64PointerValue(resp.Config.ScoreFactor)
			if resp.Config.Strategy != nil {
				r.Config.Strategy = types.StringValue(string(*resp.Config.Strategy))
			} else {
				r.Config.Strategy = types.StringNull()
			}
			r.Config.SyncRate = types.Float64PointerValue(resp.Config.SyncRate)
			r.Config.WindowSize = make([]types.Float64, 0, len(resp.Config.WindowSize))
			for _, v := range resp.Config.WindowSize {
				r.Config.WindowSize = append(r.Config.WindowSize, types.Float64Value(v))
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
