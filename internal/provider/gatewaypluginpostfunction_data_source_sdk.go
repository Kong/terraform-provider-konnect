// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginPostFunctionDataSourceModel) RefreshFromSharedPostFunctionPlugin(resp *shared.PostFunctionPlugin) {
	if resp != nil {
		r.Config.Access = make([]types.String, 0, len(resp.Config.Access))
		for _, v := range resp.Config.Access {
			r.Config.Access = append(r.Config.Access, types.StringValue(v))
		}
		r.Config.BodyFilter = make([]types.String, 0, len(resp.Config.BodyFilter))
		for _, v := range resp.Config.BodyFilter {
			r.Config.BodyFilter = append(r.Config.BodyFilter, types.StringValue(v))
		}
		r.Config.Certificate = make([]types.String, 0, len(resp.Config.Certificate))
		for _, v := range resp.Config.Certificate {
			r.Config.Certificate = append(r.Config.Certificate, types.StringValue(v))
		}
		r.Config.HeaderFilter = make([]types.String, 0, len(resp.Config.HeaderFilter))
		for _, v := range resp.Config.HeaderFilter {
			r.Config.HeaderFilter = append(r.Config.HeaderFilter, types.StringValue(v))
		}
		r.Config.Log = make([]types.String, 0, len(resp.Config.Log))
		for _, v := range resp.Config.Log {
			r.Config.Log = append(r.Config.Log, types.StringValue(v))
		}
		r.Config.Rewrite = make([]types.String, 0, len(resp.Config.Rewrite))
		for _, v := range resp.Config.Rewrite {
			r.Config.Rewrite = append(r.Config.Rewrite, types.StringValue(v))
		}
		r.Config.WsClientFrame = make([]types.String, 0, len(resp.Config.WsClientFrame))
		for _, v := range resp.Config.WsClientFrame {
			r.Config.WsClientFrame = append(r.Config.WsClientFrame, types.StringValue(v))
		}
		r.Config.WsClose = make([]types.String, 0, len(resp.Config.WsClose))
		for _, v := range resp.Config.WsClose {
			r.Config.WsClose = append(r.Config.WsClose, types.StringValue(v))
		}
		r.Config.WsHandshake = make([]types.String, 0, len(resp.Config.WsHandshake))
		for _, v := range resp.Config.WsHandshake {
			r.Config.WsHandshake = append(r.Config.WsHandshake, types.StringValue(v))
		}
		r.Config.WsUpstreamFrame = make([]types.String, 0, len(resp.Config.WsUpstreamFrame))
		for _, v := range resp.Config.WsUpstreamFrame {
			r.Config.WsUpstreamFrame = append(r.Config.WsUpstreamFrame, types.StringValue(v))
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering != nil {
			r.Ordering = make(map[string]types.String, len(resp.Ordering))
			for key, value := range resp.Ordering {
				r.Ordering[key] = types.StringValue(value)
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
