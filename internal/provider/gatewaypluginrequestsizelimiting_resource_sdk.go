// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginRequestSizeLimitingResourceModel) ToSharedRequestSizeLimitingPlugin(ctx context.Context) (*shared.RequestSizeLimitingPlugin, diag.Diagnostics) {
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
	var ordering *shared.RequestSizeLimitingPluginOrdering
	if r.Ordering != nil {
		var after *shared.RequestSizeLimitingPluginAfter
		if r.Ordering.After != nil {
			access := make([]string, 0, len(r.Ordering.After.Access))
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.RequestSizeLimitingPluginAfter{
				Access: access,
			}
		}
		var before *shared.RequestSizeLimitingPluginBefore
		if r.Ordering.Before != nil {
			access1 := make([]string, 0, len(r.Ordering.Before.Access))
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.RequestSizeLimitingPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.RequestSizeLimitingPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.RequestSizeLimitingPluginPartials
	if r.Partials != nil {
		partials = make([]shared.RequestSizeLimitingPluginPartials, 0, len(r.Partials))
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
			partials = append(partials, shared.RequestSizeLimitingPluginPartials{
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
	var config *shared.RequestSizeLimitingPluginConfig
	if r.Config != nil {
		allowedPayloadSize := new(int64)
		if !r.Config.AllowedPayloadSize.IsUnknown() && !r.Config.AllowedPayloadSize.IsNull() {
			*allowedPayloadSize = r.Config.AllowedPayloadSize.ValueInt64()
		} else {
			allowedPayloadSize = nil
		}
		requireContentLength := new(bool)
		if !r.Config.RequireContentLength.IsUnknown() && !r.Config.RequireContentLength.IsNull() {
			*requireContentLength = r.Config.RequireContentLength.ValueBool()
		} else {
			requireContentLength = nil
		}
		sizeUnit := new(shared.SizeUnit)
		if !r.Config.SizeUnit.IsUnknown() && !r.Config.SizeUnit.IsNull() {
			*sizeUnit = shared.SizeUnit(r.Config.SizeUnit.ValueString())
		} else {
			sizeUnit = nil
		}
		config = &shared.RequestSizeLimitingPluginConfig{
			AllowedPayloadSize:   allowedPayloadSize,
			RequireContentLength: requireContentLength,
			SizeUnit:             sizeUnit,
		}
	}
	var consumer *shared.RequestSizeLimitingPluginConsumer
	if r.Consumer != nil {
		id2 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id2 = r.Consumer.ID.ValueString()
		} else {
			id2 = nil
		}
		consumer = &shared.RequestSizeLimitingPluginConsumer{
			ID: id2,
		}
	}
	protocols := make([]shared.RequestSizeLimitingPluginProtocols, 0, len(r.Protocols))
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.RequestSizeLimitingPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.RequestSizeLimitingPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.RequestSizeLimitingPluginRoute{
			ID: id3,
		}
	}
	var service *shared.RequestSizeLimitingPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.RequestSizeLimitingPluginService{
			ID: id4,
		}
	}
	out := shared.RequestSizeLimitingPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Partials:     partials,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}

	return &out, diags
}

func (r *GatewayPluginRequestSizeLimitingResourceModel) ToOperationsCreateRequestsizelimitingPluginRequest(ctx context.Context) (*operations.CreateRequestsizelimitingPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	requestSizeLimitingPlugin, requestSizeLimitingPluginDiags := r.ToSharedRequestSizeLimitingPlugin(ctx)
	diags.Append(requestSizeLimitingPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateRequestsizelimitingPluginRequest{
		ControlPlaneID:            controlPlaneID,
		RequestSizeLimitingPlugin: *requestSizeLimitingPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginRequestSizeLimitingResourceModel) ToOperationsUpdateRequestsizelimitingPluginRequest(ctx context.Context) (*operations.UpdateRequestsizelimitingPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	requestSizeLimitingPlugin, requestSizeLimitingPluginDiags := r.ToSharedRequestSizeLimitingPlugin(ctx)
	diags.Append(requestSizeLimitingPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateRequestsizelimitingPluginRequest{
		PluginID:                  pluginID,
		ControlPlaneID:            controlPlaneID,
		RequestSizeLimitingPlugin: *requestSizeLimitingPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginRequestSizeLimitingResourceModel) ToOperationsGetRequestsizelimitingPluginRequest(ctx context.Context) (*operations.GetRequestsizelimitingPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetRequestsizelimitingPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginRequestSizeLimitingResourceModel) ToOperationsDeleteRequestsizelimitingPluginRequest(ctx context.Context) (*operations.DeleteRequestsizelimitingPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.DeleteRequestsizelimitingPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginRequestSizeLimitingResourceModel) RefreshFromSharedRequestSizeLimitingPlugin(ctx context.Context, resp *shared.RequestSizeLimitingPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.RequestSizeLimitingPluginConfig{}
			r.Config.AllowedPayloadSize = types.Int64PointerValue(resp.Config.AllowedPayloadSize)
			r.Config.RequireContentLength = types.BoolPointerValue(resp.Config.RequireContentLength)
			if resp.Config.SizeUnit != nil {
				r.Config.SizeUnit = types.StringValue(string(*resp.Config.SizeUnit))
			} else {
				r.Config.SizeUnit = types.StringNull()
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.Set{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
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
