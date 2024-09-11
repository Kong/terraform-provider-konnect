// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateStatsdPluginConsumerIdentifierDefault string

const (
	CreateStatsdPluginConsumerIdentifierDefaultConsumerID CreateStatsdPluginConsumerIdentifierDefault = "consumer_id"
	CreateStatsdPluginConsumerIdentifierDefaultCustomID   CreateStatsdPluginConsumerIdentifierDefault = "custom_id"
	CreateStatsdPluginConsumerIdentifierDefaultUsername   CreateStatsdPluginConsumerIdentifierDefault = "username"
)

func (e CreateStatsdPluginConsumerIdentifierDefault) ToPointer() *CreateStatsdPluginConsumerIdentifierDefault {
	return &e
}
func (e *CreateStatsdPluginConsumerIdentifierDefault) UnmarshalJSON(data []byte) error {
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
		*e = CreateStatsdPluginConsumerIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginConsumerIdentifierDefault: %v", v)
	}
}

// CreateStatsdPluginConsumerIdentifier - Authenticated user detail.
type CreateStatsdPluginConsumerIdentifier string

const (
	CreateStatsdPluginConsumerIdentifierConsumerID CreateStatsdPluginConsumerIdentifier = "consumer_id"
	CreateStatsdPluginConsumerIdentifierCustomID   CreateStatsdPluginConsumerIdentifier = "custom_id"
	CreateStatsdPluginConsumerIdentifierUsername   CreateStatsdPluginConsumerIdentifier = "username"
)

func (e CreateStatsdPluginConsumerIdentifier) ToPointer() *CreateStatsdPluginConsumerIdentifier {
	return &e
}
func (e *CreateStatsdPluginConsumerIdentifier) UnmarshalJSON(data []byte) error {
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
		*e = CreateStatsdPluginConsumerIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginConsumerIdentifier: %v", v)
	}
}

// CreateStatsdPluginName - StatsD metric’s name.
type CreateStatsdPluginName string

const (
	CreateStatsdPluginNameKongLatency                CreateStatsdPluginName = "kong_latency"
	CreateStatsdPluginNameLatency                    CreateStatsdPluginName = "latency"
	CreateStatsdPluginNameRequestCount               CreateStatsdPluginName = "request_count"
	CreateStatsdPluginNameRequestPerUser             CreateStatsdPluginName = "request_per_user"
	CreateStatsdPluginNameRequestSize                CreateStatsdPluginName = "request_size"
	CreateStatsdPluginNameResponseSize               CreateStatsdPluginName = "response_size"
	CreateStatsdPluginNameStatusCount                CreateStatsdPluginName = "status_count"
	CreateStatsdPluginNameStatusCountPerUser         CreateStatsdPluginName = "status_count_per_user"
	CreateStatsdPluginNameUniqueUsers                CreateStatsdPluginName = "unique_users"
	CreateStatsdPluginNameUpstreamLatency            CreateStatsdPluginName = "upstream_latency"
	CreateStatsdPluginNameStatusCountPerWorkspace    CreateStatsdPluginName = "status_count_per_workspace"
	CreateStatsdPluginNameStatusCountPerUserPerRoute CreateStatsdPluginName = "status_count_per_user_per_route"
	CreateStatsdPluginNameShdictUsage                CreateStatsdPluginName = "shdict_usage"
	CreateStatsdPluginNameCacheDatastoreHitsTotal    CreateStatsdPluginName = "cache_datastore_hits_total"
	CreateStatsdPluginNameCacheDatastoreMissesTotal  CreateStatsdPluginName = "cache_datastore_misses_total"
)

func (e CreateStatsdPluginName) ToPointer() *CreateStatsdPluginName {
	return &e
}
func (e *CreateStatsdPluginName) UnmarshalJSON(data []byte) error {
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
	case "request_per_user":
		fallthrough
	case "request_size":
		fallthrough
	case "response_size":
		fallthrough
	case "status_count":
		fallthrough
	case "status_count_per_user":
		fallthrough
	case "unique_users":
		fallthrough
	case "upstream_latency":
		fallthrough
	case "status_count_per_workspace":
		fallthrough
	case "status_count_per_user_per_route":
		fallthrough
	case "shdict_usage":
		fallthrough
	case "cache_datastore_hits_total":
		fallthrough
	case "cache_datastore_misses_total":
		*e = CreateStatsdPluginName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginName: %v", v)
	}
}

// CreateStatsdPluginServiceIdentifier - Service detail.
type CreateStatsdPluginServiceIdentifier string

const (
	CreateStatsdPluginServiceIdentifierServiceID         CreateStatsdPluginServiceIdentifier = "service_id"
	CreateStatsdPluginServiceIdentifierServiceName       CreateStatsdPluginServiceIdentifier = "service_name"
	CreateStatsdPluginServiceIdentifierServiceHost       CreateStatsdPluginServiceIdentifier = "service_host"
	CreateStatsdPluginServiceIdentifierServiceNameOrHost CreateStatsdPluginServiceIdentifier = "service_name_or_host"
)

func (e CreateStatsdPluginServiceIdentifier) ToPointer() *CreateStatsdPluginServiceIdentifier {
	return &e
}
func (e *CreateStatsdPluginServiceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_host":
		fallthrough
	case "service_name_or_host":
		*e = CreateStatsdPluginServiceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginServiceIdentifier: %v", v)
	}
}

// CreateStatsdPluginStatType - Determines what sort of event a metric represents.
type CreateStatsdPluginStatType string

const (
	CreateStatsdPluginStatTypeCounter   CreateStatsdPluginStatType = "counter"
	CreateStatsdPluginStatTypeGauge     CreateStatsdPluginStatType = "gauge"
	CreateStatsdPluginStatTypeHistogram CreateStatsdPluginStatType = "histogram"
	CreateStatsdPluginStatTypeMeter     CreateStatsdPluginStatType = "meter"
	CreateStatsdPluginStatTypeSet       CreateStatsdPluginStatType = "set"
	CreateStatsdPluginStatTypeTimer     CreateStatsdPluginStatType = "timer"
)

func (e CreateStatsdPluginStatType) ToPointer() *CreateStatsdPluginStatType {
	return &e
}
func (e *CreateStatsdPluginStatType) UnmarshalJSON(data []byte) error {
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
		*e = CreateStatsdPluginStatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginStatType: %v", v)
	}
}

// CreateStatsdPluginWorkspaceIdentifier - Workspace detail.
type CreateStatsdPluginWorkspaceIdentifier string

const (
	CreateStatsdPluginWorkspaceIdentifierWorkspaceID   CreateStatsdPluginWorkspaceIdentifier = "workspace_id"
	CreateStatsdPluginWorkspaceIdentifierWorkspaceName CreateStatsdPluginWorkspaceIdentifier = "workspace_name"
)

func (e CreateStatsdPluginWorkspaceIdentifier) ToPointer() *CreateStatsdPluginWorkspaceIdentifier {
	return &e
}
func (e *CreateStatsdPluginWorkspaceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = CreateStatsdPluginWorkspaceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginWorkspaceIdentifier: %v", v)
	}
}

type CreateStatsdPluginMetrics struct {
	// Authenticated user detail.
	ConsumerIdentifier *CreateStatsdPluginConsumerIdentifier `json:"consumer_identifier,omitempty"`
	// StatsD metric’s name.
	Name CreateStatsdPluginName `json:"name"`
	// Sampling rate
	SampleRate *float64 `json:"sample_rate,omitempty"`
	// Service detail.
	ServiceIdentifier *CreateStatsdPluginServiceIdentifier `json:"service_identifier,omitempty"`
	// Determines what sort of event a metric represents.
	StatType CreateStatsdPluginStatType `json:"stat_type"`
	// Workspace detail.
	WorkspaceIdentifier *CreateStatsdPluginWorkspaceIdentifier `json:"workspace_identifier,omitempty"`
}

func (o *CreateStatsdPluginMetrics) GetConsumerIdentifier() *CreateStatsdPluginConsumerIdentifier {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifier
}

func (o *CreateStatsdPluginMetrics) GetName() CreateStatsdPluginName {
	if o == nil {
		return CreateStatsdPluginName("")
	}
	return o.Name
}

func (o *CreateStatsdPluginMetrics) GetSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *CreateStatsdPluginMetrics) GetServiceIdentifier() *CreateStatsdPluginServiceIdentifier {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifier
}

func (o *CreateStatsdPluginMetrics) GetStatType() CreateStatsdPluginStatType {
	if o == nil {
		return CreateStatsdPluginStatType("")
	}
	return o.StatType
}

func (o *CreateStatsdPluginMetrics) GetWorkspaceIdentifier() *CreateStatsdPluginWorkspaceIdentifier {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifier
}

// CreateStatsdPluginConcurrencyLimit - The number of of queue delivery timers. -1 indicates unlimited.
type CreateStatsdPluginConcurrencyLimit int64

const (
	CreateStatsdPluginConcurrencyLimitMinus1 CreateStatsdPluginConcurrencyLimit = -1
	CreateStatsdPluginConcurrencyLimitOne    CreateStatsdPluginConcurrencyLimit = 1
)

func (e CreateStatsdPluginConcurrencyLimit) ToPointer() *CreateStatsdPluginConcurrencyLimit {
	return &e
}
func (e *CreateStatsdPluginConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 1:
		*e = CreateStatsdPluginConcurrencyLimit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginConcurrencyLimit: %v", v)
	}
}

type CreateStatsdPluginQueue struct {
	// The number of of queue delivery timers. -1 indicates unlimited.
	ConcurrencyLimit *CreateStatsdPluginConcurrencyLimit `json:"concurrency_limit,omitempty"`
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

func (o *CreateStatsdPluginQueue) GetConcurrencyLimit() *CreateStatsdPluginConcurrencyLimit {
	if o == nil {
		return nil
	}
	return o.ConcurrencyLimit
}

func (o *CreateStatsdPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *CreateStatsdPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *CreateStatsdPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *CreateStatsdPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *CreateStatsdPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *CreateStatsdPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *CreateStatsdPluginQueue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

type CreateStatsdPluginServiceIdentifierDefault string

const (
	CreateStatsdPluginServiceIdentifierDefaultServiceID         CreateStatsdPluginServiceIdentifierDefault = "service_id"
	CreateStatsdPluginServiceIdentifierDefaultServiceName       CreateStatsdPluginServiceIdentifierDefault = "service_name"
	CreateStatsdPluginServiceIdentifierDefaultServiceHost       CreateStatsdPluginServiceIdentifierDefault = "service_host"
	CreateStatsdPluginServiceIdentifierDefaultServiceNameOrHost CreateStatsdPluginServiceIdentifierDefault = "service_name_or_host"
)

func (e CreateStatsdPluginServiceIdentifierDefault) ToPointer() *CreateStatsdPluginServiceIdentifierDefault {
	return &e
}
func (e *CreateStatsdPluginServiceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_host":
		fallthrough
	case "service_name_or_host":
		*e = CreateStatsdPluginServiceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginServiceIdentifierDefault: %v", v)
	}
}

type CreateStatsdPluginTagStyle string

const (
	CreateStatsdPluginTagStyleDogstatsd CreateStatsdPluginTagStyle = "dogstatsd"
	CreateStatsdPluginTagStyleInfluxdb  CreateStatsdPluginTagStyle = "influxdb"
	CreateStatsdPluginTagStyleLibrato   CreateStatsdPluginTagStyle = "librato"
	CreateStatsdPluginTagStyleSignalfx  CreateStatsdPluginTagStyle = "signalfx"
)

func (e CreateStatsdPluginTagStyle) ToPointer() *CreateStatsdPluginTagStyle {
	return &e
}
func (e *CreateStatsdPluginTagStyle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dogstatsd":
		fallthrough
	case "influxdb":
		fallthrough
	case "librato":
		fallthrough
	case "signalfx":
		*e = CreateStatsdPluginTagStyle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginTagStyle: %v", v)
	}
}

type CreateStatsdPluginWorkspaceIdentifierDefault string

const (
	CreateStatsdPluginWorkspaceIdentifierDefaultWorkspaceID   CreateStatsdPluginWorkspaceIdentifierDefault = "workspace_id"
	CreateStatsdPluginWorkspaceIdentifierDefaultWorkspaceName CreateStatsdPluginWorkspaceIdentifierDefault = "workspace_name"
)

func (e CreateStatsdPluginWorkspaceIdentifierDefault) ToPointer() *CreateStatsdPluginWorkspaceIdentifierDefault {
	return &e
}
func (e *CreateStatsdPluginWorkspaceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = CreateStatsdPluginWorkspaceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginWorkspaceIdentifierDefault: %v", v)
	}
}

type CreateStatsdPluginConfig struct {
	// List of status code ranges that are allowed to be logged in metrics.
	AllowStatusCodes          []string                                     `json:"allow_status_codes,omitempty"`
	ConsumerIdentifierDefault *CreateStatsdPluginConsumerIdentifierDefault `json:"consumer_identifier_default,omitempty"`
	FlushTimeout              *float64                                     `json:"flush_timeout,omitempty"`
	// The IP address or hostname of StatsD server to send data to.
	Host             *string `json:"host,omitempty"`
	HostnameInPrefix *bool   `json:"hostname_in_prefix,omitempty"`
	// List of metrics to be logged.
	Metrics []CreateStatsdPluginMetrics `json:"metrics,omitempty"`
	// The port of StatsD server to send data to.
	Port *int64 `json:"port,omitempty"`
	// String to prefix to each metric's name.
	Prefix                     *string                                       `json:"prefix,omitempty"`
	Queue                      *CreateStatsdPluginQueue                      `json:"queue,omitempty"`
	QueueSize                  *int64                                        `json:"queue_size,omitempty"`
	RetryCount                 *int64                                        `json:"retry_count,omitempty"`
	ServiceIdentifierDefault   *CreateStatsdPluginServiceIdentifierDefault   `json:"service_identifier_default,omitempty"`
	TagStyle                   *CreateStatsdPluginTagStyle                   `json:"tag_style,omitempty"`
	UDPPacketSize              *float64                                      `json:"udp_packet_size,omitempty"`
	UseTCP                     *bool                                         `json:"use_tcp,omitempty"`
	WorkspaceIdentifierDefault *CreateStatsdPluginWorkspaceIdentifierDefault `json:"workspace_identifier_default,omitempty"`
}

func (o *CreateStatsdPluginConfig) GetAllowStatusCodes() []string {
	if o == nil {
		return nil
	}
	return o.AllowStatusCodes
}

func (o *CreateStatsdPluginConfig) GetConsumerIdentifierDefault() *CreateStatsdPluginConsumerIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifierDefault
}

func (o *CreateStatsdPluginConfig) GetFlushTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushTimeout
}

func (o *CreateStatsdPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateStatsdPluginConfig) GetHostnameInPrefix() *bool {
	if o == nil {
		return nil
	}
	return o.HostnameInPrefix
}

func (o *CreateStatsdPluginConfig) GetMetrics() []CreateStatsdPluginMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *CreateStatsdPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateStatsdPluginConfig) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *CreateStatsdPluginConfig) GetQueue() *CreateStatsdPluginQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *CreateStatsdPluginConfig) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}

func (o *CreateStatsdPluginConfig) GetRetryCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RetryCount
}

func (o *CreateStatsdPluginConfig) GetServiceIdentifierDefault() *CreateStatsdPluginServiceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifierDefault
}

func (o *CreateStatsdPluginConfig) GetTagStyle() *CreateStatsdPluginTagStyle {
	if o == nil {
		return nil
	}
	return o.TagStyle
}

func (o *CreateStatsdPluginConfig) GetUDPPacketSize() *float64 {
	if o == nil {
		return nil
	}
	return o.UDPPacketSize
}

func (o *CreateStatsdPluginConfig) GetUseTCP() *bool {
	if o == nil {
		return nil
	}
	return o.UseTCP
}

func (o *CreateStatsdPluginConfig) GetWorkspaceIdentifierDefault() *CreateStatsdPluginWorkspaceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifierDefault
}

type CreateStatsdPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateStatsdPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateStatsdPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateStatsdPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateStatsdPluginOrdering struct {
	After  *CreateStatsdPluginAfter  `json:"after,omitempty"`
	Before *CreateStatsdPluginBefore `json:"before,omitempty"`
}

func (o *CreateStatsdPluginOrdering) GetAfter() *CreateStatsdPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateStatsdPluginOrdering) GetBefore() *CreateStatsdPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateStatsdPluginProtocols string

const (
	CreateStatsdPluginProtocolsGrpc           CreateStatsdPluginProtocols = "grpc"
	CreateStatsdPluginProtocolsGrpcs          CreateStatsdPluginProtocols = "grpcs"
	CreateStatsdPluginProtocolsHTTP           CreateStatsdPluginProtocols = "http"
	CreateStatsdPluginProtocolsHTTPS          CreateStatsdPluginProtocols = "https"
	CreateStatsdPluginProtocolsTCP            CreateStatsdPluginProtocols = "tcp"
	CreateStatsdPluginProtocolsTLS            CreateStatsdPluginProtocols = "tls"
	CreateStatsdPluginProtocolsTLSPassthrough CreateStatsdPluginProtocols = "tls_passthrough"
	CreateStatsdPluginProtocolsUDP            CreateStatsdPluginProtocols = "udp"
	CreateStatsdPluginProtocolsWs             CreateStatsdPluginProtocols = "ws"
	CreateStatsdPluginProtocolsWss            CreateStatsdPluginProtocols = "wss"
)

func (e CreateStatsdPluginProtocols) ToPointer() *CreateStatsdPluginProtocols {
	return &e
}
func (e *CreateStatsdPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateStatsdPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdPluginProtocols: %v", v)
	}
}

// CreateStatsdPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateStatsdPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateStatsdPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateStatsdPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateStatsdPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateStatsdPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateStatsdPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateStatsdPlugin struct {
	Config *CreateStatsdPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                       `json:"enabled,omitempty"`
	InstanceName *string                     `json:"instance_name,omitempty"`
	name         *string                     `const:"statsd" json:"name,omitempty"`
	Ordering     *CreateStatsdPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateStatsdPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateStatsdPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateStatsdPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateStatsdPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateStatsdPluginService `json:"service,omitempty"`
}

func (c CreateStatsdPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateStatsdPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateStatsdPlugin) GetConfig() *CreateStatsdPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateStatsdPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateStatsdPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateStatsdPlugin) GetName() *string {
	return types.String("statsd")
}

func (o *CreateStatsdPlugin) GetOrdering() *CreateStatsdPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateStatsdPlugin) GetProtocols() []CreateStatsdPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateStatsdPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateStatsdPlugin) GetConsumer() *CreateStatsdPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateStatsdPlugin) GetConsumerGroup() *CreateStatsdPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateStatsdPlugin) GetRoute() *CreateStatsdPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateStatsdPlugin) GetService() *CreateStatsdPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
