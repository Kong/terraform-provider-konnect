// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginAiRateLimitingAdvancedResourceModel) ToSharedAiRateLimitingAdvancedPlugin() *shared.AiRateLimitingAdvancedPlugin {
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
	var ordering *shared.AiRateLimitingAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.AiRateLimitingAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AiRateLimitingAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.AiRateLimitingAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AiRateLimitingAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AiRateLimitingAdvancedPluginOrdering{
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
	var config *shared.AiRateLimitingAdvancedPluginConfig
	if r.Config != nil {
		dictionaryName := new(string)
		if !r.Config.DictionaryName.IsUnknown() && !r.Config.DictionaryName.IsNull() {
			*dictionaryName = r.Config.DictionaryName.ValueString()
		} else {
			dictionaryName = nil
		}
		disablePenalty := new(bool)
		if !r.Config.DisablePenalty.IsUnknown() && !r.Config.DisablePenalty.IsNull() {
			*disablePenalty = r.Config.DisablePenalty.ValueBool()
		} else {
			disablePenalty = nil
		}
		errorCode := new(float64)
		if !r.Config.ErrorCode.IsUnknown() && !r.Config.ErrorCode.IsNull() {
			*errorCode = r.Config.ErrorCode.ValueFloat64()
		} else {
			errorCode = nil
		}
		errorHideProviders := new(bool)
		if !r.Config.ErrorHideProviders.IsUnknown() && !r.Config.ErrorHideProviders.IsNull() {
			*errorHideProviders = r.Config.ErrorHideProviders.ValueBool()
		} else {
			errorHideProviders = nil
		}
		errorMessage := new(string)
		if !r.Config.ErrorMessage.IsUnknown() && !r.Config.ErrorMessage.IsNull() {
			*errorMessage = r.Config.ErrorMessage.ValueString()
		} else {
			errorMessage = nil
		}
		headerName := new(string)
		if !r.Config.HeaderName.IsUnknown() && !r.Config.HeaderName.IsNull() {
			*headerName = r.Config.HeaderName.ValueString()
		} else {
			headerName = nil
		}
		hideClientHeaders := new(bool)
		if !r.Config.HideClientHeaders.IsUnknown() && !r.Config.HideClientHeaders.IsNull() {
			*hideClientHeaders = r.Config.HideClientHeaders.ValueBool()
		} else {
			hideClientHeaders = nil
		}
		identifier := new(shared.Identifier)
		if !r.Config.Identifier.IsUnknown() && !r.Config.Identifier.IsNull() {
			*identifier = shared.Identifier(r.Config.Identifier.ValueString())
		} else {
			identifier = nil
		}
		var llmProviders []shared.LlmProviders = []shared.LlmProviders{}
		for _, llmProvidersItem := range r.Config.LlmProviders {
			var limit float64
			limit = llmProvidersItem.Limit.ValueFloat64()

			name := shared.Name(llmProvidersItem.Name.ValueString())
			var windowSize float64
			windowSize = llmProvidersItem.WindowSize.ValueFloat64()

			llmProviders = append(llmProviders, shared.LlmProviders{
				Limit:      limit,
				Name:       name,
				WindowSize: windowSize,
			})
		}
		path := new(string)
		if !r.Config.Path.IsUnknown() && !r.Config.Path.IsNull() {
			*path = r.Config.Path.ValueString()
		} else {
			path = nil
		}
		var redis *shared.Redis
		if r.Config.Redis != nil {
			clusterMaxRedirections := new(int64)
			if !r.Config.Redis.ClusterMaxRedirections.IsUnknown() && !r.Config.Redis.ClusterMaxRedirections.IsNull() {
				*clusterMaxRedirections = r.Config.Redis.ClusterMaxRedirections.ValueInt64()
			} else {
				clusterMaxRedirections = nil
			}
			var clusterNodes []shared.ClusterNodes = []shared.ClusterNodes{}
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
				clusterNodes = append(clusterNodes, shared.ClusterNodes{
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
			var sentinelNodes []shared.SentinelNodes = []shared.SentinelNodes{}
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
				sentinelNodes = append(sentinelNodes, shared.SentinelNodes{
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
			sentinelRole := new(shared.SentinelRole)
			if !r.Config.Redis.SentinelRole.IsUnknown() && !r.Config.Redis.SentinelRole.IsNull() {
				*sentinelRole = shared.SentinelRole(r.Config.Redis.SentinelRole.ValueString())
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
			redis = &shared.Redis{
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
		requestPromptCountFunction := new(string)
		if !r.Config.RequestPromptCountFunction.IsUnknown() && !r.Config.RequestPromptCountFunction.IsNull() {
			*requestPromptCountFunction = r.Config.RequestPromptCountFunction.ValueString()
		} else {
			requestPromptCountFunction = nil
		}
		retryAfterJitterMax := new(float64)
		if !r.Config.RetryAfterJitterMax.IsUnknown() && !r.Config.RetryAfterJitterMax.IsNull() {
			*retryAfterJitterMax = r.Config.RetryAfterJitterMax.ValueFloat64()
		} else {
			retryAfterJitterMax = nil
		}
		strategy := new(shared.Strategy)
		if !r.Config.Strategy.IsUnknown() && !r.Config.Strategy.IsNull() {
			*strategy = shared.Strategy(r.Config.Strategy.ValueString())
		} else {
			strategy = nil
		}
		syncRate := new(float64)
		if !r.Config.SyncRate.IsUnknown() && !r.Config.SyncRate.IsNull() {
			*syncRate = r.Config.SyncRate.ValueFloat64()
		} else {
			syncRate = nil
		}
		tokensCountStrategy := new(shared.TokensCountStrategy)
		if !r.Config.TokensCountStrategy.IsUnknown() && !r.Config.TokensCountStrategy.IsNull() {
			*tokensCountStrategy = shared.TokensCountStrategy(r.Config.TokensCountStrategy.ValueString())
		} else {
			tokensCountStrategy = nil
		}
		windowType := new(shared.WindowType)
		if !r.Config.WindowType.IsUnknown() && !r.Config.WindowType.IsNull() {
			*windowType = shared.WindowType(r.Config.WindowType.ValueString())
		} else {
			windowType = nil
		}
		config = &shared.AiRateLimitingAdvancedPluginConfig{
			DictionaryName:             dictionaryName,
			DisablePenalty:             disablePenalty,
			ErrorCode:                  errorCode,
			ErrorHideProviders:         errorHideProviders,
			ErrorMessage:               errorMessage,
			HeaderName:                 headerName,
			HideClientHeaders:          hideClientHeaders,
			Identifier:                 identifier,
			LlmProviders:               llmProviders,
			Path:                       path,
			Redis:                      redis,
			RequestPromptCountFunction: requestPromptCountFunction,
			RetryAfterJitterMax:        retryAfterJitterMax,
			Strategy:                   strategy,
			SyncRate:                   syncRate,
			TokensCountStrategy:        tokensCountStrategy,
			WindowType:                 windowType,
		}
	}
	var consumer *shared.AiRateLimitingAdvancedPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.AiRateLimitingAdvancedPluginConsumer{
			ID: id1,
		}
	}
	var consumerGroup *shared.AiRateLimitingAdvancedPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id2 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id2 = r.ConsumerGroup.ID.ValueString()
		} else {
			id2 = nil
		}
		consumerGroup = &shared.AiRateLimitingAdvancedPluginConsumerGroup{
			ID: id2,
		}
	}
	var protocols []shared.AiRateLimitingAdvancedPluginProtocols = []shared.AiRateLimitingAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AiRateLimitingAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AiRateLimitingAdvancedPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.AiRateLimitingAdvancedPluginRoute{
			ID: id3,
		}
	}
	var service *shared.AiRateLimitingAdvancedPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.AiRateLimitingAdvancedPluginService{
			ID: id4,
		}
	}
	out := shared.AiRateLimitingAdvancedPlugin{
		CreatedAt:     createdAt,
		Enabled:       enabled,
		ID:            id,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Tags:          tags,
		UpdatedAt:     updatedAt,
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginAiRateLimitingAdvancedResourceModel) RefreshFromSharedAiRateLimitingAdvancedPlugin(ctx context.Context, resp *shared.AiRateLimitingAdvancedPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.AiRateLimitingAdvancedPluginConfig{}
			r.Config.DictionaryName = types.StringPointerValue(resp.Config.DictionaryName)
			r.Config.DisablePenalty = types.BoolPointerValue(resp.Config.DisablePenalty)
			r.Config.ErrorCode = types.Float64PointerValue(resp.Config.ErrorCode)
			r.Config.ErrorHideProviders = types.BoolPointerValue(resp.Config.ErrorHideProviders)
			r.Config.ErrorMessage = types.StringPointerValue(resp.Config.ErrorMessage)
			r.Config.HeaderName = types.StringPointerValue(resp.Config.HeaderName)
			r.Config.HideClientHeaders = types.BoolPointerValue(resp.Config.HideClientHeaders)
			if resp.Config.Identifier != nil {
				r.Config.Identifier = types.StringValue(string(*resp.Config.Identifier))
			} else {
				r.Config.Identifier = types.StringNull()
			}
			r.Config.LlmProviders = []tfTypes.LlmProviders{}
			if len(r.Config.LlmProviders) > len(resp.Config.LlmProviders) {
				r.Config.LlmProviders = r.Config.LlmProviders[:len(resp.Config.LlmProviders)]
			}
			for llmProvidersCount, llmProvidersItem := range resp.Config.LlmProviders {
				var llmProviders tfTypes.LlmProviders
				llmProviders.Limit = types.Float64Value(llmProvidersItem.Limit)
				llmProviders.Name = types.StringValue(string(llmProvidersItem.Name))
				llmProviders.WindowSize = types.Float64Value(llmProvidersItem.WindowSize)
				if llmProvidersCount+1 > len(r.Config.LlmProviders) {
					r.Config.LlmProviders = append(r.Config.LlmProviders, llmProviders)
				} else {
					r.Config.LlmProviders[llmProvidersCount].Limit = llmProviders.Limit
					r.Config.LlmProviders[llmProvidersCount].Name = llmProviders.Name
					r.Config.LlmProviders[llmProvidersCount].WindowSize = llmProviders.WindowSize
				}
			}
			r.Config.Path = types.StringPointerValue(resp.Config.Path)
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
			r.Config.RequestPromptCountFunction = types.StringPointerValue(resp.Config.RequestPromptCountFunction)
			r.Config.RetryAfterJitterMax = types.Float64PointerValue(resp.Config.RetryAfterJitterMax)
			if resp.Config.Strategy != nil {
				r.Config.Strategy = types.StringValue(string(*resp.Config.Strategy))
			} else {
				r.Config.Strategy = types.StringNull()
			}
			r.Config.SyncRate = types.Float64PointerValue(resp.Config.SyncRate)
			if resp.Config.TokensCountStrategy != nil {
				r.Config.TokensCountStrategy = types.StringValue(string(*resp.Config.TokensCountStrategy))
			} else {
				r.Config.TokensCountStrategy = types.StringNull()
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

	return diags
}
