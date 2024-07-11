// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginPreFunctionDataSourceModel) RefreshFromSharedPreFunctionPlugin(resp *shared.PreFunctionPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreatePostFunctionPluginConfig{}
			r.Config.Access = []types.String{}
			for _, v := range resp.Config.Access {
				r.Config.Access = append(r.Config.Access, types.StringValue(v))
			}
			r.Config.BodyFilter = []types.String{}
			for _, v := range resp.Config.BodyFilter {
				r.Config.BodyFilter = append(r.Config.BodyFilter, types.StringValue(v))
			}
			r.Config.Certificate = []types.String{}
			for _, v := range resp.Config.Certificate {
				r.Config.Certificate = append(r.Config.Certificate, types.StringValue(v))
			}
			r.Config.HeaderFilter = []types.String{}
			for _, v := range resp.Config.HeaderFilter {
				r.Config.HeaderFilter = append(r.Config.HeaderFilter, types.StringValue(v))
			}
			r.Config.Log = []types.String{}
			for _, v := range resp.Config.Log {
				r.Config.Log = append(r.Config.Log, types.StringValue(v))
			}
			r.Config.Rewrite = []types.String{}
			for _, v := range resp.Config.Rewrite {
				r.Config.Rewrite = append(r.Config.Rewrite, types.StringValue(v))
			}
			r.Config.WsClientFrame = []types.String{}
			for _, v := range resp.Config.WsClientFrame {
				r.Config.WsClientFrame = append(r.Config.WsClientFrame, types.StringValue(v))
			}
			r.Config.WsClose = []types.String{}
			for _, v := range resp.Config.WsClose {
				r.Config.WsClose = append(r.Config.WsClose, types.StringValue(v))
			}
			r.Config.WsHandshake = []types.String{}
			for _, v := range resp.Config.WsHandshake {
				r.Config.WsHandshake = append(r.Config.WsHandshake, types.StringValue(v))
			}
			r.Config.WsUpstreamFrame = []types.String{}
			for _, v := range resp.Config.WsUpstreamFrame {
				r.Config.WsUpstreamFrame = append(r.Config.WsUpstreamFrame, types.StringValue(v))
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
