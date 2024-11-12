// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

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

// DefaultFormat - The default header format to use when extractors did not match any format in the incoming headers and `inject` is configured with the value: `preserve`. This can happen when no tracing header was found in the request, or the incoming tracing header formats were not included in `extract`.
type DefaultFormat string

const (
	DefaultFormatW3c      DefaultFormat = "w3c"
	DefaultFormatDatadog  DefaultFormat = "datadog"
	DefaultFormatB3       DefaultFormat = "b3"
	DefaultFormatGcp      DefaultFormat = "gcp"
	DefaultFormatB3Single DefaultFormat = "b3-single"
	DefaultFormatJaeger   DefaultFormat = "jaeger"
	DefaultFormatAws      DefaultFormat = "aws"
	DefaultFormatOt       DefaultFormat = "ot"
)

func (e DefaultFormat) ToPointer() *DefaultFormat {
	return &e
}
func (e *DefaultFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "w3c":
		fallthrough
	case "datadog":
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
		*e = DefaultFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DefaultFormat: %v", v)
	}
}

type Extract string

const (
	ExtractW3c     Extract = "w3c"
	ExtractDatadog Extract = "datadog"
	ExtractB3      Extract = "b3"
	ExtractGcp     Extract = "gcp"
	ExtractJaeger  Extract = "jaeger"
	ExtractAws     Extract = "aws"
	ExtractOt      Extract = "ot"
)

func (e Extract) ToPointer() *Extract {
	return &e
}
func (e *Extract) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "w3c":
		fallthrough
	case "datadog":
		fallthrough
	case "b3":
		fallthrough
	case "gcp":
		fallthrough
	case "jaeger":
		fallthrough
	case "aws":
		fallthrough
	case "ot":
		*e = Extract(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Extract: %v", v)
	}
}

type Inject string

const (
	InjectPreserve Inject = "preserve"
	InjectW3c      Inject = "w3c"
	InjectDatadog  Inject = "datadog"
	InjectB3       Inject = "b3"
	InjectGcp      Inject = "gcp"
	InjectB3Single Inject = "b3-single"
	InjectJaeger   Inject = "jaeger"
	InjectAws      Inject = "aws"
	InjectOt       Inject = "ot"
)

func (e Inject) ToPointer() *Inject {
	return &e
}
func (e *Inject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preserve":
		fallthrough
	case "w3c":
		fallthrough
	case "datadog":
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
		*e = Inject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Inject: %v", v)
	}
}

type Propagation struct {
	// Header names to clear after context extraction. This allows to extract the context from a certain header and then remove it from the request, useful when extraction and injection are performed on different header formats and the original header should not be sent to the upstream. If left empty, no headers are cleared.
	Clear []string `json:"clear,omitempty"`
	// The default header format to use when extractors did not match any format in the incoming headers and `inject` is configured with the value: `preserve`. This can happen when no tracing header was found in the request, or the incoming tracing header formats were not included in `extract`.
	DefaultFormat DefaultFormat `json:"default_format"`
	// Header formats used to extract tracing context from incoming requests. If multiple values are specified, the first one found will be used for extraction. If left empty, Kong will not extract any tracing context information from incoming requests and generate a trace with no parent and a new trace ID.
	Extract []Extract `json:"extract,omitempty"`
	// Header formats used to inject tracing context. The value `preserve` will use the same header format as the incoming request. If multiple values are specified, all of them will be used during injection. If left empty, Kong will not inject any tracing context information in outgoing requests.
	Inject []Inject `json:"inject,omitempty"`
}

func (o *Propagation) GetClear() []string {
	if o == nil {
		return nil
	}
	return o.Clear
}

func (o *Propagation) GetDefaultFormat() DefaultFormat {
	if o == nil {
		return DefaultFormat("")
	}
	return o.DefaultFormat
}

func (o *Propagation) GetExtract() []Extract {
	if o == nil {
		return nil
	}
	return o.Extract
}

func (o *Propagation) GetInject() []Inject {
	if o == nil {
		return nil
	}
	return o.Inject
}

// OpentelemetryPluginConcurrencyLimit - The number of of queue delivery timers. -1 indicates unlimited.
type OpentelemetryPluginConcurrencyLimit int64

const (
	OpentelemetryPluginConcurrencyLimitMinus1 OpentelemetryPluginConcurrencyLimit = -1
	OpentelemetryPluginConcurrencyLimitOne    OpentelemetryPluginConcurrencyLimit = 1
)

func (e OpentelemetryPluginConcurrencyLimit) ToPointer() *OpentelemetryPluginConcurrencyLimit {
	return &e
}
func (e *OpentelemetryPluginConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 1:
		*e = OpentelemetryPluginConcurrencyLimit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OpentelemetryPluginConcurrencyLimit: %v", v)
	}
}

type OpentelemetryPluginQueue struct {
	// The number of of queue delivery timers. -1 indicates unlimited.
	ConcurrencyLimit *OpentelemetryPluginConcurrencyLimit `json:"concurrency_limit,omitempty"`
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

func (o *OpentelemetryPluginQueue) GetConcurrencyLimit() *OpentelemetryPluginConcurrencyLimit {
	if o == nil {
		return nil
	}
	return o.ConcurrencyLimit
}

func (o *OpentelemetryPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *OpentelemetryPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *OpentelemetryPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *OpentelemetryPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *OpentelemetryPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *OpentelemetryPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *OpentelemetryPluginQueue) GetMaxRetryTime() *float64 {
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
	ConnectTimeout *int64      `json:"connect_timeout,omitempty"`
	HeaderType     *HeaderType `json:"header_type,omitempty"`
	// The custom headers to be added in the HTTP request sent to the OTLP server. This setting is useful for adding the authentication headers (token) for the APM backend.
	Headers                      map[string]any `json:"headers,omitempty"`
	HTTPResponseHeaderForTraceid *string        `json:"http_response_header_for_traceid,omitempty"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	LogsEndpoint *string                   `json:"logs_endpoint,omitempty"`
	Propagation  *Propagation              `json:"propagation,omitempty"`
	Queue        *OpentelemetryPluginQueue `json:"queue,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout        *int64         `json:"read_timeout,omitempty"`
	ResourceAttributes map[string]any `json:"resource_attributes,omitempty"`
	// Tracing sampling rate for configuring the probability-based sampler. When set, this value supersedes the global `tracing_sampling_rate` setting from kong.conf.
	SamplingRate *float64 `json:"sampling_rate,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	TracesEndpoint *string `json:"traces_endpoint,omitempty"`
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

func (o *OpentelemetryPluginConfig) GetHeaderType() *HeaderType {
	if o == nil {
		return nil
	}
	return o.HeaderType
}

func (o *OpentelemetryPluginConfig) GetHeaders() map[string]any {
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

func (o *OpentelemetryPluginConfig) GetLogsEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.LogsEndpoint
}

func (o *OpentelemetryPluginConfig) GetPropagation() *Propagation {
	if o == nil {
		return nil
	}
	return o.Propagation
}

func (o *OpentelemetryPluginConfig) GetQueue() *OpentelemetryPluginQueue {
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

func (o *OpentelemetryPluginConfig) GetResourceAttributes() map[string]any {
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

func (o *OpentelemetryPluginConfig) GetTracesEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.TracesEndpoint
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

type OpentelemetryPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *OpentelemetryPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type OpentelemetryPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *OpentelemetryPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type OpentelemetryPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *OpentelemetryPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type OpentelemetryPluginOrdering struct {
	After  *OpentelemetryPluginAfter  `json:"after,omitempty"`
	Before *OpentelemetryPluginBefore `json:"before,omitempty"`
}

func (o *OpentelemetryPluginOrdering) GetAfter() *OpentelemetryPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *OpentelemetryPluginOrdering) GetBefore() *OpentelemetryPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

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

// OpentelemetryPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type OpentelemetryPlugin struct {
	Config OpentelemetryPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *OpentelemetryPluginConsumer      `json:"consumer"`
	ConsumerGroup *OpentelemetryPluginConsumerGroup `json:"consumer_group"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	ID           *string                      `json:"id,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         string                       `const:"opentelemetry" json:"name"`
	Ordering     *OpentelemetryPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []OpentelemetryPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *OpentelemetryPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *OpentelemetryPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
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

func (o *OpentelemetryPlugin) GetConfig() OpentelemetryPluginConfig {
	if o == nil {
		return OpentelemetryPluginConfig{}
	}
	return o.Config
}

func (o *OpentelemetryPlugin) GetConsumer() *OpentelemetryPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *OpentelemetryPlugin) GetConsumerGroup() *OpentelemetryPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *OpentelemetryPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *OpentelemetryPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *OpentelemetryPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OpentelemetryPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *OpentelemetryPlugin) GetName() string {
	return "opentelemetry"
}

func (o *OpentelemetryPlugin) GetOrdering() *OpentelemetryPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *OpentelemetryPlugin) GetProtocols() []OpentelemetryPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
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

func (o *OpentelemetryPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *OpentelemetryPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// OpentelemetryPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type OpentelemetryPluginInput struct {
	Config OpentelemetryPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *OpentelemetryPluginConsumer      `json:"consumer"`
	ConsumerGroup *OpentelemetryPluginConsumerGroup `json:"consumer_group"`
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	ID           *string                      `json:"id,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         string                       `const:"opentelemetry" json:"name"`
	Ordering     *OpentelemetryPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []OpentelemetryPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *OpentelemetryPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *OpentelemetryPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (o OpentelemetryPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OpentelemetryPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OpentelemetryPluginInput) GetConfig() OpentelemetryPluginConfig {
	if o == nil {
		return OpentelemetryPluginConfig{}
	}
	return o.Config
}

func (o *OpentelemetryPluginInput) GetConsumer() *OpentelemetryPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *OpentelemetryPluginInput) GetConsumerGroup() *OpentelemetryPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *OpentelemetryPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *OpentelemetryPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OpentelemetryPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *OpentelemetryPluginInput) GetName() string {
	return "opentelemetry"
}

func (o *OpentelemetryPluginInput) GetOrdering() *OpentelemetryPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *OpentelemetryPluginInput) GetProtocols() []OpentelemetryPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *OpentelemetryPluginInput) GetRoute() *OpentelemetryPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *OpentelemetryPluginInput) GetService() *OpentelemetryPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *OpentelemetryPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
