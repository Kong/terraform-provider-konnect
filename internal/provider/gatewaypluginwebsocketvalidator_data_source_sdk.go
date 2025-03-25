// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginWebsocketValidatorDataSourceModel) RefreshFromSharedWebsocketValidatorPlugin(resp *shared.WebsocketValidatorPlugin) {
	if resp != nil {
		if resp.Config.Client == nil {
			r.Config.Client = nil
		} else {
			r.Config.Client = &tfTypes.WebsocketValidatorPluginClient{}
			if resp.Config.Client.Binary == nil {
				r.Config.Client.Binary = nil
			} else {
				r.Config.Client.Binary = &tfTypes.Binary{}
				r.Config.Client.Binary.Schema = types.StringValue(resp.Config.Client.Binary.Schema)
				r.Config.Client.Binary.Type = types.StringValue(string(resp.Config.Client.Binary.Type))
			}
			if resp.Config.Client.Text == nil {
				r.Config.Client.Text = nil
			} else {
				r.Config.Client.Text = &tfTypes.Binary{}
				r.Config.Client.Text.Schema = types.StringValue(resp.Config.Client.Text.Schema)
				r.Config.Client.Text.Type = types.StringValue(string(resp.Config.Client.Text.Type))
			}
		}
		if resp.Config.Upstream == nil {
			r.Config.Upstream = nil
		} else {
			r.Config.Upstream = &tfTypes.WebsocketValidatorPluginClient{}
			if resp.Config.Upstream.Binary == nil {
				r.Config.Upstream.Binary = nil
			} else {
				r.Config.Upstream.Binary = &tfTypes.Binary{}
				r.Config.Upstream.Binary.Schema = types.StringValue(resp.Config.Upstream.Binary.Schema)
				r.Config.Upstream.Binary.Type = types.StringValue(string(resp.Config.Upstream.Binary.Type))
			}
			if resp.Config.Upstream.Text == nil {
				r.Config.Upstream.Text = nil
			} else {
				r.Config.Upstream.Text = &tfTypes.Binary{}
				r.Config.Upstream.Text.Schema = types.StringValue(resp.Config.Upstream.Text.Schema)
				r.Config.Upstream.Text.Type = types.StringValue(string(resp.Config.Upstream.Text.Type))
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
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
