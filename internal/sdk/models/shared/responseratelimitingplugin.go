// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type ResponseRatelimitingPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *ResponseRatelimitingPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type ResponseRatelimitingPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *ResponseRatelimitingPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type ResponseRatelimitingPluginOrdering struct {
	After  *ResponseRatelimitingPluginAfter  `json:"after,omitempty"`
	Before *ResponseRatelimitingPluginBefore `json:"before,omitempty"`
}

func (o *ResponseRatelimitingPluginOrdering) GetAfter() *ResponseRatelimitingPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ResponseRatelimitingPluginOrdering) GetBefore() *ResponseRatelimitingPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type ResponseRatelimitingPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *ResponseRatelimitingPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ResponseRatelimitingPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ResponseRatelimitingPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// ResponseRatelimitingPluginLimitBy - The entity that will be used when aggregating the limits: `consumer`, `credential`, `ip`. If the `consumer` or the `credential` cannot be determined, the system will always fallback to `ip`.
type ResponseRatelimitingPluginLimitBy string

const (
	ResponseRatelimitingPluginLimitByConsumer   ResponseRatelimitingPluginLimitBy = "consumer"
	ResponseRatelimitingPluginLimitByCredential ResponseRatelimitingPluginLimitBy = "credential"
	ResponseRatelimitingPluginLimitByIP         ResponseRatelimitingPluginLimitBy = "ip"
)

func (e ResponseRatelimitingPluginLimitBy) ToPointer() *ResponseRatelimitingPluginLimitBy {
	return &e
}
func (e *ResponseRatelimitingPluginLimitBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer":
		fallthrough
	case "credential":
		fallthrough
	case "ip":
		*e = ResponseRatelimitingPluginLimitBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseRatelimitingPluginLimitBy: %v", v)
	}
}

// ResponseRatelimitingPluginPolicy - The rate-limiting policies to use for retrieving and incrementing the limits.
type ResponseRatelimitingPluginPolicy string

const (
	ResponseRatelimitingPluginPolicyCluster ResponseRatelimitingPluginPolicy = "cluster"
	ResponseRatelimitingPluginPolicyLocal   ResponseRatelimitingPluginPolicy = "local"
	ResponseRatelimitingPluginPolicyRedis   ResponseRatelimitingPluginPolicy = "redis"
)

func (e ResponseRatelimitingPluginPolicy) ToPointer() *ResponseRatelimitingPluginPolicy {
	return &e
}
func (e *ResponseRatelimitingPluginPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cluster":
		fallthrough
	case "local":
		fallthrough
	case "redis":
		*e = ResponseRatelimitingPluginPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseRatelimitingPluginPolicy: %v", v)
	}
}

// ResponseRatelimitingPluginRedis - Redis configuration
type ResponseRatelimitingPluginRedis struct {
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	Timeout *int64 `json:"timeout,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *ResponseRatelimitingPluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *ResponseRatelimitingPluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *ResponseRatelimitingPluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ResponseRatelimitingPluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ResponseRatelimitingPluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *ResponseRatelimitingPluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *ResponseRatelimitingPluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *ResponseRatelimitingPluginRedis) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *ResponseRatelimitingPluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type ResponseRatelimitingPluginConfig struct {
	// A boolean value that determines if the requests should be blocked as soon as one limit is being exceeded. This will block requests that are supposed to consume other limits too.
	BlockOnFirstViolation *bool `json:"block_on_first_violation,omitempty"`
	// A boolean value that determines if the requests should be proxied even if Kong has troubles connecting a third-party datastore. If `true`, requests will be proxied anyway, effectively disabling the rate-limiting function until the datastore is working again. If `false`, then the clients will see `500` errors.
	FaultTolerant *bool `json:"fault_tolerant,omitempty"`
	// The name of the response header used to increment the counters.
	HeaderName *string `json:"header_name,omitempty"`
	// Optionally hide informative response headers.
	HideClientHeaders *bool `json:"hide_client_headers,omitempty"`
	// The entity that will be used when aggregating the limits: `consumer`, `credential`, `ip`. If the `consumer` or the `credential` cannot be determined, the system will always fallback to `ip`.
	LimitBy *ResponseRatelimitingPluginLimitBy `json:"limit_by,omitempty"`
	// A map that defines rate limits for the plugin.
	Limits map[string]any `json:"limits,omitempty"`
	// The rate-limiting policies to use for retrieving and incrementing the limits.
	Policy *ResponseRatelimitingPluginPolicy `json:"policy,omitempty"`
	// Redis configuration
	Redis *ResponseRatelimitingPluginRedis `json:"redis,omitempty"`
}

func (o *ResponseRatelimitingPluginConfig) GetBlockOnFirstViolation() *bool {
	if o == nil {
		return nil
	}
	return o.BlockOnFirstViolation
}

func (o *ResponseRatelimitingPluginConfig) GetFaultTolerant() *bool {
	if o == nil {
		return nil
	}
	return o.FaultTolerant
}

func (o *ResponseRatelimitingPluginConfig) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *ResponseRatelimitingPluginConfig) GetHideClientHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.HideClientHeaders
}

func (o *ResponseRatelimitingPluginConfig) GetLimitBy() *ResponseRatelimitingPluginLimitBy {
	if o == nil {
		return nil
	}
	return o.LimitBy
}

func (o *ResponseRatelimitingPluginConfig) GetLimits() map[string]any {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *ResponseRatelimitingPluginConfig) GetPolicy() *ResponseRatelimitingPluginPolicy {
	if o == nil {
		return nil
	}
	return o.Policy
}

func (o *ResponseRatelimitingPluginConfig) GetRedis() *ResponseRatelimitingPluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

// ResponseRatelimitingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type ResponseRatelimitingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *ResponseRatelimitingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ResponseRatelimitingPluginProtocols string

const (
	ResponseRatelimitingPluginProtocolsGrpc  ResponseRatelimitingPluginProtocols = "grpc"
	ResponseRatelimitingPluginProtocolsGrpcs ResponseRatelimitingPluginProtocols = "grpcs"
	ResponseRatelimitingPluginProtocolsHTTP  ResponseRatelimitingPluginProtocols = "http"
	ResponseRatelimitingPluginProtocolsHTTPS ResponseRatelimitingPluginProtocols = "https"
)

func (e ResponseRatelimitingPluginProtocols) ToPointer() *ResponseRatelimitingPluginProtocols {
	return &e
}
func (e *ResponseRatelimitingPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = ResponseRatelimitingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseRatelimitingPluginProtocols: %v", v)
	}
}

// ResponseRatelimitingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type ResponseRatelimitingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *ResponseRatelimitingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ResponseRatelimitingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type ResponseRatelimitingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *ResponseRatelimitingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ResponseRatelimitingPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ResponseRatelimitingPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                                `json:"enabled,omitempty"`
	ID           *string                              `json:"id,omitempty"`
	InstanceName *string                              `json:"instance_name,omitempty"`
	name         string                               `const:"response-ratelimiting" json:"name"`
	Ordering     *ResponseRatelimitingPluginOrdering  `json:"ordering,omitempty"`
	Partials     []ResponseRatelimitingPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                            `json:"updated_at,omitempty"`
	Config    *ResponseRatelimitingPluginConfig `json:"config,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *ResponseRatelimitingPluginConsumer `json:"consumer"`
	// A set of strings representing HTTP protocols.
	Protocols []ResponseRatelimitingPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *ResponseRatelimitingPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ResponseRatelimitingPluginService `json:"service"`
}

func (r ResponseRatelimitingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResponseRatelimitingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResponseRatelimitingPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ResponseRatelimitingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ResponseRatelimitingPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ResponseRatelimitingPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *ResponseRatelimitingPlugin) GetName() string {
	return "response-ratelimiting"
}

func (o *ResponseRatelimitingPlugin) GetOrdering() *ResponseRatelimitingPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *ResponseRatelimitingPlugin) GetPartials() []ResponseRatelimitingPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *ResponseRatelimitingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ResponseRatelimitingPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ResponseRatelimitingPlugin) GetConfig() *ResponseRatelimitingPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *ResponseRatelimitingPlugin) GetConsumer() *ResponseRatelimitingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ResponseRatelimitingPlugin) GetProtocols() []ResponseRatelimitingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *ResponseRatelimitingPlugin) GetRoute() *ResponseRatelimitingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ResponseRatelimitingPlugin) GetService() *ResponseRatelimitingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
