// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginOauth2DataSourceModel) RefreshFromSharedOauth2Plugin(resp *shared.Oauth2Plugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateOauth2PluginConfig{}
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
			r.Ordering = &tfTypes.CreateACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.CreateACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.CreateACLPluginAfter{}
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
