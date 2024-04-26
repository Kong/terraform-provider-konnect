// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type RateLimitingAdvancedPluginProtocols string

const (
	RateLimitingAdvancedPluginProtocolsGrpc           RateLimitingAdvancedPluginProtocols = "grpc"
	RateLimitingAdvancedPluginProtocolsGrpcs          RateLimitingAdvancedPluginProtocols = "grpcs"
	RateLimitingAdvancedPluginProtocolsHTTP           RateLimitingAdvancedPluginProtocols = "http"
	RateLimitingAdvancedPluginProtocolsHTTPS          RateLimitingAdvancedPluginProtocols = "https"
	RateLimitingAdvancedPluginProtocolsTCP            RateLimitingAdvancedPluginProtocols = "tcp"
	RateLimitingAdvancedPluginProtocolsTLS            RateLimitingAdvancedPluginProtocols = "tls"
	RateLimitingAdvancedPluginProtocolsTLSPassthrough RateLimitingAdvancedPluginProtocols = "tls_passthrough"
	RateLimitingAdvancedPluginProtocolsUDP            RateLimitingAdvancedPluginProtocols = "udp"
	RateLimitingAdvancedPluginProtocolsWs             RateLimitingAdvancedPluginProtocols = "ws"
	RateLimitingAdvancedPluginProtocolsWss            RateLimitingAdvancedPluginProtocols = "wss"
)

func (e RateLimitingAdvancedPluginProtocols) ToPointer() *RateLimitingAdvancedPluginProtocols {
	return &e
}

func (e *RateLimitingAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = RateLimitingAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RateLimitingAdvancedPluginProtocols: %v", v)
	}
}

// RateLimitingAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RateLimitingAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RateLimitingAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RateLimitingAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type RateLimitingAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RateLimitingAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RateLimitingAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RateLimitingAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RateLimitingAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// Identifier - The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be `ip`, `credential`, `consumer`, `service`, `header`, `path` or `consumer-group`.
type Identifier string

const (
	IdentifierIP            Identifier = "ip"
	IdentifierCredential    Identifier = "credential"
	IdentifierConsumer      Identifier = "consumer"
	IdentifierService       Identifier = "service"
	IdentifierHeader        Identifier = "header"
	IdentifierPath          Identifier = "path"
	IdentifierConsumerGroup Identifier = "consumer-group"
)

func (e Identifier) ToPointer() *Identifier {
	return &e
}

func (e *Identifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ip":
		fallthrough
	case "credential":
		fallthrough
	case "consumer":
		fallthrough
	case "service":
		fallthrough
	case "header":
		fallthrough
	case "path":
		fallthrough
	case "consumer-group":
		*e = Identifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Identifier: %v", v)
	}
}

// SentinelRole - Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
type SentinelRole string

const (
	SentinelRoleMaster SentinelRole = "master"
	SentinelRoleSlave  SentinelRole = "slave"
	SentinelRoleAny    SentinelRole = "any"
)

func (e SentinelRole) ToPointer() *SentinelRole {
	return &e
}

func (e *SentinelRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "master":
		fallthrough
	case "slave":
		fallthrough
	case "any":
		*e = SentinelRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SentinelRole: %v", v)
	}
}

type Redis struct {
	// Cluster addresses to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Cluster. Each string element must be a hostname. The minimum length of the array is 1 element.
	ClusterAddresses []string `json:"cluster_addresses,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `default:"0" json:"database"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return `nil`. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than `keepalive_pool_size`. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than `keepalive_pool_size`.
	KeepaliveBacklog *int64 `json:"keepalive_backlog,omitempty"`
	// The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither `keepalive_pool_size` nor `keepalive_backlog` is specified, no pool is created. If `keepalive_pool_size` isn't specified but `keepalive_backlog` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.
	KeepalivePoolSize *int64 `default:"256" json:"keepalive_pool_size"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// Sentinel addresses to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel. Each string element must be a hostname. The minimum length of the array is 1 element.
	SentinelAddresses []string `json:"sentinel_addresses,omitempty"`
	// Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.
	SentinelMaster *string `json:"sentinel_master,omitempty"`
	// Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.
	SentinelPassword *string `json:"sentinel_password,omitempty"`
	// Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
	SentinelRole *SentinelRole `json:"sentinel_role,omitempty"`
	// Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.
	SentinelUsername *string `json:"sentinel_username,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `default:"false" json:"ssl"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `default:"false" json:"ssl_verify"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	Timeout *int64 `default:"2000" json:"timeout"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (r Redis) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Redis) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Redis) GetClusterAddresses() []string {
	if o == nil {
		return nil
	}
	return o.ClusterAddresses
}

func (o *Redis) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *Redis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *Redis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *Redis) GetKeepaliveBacklog() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepaliveBacklog
}

func (o *Redis) GetKeepalivePoolSize() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepalivePoolSize
}

func (o *Redis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *Redis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Redis) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *Redis) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

func (o *Redis) GetSentinelAddresses() []string {
	if o == nil {
		return nil
	}
	return o.SentinelAddresses
}

func (o *Redis) GetSentinelMaster() *string {
	if o == nil {
		return nil
	}
	return o.SentinelMaster
}

func (o *Redis) GetSentinelPassword() *string {
	if o == nil {
		return nil
	}
	return o.SentinelPassword
}

func (o *Redis) GetSentinelRole() *SentinelRole {
	if o == nil {
		return nil
	}
	return o.SentinelRole
}

func (o *Redis) GetSentinelUsername() *string {
	if o == nil {
		return nil
	}
	return o.SentinelUsername
}

func (o *Redis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *Redis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *Redis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *Redis) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *Redis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// RateLimitingAdvancedPluginStrategy - The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: `local` and `cluster`.
type RateLimitingAdvancedPluginStrategy string

const (
	RateLimitingAdvancedPluginStrategyCluster RateLimitingAdvancedPluginStrategy = "cluster"
	RateLimitingAdvancedPluginStrategyRedis   RateLimitingAdvancedPluginStrategy = "redis"
	RateLimitingAdvancedPluginStrategyLocal   RateLimitingAdvancedPluginStrategy = "local"
)

func (e RateLimitingAdvancedPluginStrategy) ToPointer() *RateLimitingAdvancedPluginStrategy {
	return &e
}

func (e *RateLimitingAdvancedPluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cluster":
		fallthrough
	case "redis":
		fallthrough
	case "local":
		*e = RateLimitingAdvancedPluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RateLimitingAdvancedPluginStrategy: %v", v)
	}
}

// WindowType - Sets the time window type to either `sliding` (default) or `fixed`. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters.
type WindowType string

const (
	WindowTypeFixed   WindowType = "fixed"
	WindowTypeSliding WindowType = "sliding"
)

func (e WindowType) ToPointer() *WindowType {
	return &e
}

func (e *WindowType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fixed":
		fallthrough
	case "sliding":
		*e = WindowType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WindowType: %v", v)
	}
}

type RateLimitingAdvancedPluginConfig struct {
	// List of consumer groups allowed to override the rate limiting settings for the given Route or Service. Required if `enforce_consumer_groups` is set to `true`.
	ConsumerGroups []string `json:"consumer_groups,omitempty"`
	// The shared dictionary where counters are stored. When the plugin is configured to synchronize counter data externally (that is `config.strategy` is `cluster` or `redis` and `config.sync_rate` isn't `-1`), this dictionary serves as a buffer to populate counters in the data store on each synchronization cycle.
	DictionaryName *string `default:"kong_rate_limiting_counters" json:"dictionary_name"`
	// If set to `true`, this doesn't count denied requests (status = `429`). If set to `false`, all requests, including denied ones, are counted. This parameter only affects the `sliding` window_type.
	DisablePenalty *bool `default:"false" json:"disable_penalty"`
	// Determines if consumer groups are allowed to override the rate limiting settings for the given Route or Service. Flipping `enforce_consumer_groups` from `true` to `false` disables the group override, but does not clear the list of consumer groups. You can then flip `enforce_consumer_groups` to `true` to re-enforce the groups.
	EnforceConsumerGroups *bool `default:"false" json:"enforce_consumer_groups"`
	// Set a custom error code to return when the rate limit is exceeded.
	ErrorCode *float64 `default:"429" json:"error_code"`
	// Set a custom error message to return when the rate limit is exceeded.
	ErrorMessage *string `default:"API rate limit exceeded" json:"error_message"`
	// A string representing an HTTP header name.
	HeaderName *string `json:"header_name,omitempty"`
	// Optionally hide informative response headers that would otherwise provide information about the current status of limits and counters.
	HideClientHeaders *bool `default:"false" json:"hide_client_headers"`
	// The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be `ip`, `credential`, `consumer`, `service`, `header`, `path` or `consumer-group`.
	Identifier *Identifier `default:"consumer" json:"identifier"`
	// One or more requests-per-window limits to apply. There must be a matching number of window limits and sizes specified.
	Limit []float64 `json:"limit,omitempty"`
	// The rate limiting library namespace to use for this plugin instance. Counter data and sync configuration is isolated in each namespace. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. `strategy`, `redis`, `sync_rate`, `window_size`, `dictionary_name`, need to be the same.
	Namespace *string `json:"namespace,omitempty"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	Path  *string `json:"path,omitempty"`
	Redis *Redis  `json:"redis,omitempty"`
	// The upper bound of a jitter (random delay) in seconds to be added to the `Retry-After` header of denied requests (status = `429`) in order to prevent all the clients from coming back at the same time. The lower bound of the jitter is `0`; in this case, the `Retry-After` header is equal to the `RateLimit-Reset` header.
	RetryAfterJitterMax *float64 `default:"0" json:"retry_after_jitter_max"`
	// The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: `local` and `cluster`.
	Strategy *RateLimitingAdvancedPluginStrategy `default:"local" json:"strategy"`
	// How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 will sync the counters in the specified number of seconds. The minimum allowed interval is 0.02 seconds (20ms).
	SyncRate *float64 `json:"sync_rate,omitempty"`
	// One or more window sizes to apply a limit to (defined in seconds). There must be a matching number of window limits and sizes specified.
	WindowSize []float64 `json:"window_size,omitempty"`
	// Sets the time window type to either `sliding` (default) or `fixed`. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters.
	WindowType *WindowType `default:"sliding" json:"window_type"`
}

func (r RateLimitingAdvancedPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RateLimitingAdvancedPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RateLimitingAdvancedPluginConfig) GetConsumerGroups() []string {
	if o == nil {
		return nil
	}
	return o.ConsumerGroups
}

func (o *RateLimitingAdvancedPluginConfig) GetDictionaryName() *string {
	if o == nil {
		return nil
	}
	return o.DictionaryName
}

func (o *RateLimitingAdvancedPluginConfig) GetDisablePenalty() *bool {
	if o == nil {
		return nil
	}
	return o.DisablePenalty
}

func (o *RateLimitingAdvancedPluginConfig) GetEnforceConsumerGroups() *bool {
	if o == nil {
		return nil
	}
	return o.EnforceConsumerGroups
}

func (o *RateLimitingAdvancedPluginConfig) GetErrorCode() *float64 {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *RateLimitingAdvancedPluginConfig) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *RateLimitingAdvancedPluginConfig) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *RateLimitingAdvancedPluginConfig) GetHideClientHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.HideClientHeaders
}

func (o *RateLimitingAdvancedPluginConfig) GetIdentifier() *Identifier {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *RateLimitingAdvancedPluginConfig) GetLimit() []float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *RateLimitingAdvancedPluginConfig) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *RateLimitingAdvancedPluginConfig) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *RateLimitingAdvancedPluginConfig) GetRedis() *Redis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *RateLimitingAdvancedPluginConfig) GetRetryAfterJitterMax() *float64 {
	if o == nil {
		return nil
	}
	return o.RetryAfterJitterMax
}

func (o *RateLimitingAdvancedPluginConfig) GetStrategy() *RateLimitingAdvancedPluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *RateLimitingAdvancedPluginConfig) GetSyncRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SyncRate
}

func (o *RateLimitingAdvancedPluginConfig) GetWindowSize() []float64 {
	if o == nil {
		return nil
	}
	return o.WindowSize
}

func (o *RateLimitingAdvancedPluginConfig) GetWindowType() *WindowType {
	if o == nil {
		return nil
	}
	return o.WindowType
}

// RateLimitingAdvancedPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RateLimitingAdvancedPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"rate-limiting-advanced" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []RateLimitingAdvancedPluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *RateLimitingAdvancedPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *RateLimitingAdvancedPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RateLimitingAdvancedPluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64                           `json:"created_at,omitempty"`
	ID        *string                          `json:"id,omitempty"`
	Config    RateLimitingAdvancedPluginConfig `json:"config"`
}

func (r RateLimitingAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RateLimitingAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RateLimitingAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RateLimitingAdvancedPlugin) GetName() string {
	return "rate-limiting-advanced"
}

func (o *RateLimitingAdvancedPlugin) GetProtocols() []RateLimitingAdvancedPluginProtocols {
	if o == nil {
		return []RateLimitingAdvancedPluginProtocols{}
	}
	return o.Protocols
}

func (o *RateLimitingAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RateLimitingAdvancedPlugin) GetConsumer() *RateLimitingAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RateLimitingAdvancedPlugin) GetRoute() *RateLimitingAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RateLimitingAdvancedPlugin) GetService() *RateLimitingAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *RateLimitingAdvancedPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RateLimitingAdvancedPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RateLimitingAdvancedPlugin) GetConfig() RateLimitingAdvancedPluginConfig {
	if o == nil {
		return RateLimitingAdvancedPluginConfig{}
	}
	return o.Config
}
