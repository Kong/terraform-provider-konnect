// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateOpentelemetryPluginHeaderType string

const (
	CreateOpentelemetryPluginHeaderTypePreserve CreateOpentelemetryPluginHeaderType = "preserve"
	CreateOpentelemetryPluginHeaderTypeIgnore   CreateOpentelemetryPluginHeaderType = "ignore"
	CreateOpentelemetryPluginHeaderTypeB3       CreateOpentelemetryPluginHeaderType = "b3"
	CreateOpentelemetryPluginHeaderTypeB3Single CreateOpentelemetryPluginHeaderType = "b3-single"
	CreateOpentelemetryPluginHeaderTypeW3c      CreateOpentelemetryPluginHeaderType = "w3c"
	CreateOpentelemetryPluginHeaderTypeJaeger   CreateOpentelemetryPluginHeaderType = "jaeger"
	CreateOpentelemetryPluginHeaderTypeOt       CreateOpentelemetryPluginHeaderType = "ot"
	CreateOpentelemetryPluginHeaderTypeAws      CreateOpentelemetryPluginHeaderType = "aws"
	CreateOpentelemetryPluginHeaderTypeGcp      CreateOpentelemetryPluginHeaderType = "gcp"
	CreateOpentelemetryPluginHeaderTypeDatadog  CreateOpentelemetryPluginHeaderType = "datadog"
)

func (e CreateOpentelemetryPluginHeaderType) ToPointer() *CreateOpentelemetryPluginHeaderType {
	return &e
}
func (e *CreateOpentelemetryPluginHeaderType) UnmarshalJSON(data []byte) error {
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
		*e = CreateOpentelemetryPluginHeaderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOpentelemetryPluginHeaderType: %v", v)
	}
}

// CreateOpentelemetryPluginDefaultFormat - The default header format to use when extractors did not match any format in the incoming headers and `inject` is configured with the value: `preserve`. This can happen when no tracing header was found in the request, or the incoming tracing header formats were not included in `extract`.
type CreateOpentelemetryPluginDefaultFormat string

const (
	CreateOpentelemetryPluginDefaultFormatB3       CreateOpentelemetryPluginDefaultFormat = "b3"
	CreateOpentelemetryPluginDefaultFormatGcp      CreateOpentelemetryPluginDefaultFormat = "gcp"
	CreateOpentelemetryPluginDefaultFormatB3Single CreateOpentelemetryPluginDefaultFormat = "b3-single"
	CreateOpentelemetryPluginDefaultFormatJaeger   CreateOpentelemetryPluginDefaultFormat = "jaeger"
	CreateOpentelemetryPluginDefaultFormatAws      CreateOpentelemetryPluginDefaultFormat = "aws"
	CreateOpentelemetryPluginDefaultFormatOt       CreateOpentelemetryPluginDefaultFormat = "ot"
	CreateOpentelemetryPluginDefaultFormatW3c      CreateOpentelemetryPluginDefaultFormat = "w3c"
	CreateOpentelemetryPluginDefaultFormatDatadog  CreateOpentelemetryPluginDefaultFormat = "datadog"
)

func (e CreateOpentelemetryPluginDefaultFormat) ToPointer() *CreateOpentelemetryPluginDefaultFormat {
	return &e
}
func (e *CreateOpentelemetryPluginDefaultFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "b3":
		fallthrough
	case "gcp":
		fallthrough
	case "b3-single":
		fallthrough
	case "jaeger":
		fallthrough
	case "aws":
		fallthrough
	case "ot":
		fallthrough
	case "w3c":
		fallthrough
	case "datadog":
		*e = CreateOpentelemetryPluginDefaultFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOpentelemetryPluginDefaultFormat: %v", v)
	}
}

type CreateOpentelemetryPluginExtract string

const (
	CreateOpentelemetryPluginExtractB3      CreateOpentelemetryPluginExtract = "b3"
	CreateOpentelemetryPluginExtractGcp     CreateOpentelemetryPluginExtract = "gcp"
	CreateOpentelemetryPluginExtractJaeger  CreateOpentelemetryPluginExtract = "jaeger"
	CreateOpentelemetryPluginExtractAws     CreateOpentelemetryPluginExtract = "aws"
	CreateOpentelemetryPluginExtractOt      CreateOpentelemetryPluginExtract = "ot"
	CreateOpentelemetryPluginExtractW3c     CreateOpentelemetryPluginExtract = "w3c"
	CreateOpentelemetryPluginExtractDatadog CreateOpentelemetryPluginExtract = "datadog"
)

func (e CreateOpentelemetryPluginExtract) ToPointer() *CreateOpentelemetryPluginExtract {
	return &e
}
func (e *CreateOpentelemetryPluginExtract) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "b3":
		fallthrough
	case "gcp":
		fallthrough
	case "jaeger":
		fallthrough
	case "aws":
		fallthrough
	case "ot":
		fallthrough
	case "w3c":
		fallthrough
	case "datadog":
		*e = CreateOpentelemetryPluginExtract(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOpentelemetryPluginExtract: %v", v)
	}
}

type CreateOpentelemetryPluginInject string

const (
	CreateOpentelemetryPluginInjectPreserve CreateOpentelemetryPluginInject = "preserve"
	CreateOpentelemetryPluginInjectB3       CreateOpentelemetryPluginInject = "b3"
	CreateOpentelemetryPluginInjectGcp      CreateOpentelemetryPluginInject = "gcp"
	CreateOpentelemetryPluginInjectB3Single CreateOpentelemetryPluginInject = "b3-single"
	CreateOpentelemetryPluginInjectJaeger   CreateOpentelemetryPluginInject = "jaeger"
	CreateOpentelemetryPluginInjectAws      CreateOpentelemetryPluginInject = "aws"
	CreateOpentelemetryPluginInjectOt       CreateOpentelemetryPluginInject = "ot"
	CreateOpentelemetryPluginInjectW3c      CreateOpentelemetryPluginInject = "w3c"
	CreateOpentelemetryPluginInjectDatadog  CreateOpentelemetryPluginInject = "datadog"
)

func (e CreateOpentelemetryPluginInject) ToPointer() *CreateOpentelemetryPluginInject {
	return &e
}
func (e *CreateOpentelemetryPluginInject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preserve":
		fallthrough
	case "b3":
		fallthrough
	case "gcp":
		fallthrough
	case "b3-single":
		fallthrough
	case "jaeger":
		fallthrough
	case "aws":
		fallthrough
	case "ot":
		fallthrough
	case "w3c":
		fallthrough
	case "datadog":
		*e = CreateOpentelemetryPluginInject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOpentelemetryPluginInject: %v", v)
	}
}

type CreateOpentelemetryPluginPropagation struct {
	// Header names to clear after context extraction. This allows to extract the context from a certain header and then remove it from the request, useful when extraction and injection are performed on different header formats and the original header should not be sent to the upstream. If left empty, no headers are cleared.
	Clear []string `json:"clear,omitempty"`
	// The default header format to use when extractors did not match any format in the incoming headers and `inject` is configured with the value: `preserve`. This can happen when no tracing header was found in the request, or the incoming tracing header formats were not included in `extract`.
	DefaultFormat CreateOpentelemetryPluginDefaultFormat `json:"default_format"`
	// Header formats used to extract tracing context from incoming requests. If multiple values are specified, the first one found will be used for extraction. If left empty, Kong will not extract any tracing context information from incoming requests and generate a trace with no parent and a new trace ID.
	Extract []CreateOpentelemetryPluginExtract `json:"extract,omitempty"`
	// Header formats used to inject tracing context. The value `preserve` will use the same header format as the incoming request. If multiple values are specified, all of them will be used during injection. If left empty, Kong will not inject any tracing context information in outgoing requests.
	Inject []CreateOpentelemetryPluginInject `json:"inject,omitempty"`
}

func (o *CreateOpentelemetryPluginPropagation) GetClear() []string {
	if o == nil {
		return nil
	}
	return o.Clear
}

func (o *CreateOpentelemetryPluginPropagation) GetDefaultFormat() CreateOpentelemetryPluginDefaultFormat {
	if o == nil {
		return CreateOpentelemetryPluginDefaultFormat("")
	}
	return o.DefaultFormat
}

func (o *CreateOpentelemetryPluginPropagation) GetExtract() []CreateOpentelemetryPluginExtract {
	if o == nil {
		return nil
	}
	return o.Extract
}

func (o *CreateOpentelemetryPluginPropagation) GetInject() []CreateOpentelemetryPluginInject {
	if o == nil {
		return nil
	}
	return o.Inject
}

type CreateOpentelemetryPluginQueue struct {
	// Time in seconds before the initial retry is made for a failing batch.
	InitialRetryDelay *float64 `json:"initial_retry_delay,omitempty"`
	// Maximum number of entries that can be processed at a time.
	MaxBatchSize *int64 `json:"max_batch_size,omitempty"`
	// Maximum number of bytes that can be waiting on a queue, requires string content.
	MaxBytes *int64 `json:"max_bytes,omitempty"`
	// Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.
	MaxCoalescingDelay *float64 `json:"max_coalescing_delay,omitempty"`
	// Maximum number of entries that can be waiting on the queue.
	MaxEntries *int64 `json:"max_entries,omitempty"`
	// Maximum time in seconds between retries, caps exponential backoff.
	MaxRetryDelay *float64 `json:"max_retry_delay,omitempty"`
	// Time in seconds before the queue gives up calling a failed handler for a batch.
	MaxRetryTime *float64 `json:"max_retry_time,omitempty"`
}

func (o *CreateOpentelemetryPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *CreateOpentelemetryPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *CreateOpentelemetryPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *CreateOpentelemetryPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *CreateOpentelemetryPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *CreateOpentelemetryPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *CreateOpentelemetryPluginQueue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

type CreateOpentelemetryPluginConfig struct {
	// The delay, in seconds, between two consecutive batches.
	BatchFlushDelay *int64 `json:"batch_flush_delay,omitempty"`
	// The number of spans to be sent in a single batch.
	BatchSpanCount *int64 `json:"batch_span_count,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	Endpoint   *string                              `json:"endpoint,omitempty"`
	HeaderType *CreateOpentelemetryPluginHeaderType `json:"header_type,omitempty"`
	// The custom headers to be added in the HTTP request sent to the OTLP server. This setting is useful for adding the authentication headers (token) for the APM backend.
	Headers                      map[string]any                        `json:"headers,omitempty"`
	HTTPResponseHeaderForTraceid *string                               `json:"http_response_header_for_traceid,omitempty"`
	Propagation                  *CreateOpentelemetryPluginPropagation `json:"propagation,omitempty"`
	Queue                        *CreateOpentelemetryPluginQueue       `json:"queue,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout        *int64         `json:"read_timeout,omitempty"`
	ResourceAttributes map[string]any `json:"resource_attributes,omitempty"`
	// Tracing sampling rate for configuring the probability-based sampler. When set, this value supersedes the global `tracing_sampling_rate` setting from kong.conf.
	SamplingRate *float64 `json:"sampling_rate,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
}

func (o *CreateOpentelemetryPluginConfig) GetBatchFlushDelay() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchFlushDelay
}

func (o *CreateOpentelemetryPluginConfig) GetBatchSpanCount() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchSpanCount
}

func (o *CreateOpentelemetryPluginConfig) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *CreateOpentelemetryPluginConfig) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *CreateOpentelemetryPluginConfig) GetHeaderType() *CreateOpentelemetryPluginHeaderType {
	if o == nil {
		return nil
	}
	return o.HeaderType
}

func (o *CreateOpentelemetryPluginConfig) GetHeaders() map[string]any {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateOpentelemetryPluginConfig) GetHTTPResponseHeaderForTraceid() *string {
	if o == nil {
		return nil
	}
	return o.HTTPResponseHeaderForTraceid
}

func (o *CreateOpentelemetryPluginConfig) GetPropagation() *CreateOpentelemetryPluginPropagation {
	if o == nil {
		return nil
	}
	return o.Propagation
}

func (o *CreateOpentelemetryPluginConfig) GetQueue() *CreateOpentelemetryPluginQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *CreateOpentelemetryPluginConfig) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *CreateOpentelemetryPluginConfig) GetResourceAttributes() map[string]any {
	if o == nil {
		return nil
	}
	return o.ResourceAttributes
}

func (o *CreateOpentelemetryPluginConfig) GetSamplingRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SamplingRate
}

func (o *CreateOpentelemetryPluginConfig) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

type CreateOpentelemetryPluginProtocols string

const (
	CreateOpentelemetryPluginProtocolsGrpc           CreateOpentelemetryPluginProtocols = "grpc"
	CreateOpentelemetryPluginProtocolsGrpcs          CreateOpentelemetryPluginProtocols = "grpcs"
	CreateOpentelemetryPluginProtocolsHTTP           CreateOpentelemetryPluginProtocols = "http"
	CreateOpentelemetryPluginProtocolsHTTPS          CreateOpentelemetryPluginProtocols = "https"
	CreateOpentelemetryPluginProtocolsTCP            CreateOpentelemetryPluginProtocols = "tcp"
	CreateOpentelemetryPluginProtocolsTLS            CreateOpentelemetryPluginProtocols = "tls"
	CreateOpentelemetryPluginProtocolsTLSPassthrough CreateOpentelemetryPluginProtocols = "tls_passthrough"
	CreateOpentelemetryPluginProtocolsUDP            CreateOpentelemetryPluginProtocols = "udp"
	CreateOpentelemetryPluginProtocolsWs             CreateOpentelemetryPluginProtocols = "ws"
	CreateOpentelemetryPluginProtocolsWss            CreateOpentelemetryPluginProtocols = "wss"
)

func (e CreateOpentelemetryPluginProtocols) ToPointer() *CreateOpentelemetryPluginProtocols {
	return &e
}
func (e *CreateOpentelemetryPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateOpentelemetryPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOpentelemetryPluginProtocols: %v", v)
	}
}

// CreateOpentelemetryPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateOpentelemetryPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpentelemetryPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateOpentelemetryPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpentelemetryPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateOpentelemetryPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateOpentelemetryPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpentelemetryPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateOpentelemetryPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateOpentelemetryPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpentelemetryPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateOpentelemetryPlugin struct {
	Config *CreateOpentelemetryPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"opentelemetry" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateOpentelemetryPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateOpentelemetryPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateOpentelemetryPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateOpentelemetryPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateOpentelemetryPluginService `json:"service,omitempty"`
}

func (c CreateOpentelemetryPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateOpentelemetryPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateOpentelemetryPlugin) GetConfig() *CreateOpentelemetryPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateOpentelemetryPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateOpentelemetryPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateOpentelemetryPlugin) GetName() *string {
	return types.String("opentelemetry")
}

func (o *CreateOpentelemetryPlugin) GetProtocols() []CreateOpentelemetryPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateOpentelemetryPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateOpentelemetryPlugin) GetConsumer() *CreateOpentelemetryPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateOpentelemetryPlugin) GetConsumerGroup() *CreateOpentelemetryPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateOpentelemetryPlugin) GetRoute() *CreateOpentelemetryPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateOpentelemetryPlugin) GetService() *CreateOpentelemetryPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
