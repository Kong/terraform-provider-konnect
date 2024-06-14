// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginResponseTransformerAdvancedDataSourceModel) RefreshFromSharedResponseTransformerAdvancedPlugin(resp *shared.ResponseTransformerAdvancedPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateResponseTransformerAdvancedPluginConfig{}
			if resp.Config.Add == nil {
				r.Config.Add = nil
			} else {
				r.Config.Add = &tfTypes.CreateResponseTransformerAdvancedPluginAdd{}
				r.Config.Add.Headers = []types.String{}
				for _, v := range resp.Config.Add.Headers {
					r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
				}
				r.Config.Add.IfStatus = []types.String{}
				for _, v := range resp.Config.Add.IfStatus {
					r.Config.Add.IfStatus = append(r.Config.Add.IfStatus, types.StringValue(v))
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
			if resp.Config.Allow == nil {
				r.Config.Allow = nil
			} else {
				r.Config.Allow = &tfTypes.CreateResponseTransformerAdvancedPluginAllow{}
				r.Config.Allow.JSON = []types.String{}
				for _, v := range resp.Config.Allow.JSON {
					r.Config.Allow.JSON = append(r.Config.Allow.JSON, types.StringValue(v))
				}
			}
			if resp.Config.Append == nil {
				r.Config.Append = nil
			} else {
				r.Config.Append = &tfTypes.CreateResponseTransformerAdvancedPluginAdd{}
				r.Config.Append.Headers = []types.String{}
				for _, v := range resp.Config.Append.Headers {
					r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
				}
				r.Config.Append.IfStatus = []types.String{}
				for _, v := range resp.Config.Append.IfStatus {
					r.Config.Append.IfStatus = append(r.Config.Append.IfStatus, types.StringValue(v))
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
			r.Config.DotsInKeys = types.BoolPointerValue(resp.Config.DotsInKeys)
			if resp.Config.Remove == nil {
				r.Config.Remove = nil
			} else {
				r.Config.Remove = &tfTypes.CreateResponseTransformerAdvancedPluginRemove{}
				r.Config.Remove.Headers = []types.String{}
				for _, v := range resp.Config.Remove.Headers {
					r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
				}
				r.Config.Remove.IfStatus = []types.String{}
				for _, v := range resp.Config.Remove.IfStatus {
					r.Config.Remove.IfStatus = append(r.Config.Remove.IfStatus, types.StringValue(v))
				}
				r.Config.Remove.JSON = []types.String{}
				for _, v := range resp.Config.Remove.JSON {
					r.Config.Remove.JSON = append(r.Config.Remove.JSON, types.StringValue(v))
				}
			}
			if resp.Config.Rename == nil {
				r.Config.Rename = nil
			} else {
				r.Config.Rename = &tfTypes.CreateResponseTransformerAdvancedPluginRename{}
				r.Config.Rename.Headers = []types.String{}
				for _, v := range resp.Config.Rename.Headers {
					r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
				}
				r.Config.Rename.IfStatus = []types.String{}
				for _, v := range resp.Config.Rename.IfStatus {
					r.Config.Rename.IfStatus = append(r.Config.Rename.IfStatus, types.StringValue(v))
				}
			}
			if resp.Config.Replace == nil {
				r.Config.Replace = nil
			} else {
				r.Config.Replace = &tfTypes.CreateResponseTransformerAdvancedPluginReplace{}
				r.Config.Replace.Body = types.StringPointerValue(resp.Config.Replace.Body)
				r.Config.Replace.Headers = []types.String{}
				for _, v := range resp.Config.Replace.Headers {
					r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
				}
				r.Config.Replace.IfStatus = []types.String{}
				for _, v := range resp.Config.Replace.IfStatus {
					r.Config.Replace.IfStatus = append(r.Config.Replace.IfStatus, types.StringValue(v))
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
			if resp.Config.Transform == nil {
				r.Config.Transform = nil
			} else {
				r.Config.Transform = &tfTypes.CreateResponseTransformerAdvancedPluginTransform{}
				r.Config.Transform.Functions = []types.String{}
				for _, v := range resp.Config.Transform.Functions {
					r.Config.Transform.Functions = append(r.Config.Transform.Functions, types.StringValue(v))
				}
				r.Config.Transform.IfStatus = []types.String{}
				for _, v := range resp.Config.Transform.IfStatus {
					r.Config.Transform.IfStatus = append(r.Config.Transform.IfStatus, types.StringValue(v))
				}
				r.Config.Transform.JSON = []types.String{}
				for _, v := range resp.Config.Transform.JSON {
					r.Config.Transform.JSON = append(r.Config.Transform.JSON, types.StringValue(v))
				}
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
