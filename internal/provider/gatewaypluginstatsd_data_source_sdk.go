// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginStatsdDataSourceModel) RefreshFromSharedStatsdPlugin(ctx context.Context, resp *shared.StatsdPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.StatsdPluginConfig{}
			r.Config.AllowStatusCodes = make([]types.String, 0, len(resp.Config.AllowStatusCodes))
			for _, v := range resp.Config.AllowStatusCodes {
				r.Config.AllowStatusCodes = append(r.Config.AllowStatusCodes, types.StringValue(v))
			}
			if resp.Config.ConsumerIdentifierDefault != nil {
				r.Config.ConsumerIdentifierDefault = types.StringValue(string(*resp.Config.ConsumerIdentifierDefault))
			} else {
				r.Config.ConsumerIdentifierDefault = types.StringNull()
			}
			r.Config.FlushTimeout = types.Float64PointerValue(resp.Config.FlushTimeout)
			r.Config.Host = types.StringPointerValue(resp.Config.Host)
			r.Config.HostnameInPrefix = types.BoolPointerValue(resp.Config.HostnameInPrefix)
			r.Config.Metrics = []tfTypes.StatsdPluginMetrics{}
			if len(r.Config.Metrics) > len(resp.Config.Metrics) {
				r.Config.Metrics = r.Config.Metrics[:len(resp.Config.Metrics)]
			}
			for metricsCount, metricsItem := range resp.Config.Metrics {
				var metrics tfTypes.StatsdPluginMetrics
				if metricsItem.ConsumerIdentifier != nil {
					metrics.ConsumerIdentifier = types.StringValue(string(*metricsItem.ConsumerIdentifier))
				} else {
					metrics.ConsumerIdentifier = types.StringNull()
				}
				metrics.Name = types.StringValue(string(metricsItem.Name))
				metrics.SampleRate = types.Float64PointerValue(metricsItem.SampleRate)
				if metricsItem.ServiceIdentifier != nil {
					metrics.ServiceIdentifier = types.StringValue(string(*metricsItem.ServiceIdentifier))
				} else {
					metrics.ServiceIdentifier = types.StringNull()
				}
				metrics.StatType = types.StringValue(string(metricsItem.StatType))
				if metricsItem.WorkspaceIdentifier != nil {
					metrics.WorkspaceIdentifier = types.StringValue(string(*metricsItem.WorkspaceIdentifier))
				} else {
					metrics.WorkspaceIdentifier = types.StringNull()
				}
				if metricsCount+1 > len(r.Config.Metrics) {
					r.Config.Metrics = append(r.Config.Metrics, metrics)
				} else {
					r.Config.Metrics[metricsCount].ConsumerIdentifier = metrics.ConsumerIdentifier
					r.Config.Metrics[metricsCount].Name = metrics.Name
					r.Config.Metrics[metricsCount].SampleRate = metrics.SampleRate
					r.Config.Metrics[metricsCount].ServiceIdentifier = metrics.ServiceIdentifier
					r.Config.Metrics[metricsCount].StatType = metrics.StatType
					r.Config.Metrics[metricsCount].WorkspaceIdentifier = metrics.WorkspaceIdentifier
				}
			}
			r.Config.Port = types.Int64PointerValue(resp.Config.Port)
			r.Config.Prefix = types.StringPointerValue(resp.Config.Prefix)
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
			r.Config.QueueSize = types.Int64PointerValue(resp.Config.QueueSize)
			r.Config.RetryCount = types.Int64PointerValue(resp.Config.RetryCount)
			if resp.Config.ServiceIdentifierDefault != nil {
				r.Config.ServiceIdentifierDefault = types.StringValue(string(*resp.Config.ServiceIdentifierDefault))
			} else {
				r.Config.ServiceIdentifierDefault = types.StringNull()
			}
			if resp.Config.TagStyle != nil {
				r.Config.TagStyle = types.StringValue(string(*resp.Config.TagStyle))
			} else {
				r.Config.TagStyle = types.StringNull()
			}
			r.Config.UDPPacketSize = types.Float64PointerValue(resp.Config.UDPPacketSize)
			r.Config.UseTCP = types.BoolPointerValue(resp.Config.UseTCP)
			if resp.Config.WorkspaceIdentifierDefault != nil {
				r.Config.WorkspaceIdentifierDefault = types.StringValue(string(*resp.Config.WorkspaceIdentifierDefault))
			} else {
				r.Config.WorkspaceIdentifierDefault = types.StringNull()
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
