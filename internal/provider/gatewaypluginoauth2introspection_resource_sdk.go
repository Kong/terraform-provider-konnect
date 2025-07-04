// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginOauth2IntrospectionResourceModel) ToSharedOauth2IntrospectionPlugin(ctx context.Context) (*shared.Oauth2IntrospectionPlugin, diag.Diagnostics) {
	var diags diag.Diagnostics

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
	var ordering *shared.Oauth2IntrospectionPluginOrdering
	if r.Ordering != nil {
		var after *shared.Oauth2IntrospectionPluginAfter
		if r.Ordering.After != nil {
			access := make([]string, 0, len(r.Ordering.After.Access))
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.Oauth2IntrospectionPluginAfter{
				Access: access,
			}
		}
		var before *shared.Oauth2IntrospectionPluginBefore
		if r.Ordering.Before != nil {
			access1 := make([]string, 0, len(r.Ordering.Before.Access))
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.Oauth2IntrospectionPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.Oauth2IntrospectionPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.Oauth2IntrospectionPluginPartials
	if r.Partials != nil {
		partials = make([]shared.Oauth2IntrospectionPluginPartials, 0, len(r.Partials))
		for _, partialsItem := range r.Partials {
			id1 := new(string)
			if !partialsItem.ID.IsUnknown() && !partialsItem.ID.IsNull() {
				*id1 = partialsItem.ID.ValueString()
			} else {
				id1 = nil
			}
			name := new(string)
			if !partialsItem.Name.IsUnknown() && !partialsItem.Name.IsNull() {
				*name = partialsItem.Name.ValueString()
			} else {
				name = nil
			}
			path := new(string)
			if !partialsItem.Path.IsUnknown() && !partialsItem.Path.IsNull() {
				*path = partialsItem.Path.ValueString()
			} else {
				path = nil
			}
			partials = append(partials, shared.Oauth2IntrospectionPluginPartials{
				ID:   id1,
				Name: name,
				Path: path,
			})
		}
	}
	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.Oauth2IntrospectionPluginConfig
	if r.Config != nil {
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		authorizationValue := new(string)
		if !r.Config.AuthorizationValue.IsUnknown() && !r.Config.AuthorizationValue.IsNull() {
			*authorizationValue = r.Config.AuthorizationValue.ValueString()
		} else {
			authorizationValue = nil
		}
		consumerBy := new(shared.Oauth2IntrospectionPluginConsumerBy)
		if !r.Config.ConsumerBy.IsUnknown() && !r.Config.ConsumerBy.IsNull() {
			*consumerBy = shared.Oauth2IntrospectionPluginConsumerBy(r.Config.ConsumerBy.ValueString())
		} else {
			consumerBy = nil
		}
		customClaimsForward := make([]string, 0, len(r.Config.CustomClaimsForward))
		for _, customClaimsForwardItem := range r.Config.CustomClaimsForward {
			customClaimsForward = append(customClaimsForward, customClaimsForwardItem.ValueString())
		}
		customIntrospectionHeaders := make(map[string]interface{})
		for customIntrospectionHeadersKey, customIntrospectionHeadersValue := range r.Config.CustomIntrospectionHeaders {
			var customIntrospectionHeadersInst interface{}
			_ = json.Unmarshal([]byte(customIntrospectionHeadersValue.ValueString()), &customIntrospectionHeadersInst)
			customIntrospectionHeaders[customIntrospectionHeadersKey] = customIntrospectionHeadersInst
		}
		hideCredentials := new(bool)
		if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
			*hideCredentials = r.Config.HideCredentials.ValueBool()
		} else {
			hideCredentials = nil
		}
		introspectRequest := new(bool)
		if !r.Config.IntrospectRequest.IsUnknown() && !r.Config.IntrospectRequest.IsNull() {
			*introspectRequest = r.Config.IntrospectRequest.ValueBool()
		} else {
			introspectRequest = nil
		}
		introspectionURL := new(string)
		if !r.Config.IntrospectionURL.IsUnknown() && !r.Config.IntrospectionURL.IsNull() {
			*introspectionURL = r.Config.IntrospectionURL.ValueString()
		} else {
			introspectionURL = nil
		}
		keepalive := new(int64)
		if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
			*keepalive = r.Config.Keepalive.ValueInt64()
		} else {
			keepalive = nil
		}
		runOnPreflight := new(bool)
		if !r.Config.RunOnPreflight.IsUnknown() && !r.Config.RunOnPreflight.IsNull() {
			*runOnPreflight = r.Config.RunOnPreflight.ValueBool()
		} else {
			runOnPreflight = nil
		}
		timeout := new(int64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout = r.Config.Timeout.ValueInt64()
		} else {
			timeout = nil
		}
		tokenTypeHint := new(string)
		if !r.Config.TokenTypeHint.IsUnknown() && !r.Config.TokenTypeHint.IsNull() {
			*tokenTypeHint = r.Config.TokenTypeHint.ValueString()
		} else {
			tokenTypeHint = nil
		}
		ttl := new(float64)
		if !r.Config.TTL.IsUnknown() && !r.Config.TTL.IsNull() {
			*ttl = r.Config.TTL.ValueFloat64()
		} else {
			ttl = nil
		}
		config = &shared.Oauth2IntrospectionPluginConfig{
			Anonymous:                  anonymous,
			AuthorizationValue:         authorizationValue,
			ConsumerBy:                 consumerBy,
			CustomClaimsForward:        customClaimsForward,
			CustomIntrospectionHeaders: customIntrospectionHeaders,
			HideCredentials:            hideCredentials,
			IntrospectRequest:          introspectRequest,
			IntrospectionURL:           introspectionURL,
			Keepalive:                  keepalive,
			RunOnPreflight:             runOnPreflight,
			Timeout:                    timeout,
			TokenTypeHint:              tokenTypeHint,
			TTL:                        ttl,
		}
	}
	protocols := make([]shared.Oauth2IntrospectionPluginProtocols, 0, len(r.Protocols))
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.Oauth2IntrospectionPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.Oauth2IntrospectionPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.Oauth2IntrospectionPluginRoute{
			ID: id2,
		}
	}
	var service *shared.Oauth2IntrospectionPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.Oauth2IntrospectionPluginService{
			ID: id3,
		}
	}
	out := shared.Oauth2IntrospectionPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Partials:     partials,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}

	return &out, diags
}

func (r *GatewayPluginOauth2IntrospectionResourceModel) ToOperationsCreateOauth2introspectionPluginRequest(ctx context.Context) (*operations.CreateOauth2introspectionPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	oauth2IntrospectionPlugin, oauth2IntrospectionPluginDiags := r.ToSharedOauth2IntrospectionPlugin(ctx)
	diags.Append(oauth2IntrospectionPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateOauth2introspectionPluginRequest{
		ControlPlaneID:            controlPlaneID,
		Oauth2IntrospectionPlugin: *oauth2IntrospectionPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginOauth2IntrospectionResourceModel) ToOperationsUpdateOauth2introspectionPluginRequest(ctx context.Context) (*operations.UpdateOauth2introspectionPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	oauth2IntrospectionPlugin, oauth2IntrospectionPluginDiags := r.ToSharedOauth2IntrospectionPlugin(ctx)
	diags.Append(oauth2IntrospectionPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateOauth2introspectionPluginRequest{
		PluginID:                  pluginID,
		ControlPlaneID:            controlPlaneID,
		Oauth2IntrospectionPlugin: *oauth2IntrospectionPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginOauth2IntrospectionResourceModel) ToOperationsGetOauth2introspectionPluginRequest(ctx context.Context) (*operations.GetOauth2introspectionPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetOauth2introspectionPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginOauth2IntrospectionResourceModel) ToOperationsDeleteOauth2introspectionPluginRequest(ctx context.Context) (*operations.DeleteOauth2introspectionPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.DeleteOauth2introspectionPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginOauth2IntrospectionResourceModel) RefreshFromSharedOauth2IntrospectionPlugin(ctx context.Context, resp *shared.Oauth2IntrospectionPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.Oauth2IntrospectionPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.AuthorizationValue = types.StringPointerValue(resp.Config.AuthorizationValue)
			if resp.Config.ConsumerBy != nil {
				r.Config.ConsumerBy = types.StringValue(string(*resp.Config.ConsumerBy))
			} else {
				r.Config.ConsumerBy = types.StringNull()
			}
			r.Config.CustomClaimsForward = make([]types.String, 0, len(resp.Config.CustomClaimsForward))
			for _, v := range resp.Config.CustomClaimsForward {
				r.Config.CustomClaimsForward = append(r.Config.CustomClaimsForward, types.StringValue(v))
			}
			if len(resp.Config.CustomIntrospectionHeaders) > 0 {
				r.Config.CustomIntrospectionHeaders = make(map[string]types.String, len(resp.Config.CustomIntrospectionHeaders))
				for key, value := range resp.Config.CustomIntrospectionHeaders {
					result, _ := json.Marshal(value)
					r.Config.CustomIntrospectionHeaders[key] = types.StringValue(string(result))
				}
			}
			r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
			r.Config.IntrospectRequest = types.BoolPointerValue(resp.Config.IntrospectRequest)
			r.Config.IntrospectionURL = types.StringPointerValue(resp.Config.IntrospectionURL)
			r.Config.Keepalive = types.Int64PointerValue(resp.Config.Keepalive)
			r.Config.RunOnPreflight = types.BoolPointerValue(resp.Config.RunOnPreflight)
			r.Config.Timeout = types.Int64PointerValue(resp.Config.Timeout)
			r.Config.TokenTypeHint = types.StringPointerValue(resp.Config.TokenTypeHint)
			r.Config.TTL = types.Float64PointerValue(resp.Config.TTL)
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
		if resp.Partials != nil {
			r.Partials = []tfTypes.Partials{}
			if len(r.Partials) > len(resp.Partials) {
				r.Partials = r.Partials[:len(resp.Partials)]
			}
			for partialsCount, partialsItem := range resp.Partials {
				var partials tfTypes.Partials
				partials.ID = types.StringPointerValue(partialsItem.ID)
				partials.Name = types.StringPointerValue(partialsItem.Name)
				partials.Path = types.StringPointerValue(partialsItem.Path)
				if partialsCount+1 > len(r.Partials) {
					r.Partials = append(r.Partials, partials)
				} else {
					r.Partials[partialsCount].ID = partials.ID
					r.Partials[partialsCount].Name = partials.Name
					r.Partials[partialsCount].Path = partials.Path
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
			r.Route = &tfTypes.Set{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.Set{}
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
