// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

// StatsdAdvancedPluginConsumerIdentifierDefault - The default consumer identifier for metrics. This will take effect when a metric's consumer identifier is omitted. Allowed values are `custom_id`, `consumer_id`, `username`.
type StatsdAdvancedPluginConsumerIdentifierDefault string

const (
	StatsdAdvancedPluginConsumerIdentifierDefaultConsumerID StatsdAdvancedPluginConsumerIdentifierDefault = "consumer_id"
	StatsdAdvancedPluginConsumerIdentifierDefaultCustomID   StatsdAdvancedPluginConsumerIdentifierDefault = "custom_id"
	StatsdAdvancedPluginConsumerIdentifierDefaultUsername   StatsdAdvancedPluginConsumerIdentifierDefault = "username"
)

func (e StatsdAdvancedPluginConsumerIdentifierDefault) ToPointer() *StatsdAdvancedPluginConsumerIdentifierDefault {
	return &e
}
func (e *StatsdAdvancedPluginConsumerIdentifierDefault) UnmarshalJSON(data []byte) error {
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
		*e = StatsdAdvancedPluginConsumerIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginConsumerIdentifierDefault: %v", v)
	}
}

type StatsdAdvancedPluginConsumerIdentifier string

const (
	StatsdAdvancedPluginConsumerIdentifierConsumerID StatsdAdvancedPluginConsumerIdentifier = "consumer_id"
	StatsdAdvancedPluginConsumerIdentifierCustomID   StatsdAdvancedPluginConsumerIdentifier = "custom_id"
	StatsdAdvancedPluginConsumerIdentifierUsername   StatsdAdvancedPluginConsumerIdentifier = "username"
)

func (e StatsdAdvancedPluginConsumerIdentifier) ToPointer() *StatsdAdvancedPluginConsumerIdentifier {
	return &e
}
func (e *StatsdAdvancedPluginConsumerIdentifier) UnmarshalJSON(data []byte) error {
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
		*e = StatsdAdvancedPluginConsumerIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginConsumerIdentifier: %v", v)
	}
}

type StatsdAdvancedPluginName string

const (
	StatsdAdvancedPluginNameCacheDatastoreHitsTotal    StatsdAdvancedPluginName = "cache_datastore_hits_total"
	StatsdAdvancedPluginNameCacheDatastoreMissesTotal  StatsdAdvancedPluginName = "cache_datastore_misses_total"
	StatsdAdvancedPluginNameKongLatency                StatsdAdvancedPluginName = "kong_latency"
	StatsdAdvancedPluginNameLatency                    StatsdAdvancedPluginName = "latency"
	StatsdAdvancedPluginNameRequestCount               StatsdAdvancedPluginName = "request_count"
	StatsdAdvancedPluginNameRequestPerUser             StatsdAdvancedPluginName = "request_per_user"
	StatsdAdvancedPluginNameRequestSize                StatsdAdvancedPluginName = "request_size"
	StatsdAdvancedPluginNameResponseSize               StatsdAdvancedPluginName = "response_size"
	StatsdAdvancedPluginNameShdictUsage                StatsdAdvancedPluginName = "shdict_usage"
	StatsdAdvancedPluginNameStatusCount                StatsdAdvancedPluginName = "status_count"
	StatsdAdvancedPluginNameStatusCountPerUser         StatsdAdvancedPluginName = "status_count_per_user"
	StatsdAdvancedPluginNameStatusCountPerUserPerRoute StatsdAdvancedPluginName = "status_count_per_user_per_route"
	StatsdAdvancedPluginNameStatusCountPerWorkspace    StatsdAdvancedPluginName = "status_count_per_workspace"
	StatsdAdvancedPluginNameUniqueUsers                StatsdAdvancedPluginName = "unique_users"
	StatsdAdvancedPluginNameUpstreamLatency            StatsdAdvancedPluginName = "upstream_latency"
)

func (e StatsdAdvancedPluginName) ToPointer() *StatsdAdvancedPluginName {
	return &e
}
func (e *StatsdAdvancedPluginName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cache_datastore_hits_total":
		fallthrough
	case "cache_datastore_misses_total":
		fallthrough
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
	case "shdict_usage":
		fallthrough
	case "status_count":
		fallthrough
	case "status_count_per_user":
		fallthrough
	case "status_count_per_user_per_route":
		fallthrough
	case "status_count_per_workspace":
		fallthrough
	case "unique_users":
		fallthrough
	case "upstream_latency":
		*e = StatsdAdvancedPluginName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginName: %v", v)
	}
}

type StatsdAdvancedPluginServiceIdentifier string

const (
	StatsdAdvancedPluginServiceIdentifierServiceHost       StatsdAdvancedPluginServiceIdentifier = "service_host"
	StatsdAdvancedPluginServiceIdentifierServiceID         StatsdAdvancedPluginServiceIdentifier = "service_id"
	StatsdAdvancedPluginServiceIdentifierServiceName       StatsdAdvancedPluginServiceIdentifier = "service_name"
	StatsdAdvancedPluginServiceIdentifierServiceNameOrHost StatsdAdvancedPluginServiceIdentifier = "service_name_or_host"
)

func (e StatsdAdvancedPluginServiceIdentifier) ToPointer() *StatsdAdvancedPluginServiceIdentifier {
	return &e
}
func (e *StatsdAdvancedPluginServiceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_host":
		fallthrough
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_name_or_host":
		*e = StatsdAdvancedPluginServiceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginServiceIdentifier: %v", v)
	}
}

type StatsdAdvancedPluginStatType string

const (
	StatsdAdvancedPluginStatTypeCounter   StatsdAdvancedPluginStatType = "counter"
	StatsdAdvancedPluginStatTypeGauge     StatsdAdvancedPluginStatType = "gauge"
	StatsdAdvancedPluginStatTypeHistogram StatsdAdvancedPluginStatType = "histogram"
	StatsdAdvancedPluginStatTypeMeter     StatsdAdvancedPluginStatType = "meter"
	StatsdAdvancedPluginStatTypeSet       StatsdAdvancedPluginStatType = "set"
	StatsdAdvancedPluginStatTypeTimer     StatsdAdvancedPluginStatType = "timer"
)

func (e StatsdAdvancedPluginStatType) ToPointer() *StatsdAdvancedPluginStatType {
	return &e
}
func (e *StatsdAdvancedPluginStatType) UnmarshalJSON(data []byte) error {
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
		*e = StatsdAdvancedPluginStatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginStatType: %v", v)
	}
}

type StatsdAdvancedPluginWorkspaceIdentifier string

const (
	StatsdAdvancedPluginWorkspaceIdentifierWorkspaceID   StatsdAdvancedPluginWorkspaceIdentifier = "workspace_id"
	StatsdAdvancedPluginWorkspaceIdentifierWorkspaceName StatsdAdvancedPluginWorkspaceIdentifier = "workspace_name"
)

func (e StatsdAdvancedPluginWorkspaceIdentifier) ToPointer() *StatsdAdvancedPluginWorkspaceIdentifier {
	return &e
}
func (e *StatsdAdvancedPluginWorkspaceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = StatsdAdvancedPluginWorkspaceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginWorkspaceIdentifier: %v", v)
	}
}

type StatsdAdvancedPluginMetrics struct {
	ConsumerIdentifier  *StatsdAdvancedPluginConsumerIdentifier  `json:"consumer_identifier,omitempty"`
	Name                StatsdAdvancedPluginName                 `json:"name"`
	SampleRate          *float64                                 `json:"sample_rate,omitempty"`
	ServiceIdentifier   *StatsdAdvancedPluginServiceIdentifier   `json:"service_identifier,omitempty"`
	StatType            StatsdAdvancedPluginStatType             `json:"stat_type"`
	WorkspaceIdentifier *StatsdAdvancedPluginWorkspaceIdentifier `json:"workspace_identifier,omitempty"`
}

func (o *StatsdAdvancedPluginMetrics) GetConsumerIdentifier() *StatsdAdvancedPluginConsumerIdentifier {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifier
}

func (o *StatsdAdvancedPluginMetrics) GetName() StatsdAdvancedPluginName {
	if o == nil {
		return StatsdAdvancedPluginName("")
	}
	return o.Name
}

func (o *StatsdAdvancedPluginMetrics) GetSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *StatsdAdvancedPluginMetrics) GetServiceIdentifier() *StatsdAdvancedPluginServiceIdentifier {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifier
}

func (o *StatsdAdvancedPluginMetrics) GetStatType() StatsdAdvancedPluginStatType {
	if o == nil {
		return StatsdAdvancedPluginStatType("")
	}
	return o.StatType
}

func (o *StatsdAdvancedPluginMetrics) GetWorkspaceIdentifier() *StatsdAdvancedPluginWorkspaceIdentifier {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifier
}

// StatsdAdvancedPluginConcurrencyLimit - The number of of queue delivery timers. -1 indicates unlimited.
type StatsdAdvancedPluginConcurrencyLimit int64

const (
	StatsdAdvancedPluginConcurrencyLimitMinus1 StatsdAdvancedPluginConcurrencyLimit = -1
	StatsdAdvancedPluginConcurrencyLimitOne    StatsdAdvancedPluginConcurrencyLimit = 1
)

func (e StatsdAdvancedPluginConcurrencyLimit) ToPointer() *StatsdAdvancedPluginConcurrencyLimit {
	return &e
}
func (e *StatsdAdvancedPluginConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 1:
		*e = StatsdAdvancedPluginConcurrencyLimit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginConcurrencyLimit: %v", v)
	}
}

type StatsdAdvancedPluginQueue struct {
	// The number of of queue delivery timers. -1 indicates unlimited.
	ConcurrencyLimit *StatsdAdvancedPluginConcurrencyLimit `json:"concurrency_limit,omitempty"`
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

func (o *StatsdAdvancedPluginQueue) GetConcurrencyLimit() *StatsdAdvancedPluginConcurrencyLimit {
	if o == nil {
		return nil
	}
	return o.ConcurrencyLimit
}

func (o *StatsdAdvancedPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *StatsdAdvancedPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *StatsdAdvancedPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *StatsdAdvancedPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *StatsdAdvancedPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *StatsdAdvancedPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *StatsdAdvancedPluginQueue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

// StatsdAdvancedPluginServiceIdentifierDefault - The default service identifier for metrics. This will take effect when a metric's service identifier is omitted. Allowed values are `service_name_or_host`, `service_id`, `service_name`, `service_host`.
type StatsdAdvancedPluginServiceIdentifierDefault string

const (
	StatsdAdvancedPluginServiceIdentifierDefaultServiceHost       StatsdAdvancedPluginServiceIdentifierDefault = "service_host"
	StatsdAdvancedPluginServiceIdentifierDefaultServiceID         StatsdAdvancedPluginServiceIdentifierDefault = "service_id"
	StatsdAdvancedPluginServiceIdentifierDefaultServiceName       StatsdAdvancedPluginServiceIdentifierDefault = "service_name"
	StatsdAdvancedPluginServiceIdentifierDefaultServiceNameOrHost StatsdAdvancedPluginServiceIdentifierDefault = "service_name_or_host"
)

func (e StatsdAdvancedPluginServiceIdentifierDefault) ToPointer() *StatsdAdvancedPluginServiceIdentifierDefault {
	return &e
}
func (e *StatsdAdvancedPluginServiceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_host":
		fallthrough
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_name_or_host":
		*e = StatsdAdvancedPluginServiceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginServiceIdentifierDefault: %v", v)
	}
}

// StatsdAdvancedPluginWorkspaceIdentifierDefault - The default workspace identifier for metrics. This will take effect when a metric's workspace identifier is omitted. Allowed values are `workspace_id`, `workspace_name`.
type StatsdAdvancedPluginWorkspaceIdentifierDefault string

const (
	StatsdAdvancedPluginWorkspaceIdentifierDefaultWorkspaceID   StatsdAdvancedPluginWorkspaceIdentifierDefault = "workspace_id"
	StatsdAdvancedPluginWorkspaceIdentifierDefaultWorkspaceName StatsdAdvancedPluginWorkspaceIdentifierDefault = "workspace_name"
)

func (e StatsdAdvancedPluginWorkspaceIdentifierDefault) ToPointer() *StatsdAdvancedPluginWorkspaceIdentifierDefault {
	return &e
}
func (e *StatsdAdvancedPluginWorkspaceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = StatsdAdvancedPluginWorkspaceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginWorkspaceIdentifierDefault: %v", v)
	}
}

type StatsdAdvancedPluginConfig struct {
	// List of status code ranges that are allowed to be logged in metrics.
	AllowStatusCodes []string `json:"allow_status_codes,omitempty"`
	// The default consumer identifier for metrics. This will take effect when a metric's consumer identifier is omitted. Allowed values are `custom_id`, `consumer_id`, `username`.
	ConsumerIdentifierDefault *StatsdAdvancedPluginConsumerIdentifierDefault `json:"consumer_identifier_default,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Include the `hostname` in the `prefix` for each metric name.
	HostnameInPrefix *bool `json:"hostname_in_prefix,omitempty"`
	// List of Metrics to be logged.
	Metrics []StatsdAdvancedPluginMetrics `json:"metrics,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// String to prefix to each metric's name.
	Prefix *string                    `json:"prefix,omitempty"`
	Queue  *StatsdAdvancedPluginQueue `json:"queue,omitempty"`
	// The default service identifier for metrics. This will take effect when a metric's service identifier is omitted. Allowed values are `service_name_or_host`, `service_id`, `service_name`, `service_host`.
	ServiceIdentifierDefault *StatsdAdvancedPluginServiceIdentifierDefault `json:"service_identifier_default,omitempty"`
	// Combine UDP packet up to the size configured. If zero (0), don't combine the UDP packet. Must be a number between 0 and 65507 (inclusive).
	UDPPacketSize *float64 `json:"udp_packet_size,omitempty"`
	// Use TCP instead of UDP.
	UseTCP *bool `json:"use_tcp,omitempty"`
	// The default workspace identifier for metrics. This will take effect when a metric's workspace identifier is omitted. Allowed values are `workspace_id`, `workspace_name`.
	WorkspaceIdentifierDefault *StatsdAdvancedPluginWorkspaceIdentifierDefault `json:"workspace_identifier_default,omitempty"`
}

func (o *StatsdAdvancedPluginConfig) GetAllowStatusCodes() []string {
	if o == nil {
		return nil
	}
	return o.AllowStatusCodes
}

func (o *StatsdAdvancedPluginConfig) GetConsumerIdentifierDefault() *StatsdAdvancedPluginConsumerIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifierDefault
}

func (o *StatsdAdvancedPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *StatsdAdvancedPluginConfig) GetHostnameInPrefix() *bool {
	if o == nil {
		return nil
	}
	return o.HostnameInPrefix
}

func (o *StatsdAdvancedPluginConfig) GetMetrics() []StatsdAdvancedPluginMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *StatsdAdvancedPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *StatsdAdvancedPluginConfig) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *StatsdAdvancedPluginConfig) GetQueue() *StatsdAdvancedPluginQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *StatsdAdvancedPluginConfig) GetServiceIdentifierDefault() *StatsdAdvancedPluginServiceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifierDefault
}

func (o *StatsdAdvancedPluginConfig) GetUDPPacketSize() *float64 {
	if o == nil {
		return nil
	}
	return o.UDPPacketSize
}

func (o *StatsdAdvancedPluginConfig) GetUseTCP() *bool {
	if o == nil {
		return nil
	}
	return o.UseTCP
}

func (o *StatsdAdvancedPluginConfig) GetWorkspaceIdentifierDefault() *StatsdAdvancedPluginWorkspaceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifierDefault
}

// StatsdAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type StatsdAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// StatsdAdvancedPluginProtocols - A string representing a protocol, such as HTTP or HTTPS.
type StatsdAdvancedPluginProtocols string

const (
	StatsdAdvancedPluginProtocolsGrpc           StatsdAdvancedPluginProtocols = "grpc"
	StatsdAdvancedPluginProtocolsGrpcs          StatsdAdvancedPluginProtocols = "grpcs"
	StatsdAdvancedPluginProtocolsHTTP           StatsdAdvancedPluginProtocols = "http"
	StatsdAdvancedPluginProtocolsHTTPS          StatsdAdvancedPluginProtocols = "https"
	StatsdAdvancedPluginProtocolsTCP            StatsdAdvancedPluginProtocols = "tcp"
	StatsdAdvancedPluginProtocolsTLS            StatsdAdvancedPluginProtocols = "tls"
	StatsdAdvancedPluginProtocolsTLSPassthrough StatsdAdvancedPluginProtocols = "tls_passthrough"
	StatsdAdvancedPluginProtocolsUDP            StatsdAdvancedPluginProtocols = "udp"
	StatsdAdvancedPluginProtocolsWs             StatsdAdvancedPluginProtocols = "ws"
	StatsdAdvancedPluginProtocolsWss            StatsdAdvancedPluginProtocols = "wss"
)

func (e StatsdAdvancedPluginProtocols) ToPointer() *StatsdAdvancedPluginProtocols {
	return &e
}
func (e *StatsdAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = StatsdAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdAdvancedPluginProtocols: %v", v)
	}
}

// StatsdAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type StatsdAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// StatsdAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type StatsdAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// StatsdAdvancedPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type StatsdAdvancedPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"statsd-advanced" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                     `json:"updated_at,omitempty"`
	Config    StatsdAdvancedPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *StatsdAdvancedPluginConsumer `json:"consumer"`
	// A set of strings representing protocols.
	Protocols []StatsdAdvancedPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *StatsdAdvancedPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *StatsdAdvancedPluginService `json:"service"`
}

func (s StatsdAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StatsdAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StatsdAdvancedPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *StatsdAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *StatsdAdvancedPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *StatsdAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *StatsdAdvancedPlugin) GetName() string {
	return "statsd-advanced"
}

func (o *StatsdAdvancedPlugin) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *StatsdAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *StatsdAdvancedPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *StatsdAdvancedPlugin) GetConfig() StatsdAdvancedPluginConfig {
	if o == nil {
		return StatsdAdvancedPluginConfig{}
	}
	return o.Config
}

func (o *StatsdAdvancedPlugin) GetConsumer() *StatsdAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *StatsdAdvancedPlugin) GetProtocols() []StatsdAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *StatsdAdvancedPlugin) GetRoute() *StatsdAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *StatsdAdvancedPlugin) GetService() *StatsdAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// StatsdAdvancedPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type StatsdAdvancedPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"statsd-advanced" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string                   `json:"tags,omitempty"`
	Config StatsdAdvancedPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *StatsdAdvancedPluginConsumer `json:"consumer"`
	// A set of strings representing protocols.
	Protocols []StatsdAdvancedPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *StatsdAdvancedPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *StatsdAdvancedPluginService `json:"service"`
}

func (s StatsdAdvancedPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StatsdAdvancedPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StatsdAdvancedPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *StatsdAdvancedPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *StatsdAdvancedPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *StatsdAdvancedPluginInput) GetName() string {
	return "statsd-advanced"
}

func (o *StatsdAdvancedPluginInput) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *StatsdAdvancedPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *StatsdAdvancedPluginInput) GetConfig() StatsdAdvancedPluginConfig {
	if o == nil {
		return StatsdAdvancedPluginConfig{}
	}
	return o.Config
}

func (o *StatsdAdvancedPluginInput) GetConsumer() *StatsdAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *StatsdAdvancedPluginInput) GetProtocols() []StatsdAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *StatsdAdvancedPluginInput) GetRoute() *StatsdAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *StatsdAdvancedPluginInput) GetService() *StatsdAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
