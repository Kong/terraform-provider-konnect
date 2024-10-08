// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginHTTPLogResourceModel) ToSharedCreateHTTPLogPlugin() *shared.CreateHTTPLogPlugin {
	var config *shared.CreateHTTPLogPluginConfig
	if r.Config != nil {
		contentType := new(shared.CreateHTTPLogPluginContentType)
		if !r.Config.ContentType.IsUnknown() && !r.Config.ContentType.IsNull() {
			*contentType = shared.CreateHTTPLogPluginContentType(r.Config.ContentType.ValueString())
		} else {
			contentType = nil
		}
		customFieldsByLua := make(map[string]interface{})
		for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
			var customFieldsByLuaInst interface{}
			_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
			customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
		}
		flushTimeout := new(float64)
		if !r.Config.FlushTimeout.IsUnknown() && !r.Config.FlushTimeout.IsNull() {
			*flushTimeout, _ = r.Config.FlushTimeout.ValueBigFloat().Float64()
		} else {
			flushTimeout = nil
		}
		headers := make(map[string]interface{})
		for headersKey, headersValue := range r.Config.Headers {
			var headersInst interface{}
			_ = json.Unmarshal([]byte(headersValue.ValueString()), &headersInst)
			headers[headersKey] = headersInst
		}
		httpEndpoint := new(string)
		if !r.Config.HTTPEndpoint.IsUnknown() && !r.Config.HTTPEndpoint.IsNull() {
			*httpEndpoint = r.Config.HTTPEndpoint.ValueString()
		} else {
			httpEndpoint = nil
		}
		keepalive := new(float64)
		if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
			*keepalive, _ = r.Config.Keepalive.ValueBigFloat().Float64()
		} else {
			keepalive = nil
		}
		method := new(shared.CreateHTTPLogPluginMethod)
		if !r.Config.Method.IsUnknown() && !r.Config.Method.IsNull() {
			*method = shared.CreateHTTPLogPluginMethod(r.Config.Method.ValueString())
		} else {
			method = nil
		}
		var queue *shared.CreateHTTPLogPluginQueue
		if r.Config.Queue != nil {
			concurrencyLimit := new(shared.CreateHTTPLogPluginConcurrencyLimit)
			if !r.Config.Queue.ConcurrencyLimit.IsUnknown() && !r.Config.Queue.ConcurrencyLimit.IsNull() {
				*concurrencyLimit = shared.CreateHTTPLogPluginConcurrencyLimit(r.Config.Queue.ConcurrencyLimit.ValueInt64())
			} else {
				concurrencyLimit = nil
			}
			initialRetryDelay := new(float64)
			if !r.Config.Queue.InitialRetryDelay.IsUnknown() && !r.Config.Queue.InitialRetryDelay.IsNull() {
				*initialRetryDelay, _ = r.Config.Queue.InitialRetryDelay.ValueBigFloat().Float64()
			} else {
				initialRetryDelay = nil
			}
			maxBatchSize := new(int64)
			if !r.Config.Queue.MaxBatchSize.IsUnknown() && !r.Config.Queue.MaxBatchSize.IsNull() {
				*maxBatchSize = r.Config.Queue.MaxBatchSize.ValueInt64()
			} else {
				maxBatchSize = nil
			}
			maxBytes := new(int64)
			if !r.Config.Queue.MaxBytes.IsUnknown() && !r.Config.Queue.MaxBytes.IsNull() {
				*maxBytes = r.Config.Queue.MaxBytes.ValueInt64()
			} else {
				maxBytes = nil
			}
			maxCoalescingDelay := new(float64)
			if !r.Config.Queue.MaxCoalescingDelay.IsUnknown() && !r.Config.Queue.MaxCoalescingDelay.IsNull() {
				*maxCoalescingDelay, _ = r.Config.Queue.MaxCoalescingDelay.ValueBigFloat().Float64()
			} else {
				maxCoalescingDelay = nil
			}
			maxEntries := new(int64)
			if !r.Config.Queue.MaxEntries.IsUnknown() && !r.Config.Queue.MaxEntries.IsNull() {
				*maxEntries = r.Config.Queue.MaxEntries.ValueInt64()
			} else {
				maxEntries = nil
			}
			maxRetryDelay := new(float64)
			if !r.Config.Queue.MaxRetryDelay.IsUnknown() && !r.Config.Queue.MaxRetryDelay.IsNull() {
				*maxRetryDelay, _ = r.Config.Queue.MaxRetryDelay.ValueBigFloat().Float64()
			} else {
				maxRetryDelay = nil
			}
			maxRetryTime := new(float64)
			if !r.Config.Queue.MaxRetryTime.IsUnknown() && !r.Config.Queue.MaxRetryTime.IsNull() {
				*maxRetryTime, _ = r.Config.Queue.MaxRetryTime.ValueBigFloat().Float64()
			} else {
				maxRetryTime = nil
			}
			queue = &shared.CreateHTTPLogPluginQueue{
				ConcurrencyLimit:   concurrencyLimit,
				InitialRetryDelay:  initialRetryDelay,
				MaxBatchSize:       maxBatchSize,
				MaxBytes:           maxBytes,
				MaxCoalescingDelay: maxCoalescingDelay,
				MaxEntries:         maxEntries,
				MaxRetryDelay:      maxRetryDelay,
				MaxRetryTime:       maxRetryTime,
			}
		}
		queueSize := new(int64)
		if !r.Config.QueueSize.IsUnknown() && !r.Config.QueueSize.IsNull() {
			*queueSize = r.Config.QueueSize.ValueInt64()
		} else {
			queueSize = nil
		}
		retryCount := new(int64)
		if !r.Config.RetryCount.IsUnknown() && !r.Config.RetryCount.IsNull() {
			*retryCount = r.Config.RetryCount.ValueInt64()
		} else {
			retryCount = nil
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
		} else {
			timeout = nil
		}
		config = &shared.CreateHTTPLogPluginConfig{
			ContentType:       contentType,
			CustomFieldsByLua: customFieldsByLua,
			FlushTimeout:      flushTimeout,
			Headers:           headers,
			HTTPEndpoint:      httpEndpoint,
			Keepalive:         keepalive,
			Method:            method,
			Queue:             queue,
			QueueSize:         queueSize,
			RetryCount:        retryCount,
			Timeout:           timeout,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.CreateHTTPLogPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateHTTPLogPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateHTTPLogPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateHTTPLogPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateHTTPLogPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateHTTPLogPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateHTTPLogPluginProtocols = []shared.CreateHTTPLogPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateHTTPLogPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateHTTPLogPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateHTTPLogPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateHTTPLogPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateHTTPLogPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateHTTPLogPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateHTTPLogPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateHTTPLogPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateHTTPLogPluginService{
			ID: id3,
		}
	}
	out := shared.CreateHTTPLogPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginHTTPLogResourceModel) RefreshFromSharedHTTPLogPlugin(resp *shared.HTTPLogPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateHTTPLogPluginConfig{}
			if resp.Config.ContentType != nil {
				r.Config.ContentType = types.StringValue(string(*resp.Config.ContentType))
			} else {
				r.Config.ContentType = types.StringNull()
			}
			if len(resp.Config.CustomFieldsByLua) > 0 {
				r.Config.CustomFieldsByLua = make(map[string]types.String)
				for key, value := range resp.Config.CustomFieldsByLua {
					result, _ := json.Marshal(value)
					r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
				}
			}
			if resp.Config.FlushTimeout != nil {
				r.Config.FlushTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.FlushTimeout)))
			} else {
				r.Config.FlushTimeout = types.NumberNull()
			}
			if len(resp.Config.Headers) > 0 {
				r.Config.Headers = make(map[string]types.String)
				for key1, value1 := range resp.Config.Headers {
					result1, _ := json.Marshal(value1)
					r.Config.Headers[key1] = types.StringValue(string(result1))
				}
			}
			r.Config.HTTPEndpoint = types.StringPointerValue(resp.Config.HTTPEndpoint)
			if resp.Config.Keepalive != nil {
				r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
			} else {
				r.Config.Keepalive = types.NumberNull()
			}
			if resp.Config.Method != nil {
				r.Config.Method = types.StringValue(string(*resp.Config.Method))
			} else {
				r.Config.Method = types.StringNull()
			}
			if resp.Config.Queue == nil {
				r.Config.Queue = nil
			} else {
				r.Config.Queue = &tfTypes.CreateDatadogPluginQueue{}
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
			if resp.Config.Timeout != nil {
				r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
			} else {
				r.Config.Timeout = types.NumberNull()
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
			r.Ordering = &tfTypes.CreateACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.CreateACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.CreateACLPluginAfter{}
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
