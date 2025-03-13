// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginSamlResourceModel) ToSharedSamlPluginInput() *shared.SamlPluginInput {
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
	var ordering *shared.SamlPluginOrdering
	if r.Ordering != nil {
		var after *shared.SamlPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.SamlPluginAfter{
				Access: access,
			}
		}
		var before *shared.SamlPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.SamlPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.SamlPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	anonymous := new(string)
	if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
		*anonymous = r.Config.Anonymous.ValueString()
	} else {
		anonymous = nil
	}
	assertionConsumerPath := new(string)
	if !r.Config.AssertionConsumerPath.IsUnknown() && !r.Config.AssertionConsumerPath.IsNull() {
		*assertionConsumerPath = r.Config.AssertionConsumerPath.ValueString()
	} else {
		assertionConsumerPath = nil
	}
	idpCertificate := new(string)
	if !r.Config.IdpCertificate.IsUnknown() && !r.Config.IdpCertificate.IsNull() {
		*idpCertificate = r.Config.IdpCertificate.ValueString()
	} else {
		idpCertificate = nil
	}
	idpSsoURL := new(string)
	if !r.Config.IdpSsoURL.IsUnknown() && !r.Config.IdpSsoURL.IsNull() {
		*idpSsoURL = r.Config.IdpSsoURL.ValueString()
	} else {
		idpSsoURL = nil
	}
	issuer := new(string)
	if !r.Config.Issuer.IsUnknown() && !r.Config.Issuer.IsNull() {
		*issuer = r.Config.Issuer.ValueString()
	} else {
		issuer = nil
	}
	nameidFormat := new(shared.NameidFormat)
	if !r.Config.NameidFormat.IsUnknown() && !r.Config.NameidFormat.IsNull() {
		*nameidFormat = shared.NameidFormat(r.Config.NameidFormat.ValueString())
	} else {
		nameidFormat = nil
	}
	var redis *shared.SamlPluginRedis
	if r.Config.Redis != nil {
		clusterMaxRedirections := new(int64)
		if !r.Config.Redis.ClusterMaxRedirections.IsUnknown() && !r.Config.Redis.ClusterMaxRedirections.IsNull() {
			*clusterMaxRedirections = r.Config.Redis.ClusterMaxRedirections.ValueInt64()
		} else {
			clusterMaxRedirections = nil
		}
		var clusterNodes []shared.SamlPluginClusterNodes = []shared.SamlPluginClusterNodes{}
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
			clusterNodes = append(clusterNodes, shared.SamlPluginClusterNodes{
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
		prefix := new(string)
		if !r.Config.Redis.Prefix.IsUnknown() && !r.Config.Redis.Prefix.IsNull() {
			*prefix = r.Config.Redis.Prefix.ValueString()
		} else {
			prefix = nil
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
		var sentinelNodes []shared.SamlPluginSentinelNodes = []shared.SamlPluginSentinelNodes{}
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
			sentinelNodes = append(sentinelNodes, shared.SamlPluginSentinelNodes{
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
		sentinelRole := new(shared.SamlPluginSentinelRole)
		if !r.Config.Redis.SentinelRole.IsUnknown() && !r.Config.Redis.SentinelRole.IsNull() {
			*sentinelRole = shared.SamlPluginSentinelRole(r.Config.Redis.SentinelRole.ValueString())
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
		socket := new(string)
		if !r.Config.Redis.Socket.IsUnknown() && !r.Config.Redis.Socket.IsNull() {
			*socket = r.Config.Redis.Socket.ValueString()
		} else {
			socket = nil
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
		redis = &shared.SamlPluginRedis{
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
			Prefix:                 prefix,
			ReadTimeout:            readTimeout,
			SendTimeout:            sendTimeout,
			SentinelMaster:         sentinelMaster,
			SentinelNodes:          sentinelNodes,
			SentinelPassword:       sentinelPassword,
			SentinelRole:           sentinelRole,
			SentinelUsername:       sentinelUsername,
			ServerName:             serverName,
			Socket:                 socket,
			Ssl:                    ssl,
			SslVerify:              sslVerify,
			Username:               username,
		}
	}
	requestDigestAlgorithm := new(shared.RequestDigestAlgorithm)
	if !r.Config.RequestDigestAlgorithm.IsUnknown() && !r.Config.RequestDigestAlgorithm.IsNull() {
		*requestDigestAlgorithm = shared.RequestDigestAlgorithm(r.Config.RequestDigestAlgorithm.ValueString())
	} else {
		requestDigestAlgorithm = nil
	}
	requestSignatureAlgorithm := new(shared.RequestSignatureAlgorithm)
	if !r.Config.RequestSignatureAlgorithm.IsUnknown() && !r.Config.RequestSignatureAlgorithm.IsNull() {
		*requestSignatureAlgorithm = shared.RequestSignatureAlgorithm(r.Config.RequestSignatureAlgorithm.ValueString())
	} else {
		requestSignatureAlgorithm = nil
	}
	requestSigningCertificate := new(string)
	if !r.Config.RequestSigningCertificate.IsUnknown() && !r.Config.RequestSigningCertificate.IsNull() {
		*requestSigningCertificate = r.Config.RequestSigningCertificate.ValueString()
	} else {
		requestSigningCertificate = nil
	}
	requestSigningKey := new(string)
	if !r.Config.RequestSigningKey.IsUnknown() && !r.Config.RequestSigningKey.IsNull() {
		*requestSigningKey = r.Config.RequestSigningKey.ValueString()
	} else {
		requestSigningKey = nil
	}
	responseDigestAlgorithm := new(shared.ResponseDigestAlgorithm)
	if !r.Config.ResponseDigestAlgorithm.IsUnknown() && !r.Config.ResponseDigestAlgorithm.IsNull() {
		*responseDigestAlgorithm = shared.ResponseDigestAlgorithm(r.Config.ResponseDigestAlgorithm.ValueString())
	} else {
		responseDigestAlgorithm = nil
	}
	responseEncryptionKey := new(string)
	if !r.Config.ResponseEncryptionKey.IsUnknown() && !r.Config.ResponseEncryptionKey.IsNull() {
		*responseEncryptionKey = r.Config.ResponseEncryptionKey.ValueString()
	} else {
		responseEncryptionKey = nil
	}
	responseSignatureAlgorithm := new(shared.ResponseSignatureAlgorithm)
	if !r.Config.ResponseSignatureAlgorithm.IsUnknown() && !r.Config.ResponseSignatureAlgorithm.IsNull() {
		*responseSignatureAlgorithm = shared.ResponseSignatureAlgorithm(r.Config.ResponseSignatureAlgorithm.ValueString())
	} else {
		responseSignatureAlgorithm = nil
	}
	sessionAbsoluteTimeout := new(float64)
	if !r.Config.SessionAbsoluteTimeout.IsUnknown() && !r.Config.SessionAbsoluteTimeout.IsNull() {
		*sessionAbsoluteTimeout, _ = r.Config.SessionAbsoluteTimeout.ValueBigFloat().Float64()
	} else {
		sessionAbsoluteTimeout = nil
	}
	sessionAudience := new(string)
	if !r.Config.SessionAudience.IsUnknown() && !r.Config.SessionAudience.IsNull() {
		*sessionAudience = r.Config.SessionAudience.ValueString()
	} else {
		sessionAudience = nil
	}
	sessionCookieDomain := new(string)
	if !r.Config.SessionCookieDomain.IsUnknown() && !r.Config.SessionCookieDomain.IsNull() {
		*sessionCookieDomain = r.Config.SessionCookieDomain.ValueString()
	} else {
		sessionCookieDomain = nil
	}
	sessionCookieHTTPOnly := new(bool)
	if !r.Config.SessionCookieHTTPOnly.IsUnknown() && !r.Config.SessionCookieHTTPOnly.IsNull() {
		*sessionCookieHTTPOnly = r.Config.SessionCookieHTTPOnly.ValueBool()
	} else {
		sessionCookieHTTPOnly = nil
	}
	sessionCookieName := new(string)
	if !r.Config.SessionCookieName.IsUnknown() && !r.Config.SessionCookieName.IsNull() {
		*sessionCookieName = r.Config.SessionCookieName.ValueString()
	} else {
		sessionCookieName = nil
	}
	sessionCookiePath := new(string)
	if !r.Config.SessionCookiePath.IsUnknown() && !r.Config.SessionCookiePath.IsNull() {
		*sessionCookiePath = r.Config.SessionCookiePath.ValueString()
	} else {
		sessionCookiePath = nil
	}
	sessionCookieSameSite := new(shared.SamlPluginSessionCookieSameSite)
	if !r.Config.SessionCookieSameSite.IsUnknown() && !r.Config.SessionCookieSameSite.IsNull() {
		*sessionCookieSameSite = shared.SamlPluginSessionCookieSameSite(r.Config.SessionCookieSameSite.ValueString())
	} else {
		sessionCookieSameSite = nil
	}
	sessionCookieSecure := new(bool)
	if !r.Config.SessionCookieSecure.IsUnknown() && !r.Config.SessionCookieSecure.IsNull() {
		*sessionCookieSecure = r.Config.SessionCookieSecure.ValueBool()
	} else {
		sessionCookieSecure = nil
	}
	sessionEnforceSameSubject := new(bool)
	if !r.Config.SessionEnforceSameSubject.IsUnknown() && !r.Config.SessionEnforceSameSubject.IsNull() {
		*sessionEnforceSameSubject = r.Config.SessionEnforceSameSubject.ValueBool()
	} else {
		sessionEnforceSameSubject = nil
	}
	sessionHashStorageKey := new(bool)
	if !r.Config.SessionHashStorageKey.IsUnknown() && !r.Config.SessionHashStorageKey.IsNull() {
		*sessionHashStorageKey = r.Config.SessionHashStorageKey.ValueBool()
	} else {
		sessionHashStorageKey = nil
	}
	sessionHashSubject := new(bool)
	if !r.Config.SessionHashSubject.IsUnknown() && !r.Config.SessionHashSubject.IsNull() {
		*sessionHashSubject = r.Config.SessionHashSubject.ValueBool()
	} else {
		sessionHashSubject = nil
	}
	sessionIdlingTimeout := new(float64)
	if !r.Config.SessionIdlingTimeout.IsUnknown() && !r.Config.SessionIdlingTimeout.IsNull() {
		*sessionIdlingTimeout, _ = r.Config.SessionIdlingTimeout.ValueBigFloat().Float64()
	} else {
		sessionIdlingTimeout = nil
	}
	sessionMemcachedHost := new(string)
	if !r.Config.SessionMemcachedHost.IsUnknown() && !r.Config.SessionMemcachedHost.IsNull() {
		*sessionMemcachedHost = r.Config.SessionMemcachedHost.ValueString()
	} else {
		sessionMemcachedHost = nil
	}
	sessionMemcachedPort := new(int64)
	if !r.Config.SessionMemcachedPort.IsUnknown() && !r.Config.SessionMemcachedPort.IsNull() {
		*sessionMemcachedPort = r.Config.SessionMemcachedPort.ValueInt64()
	} else {
		sessionMemcachedPort = nil
	}
	sessionMemcachedPrefix := new(string)
	if !r.Config.SessionMemcachedPrefix.IsUnknown() && !r.Config.SessionMemcachedPrefix.IsNull() {
		*sessionMemcachedPrefix = r.Config.SessionMemcachedPrefix.ValueString()
	} else {
		sessionMemcachedPrefix = nil
	}
	sessionMemcachedSocket := new(string)
	if !r.Config.SessionMemcachedSocket.IsUnknown() && !r.Config.SessionMemcachedSocket.IsNull() {
		*sessionMemcachedSocket = r.Config.SessionMemcachedSocket.ValueString()
	} else {
		sessionMemcachedSocket = nil
	}
	sessionRemember := new(bool)
	if !r.Config.SessionRemember.IsUnknown() && !r.Config.SessionRemember.IsNull() {
		*sessionRemember = r.Config.SessionRemember.ValueBool()
	} else {
		sessionRemember = nil
	}
	sessionRememberAbsoluteTimeout := new(float64)
	if !r.Config.SessionRememberAbsoluteTimeout.IsUnknown() && !r.Config.SessionRememberAbsoluteTimeout.IsNull() {
		*sessionRememberAbsoluteTimeout, _ = r.Config.SessionRememberAbsoluteTimeout.ValueBigFloat().Float64()
	} else {
		sessionRememberAbsoluteTimeout = nil
	}
	sessionRememberCookieName := new(string)
	if !r.Config.SessionRememberCookieName.IsUnknown() && !r.Config.SessionRememberCookieName.IsNull() {
		*sessionRememberCookieName = r.Config.SessionRememberCookieName.ValueString()
	} else {
		sessionRememberCookieName = nil
	}
	sessionRememberRollingTimeout := new(float64)
	if !r.Config.SessionRememberRollingTimeout.IsUnknown() && !r.Config.SessionRememberRollingTimeout.IsNull() {
		*sessionRememberRollingTimeout, _ = r.Config.SessionRememberRollingTimeout.ValueBigFloat().Float64()
	} else {
		sessionRememberRollingTimeout = nil
	}
	var sessionRequestHeaders []shared.SamlPluginSessionRequestHeaders = []shared.SamlPluginSessionRequestHeaders{}
	for _, sessionRequestHeadersItem := range r.Config.SessionRequestHeaders {
		sessionRequestHeaders = append(sessionRequestHeaders, shared.SamlPluginSessionRequestHeaders(sessionRequestHeadersItem.ValueString()))
	}
	var sessionResponseHeaders []shared.SamlPluginSessionResponseHeaders = []shared.SamlPluginSessionResponseHeaders{}
	for _, sessionResponseHeadersItem := range r.Config.SessionResponseHeaders {
		sessionResponseHeaders = append(sessionResponseHeaders, shared.SamlPluginSessionResponseHeaders(sessionResponseHeadersItem.ValueString()))
	}
	sessionRollingTimeout := new(float64)
	if !r.Config.SessionRollingTimeout.IsUnknown() && !r.Config.SessionRollingTimeout.IsNull() {
		*sessionRollingTimeout, _ = r.Config.SessionRollingTimeout.ValueBigFloat().Float64()
	} else {
		sessionRollingTimeout = nil
	}
	sessionSecret := new(string)
	if !r.Config.SessionSecret.IsUnknown() && !r.Config.SessionSecret.IsNull() {
		*sessionSecret = r.Config.SessionSecret.ValueString()
	} else {
		sessionSecret = nil
	}
	sessionStorage := new(shared.SamlPluginSessionStorage)
	if !r.Config.SessionStorage.IsUnknown() && !r.Config.SessionStorage.IsNull() {
		*sessionStorage = shared.SamlPluginSessionStorage(r.Config.SessionStorage.ValueString())
	} else {
		sessionStorage = nil
	}
	sessionStoreMetadata := new(bool)
	if !r.Config.SessionStoreMetadata.IsUnknown() && !r.Config.SessionStoreMetadata.IsNull() {
		*sessionStoreMetadata = r.Config.SessionStoreMetadata.ValueBool()
	} else {
		sessionStoreMetadata = nil
	}
	validateAssertionSignature := new(bool)
	if !r.Config.ValidateAssertionSignature.IsUnknown() && !r.Config.ValidateAssertionSignature.IsNull() {
		*validateAssertionSignature = r.Config.ValidateAssertionSignature.ValueBool()
	} else {
		validateAssertionSignature = nil
	}
	config := shared.SamlPluginConfig{
		Anonymous:                      anonymous,
		AssertionConsumerPath:          assertionConsumerPath,
		IdpCertificate:                 idpCertificate,
		IdpSsoURL:                      idpSsoURL,
		Issuer:                         issuer,
		NameidFormat:                   nameidFormat,
		Redis:                          redis,
		RequestDigestAlgorithm:         requestDigestAlgorithm,
		RequestSignatureAlgorithm:      requestSignatureAlgorithm,
		RequestSigningCertificate:      requestSigningCertificate,
		RequestSigningKey:              requestSigningKey,
		ResponseDigestAlgorithm:        responseDigestAlgorithm,
		ResponseEncryptionKey:          responseEncryptionKey,
		ResponseSignatureAlgorithm:     responseSignatureAlgorithm,
		SessionAbsoluteTimeout:         sessionAbsoluteTimeout,
		SessionAudience:                sessionAudience,
		SessionCookieDomain:            sessionCookieDomain,
		SessionCookieHTTPOnly:          sessionCookieHTTPOnly,
		SessionCookieName:              sessionCookieName,
		SessionCookiePath:              sessionCookiePath,
		SessionCookieSameSite:          sessionCookieSameSite,
		SessionCookieSecure:            sessionCookieSecure,
		SessionEnforceSameSubject:      sessionEnforceSameSubject,
		SessionHashStorageKey:          sessionHashStorageKey,
		SessionHashSubject:             sessionHashSubject,
		SessionIdlingTimeout:           sessionIdlingTimeout,
		SessionMemcachedHost:           sessionMemcachedHost,
		SessionMemcachedPort:           sessionMemcachedPort,
		SessionMemcachedPrefix:         sessionMemcachedPrefix,
		SessionMemcachedSocket:         sessionMemcachedSocket,
		SessionRemember:                sessionRemember,
		SessionRememberAbsoluteTimeout: sessionRememberAbsoluteTimeout,
		SessionRememberCookieName:      sessionRememberCookieName,
		SessionRememberRollingTimeout:  sessionRememberRollingTimeout,
		SessionRequestHeaders:          sessionRequestHeaders,
		SessionResponseHeaders:         sessionResponseHeaders,
		SessionRollingTimeout:          sessionRollingTimeout,
		SessionSecret:                  sessionSecret,
		SessionStorage:                 sessionStorage,
		SessionStoreMetadata:           sessionStoreMetadata,
		ValidateAssertionSignature:     validateAssertionSignature,
	}
	var protocols []shared.SamlPluginProtocols = []shared.SamlPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.SamlPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.SamlPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.SamlPluginRoute{
			ID: id1,
		}
	}
	var service *shared.SamlPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.SamlPluginService{
			ID: id2,
		}
	}
	out := shared.SamlPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginSamlResourceModel) RefreshFromSharedSamlPlugin(resp *shared.SamlPlugin) {
	if resp != nil {
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		r.Config.AssertionConsumerPath = types.StringPointerValue(resp.Config.AssertionConsumerPath)
		r.Config.IdpCertificate = types.StringPointerValue(resp.Config.IdpCertificate)
		r.Config.IdpSsoURL = types.StringPointerValue(resp.Config.IdpSsoURL)
		r.Config.Issuer = types.StringPointerValue(resp.Config.Issuer)
		if resp.Config.NameidFormat != nil {
			r.Config.NameidFormat = types.StringValue(string(*resp.Config.NameidFormat))
		} else {
			r.Config.NameidFormat = types.StringNull()
		}
		if resp.Config.Redis == nil {
			r.Config.Redis = nil
		} else {
			r.Config.Redis = &tfTypes.OpenidConnectPluginRedis{}
			r.Config.Redis.ClusterMaxRedirections = types.Int64PointerValue(resp.Config.Redis.ClusterMaxRedirections)
			r.Config.Redis.ClusterNodes = []tfTypes.AiProxyAdvancedPluginClusterNodes{}
			if len(r.Config.Redis.ClusterNodes) > len(resp.Config.Redis.ClusterNodes) {
				r.Config.Redis.ClusterNodes = r.Config.Redis.ClusterNodes[:len(resp.Config.Redis.ClusterNodes)]
			}
			for clusterNodesCount, clusterNodesItem := range resp.Config.Redis.ClusterNodes {
				var clusterNodes1 tfTypes.AiProxyAdvancedPluginClusterNodes
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
			r.Config.Redis.Prefix = types.StringPointerValue(resp.Config.Redis.Prefix)
			r.Config.Redis.ReadTimeout = types.Int64PointerValue(resp.Config.Redis.ReadTimeout)
			r.Config.Redis.SendTimeout = types.Int64PointerValue(resp.Config.Redis.SendTimeout)
			r.Config.Redis.SentinelMaster = types.StringPointerValue(resp.Config.Redis.SentinelMaster)
			r.Config.Redis.SentinelNodes = []tfTypes.AiProxyAdvancedPluginSentinelNodes{}
			if len(r.Config.Redis.SentinelNodes) > len(resp.Config.Redis.SentinelNodes) {
				r.Config.Redis.SentinelNodes = r.Config.Redis.SentinelNodes[:len(resp.Config.Redis.SentinelNodes)]
			}
			for sentinelNodesCount, sentinelNodesItem := range resp.Config.Redis.SentinelNodes {
				var sentinelNodes1 tfTypes.AiProxyAdvancedPluginSentinelNodes
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
			r.Config.Redis.Socket = types.StringPointerValue(resp.Config.Redis.Socket)
			r.Config.Redis.Ssl = types.BoolPointerValue(resp.Config.Redis.Ssl)
			r.Config.Redis.SslVerify = types.BoolPointerValue(resp.Config.Redis.SslVerify)
			r.Config.Redis.Username = types.StringPointerValue(resp.Config.Redis.Username)
		}
		if resp.Config.RequestDigestAlgorithm != nil {
			r.Config.RequestDigestAlgorithm = types.StringValue(string(*resp.Config.RequestDigestAlgorithm))
		} else {
			r.Config.RequestDigestAlgorithm = types.StringNull()
		}
		if resp.Config.RequestSignatureAlgorithm != nil {
			r.Config.RequestSignatureAlgorithm = types.StringValue(string(*resp.Config.RequestSignatureAlgorithm))
		} else {
			r.Config.RequestSignatureAlgorithm = types.StringNull()
		}
		r.Config.RequestSigningCertificate = types.StringPointerValue(resp.Config.RequestSigningCertificate)
		r.Config.RequestSigningKey = types.StringPointerValue(resp.Config.RequestSigningKey)
		if resp.Config.ResponseDigestAlgorithm != nil {
			r.Config.ResponseDigestAlgorithm = types.StringValue(string(*resp.Config.ResponseDigestAlgorithm))
		} else {
			r.Config.ResponseDigestAlgorithm = types.StringNull()
		}
		r.Config.ResponseEncryptionKey = types.StringPointerValue(resp.Config.ResponseEncryptionKey)
		if resp.Config.ResponseSignatureAlgorithm != nil {
			r.Config.ResponseSignatureAlgorithm = types.StringValue(string(*resp.Config.ResponseSignatureAlgorithm))
		} else {
			r.Config.ResponseSignatureAlgorithm = types.StringNull()
		}
		if resp.Config.SessionAbsoluteTimeout != nil {
			r.Config.SessionAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionAbsoluteTimeout)))
		} else {
			r.Config.SessionAbsoluteTimeout = types.NumberNull()
		}
		r.Config.SessionAudience = types.StringPointerValue(resp.Config.SessionAudience)
		r.Config.SessionCookieDomain = types.StringPointerValue(resp.Config.SessionCookieDomain)
		r.Config.SessionCookieHTTPOnly = types.BoolPointerValue(resp.Config.SessionCookieHTTPOnly)
		r.Config.SessionCookieName = types.StringPointerValue(resp.Config.SessionCookieName)
		r.Config.SessionCookiePath = types.StringPointerValue(resp.Config.SessionCookiePath)
		if resp.Config.SessionCookieSameSite != nil {
			r.Config.SessionCookieSameSite = types.StringValue(string(*resp.Config.SessionCookieSameSite))
		} else {
			r.Config.SessionCookieSameSite = types.StringNull()
		}
		r.Config.SessionCookieSecure = types.BoolPointerValue(resp.Config.SessionCookieSecure)
		r.Config.SessionEnforceSameSubject = types.BoolPointerValue(resp.Config.SessionEnforceSameSubject)
		r.Config.SessionHashStorageKey = types.BoolPointerValue(resp.Config.SessionHashStorageKey)
		r.Config.SessionHashSubject = types.BoolPointerValue(resp.Config.SessionHashSubject)
		if resp.Config.SessionIdlingTimeout != nil {
			r.Config.SessionIdlingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionIdlingTimeout)))
		} else {
			r.Config.SessionIdlingTimeout = types.NumberNull()
		}
		r.Config.SessionMemcachedHost = types.StringPointerValue(resp.Config.SessionMemcachedHost)
		r.Config.SessionMemcachedPort = types.Int64PointerValue(resp.Config.SessionMemcachedPort)
		r.Config.SessionMemcachedPrefix = types.StringPointerValue(resp.Config.SessionMemcachedPrefix)
		r.Config.SessionMemcachedSocket = types.StringPointerValue(resp.Config.SessionMemcachedSocket)
		r.Config.SessionRemember = types.BoolPointerValue(resp.Config.SessionRemember)
		if resp.Config.SessionRememberAbsoluteTimeout != nil {
			r.Config.SessionRememberAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRememberAbsoluteTimeout)))
		} else {
			r.Config.SessionRememberAbsoluteTimeout = types.NumberNull()
		}
		r.Config.SessionRememberCookieName = types.StringPointerValue(resp.Config.SessionRememberCookieName)
		if resp.Config.SessionRememberRollingTimeout != nil {
			r.Config.SessionRememberRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRememberRollingTimeout)))
		} else {
			r.Config.SessionRememberRollingTimeout = types.NumberNull()
		}
		r.Config.SessionRequestHeaders = make([]types.String, 0, len(resp.Config.SessionRequestHeaders))
		for _, v := range resp.Config.SessionRequestHeaders {
			r.Config.SessionRequestHeaders = append(r.Config.SessionRequestHeaders, types.StringValue(string(v)))
		}
		r.Config.SessionResponseHeaders = make([]types.String, 0, len(resp.Config.SessionResponseHeaders))
		for _, v := range resp.Config.SessionResponseHeaders {
			r.Config.SessionResponseHeaders = append(r.Config.SessionResponseHeaders, types.StringValue(string(v)))
		}
		if resp.Config.SessionRollingTimeout != nil {
			r.Config.SessionRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.SessionRollingTimeout)))
		} else {
			r.Config.SessionRollingTimeout = types.NumberNull()
		}
		r.Config.SessionSecret = types.StringPointerValue(resp.Config.SessionSecret)
		if resp.Config.SessionStorage != nil {
			r.Config.SessionStorage = types.StringValue(string(*resp.Config.SessionStorage))
		} else {
			r.Config.SessionStorage = types.StringNull()
		}
		r.Config.SessionStoreMetadata = types.BoolPointerValue(resp.Config.SessionStoreMetadata)
		r.Config.ValidateAssertionSignature = types.BoolPointerValue(resp.Config.ValidateAssertionSignature)
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
