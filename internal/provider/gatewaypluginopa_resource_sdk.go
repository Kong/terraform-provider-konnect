// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginOpaResourceModel) ToSharedOpaPlugin() *shared.OpaPlugin {
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
	var ordering *shared.OpaPluginOrdering
	if r.Ordering != nil {
		var after *shared.OpaPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.OpaPluginAfter{
				Access: access,
			}
		}
		var before *shared.OpaPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.OpaPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.OpaPluginOrdering{
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
	var config *shared.OpaPluginConfig
	if r.Config != nil {
		includeBodyInOpaInput := new(bool)
		if !r.Config.IncludeBodyInOpaInput.IsUnknown() && !r.Config.IncludeBodyInOpaInput.IsNull() {
			*includeBodyInOpaInput = r.Config.IncludeBodyInOpaInput.ValueBool()
		} else {
			includeBodyInOpaInput = nil
		}
		includeConsumerInOpaInput := new(bool)
		if !r.Config.IncludeConsumerInOpaInput.IsUnknown() && !r.Config.IncludeConsumerInOpaInput.IsNull() {
			*includeConsumerInOpaInput = r.Config.IncludeConsumerInOpaInput.ValueBool()
		} else {
			includeConsumerInOpaInput = nil
		}
		includeParsedJSONBodyInOpaInput := new(bool)
		if !r.Config.IncludeParsedJSONBodyInOpaInput.IsUnknown() && !r.Config.IncludeParsedJSONBodyInOpaInput.IsNull() {
			*includeParsedJSONBodyInOpaInput = r.Config.IncludeParsedJSONBodyInOpaInput.ValueBool()
		} else {
			includeParsedJSONBodyInOpaInput = nil
		}
		includeRouteInOpaInput := new(bool)
		if !r.Config.IncludeRouteInOpaInput.IsUnknown() && !r.Config.IncludeRouteInOpaInput.IsNull() {
			*includeRouteInOpaInput = r.Config.IncludeRouteInOpaInput.ValueBool()
		} else {
			includeRouteInOpaInput = nil
		}
		includeServiceInOpaInput := new(bool)
		if !r.Config.IncludeServiceInOpaInput.IsUnknown() && !r.Config.IncludeServiceInOpaInput.IsNull() {
			*includeServiceInOpaInput = r.Config.IncludeServiceInOpaInput.ValueBool()
		} else {
			includeServiceInOpaInput = nil
		}
		includeURICapturesInOpaInput := new(bool)
		if !r.Config.IncludeURICapturesInOpaInput.IsUnknown() && !r.Config.IncludeURICapturesInOpaInput.IsNull() {
			*includeURICapturesInOpaInput = r.Config.IncludeURICapturesInOpaInput.ValueBool()
		} else {
			includeURICapturesInOpaInput = nil
		}
		opaHost := new(string)
		if !r.Config.OpaHost.IsUnknown() && !r.Config.OpaHost.IsNull() {
			*opaHost = r.Config.OpaHost.ValueString()
		} else {
			opaHost = nil
		}
		opaPath := new(string)
		if !r.Config.OpaPath.IsUnknown() && !r.Config.OpaPath.IsNull() {
			*opaPath = r.Config.OpaPath.ValueString()
		} else {
			opaPath = nil
		}
		opaPort := new(int64)
		if !r.Config.OpaPort.IsUnknown() && !r.Config.OpaPort.IsNull() {
			*opaPort = r.Config.OpaPort.ValueInt64()
		} else {
			opaPort = nil
		}
		opaProtocol := new(shared.OpaProtocol)
		if !r.Config.OpaProtocol.IsUnknown() && !r.Config.OpaProtocol.IsNull() {
			*opaProtocol = shared.OpaProtocol(r.Config.OpaProtocol.ValueString())
		} else {
			opaProtocol = nil
		}
		sslVerify := new(bool)
		if !r.Config.SslVerify.IsUnknown() && !r.Config.SslVerify.IsNull() {
			*sslVerify = r.Config.SslVerify.ValueBool()
		} else {
			sslVerify = nil
		}
		config = &shared.OpaPluginConfig{
			IncludeBodyInOpaInput:           includeBodyInOpaInput,
			IncludeConsumerInOpaInput:       includeConsumerInOpaInput,
			IncludeParsedJSONBodyInOpaInput: includeParsedJSONBodyInOpaInput,
			IncludeRouteInOpaInput:          includeRouteInOpaInput,
			IncludeServiceInOpaInput:        includeServiceInOpaInput,
			IncludeURICapturesInOpaInput:    includeURICapturesInOpaInput,
			OpaHost:                         opaHost,
			OpaPath:                         opaPath,
			OpaPort:                         opaPort,
			OpaProtocol:                     opaProtocol,
			SslVerify:                       sslVerify,
		}
	}
	var protocols []shared.OpaPluginProtocols = []shared.OpaPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.OpaPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.OpaPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.OpaPluginRoute{
			ID: id1,
		}
	}
	var service *shared.OpaPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.OpaPluginService{
			ID: id2,
		}
	}
	out := shared.OpaPlugin{
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

func (r *GatewayPluginOpaResourceModel) RefreshFromSharedOpaPlugin(ctx context.Context, resp *shared.OpaPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.OpaPluginConfig{}
			r.Config.IncludeBodyInOpaInput = types.BoolPointerValue(resp.Config.IncludeBodyInOpaInput)
			r.Config.IncludeConsumerInOpaInput = types.BoolPointerValue(resp.Config.IncludeConsumerInOpaInput)
			r.Config.IncludeParsedJSONBodyInOpaInput = types.BoolPointerValue(resp.Config.IncludeParsedJSONBodyInOpaInput)
			r.Config.IncludeRouteInOpaInput = types.BoolPointerValue(resp.Config.IncludeRouteInOpaInput)
			r.Config.IncludeServiceInOpaInput = types.BoolPointerValue(resp.Config.IncludeServiceInOpaInput)
			r.Config.IncludeURICapturesInOpaInput = types.BoolPointerValue(resp.Config.IncludeURICapturesInOpaInput)
			r.Config.OpaHost = types.StringPointerValue(resp.Config.OpaHost)
			r.Config.OpaPath = types.StringPointerValue(resp.Config.OpaPath)
			r.Config.OpaPort = types.Int64PointerValue(resp.Config.OpaPort)
			if resp.Config.OpaProtocol != nil {
				r.Config.OpaProtocol = types.StringValue(string(*resp.Config.OpaProtocol))
			} else {
				r.Config.OpaProtocol = types.StringNull()
			}
			r.Config.SslVerify = types.BoolPointerValue(resp.Config.SslVerify)
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
