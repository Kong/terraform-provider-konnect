// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type OpentelemetryPluginProtocols string

const (
	OpentelemetryPluginProtocolsGrpc           OpentelemetryPluginProtocols = "grpc"
	OpentelemetryPluginProtocolsGrpcs          OpentelemetryPluginProtocols = "grpcs"
	OpentelemetryPluginProtocolsHTTP           OpentelemetryPluginProtocols = "http"
	OpentelemetryPluginProtocolsHTTPS          OpentelemetryPluginProtocols = "https"
	OpentelemetryPluginProtocolsTCP            OpentelemetryPluginProtocols = "tcp"
	OpentelemetryPluginProtocolsTLS            OpentelemetryPluginProtocols = "tls"
	OpentelemetryPluginProtocolsTLSPassthrough OpentelemetryPluginProtocols = "tls_passthrough"
	OpentelemetryPluginProtocolsUDP            OpentelemetryPluginProtocols = "udp"
	OpentelemetryPluginProtocolsWs             OpentelemetryPluginProtocols = "ws"
	OpentelemetryPluginProtocolsWss            OpentelemetryPluginProtocols = "wss"
)

func (e OpentelemetryPluginProtocols) ToPointer() *OpentelemetryPluginProtocols {
	return &e
}

func (e *OpentelemetryPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = OpentelemetryPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OpentelemetryPluginProtocols: %v", v)
	}
}

// OpentelemetryPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type OpentelemetryPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *OpentelemetryPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// OpentelemetryPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type OpentelemetryPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *OpentelemetryPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// OpentelemetryPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type OpentelemetryPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *OpentelemetryPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type HeaderType string

const (
	HeaderTypePreserve HeaderType = "preserve"
	HeaderTypeIgnore   HeaderType = "ignore"
	HeaderTypeB3       HeaderType = "b3"
	HeaderTypeB3Single HeaderType = "b3-single"
	HeaderTypeW3c      HeaderType = "w3c"
	HeaderTypeJaeger   HeaderType = "jaeger"
	HeaderTypeOt       HeaderType = "ot"
	HeaderTypeAws      HeaderType = "aws"
	HeaderTypeGcp      HeaderType = "gcp"
	HeaderTypeDatadog  HeaderType = "datadog"
)

func (e HeaderType) ToPointer() *HeaderType {
	return &e
}

func (e *HeaderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preserve":
		fallthrough
	case "ignore":
		fallthrough
	case "b3":
		fallthrough
	case "b3-single":
		fallthrough
	case "w3c":
		fallthrough
	case "jaeger":
		fallthrough
	case "ot":
		fallthrough
	case "aws":
		fallthrough
	case "gcp":
		fallthrough
	case "datadog":
		*e = HeaderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HeaderType: %v", v)
	}
}

type Queue struct {
	// Time in seconds before the initial retry is made for a failing batch.
	InitialRetryDelay *float64 `default:"0.01" json:"initial_retry_delay"`
	// Maximum number of entries that can be processed at a time.
	MaxBatchSize *int64 `default:"1" json:"max_batch_size"`
	// Maximum number of bytes that can be waiting on a queue, requires string content.
	MaxBytes *int64 `json:"max_bytes,omitempty"`
	// Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.
	MaxCoalescingDelay *float64 `default:"1" json:"max_coalescing_delay"`
	// Maximum number of entries that can be waiting on the queue.
	MaxEntries *int64 `default:"10000" json:"max_entries"`
	// Maximum time in seconds between retries, caps exponential backoff.
	MaxRetryDelay *float64 `default:"60" json:"max_retry_delay"`
	// Time in seconds before the queue gives up calling a failed handler for a batch.
	MaxRetryTime *float64 `default:"60" json:"max_retry_time"`
}

func (q Queue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(q, "", false)
}

func (q *Queue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &q, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Queue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *Queue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *Queue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *Queue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *Queue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *Queue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *Queue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

type OpentelemetryPluginConfig struct {
	// The delay, in seconds, between two consecutive batches.
	BatchFlushDelay *int64 `json:"batch_flush_delay,omitempty"`
	// The number of spans to be sent in a single batch.
	BatchSpanCount *int64 `json:"batch_span_count,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ConnectTimeout *int64 `default:"1000" json:"connect_timeout"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	Endpoint   *string     `json:"endpoint,omitempty"`
	HeaderType *HeaderType `default:"preserve" json:"header_type"`
	// The custom headers to be added in the HTTP request sent to the OTLP server. This setting is useful for adding the authentication headers (token) for the APM backend.
	Headers                      map[string]interface{} `json:"headers,omitempty"`
	HTTPResponseHeaderForTraceid *string                `json:"http_response_header_for_traceid,omitempty"`
	Queue                        *Queue                 `json:"queue,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout        *int64                 `default:"5000" json:"read_timeout"`
	ResourceAttributes map[string]interface{} `json:"resource_attributes,omitempty"`
	// Tracing sampling rate for configuring the probability-based sampler. When set, this value supersedes the global `tracing_sampling_rate` setting from kong.conf.
	SamplingRate *float64 `json:"sampling_rate,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `default:"5000" json:"send_timeout"`
}

func (o OpentelemetryPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OpentelemetryPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OpentelemetryPluginConfig) GetBatchFlushDelay() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchFlushDelay
}

func (o *OpentelemetryPluginConfig) GetBatchSpanCount() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchSpanCount
}

func (o *OpentelemetryPluginConfig) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *OpentelemetryPluginConfig) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *OpentelemetryPluginConfig) GetHeaderType() *HeaderType {
	if o == nil {
		return nil
	}
	return o.HeaderType
}

func (o *OpentelemetryPluginConfig) GetHeaders() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *OpentelemetryPluginConfig) GetHTTPResponseHeaderForTraceid() *string {
	if o == nil {
		return nil
	}
	return o.HTTPResponseHeaderForTraceid
}

func (o *OpentelemetryPluginConfig) GetQueue() *Queue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *OpentelemetryPluginConfig) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *OpentelemetryPluginConfig) GetResourceAttributes() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.ResourceAttributes
}

func (o *OpentelemetryPluginConfig) GetSamplingRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SamplingRate
}

func (o *OpentelemetryPluginConfig) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

// OpentelemetryPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type OpentelemetryPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"opentelemetry" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []OpentelemetryPluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *OpentelemetryPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *OpentelemetryPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *OpentelemetryPluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64                    `json:"created_at,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	Config    OpentelemetryPluginConfig `json:"config"`
}

func (o OpentelemetryPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OpentelemetryPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OpentelemetryPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *OpentelemetryPlugin) GetName() string {
	return "opentelemetry"
}

func (o *OpentelemetryPlugin) GetProtocols() []OpentelemetryPluginProtocols {
	if o == nil {
		return []OpentelemetryPluginProtocols{}
	}
	return o.Protocols
}

func (o *OpentelemetryPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *OpentelemetryPlugin) GetConsumer() *OpentelemetryPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *OpentelemetryPlugin) GetRoute() *OpentelemetryPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *OpentelemetryPlugin) GetService() *OpentelemetryPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *OpentelemetryPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *OpentelemetryPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OpentelemetryPlugin) GetConfig() OpentelemetryPluginConfig {
	if o == nil {
		return OpentelemetryPluginConfig{}
	}
	return o.Config
}
