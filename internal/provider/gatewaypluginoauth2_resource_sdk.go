// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginOauth2ResourceModel) ToSharedOauth2PluginInput() *shared.Oauth2PluginInput {
	acceptHTTPIfAlreadyTerminated := new(bool)
	if !r.Config.AcceptHTTPIfAlreadyTerminated.IsUnknown() && !r.Config.AcceptHTTPIfAlreadyTerminated.IsNull() {
		*acceptHTTPIfAlreadyTerminated = r.Config.AcceptHTTPIfAlreadyTerminated.ValueBool()
	} else {
		acceptHTTPIfAlreadyTerminated = nil
	}
	anonymous := new(string)
	if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
		*anonymous = r.Config.Anonymous.ValueString()
	} else {
		anonymous = nil
	}
	authHeaderName := new(string)
	if !r.Config.AuthHeaderName.IsUnknown() && !r.Config.AuthHeaderName.IsNull() {
		*authHeaderName = r.Config.AuthHeaderName.ValueString()
	} else {
		authHeaderName = nil
	}
	enableAuthorizationCode := new(bool)
	if !r.Config.EnableAuthorizationCode.IsUnknown() && !r.Config.EnableAuthorizationCode.IsNull() {
		*enableAuthorizationCode = r.Config.EnableAuthorizationCode.ValueBool()
	} else {
		enableAuthorizationCode = nil
	}
	enableClientCredentials := new(bool)
	if !r.Config.EnableClientCredentials.IsUnknown() && !r.Config.EnableClientCredentials.IsNull() {
		*enableClientCredentials = r.Config.EnableClientCredentials.ValueBool()
	} else {
		enableClientCredentials = nil
	}
	enableImplicitGrant := new(bool)
	if !r.Config.EnableImplicitGrant.IsUnknown() && !r.Config.EnableImplicitGrant.IsNull() {
		*enableImplicitGrant = r.Config.EnableImplicitGrant.ValueBool()
	} else {
		enableImplicitGrant = nil
	}
	enablePasswordGrant := new(bool)
	if !r.Config.EnablePasswordGrant.IsUnknown() && !r.Config.EnablePasswordGrant.IsNull() {
		*enablePasswordGrant = r.Config.EnablePasswordGrant.ValueBool()
	} else {
		enablePasswordGrant = nil
	}
	globalCredentials := new(bool)
	if !r.Config.GlobalCredentials.IsUnknown() && !r.Config.GlobalCredentials.IsNull() {
		*globalCredentials = r.Config.GlobalCredentials.ValueBool()
	} else {
		globalCredentials = nil
	}
	hideCredentials := new(bool)
	if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
		*hideCredentials = r.Config.HideCredentials.ValueBool()
	} else {
		hideCredentials = nil
	}
	mandatoryScope := new(bool)
	if !r.Config.MandatoryScope.IsUnknown() && !r.Config.MandatoryScope.IsNull() {
		*mandatoryScope = r.Config.MandatoryScope.ValueBool()
	} else {
		mandatoryScope = nil
	}
	persistentRefreshToken := new(bool)
	if !r.Config.PersistentRefreshToken.IsUnknown() && !r.Config.PersistentRefreshToken.IsNull() {
		*persistentRefreshToken = r.Config.PersistentRefreshToken.ValueBool()
	} else {
		persistentRefreshToken = nil
	}
	pkce := new(shared.Pkce)
	if !r.Config.Pkce.IsUnknown() && !r.Config.Pkce.IsNull() {
		*pkce = shared.Pkce(r.Config.Pkce.ValueString())
	} else {
		pkce = nil
	}
	provisionKey := new(string)
	if !r.Config.ProvisionKey.IsUnknown() && !r.Config.ProvisionKey.IsNull() {
		*provisionKey = r.Config.ProvisionKey.ValueString()
	} else {
		provisionKey = nil
	}
	realm := new(string)
	if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
		*realm = r.Config.Realm.ValueString()
	} else {
		realm = nil
	}
	refreshTokenTTL := new(float64)
	if !r.Config.RefreshTokenTTL.IsUnknown() && !r.Config.RefreshTokenTTL.IsNull() {
		*refreshTokenTTL, _ = r.Config.RefreshTokenTTL.ValueBigFloat().Float64()
	} else {
		refreshTokenTTL = nil
	}
	reuseRefreshToken := new(bool)
	if !r.Config.ReuseRefreshToken.IsUnknown() && !r.Config.ReuseRefreshToken.IsNull() {
		*reuseRefreshToken = r.Config.ReuseRefreshToken.ValueBool()
	} else {
		reuseRefreshToken = nil
	}
	var scopes []string = []string{}
	for _, scopesItem := range r.Config.Scopes {
		scopes = append(scopes, scopesItem.ValueString())
	}
	tokenExpiration := new(float64)
	if !r.Config.TokenExpiration.IsUnknown() && !r.Config.TokenExpiration.IsNull() {
		*tokenExpiration, _ = r.Config.TokenExpiration.ValueBigFloat().Float64()
	} else {
		tokenExpiration = nil
	}
	config := shared.Oauth2PluginConfig{
		AcceptHTTPIfAlreadyTerminated: acceptHTTPIfAlreadyTerminated,
		Anonymous:                     anonymous,
		AuthHeaderName:                authHeaderName,
		EnableAuthorizationCode:       enableAuthorizationCode,
		EnableClientCredentials:       enableClientCredentials,
		EnableImplicitGrant:           enableImplicitGrant,
		EnablePasswordGrant:           enablePasswordGrant,
		GlobalCredentials:             globalCredentials,
		HideCredentials:               hideCredentials,
		MandatoryScope:                mandatoryScope,
		PersistentRefreshToken:        persistentRefreshToken,
		Pkce:                          pkce,
		ProvisionKey:                  provisionKey,
		Realm:                         realm,
		RefreshTokenTTL:               refreshTokenTTL,
		ReuseRefreshToken:             reuseRefreshToken,
		Scopes:                        scopes,
		TokenExpiration:               tokenExpiration,
	}
	var consumer *shared.Oauth2PluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.Oauth2PluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.Oauth2PluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.Oauth2PluginConsumerGroup{
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
	var ordering *shared.Oauth2PluginOrdering
	if r.Ordering != nil {
		var after *shared.Oauth2PluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.Oauth2PluginAfter{
				Access: access,
			}
		}
		var before *shared.Oauth2PluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.Oauth2PluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.Oauth2PluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.Oauth2PluginProtocols = []shared.Oauth2PluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.Oauth2PluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.Oauth2PluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.Oauth2PluginRoute{
			ID: id3,
		}
	}
	var service *shared.Oauth2PluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.Oauth2PluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.Oauth2PluginInput{
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

func (r *GatewayPluginOauth2ResourceModel) RefreshFromSharedOauth2Plugin(resp *shared.Oauth2Plugin) {
	if resp != nil {
		r.Config.AcceptHTTPIfAlreadyTerminated = types.BoolPointerValue(resp.Config.AcceptHTTPIfAlreadyTerminated)
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		r.Config.AuthHeaderName = types.StringPointerValue(resp.Config.AuthHeaderName)
		r.Config.EnableAuthorizationCode = types.BoolPointerValue(resp.Config.EnableAuthorizationCode)
		r.Config.EnableClientCredentials = types.BoolPointerValue(resp.Config.EnableClientCredentials)
		r.Config.EnableImplicitGrant = types.BoolPointerValue(resp.Config.EnableImplicitGrant)
		r.Config.EnablePasswordGrant = types.BoolPointerValue(resp.Config.EnablePasswordGrant)
		r.Config.GlobalCredentials = types.BoolPointerValue(resp.Config.GlobalCredentials)
		r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
		r.Config.MandatoryScope = types.BoolPointerValue(resp.Config.MandatoryScope)
		r.Config.PersistentRefreshToken = types.BoolPointerValue(resp.Config.PersistentRefreshToken)
		if resp.Config.Pkce != nil {
			r.Config.Pkce = types.StringValue(string(*resp.Config.Pkce))
		} else {
			r.Config.Pkce = types.StringNull()
		}
		r.Config.ProvisionKey = types.StringPointerValue(resp.Config.ProvisionKey)
		r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
		if resp.Config.RefreshTokenTTL != nil {
			r.Config.RefreshTokenTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.RefreshTokenTTL)))
		} else {
			r.Config.RefreshTokenTTL = types.NumberNull()
		}
		r.Config.ReuseRefreshToken = types.BoolPointerValue(resp.Config.ReuseRefreshToken)
		r.Config.Scopes = []types.String{}
		for _, v := range resp.Config.Scopes {
			r.Config.Scopes = append(r.Config.Scopes, types.StringValue(v))
		}
		if resp.Config.TokenExpiration != nil {
			r.Config.TokenExpiration = types.NumberValue(big.NewFloat(float64(*resp.Config.TokenExpiration)))
		} else {
			r.Config.TokenExpiration = types.NumberNull()
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
