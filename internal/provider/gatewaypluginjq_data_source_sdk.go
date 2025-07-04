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

func (r *GatewayPluginJqDataSourceModel) ToOperationsGetJqPluginRequest(ctx context.Context) (*operations.GetJqPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetJqPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginJqDataSourceModel) RefreshFromSharedJqPlugin(ctx context.Context, resp *shared.JqPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.JqPluginConfig{}
			r.Config.RequestIfMediaType = make([]types.String, 0, len(resp.Config.RequestIfMediaType))
			for _, v := range resp.Config.RequestIfMediaType {
				r.Config.RequestIfMediaType = append(r.Config.RequestIfMediaType, types.StringValue(v))
			}
			r.Config.RequestJqProgram = types.StringPointerValue(resp.Config.RequestJqProgram)
			if resp.Config.RequestJqProgramOptions == nil {
				r.Config.RequestJqProgramOptions = nil
			} else {
				r.Config.RequestJqProgramOptions = &tfTypes.RequestJqProgramOptions{}
				r.Config.RequestJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.ASCIIOutput)
				r.Config.RequestJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.CompactOutput)
				r.Config.RequestJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.JoinOutput)
				r.Config.RequestJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.RawOutput)
				r.Config.RequestJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.SortKeys)
			}
			r.Config.ResponseIfMediaType = make([]types.String, 0, len(resp.Config.ResponseIfMediaType))
			for _, v := range resp.Config.ResponseIfMediaType {
				r.Config.ResponseIfMediaType = append(r.Config.ResponseIfMediaType, types.StringValue(v))
			}
			r.Config.ResponseIfStatusCode = make([]types.Int64, 0, len(resp.Config.ResponseIfStatusCode))
			for _, v := range resp.Config.ResponseIfStatusCode {
				r.Config.ResponseIfStatusCode = append(r.Config.ResponseIfStatusCode, types.Int64Value(v))
			}
			r.Config.ResponseJqProgram = types.StringPointerValue(resp.Config.ResponseJqProgram)
			if resp.Config.ResponseJqProgramOptions == nil {
				r.Config.ResponseJqProgramOptions = nil
			} else {
				r.Config.ResponseJqProgramOptions = &tfTypes.RequestJqProgramOptions{}
				r.Config.ResponseJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.ASCIIOutput)
				r.Config.ResponseJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.CompactOutput)
				r.Config.ResponseJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.JoinOutput)
				r.Config.ResponseJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.RawOutput)
				r.Config.ResponseJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.SortKeys)
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
