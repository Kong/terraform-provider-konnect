// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginUpstreamOauthResourceModel) ToSharedUpstreamOauthPlugin() *shared.UpstreamOauthPlugin {
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
	var ordering *shared.UpstreamOauthPluginOrdering
	if r.Ordering != nil {
		var after *shared.UpstreamOauthPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.UpstreamOauthPluginAfter{
				Access: access,
			}
		}
		var before *shared.UpstreamOauthPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.UpstreamOauthPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.UpstreamOauthPluginOrdering{
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
	var config *shared.UpstreamOauthPluginConfig
	if r.Config != nil {
		var behavior *shared.Behavior
		if r.Config.Behavior != nil {
			idpErrorResponseBodyTemplate := new(string)
			if !r.Config.Behavior.IdpErrorResponseBodyTemplate.IsUnknown() && !r.Config.Behavior.IdpErrorResponseBodyTemplate.IsNull() {
				*idpErrorResponseBodyTemplate = r.Config.Behavior.IdpErrorResponseBodyTemplate.ValueString()
			} else {
				idpErrorResponseBodyTemplate = nil
			}
			idpErrorResponseContentType := new(string)
			if !r.Config.Behavior.IdpErrorResponseContentType.IsUnknown() && !r.Config.Behavior.IdpErrorResponseContentType.IsNull() {
				*idpErrorResponseContentType = r.Config.Behavior.IdpErrorResponseContentType.ValueString()
			} else {
				idpErrorResponseContentType = nil
			}
			idpErrorResponseMessage := new(string)
			if !r.Config.Behavior.IdpErrorResponseMessage.IsUnknown() && !r.Config.Behavior.IdpErrorResponseMessage.IsNull() {
				*idpErrorResponseMessage = r.Config.Behavior.IdpErrorResponseMessage.ValueString()
			} else {
				idpErrorResponseMessage = nil
			}
			idpErrorResponseStatusCode := new(int64)
			if !r.Config.Behavior.IdpErrorResponseStatusCode.IsUnknown() && !r.Config.Behavior.IdpErrorResponseStatusCode.IsNull() {
				*idpErrorResponseStatusCode = r.Config.Behavior.IdpErrorResponseStatusCode.ValueInt64()
			} else {
				idpErrorResponseStatusCode = nil
			}
			var purgeTokenOnUpstreamStatusCodes []int64 = []int64{}
			for _, purgeTokenOnUpstreamStatusCodesItem := range r.Config.Behavior.PurgeTokenOnUpstreamStatusCodes {
				purgeTokenOnUpstreamStatusCodes = append(purgeTokenOnUpstreamStatusCodes, purgeTokenOnUpstreamStatusCodesItem.ValueInt64())
			}
			upstreamAccessTokenHeaderName := new(string)
			if !r.Config.Behavior.UpstreamAccessTokenHeaderName.IsUnknown() && !r.Config.Behavior.UpstreamAccessTokenHeaderName.IsNull() {
				*upstreamAccessTokenHeaderName = r.Config.Behavior.UpstreamAccessTokenHeaderName.ValueString()
			} else {
				upstreamAccessTokenHeaderName = nil
			}
			behavior = &shared.Behavior{
				IdpErrorResponseBodyTemplate:    idpErrorResponseBodyTemplate,
				IdpErrorResponseContentType:     idpErrorResponseContentType,
				IdpErrorResponseMessage:         idpErrorResponseMessage,
				IdpErrorResponseStatusCode:      idpErrorResponseStatusCode,
				PurgeTokenOnUpstreamStatusCodes: purgeTokenOnUpstreamStatusCodes,
				UpstreamAccessTokenHeaderName:   upstreamAccessTokenHeaderName,
			}
		}
		var cache *shared.Cache
		if r.Config.Cache != nil {
			defaultTTL := new(float64)
			if !r.Config.Cache.DefaultTTL.IsUnknown() && !r.Config.Cache.DefaultTTL.IsNull() {
				*defaultTTL = r.Config.Cache.DefaultTTL.ValueFloat64()
			} else {
				defaultTTL = nil
			}
			eagerlyExpire := new(int64)
			if !r.Config.Cache.EagerlyExpire.IsUnknown() && !r.Config.Cache.EagerlyExpire.IsNull() {
				*eagerlyExpire = r.Config.Cache.EagerlyExpire.ValueInt64()
			} else {
				eagerlyExpire = nil
			}
			var memory *shared.UpstreamOauthPluginMemory
			if r.Config.Cache.Memory != nil {
				dictionaryName := new(string)
				if !r.Config.Cache.Memory.DictionaryName.IsUnknown() && !r.Config.Cache.Memory.DictionaryName.IsNull() {
					*dictionaryName = r.Config.Cache.Memory.DictionaryName.ValueString()
				} else {
					dictionaryName = nil
				}
				memory = &shared.UpstreamOauthPluginMemory{
					DictionaryName: dictionaryName,
				}
			}
			var redis *shared.UpstreamOauthPluginRedis
			if r.Config.Cache.Redis != nil {
				clusterMaxRedirections := new(int64)
				if !r.Config.Cache.Redis.ClusterMaxRedirections.IsUnknown() && !r.Config.Cache.Redis.ClusterMaxRedirections.IsNull() {
					*clusterMaxRedirections = r.Config.Cache.Redis.ClusterMaxRedirections.ValueInt64()
				} else {
					clusterMaxRedirections = nil
				}
				var clusterNodes []shared.UpstreamOauthPluginClusterNodes = []shared.UpstreamOauthPluginClusterNodes{}
				for _, clusterNodesItem := range r.Config.Cache.Redis.ClusterNodes {
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
					clusterNodes = append(clusterNodes, shared.UpstreamOauthPluginClusterNodes{
						IP:   ip,
						Port: port,
					})
				}
				connectTimeout := new(int64)
				if !r.Config.Cache.Redis.ConnectTimeout.IsUnknown() && !r.Config.Cache.Redis.ConnectTimeout.IsNull() {
					*connectTimeout = r.Config.Cache.Redis.ConnectTimeout.ValueInt64()
				} else {
					connectTimeout = nil
				}
				connectionIsProxied := new(bool)
				if !r.Config.Cache.Redis.ConnectionIsProxied.IsUnknown() && !r.Config.Cache.Redis.ConnectionIsProxied.IsNull() {
					*connectionIsProxied = r.Config.Cache.Redis.ConnectionIsProxied.ValueBool()
				} else {
					connectionIsProxied = nil
				}
				database := new(int64)
				if !r.Config.Cache.Redis.Database.IsUnknown() && !r.Config.Cache.Redis.Database.IsNull() {
					*database = r.Config.Cache.Redis.Database.ValueInt64()
				} else {
					database = nil
				}
				host := new(string)
				if !r.Config.Cache.Redis.Host.IsUnknown() && !r.Config.Cache.Redis.Host.IsNull() {
					*host = r.Config.Cache.Redis.Host.ValueString()
				} else {
					host = nil
				}
				keepaliveBacklog := new(int64)
				if !r.Config.Cache.Redis.KeepaliveBacklog.IsUnknown() && !r.Config.Cache.Redis.KeepaliveBacklog.IsNull() {
					*keepaliveBacklog = r.Config.Cache.Redis.KeepaliveBacklog.ValueInt64()
				} else {
					keepaliveBacklog = nil
				}
				keepalivePoolSize := new(int64)
				if !r.Config.Cache.Redis.KeepalivePoolSize.IsUnknown() && !r.Config.Cache.Redis.KeepalivePoolSize.IsNull() {
					*keepalivePoolSize = r.Config.Cache.Redis.KeepalivePoolSize.ValueInt64()
				} else {
					keepalivePoolSize = nil
				}
				password := new(string)
				if !r.Config.Cache.Redis.Password.IsUnknown() && !r.Config.Cache.Redis.Password.IsNull() {
					*password = r.Config.Cache.Redis.Password.ValueString()
				} else {
					password = nil
				}
				port1 := new(int64)
				if !r.Config.Cache.Redis.Port.IsUnknown() && !r.Config.Cache.Redis.Port.IsNull() {
					*port1 = r.Config.Cache.Redis.Port.ValueInt64()
				} else {
					port1 = nil
				}
				readTimeout := new(int64)
				if !r.Config.Cache.Redis.ReadTimeout.IsUnknown() && !r.Config.Cache.Redis.ReadTimeout.IsNull() {
					*readTimeout = r.Config.Cache.Redis.ReadTimeout.ValueInt64()
				} else {
					readTimeout = nil
				}
				sendTimeout := new(int64)
				if !r.Config.Cache.Redis.SendTimeout.IsUnknown() && !r.Config.Cache.Redis.SendTimeout.IsNull() {
					*sendTimeout = r.Config.Cache.Redis.SendTimeout.ValueInt64()
				} else {
					sendTimeout = nil
				}
				sentinelMaster := new(string)
				if !r.Config.Cache.Redis.SentinelMaster.IsUnknown() && !r.Config.Cache.Redis.SentinelMaster.IsNull() {
					*sentinelMaster = r.Config.Cache.Redis.SentinelMaster.ValueString()
				} else {
					sentinelMaster = nil
				}
				var sentinelNodes []shared.UpstreamOauthPluginSentinelNodes = []shared.UpstreamOauthPluginSentinelNodes{}
				for _, sentinelNodesItem := range r.Config.Cache.Redis.SentinelNodes {
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
					sentinelNodes = append(sentinelNodes, shared.UpstreamOauthPluginSentinelNodes{
						Host: host1,
						Port: port2,
					})
				}
				sentinelPassword := new(string)
				if !r.Config.Cache.Redis.SentinelPassword.IsUnknown() && !r.Config.Cache.Redis.SentinelPassword.IsNull() {
					*sentinelPassword = r.Config.Cache.Redis.SentinelPassword.ValueString()
				} else {
					sentinelPassword = nil
				}
				sentinelRole := new(shared.UpstreamOauthPluginSentinelRole)
				if !r.Config.Cache.Redis.SentinelRole.IsUnknown() && !r.Config.Cache.Redis.SentinelRole.IsNull() {
					*sentinelRole = shared.UpstreamOauthPluginSentinelRole(r.Config.Cache.Redis.SentinelRole.ValueString())
				} else {
					sentinelRole = nil
				}
				sentinelUsername := new(string)
				if !r.Config.Cache.Redis.SentinelUsername.IsUnknown() && !r.Config.Cache.Redis.SentinelUsername.IsNull() {
					*sentinelUsername = r.Config.Cache.Redis.SentinelUsername.ValueString()
				} else {
					sentinelUsername = nil
				}
				serverName := new(string)
				if !r.Config.Cache.Redis.ServerName.IsUnknown() && !r.Config.Cache.Redis.ServerName.IsNull() {
					*serverName = r.Config.Cache.Redis.ServerName.ValueString()
				} else {
					serverName = nil
				}
				ssl := new(bool)
				if !r.Config.Cache.Redis.Ssl.IsUnknown() && !r.Config.Cache.Redis.Ssl.IsNull() {
					*ssl = r.Config.Cache.Redis.Ssl.ValueBool()
				} else {
					ssl = nil
				}
				sslVerify := new(bool)
				if !r.Config.Cache.Redis.SslVerify.IsUnknown() && !r.Config.Cache.Redis.SslVerify.IsNull() {
					*sslVerify = r.Config.Cache.Redis.SslVerify.ValueBool()
				} else {
					sslVerify = nil
				}
				username := new(string)
				if !r.Config.Cache.Redis.Username.IsUnknown() && !r.Config.Cache.Redis.Username.IsNull() {
					*username = r.Config.Cache.Redis.Username.ValueString()
				} else {
					username = nil
				}
				redis = &shared.UpstreamOauthPluginRedis{
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
			strategy := new(shared.UpstreamOauthPluginStrategy)
			if !r.Config.Cache.Strategy.IsUnknown() && !r.Config.Cache.Strategy.IsNull() {
				*strategy = shared.UpstreamOauthPluginStrategy(r.Config.Cache.Strategy.ValueString())
			} else {
				strategy = nil
			}
			cache = &shared.Cache{
				DefaultTTL:    defaultTTL,
				EagerlyExpire: eagerlyExpire,
				Memory:        memory,
				Redis:         redis,
				Strategy:      strategy,
			}
		}
		var client *shared.Client
		if r.Config.Client != nil {
			authMethod := new(shared.AuthMethod)
			if !r.Config.Client.AuthMethod.IsUnknown() && !r.Config.Client.AuthMethod.IsNull() {
				*authMethod = shared.AuthMethod(r.Config.Client.AuthMethod.ValueString())
			} else {
				authMethod = nil
			}
			clientSecretJwtAlg := new(shared.ClientSecretJwtAlg)
			if !r.Config.Client.ClientSecretJwtAlg.IsUnknown() && !r.Config.Client.ClientSecretJwtAlg.IsNull() {
				*clientSecretJwtAlg = shared.ClientSecretJwtAlg(r.Config.Client.ClientSecretJwtAlg.ValueString())
			} else {
				clientSecretJwtAlg = nil
			}
			httpProxy := new(string)
			if !r.Config.Client.HTTPProxy.IsUnknown() && !r.Config.Client.HTTPProxy.IsNull() {
				*httpProxy = r.Config.Client.HTTPProxy.ValueString()
			} else {
				httpProxy = nil
			}
			httpProxyAuthorization := new(string)
			if !r.Config.Client.HTTPProxyAuthorization.IsUnknown() && !r.Config.Client.HTTPProxyAuthorization.IsNull() {
				*httpProxyAuthorization = r.Config.Client.HTTPProxyAuthorization.ValueString()
			} else {
				httpProxyAuthorization = nil
			}
			httpVersion := new(float64)
			if !r.Config.Client.HTTPVersion.IsUnknown() && !r.Config.Client.HTTPVersion.IsNull() {
				*httpVersion = r.Config.Client.HTTPVersion.ValueFloat64()
			} else {
				httpVersion = nil
			}
			httpsProxy := new(string)
			if !r.Config.Client.HTTPSProxy.IsUnknown() && !r.Config.Client.HTTPSProxy.IsNull() {
				*httpsProxy = r.Config.Client.HTTPSProxy.ValueString()
			} else {
				httpsProxy = nil
			}
			httpsProxyAuthorization := new(string)
			if !r.Config.Client.HTTPSProxyAuthorization.IsUnknown() && !r.Config.Client.HTTPSProxyAuthorization.IsNull() {
				*httpsProxyAuthorization = r.Config.Client.HTTPSProxyAuthorization.ValueString()
			} else {
				httpsProxyAuthorization = nil
			}
			keepAlive := new(bool)
			if !r.Config.Client.KeepAlive.IsUnknown() && !r.Config.Client.KeepAlive.IsNull() {
				*keepAlive = r.Config.Client.KeepAlive.ValueBool()
			} else {
				keepAlive = nil
			}
			noProxy := new(string)
			if !r.Config.Client.NoProxy.IsUnknown() && !r.Config.Client.NoProxy.IsNull() {
				*noProxy = r.Config.Client.NoProxy.ValueString()
			} else {
				noProxy = nil
			}
			sslVerify1 := new(bool)
			if !r.Config.Client.SslVerify.IsUnknown() && !r.Config.Client.SslVerify.IsNull() {
				*sslVerify1 = r.Config.Client.SslVerify.ValueBool()
			} else {
				sslVerify1 = nil
			}
			timeout := new(int64)
			if !r.Config.Client.Timeout.IsUnknown() && !r.Config.Client.Timeout.IsNull() {
				*timeout = r.Config.Client.Timeout.ValueInt64()
			} else {
				timeout = nil
			}
			client = &shared.Client{
				AuthMethod:              authMethod,
				ClientSecretJwtAlg:      clientSecretJwtAlg,
				HTTPProxy:               httpProxy,
				HTTPProxyAuthorization:  httpProxyAuthorization,
				HTTPVersion:             httpVersion,
				HTTPSProxy:              httpsProxy,
				HTTPSProxyAuthorization: httpsProxyAuthorization,
				KeepAlive:               keepAlive,
				NoProxy:                 noProxy,
				SslVerify:               sslVerify1,
				Timeout:                 timeout,
			}
		}
		var oauth *shared.Oauth
		if r.Config.Oauth != nil {
			var audience []string = []string{}
			for _, audienceItem := range r.Config.Oauth.Audience {
				audience = append(audience, audienceItem.ValueString())
			}
			clientID := new(string)
			if !r.Config.Oauth.ClientID.IsUnknown() && !r.Config.Oauth.ClientID.IsNull() {
				*clientID = r.Config.Oauth.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Config.Oauth.ClientSecret.IsUnknown() && !r.Config.Oauth.ClientSecret.IsNull() {
				*clientSecret = r.Config.Oauth.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			grantType := new(shared.GrantType)
			if !r.Config.Oauth.GrantType.IsUnknown() && !r.Config.Oauth.GrantType.IsNull() {
				*grantType = shared.GrantType(r.Config.Oauth.GrantType.ValueString())
			} else {
				grantType = nil
			}
			password1 := new(string)
			if !r.Config.Oauth.Password.IsUnknown() && !r.Config.Oauth.Password.IsNull() {
				*password1 = r.Config.Oauth.Password.ValueString()
			} else {
				password1 = nil
			}
			var scopes []string = []string{}
			for _, scopesItem := range r.Config.Oauth.Scopes {
				scopes = append(scopes, scopesItem.ValueString())
			}
			tokenEndpoint := new(string)
			if !r.Config.Oauth.TokenEndpoint.IsUnknown() && !r.Config.Oauth.TokenEndpoint.IsNull() {
				*tokenEndpoint = r.Config.Oauth.TokenEndpoint.ValueString()
			} else {
				tokenEndpoint = nil
			}
			tokenHeaders := make(map[string]interface{})
			for tokenHeadersKey, tokenHeadersValue := range r.Config.Oauth.TokenHeaders {
				var tokenHeadersInst interface{}
				_ = json.Unmarshal([]byte(tokenHeadersValue.ValueString()), &tokenHeadersInst)
				tokenHeaders[tokenHeadersKey] = tokenHeadersInst
			}
			tokenPostArgs := make(map[string]interface{})
			for tokenPostArgsKey, tokenPostArgsValue := range r.Config.Oauth.TokenPostArgs {
				var tokenPostArgsInst interface{}
				_ = json.Unmarshal([]byte(tokenPostArgsValue.ValueString()), &tokenPostArgsInst)
				tokenPostArgs[tokenPostArgsKey] = tokenPostArgsInst
			}
			username1 := new(string)
			if !r.Config.Oauth.Username.IsUnknown() && !r.Config.Oauth.Username.IsNull() {
				*username1 = r.Config.Oauth.Username.ValueString()
			} else {
				username1 = nil
			}
			oauth = &shared.Oauth{
				Audience:      audience,
				ClientID:      clientID,
				ClientSecret:  clientSecret,
				GrantType:     grantType,
				Password:      password1,
				Scopes:        scopes,
				TokenEndpoint: tokenEndpoint,
				TokenHeaders:  tokenHeaders,
				TokenPostArgs: tokenPostArgs,
				Username:      username1,
			}
		}
		config = &shared.UpstreamOauthPluginConfig{
			Behavior: behavior,
			Cache:    cache,
			Client:   client,
			Oauth:    oauth,
		}
	}
	var consumer *shared.UpstreamOauthPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.UpstreamOauthPluginConsumer{
			ID: id1,
		}
	}
	var consumerGroup *shared.UpstreamOauthPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id2 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id2 = r.ConsumerGroup.ID.ValueString()
		} else {
			id2 = nil
		}
		consumerGroup = &shared.UpstreamOauthPluginConsumerGroup{
			ID: id2,
		}
	}
	var protocols []shared.UpstreamOauthPluginProtocols = []shared.UpstreamOauthPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.UpstreamOauthPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.UpstreamOauthPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.UpstreamOauthPluginRoute{
			ID: id3,
		}
	}
	var service *shared.UpstreamOauthPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.UpstreamOauthPluginService{
			ID: id4,
		}
	}
	out := shared.UpstreamOauthPlugin{
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

func (r *GatewayPluginUpstreamOauthResourceModel) RefreshFromSharedUpstreamOauthPlugin(ctx context.Context, resp *shared.UpstreamOauthPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.UpstreamOauthPluginConfig{}
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
				r.Config.Cache.DefaultTTL = types.Float64PointerValue(resp.Config.Cache.DefaultTTL)
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
						var clusterNodes tfTypes.AiProxyAdvancedPluginClusterNodes
						clusterNodes.IP = types.StringPointerValue(clusterNodesItem.IP)
						clusterNodes.Port = types.Int64PointerValue(clusterNodesItem.Port)
						if clusterNodesCount+1 > len(r.Config.Cache.Redis.ClusterNodes) {
							r.Config.Cache.Redis.ClusterNodes = append(r.Config.Cache.Redis.ClusterNodes, clusterNodes)
						} else {
							r.Config.Cache.Redis.ClusterNodes[clusterNodesCount].IP = clusterNodes.IP
							r.Config.Cache.Redis.ClusterNodes[clusterNodesCount].Port = clusterNodes.Port
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
						var sentinelNodes tfTypes.AiProxyAdvancedPluginSentinelNodes
						sentinelNodes.Host = types.StringPointerValue(sentinelNodesItem.Host)
						sentinelNodes.Port = types.Int64PointerValue(sentinelNodesItem.Port)
						if sentinelNodesCount+1 > len(r.Config.Cache.Redis.SentinelNodes) {
							r.Config.Cache.Redis.SentinelNodes = append(r.Config.Cache.Redis.SentinelNodes, sentinelNodes)
						} else {
							r.Config.Cache.Redis.SentinelNodes[sentinelNodesCount].Host = sentinelNodes.Host
							r.Config.Cache.Redis.SentinelNodes[sentinelNodesCount].Port = sentinelNodes.Port
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
				r.Config.Client.HTTPVersion = types.Float64PointerValue(resp.Config.Client.HTTPVersion)
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
				if resp.Config.Oauth.Scopes != nil {
					r.Config.Oauth.Scopes = make([]types.String, 0, len(resp.Config.Oauth.Scopes))
					for _, v := range resp.Config.Oauth.Scopes {
						r.Config.Oauth.Scopes = append(r.Config.Oauth.Scopes, types.StringValue(v))
					}
				}
				r.Config.Oauth.TokenEndpoint = types.StringPointerValue(resp.Config.Oauth.TokenEndpoint)
				if len(resp.Config.Oauth.TokenHeaders) > 0 {
					r.Config.Oauth.TokenHeaders = make(map[string]types.String, len(resp.Config.Oauth.TokenHeaders))
					for key, value := range resp.Config.Oauth.TokenHeaders {
						result, _ := json.Marshal(value)
						r.Config.Oauth.TokenHeaders[key] = types.StringValue(string(result))
					}
				}
				if len(resp.Config.Oauth.TokenPostArgs) > 0 {
					r.Config.Oauth.TokenPostArgs = make(map[string]types.String, len(resp.Config.Oauth.TokenPostArgs))
					for key1, value1 := range resp.Config.Oauth.TokenPostArgs {
						result1, _ := json.Marshal(value1)
						r.Config.Oauth.TokenPostArgs[key1] = types.StringValue(string(result1))
					}
				}
				r.Config.Oauth.Username = types.StringPointerValue(resp.Config.Oauth.Username)
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
