// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

// ConsumerIdentifier - Authenticated user detail
type ConsumerIdentifier string

const (
	ConsumerIdentifierConsumerID ConsumerIdentifier = "consumer_id"
	ConsumerIdentifierCustomID   ConsumerIdentifier = "custom_id"
	ConsumerIdentifierUsername   ConsumerIdentifier = "username"
)

func (e ConsumerIdentifier) ToPointer() *ConsumerIdentifier {
	return &e
}
func (e *ConsumerIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer_id":
		fallthrough
	case "custom_id":
		fallthrough
	case "username":
		*e = ConsumerIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConsumerIdentifier: %v", v)
	}
}

// Name - Datadog metric’s name
type Name string

const (
	NameKongLatency     Name = "kong_latency"
	NameLatency         Name = "latency"
	NameRequestCount    Name = "request_count"
	NameRequestSize     Name = "request_size"
	NameResponseSize    Name = "response_size"
	NameUpstreamLatency Name = "upstream_latency"
)

func (e Name) ToPointer() *Name {
	return &e
}
func (e *Name) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kong_latency":
		fallthrough
	case "latency":
		fallthrough
	case "request_count":
		fallthrough
	case "request_size":
		fallthrough
	case "response_size":
		fallthrough
	case "upstream_latency":
		*e = Name(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Name: %v", v)
	}
}

// StatType - Determines what sort of event the metric represents
type StatType string

const (
	StatTypeCounter      StatType = "counter"
	StatTypeGauge        StatType = "gauge"
	StatTypeHistogram    StatType = "histogram"
	StatTypeMeter        StatType = "meter"
	StatTypeSet          StatType = "set"
	StatTypeTimer        StatType = "timer"
	StatTypeDistribution StatType = "distribution"
)

func (e StatType) ToPointer() *StatType {
	return &e
}
func (e *StatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "counter":
		fallthrough
	case "gauge":
		fallthrough
	case "histogram":
		fallthrough
	case "meter":
		fallthrough
	case "set":
		fallthrough
	case "timer":
		fallthrough
	case "distribution":
		*e = StatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatType: %v", v)
	}
}

type Metrics struct {
	// Authenticated user detail
	ConsumerIdentifier *ConsumerIdentifier `json:"consumer_identifier,omitempty"`
	// Datadog metric’s name
	Name Name `json:"name"`
	// Sampling rate
	SampleRate *float64 `json:"sample_rate,omitempty"`
	// Determines what sort of event the metric represents
	StatType StatType `json:"stat_type"`
	// List of tags
	Tags []string `json:"tags,omitempty"`
}

func (o *Metrics) GetConsumerIdentifier() *ConsumerIdentifier {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifier
}

func (o *Metrics) GetName() Name {
	if o == nil {
		return Name("")
	}
	return o.Name
}

func (o *Metrics) GetSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *Metrics) GetStatType() StatType {
	if o == nil {
		return StatType("")
	}
	return o.StatType
}

func (o *Metrics) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// ConcurrencyLimit - The number of of queue delivery timers. -1 indicates unlimited.
type ConcurrencyLimit int64

const (
	ConcurrencyLimitMinus1 ConcurrencyLimit = -1
	ConcurrencyLimitOne    ConcurrencyLimit = 1
)

func (e ConcurrencyLimit) ToPointer() *ConcurrencyLimit {
	return &e
}
func (e *ConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 1:
		*e = ConcurrencyLimit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConcurrencyLimit: %v", v)
	}
}

type Queue struct {
	// The number of of queue delivery timers. -1 indicates unlimited.
	ConcurrencyLimit *ConcurrencyLimit `json:"concurrency_limit,omitempty"`
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

func (o *Queue) GetConcurrencyLimit() *ConcurrencyLimit {
	if o == nil {
		return nil
	}
	return o.ConcurrencyLimit
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

type DatadogPluginConfig struct {
	// String to be attached as tag of the consumer.
	ConsumerTag *string `json:"consumer_tag,omitempty"`
	// Optional time in seconds. If `queue_size` > 1, this is the max idle time before sending a log with less than `queue_size` records.
	FlushTimeout *float64 `json:"flush_timeout,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// List of metrics to be logged.
	Metrics []Metrics `json:"metrics,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// String to be attached as a prefix to a metric's name.
	Prefix *string `json:"prefix,omitempty"`
	Queue  *Queue  `json:"queue,omitempty"`
	// Maximum number of log entries to be sent on each message to the upstream server.
	QueueSize *int64 `json:"queue_size,omitempty"`
	// Number of times to retry when sending data to the upstream server.
	RetryCount *int64 `json:"retry_count,omitempty"`
	// String to be attached as the name of the service.
	ServiceNameTag *string `json:"service_name_tag,omitempty"`
	// String to be attached as the tag of the HTTP status.
	StatusTag *string `json:"status_tag,omitempty"`
}

func (o *DatadogPluginConfig) GetConsumerTag() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerTag
}

func (o *DatadogPluginConfig) GetFlushTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushTimeout
}

func (o *DatadogPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *DatadogPluginConfig) GetMetrics() []Metrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *DatadogPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DatadogPluginConfig) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *DatadogPluginConfig) GetQueue() *Queue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *DatadogPluginConfig) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}

func (o *DatadogPluginConfig) GetRetryCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RetryCount
}

func (o *DatadogPluginConfig) GetServiceNameTag() *string {
	if o == nil {
		return nil
	}
	return o.ServiceNameTag
}

func (o *DatadogPluginConfig) GetStatusTag() *string {
	if o == nil {
		return nil
	}
	return o.StatusTag
}

// DatadogPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type DatadogPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *DatadogPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type DatadogPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *DatadogPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type DatadogPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *DatadogPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type DatadogPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *DatadogPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type DatadogPluginOrdering struct {
	After  *DatadogPluginAfter  `json:"after,omitempty"`
	Before *DatadogPluginBefore `json:"before,omitempty"`
}

func (o *DatadogPluginOrdering) GetAfter() *DatadogPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *DatadogPluginOrdering) GetBefore() *DatadogPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type DatadogPluginProtocols string

const (
	DatadogPluginProtocolsGrpc           DatadogPluginProtocols = "grpc"
	DatadogPluginProtocolsGrpcs          DatadogPluginProtocols = "grpcs"
	DatadogPluginProtocolsHTTP           DatadogPluginProtocols = "http"
	DatadogPluginProtocolsHTTPS          DatadogPluginProtocols = "https"
	DatadogPluginProtocolsTCP            DatadogPluginProtocols = "tcp"
	DatadogPluginProtocolsTLS            DatadogPluginProtocols = "tls"
	DatadogPluginProtocolsTLSPassthrough DatadogPluginProtocols = "tls_passthrough"
	DatadogPluginProtocolsUDP            DatadogPluginProtocols = "udp"
	DatadogPluginProtocolsWs             DatadogPluginProtocols = "ws"
	DatadogPluginProtocolsWss            DatadogPluginProtocols = "wss"
)

func (e DatadogPluginProtocols) ToPointer() *DatadogPluginProtocols {
	return &e
}
func (e *DatadogPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = DatadogPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatadogPluginProtocols: %v", v)
	}
}

// DatadogPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type DatadogPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *DatadogPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// DatadogPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type DatadogPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *DatadogPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// DatadogPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type DatadogPlugin struct {
	Config DatadogPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *DatadogPluginConsumer      `json:"consumer"`
	ConsumerGroup *DatadogPluginConsumerGroup `json:"consumer_group"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                  `json:"enabled,omitempty"`
	ID           *string                `json:"id,omitempty"`
	InstanceName *string                `json:"instance_name,omitempty"`
	name         string                 `const:"datadog" json:"name"`
	Ordering     *DatadogPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []DatadogPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *DatadogPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *DatadogPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (d DatadogPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DatadogPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DatadogPlugin) GetConfig() DatadogPluginConfig {
	if o == nil {
		return DatadogPluginConfig{}
	}
	return o.Config
}

func (o *DatadogPlugin) GetConsumer() *DatadogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *DatadogPlugin) GetConsumerGroup() *DatadogPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *DatadogPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *DatadogPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *DatadogPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DatadogPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *DatadogPlugin) GetName() string {
	return "datadog"
}

func (o *DatadogPlugin) GetOrdering() *DatadogPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *DatadogPlugin) GetProtocols() []DatadogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *DatadogPlugin) GetRoute() *DatadogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *DatadogPlugin) GetService() *DatadogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *DatadogPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *DatadogPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// DatadogPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type DatadogPluginInput struct {
	Config DatadogPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *DatadogPluginConsumer      `json:"consumer"`
	ConsumerGroup *DatadogPluginConsumerGroup `json:"consumer_group"`
	// Whether the plugin is applied.
	Enabled      *bool                  `json:"enabled,omitempty"`
	ID           *string                `json:"id,omitempty"`
	InstanceName *string                `json:"instance_name,omitempty"`
	name         string                 `const:"datadog" json:"name"`
	Ordering     *DatadogPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []DatadogPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *DatadogPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *DatadogPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (d DatadogPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DatadogPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DatadogPluginInput) GetConfig() DatadogPluginConfig {
	if o == nil {
		return DatadogPluginConfig{}
	}
	return o.Config
}

func (o *DatadogPluginInput) GetConsumer() *DatadogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *DatadogPluginInput) GetConsumerGroup() *DatadogPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *DatadogPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *DatadogPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DatadogPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *DatadogPluginInput) GetName() string {
	return "datadog"
}

func (o *DatadogPluginInput) GetOrdering() *DatadogPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *DatadogPluginInput) GetProtocols() []DatadogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *DatadogPluginInput) GetRoute() *DatadogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *DatadogPluginInput) GetService() *DatadogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *DatadogPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
