// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginLogglyDataSourceModel) RefreshFromSharedLogglyPlugin(resp *shared.LogglyPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.LogglyPluginConfig{}
			if resp.Config.ClientErrorsSeverity != nil {
				r.Config.ClientErrorsSeverity = types.StringValue(string(*resp.Config.ClientErrorsSeverity))
			} else {
				r.Config.ClientErrorsSeverity = types.StringNull()
			}
			if len(resp.Config.CustomFieldsByLua) > 0 {
				r.Config.CustomFieldsByLua = make(map[string]types.String, len(resp.Config.CustomFieldsByLua))
				for key, value := range resp.Config.CustomFieldsByLua {
					result, _ := json.Marshal(value)
					r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
				}
			}
			r.Config.Host = types.StringPointerValue(resp.Config.Host)
			r.Config.Key = types.StringPointerValue(resp.Config.Key)
			if resp.Config.LogLevel != nil {
				r.Config.LogLevel = types.StringValue(string(*resp.Config.LogLevel))
			} else {
				r.Config.LogLevel = types.StringNull()
			}
			r.Config.Port = types.Int64PointerValue(resp.Config.Port)
			if resp.Config.ServerErrorsSeverity != nil {
				r.Config.ServerErrorsSeverity = types.StringValue(string(*resp.Config.ServerErrorsSeverity))
			} else {
				r.Config.ServerErrorsSeverity = types.StringNull()
			}
			if resp.Config.SuccessfulSeverity != nil {
				r.Config.SuccessfulSeverity = types.StringValue(string(*resp.Config.SuccessfulSeverity))
			} else {
				r.Config.SuccessfulSeverity = types.StringNull()
			}
			r.Config.Tags = make([]types.String, 0, len(resp.Config.Tags))
			for _, v := range resp.Config.Tags {
				r.Config.Tags = append(r.Config.Tags, types.StringValue(v))
			}
			if resp.Config.Timeout != nil {
				r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
			} else {
				r.Config.Timeout = types.NumberNull()
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
}
