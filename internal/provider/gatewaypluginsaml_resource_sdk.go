// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginSamlResourceModel) ToSharedCreateSamlPlugin() *shared.CreateSamlPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateSamlPluginProtocols = []shared.CreateSamlPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateSamlPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateSamlPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateSamlPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateSamlPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateSamlPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateSamlPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateSamlPluginService{
			ID: id2,
		}
	}
	assertionConsumerPath := new(string)
	if !r.Config.AssertionConsumerPath.IsUnknown() && !r.Config.AssertionConsumerPath.IsNull() {
		*assertionConsumerPath = r.Config.AssertionConsumerPath.ValueString()
	} else {
		assertionConsumerPath = nil
	}
	idpSsoURL := new(string)
	if !r.Config.IdpSsoURL.IsUnknown() && !r.Config.IdpSsoURL.IsNull() {
		*idpSsoURL = r.Config.IdpSsoURL.ValueString()
	} else {
		idpSsoURL = nil
	}
	idpCertificate := new(string)
	if !r.Config.IdpCertificate.IsUnknown() && !r.Config.IdpCertificate.IsNull() {
		*idpCertificate = r.Config.IdpCertificate.ValueString()
	} else {
		idpCertificate = nil
	}
	responseEncryptionKey := new(string)
	if !r.Config.ResponseEncryptionKey.IsUnknown() && !r.Config.ResponseEncryptionKey.IsNull() {
		*responseEncryptionKey = r.Config.ResponseEncryptionKey.ValueString()
	} else {
		responseEncryptionKey = nil
	}
	requestSigningKey := new(string)
	if !r.Config.RequestSigningKey.IsUnknown() && !r.Config.RequestSigningKey.IsNull() {
		*requestSigningKey = r.Config.RequestSigningKey.ValueString()
	} else {
		requestSigningKey = nil
	}
	requestSigningCertificate := new(string)
	if !r.Config.RequestSigningCertificate.IsUnknown() && !r.Config.RequestSigningCertificate.IsNull() {
		*requestSigningCertificate = r.Config.RequestSigningCertificate.ValueString()
	} else {
		requestSigningCertificate = nil
	}
	requestSignatureAlgorithm := new(shared.RequestSignatureAlgorithm)
	if !r.Config.RequestSignatureAlgorithm.IsUnknown() && !r.Config.RequestSignatureAlgorithm.IsNull() {
		*requestSignatureAlgorithm = shared.RequestSignatureAlgorithm(r.Config.RequestSignatureAlgorithm.ValueString())
	} else {
		requestSignatureAlgorithm = nil
	}
	requestDigestAlgorithm := new(shared.RequestDigestAlgorithm)
	if !r.Config.RequestDigestAlgorithm.IsUnknown() && !r.Config.RequestDigestAlgorithm.IsNull() {
		*requestDigestAlgorithm = shared.RequestDigestAlgorithm(r.Config.RequestDigestAlgorithm.ValueString())
	} else {
		requestDigestAlgorithm = nil
	}
	responseSignatureAlgorithm := new(shared.ResponseSignatureAlgorithm)
	if !r.Config.ResponseSignatureAlgorithm.IsUnknown() && !r.Config.ResponseSignatureAlgorithm.IsNull() {
		*responseSignatureAlgorithm = shared.ResponseSignatureAlgorithm(r.Config.ResponseSignatureAlgorithm.ValueString())
	} else {
		responseSignatureAlgorithm = nil
	}
	responseDigestAlgorithm := new(shared.ResponseDigestAlgorithm)
	if !r.Config.ResponseDigestAlgorithm.IsUnknown() && !r.Config.ResponseDigestAlgorithm.IsNull() {
		*responseDigestAlgorithm = shared.ResponseDigestAlgorithm(r.Config.ResponseDigestAlgorithm.ValueString())
	} else {
		responseDigestAlgorithm = nil
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
	validateAssertionSignature := new(bool)
	if !r.Config.ValidateAssertionSignature.IsUnknown() && !r.Config.ValidateAssertionSignature.IsNull() {
		*validateAssertionSignature = r.Config.ValidateAssertionSignature.ValueBool()
	} else {
		validateAssertionSignature = nil
	}
	anonymous := new(string)
	if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
		*anonymous = r.Config.Anonymous.ValueString()
	} else {
		anonymous = nil
	}
	sessionSecret := new(string)
	if !r.Config.SessionSecret.IsUnknown() && !r.Config.SessionSecret.IsNull() {
		*sessionSecret = r.Config.SessionSecret.ValueString()
	} else {
		sessionSecret = nil
	}
	sessionAudience := new(string)
	if !r.Config.SessionAudience.IsUnknown() && !r.Config.SessionAudience.IsNull() {
		*sessionAudience = r.Config.SessionAudience.ValueString()
	} else {
		sessionAudience = nil
	}
	sessionCookieName := new(string)
	if !r.Config.SessionCookieName.IsUnknown() && !r.Config.SessionCookieName.IsNull() {
		*sessionCookieName = r.Config.SessionCookieName.ValueString()
	} else {
		sessionCookieName = nil
	}
	sessionRemember := new(bool)
	if !r.Config.SessionRemember.IsUnknown() && !r.Config.SessionRemember.IsNull() {
		*sessionRemember = r.Config.SessionRemember.ValueBool()
	} else {
		sessionRemember = nil
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
	sessionRememberAbsoluteTimeout := new(float64)
	if !r.Config.SessionRememberAbsoluteTimeout.IsUnknown() && !r.Config.SessionRememberAbsoluteTimeout.IsNull() {
		*sessionRememberAbsoluteTimeout, _ = r.Config.SessionRememberAbsoluteTimeout.ValueBigFloat().Float64()
	} else {
		sessionRememberAbsoluteTimeout = nil
	}
	sessionIdlingTimeout := new(float64)
	if !r.Config.SessionIdlingTimeout.IsUnknown() && !r.Config.SessionIdlingTimeout.IsNull() {
		*sessionIdlingTimeout, _ = r.Config.SessionIdlingTimeout.ValueBigFloat().Float64()
	} else {
		sessionIdlingTimeout = nil
	}
	sessionRollingTimeout := new(float64)
	if !r.Config.SessionRollingTimeout.IsUnknown() && !r.Config.SessionRollingTimeout.IsNull() {
		*sessionRollingTimeout, _ = r.Config.SessionRollingTimeout.ValueBigFloat().Float64()
	} else {
		sessionRollingTimeout = nil
	}
	sessionAbsoluteTimeout := new(float64)
	if !r.Config.SessionAbsoluteTimeout.IsUnknown() && !r.Config.SessionAbsoluteTimeout.IsNull() {
		*sessionAbsoluteTimeout, _ = r.Config.SessionAbsoluteTimeout.ValueBigFloat().Float64()
	} else {
		sessionAbsoluteTimeout = nil
	}
	sessionCookiePath := new(string)
	if !r.Config.SessionCookiePath.IsUnknown() && !r.Config.SessionCookiePath.IsNull() {
		*sessionCookiePath = r.Config.SessionCookiePath.ValueString()
	} else {
		sessionCookiePath = nil
	}
	sessionCookieDomain := new(string)
	if !r.Config.SessionCookieDomain.IsUnknown() && !r.Config.SessionCookieDomain.IsNull() {
		*sessionCookieDomain = r.Config.SessionCookieDomain.ValueString()
	} else {
		sessionCookieDomain = nil
	}
	sessionCookieSameSite := new(shared.CreateSamlPluginSessionCookieSameSite)
	if !r.Config.SessionCookieSameSite.IsUnknown() && !r.Config.SessionCookieSameSite.IsNull() {
		*sessionCookieSameSite = shared.CreateSamlPluginSessionCookieSameSite(r.Config.SessionCookieSameSite.ValueString())
	} else {
		sessionCookieSameSite = nil
	}
	sessionCookieHTTPOnly := new(bool)
	if !r.Config.SessionCookieHTTPOnly.IsUnknown() && !r.Config.SessionCookieHTTPOnly.IsNull() {
		*sessionCookieHTTPOnly = r.Config.SessionCookieHTTPOnly.ValueBool()
	} else {
		sessionCookieHTTPOnly = nil
	}
	sessionCookieSecure := new(bool)
	if !r.Config.SessionCookieSecure.IsUnknown() && !r.Config.SessionCookieSecure.IsNull() {
		*sessionCookieSecure = r.Config.SessionCookieSecure.ValueBool()
	} else {
		sessionCookieSecure = nil
	}
	var sessionRequestHeaders []shared.CreateSamlPluginSessionRequestHeaders = []shared.CreateSamlPluginSessionRequestHeaders{}
	for _, sessionRequestHeadersItem := range r.Config.SessionRequestHeaders {
		sessionRequestHeaders = append(sessionRequestHeaders, shared.CreateSamlPluginSessionRequestHeaders(sessionRequestHeadersItem.ValueString()))
	}
	var sessionResponseHeaders []shared.CreateSamlPluginSessionResponseHeaders = []shared.CreateSamlPluginSessionResponseHeaders{}
	for _, sessionResponseHeadersItem := range r.Config.SessionResponseHeaders {
		sessionResponseHeaders = append(sessionResponseHeaders, shared.CreateSamlPluginSessionResponseHeaders(sessionResponseHeadersItem.ValueString()))
	}
	sessionStorage := new(shared.CreateSamlPluginSessionStorage)
	if !r.Config.SessionStorage.IsUnknown() && !r.Config.SessionStorage.IsNull() {
		*sessionStorage = shared.CreateSamlPluginSessionStorage(r.Config.SessionStorage.ValueString())
	} else {
		sessionStorage = nil
	}
	sessionStoreMetadata := new(bool)
	if !r.Config.SessionStoreMetadata.IsUnknown() && !r.Config.SessionStoreMetadata.IsNull() {
		*sessionStoreMetadata = r.Config.SessionStoreMetadata.ValueBool()
	} else {
		sessionStoreMetadata = nil
	}
	sessionEnforceSameSubject := new(bool)
	if !r.Config.SessionEnforceSameSubject.IsUnknown() && !r.Config.SessionEnforceSameSubject.IsNull() {
		*sessionEnforceSameSubject = r.Config.SessionEnforceSameSubject.ValueBool()
	} else {
		sessionEnforceSameSubject = nil
	}
	sessionHashSubject := new(bool)
	if !r.Config.SessionHashSubject.IsUnknown() && !r.Config.SessionHashSubject.IsNull() {
		*sessionHashSubject = r.Config.SessionHashSubject.ValueBool()
	} else {
		sessionHashSubject = nil
	}
	sessionHashStorageKey := new(bool)
	if !r.Config.SessionHashStorageKey.IsUnknown() && !r.Config.SessionHashStorageKey.IsNull() {
		*sessionHashStorageKey = r.Config.SessionHashStorageKey.ValueBool()
	} else {
		sessionHashStorageKey = nil
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
	sessionRedisPrefix := new(string)
	if !r.Config.SessionRedisPrefix.IsUnknown() && !r.Config.SessionRedisPrefix.IsNull() {
		*sessionRedisPrefix = r.Config.SessionRedisPrefix.ValueString()
	} else {
		sessionRedisPrefix = nil
	}
	sessionRedisSocket := new(string)
	if !r.Config.SessionRedisSocket.IsUnknown() && !r.Config.SessionRedisSocket.IsNull() {
		*sessionRedisSocket = r.Config.SessionRedisSocket.ValueString()
	} else {
		sessionRedisSocket = nil
	}
	sessionRedisHost := new(string)
	if !r.Config.SessionRedisHost.IsUnknown() && !r.Config.SessionRedisHost.IsNull() {
		*sessionRedisHost = r.Config.SessionRedisHost.ValueString()
	} else {
		sessionRedisHost = nil
	}
	sessionRedisPort := new(int64)
	if !r.Config.SessionRedisPort.IsUnknown() && !r.Config.SessionRedisPort.IsNull() {
		*sessionRedisPort = r.Config.SessionRedisPort.ValueInt64()
	} else {
		sessionRedisPort = nil
	}
	sessionRedisUsername := new(string)
	if !r.Config.SessionRedisUsername.IsUnknown() && !r.Config.SessionRedisUsername.IsNull() {
		*sessionRedisUsername = r.Config.SessionRedisUsername.ValueString()
	} else {
		sessionRedisUsername = nil
	}
	sessionRedisPassword := new(string)
	if !r.Config.SessionRedisPassword.IsUnknown() && !r.Config.SessionRedisPassword.IsNull() {
		*sessionRedisPassword = r.Config.SessionRedisPassword.ValueString()
	} else {
		sessionRedisPassword = nil
	}
	sessionRedisConnectTimeout := new(int64)
	if !r.Config.SessionRedisConnectTimeout.IsUnknown() && !r.Config.SessionRedisConnectTimeout.IsNull() {
		*sessionRedisConnectTimeout = r.Config.SessionRedisConnectTimeout.ValueInt64()
	} else {
		sessionRedisConnectTimeout = nil
	}
	sessionRedisReadTimeout := new(int64)
	if !r.Config.SessionRedisReadTimeout.IsUnknown() && !r.Config.SessionRedisReadTimeout.IsNull() {
		*sessionRedisReadTimeout = r.Config.SessionRedisReadTimeout.ValueInt64()
	} else {
		sessionRedisReadTimeout = nil
	}
	sessionRedisSendTimeout := new(int64)
	if !r.Config.SessionRedisSendTimeout.IsUnknown() && !r.Config.SessionRedisSendTimeout.IsNull() {
		*sessionRedisSendTimeout = r.Config.SessionRedisSendTimeout.ValueInt64()
	} else {
		sessionRedisSendTimeout = nil
	}
	sessionRedisSsl := new(bool)
	if !r.Config.SessionRedisSsl.IsUnknown() && !r.Config.SessionRedisSsl.IsNull() {
		*sessionRedisSsl = r.Config.SessionRedisSsl.ValueBool()
	} else {
		sessionRedisSsl = nil
	}
	sessionRedisSslVerify := new(bool)
	if !r.Config.SessionRedisSslVerify.IsUnknown() && !r.Config.SessionRedisSslVerify.IsNull() {
		*sessionRedisSslVerify = r.Config.SessionRedisSslVerify.ValueBool()
	} else {
		sessionRedisSslVerify = nil
	}
	sessionRedisServerName := new(string)
	if !r.Config.SessionRedisServerName.IsUnknown() && !r.Config.SessionRedisServerName.IsNull() {
		*sessionRedisServerName = r.Config.SessionRedisServerName.ValueString()
	} else {
		sessionRedisServerName = nil
	}
	var sessionRedisClusterNodes []shared.CreateSamlPluginSessionRedisClusterNodes = []shared.CreateSamlPluginSessionRedisClusterNodes{}
	for _, sessionRedisClusterNodesItem := range r.Config.SessionRedisClusterNodes {
		ip := new(string)
		if !sessionRedisClusterNodesItem.IP.IsUnknown() && !sessionRedisClusterNodesItem.IP.IsNull() {
			*ip = sessionRedisClusterNodesItem.IP.ValueString()
		} else {
			ip = nil
		}
		port := new(int64)
		if !sessionRedisClusterNodesItem.Port.IsUnknown() && !sessionRedisClusterNodesItem.Port.IsNull() {
			*port = sessionRedisClusterNodesItem.Port.ValueInt64()
		} else {
			port = nil
		}
		sessionRedisClusterNodes = append(sessionRedisClusterNodes, shared.CreateSamlPluginSessionRedisClusterNodes{
			IP:   ip,
			Port: port,
		})
	}
	sessionRedisClusterMaxRedirections := new(int64)
	if !r.Config.SessionRedisClusterMaxRedirections.IsUnknown() && !r.Config.SessionRedisClusterMaxRedirections.IsNull() {
		*sessionRedisClusterMaxRedirections = r.Config.SessionRedisClusterMaxRedirections.ValueInt64()
	} else {
		sessionRedisClusterMaxRedirections = nil
	}
	config := shared.CreateSamlPluginConfig{
		AssertionConsumerPath:              assertionConsumerPath,
		IdpSsoURL:                          idpSsoURL,
		IdpCertificate:                     idpCertificate,
		ResponseEncryptionKey:              responseEncryptionKey,
		RequestSigningKey:                  requestSigningKey,
		RequestSigningCertificate:          requestSigningCertificate,
		RequestSignatureAlgorithm:          requestSignatureAlgorithm,
		RequestDigestAlgorithm:             requestDigestAlgorithm,
		ResponseSignatureAlgorithm:         responseSignatureAlgorithm,
		ResponseDigestAlgorithm:            responseDigestAlgorithm,
		Issuer:                             issuer,
		NameidFormat:                       nameidFormat,
		ValidateAssertionSignature:         validateAssertionSignature,
		Anonymous:                          anonymous,
		SessionSecret:                      sessionSecret,
		SessionAudience:                    sessionAudience,
		SessionCookieName:                  sessionCookieName,
		SessionRemember:                    sessionRemember,
		SessionRememberCookieName:          sessionRememberCookieName,
		SessionRememberRollingTimeout:      sessionRememberRollingTimeout,
		SessionRememberAbsoluteTimeout:     sessionRememberAbsoluteTimeout,
		SessionIdlingTimeout:               sessionIdlingTimeout,
		SessionRollingTimeout:              sessionRollingTimeout,
		SessionAbsoluteTimeout:             sessionAbsoluteTimeout,
		SessionCookiePath:                  sessionCookiePath,
		SessionCookieDomain:                sessionCookieDomain,
		SessionCookieSameSite:              sessionCookieSameSite,
		SessionCookieHTTPOnly:              sessionCookieHTTPOnly,
		SessionCookieSecure:                sessionCookieSecure,
		SessionRequestHeaders:              sessionRequestHeaders,
		SessionResponseHeaders:             sessionResponseHeaders,
		SessionStorage:                     sessionStorage,
		SessionStoreMetadata:               sessionStoreMetadata,
		SessionEnforceSameSubject:          sessionEnforceSameSubject,
		SessionHashSubject:                 sessionHashSubject,
		SessionHashStorageKey:              sessionHashStorageKey,
		SessionMemcachedPrefix:             sessionMemcachedPrefix,
		SessionMemcachedSocket:             sessionMemcachedSocket,
		SessionMemcachedHost:               sessionMemcachedHost,
		SessionMemcachedPort:               sessionMemcachedPort,
		SessionRedisPrefix:                 sessionRedisPrefix,
		SessionRedisSocket:                 sessionRedisSocket,
		SessionRedisHost:                   sessionRedisHost,
		SessionRedisPort:                   sessionRedisPort,
		SessionRedisUsername:               sessionRedisUsername,
		SessionRedisPassword:               sessionRedisPassword,
		SessionRedisConnectTimeout:         sessionRedisConnectTimeout,
		SessionRedisReadTimeout:            sessionRedisReadTimeout,
		SessionRedisSendTimeout:            sessionRedisSendTimeout,
		SessionRedisSsl:                    sessionRedisSsl,
		SessionRedisSslVerify:              sessionRedisSslVerify,
		SessionRedisServerName:             sessionRedisServerName,
		SessionRedisClusterNodes:           sessionRedisClusterNodes,
		SessionRedisClusterMaxRedirections: sessionRedisClusterMaxRedirections,
	}
	out := shared.CreateSamlPlugin{
		Enabled:   enabled,
		Protocols: protocols,
		Tags:      tags,
		Consumer:  consumer,
		Route:     route,
		Service:   service,
		Config:    config,
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
		r.Config.SessionRedisClusterMaxRedirections = types.Int64PointerValue(resp.Config.SessionRedisClusterMaxRedirections)
		if len(r.Config.SessionRedisClusterNodes) > len(resp.Config.SessionRedisClusterNodes) {
			r.Config.SessionRedisClusterNodes = r.Config.SessionRedisClusterNodes[:len(resp.Config.SessionRedisClusterNodes)]
		}
		for sessionRedisClusterNodesCount, sessionRedisClusterNodesItem := range resp.Config.SessionRedisClusterNodes {
			var sessionRedisClusterNodes1 tfTypes.SessionRedisClusterNodes
			sessionRedisClusterNodes1.IP = types.StringPointerValue(sessionRedisClusterNodesItem.IP)
			sessionRedisClusterNodes1.Port = types.Int64PointerValue(sessionRedisClusterNodesItem.Port)
			if sessionRedisClusterNodesCount+1 > len(r.Config.SessionRedisClusterNodes) {
				r.Config.SessionRedisClusterNodes = append(r.Config.SessionRedisClusterNodes, sessionRedisClusterNodes1)
			} else {
				r.Config.SessionRedisClusterNodes[sessionRedisClusterNodesCount].IP = sessionRedisClusterNodes1.IP
				r.Config.SessionRedisClusterNodes[sessionRedisClusterNodesCount].Port = sessionRedisClusterNodes1.Port
			}
		}
		r.Config.SessionRedisConnectTimeout = types.Int64PointerValue(resp.Config.SessionRedisConnectTimeout)
		r.Config.SessionRedisHost = types.StringPointerValue(resp.Config.SessionRedisHost)
		r.Config.SessionRedisPassword = types.StringPointerValue(resp.Config.SessionRedisPassword)
		r.Config.SessionRedisPort = types.Int64PointerValue(resp.Config.SessionRedisPort)
		r.Config.SessionRedisPrefix = types.StringPointerValue(resp.Config.SessionRedisPrefix)
		r.Config.SessionRedisReadTimeout = types.Int64PointerValue(resp.Config.SessionRedisReadTimeout)
		r.Config.SessionRedisSendTimeout = types.Int64PointerValue(resp.Config.SessionRedisSendTimeout)
		r.Config.SessionRedisServerName = types.StringPointerValue(resp.Config.SessionRedisServerName)
		r.Config.SessionRedisSocket = types.StringPointerValue(resp.Config.SessionRedisSocket)
		r.Config.SessionRedisSsl = types.BoolPointerValue(resp.Config.SessionRedisSsl)
		r.Config.SessionRedisSslVerify = types.BoolPointerValue(resp.Config.SessionRedisSslVerify)
		r.Config.SessionRedisUsername = types.StringPointerValue(resp.Config.SessionRedisUsername)
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
		r.Config.SessionRequestHeaders = []types.String{}
		for _, v := range resp.Config.SessionRequestHeaders {
			r.Config.SessionRequestHeaders = append(r.Config.SessionRequestHeaders, types.StringValue(string(v)))
		}
		r.Config.SessionResponseHeaders = []types.String{}
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
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
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
	}
}
