// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginJWTSignerResourceModel) ToSharedCreateJWTSignerPlugin() *shared.CreateJWTSignerPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateJWTSignerPluginProtocols = []shared.CreateJWTSignerPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateJWTSignerPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateJWTSignerPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateJWTSignerPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateJWTSignerPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateJWTSignerPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateJWTSignerPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateJWTSignerPluginService{
			ID: id2,
		}
	}
	realm := new(string)
	if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
		*realm = r.Config.Realm.ValueString()
	} else {
		realm = nil
	}
	enableHsSignatures := new(bool)
	if !r.Config.EnableHsSignatures.IsUnknown() && !r.Config.EnableHsSignatures.IsNull() {
		*enableHsSignatures = r.Config.EnableHsSignatures.ValueBool()
	} else {
		enableHsSignatures = nil
	}
	enableInstrumentation := new(bool)
	if !r.Config.EnableInstrumentation.IsUnknown() && !r.Config.EnableInstrumentation.IsNull() {
		*enableInstrumentation = r.Config.EnableInstrumentation.ValueBool()
	} else {
		enableInstrumentation = nil
	}
	accessTokenIssuer := new(string)
	if !r.Config.AccessTokenIssuer.IsUnknown() && !r.Config.AccessTokenIssuer.IsNull() {
		*accessTokenIssuer = r.Config.AccessTokenIssuer.ValueString()
	} else {
		accessTokenIssuer = nil
	}
	accessTokenKeyset := new(string)
	if !r.Config.AccessTokenKeyset.IsUnknown() && !r.Config.AccessTokenKeyset.IsNull() {
		*accessTokenKeyset = r.Config.AccessTokenKeyset.ValueString()
	} else {
		accessTokenKeyset = nil
	}
	accessTokenJwksURI := new(string)
	if !r.Config.AccessTokenJwksURI.IsUnknown() && !r.Config.AccessTokenJwksURI.IsNull() {
		*accessTokenJwksURI = r.Config.AccessTokenJwksURI.ValueString()
	} else {
		accessTokenJwksURI = nil
	}
	accessTokenRequestHeader := new(string)
	if !r.Config.AccessTokenRequestHeader.IsUnknown() && !r.Config.AccessTokenRequestHeader.IsNull() {
		*accessTokenRequestHeader = r.Config.AccessTokenRequestHeader.ValueString()
	} else {
		accessTokenRequestHeader = nil
	}
	accessTokenLeeway := new(float64)
	if !r.Config.AccessTokenLeeway.IsUnknown() && !r.Config.AccessTokenLeeway.IsNull() {
		*accessTokenLeeway, _ = r.Config.AccessTokenLeeway.ValueBigFloat().Float64()
	} else {
		accessTokenLeeway = nil
	}
	var accessTokenScopesRequired []string = []string{}
	for _, accessTokenScopesRequiredItem := range r.Config.AccessTokenScopesRequired {
		accessTokenScopesRequired = append(accessTokenScopesRequired, accessTokenScopesRequiredItem.ValueString())
	}
	var accessTokenScopesClaim []string = []string{}
	for _, accessTokenScopesClaimItem := range r.Config.AccessTokenScopesClaim {
		accessTokenScopesClaim = append(accessTokenScopesClaim, accessTokenScopesClaimItem.ValueString())
	}
	var accessTokenConsumerClaim []string = []string{}
	for _, accessTokenConsumerClaimItem := range r.Config.AccessTokenConsumerClaim {
		accessTokenConsumerClaim = append(accessTokenConsumerClaim, accessTokenConsumerClaimItem.ValueString())
	}
	var accessTokenConsumerBy []shared.AccessTokenConsumerBy = []shared.AccessTokenConsumerBy{}
	for _, accessTokenConsumerByItem := range r.Config.AccessTokenConsumerBy {
		accessTokenConsumerBy = append(accessTokenConsumerBy, shared.AccessTokenConsumerBy(accessTokenConsumerByItem.ValueString()))
	}
	accessTokenUpstreamHeader := new(string)
	if !r.Config.AccessTokenUpstreamHeader.IsUnknown() && !r.Config.AccessTokenUpstreamHeader.IsNull() {
		*accessTokenUpstreamHeader = r.Config.AccessTokenUpstreamHeader.ValueString()
	} else {
		accessTokenUpstreamHeader = nil
	}
	accessTokenUpstreamLeeway := new(float64)
	if !r.Config.AccessTokenUpstreamLeeway.IsUnknown() && !r.Config.AccessTokenUpstreamLeeway.IsNull() {
		*accessTokenUpstreamLeeway, _ = r.Config.AccessTokenUpstreamLeeway.ValueBigFloat().Float64()
	} else {
		accessTokenUpstreamLeeway = nil
	}
	accessTokenIntrospectionEndpoint := new(string)
	if !r.Config.AccessTokenIntrospectionEndpoint.IsUnknown() && !r.Config.AccessTokenIntrospectionEndpoint.IsNull() {
		*accessTokenIntrospectionEndpoint = r.Config.AccessTokenIntrospectionEndpoint.ValueString()
	} else {
		accessTokenIntrospectionEndpoint = nil
	}
	accessTokenIntrospectionAuthorization := new(string)
	if !r.Config.AccessTokenIntrospectionAuthorization.IsUnknown() && !r.Config.AccessTokenIntrospectionAuthorization.IsNull() {
		*accessTokenIntrospectionAuthorization = r.Config.AccessTokenIntrospectionAuthorization.ValueString()
	} else {
		accessTokenIntrospectionAuthorization = nil
	}
	accessTokenIntrospectionBodyArgs := new(string)
	if !r.Config.AccessTokenIntrospectionBodyArgs.IsUnknown() && !r.Config.AccessTokenIntrospectionBodyArgs.IsNull() {
		*accessTokenIntrospectionBodyArgs = r.Config.AccessTokenIntrospectionBodyArgs.ValueString()
	} else {
		accessTokenIntrospectionBodyArgs = nil
	}
	accessTokenIntrospectionHint := new(string)
	if !r.Config.AccessTokenIntrospectionHint.IsUnknown() && !r.Config.AccessTokenIntrospectionHint.IsNull() {
		*accessTokenIntrospectionHint = r.Config.AccessTokenIntrospectionHint.ValueString()
	} else {
		accessTokenIntrospectionHint = nil
	}
	var accessTokenIntrospectionJwtClaim []string = []string{}
	for _, accessTokenIntrospectionJwtClaimItem := range r.Config.AccessTokenIntrospectionJwtClaim {
		accessTokenIntrospectionJwtClaim = append(accessTokenIntrospectionJwtClaim, accessTokenIntrospectionJwtClaimItem.ValueString())
	}
	var accessTokenIntrospectionScopesRequired []string = []string{}
	for _, accessTokenIntrospectionScopesRequiredItem := range r.Config.AccessTokenIntrospectionScopesRequired {
		accessTokenIntrospectionScopesRequired = append(accessTokenIntrospectionScopesRequired, accessTokenIntrospectionScopesRequiredItem.ValueString())
	}
	var accessTokenIntrospectionScopesClaim []string = []string{}
	for _, accessTokenIntrospectionScopesClaimItem := range r.Config.AccessTokenIntrospectionScopesClaim {
		accessTokenIntrospectionScopesClaim = append(accessTokenIntrospectionScopesClaim, accessTokenIntrospectionScopesClaimItem.ValueString())
	}
	var accessTokenIntrospectionConsumerClaim []string = []string{}
	for _, accessTokenIntrospectionConsumerClaimItem := range r.Config.AccessTokenIntrospectionConsumerClaim {
		accessTokenIntrospectionConsumerClaim = append(accessTokenIntrospectionConsumerClaim, accessTokenIntrospectionConsumerClaimItem.ValueString())
	}
	var accessTokenIntrospectionConsumerBy []shared.AccessTokenIntrospectionConsumerBy = []shared.AccessTokenIntrospectionConsumerBy{}
	for _, accessTokenIntrospectionConsumerByItem := range r.Config.AccessTokenIntrospectionConsumerBy {
		accessTokenIntrospectionConsumerBy = append(accessTokenIntrospectionConsumerBy, shared.AccessTokenIntrospectionConsumerBy(accessTokenIntrospectionConsumerByItem.ValueString()))
	}
	accessTokenIntrospectionLeeway := new(float64)
	if !r.Config.AccessTokenIntrospectionLeeway.IsUnknown() && !r.Config.AccessTokenIntrospectionLeeway.IsNull() {
		*accessTokenIntrospectionLeeway, _ = r.Config.AccessTokenIntrospectionLeeway.ValueBigFloat().Float64()
	} else {
		accessTokenIntrospectionLeeway = nil
	}
	accessTokenIntrospectionTimeout := new(float64)
	if !r.Config.AccessTokenIntrospectionTimeout.IsUnknown() && !r.Config.AccessTokenIntrospectionTimeout.IsNull() {
		*accessTokenIntrospectionTimeout, _ = r.Config.AccessTokenIntrospectionTimeout.ValueBigFloat().Float64()
	} else {
		accessTokenIntrospectionTimeout = nil
	}
	accessTokenSigningAlgorithm := new(shared.AccessTokenSigningAlgorithm)
	if !r.Config.AccessTokenSigningAlgorithm.IsUnknown() && !r.Config.AccessTokenSigningAlgorithm.IsNull() {
		*accessTokenSigningAlgorithm = shared.AccessTokenSigningAlgorithm(r.Config.AccessTokenSigningAlgorithm.ValueString())
	} else {
		accessTokenSigningAlgorithm = nil
	}
	accessTokenOptional := new(bool)
	if !r.Config.AccessTokenOptional.IsUnknown() && !r.Config.AccessTokenOptional.IsNull() {
		*accessTokenOptional = r.Config.AccessTokenOptional.ValueBool()
	} else {
		accessTokenOptional = nil
	}
	verifyAccessTokenSignature := new(bool)
	if !r.Config.VerifyAccessTokenSignature.IsUnknown() && !r.Config.VerifyAccessTokenSignature.IsNull() {
		*verifyAccessTokenSignature = r.Config.VerifyAccessTokenSignature.ValueBool()
	} else {
		verifyAccessTokenSignature = nil
	}
	verifyAccessTokenExpiry := new(bool)
	if !r.Config.VerifyAccessTokenExpiry.IsUnknown() && !r.Config.VerifyAccessTokenExpiry.IsNull() {
		*verifyAccessTokenExpiry = r.Config.VerifyAccessTokenExpiry.ValueBool()
	} else {
		verifyAccessTokenExpiry = nil
	}
	verifyAccessTokenScopes := new(bool)
	if !r.Config.VerifyAccessTokenScopes.IsUnknown() && !r.Config.VerifyAccessTokenScopes.IsNull() {
		*verifyAccessTokenScopes = r.Config.VerifyAccessTokenScopes.ValueBool()
	} else {
		verifyAccessTokenScopes = nil
	}
	verifyAccessTokenIntrospectionExpiry := new(bool)
	if !r.Config.VerifyAccessTokenIntrospectionExpiry.IsUnknown() && !r.Config.VerifyAccessTokenIntrospectionExpiry.IsNull() {
		*verifyAccessTokenIntrospectionExpiry = r.Config.VerifyAccessTokenIntrospectionExpiry.ValueBool()
	} else {
		verifyAccessTokenIntrospectionExpiry = nil
	}
	verifyAccessTokenIntrospectionScopes := new(bool)
	if !r.Config.VerifyAccessTokenIntrospectionScopes.IsUnknown() && !r.Config.VerifyAccessTokenIntrospectionScopes.IsNull() {
		*verifyAccessTokenIntrospectionScopes = r.Config.VerifyAccessTokenIntrospectionScopes.ValueBool()
	} else {
		verifyAccessTokenIntrospectionScopes = nil
	}
	cacheAccessTokenIntrospection := new(bool)
	if !r.Config.CacheAccessTokenIntrospection.IsUnknown() && !r.Config.CacheAccessTokenIntrospection.IsNull() {
		*cacheAccessTokenIntrospection = r.Config.CacheAccessTokenIntrospection.ValueBool()
	} else {
		cacheAccessTokenIntrospection = nil
	}
	trustAccessTokenIntrospection := new(bool)
	if !r.Config.TrustAccessTokenIntrospection.IsUnknown() && !r.Config.TrustAccessTokenIntrospection.IsNull() {
		*trustAccessTokenIntrospection = r.Config.TrustAccessTokenIntrospection.ValueBool()
	} else {
		trustAccessTokenIntrospection = nil
	}
	enableAccessTokenIntrospection := new(bool)
	if !r.Config.EnableAccessTokenIntrospection.IsUnknown() && !r.Config.EnableAccessTokenIntrospection.IsNull() {
		*enableAccessTokenIntrospection = r.Config.EnableAccessTokenIntrospection.ValueBool()
	} else {
		enableAccessTokenIntrospection = nil
	}
	channelTokenIssuer := new(string)
	if !r.Config.ChannelTokenIssuer.IsUnknown() && !r.Config.ChannelTokenIssuer.IsNull() {
		*channelTokenIssuer = r.Config.ChannelTokenIssuer.ValueString()
	} else {
		channelTokenIssuer = nil
	}
	channelTokenKeyset := new(string)
	if !r.Config.ChannelTokenKeyset.IsUnknown() && !r.Config.ChannelTokenKeyset.IsNull() {
		*channelTokenKeyset = r.Config.ChannelTokenKeyset.ValueString()
	} else {
		channelTokenKeyset = nil
	}
	channelTokenJwksURI := new(string)
	if !r.Config.ChannelTokenJwksURI.IsUnknown() && !r.Config.ChannelTokenJwksURI.IsNull() {
		*channelTokenJwksURI = r.Config.ChannelTokenJwksURI.ValueString()
	} else {
		channelTokenJwksURI = nil
	}
	channelTokenRequestHeader := new(string)
	if !r.Config.ChannelTokenRequestHeader.IsUnknown() && !r.Config.ChannelTokenRequestHeader.IsNull() {
		*channelTokenRequestHeader = r.Config.ChannelTokenRequestHeader.ValueString()
	} else {
		channelTokenRequestHeader = nil
	}
	channelTokenLeeway := new(float64)
	if !r.Config.ChannelTokenLeeway.IsUnknown() && !r.Config.ChannelTokenLeeway.IsNull() {
		*channelTokenLeeway, _ = r.Config.ChannelTokenLeeway.ValueBigFloat().Float64()
	} else {
		channelTokenLeeway = nil
	}
	var channelTokenScopesRequired []string = []string{}
	for _, channelTokenScopesRequiredItem := range r.Config.ChannelTokenScopesRequired {
		channelTokenScopesRequired = append(channelTokenScopesRequired, channelTokenScopesRequiredItem.ValueString())
	}
	var channelTokenScopesClaim []string = []string{}
	for _, channelTokenScopesClaimItem := range r.Config.ChannelTokenScopesClaim {
		channelTokenScopesClaim = append(channelTokenScopesClaim, channelTokenScopesClaimItem.ValueString())
	}
	var channelTokenConsumerClaim []string = []string{}
	for _, channelTokenConsumerClaimItem := range r.Config.ChannelTokenConsumerClaim {
		channelTokenConsumerClaim = append(channelTokenConsumerClaim, channelTokenConsumerClaimItem.ValueString())
	}
	var channelTokenConsumerBy []shared.ChannelTokenConsumerBy = []shared.ChannelTokenConsumerBy{}
	for _, channelTokenConsumerByItem := range r.Config.ChannelTokenConsumerBy {
		channelTokenConsumerBy = append(channelTokenConsumerBy, shared.ChannelTokenConsumerBy(channelTokenConsumerByItem.ValueString()))
	}
	channelTokenUpstreamHeader := new(string)
	if !r.Config.ChannelTokenUpstreamHeader.IsUnknown() && !r.Config.ChannelTokenUpstreamHeader.IsNull() {
		*channelTokenUpstreamHeader = r.Config.ChannelTokenUpstreamHeader.ValueString()
	} else {
		channelTokenUpstreamHeader = nil
	}
	channelTokenUpstreamLeeway := new(float64)
	if !r.Config.ChannelTokenUpstreamLeeway.IsUnknown() && !r.Config.ChannelTokenUpstreamLeeway.IsNull() {
		*channelTokenUpstreamLeeway, _ = r.Config.ChannelTokenUpstreamLeeway.ValueBigFloat().Float64()
	} else {
		channelTokenUpstreamLeeway = nil
	}
	channelTokenIntrospectionEndpoint := new(string)
	if !r.Config.ChannelTokenIntrospectionEndpoint.IsUnknown() && !r.Config.ChannelTokenIntrospectionEndpoint.IsNull() {
		*channelTokenIntrospectionEndpoint = r.Config.ChannelTokenIntrospectionEndpoint.ValueString()
	} else {
		channelTokenIntrospectionEndpoint = nil
	}
	channelTokenIntrospectionAuthorization := new(string)
	if !r.Config.ChannelTokenIntrospectionAuthorization.IsUnknown() && !r.Config.ChannelTokenIntrospectionAuthorization.IsNull() {
		*channelTokenIntrospectionAuthorization = r.Config.ChannelTokenIntrospectionAuthorization.ValueString()
	} else {
		channelTokenIntrospectionAuthorization = nil
	}
	channelTokenIntrospectionBodyArgs := new(string)
	if !r.Config.ChannelTokenIntrospectionBodyArgs.IsUnknown() && !r.Config.ChannelTokenIntrospectionBodyArgs.IsNull() {
		*channelTokenIntrospectionBodyArgs = r.Config.ChannelTokenIntrospectionBodyArgs.ValueString()
	} else {
		channelTokenIntrospectionBodyArgs = nil
	}
	channelTokenIntrospectionHint := new(string)
	if !r.Config.ChannelTokenIntrospectionHint.IsUnknown() && !r.Config.ChannelTokenIntrospectionHint.IsNull() {
		*channelTokenIntrospectionHint = r.Config.ChannelTokenIntrospectionHint.ValueString()
	} else {
		channelTokenIntrospectionHint = nil
	}
	var channelTokenIntrospectionJwtClaim []string = []string{}
	for _, channelTokenIntrospectionJwtClaimItem := range r.Config.ChannelTokenIntrospectionJwtClaim {
		channelTokenIntrospectionJwtClaim = append(channelTokenIntrospectionJwtClaim, channelTokenIntrospectionJwtClaimItem.ValueString())
	}
	var channelTokenIntrospectionScopesRequired []string = []string{}
	for _, channelTokenIntrospectionScopesRequiredItem := range r.Config.ChannelTokenIntrospectionScopesRequired {
		channelTokenIntrospectionScopesRequired = append(channelTokenIntrospectionScopesRequired, channelTokenIntrospectionScopesRequiredItem.ValueString())
	}
	var channelTokenIntrospectionScopesClaim []string = []string{}
	for _, channelTokenIntrospectionScopesClaimItem := range r.Config.ChannelTokenIntrospectionScopesClaim {
		channelTokenIntrospectionScopesClaim = append(channelTokenIntrospectionScopesClaim, channelTokenIntrospectionScopesClaimItem.ValueString())
	}
	var channelTokenIntrospectionConsumerClaim []string = []string{}
	for _, channelTokenIntrospectionConsumerClaimItem := range r.Config.ChannelTokenIntrospectionConsumerClaim {
		channelTokenIntrospectionConsumerClaim = append(channelTokenIntrospectionConsumerClaim, channelTokenIntrospectionConsumerClaimItem.ValueString())
	}
	var channelTokenIntrospectionConsumerBy []shared.ChannelTokenIntrospectionConsumerBy = []shared.ChannelTokenIntrospectionConsumerBy{}
	for _, channelTokenIntrospectionConsumerByItem := range r.Config.ChannelTokenIntrospectionConsumerBy {
		channelTokenIntrospectionConsumerBy = append(channelTokenIntrospectionConsumerBy, shared.ChannelTokenIntrospectionConsumerBy(channelTokenIntrospectionConsumerByItem.ValueString()))
	}
	channelTokenIntrospectionLeeway := new(float64)
	if !r.Config.ChannelTokenIntrospectionLeeway.IsUnknown() && !r.Config.ChannelTokenIntrospectionLeeway.IsNull() {
		*channelTokenIntrospectionLeeway, _ = r.Config.ChannelTokenIntrospectionLeeway.ValueBigFloat().Float64()
	} else {
		channelTokenIntrospectionLeeway = nil
	}
	channelTokenIntrospectionTimeout := new(float64)
	if !r.Config.ChannelTokenIntrospectionTimeout.IsUnknown() && !r.Config.ChannelTokenIntrospectionTimeout.IsNull() {
		*channelTokenIntrospectionTimeout, _ = r.Config.ChannelTokenIntrospectionTimeout.ValueBigFloat().Float64()
	} else {
		channelTokenIntrospectionTimeout = nil
	}
	channelTokenSigningAlgorithm := new(shared.ChannelTokenSigningAlgorithm)
	if !r.Config.ChannelTokenSigningAlgorithm.IsUnknown() && !r.Config.ChannelTokenSigningAlgorithm.IsNull() {
		*channelTokenSigningAlgorithm = shared.ChannelTokenSigningAlgorithm(r.Config.ChannelTokenSigningAlgorithm.ValueString())
	} else {
		channelTokenSigningAlgorithm = nil
	}
	channelTokenOptional := new(bool)
	if !r.Config.ChannelTokenOptional.IsUnknown() && !r.Config.ChannelTokenOptional.IsNull() {
		*channelTokenOptional = r.Config.ChannelTokenOptional.ValueBool()
	} else {
		channelTokenOptional = nil
	}
	verifyChannelTokenSignature := new(bool)
	if !r.Config.VerifyChannelTokenSignature.IsUnknown() && !r.Config.VerifyChannelTokenSignature.IsNull() {
		*verifyChannelTokenSignature = r.Config.VerifyChannelTokenSignature.ValueBool()
	} else {
		verifyChannelTokenSignature = nil
	}
	verifyChannelTokenExpiry := new(bool)
	if !r.Config.VerifyChannelTokenExpiry.IsUnknown() && !r.Config.VerifyChannelTokenExpiry.IsNull() {
		*verifyChannelTokenExpiry = r.Config.VerifyChannelTokenExpiry.ValueBool()
	} else {
		verifyChannelTokenExpiry = nil
	}
	verifyChannelTokenScopes := new(bool)
	if !r.Config.VerifyChannelTokenScopes.IsUnknown() && !r.Config.VerifyChannelTokenScopes.IsNull() {
		*verifyChannelTokenScopes = r.Config.VerifyChannelTokenScopes.ValueBool()
	} else {
		verifyChannelTokenScopes = nil
	}
	verifyChannelTokenIntrospectionExpiry := new(bool)
	if !r.Config.VerifyChannelTokenIntrospectionExpiry.IsUnknown() && !r.Config.VerifyChannelTokenIntrospectionExpiry.IsNull() {
		*verifyChannelTokenIntrospectionExpiry = r.Config.VerifyChannelTokenIntrospectionExpiry.ValueBool()
	} else {
		verifyChannelTokenIntrospectionExpiry = nil
	}
	verifyChannelTokenIntrospectionScopes := new(bool)
	if !r.Config.VerifyChannelTokenIntrospectionScopes.IsUnknown() && !r.Config.VerifyChannelTokenIntrospectionScopes.IsNull() {
		*verifyChannelTokenIntrospectionScopes = r.Config.VerifyChannelTokenIntrospectionScopes.ValueBool()
	} else {
		verifyChannelTokenIntrospectionScopes = nil
	}
	cacheChannelTokenIntrospection := new(bool)
	if !r.Config.CacheChannelTokenIntrospection.IsUnknown() && !r.Config.CacheChannelTokenIntrospection.IsNull() {
		*cacheChannelTokenIntrospection = r.Config.CacheChannelTokenIntrospection.ValueBool()
	} else {
		cacheChannelTokenIntrospection = nil
	}
	trustChannelTokenIntrospection := new(bool)
	if !r.Config.TrustChannelTokenIntrospection.IsUnknown() && !r.Config.TrustChannelTokenIntrospection.IsNull() {
		*trustChannelTokenIntrospection = r.Config.TrustChannelTokenIntrospection.ValueBool()
	} else {
		trustChannelTokenIntrospection = nil
	}
	enableChannelTokenIntrospection := new(bool)
	if !r.Config.EnableChannelTokenIntrospection.IsUnknown() && !r.Config.EnableChannelTokenIntrospection.IsNull() {
		*enableChannelTokenIntrospection = r.Config.EnableChannelTokenIntrospection.ValueBool()
	} else {
		enableChannelTokenIntrospection = nil
	}
	addClaims := make(map[string]interface{})
	for addClaimsKey, addClaimsValue := range r.Config.AddClaims {
		var addClaimsInst interface{}
		_ = json.Unmarshal([]byte(addClaimsValue.ValueString()), &addClaimsInst)
		addClaims[addClaimsKey] = addClaimsInst
	}
	setClaims := make(map[string]interface{})
	for setClaimsKey, setClaimsValue := range r.Config.SetClaims {
		var setClaimsInst interface{}
		_ = json.Unmarshal([]byte(setClaimsValue.ValueString()), &setClaimsInst)
		setClaims[setClaimsKey] = setClaimsInst
	}
	config := shared.CreateJWTSignerPluginConfig{
		Realm:                                   realm,
		EnableHsSignatures:                      enableHsSignatures,
		EnableInstrumentation:                   enableInstrumentation,
		AccessTokenIssuer:                       accessTokenIssuer,
		AccessTokenKeyset:                       accessTokenKeyset,
		AccessTokenJwksURI:                      accessTokenJwksURI,
		AccessTokenRequestHeader:                accessTokenRequestHeader,
		AccessTokenLeeway:                       accessTokenLeeway,
		AccessTokenScopesRequired:               accessTokenScopesRequired,
		AccessTokenScopesClaim:                  accessTokenScopesClaim,
		AccessTokenConsumerClaim:                accessTokenConsumerClaim,
		AccessTokenConsumerBy:                   accessTokenConsumerBy,
		AccessTokenUpstreamHeader:               accessTokenUpstreamHeader,
		AccessTokenUpstreamLeeway:               accessTokenUpstreamLeeway,
		AccessTokenIntrospectionEndpoint:        accessTokenIntrospectionEndpoint,
		AccessTokenIntrospectionAuthorization:   accessTokenIntrospectionAuthorization,
		AccessTokenIntrospectionBodyArgs:        accessTokenIntrospectionBodyArgs,
		AccessTokenIntrospectionHint:            accessTokenIntrospectionHint,
		AccessTokenIntrospectionJwtClaim:        accessTokenIntrospectionJwtClaim,
		AccessTokenIntrospectionScopesRequired:  accessTokenIntrospectionScopesRequired,
		AccessTokenIntrospectionScopesClaim:     accessTokenIntrospectionScopesClaim,
		AccessTokenIntrospectionConsumerClaim:   accessTokenIntrospectionConsumerClaim,
		AccessTokenIntrospectionConsumerBy:      accessTokenIntrospectionConsumerBy,
		AccessTokenIntrospectionLeeway:          accessTokenIntrospectionLeeway,
		AccessTokenIntrospectionTimeout:         accessTokenIntrospectionTimeout,
		AccessTokenSigningAlgorithm:             accessTokenSigningAlgorithm,
		AccessTokenOptional:                     accessTokenOptional,
		VerifyAccessTokenSignature:              verifyAccessTokenSignature,
		VerifyAccessTokenExpiry:                 verifyAccessTokenExpiry,
		VerifyAccessTokenScopes:                 verifyAccessTokenScopes,
		VerifyAccessTokenIntrospectionExpiry:    verifyAccessTokenIntrospectionExpiry,
		VerifyAccessTokenIntrospectionScopes:    verifyAccessTokenIntrospectionScopes,
		CacheAccessTokenIntrospection:           cacheAccessTokenIntrospection,
		TrustAccessTokenIntrospection:           trustAccessTokenIntrospection,
		EnableAccessTokenIntrospection:          enableAccessTokenIntrospection,
		ChannelTokenIssuer:                      channelTokenIssuer,
		ChannelTokenKeyset:                      channelTokenKeyset,
		ChannelTokenJwksURI:                     channelTokenJwksURI,
		ChannelTokenRequestHeader:               channelTokenRequestHeader,
		ChannelTokenLeeway:                      channelTokenLeeway,
		ChannelTokenScopesRequired:              channelTokenScopesRequired,
		ChannelTokenScopesClaim:                 channelTokenScopesClaim,
		ChannelTokenConsumerClaim:               channelTokenConsumerClaim,
		ChannelTokenConsumerBy:                  channelTokenConsumerBy,
		ChannelTokenUpstreamHeader:              channelTokenUpstreamHeader,
		ChannelTokenUpstreamLeeway:              channelTokenUpstreamLeeway,
		ChannelTokenIntrospectionEndpoint:       channelTokenIntrospectionEndpoint,
		ChannelTokenIntrospectionAuthorization:  channelTokenIntrospectionAuthorization,
		ChannelTokenIntrospectionBodyArgs:       channelTokenIntrospectionBodyArgs,
		ChannelTokenIntrospectionHint:           channelTokenIntrospectionHint,
		ChannelTokenIntrospectionJwtClaim:       channelTokenIntrospectionJwtClaim,
		ChannelTokenIntrospectionScopesRequired: channelTokenIntrospectionScopesRequired,
		ChannelTokenIntrospectionScopesClaim:    channelTokenIntrospectionScopesClaim,
		ChannelTokenIntrospectionConsumerClaim:  channelTokenIntrospectionConsumerClaim,
		ChannelTokenIntrospectionConsumerBy:     channelTokenIntrospectionConsumerBy,
		ChannelTokenIntrospectionLeeway:         channelTokenIntrospectionLeeway,
		ChannelTokenIntrospectionTimeout:        channelTokenIntrospectionTimeout,
		ChannelTokenSigningAlgorithm:            channelTokenSigningAlgorithm,
		ChannelTokenOptional:                    channelTokenOptional,
		VerifyChannelTokenSignature:             verifyChannelTokenSignature,
		VerifyChannelTokenExpiry:                verifyChannelTokenExpiry,
		VerifyChannelTokenScopes:                verifyChannelTokenScopes,
		VerifyChannelTokenIntrospectionExpiry:   verifyChannelTokenIntrospectionExpiry,
		VerifyChannelTokenIntrospectionScopes:   verifyChannelTokenIntrospectionScopes,
		CacheChannelTokenIntrospection:          cacheChannelTokenIntrospection,
		TrustChannelTokenIntrospection:          trustChannelTokenIntrospection,
		EnableChannelTokenIntrospection:         enableChannelTokenIntrospection,
		AddClaims:                               addClaims,
		SetClaims:                               setClaims,
	}
	out := shared.CreateJWTSignerPlugin{
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

func (r *GatewayPluginJWTSignerResourceModel) RefreshFromSharedJWTSignerPlugin(resp *shared.JWTSignerPlugin) {
	if resp != nil {
		r.Config.AccessTokenConsumerBy = []types.String{}
		for _, v := range resp.Config.AccessTokenConsumerBy {
			r.Config.AccessTokenConsumerBy = append(r.Config.AccessTokenConsumerBy, types.StringValue(string(v)))
		}
		r.Config.AccessTokenConsumerClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenConsumerClaim {
			r.Config.AccessTokenConsumerClaim = append(r.Config.AccessTokenConsumerClaim, types.StringValue(v))
		}
		r.Config.AccessTokenIntrospectionAuthorization = types.StringPointerValue(resp.Config.AccessTokenIntrospectionAuthorization)
		r.Config.AccessTokenIntrospectionBodyArgs = types.StringPointerValue(resp.Config.AccessTokenIntrospectionBodyArgs)
		r.Config.AccessTokenIntrospectionConsumerBy = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionConsumerBy {
			r.Config.AccessTokenIntrospectionConsumerBy = append(r.Config.AccessTokenIntrospectionConsumerBy, types.StringValue(string(v)))
		}
		r.Config.AccessTokenIntrospectionConsumerClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionConsumerClaim {
			r.Config.AccessTokenIntrospectionConsumerClaim = append(r.Config.AccessTokenIntrospectionConsumerClaim, types.StringValue(v))
		}
		r.Config.AccessTokenIntrospectionEndpoint = types.StringPointerValue(resp.Config.AccessTokenIntrospectionEndpoint)
		r.Config.AccessTokenIntrospectionHint = types.StringPointerValue(resp.Config.AccessTokenIntrospectionHint)
		r.Config.AccessTokenIntrospectionJwtClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionJwtClaim {
			r.Config.AccessTokenIntrospectionJwtClaim = append(r.Config.AccessTokenIntrospectionJwtClaim, types.StringValue(v))
		}
		if resp.Config.AccessTokenIntrospectionLeeway != nil {
			r.Config.AccessTokenIntrospectionLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenIntrospectionLeeway)))
		} else {
			r.Config.AccessTokenIntrospectionLeeway = types.NumberNull()
		}
		r.Config.AccessTokenIntrospectionScopesClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionScopesClaim {
			r.Config.AccessTokenIntrospectionScopesClaim = append(r.Config.AccessTokenIntrospectionScopesClaim, types.StringValue(v))
		}
		r.Config.AccessTokenIntrospectionScopesRequired = []types.String{}
		for _, v := range resp.Config.AccessTokenIntrospectionScopesRequired {
			r.Config.AccessTokenIntrospectionScopesRequired = append(r.Config.AccessTokenIntrospectionScopesRequired, types.StringValue(v))
		}
		if resp.Config.AccessTokenIntrospectionTimeout != nil {
			r.Config.AccessTokenIntrospectionTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenIntrospectionTimeout)))
		} else {
			r.Config.AccessTokenIntrospectionTimeout = types.NumberNull()
		}
		r.Config.AccessTokenIssuer = types.StringPointerValue(resp.Config.AccessTokenIssuer)
		r.Config.AccessTokenJwksURI = types.StringPointerValue(resp.Config.AccessTokenJwksURI)
		r.Config.AccessTokenKeyset = types.StringPointerValue(resp.Config.AccessTokenKeyset)
		if resp.Config.AccessTokenLeeway != nil {
			r.Config.AccessTokenLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenLeeway)))
		} else {
			r.Config.AccessTokenLeeway = types.NumberNull()
		}
		r.Config.AccessTokenOptional = types.BoolPointerValue(resp.Config.AccessTokenOptional)
		r.Config.AccessTokenRequestHeader = types.StringPointerValue(resp.Config.AccessTokenRequestHeader)
		r.Config.AccessTokenScopesClaim = []types.String{}
		for _, v := range resp.Config.AccessTokenScopesClaim {
			r.Config.AccessTokenScopesClaim = append(r.Config.AccessTokenScopesClaim, types.StringValue(v))
		}
		r.Config.AccessTokenScopesRequired = []types.String{}
		for _, v := range resp.Config.AccessTokenScopesRequired {
			r.Config.AccessTokenScopesRequired = append(r.Config.AccessTokenScopesRequired, types.StringValue(v))
		}
		if resp.Config.AccessTokenSigningAlgorithm != nil {
			r.Config.AccessTokenSigningAlgorithm = types.StringValue(string(*resp.Config.AccessTokenSigningAlgorithm))
		} else {
			r.Config.AccessTokenSigningAlgorithm = types.StringNull()
		}
		r.Config.AccessTokenUpstreamHeader = types.StringPointerValue(resp.Config.AccessTokenUpstreamHeader)
		if resp.Config.AccessTokenUpstreamLeeway != nil {
			r.Config.AccessTokenUpstreamLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.AccessTokenUpstreamLeeway)))
		} else {
			r.Config.AccessTokenUpstreamLeeway = types.NumberNull()
		}
		if len(resp.Config.AddClaims) > 0 {
			r.Config.AddClaims = make(map[string]types.String)
			for key, value := range resp.Config.AddClaims {
				result, _ := json.Marshal(value)
				r.Config.AddClaims[key] = types.StringValue(string(result))
			}
		}
		r.Config.CacheAccessTokenIntrospection = types.BoolPointerValue(resp.Config.CacheAccessTokenIntrospection)
		r.Config.CacheChannelTokenIntrospection = types.BoolPointerValue(resp.Config.CacheChannelTokenIntrospection)
		r.Config.ChannelTokenConsumerBy = []types.String{}
		for _, v := range resp.Config.ChannelTokenConsumerBy {
			r.Config.ChannelTokenConsumerBy = append(r.Config.ChannelTokenConsumerBy, types.StringValue(string(v)))
		}
		r.Config.ChannelTokenConsumerClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenConsumerClaim {
			r.Config.ChannelTokenConsumerClaim = append(r.Config.ChannelTokenConsumerClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenIntrospectionAuthorization = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionAuthorization)
		r.Config.ChannelTokenIntrospectionBodyArgs = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionBodyArgs)
		r.Config.ChannelTokenIntrospectionConsumerBy = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionConsumerBy {
			r.Config.ChannelTokenIntrospectionConsumerBy = append(r.Config.ChannelTokenIntrospectionConsumerBy, types.StringValue(string(v)))
		}
		r.Config.ChannelTokenIntrospectionConsumerClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionConsumerClaim {
			r.Config.ChannelTokenIntrospectionConsumerClaim = append(r.Config.ChannelTokenIntrospectionConsumerClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenIntrospectionEndpoint = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionEndpoint)
		r.Config.ChannelTokenIntrospectionHint = types.StringPointerValue(resp.Config.ChannelTokenIntrospectionHint)
		r.Config.ChannelTokenIntrospectionJwtClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionJwtClaim {
			r.Config.ChannelTokenIntrospectionJwtClaim = append(r.Config.ChannelTokenIntrospectionJwtClaim, types.StringValue(v))
		}
		if resp.Config.ChannelTokenIntrospectionLeeway != nil {
			r.Config.ChannelTokenIntrospectionLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenIntrospectionLeeway)))
		} else {
			r.Config.ChannelTokenIntrospectionLeeway = types.NumberNull()
		}
		r.Config.ChannelTokenIntrospectionScopesClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionScopesClaim {
			r.Config.ChannelTokenIntrospectionScopesClaim = append(r.Config.ChannelTokenIntrospectionScopesClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenIntrospectionScopesRequired = []types.String{}
		for _, v := range resp.Config.ChannelTokenIntrospectionScopesRequired {
			r.Config.ChannelTokenIntrospectionScopesRequired = append(r.Config.ChannelTokenIntrospectionScopesRequired, types.StringValue(v))
		}
		if resp.Config.ChannelTokenIntrospectionTimeout != nil {
			r.Config.ChannelTokenIntrospectionTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenIntrospectionTimeout)))
		} else {
			r.Config.ChannelTokenIntrospectionTimeout = types.NumberNull()
		}
		r.Config.ChannelTokenIssuer = types.StringPointerValue(resp.Config.ChannelTokenIssuer)
		r.Config.ChannelTokenJwksURI = types.StringPointerValue(resp.Config.ChannelTokenJwksURI)
		r.Config.ChannelTokenKeyset = types.StringPointerValue(resp.Config.ChannelTokenKeyset)
		if resp.Config.ChannelTokenLeeway != nil {
			r.Config.ChannelTokenLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenLeeway)))
		} else {
			r.Config.ChannelTokenLeeway = types.NumberNull()
		}
		r.Config.ChannelTokenOptional = types.BoolPointerValue(resp.Config.ChannelTokenOptional)
		r.Config.ChannelTokenRequestHeader = types.StringPointerValue(resp.Config.ChannelTokenRequestHeader)
		r.Config.ChannelTokenScopesClaim = []types.String{}
		for _, v := range resp.Config.ChannelTokenScopesClaim {
			r.Config.ChannelTokenScopesClaim = append(r.Config.ChannelTokenScopesClaim, types.StringValue(v))
		}
		r.Config.ChannelTokenScopesRequired = []types.String{}
		for _, v := range resp.Config.ChannelTokenScopesRequired {
			r.Config.ChannelTokenScopesRequired = append(r.Config.ChannelTokenScopesRequired, types.StringValue(v))
		}
		if resp.Config.ChannelTokenSigningAlgorithm != nil {
			r.Config.ChannelTokenSigningAlgorithm = types.StringValue(string(*resp.Config.ChannelTokenSigningAlgorithm))
		} else {
			r.Config.ChannelTokenSigningAlgorithm = types.StringNull()
		}
		r.Config.ChannelTokenUpstreamHeader = types.StringPointerValue(resp.Config.ChannelTokenUpstreamHeader)
		if resp.Config.ChannelTokenUpstreamLeeway != nil {
			r.Config.ChannelTokenUpstreamLeeway = types.NumberValue(big.NewFloat(float64(*resp.Config.ChannelTokenUpstreamLeeway)))
		} else {
			r.Config.ChannelTokenUpstreamLeeway = types.NumberNull()
		}
		r.Config.EnableAccessTokenIntrospection = types.BoolPointerValue(resp.Config.EnableAccessTokenIntrospection)
		r.Config.EnableChannelTokenIntrospection = types.BoolPointerValue(resp.Config.EnableChannelTokenIntrospection)
		r.Config.EnableHsSignatures = types.BoolPointerValue(resp.Config.EnableHsSignatures)
		r.Config.EnableInstrumentation = types.BoolPointerValue(resp.Config.EnableInstrumentation)
		r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
		if len(resp.Config.SetClaims) > 0 {
			r.Config.SetClaims = make(map[string]types.String)
			for key1, value1 := range resp.Config.SetClaims {
				result1, _ := json.Marshal(value1)
				r.Config.SetClaims[key1] = types.StringValue(string(result1))
			}
		}
		r.Config.TrustAccessTokenIntrospection = types.BoolPointerValue(resp.Config.TrustAccessTokenIntrospection)
		r.Config.TrustChannelTokenIntrospection = types.BoolPointerValue(resp.Config.TrustChannelTokenIntrospection)
		r.Config.VerifyAccessTokenExpiry = types.BoolPointerValue(resp.Config.VerifyAccessTokenExpiry)
		r.Config.VerifyAccessTokenIntrospectionExpiry = types.BoolPointerValue(resp.Config.VerifyAccessTokenIntrospectionExpiry)
		r.Config.VerifyAccessTokenIntrospectionScopes = types.BoolPointerValue(resp.Config.VerifyAccessTokenIntrospectionScopes)
		r.Config.VerifyAccessTokenScopes = types.BoolPointerValue(resp.Config.VerifyAccessTokenScopes)
		r.Config.VerifyAccessTokenSignature = types.BoolPointerValue(resp.Config.VerifyAccessTokenSignature)
		r.Config.VerifyChannelTokenExpiry = types.BoolPointerValue(resp.Config.VerifyChannelTokenExpiry)
		r.Config.VerifyChannelTokenIntrospectionExpiry = types.BoolPointerValue(resp.Config.VerifyChannelTokenIntrospectionExpiry)
		r.Config.VerifyChannelTokenIntrospectionScopes = types.BoolPointerValue(resp.Config.VerifyChannelTokenIntrospectionScopes)
		r.Config.VerifyChannelTokenScopes = types.BoolPointerValue(resp.Config.VerifyChannelTokenScopes)
		r.Config.VerifyChannelTokenSignature = types.BoolPointerValue(resp.Config.VerifyChannelTokenSignature)
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
