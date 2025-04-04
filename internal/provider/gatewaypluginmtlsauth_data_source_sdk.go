// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginMtlsAuthDataSourceModel) RefreshFromSharedMtlsAuthPlugin(resp *shared.MtlsAuthPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.MtlsAuthPluginConfig{}
			r.Config.AllowPartialChain = types.BoolPointerValue(resp.Config.AllowPartialChain)
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			if resp.Config.AuthenticatedGroupBy != nil {
				r.Config.AuthenticatedGroupBy = types.StringValue(string(*resp.Config.AuthenticatedGroupBy))
			} else {
				r.Config.AuthenticatedGroupBy = types.StringNull()
			}
			r.Config.CaCertificates = make([]types.String, 0, len(resp.Config.CaCertificates))
			for _, v := range resp.Config.CaCertificates {
				r.Config.CaCertificates = append(r.Config.CaCertificates, types.StringValue(v))
			}
			if resp.Config.CacheTTL != nil {
				r.Config.CacheTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.CacheTTL)))
			} else {
				r.Config.CacheTTL = types.NumberNull()
			}
			if resp.Config.CertCacheTTL != nil {
				r.Config.CertCacheTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.CertCacheTTL)))
			} else {
				r.Config.CertCacheTTL = types.NumberNull()
			}
			r.Config.ConsumerBy = make([]types.String, 0, len(resp.Config.ConsumerBy))
			for _, v := range resp.Config.ConsumerBy {
				r.Config.ConsumerBy = append(r.Config.ConsumerBy, types.StringValue(string(v)))
			}
			r.Config.DefaultConsumer = types.StringPointerValue(resp.Config.DefaultConsumer)
			r.Config.HTTPProxyHost = types.StringPointerValue(resp.Config.HTTPProxyHost)
			r.Config.HTTPProxyPort = types.Int64PointerValue(resp.Config.HTTPProxyPort)
			if resp.Config.HTTPTimeout != nil {
				r.Config.HTTPTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.HTTPTimeout)))
			} else {
				r.Config.HTTPTimeout = types.NumberNull()
			}
			r.Config.HTTPSProxyHost = types.StringPointerValue(resp.Config.HTTPSProxyHost)
			r.Config.HTTPSProxyPort = types.Int64PointerValue(resp.Config.HTTPSProxyPort)
			if resp.Config.RevocationCheckMode != nil {
				r.Config.RevocationCheckMode = types.StringValue(string(*resp.Config.RevocationCheckMode))
			} else {
				r.Config.RevocationCheckMode = types.StringNull()
			}
			r.Config.SendCaDn = types.BoolPointerValue(resp.Config.SendCaDn)
			r.Config.SkipConsumerLookup = types.BoolPointerValue(resp.Config.SkipConsumerLookup)
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
}
