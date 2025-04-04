// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginRequestTransformerDataSourceModel) RefreshFromSharedRequestTransformerPlugin(resp *shared.RequestTransformerPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.RequestTransformerPluginConfig{}
			if resp.Config.Add == nil {
				r.Config.Add = nil
			} else {
				r.Config.Add = &tfTypes.Add{}
				r.Config.Add.Body = make([]types.String, 0, len(resp.Config.Add.Body))
				for _, v := range resp.Config.Add.Body {
					r.Config.Add.Body = append(r.Config.Add.Body, types.StringValue(v))
				}
				r.Config.Add.Headers = make([]types.String, 0, len(resp.Config.Add.Headers))
				for _, v := range resp.Config.Add.Headers {
					r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
				}
				r.Config.Add.Querystring = make([]types.String, 0, len(resp.Config.Add.Querystring))
				for _, v := range resp.Config.Add.Querystring {
					r.Config.Add.Querystring = append(r.Config.Add.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Append == nil {
				r.Config.Append = nil
			} else {
				r.Config.Append = &tfTypes.Add{}
				r.Config.Append.Body = make([]types.String, 0, len(resp.Config.Append.Body))
				for _, v := range resp.Config.Append.Body {
					r.Config.Append.Body = append(r.Config.Append.Body, types.StringValue(v))
				}
				r.Config.Append.Headers = make([]types.String, 0, len(resp.Config.Append.Headers))
				for _, v := range resp.Config.Append.Headers {
					r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
				}
				r.Config.Append.Querystring = make([]types.String, 0, len(resp.Config.Append.Querystring))
				for _, v := range resp.Config.Append.Querystring {
					r.Config.Append.Querystring = append(r.Config.Append.Querystring, types.StringValue(v))
				}
			}
			r.Config.HTTPMethod = types.StringPointerValue(resp.Config.HTTPMethod)
			if resp.Config.Remove == nil {
				r.Config.Remove = nil
			} else {
				r.Config.Remove = &tfTypes.Add{}
				r.Config.Remove.Body = make([]types.String, 0, len(resp.Config.Remove.Body))
				for _, v := range resp.Config.Remove.Body {
					r.Config.Remove.Body = append(r.Config.Remove.Body, types.StringValue(v))
				}
				r.Config.Remove.Headers = make([]types.String, 0, len(resp.Config.Remove.Headers))
				for _, v := range resp.Config.Remove.Headers {
					r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
				}
				r.Config.Remove.Querystring = make([]types.String, 0, len(resp.Config.Remove.Querystring))
				for _, v := range resp.Config.Remove.Querystring {
					r.Config.Remove.Querystring = append(r.Config.Remove.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Rename == nil {
				r.Config.Rename = nil
			} else {
				r.Config.Rename = &tfTypes.Add{}
				r.Config.Rename.Body = make([]types.String, 0, len(resp.Config.Rename.Body))
				for _, v := range resp.Config.Rename.Body {
					r.Config.Rename.Body = append(r.Config.Rename.Body, types.StringValue(v))
				}
				r.Config.Rename.Headers = make([]types.String, 0, len(resp.Config.Rename.Headers))
				for _, v := range resp.Config.Rename.Headers {
					r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
				}
				r.Config.Rename.Querystring = make([]types.String, 0, len(resp.Config.Rename.Querystring))
				for _, v := range resp.Config.Rename.Querystring {
					r.Config.Rename.Querystring = append(r.Config.Rename.Querystring, types.StringValue(v))
				}
			}
			if resp.Config.Replace == nil {
				r.Config.Replace = nil
			} else {
				r.Config.Replace = &tfTypes.Replace{}
				r.Config.Replace.Body = make([]types.String, 0, len(resp.Config.Replace.Body))
				for _, v := range resp.Config.Replace.Body {
					r.Config.Replace.Body = append(r.Config.Replace.Body, types.StringValue(v))
				}
				r.Config.Replace.Headers = make([]types.String, 0, len(resp.Config.Replace.Headers))
				for _, v := range resp.Config.Replace.Headers {
					r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
				}
				r.Config.Replace.Querystring = make([]types.String, 0, len(resp.Config.Replace.Querystring))
				for _, v := range resp.Config.Replace.Querystring {
					r.Config.Replace.Querystring = append(r.Config.Replace.Querystring, types.StringValue(v))
				}
				r.Config.Replace.URI = types.StringPointerValue(resp.Config.Replace.URI)
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLWithoutParentsConsumer{}
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
