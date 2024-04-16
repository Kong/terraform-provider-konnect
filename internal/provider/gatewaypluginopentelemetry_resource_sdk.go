// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginOpentelemetryResourceModel) ToSharedCreateOpentelemetryPlugin() *shared.CreateOpentelemetryPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateOpentelemetryPluginProtocols = []shared.CreateOpentelemetryPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateOpentelemetryPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateOpentelemetryPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateOpentelemetryPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateOpentelemetryPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateOpentelemetryPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateOpentelemetryPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateOpentelemetryPluginService{
			ID: id2,
		}
	}
	endpoint := new(string)
	if !r.Config.Endpoint.IsUnknown() && !r.Config.Endpoint.IsNull() {
		*endpoint = r.Config.Endpoint.ValueString()
	} else {
		endpoint = nil
	}
	headers := make(map[string]interface{})
	for headersKey, headersValue := range r.Config.Headers {
		var headersInst interface{}
		_ = json.Unmarshal([]byte(headersValue.ValueString()), &headersInst)
		headers[headersKey] = headersInst
	}
	resourceAttributes := make(map[string]interface{})
	for resourceAttributesKey, resourceAttributesValue := range r.Config.ResourceAttributes {
		var resourceAttributesInst interface{}
		_ = json.Unmarshal([]byte(resourceAttributesValue.ValueString()), &resourceAttributesInst)
		resourceAttributes[resourceAttributesKey] = resourceAttributesInst
	}
	var queue *shared.Queue
	if r.Config.Queue != nil {
		maxBatchSize := new(int64)
		if !r.Config.Queue.MaxBatchSize.IsUnknown() && !r.Config.Queue.MaxBatchSize.IsNull() {
			*maxBatchSize = r.Config.Queue.MaxBatchSize.ValueInt64()
		} else {
			maxBatchSize = nil
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
		maxBytes := new(int64)
		if !r.Config.Queue.MaxBytes.IsUnknown() && !r.Config.Queue.MaxBytes.IsNull() {
			*maxBytes = r.Config.Queue.MaxBytes.ValueInt64()
		} else {
			maxBytes = nil
		}
		maxRetryTime := new(float64)
		if !r.Config.Queue.MaxRetryTime.IsUnknown() && !r.Config.Queue.MaxRetryTime.IsNull() {
			*maxRetryTime, _ = r.Config.Queue.MaxRetryTime.ValueBigFloat().Float64()
		} else {
			maxRetryTime = nil
		}
		initialRetryDelay := new(float64)
		if !r.Config.Queue.InitialRetryDelay.IsUnknown() && !r.Config.Queue.InitialRetryDelay.IsNull() {
			*initialRetryDelay, _ = r.Config.Queue.InitialRetryDelay.ValueBigFloat().Float64()
		} else {
			initialRetryDelay = nil
		}
		maxRetryDelay := new(float64)
		if !r.Config.Queue.MaxRetryDelay.IsUnknown() && !r.Config.Queue.MaxRetryDelay.IsNull() {
			*maxRetryDelay, _ = r.Config.Queue.MaxRetryDelay.ValueBigFloat().Float64()
		} else {
			maxRetryDelay = nil
		}
		queue = &shared.Queue{
			MaxBatchSize:       maxBatchSize,
			MaxCoalescingDelay: maxCoalescingDelay,
			MaxEntries:         maxEntries,
			MaxBytes:           maxBytes,
			MaxRetryTime:       maxRetryTime,
			InitialRetryDelay:  initialRetryDelay,
			MaxRetryDelay:      maxRetryDelay,
		}
	}
	batchSpanCount := new(int64)
	if !r.Config.BatchSpanCount.IsUnknown() && !r.Config.BatchSpanCount.IsNull() {
		*batchSpanCount = r.Config.BatchSpanCount.ValueInt64()
	} else {
		batchSpanCount = nil
	}
	batchFlushDelay := new(int64)
	if !r.Config.BatchFlushDelay.IsUnknown() && !r.Config.BatchFlushDelay.IsNull() {
		*batchFlushDelay = r.Config.BatchFlushDelay.ValueInt64()
	} else {
		batchFlushDelay = nil
	}
	connectTimeout := new(int64)
	if !r.Config.ConnectTimeout.IsUnknown() && !r.Config.ConnectTimeout.IsNull() {
		*connectTimeout = r.Config.ConnectTimeout.ValueInt64()
	} else {
		connectTimeout = nil
	}
	sendTimeout := new(int64)
	if !r.Config.SendTimeout.IsUnknown() && !r.Config.SendTimeout.IsNull() {
		*sendTimeout = r.Config.SendTimeout.ValueInt64()
	} else {
		sendTimeout = nil
	}
	readTimeout := new(int64)
	if !r.Config.ReadTimeout.IsUnknown() && !r.Config.ReadTimeout.IsNull() {
		*readTimeout = r.Config.ReadTimeout.ValueInt64()
	} else {
		readTimeout = nil
	}
	httpResponseHeaderForTraceid := new(string)
	if !r.Config.HTTPResponseHeaderForTraceid.IsUnknown() && !r.Config.HTTPResponseHeaderForTraceid.IsNull() {
		*httpResponseHeaderForTraceid = r.Config.HTTPResponseHeaderForTraceid.ValueString()
	} else {
		httpResponseHeaderForTraceid = nil
	}
	headerType := new(shared.HeaderType)
	if !r.Config.HeaderType.IsUnknown() && !r.Config.HeaderType.IsNull() {
		*headerType = shared.HeaderType(r.Config.HeaderType.ValueString())
	} else {
		headerType = nil
	}
	samplingRate := new(float64)
	if !r.Config.SamplingRate.IsUnknown() && !r.Config.SamplingRate.IsNull() {
		*samplingRate, _ = r.Config.SamplingRate.ValueBigFloat().Float64()
	} else {
		samplingRate = nil
	}
	config := shared.CreateOpentelemetryPluginConfig{
		Endpoint:                     endpoint,
		Headers:                      headers,
		ResourceAttributes:           resourceAttributes,
		Queue:                        queue,
		BatchSpanCount:               batchSpanCount,
		BatchFlushDelay:              batchFlushDelay,
		ConnectTimeout:               connectTimeout,
		SendTimeout:                  sendTimeout,
		ReadTimeout:                  readTimeout,
		HTTPResponseHeaderForTraceid: httpResponseHeaderForTraceid,
		HeaderType:                   headerType,
		SamplingRate:                 samplingRate,
	}
	out := shared.CreateOpentelemetryPlugin{
		Enabled:   enabled,
		Protocols: protocols,
		Tags:      tags,
		Consumer:  consumer,
		Route:     route,
		Service:   service,
		Config:    config,
	}
	return &out
}

func (r *GatewayPluginOpentelemetryResourceModel) RefreshFromSharedOpentelemetryPlugin(resp *shared.OpentelemetryPlugin) {
	if resp != nil {
		r.Config.BatchFlushDelay = types.Int64PointerValue(resp.Config.BatchFlushDelay)
		r.Config.BatchSpanCount = types.Int64PointerValue(resp.Config.BatchSpanCount)
		r.Config.ConnectTimeout = types.Int64PointerValue(resp.Config.ConnectTimeout)
		r.Config.Endpoint = types.StringPointerValue(resp.Config.Endpoint)
		if resp.Config.HeaderType != nil {
			r.Config.HeaderType = types.StringValue(string(*resp.Config.HeaderType))
		} else {
			r.Config.HeaderType = types.StringNull()
		}
		if len(resp.Config.Headers) > 0 {
			r.Config.Headers = make(map[string]types.String)
			for key, value := range resp.Config.Headers {
				result, _ := json.Marshal(value)
				r.Config.Headers[key] = types.StringValue(string(result))
			}
		}
		r.Config.HTTPResponseHeaderForTraceid = types.StringPointerValue(resp.Config.HTTPResponseHeaderForTraceid)
		if resp.Config.Queue == nil {
			r.Config.Queue = nil
		} else {
			r.Config.Queue = &tfTypes.Queue{}
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
		r.Config.ReadTimeout = types.Int64PointerValue(resp.Config.ReadTimeout)
		if len(resp.Config.ResourceAttributes) > 0 {
			r.Config.ResourceAttributes = make(map[string]types.String)
			for key1, value1 := range resp.Config.ResourceAttributes {
				result1, _ := json.Marshal(value1)
				r.Config.ResourceAttributes[key1] = types.StringValue(string(result1))
			}
		}
		if resp.Config.SamplingRate != nil {
			r.Config.SamplingRate = types.NumberValue(big.NewFloat(float64(*resp.Config.SamplingRate)))
		} else {
			r.Config.SamplingRate = types.NumberNull()
		}
		r.Config.SendTimeout = types.Int64PointerValue(resp.Config.SendTimeout)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
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
	}
}
