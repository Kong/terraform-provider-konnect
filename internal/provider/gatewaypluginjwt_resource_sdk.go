// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginJWTResourceModel) ToSharedCreateJWTPlugin() *shared.CreateJWTPlugin {
	var config *shared.CreateJWTPluginConfig
	if r.Config != nil {
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		var claimsToVerify []shared.CreateJWTPluginClaimsToVerify = []shared.CreateJWTPluginClaimsToVerify{}
		for _, claimsToVerifyItem := range r.Config.ClaimsToVerify {
			claimsToVerify = append(claimsToVerify, shared.CreateJWTPluginClaimsToVerify(claimsToVerifyItem.ValueString()))
		}
		var cookieNames []string = []string{}
		for _, cookieNamesItem := range r.Config.CookieNames {
			cookieNames = append(cookieNames, cookieNamesItem.ValueString())
		}
		var headerNames []string = []string{}
		for _, headerNamesItem := range r.Config.HeaderNames {
			headerNames = append(headerNames, headerNamesItem.ValueString())
		}
		keyClaimName := new(string)
		if !r.Config.KeyClaimName.IsUnknown() && !r.Config.KeyClaimName.IsNull() {
			*keyClaimName = r.Config.KeyClaimName.ValueString()
		} else {
			keyClaimName = nil
		}
		maximumExpiration := new(float64)
		if !r.Config.MaximumExpiration.IsUnknown() && !r.Config.MaximumExpiration.IsNull() {
			*maximumExpiration, _ = r.Config.MaximumExpiration.ValueBigFloat().Float64()
		} else {
			maximumExpiration = nil
		}
		runOnPreflight := new(bool)
		if !r.Config.RunOnPreflight.IsUnknown() && !r.Config.RunOnPreflight.IsNull() {
			*runOnPreflight = r.Config.RunOnPreflight.ValueBool()
		} else {
			runOnPreflight = nil
		}
		secretIsBase64 := new(bool)
		if !r.Config.SecretIsBase64.IsUnknown() && !r.Config.SecretIsBase64.IsNull() {
			*secretIsBase64 = r.Config.SecretIsBase64.ValueBool()
		} else {
			secretIsBase64 = nil
		}
		var uriParamNames []string = []string{}
		for _, uriParamNamesItem := range r.Config.URIParamNames {
			uriParamNames = append(uriParamNames, uriParamNamesItem.ValueString())
		}
		config = &shared.CreateJWTPluginConfig{
			Anonymous:         anonymous,
			ClaimsToVerify:    claimsToVerify,
			CookieNames:       cookieNames,
			HeaderNames:       headerNames,
			KeyClaimName:      keyClaimName,
			MaximumExpiration: maximumExpiration,
			RunOnPreflight:    runOnPreflight,
			SecretIsBase64:    secretIsBase64,
			URIParamNames:     uriParamNames,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var protocols []shared.CreateJWTPluginProtocols = []shared.CreateJWTPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateJWTPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateJWTPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateJWTPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateJWTPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateJWTPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateJWTPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateJWTPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateJWTPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateJWTPluginService{
			ID: id3,
		}
	}
	out := shared.CreateJWTPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginJWTResourceModel) RefreshFromSharedJWTPlugin(resp *shared.JWTPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateJWTPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.ClaimsToVerify = []types.String{}
			for _, v := range resp.Config.ClaimsToVerify {
				r.Config.ClaimsToVerify = append(r.Config.ClaimsToVerify, types.StringValue(string(v)))
			}
			r.Config.CookieNames = []types.String{}
			for _, v := range resp.Config.CookieNames {
				r.Config.CookieNames = append(r.Config.CookieNames, types.StringValue(v))
			}
			r.Config.HeaderNames = []types.String{}
			for _, v := range resp.Config.HeaderNames {
				r.Config.HeaderNames = append(r.Config.HeaderNames, types.StringValue(v))
			}
			r.Config.KeyClaimName = types.StringPointerValue(resp.Config.KeyClaimName)
			if resp.Config.MaximumExpiration != nil {
				r.Config.MaximumExpiration = types.NumberValue(big.NewFloat(float64(*resp.Config.MaximumExpiration)))
			} else {
				r.Config.MaximumExpiration = types.NumberNull()
			}
			r.Config.RunOnPreflight = types.BoolPointerValue(resp.Config.RunOnPreflight)
			r.Config.SecretIsBase64 = types.BoolPointerValue(resp.Config.SecretIsBase64)
			r.Config.URIParamNames = []types.String{}
			for _, v := range resp.Config.URIParamNames {
				r.Config.URIParamNames = append(r.Config.URIParamNames, types.StringValue(v))
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
