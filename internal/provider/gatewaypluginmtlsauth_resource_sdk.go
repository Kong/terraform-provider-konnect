// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginMtlsAuthResourceModel) ToSharedMtlsAuthPluginInput() *shared.MtlsAuthPluginInput {
	allowPartialChain := new(bool)
	if !r.Config.AllowPartialChain.IsUnknown() && !r.Config.AllowPartialChain.IsNull() {
		*allowPartialChain = r.Config.AllowPartialChain.ValueBool()
	} else {
		allowPartialChain = nil
	}
	anonymous := new(string)
	if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
		*anonymous = r.Config.Anonymous.ValueString()
	} else {
		anonymous = nil
	}
	authenticatedGroupBy := new(shared.AuthenticatedGroupBy)
	if !r.Config.AuthenticatedGroupBy.IsUnknown() && !r.Config.AuthenticatedGroupBy.IsNull() {
		*authenticatedGroupBy = shared.AuthenticatedGroupBy(r.Config.AuthenticatedGroupBy.ValueString())
	} else {
		authenticatedGroupBy = nil
	}
	var caCertificates []string = []string{}
	for _, caCertificatesItem := range r.Config.CaCertificates {
		caCertificates = append(caCertificates, caCertificatesItem.ValueString())
	}
	cacheTTL := new(float64)
	if !r.Config.CacheTTL.IsUnknown() && !r.Config.CacheTTL.IsNull() {
		*cacheTTL, _ = r.Config.CacheTTL.ValueBigFloat().Float64()
	} else {
		cacheTTL = nil
	}
	certCacheTTL := new(float64)
	if !r.Config.CertCacheTTL.IsUnknown() && !r.Config.CertCacheTTL.IsNull() {
		*certCacheTTL, _ = r.Config.CertCacheTTL.ValueBigFloat().Float64()
	} else {
		certCacheTTL = nil
	}
	var consumerBy []shared.MtlsAuthPluginConsumerBy = []shared.MtlsAuthPluginConsumerBy{}
	for _, consumerByItem := range r.Config.ConsumerBy {
		consumerBy = append(consumerBy, shared.MtlsAuthPluginConsumerBy(consumerByItem.ValueString()))
	}
	defaultConsumer := new(string)
	if !r.Config.DefaultConsumer.IsUnknown() && !r.Config.DefaultConsumer.IsNull() {
		*defaultConsumer = r.Config.DefaultConsumer.ValueString()
	} else {
		defaultConsumer = nil
	}
	httpProxyHost := new(string)
	if !r.Config.HTTPProxyHost.IsUnknown() && !r.Config.HTTPProxyHost.IsNull() {
		*httpProxyHost = r.Config.HTTPProxyHost.ValueString()
	} else {
		httpProxyHost = nil
	}
	httpProxyPort := new(int64)
	if !r.Config.HTTPProxyPort.IsUnknown() && !r.Config.HTTPProxyPort.IsNull() {
		*httpProxyPort = r.Config.HTTPProxyPort.ValueInt64()
	} else {
		httpProxyPort = nil
	}
	httpTimeout := new(float64)
	if !r.Config.HTTPTimeout.IsUnknown() && !r.Config.HTTPTimeout.IsNull() {
		*httpTimeout, _ = r.Config.HTTPTimeout.ValueBigFloat().Float64()
	} else {
		httpTimeout = nil
	}
	httpsProxyHost := new(string)
	if !r.Config.HTTPSProxyHost.IsUnknown() && !r.Config.HTTPSProxyHost.IsNull() {
		*httpsProxyHost = r.Config.HTTPSProxyHost.ValueString()
	} else {
		httpsProxyHost = nil
	}
	httpsProxyPort := new(int64)
	if !r.Config.HTTPSProxyPort.IsUnknown() && !r.Config.HTTPSProxyPort.IsNull() {
		*httpsProxyPort = r.Config.HTTPSProxyPort.ValueInt64()
	} else {
		httpsProxyPort = nil
	}
	revocationCheckMode := new(shared.RevocationCheckMode)
	if !r.Config.RevocationCheckMode.IsUnknown() && !r.Config.RevocationCheckMode.IsNull() {
		*revocationCheckMode = shared.RevocationCheckMode(r.Config.RevocationCheckMode.ValueString())
	} else {
		revocationCheckMode = nil
	}
	sendCaDn := new(bool)
	if !r.Config.SendCaDn.IsUnknown() && !r.Config.SendCaDn.IsNull() {
		*sendCaDn = r.Config.SendCaDn.ValueBool()
	} else {
		sendCaDn = nil
	}
	skipConsumerLookup := new(bool)
	if !r.Config.SkipConsumerLookup.IsUnknown() && !r.Config.SkipConsumerLookup.IsNull() {
		*skipConsumerLookup = r.Config.SkipConsumerLookup.ValueBool()
	} else {
		skipConsumerLookup = nil
	}
	config := shared.MtlsAuthPluginConfig{
		AllowPartialChain:    allowPartialChain,
		Anonymous:            anonymous,
		AuthenticatedGroupBy: authenticatedGroupBy,
		CaCertificates:       caCertificates,
		CacheTTL:             cacheTTL,
		CertCacheTTL:         certCacheTTL,
		ConsumerBy:           consumerBy,
		DefaultConsumer:      defaultConsumer,
		HTTPProxyHost:        httpProxyHost,
		HTTPProxyPort:        httpProxyPort,
		HTTPTimeout:          httpTimeout,
		HTTPSProxyHost:       httpsProxyHost,
		HTTPSProxyPort:       httpsProxyPort,
		RevocationCheckMode:  revocationCheckMode,
		SendCaDn:             sendCaDn,
		SkipConsumerLookup:   skipConsumerLookup,
	}
	var consumer *shared.MtlsAuthPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.MtlsAuthPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.MtlsAuthPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.MtlsAuthPluginConsumerGroup{
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
	var ordering *shared.MtlsAuthPluginOrdering
	if r.Ordering != nil {
		var after *shared.MtlsAuthPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.MtlsAuthPluginAfter{
				Access: access,
			}
		}
		var before *shared.MtlsAuthPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.MtlsAuthPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.MtlsAuthPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.MtlsAuthPluginProtocols = []shared.MtlsAuthPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.MtlsAuthPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.MtlsAuthPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.MtlsAuthPluginRoute{
			ID: id3,
		}
	}
	var service *shared.MtlsAuthPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.MtlsAuthPluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.MtlsAuthPluginInput{
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

func (r *GatewayPluginMtlsAuthResourceModel) RefreshFromSharedMtlsAuthPlugin(resp *shared.MtlsAuthPlugin) {
	if resp != nil {
		r.Config.AllowPartialChain = types.BoolPointerValue(resp.Config.AllowPartialChain)
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		if resp.Config.AuthenticatedGroupBy != nil {
			r.Config.AuthenticatedGroupBy = types.StringValue(string(*resp.Config.AuthenticatedGroupBy))
		} else {
			r.Config.AuthenticatedGroupBy = types.StringNull()
		}
		r.Config.CaCertificates = []types.String{}
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
		r.Config.ConsumerBy = []types.String{}
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
