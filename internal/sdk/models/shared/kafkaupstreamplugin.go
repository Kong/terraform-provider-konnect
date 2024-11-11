// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

// KafkaUpstreamPluginMechanism - The SASL authentication mechanism.  Supported options: `PLAIN`, `SCRAM-SHA-256`, or `SCRAM-SHA-512`.
type KafkaUpstreamPluginMechanism string

const (
	KafkaUpstreamPluginMechanismPlain       KafkaUpstreamPluginMechanism = "PLAIN"
	KafkaUpstreamPluginMechanismScramSha256 KafkaUpstreamPluginMechanism = "SCRAM-SHA-256"
	KafkaUpstreamPluginMechanismScramSha512 KafkaUpstreamPluginMechanism = "SCRAM-SHA-512"
)

func (e KafkaUpstreamPluginMechanism) ToPointer() *KafkaUpstreamPluginMechanism {
	return &e
}
func (e *KafkaUpstreamPluginMechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PLAIN":
		fallthrough
	case "SCRAM-SHA-256":
		fallthrough
	case "SCRAM-SHA-512":
		*e = KafkaUpstreamPluginMechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaUpstreamPluginMechanism: %v", v)
	}
}

// KafkaUpstreamPluginStrategy - The authentication strategy for the plugin, the only option for the value is `sasl`.
type KafkaUpstreamPluginStrategy string

const (
	KafkaUpstreamPluginStrategySasl KafkaUpstreamPluginStrategy = "sasl"
)

func (e KafkaUpstreamPluginStrategy) ToPointer() *KafkaUpstreamPluginStrategy {
	return &e
}
func (e *KafkaUpstreamPluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sasl":
		*e = KafkaUpstreamPluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaUpstreamPluginStrategy: %v", v)
	}
}

type KafkaUpstreamPluginAuthentication struct {
	// The SASL authentication mechanism.  Supported options: `PLAIN`, `SCRAM-SHA-256`, or `SCRAM-SHA-512`.
	Mechanism *KafkaUpstreamPluginMechanism `json:"mechanism,omitempty"`
	// Password for SASL authentication.
	Password *string `json:"password,omitempty"`
	// The authentication strategy for the plugin, the only option for the value is `sasl`.
	Strategy *KafkaUpstreamPluginStrategy `json:"strategy,omitempty"`
	// Enable this to indicate `DelegationToken` authentication.
	Tokenauth *bool `json:"tokenauth,omitempty"`
	// Username for SASL authentication.
	User *string `json:"user,omitempty"`
}

func (o *KafkaUpstreamPluginAuthentication) GetMechanism() *KafkaUpstreamPluginMechanism {
	if o == nil {
		return nil
	}
	return o.Mechanism
}

func (o *KafkaUpstreamPluginAuthentication) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *KafkaUpstreamPluginAuthentication) GetStrategy() *KafkaUpstreamPluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *KafkaUpstreamPluginAuthentication) GetTokenauth() *bool {
	if o == nil {
		return nil
	}
	return o.Tokenauth
}

func (o *KafkaUpstreamPluginAuthentication) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

type KafkaUpstreamPluginBootstrapServers struct {
	// A string representing a host name, such as example.com.
	Host string `json:"host"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port int64 `json:"port"`
}

func (o *KafkaUpstreamPluginBootstrapServers) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *KafkaUpstreamPluginBootstrapServers) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

// KafkaUpstreamPluginProducerRequestAcks - The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set).
type KafkaUpstreamPluginProducerRequestAcks int64

const (
	KafkaUpstreamPluginProducerRequestAcksMinus1 KafkaUpstreamPluginProducerRequestAcks = -1
	KafkaUpstreamPluginProducerRequestAcksZero   KafkaUpstreamPluginProducerRequestAcks = 0
	KafkaUpstreamPluginProducerRequestAcksOne    KafkaUpstreamPluginProducerRequestAcks = 1
)

func (e KafkaUpstreamPluginProducerRequestAcks) ToPointer() *KafkaUpstreamPluginProducerRequestAcks {
	return &e
}
func (e *KafkaUpstreamPluginProducerRequestAcks) UnmarshalJSON(data []byte) error {
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
		*e = KafkaUpstreamPluginProducerRequestAcks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaUpstreamPluginProducerRequestAcks: %v", v)
	}
}

type KafkaUpstreamPluginSecurity struct {
	// UUID of certificate entity for mTLS authentication.
	CertificateID *string `json:"certificate_id,omitempty"`
	// Enables TLS.
	Ssl *bool `json:"ssl,omitempty"`
}

func (o *KafkaUpstreamPluginSecurity) GetCertificateID() *string {
	if o == nil {
		return nil
	}
	return o.CertificateID
}

func (o *KafkaUpstreamPluginSecurity) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

type KafkaUpstreamPluginConfig struct {
	Authentication *KafkaUpstreamPluginAuthentication `json:"authentication,omitempty"`
	// Set of bootstrap brokers in a `{host: host, port: port}` list format.
	BootstrapServers []KafkaUpstreamPluginBootstrapServers `json:"bootstrap_servers,omitempty"`
	// An identifier for the Kafka cluster. By default, this field generates a random string. You can also set your own custom cluster identifier.  If more than one Kafka plugin is configured without a `cluster_name` (that is, if the default autogenerated value is removed), these plugins will use the same producer, and by extension, the same cluster. Logs will be sent to the leader of the cluster.
	ClusterName *string `json:"cluster_name,omitempty"`
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
	ProducerRequestAcks *KafkaUpstreamPluginProducerRequestAcks `json:"producer_request_acks,omitempty"`
	// Maximum size of a Produce request in bytes.
	ProducerRequestLimitsBytesPerRequest *int64 `json:"producer_request_limits_bytes_per_request,omitempty"`
	// Maximum number of messages to include into a single producer request.
	ProducerRequestLimitsMessagesPerRequest *int64 `json:"producer_request_limits_messages_per_request,omitempty"`
	// Backoff interval between retry attempts in milliseconds.
	ProducerRequestRetriesBackoffTimeout *int64 `json:"producer_request_retries_backoff_timeout,omitempty"`
	// Maximum number of retry attempts per single Produce request.
	ProducerRequestRetriesMaxAttempts *int64 `json:"producer_request_retries_max_attempts,omitempty"`
	// Time to wait for a Produce response in milliseconds.
	ProducerRequestTimeout *int64                       `json:"producer_request_timeout,omitempty"`
	Security               *KafkaUpstreamPluginSecurity `json:"security,omitempty"`
	// Socket timeout in milliseconds.
	Timeout *int64 `json:"timeout,omitempty"`
	// The Kafka topic to publish to.
	Topic *string `json:"topic,omitempty"`
}

func (o *KafkaUpstreamPluginConfig) GetAuthentication() *KafkaUpstreamPluginAuthentication {
	if o == nil {
		return nil
	}
	return o.Authentication
}

func (o *KafkaUpstreamPluginConfig) GetBootstrapServers() []KafkaUpstreamPluginBootstrapServers {
	if o == nil {
		return nil
	}
	return o.BootstrapServers
}

func (o *KafkaUpstreamPluginConfig) GetClusterName() *string {
	if o == nil {
		return nil
	}
	return o.ClusterName
}

func (o *KafkaUpstreamPluginConfig) GetForwardBody() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardBody
}

func (o *KafkaUpstreamPluginConfig) GetForwardHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardHeaders
}

func (o *KafkaUpstreamPluginConfig) GetForwardMethod() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardMethod
}

func (o *KafkaUpstreamPluginConfig) GetForwardURI() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardURI
}

func (o *KafkaUpstreamPluginConfig) GetKeepalive() *int64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *KafkaUpstreamPluginConfig) GetKeepaliveEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.KeepaliveEnabled
}

func (o *KafkaUpstreamPluginConfig) GetProducerAsync() *bool {
	if o == nil {
		return nil
	}
	return o.ProducerAsync
}

func (o *KafkaUpstreamPluginConfig) GetProducerAsyncBufferingLimitsMessagesInMemory() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerAsyncBufferingLimitsMessagesInMemory
}

func (o *KafkaUpstreamPluginConfig) GetProducerAsyncFlushTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerAsyncFlushTimeout
}

func (o *KafkaUpstreamPluginConfig) GetProducerRequestAcks() *KafkaUpstreamPluginProducerRequestAcks {
	if o == nil {
		return nil
	}
	return o.ProducerRequestAcks
}

func (o *KafkaUpstreamPluginConfig) GetProducerRequestLimitsBytesPerRequest() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestLimitsBytesPerRequest
}

func (o *KafkaUpstreamPluginConfig) GetProducerRequestLimitsMessagesPerRequest() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestLimitsMessagesPerRequest
}

func (o *KafkaUpstreamPluginConfig) GetProducerRequestRetriesBackoffTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestRetriesBackoffTimeout
}

func (o *KafkaUpstreamPluginConfig) GetProducerRequestRetriesMaxAttempts() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestRetriesMaxAttempts
}

func (o *KafkaUpstreamPluginConfig) GetProducerRequestTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestTimeout
}

func (o *KafkaUpstreamPluginConfig) GetSecurity() *KafkaUpstreamPluginSecurity {
	if o == nil {
		return nil
	}
	return o.Security
}

func (o *KafkaUpstreamPluginConfig) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *KafkaUpstreamPluginConfig) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}

// KafkaUpstreamPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type KafkaUpstreamPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *KafkaUpstreamPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KafkaUpstreamPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *KafkaUpstreamPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KafkaUpstreamPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *KafkaUpstreamPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type KafkaUpstreamPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *KafkaUpstreamPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type KafkaUpstreamPluginOrdering struct {
	After  *KafkaUpstreamPluginAfter  `json:"after,omitempty"`
	Before *KafkaUpstreamPluginBefore `json:"before,omitempty"`
}

func (o *KafkaUpstreamPluginOrdering) GetAfter() *KafkaUpstreamPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *KafkaUpstreamPluginOrdering) GetBefore() *KafkaUpstreamPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type KafkaUpstreamPluginProtocols string

const (
	KafkaUpstreamPluginProtocolsGrpc           KafkaUpstreamPluginProtocols = "grpc"
	KafkaUpstreamPluginProtocolsGrpcs          KafkaUpstreamPluginProtocols = "grpcs"
	KafkaUpstreamPluginProtocolsHTTP           KafkaUpstreamPluginProtocols = "http"
	KafkaUpstreamPluginProtocolsHTTPS          KafkaUpstreamPluginProtocols = "https"
	KafkaUpstreamPluginProtocolsTCP            KafkaUpstreamPluginProtocols = "tcp"
	KafkaUpstreamPluginProtocolsTLS            KafkaUpstreamPluginProtocols = "tls"
	KafkaUpstreamPluginProtocolsTLSPassthrough KafkaUpstreamPluginProtocols = "tls_passthrough"
	KafkaUpstreamPluginProtocolsUDP            KafkaUpstreamPluginProtocols = "udp"
	KafkaUpstreamPluginProtocolsWs             KafkaUpstreamPluginProtocols = "ws"
	KafkaUpstreamPluginProtocolsWss            KafkaUpstreamPluginProtocols = "wss"
)

func (e KafkaUpstreamPluginProtocols) ToPointer() *KafkaUpstreamPluginProtocols {
	return &e
}
func (e *KafkaUpstreamPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = KafkaUpstreamPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaUpstreamPluginProtocols: %v", v)
	}
}

// KafkaUpstreamPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type KafkaUpstreamPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *KafkaUpstreamPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KafkaUpstreamPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type KafkaUpstreamPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *KafkaUpstreamPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KafkaUpstreamPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type KafkaUpstreamPlugin struct {
	Config KafkaUpstreamPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *KafkaUpstreamPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *KafkaUpstreamPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	ID           *string                      `json:"id,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         string                       `const:"kafka-upstream" json:"name"`
	Ordering     *KafkaUpstreamPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []KafkaUpstreamPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *KafkaUpstreamPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *KafkaUpstreamPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (k KafkaUpstreamPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KafkaUpstreamPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KafkaUpstreamPlugin) GetConfig() KafkaUpstreamPluginConfig {
	if o == nil {
		return KafkaUpstreamPluginConfig{}
	}
	return o.Config
}

func (o *KafkaUpstreamPlugin) GetConsumer() *KafkaUpstreamPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *KafkaUpstreamPlugin) GetConsumerGroup() *KafkaUpstreamPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *KafkaUpstreamPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *KafkaUpstreamPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *KafkaUpstreamPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KafkaUpstreamPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *KafkaUpstreamPlugin) GetName() string {
	return "kafka-upstream"
}

func (o *KafkaUpstreamPlugin) GetOrdering() *KafkaUpstreamPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *KafkaUpstreamPlugin) GetProtocols() []KafkaUpstreamPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *KafkaUpstreamPlugin) GetRoute() *KafkaUpstreamPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *KafkaUpstreamPlugin) GetService() *KafkaUpstreamPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *KafkaUpstreamPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *KafkaUpstreamPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// KafkaUpstreamPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type KafkaUpstreamPluginInput struct {
	Config KafkaUpstreamPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *KafkaUpstreamPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *KafkaUpstreamPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	ID           *string                      `json:"id,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         string                       `const:"kafka-upstream" json:"name"`
	Ordering     *KafkaUpstreamPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []KafkaUpstreamPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *KafkaUpstreamPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *KafkaUpstreamPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (k KafkaUpstreamPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KafkaUpstreamPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KafkaUpstreamPluginInput) GetConfig() KafkaUpstreamPluginConfig {
	if o == nil {
		return KafkaUpstreamPluginConfig{}
	}
	return o.Config
}

func (o *KafkaUpstreamPluginInput) GetConsumer() *KafkaUpstreamPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *KafkaUpstreamPluginInput) GetConsumerGroup() *KafkaUpstreamPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *KafkaUpstreamPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *KafkaUpstreamPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KafkaUpstreamPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *KafkaUpstreamPluginInput) GetName() string {
	return "kafka-upstream"
}

func (o *KafkaUpstreamPluginInput) GetOrdering() *KafkaUpstreamPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *KafkaUpstreamPluginInput) GetProtocols() []KafkaUpstreamPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *KafkaUpstreamPluginInput) GetRoute() *KafkaUpstreamPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *KafkaUpstreamPluginInput) GetService() *KafkaUpstreamPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *KafkaUpstreamPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}