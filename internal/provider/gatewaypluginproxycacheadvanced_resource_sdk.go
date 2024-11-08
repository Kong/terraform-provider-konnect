// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginProxyCacheAdvancedResourceModel) ToSharedProxyCacheAdvancedPluginInput() *shared.ProxyCacheAdvancedPluginInput {
	bypassOnErr := new(bool)
	if !r.Config.BypassOnErr.IsUnknown() && !r.Config.BypassOnErr.IsNull() {
		*bypassOnErr = r.Config.BypassOnErr.ValueBool()
	} else {
		bypassOnErr = nil
	}
	cacheControl := new(bool)
	if !r.Config.CacheControl.IsUnknown() && !r.Config.CacheControl.IsNull() {
		*cacheControl = r.Config.CacheControl.ValueBool()
	} else {
		cacheControl = nil
	}
	cacheTTL := new(int64)
	if !r.Config.CacheTTL.IsUnknown() && !r.Config.CacheTTL.IsNull() {
		*cacheTTL = r.Config.CacheTTL.ValueInt64()
	} else {
		cacheTTL = nil
	}
	var contentType []string = []string{}
	for _, contentTypeItem := range r.Config.ContentType {
		contentType = append(contentType, contentTypeItem.ValueString())
	}
	ignoreURICase := new(bool)
	if !r.Config.IgnoreURICase.IsUnknown() && !r.Config.IgnoreURICase.IsNull() {
		*ignoreURICase = r.Config.IgnoreURICase.ValueBool()
	} else {
		ignoreURICase = nil
	}
	var memory *shared.ProxyCacheAdvancedPluginMemory
	if r.Config.Memory != nil {
		dictionaryName := new(string)
		if !r.Config.Memory.DictionaryName.IsUnknown() && !r.Config.Memory.DictionaryName.IsNull() {
			*dictionaryName = r.Config.Memory.DictionaryName.ValueString()
		} else {
			dictionaryName = nil
		}
		memory = &shared.ProxyCacheAdvancedPluginMemory{
			DictionaryName: dictionaryName,
		}
	}
	var redis *shared.ProxyCacheAdvancedPluginRedis
	if r.Config.Redis != nil {
		clusterMaxRedirections := new(int64)
		if !r.Config.Redis.ClusterMaxRedirections.IsUnknown() && !r.Config.Redis.ClusterMaxRedirections.IsNull() {
			*clusterMaxRedirections = r.Config.Redis.ClusterMaxRedirections.ValueInt64()
		} else {
			clusterMaxRedirections = nil
		}
		var clusterNodes []shared.ProxyCacheAdvancedPluginClusterNodes = []shared.ProxyCacheAdvancedPluginClusterNodes{}
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
			clusterNodes = append(clusterNodes, shared.ProxyCacheAdvancedPluginClusterNodes{
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
		var sentinelNodes []shared.ProxyCacheAdvancedPluginSentinelNodes = []shared.ProxyCacheAdvancedPluginSentinelNodes{}
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
			sentinelNodes = append(sentinelNodes, shared.ProxyCacheAdvancedPluginSentinelNodes{
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
		sentinelRole := new(shared.ProxyCacheAdvancedPluginSentinelRole)
		if !r.Config.Redis.SentinelRole.IsUnknown() && !r.Config.Redis.SentinelRole.IsNull() {
			*sentinelRole = shared.ProxyCacheAdvancedPluginSentinelRole(r.Config.Redis.SentinelRole.ValueString())
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
		redis = &shared.ProxyCacheAdvancedPluginRedis{
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
	var requestMethod []shared.ProxyCacheAdvancedPluginRequestMethod = []shared.ProxyCacheAdvancedPluginRequestMethod{}
	for _, requestMethodItem := range r.Config.RequestMethod {
		requestMethod = append(requestMethod, shared.ProxyCacheAdvancedPluginRequestMethod(requestMethodItem.ValueString()))
	}
	var responseCode []int64 = []int64{}
	for _, responseCodeItem := range r.Config.ResponseCode {
		responseCode = append(responseCode, responseCodeItem.ValueInt64())
	}
	var responseHeaders *shared.ProxyCacheAdvancedPluginResponseHeaders
	if r.Config.ResponseHeaders != nil {
		xCacheKey := new(bool)
		if !r.Config.ResponseHeaders.XCacheKey.IsUnknown() && !r.Config.ResponseHeaders.XCacheKey.IsNull() {
			*xCacheKey = r.Config.ResponseHeaders.XCacheKey.ValueBool()
		} else {
			xCacheKey = nil
		}
		xCacheStatus := new(bool)
		if !r.Config.ResponseHeaders.XCacheStatus.IsUnknown() && !r.Config.ResponseHeaders.XCacheStatus.IsNull() {
			*xCacheStatus = r.Config.ResponseHeaders.XCacheStatus.ValueBool()
		} else {
			xCacheStatus = nil
		}
		age := new(bool)
		if !r.Config.ResponseHeaders.Age.IsUnknown() && !r.Config.ResponseHeaders.Age.IsNull() {
			*age = r.Config.ResponseHeaders.Age.ValueBool()
		} else {
			age = nil
		}
		responseHeaders = &shared.ProxyCacheAdvancedPluginResponseHeaders{
			XCacheKey:    xCacheKey,
			XCacheStatus: xCacheStatus,
			Age:          age,
		}
	}
	storageTTL := new(int64)
	if !r.Config.StorageTTL.IsUnknown() && !r.Config.StorageTTL.IsNull() {
		*storageTTL = r.Config.StorageTTL.ValueInt64()
	} else {
		storageTTL = nil
	}
	strategy := new(shared.ProxyCacheAdvancedPluginStrategy)
	if !r.Config.Strategy.IsUnknown() && !r.Config.Strategy.IsNull() {
		*strategy = shared.ProxyCacheAdvancedPluginStrategy(r.Config.Strategy.ValueString())
	} else {
		strategy = nil
	}
	var varyHeaders []string = []string{}
	for _, varyHeadersItem := range r.Config.VaryHeaders {
		varyHeaders = append(varyHeaders, varyHeadersItem.ValueString())
	}
	var varyQueryParams []string = []string{}
	for _, varyQueryParamsItem := range r.Config.VaryQueryParams {
		varyQueryParams = append(varyQueryParams, varyQueryParamsItem.ValueString())
	}
	config := shared.ProxyCacheAdvancedPluginConfig{
		BypassOnErr:     bypassOnErr,
		CacheControl:    cacheControl,
		CacheTTL:        cacheTTL,
		ContentType:     contentType,
		IgnoreURICase:   ignoreURICase,
		Memory:          memory,
		Redis:           redis,
		RequestMethod:   requestMethod,
		ResponseCode:    responseCode,
		ResponseHeaders: responseHeaders,
		StorageTTL:      storageTTL,
		Strategy:        strategy,
		VaryHeaders:     varyHeaders,
		VaryQueryParams: varyQueryParams,
	}
	var consumer *shared.ProxyCacheAdvancedPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.ProxyCacheAdvancedPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.ProxyCacheAdvancedPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.ProxyCacheAdvancedPluginConsumerGroup{
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
	var ordering *shared.ProxyCacheAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.ProxyCacheAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ProxyCacheAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.ProxyCacheAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ProxyCacheAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ProxyCacheAdvancedPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.ProxyCacheAdvancedPluginProtocols = []shared.ProxyCacheAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ProxyCacheAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ProxyCacheAdvancedPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.ProxyCacheAdvancedPluginRoute{
			ID: id3,
		}
	}
	var service *shared.ProxyCacheAdvancedPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.ProxyCacheAdvancedPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.ProxyCacheAdvancedPluginInput{
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

func (r *GatewayPluginProxyCacheAdvancedResourceModel) RefreshFromSharedProxyCacheAdvancedPlugin(resp *shared.ProxyCacheAdvancedPlugin) {
	if resp != nil {
		r.Config.BypassOnErr = types.BoolPointerValue(resp.Config.BypassOnErr)
		r.Config.CacheControl = types.BoolPointerValue(resp.Config.CacheControl)
		r.Config.CacheTTL = types.Int64PointerValue(resp.Config.CacheTTL)
		r.Config.ContentType = []types.String{}
		for _, v := range resp.Config.ContentType {
			r.Config.ContentType = append(r.Config.ContentType, types.StringValue(v))
		}
		r.Config.IgnoreURICase = types.BoolPointerValue(resp.Config.IgnoreURICase)
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
		r.Config.RequestMethod = []types.String{}
		for _, v := range resp.Config.RequestMethod {
			r.Config.RequestMethod = append(r.Config.RequestMethod, types.StringValue(string(v)))
		}
		r.Config.ResponseCode = []types.Int64{}
		for _, v := range resp.Config.ResponseCode {
			r.Config.ResponseCode = append(r.Config.ResponseCode, types.Int64Value(v))
		}
		if resp.Config.ResponseHeaders == nil {
			r.Config.ResponseHeaders = nil
		} else {
			r.Config.ResponseHeaders = &tfTypes.ResponseHeaders{}
			r.Config.ResponseHeaders.Age = types.BoolPointerValue(resp.Config.ResponseHeaders.Age)
			r.Config.ResponseHeaders.XCacheKey = types.BoolPointerValue(resp.Config.ResponseHeaders.XCacheKey)
			r.Config.ResponseHeaders.XCacheStatus = types.BoolPointerValue(resp.Config.ResponseHeaders.XCacheStatus)
		}
		r.Config.StorageTTL = types.Int64PointerValue(resp.Config.StorageTTL)
		if resp.Config.Strategy != nil {
			r.Config.Strategy = types.StringValue(string(*resp.Config.Strategy))
		} else {
			r.Config.Strategy = types.StringNull()
		}
		r.Config.VaryHeaders = []types.String{}
		for _, v := range resp.Config.VaryHeaders {
			r.Config.VaryHeaders = append(r.Config.VaryHeaders, types.StringValue(v))
		}
		r.Config.VaryQueryParams = []types.String{}
		for _, v := range resp.Config.VaryQueryParams {
			r.Config.VaryQueryParams = append(r.Config.VaryQueryParams, types.StringValue(v))
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
