// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type ConsumerIdentifierDefault string

const (
	ConsumerIdentifierDefaultConsumerID ConsumerIdentifierDefault = "consumer_id"
	ConsumerIdentifierDefaultCustomID   ConsumerIdentifierDefault = "custom_id"
	ConsumerIdentifierDefaultUsername   ConsumerIdentifierDefault = "username"
)

func (e ConsumerIdentifierDefault) ToPointer() *ConsumerIdentifierDefault {
	return &e
}
func (e *ConsumerIdentifierDefault) UnmarshalJSON(data []byte) error {
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
		*e = ConsumerIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConsumerIdentifierDefault: %v", v)
	}
}

// ConsumerIdentifier - Authenticated user detail.
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

// Name - StatsD metric’s name.
type Name string

const (
	NameKongLatency                Name = "kong_latency"
	NameLatency                    Name = "latency"
	NameRequestCount               Name = "request_count"
	NameRequestPerUser             Name = "request_per_user"
	NameRequestSize                Name = "request_size"
	NameResponseSize               Name = "response_size"
	NameStatusCount                Name = "status_count"
	NameStatusCountPerUser         Name = "status_count_per_user"
	NameUniqueUsers                Name = "unique_users"
	NameUpstreamLatency            Name = "upstream_latency"
	NameStatusCountPerWorkspace    Name = "status_count_per_workspace"
	NameStatusCountPerUserPerRoute Name = "status_count_per_user_per_route"
	NameShdictUsage                Name = "shdict_usage"
	NameCacheDatastoreHitsTotal    Name = "cache_datastore_hits_total"
	NameCacheDatastoreMissesTotal  Name = "cache_datastore_misses_total"
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
		*e = Name(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Name: %v", v)
	}
}

// ServiceIdentifier - Service detail.
type ServiceIdentifier string

const (
	ServiceIdentifierServiceID         ServiceIdentifier = "service_id"
	ServiceIdentifierServiceName       ServiceIdentifier = "service_name"
	ServiceIdentifierServiceHost       ServiceIdentifier = "service_host"
	ServiceIdentifierServiceNameOrHost ServiceIdentifier = "service_name_or_host"
)

func (e ServiceIdentifier) ToPointer() *ServiceIdentifier {
	return &e
}
func (e *ServiceIdentifier) UnmarshalJSON(data []byte) error {
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
		*e = ServiceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServiceIdentifier: %v", v)
	}
}

// StatType - Determines what sort of event a metric represents.
type StatType string

const (
	StatTypeCounter   StatType = "counter"
	StatTypeGauge     StatType = "gauge"
	StatTypeHistogram StatType = "histogram"
	StatTypeMeter     StatType = "meter"
	StatTypeSet       StatType = "set"
	StatTypeTimer     StatType = "timer"
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
		*e = StatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatType: %v", v)
	}
}

// WorkspaceIdentifier - Workspace detail.
type WorkspaceIdentifier string

const (
	WorkspaceIdentifierWorkspaceID   WorkspaceIdentifier = "workspace_id"
	WorkspaceIdentifierWorkspaceName WorkspaceIdentifier = "workspace_name"
)

func (e WorkspaceIdentifier) ToPointer() *WorkspaceIdentifier {
	return &e
}
func (e *WorkspaceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = WorkspaceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WorkspaceIdentifier: %v", v)
	}
}

type Metrics struct {
	// Authenticated user detail.
	ConsumerIdentifier *ConsumerIdentifier `json:"consumer_identifier,omitempty"`
	// StatsD metric’s name.
	Name Name `json:"name"`
	// Sampling rate
	SampleRate *float64 `json:"sample_rate,omitempty"`
	// Service detail.
	ServiceIdentifier *ServiceIdentifier `json:"service_identifier,omitempty"`
	// Determines what sort of event a metric represents.
	StatType StatType `json:"stat_type"`
	// Workspace detail.
	WorkspaceIdentifier *WorkspaceIdentifier `json:"workspace_identifier,omitempty"`
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

func (o *Metrics) GetServiceIdentifier() *ServiceIdentifier {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifier
}

func (o *Metrics) GetStatType() StatType {
	if o == nil {
		return StatType("")
	}
	return o.StatType
}

func (o *Metrics) GetWorkspaceIdentifier() *WorkspaceIdentifier {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifier
}

type StatsdPluginQueue struct {
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

func (o *StatsdPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *StatsdPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *StatsdPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *StatsdPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *StatsdPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *StatsdPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *StatsdPluginQueue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

type ServiceIdentifierDefault string

const (
	ServiceIdentifierDefaultServiceID         ServiceIdentifierDefault = "service_id"
	ServiceIdentifierDefaultServiceName       ServiceIdentifierDefault = "service_name"
	ServiceIdentifierDefaultServiceHost       ServiceIdentifierDefault = "service_host"
	ServiceIdentifierDefaultServiceNameOrHost ServiceIdentifierDefault = "service_name_or_host"
)

func (e ServiceIdentifierDefault) ToPointer() *ServiceIdentifierDefault {
	return &e
}
func (e *ServiceIdentifierDefault) UnmarshalJSON(data []byte) error {
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
		*e = ServiceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServiceIdentifierDefault: %v", v)
	}
}

type TagStyle string

const (
	TagStyleDogstatsd TagStyle = "dogstatsd"
	TagStyleInfluxdb  TagStyle = "influxdb"
	TagStyleLibrato   TagStyle = "librato"
	TagStyleSignalfx  TagStyle = "signalfx"
)

func (e TagStyle) ToPointer() *TagStyle {
	return &e
}
func (e *TagStyle) UnmarshalJSON(data []byte) error {
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
		*e = TagStyle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TagStyle: %v", v)
	}
}

type WorkspaceIdentifierDefault string

const (
	WorkspaceIdentifierDefaultWorkspaceID   WorkspaceIdentifierDefault = "workspace_id"
	WorkspaceIdentifierDefaultWorkspaceName WorkspaceIdentifierDefault = "workspace_name"
)

func (e WorkspaceIdentifierDefault) ToPointer() *WorkspaceIdentifierDefault {
	return &e
}
func (e *WorkspaceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = WorkspaceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WorkspaceIdentifierDefault: %v", v)
	}
}

type StatsdPluginConfig struct {
	// List of status code ranges that are allowed to be logged in metrics.
	AllowStatusCodes          []string                   `json:"allow_status_codes,omitempty"`
	ConsumerIdentifierDefault *ConsumerIdentifierDefault `json:"consumer_identifier_default,omitempty"`
	FlushTimeout              *float64                   `json:"flush_timeout,omitempty"`
	// The IP address or hostname of StatsD server to send data to.
	Host             *string `json:"host,omitempty"`
	HostnameInPrefix *bool   `json:"hostname_in_prefix,omitempty"`
	// List of metrics to be logged.
	Metrics []Metrics `json:"metrics,omitempty"`
	// The port of StatsD server to send data to.
	Port *int64 `json:"port,omitempty"`
	// String to prefix to each metric's name.
	Prefix                     *string                     `json:"prefix,omitempty"`
	Queue                      *StatsdPluginQueue          `json:"queue,omitempty"`
	QueueSize                  *int64                      `json:"queue_size,omitempty"`
	RetryCount                 *int64                      `json:"retry_count,omitempty"`
	ServiceIdentifierDefault   *ServiceIdentifierDefault   `json:"service_identifier_default,omitempty"`
	TagStyle                   *TagStyle                   `json:"tag_style,omitempty"`
	UDPPacketSize              *float64                    `json:"udp_packet_size,omitempty"`
	UseTCP                     *bool                       `json:"use_tcp,omitempty"`
	WorkspaceIdentifierDefault *WorkspaceIdentifierDefault `json:"workspace_identifier_default,omitempty"`
}

func (o *StatsdPluginConfig) GetAllowStatusCodes() []string {
	if o == nil {
		return nil
	}
	return o.AllowStatusCodes
}

func (o *StatsdPluginConfig) GetConsumerIdentifierDefault() *ConsumerIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifierDefault
}

func (o *StatsdPluginConfig) GetFlushTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushTimeout
}

func (o *StatsdPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *StatsdPluginConfig) GetHostnameInPrefix() *bool {
	if o == nil {
		return nil
	}
	return o.HostnameInPrefix
}

func (o *StatsdPluginConfig) GetMetrics() []Metrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *StatsdPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *StatsdPluginConfig) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *StatsdPluginConfig) GetQueue() *StatsdPluginQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *StatsdPluginConfig) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}

func (o *StatsdPluginConfig) GetRetryCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RetryCount
}

func (o *StatsdPluginConfig) GetServiceIdentifierDefault() *ServiceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifierDefault
}

func (o *StatsdPluginConfig) GetTagStyle() *TagStyle {
	if o == nil {
		return nil
	}
	return o.TagStyle
}

func (o *StatsdPluginConfig) GetUDPPacketSize() *float64 {
	if o == nil {
		return nil
	}
	return o.UDPPacketSize
}

func (o *StatsdPluginConfig) GetUseTCP() *bool {
	if o == nil {
		return nil
	}
	return o.UseTCP
}

func (o *StatsdPluginConfig) GetWorkspaceIdentifierDefault() *WorkspaceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifierDefault
}

type StatsdPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *StatsdPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type StatsdPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *StatsdPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type StatsdPluginOrdering struct {
	After  *StatsdPluginAfter  `json:"after,omitempty"`
	Before *StatsdPluginBefore `json:"before,omitempty"`
}

func (o *StatsdPluginOrdering) GetAfter() *StatsdPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *StatsdPluginOrdering) GetBefore() *StatsdPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type StatsdPluginProtocols string

const (
	StatsdPluginProtocolsGrpc           StatsdPluginProtocols = "grpc"
	StatsdPluginProtocolsGrpcs          StatsdPluginProtocols = "grpcs"
	StatsdPluginProtocolsHTTP           StatsdPluginProtocols = "http"
	StatsdPluginProtocolsHTTPS          StatsdPluginProtocols = "https"
	StatsdPluginProtocolsTCP            StatsdPluginProtocols = "tcp"
	StatsdPluginProtocolsTLS            StatsdPluginProtocols = "tls"
	StatsdPluginProtocolsTLSPassthrough StatsdPluginProtocols = "tls_passthrough"
	StatsdPluginProtocolsUDP            StatsdPluginProtocols = "udp"
	StatsdPluginProtocolsWs             StatsdPluginProtocols = "ws"
	StatsdPluginProtocolsWss            StatsdPluginProtocols = "wss"
)

func (e StatsdPluginProtocols) ToPointer() *StatsdPluginProtocols {
	return &e
}
func (e *StatsdPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = StatsdPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsdPluginProtocols: %v", v)
	}
}

// StatsdPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type StatsdPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type StatsdPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// StatsdPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type StatsdPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// StatsdPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type StatsdPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *StatsdPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type StatsdPlugin struct {
	Config *StatsdPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                 `json:"enabled,omitempty"`
	ID           *string               `json:"id,omitempty"`
	InstanceName *string               `json:"instance_name,omitempty"`
	name         *string               `const:"statsd" json:"name,omitempty"`
	Ordering     *StatsdPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []StatsdPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *StatsdPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *StatsdPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *StatsdPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *StatsdPluginService `json:"service,omitempty"`
}

func (s StatsdPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StatsdPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StatsdPlugin) GetConfig() *StatsdPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *StatsdPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *StatsdPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *StatsdPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *StatsdPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *StatsdPlugin) GetName() *string {
	return types.String("statsd")
}

func (o *StatsdPlugin) GetOrdering() *StatsdPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *StatsdPlugin) GetProtocols() []StatsdPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *StatsdPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *StatsdPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *StatsdPlugin) GetConsumer() *StatsdPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *StatsdPlugin) GetConsumerGroup() *StatsdPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *StatsdPlugin) GetRoute() *StatsdPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *StatsdPlugin) GetService() *StatsdPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
