// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginOasValidationDataSourceModel) RefreshFromSharedOasValidationPlugin(resp *shared.OasValidationPlugin) {
	if resp != nil {
		r.Config.AllowedHeaderParameters = types.StringPointerValue(resp.Config.AllowedHeaderParameters)
		r.Config.APISpec = types.StringPointerValue(resp.Config.APISpec)
		r.Config.APISpecEncoded = types.BoolPointerValue(resp.Config.APISpecEncoded)
		r.Config.CustomBasePath = types.StringPointerValue(resp.Config.CustomBasePath)
		r.Config.HeaderParameterCheck = types.BoolPointerValue(resp.Config.HeaderParameterCheck)
		r.Config.IncludeBasePath = types.BoolPointerValue(resp.Config.IncludeBasePath)
		r.Config.NotifyOnlyRequestValidationFailure = types.BoolPointerValue(resp.Config.NotifyOnlyRequestValidationFailure)
		r.Config.NotifyOnlyResponseBodyValidationFailure = types.BoolPointerValue(resp.Config.NotifyOnlyResponseBodyValidationFailure)
		r.Config.QueryParameterCheck = types.BoolPointerValue(resp.Config.QueryParameterCheck)
		r.Config.ValidateRequestBody = types.BoolPointerValue(resp.Config.ValidateRequestBody)
		r.Config.ValidateRequestHeaderParams = types.BoolPointerValue(resp.Config.ValidateRequestHeaderParams)
		r.Config.ValidateRequestQueryParams = types.BoolPointerValue(resp.Config.ValidateRequestQueryParams)
		r.Config.ValidateRequestURIParams = types.BoolPointerValue(resp.Config.ValidateRequestURIParams)
		r.Config.ValidateResponseBody = types.BoolPointerValue(resp.Config.ValidateResponseBody)
		r.Config.VerboseResponse = types.BoolPointerValue(resp.Config.VerboseResponse)
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
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}