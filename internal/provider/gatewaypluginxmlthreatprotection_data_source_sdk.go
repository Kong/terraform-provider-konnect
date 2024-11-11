// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginXMLThreatProtectionDataSourceModel) RefreshFromSharedXMLThreatProtectionPlugin(resp *shared.XMLThreatProtectionPlugin) {
	if resp != nil {
		r.Config.AllowDtd = types.BoolPointerValue(resp.Config.AllowDtd)
		r.Config.AllowedContentTypes = []types.String{}
		for _, v := range resp.Config.AllowedContentTypes {
			r.Config.AllowedContentTypes = append(r.Config.AllowedContentTypes, types.StringValue(v))
		}
		r.Config.Attribute = types.Int64PointerValue(resp.Config.Attribute)
		if resp.Config.BlaMaxAmplification != nil {
			r.Config.BlaMaxAmplification = types.NumberValue(big.NewFloat(float64(*resp.Config.BlaMaxAmplification)))
		} else {
			r.Config.BlaMaxAmplification = types.NumberNull()
		}
		r.Config.BlaThreshold = types.Int64PointerValue(resp.Config.BlaThreshold)
		r.Config.Buffer = types.Int64PointerValue(resp.Config.Buffer)
		r.Config.CheckedContentTypes = []types.String{}
		for _, v := range resp.Config.CheckedContentTypes {
			r.Config.CheckedContentTypes = append(r.Config.CheckedContentTypes, types.StringValue(v))
		}
		r.Config.Comment = types.Int64PointerValue(resp.Config.Comment)
		r.Config.Document = types.Int64PointerValue(resp.Config.Document)
		r.Config.Entity = types.Int64PointerValue(resp.Config.Entity)
		r.Config.Entityname = types.Int64PointerValue(resp.Config.Entityname)
		r.Config.Entityproperty = types.Int64PointerValue(resp.Config.Entityproperty)
		r.Config.Localname = types.Int64PointerValue(resp.Config.Localname)
		r.Config.MaxAttributes = types.Int64PointerValue(resp.Config.MaxAttributes)
		r.Config.MaxChildren = types.Int64PointerValue(resp.Config.MaxChildren)
		r.Config.MaxDepth = types.Int64PointerValue(resp.Config.MaxDepth)
		r.Config.MaxNamespaces = types.Int64PointerValue(resp.Config.MaxNamespaces)
		r.Config.NamespaceAware = types.BoolPointerValue(resp.Config.NamespaceAware)
		r.Config.Namespaceuri = types.Int64PointerValue(resp.Config.Namespaceuri)
		r.Config.Pidata = types.Int64PointerValue(resp.Config.Pidata)
		r.Config.Pitarget = types.Int64PointerValue(resp.Config.Pitarget)
		r.Config.Prefix = types.Int64PointerValue(resp.Config.Prefix)
		r.Config.Text = types.Int64PointerValue(resp.Config.Text)
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
