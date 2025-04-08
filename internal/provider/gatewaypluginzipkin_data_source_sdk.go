// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginZipkinDataSourceModel) RefreshFromSharedZipkinPlugin(ctx context.Context, resp *shared.ZipkinPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.ZipkinPluginConfig{}
			r.Config.ConnectTimeout = types.Int64PointerValue(resp.Config.ConnectTimeout)
			if resp.Config.DefaultHeaderType != nil {
				r.Config.DefaultHeaderType = types.StringValue(string(*resp.Config.DefaultHeaderType))
			} else {
				r.Config.DefaultHeaderType = types.StringNull()
			}
			r.Config.DefaultServiceName = types.StringPointerValue(resp.Config.DefaultServiceName)
			if resp.Config.HeaderType != nil {
				r.Config.HeaderType = types.StringValue(string(*resp.Config.HeaderType))
			} else {
				r.Config.HeaderType = types.StringNull()
			}
			r.Config.HTTPEndpoint = types.StringPointerValue(resp.Config.HTTPEndpoint)
			r.Config.HTTPResponseHeaderForTraceid = types.StringPointerValue(resp.Config.HTTPResponseHeaderForTraceid)
			if resp.Config.HTTPSpanName != nil {
				r.Config.HTTPSpanName = types.StringValue(string(*resp.Config.HTTPSpanName))
			} else {
				r.Config.HTTPSpanName = types.StringNull()
			}
			r.Config.IncludeCredential = types.BoolPointerValue(resp.Config.IncludeCredential)
			r.Config.LocalServiceName = types.StringPointerValue(resp.Config.LocalServiceName)
			if resp.Config.PhaseDurationFlavor != nil {
				r.Config.PhaseDurationFlavor = types.StringValue(string(*resp.Config.PhaseDurationFlavor))
			} else {
				r.Config.PhaseDurationFlavor = types.StringNull()
			}
			if resp.Config.Propagation == nil {
				r.Config.Propagation = nil
			} else {
				r.Config.Propagation = &tfTypes.Propagation{}
				r.Config.Propagation.Clear = make([]types.String, 0, len(resp.Config.Propagation.Clear))
				for _, v := range resp.Config.Propagation.Clear {
					r.Config.Propagation.Clear = append(r.Config.Propagation.Clear, types.StringValue(v))
				}
				r.Config.Propagation.DefaultFormat = types.StringValue(string(resp.Config.Propagation.DefaultFormat))
				r.Config.Propagation.Extract = make([]types.String, 0, len(resp.Config.Propagation.Extract))
				for _, v := range resp.Config.Propagation.Extract {
					r.Config.Propagation.Extract = append(r.Config.Propagation.Extract, types.StringValue(string(v)))
				}
				r.Config.Propagation.Inject = make([]types.String, 0, len(resp.Config.Propagation.Inject))
				for _, v := range resp.Config.Propagation.Inject {
					r.Config.Propagation.Inject = append(r.Config.Propagation.Inject, types.StringValue(string(v)))
				}
			}
			if resp.Config.Queue == nil {
				r.Config.Queue = nil
			} else {
				r.Config.Queue = &tfTypes.Queue{}
				if resp.Config.Queue.ConcurrencyLimit != nil {
					r.Config.Queue.ConcurrencyLimit = types.Int64Value(int64(*resp.Config.Queue.ConcurrencyLimit))
				} else {
					r.Config.Queue.ConcurrencyLimit = types.Int64Null()
				}
				r.Config.Queue.InitialRetryDelay = types.Float64PointerValue(resp.Config.Queue.InitialRetryDelay)
				r.Config.Queue.MaxBatchSize = types.Int64PointerValue(resp.Config.Queue.MaxBatchSize)
				r.Config.Queue.MaxBytes = types.Int64PointerValue(resp.Config.Queue.MaxBytes)
				r.Config.Queue.MaxCoalescingDelay = types.Float64PointerValue(resp.Config.Queue.MaxCoalescingDelay)
				r.Config.Queue.MaxEntries = types.Int64PointerValue(resp.Config.Queue.MaxEntries)
				r.Config.Queue.MaxRetryDelay = types.Float64PointerValue(resp.Config.Queue.MaxRetryDelay)
				r.Config.Queue.MaxRetryTime = types.Float64PointerValue(resp.Config.Queue.MaxRetryTime)
			}
			r.Config.ReadTimeout = types.Int64PointerValue(resp.Config.ReadTimeout)
			r.Config.SampleRatio = types.Float64PointerValue(resp.Config.SampleRatio)
			r.Config.SendTimeout = types.Int64PointerValue(resp.Config.SendTimeout)
			r.Config.StaticTags = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
			if len(r.Config.StaticTags) > len(resp.Config.StaticTags) {
				r.Config.StaticTags = r.Config.StaticTags[:len(resp.Config.StaticTags)]
			}
			for staticTagsCount, staticTagsItem := range resp.Config.StaticTags {
				var staticTags tfTypes.ConfigurationDataPlaneGroupEnvironmentField
				staticTags.Name = types.StringValue(staticTagsItem.Name)
				staticTags.Value = types.StringValue(staticTagsItem.Value)
				if staticTagsCount+1 > len(r.Config.StaticTags) {
					r.Config.StaticTags = append(r.Config.StaticTags, staticTags)
				} else {
					r.Config.StaticTags[staticTagsCount].Name = staticTags.Name
					r.Config.StaticTags[staticTagsCount].Value = staticTags.Value
				}
			}
			r.Config.TagsHeader = types.StringPointerValue(resp.Config.TagsHeader)
			if resp.Config.TraceidByteCount != nil {
				r.Config.TraceidByteCount = types.Int64Value(int64(*resp.Config.TraceidByteCount))
			} else {
				r.Config.TraceidByteCount = types.Int64Null()
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

	return diags
}
