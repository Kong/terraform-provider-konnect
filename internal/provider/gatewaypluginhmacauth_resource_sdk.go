// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginHmacAuthResourceModel) ToSharedHmacAuthPluginInput() *shared.HmacAuthPluginInput {
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
	var ordering *shared.HmacAuthPluginOrdering
	if r.Ordering != nil {
		var after *shared.HmacAuthPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.HmacAuthPluginAfter{
				Access: access,
			}
		}
		var before *shared.HmacAuthPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.HmacAuthPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.HmacAuthPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var algorithms []shared.Algorithms = []shared.Algorithms{}
	for _, algorithmsItem := range r.Config.Algorithms {
		algorithms = append(algorithms, shared.Algorithms(algorithmsItem.ValueString()))
	}
	anonymous := new(string)
	if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
		*anonymous = r.Config.Anonymous.ValueString()
	} else {
		anonymous = nil
	}
	clockSkew := new(float64)
	if !r.Config.ClockSkew.IsUnknown() && !r.Config.ClockSkew.IsNull() {
		*clockSkew, _ = r.Config.ClockSkew.ValueBigFloat().Float64()
	} else {
		clockSkew = nil
	}
	var enforceHeaders []string = []string{}
	for _, enforceHeadersItem := range r.Config.EnforceHeaders {
		enforceHeaders = append(enforceHeaders, enforceHeadersItem.ValueString())
	}
	hideCredentials := new(bool)
	if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
		*hideCredentials = r.Config.HideCredentials.ValueBool()
	} else {
		hideCredentials = nil
	}
	realm := new(string)
	if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
		*realm = r.Config.Realm.ValueString()
	} else {
		realm = nil
	}
	validateRequestBody := new(bool)
	if !r.Config.ValidateRequestBody.IsUnknown() && !r.Config.ValidateRequestBody.IsNull() {
		*validateRequestBody = r.Config.ValidateRequestBody.ValueBool()
	} else {
		validateRequestBody = nil
	}
	config := shared.HmacAuthPluginConfig{
		Algorithms:          algorithms,
		Anonymous:           anonymous,
		ClockSkew:           clockSkew,
		EnforceHeaders:      enforceHeaders,
		HideCredentials:     hideCredentials,
		Realm:               realm,
		ValidateRequestBody: validateRequestBody,
	}
	var protocols []shared.HmacAuthPluginProtocols = []shared.HmacAuthPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.HmacAuthPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.HmacAuthPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.HmacAuthPluginRoute{
			ID: id1,
		}
	}
	var service *shared.HmacAuthPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.HmacAuthPluginService{
			ID: id2,
		}
	}
	out := shared.HmacAuthPluginInput{
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

func (r *GatewayPluginHmacAuthResourceModel) RefreshFromSharedHmacAuthPlugin(resp *shared.HmacAuthPlugin) {
	if resp != nil {
		r.Config.Algorithms = []types.String{}
		for _, v := range resp.Config.Algorithms {
			r.Config.Algorithms = append(r.Config.Algorithms, types.StringValue(string(v)))
		}
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		if resp.Config.ClockSkew != nil {
			r.Config.ClockSkew = types.NumberValue(big.NewFloat(float64(*resp.Config.ClockSkew)))
		} else {
			r.Config.ClockSkew = types.NumberNull()
		}
		r.Config.EnforceHeaders = []types.String{}
		for _, v := range resp.Config.EnforceHeaders {
			r.Config.EnforceHeaders = append(r.Config.EnforceHeaders, types.StringValue(v))
		}
		r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
		r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
		r.Config.ValidateRequestBody = types.BoolPointerValue(resp.Config.ValidateRequestBody)
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
