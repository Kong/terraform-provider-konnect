// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginStatsdDataSourceModel) RefreshFromSharedStatsdPlugin(resp *shared.StatsdPlugin) {
	if resp != nil {
		r.Config.AllowStatusCodes = []types.String{}
		for _, v := range resp.Config.AllowStatusCodes {
			r.Config.AllowStatusCodes = append(r.Config.AllowStatusCodes, types.StringValue(v))
		}
		if resp.Config.ConsumerIdentifierDefault != nil {
			r.Config.ConsumerIdentifierDefault = types.StringValue(string(*resp.Config.ConsumerIdentifierDefault))
		} else {
			r.Config.ConsumerIdentifierDefault = types.StringNull()
		}
		if resp.Config.FlushTimeout != nil {
			r.Config.FlushTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.FlushTimeout)))
		} else {
			r.Config.FlushTimeout = types.NumberNull()
		}
		r.Config.Host = types.StringPointerValue(resp.Config.Host)
		r.Config.HostnameInPrefix = types.BoolPointerValue(resp.Config.HostnameInPrefix)
		r.Config.Metrics = []tfTypes.StatsdPluginMetrics{}
		if len(r.Config.Metrics) > len(resp.Config.Metrics) {
			r.Config.Metrics = r.Config.Metrics[:len(resp.Config.Metrics)]
		}
		for metricsCount, metricsItem := range resp.Config.Metrics {
			var metrics1 tfTypes.StatsdPluginMetrics
			if metricsItem.ConsumerIdentifier != nil {
				metrics1.ConsumerIdentifier = types.StringValue(string(*metricsItem.ConsumerIdentifier))
			} else {
				metrics1.ConsumerIdentifier = types.StringNull()
			}
			metrics1.Name = types.StringValue(string(metricsItem.Name))
			if metricsItem.SampleRate != nil {
				metrics1.SampleRate = types.NumberValue(big.NewFloat(float64(*metricsItem.SampleRate)))
			} else {
				metrics1.SampleRate = types.NumberNull()
			}
			if metricsItem.ServiceIdentifier != nil {
				metrics1.ServiceIdentifier = types.StringValue(string(*metricsItem.ServiceIdentifier))
			} else {
				metrics1.ServiceIdentifier = types.StringNull()
			}
			metrics1.StatType = types.StringValue(string(metricsItem.StatType))
			if metricsItem.WorkspaceIdentifier != nil {
				metrics1.WorkspaceIdentifier = types.StringValue(string(*metricsItem.WorkspaceIdentifier))
			} else {
				metrics1.WorkspaceIdentifier = types.StringNull()
			}
			if metricsCount+1 > len(r.Config.Metrics) {
				r.Config.Metrics = append(r.Config.Metrics, metrics1)
			} else {
				r.Config.Metrics[metricsCount].ConsumerIdentifier = metrics1.ConsumerIdentifier
				r.Config.Metrics[metricsCount].Name = metrics1.Name
				r.Config.Metrics[metricsCount].SampleRate = metrics1.SampleRate
				r.Config.Metrics[metricsCount].ServiceIdentifier = metrics1.ServiceIdentifier
				r.Config.Metrics[metricsCount].StatType = metrics1.StatType
				r.Config.Metrics[metricsCount].WorkspaceIdentifier = metrics1.WorkspaceIdentifier
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
			if resp.Config.Queue.InitialRetryDelay != nil {
				r.Config.Queue.InitialRetryDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.InitialRetryDelay)))
			} else {
				r.Config.Queue.InitialRetryDelay = types.NumberNull()
			}
			r.Config.Queue.MaxBatchSize = types.Int64PointerValue(resp.Config.Queue.MaxBatchSize)
			r.Config.Queue.MaxBytes = types.Int64PointerValue(resp.Config.Queue.MaxBytes)
			if resp.Config.Queue.MaxCoalescingDelay != nil {
				r.Config.Queue.MaxCoalescingDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxCoalescingDelay)))
			} else {
				r.Config.Queue.MaxCoalescingDelay = types.NumberNull()
			}
			r.Config.Queue.MaxEntries = types.Int64PointerValue(resp.Config.Queue.MaxEntries)
			if resp.Config.Queue.MaxRetryDelay != nil {
				r.Config.Queue.MaxRetryDelay = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxRetryDelay)))
			} else {
				r.Config.Queue.MaxRetryDelay = types.NumberNull()
			}
			if resp.Config.Queue.MaxRetryTime != nil {
				r.Config.Queue.MaxRetryTime = types.NumberValue(big.NewFloat(float64(*resp.Config.Queue.MaxRetryTime)))
			} else {
				r.Config.Queue.MaxRetryTime = types.NumberNull()
			}
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
		if resp.Config.UDPPacketSize != nil {
			r.Config.UDPPacketSize = types.NumberValue(big.NewFloat(float64(*resp.Config.UDPPacketSize)))
		} else {
			r.Config.UDPPacketSize = types.NumberNull()
		}
		r.Config.UseTCP = types.BoolPointerValue(resp.Config.UseTCP)
		if resp.Config.WorkspaceIdentifierDefault != nil {
			r.Config.WorkspaceIdentifierDefault = types.StringValue(string(*resp.Config.WorkspaceIdentifierDefault))
		} else {
			r.Config.WorkspaceIdentifierDefault = types.StringNull()
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