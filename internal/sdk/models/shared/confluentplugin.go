// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type BootstrapServers struct {
	// A string representing a host name, such as example.com.
	Host string `json:"host"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port int64 `json:"port"`
}

func (o *BootstrapServers) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *BootstrapServers) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

// ProducerRequestAcks - The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set).
type ProducerRequestAcks int64

const (
	ProducerRequestAcksMinus1 ProducerRequestAcks = -1
	ProducerRequestAcksZero   ProducerRequestAcks = 0
	ProducerRequestAcksOne    ProducerRequestAcks = 1
)

func (e ProducerRequestAcks) ToPointer() *ProducerRequestAcks {
	return &e
}
func (e *ProducerRequestAcks) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 0:
		fallthrough
	case 1:
		*e = ProducerRequestAcks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProducerRequestAcks: %v", v)
	}
}

type ConfluentPluginConfig struct {
	// Set of bootstrap brokers in a `{host: host, port: port}` list format.
	BootstrapServers []BootstrapServers `json:"bootstrap_servers,omitempty"`
	// Username/Apikey for SASL authentication.
	ClusterAPIKey *string `json:"cluster_api_key,omitempty"`
	// Password/ApiSecret for SASL authentication.
	ClusterAPISecret *string `json:"cluster_api_secret,omitempty"`
	// An identifier for the Kafka cluster. By default, this field generates a random string. You can also set your own custom cluster identifier.  If more than one Kafka plugin is configured without a `cluster_name` (that is, if the default autogenerated value is removed), these plugins will use the same producer, and by extension, the same cluster. Logs will be sent to the leader of the cluster.
	ClusterName *string `json:"cluster_name,omitempty"`
	// Apikey for authentication with Confluent Cloud. This allows for management tasks such as creating topics, ACLs, etc.
	ConfluentCloudAPIKey *string `json:"confluent_cloud_api_key,omitempty"`
	// The corresponding secret for the Confluent Cloud API key.
	ConfluentCloudAPISecret *string `json:"confluent_cloud_api_secret,omitempty"`
	// Include the request body in the message. At least one of these must be true: `forward_method`, `forward_uri`, `forward_headers`, `forward_body`.
	ForwardBody *bool `json:"forward_body,omitempty"`
	// Include the request headers in the message. At least one of these must be true: `forward_method`, `forward_uri`, `forward_headers`, `forward_body`.
	ForwardHeaders *bool `json:"forward_headers,omitempty"`
	// Include the request method in the message. At least one of these must be true: `forward_method`, `forward_uri`, `forward_headers`, `forward_body`.
	ForwardMethod *bool `json:"forward_method,omitempty"`
	// Include the request URI and URI arguments (as in, query arguments) in the message. At least one of these must be true: `forward_method`, `forward_uri`, `forward_headers`, `forward_body`.
	ForwardURI *bool `json:"forward_uri,omitempty"`
	// Keepalive timeout in milliseconds.
	Keepalive        *int64 `json:"keepalive,omitempty"`
	KeepaliveEnabled *bool  `json:"keepalive_enabled,omitempty"`
	// Flag to enable asynchronous mode.
	ProducerAsync *bool `json:"producer_async,omitempty"`
	// Maximum number of messages that can be buffered in memory in asynchronous mode.
	ProducerAsyncBufferingLimitsMessagesInMemory *int64 `json:"producer_async_buffering_limits_messages_in_memory,omitempty"`
	// Maximum time interval in milliseconds between buffer flushes in asynchronous mode.
	ProducerAsyncFlushTimeout *int64 `json:"producer_async_flush_timeout,omitempty"`
	// The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set).
	ProducerRequestAcks *ProducerRequestAcks `json:"producer_request_acks,omitempty"`
	// Maximum size of a Produce request in bytes.
	ProducerRequestLimitsBytesPerRequest *int64 `json:"producer_request_limits_bytes_per_request,omitempty"`
	// Maximum number of messages to include into a single producer request.
	ProducerRequestLimitsMessagesPerRequest *int64 `json:"producer_request_limits_messages_per_request,omitempty"`
	// Backoff interval between retry attempts in milliseconds.
	ProducerRequestRetriesBackoffTimeout *int64 `json:"producer_request_retries_backoff_timeout,omitempty"`
	// Maximum number of retry attempts per single Produce request.
	ProducerRequestRetriesMaxAttempts *int64 `json:"producer_request_retries_max_attempts,omitempty"`
	// Time to wait for a Produce response in milliseconds.
	ProducerRequestTimeout *int64 `json:"producer_request_timeout,omitempty"`
	// Socket timeout in milliseconds.
	Timeout *int64 `json:"timeout,omitempty"`
	// The Kafka topic to publish to.
	Topic *string `json:"topic,omitempty"`
}

func (o *ConfluentPluginConfig) GetBootstrapServers() []BootstrapServers {
	if o == nil {
		return nil
	}
	return o.BootstrapServers
}

func (o *ConfluentPluginConfig) GetClusterAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.ClusterAPIKey
}

func (o *ConfluentPluginConfig) GetClusterAPISecret() *string {
	if o == nil {
		return nil
	}
	return o.ClusterAPISecret
}

func (o *ConfluentPluginConfig) GetClusterName() *string {
	if o == nil {
		return nil
	}
	return o.ClusterName
}

func (o *ConfluentPluginConfig) GetConfluentCloudAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.ConfluentCloudAPIKey
}

func (o *ConfluentPluginConfig) GetConfluentCloudAPISecret() *string {
	if o == nil {
		return nil
	}
	return o.ConfluentCloudAPISecret
}

func (o *ConfluentPluginConfig) GetForwardBody() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardBody
}

func (o *ConfluentPluginConfig) GetForwardHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardHeaders
}

func (o *ConfluentPluginConfig) GetForwardMethod() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardMethod
}

func (o *ConfluentPluginConfig) GetForwardURI() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardURI
}

func (o *ConfluentPluginConfig) GetKeepalive() *int64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *ConfluentPluginConfig) GetKeepaliveEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.KeepaliveEnabled
}

func (o *ConfluentPluginConfig) GetProducerAsync() *bool {
	if o == nil {
		return nil
	}
	return o.ProducerAsync
}

func (o *ConfluentPluginConfig) GetProducerAsyncBufferingLimitsMessagesInMemory() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerAsyncBufferingLimitsMessagesInMemory
}

func (o *ConfluentPluginConfig) GetProducerAsyncFlushTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerAsyncFlushTimeout
}

func (o *ConfluentPluginConfig) GetProducerRequestAcks() *ProducerRequestAcks {
	if o == nil {
		return nil
	}
	return o.ProducerRequestAcks
}

func (o *ConfluentPluginConfig) GetProducerRequestLimitsBytesPerRequest() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestLimitsBytesPerRequest
}

func (o *ConfluentPluginConfig) GetProducerRequestLimitsMessagesPerRequest() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestLimitsMessagesPerRequest
}

func (o *ConfluentPluginConfig) GetProducerRequestRetriesBackoffTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestRetriesBackoffTimeout
}

func (o *ConfluentPluginConfig) GetProducerRequestRetriesMaxAttempts() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestRetriesMaxAttempts
}

func (o *ConfluentPluginConfig) GetProducerRequestTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestTimeout
}

func (o *ConfluentPluginConfig) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *ConfluentPluginConfig) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}

// ConfluentPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type ConfluentPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *ConfluentPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ConfluentPluginProtocols string

const (
	ConfluentPluginProtocolsGrpc  ConfluentPluginProtocols = "grpc"
	ConfluentPluginProtocolsGrpcs ConfluentPluginProtocols = "grpcs"
	ConfluentPluginProtocolsHTTP  ConfluentPluginProtocols = "http"
	ConfluentPluginProtocolsHTTPS ConfluentPluginProtocols = "https"
)

func (e ConfluentPluginProtocols) ToPointer() *ConfluentPluginProtocols {
	return &e
}
func (e *ConfluentPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = ConfluentPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConfluentPluginProtocols: %v", v)
	}
}

// ConfluentPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type ConfluentPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *ConfluentPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ConfluentPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type ConfluentPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *ConfluentPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ConfluentPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ConfluentPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"confluent" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                `json:"updated_at,omitempty"`
	Config    ConfluentPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *ConfluentPluginConsumer `json:"consumer"`
	// A set of strings representing HTTP protocols.
	Protocols []ConfluentPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *ConfluentPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ConfluentPluginService `json:"service"`
}

func (c ConfluentPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConfluentPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConfluentPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ConfluentPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ConfluentPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConfluentPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *ConfluentPlugin) GetName() string {
	return "confluent"
}

func (o *ConfluentPlugin) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *ConfluentPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ConfluentPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ConfluentPlugin) GetConfig() ConfluentPluginConfig {
	if o == nil {
		return ConfluentPluginConfig{}
	}
	return o.Config
}

func (o *ConfluentPlugin) GetConsumer() *ConfluentPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ConfluentPlugin) GetProtocols() []ConfluentPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *ConfluentPlugin) GetRoute() *ConfluentPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ConfluentPlugin) GetService() *ConfluentPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// ConfluentPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ConfluentPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"confluent" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string              `json:"tags,omitempty"`
	Config ConfluentPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *ConfluentPluginConsumer `json:"consumer"`
	// A set of strings representing HTTP protocols.
	Protocols []ConfluentPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *ConfluentPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ConfluentPluginService `json:"service"`
}

func (c ConfluentPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConfluentPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConfluentPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ConfluentPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConfluentPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *ConfluentPluginInput) GetName() string {
	return "confluent"
}

func (o *ConfluentPluginInput) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *ConfluentPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ConfluentPluginInput) GetConfig() ConfluentPluginConfig {
	if o == nil {
		return ConfluentPluginConfig{}
	}
	return o.Config
}

func (o *ConfluentPluginInput) GetConsumer() *ConfluentPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ConfluentPluginInput) GetProtocols() []ConfluentPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *ConfluentPluginInput) GetRoute() *ConfluentPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ConfluentPluginInput) GetService() *ConfluentPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
