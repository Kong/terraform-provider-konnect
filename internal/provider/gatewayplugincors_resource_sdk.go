// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginCorsResourceModel) ToSharedCorsPlugin() *shared.CorsPlugin {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
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
	var ordering *shared.CorsPluginOrdering
	if r.Ordering != nil {
		var after *shared.CorsPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CorsPluginAfter{
				Access: access,
			}
		}
		var before *shared.CorsPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CorsPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CorsPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.CorsPluginConfig
	if r.Config != nil {
		credentials := new(bool)
		if !r.Config.Credentials.IsUnknown() && !r.Config.Credentials.IsNull() {
			*credentials = r.Config.Credentials.ValueBool()
		} else {
			credentials = nil
		}
		var exposedHeaders []string = []string{}
		for _, exposedHeadersItem := range r.Config.ExposedHeaders {
			exposedHeaders = append(exposedHeaders, exposedHeadersItem.ValueString())
		}
		var headers []string = []string{}
		for _, headersItem := range r.Config.Headers {
			headers = append(headers, headersItem.ValueString())
		}
		maxAge := new(float64)
		if !r.Config.MaxAge.IsUnknown() && !r.Config.MaxAge.IsNull() {
			*maxAge = r.Config.MaxAge.ValueFloat64()
		} else {
			maxAge = nil
		}
		var methods []shared.Methods = []shared.Methods{}
		for _, methodsItem := range r.Config.Methods {
			methods = append(methods, shared.Methods(methodsItem.ValueString()))
		}
		var origins []string = []string{}
		for _, originsItem := range r.Config.Origins {
			origins = append(origins, originsItem.ValueString())
		}
		preflightContinue := new(bool)
		if !r.Config.PreflightContinue.IsUnknown() && !r.Config.PreflightContinue.IsNull() {
			*preflightContinue = r.Config.PreflightContinue.ValueBool()
		} else {
			preflightContinue = nil
		}
		privateNetwork := new(bool)
		if !r.Config.PrivateNetwork.IsUnknown() && !r.Config.PrivateNetwork.IsNull() {
			*privateNetwork = r.Config.PrivateNetwork.ValueBool()
		} else {
			privateNetwork = nil
		}
		config = &shared.CorsPluginConfig{
			Credentials:       credentials,
			ExposedHeaders:    exposedHeaders,
			Headers:           headers,
			MaxAge:            maxAge,
			Methods:           methods,
			Origins:           origins,
			PreflightContinue: preflightContinue,
			PrivateNetwork:    privateNetwork,
		}
	}
	var protocols []shared.CorsPluginProtocols = []shared.CorsPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CorsPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.CorsPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CorsPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CorsPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CorsPluginService{
			ID: id2,
		}
	}
	out := shared.CorsPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginCorsResourceModel) RefreshFromSharedCorsPlugin(ctx context.Context, resp *shared.CorsPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CorsPluginConfig{}
			r.Config.Credentials = types.BoolPointerValue(resp.Config.Credentials)
			r.Config.ExposedHeaders = make([]types.String, 0, len(resp.Config.ExposedHeaders))
			for _, v := range resp.Config.ExposedHeaders {
				r.Config.ExposedHeaders = append(r.Config.ExposedHeaders, types.StringValue(v))
			}
			r.Config.Headers = make([]types.String, 0, len(resp.Config.Headers))
			for _, v := range resp.Config.Headers {
				r.Config.Headers = append(r.Config.Headers, types.StringValue(v))
			}
			r.Config.MaxAge = types.Float64PointerValue(resp.Config.MaxAge)
			r.Config.Methods = make([]types.String, 0, len(resp.Config.Methods))
			for _, v := range resp.Config.Methods {
				r.Config.Methods = append(r.Config.Methods, types.StringValue(string(v)))
			}
			r.Config.Origins = make([]types.String, 0, len(resp.Config.Origins))
			for _, v := range resp.Config.Origins {
				r.Config.Origins = append(r.Config.Origins, types.StringValue(v))
			}
			r.Config.PreflightContinue = types.BoolPointerValue(resp.Config.PreflightContinue)
			r.Config.PrivateNetwork = types.BoolPointerValue(resp.Config.PrivateNetwork)
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

	return diags
}
