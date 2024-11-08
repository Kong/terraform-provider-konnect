// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginResponseTransformerDataSourceModel) RefreshFromSharedResponseTransformerPlugin(resp *shared.ResponseTransformerPlugin) {
	if resp != nil {
		if resp.Config.Add == nil {
			r.Config.Add = nil
		} else {
			r.Config.Add = &tfTypes.ResponseTransformerPluginAdd{}
			r.Config.Add.Headers = []types.String{}
			for _, v := range resp.Config.Add.Headers {
				r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
			}
			r.Config.Add.JSON = []types.String{}
			for _, v := range resp.Config.Add.JSON {
				r.Config.Add.JSON = append(r.Config.Add.JSON, types.StringValue(v))
			}
			r.Config.Add.JSONTypes = []types.String{}
			for _, v := range resp.Config.Add.JSONTypes {
				r.Config.Add.JSONTypes = append(r.Config.Add.JSONTypes, types.StringValue(string(v)))
			}
		}
		if resp.Config.Append == nil {
			r.Config.Append = nil
		} else {
			r.Config.Append = &tfTypes.ResponseTransformerPluginAdd{}
			r.Config.Append.Headers = []types.String{}
			for _, v := range resp.Config.Append.Headers {
				r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
			}
			r.Config.Append.JSON = []types.String{}
			for _, v := range resp.Config.Append.JSON {
				r.Config.Append.JSON = append(r.Config.Append.JSON, types.StringValue(v))
			}
			r.Config.Append.JSONTypes = []types.String{}
			for _, v := range resp.Config.Append.JSONTypes {
				r.Config.Append.JSONTypes = append(r.Config.Append.JSONTypes, types.StringValue(string(v)))
			}
		}
		if resp.Config.Remove == nil {
			r.Config.Remove = nil
		} else {
			r.Config.Remove = &tfTypes.ResponseTransformerPluginRemove{}
			r.Config.Remove.Headers = []types.String{}
			for _, v := range resp.Config.Remove.Headers {
				r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
			}
			r.Config.Remove.JSON = []types.String{}
			for _, v := range resp.Config.Remove.JSON {
				r.Config.Remove.JSON = append(r.Config.Remove.JSON, types.StringValue(v))
			}
		}
		if resp.Config.Rename == nil {
			r.Config.Rename = nil
		} else {
			r.Config.Rename = &tfTypes.ResponseTransformerPluginRemove{}
			r.Config.Rename.Headers = []types.String{}
			for _, v := range resp.Config.Rename.Headers {
				r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
			}
			r.Config.Rename.JSON = []types.String{}
			for _, v := range resp.Config.Rename.JSON {
				r.Config.Rename.JSON = append(r.Config.Rename.JSON, types.StringValue(v))
			}
		}
		if resp.Config.Replace == nil {
			r.Config.Replace = nil
		} else {
			r.Config.Replace = &tfTypes.ResponseTransformerPluginAdd{}
			r.Config.Replace.Headers = []types.String{}
			for _, v := range resp.Config.Replace.Headers {
				r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
			}
			r.Config.Replace.JSON = []types.String{}
			for _, v := range resp.Config.Replace.JSON {
				r.Config.Replace.JSON = append(r.Config.Replace.JSON, types.StringValue(v))
			}
			r.Config.Replace.JSONTypes = []types.String{}
			for _, v := range resp.Config.Replace.JSONTypes {
				r.Config.Replace.JSONTypes = append(r.Config.Replace.JSONTypes, types.StringValue(string(v)))
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
