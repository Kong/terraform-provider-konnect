// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type Behavior struct {
	// The template to use to create the body of the response to return to the consumer if Kong fails to obtain a token from the IdP.
	IdpErrorResponseBodyTemplate *string `json:"idp_error_response_body_template,omitempty"`
	// The Content-Type of the response to return to the consumer if Kong fails to obtain a token from the IdP.
	IdpErrorResponseContentType *string `json:"idp_error_response_content_type,omitempty"`
	// The message to embed in the body of the response to return to the consumer if Kong fails to obtain a token from the IdP.
	IdpErrorResponseMessage *string `json:"idp_error_response_message,omitempty"`
	// The response code to return to the consumer if Kong fails to obtain a token from the IdP.
	IdpErrorResponseStatusCode *int64 `json:"idp_error_response_status_code,omitempty"`
	// An array of status codes which will force an access token to be purged when returned by the upstream. An empty array will disable this functionality.
	PurgeTokenOnUpstreamStatusCodes []int64 `json:"purge_token_on_upstream_status_codes,omitempty"`
	// The name of the header used to send the access token (obtained from the IdP) to the upstream service.
	UpstreamAccessTokenHeaderName *string `json:"upstream_access_token_header_name,omitempty"`
}

func (o *Behavior) GetIdpErrorResponseBodyTemplate() *string {
	if o == nil {
		return nil
	}
	return o.IdpErrorResponseBodyTemplate
}

func (o *Behavior) GetIdpErrorResponseContentType() *string {
	if o == nil {
		return nil
	}
	return o.IdpErrorResponseContentType
}

func (o *Behavior) GetIdpErrorResponseMessage() *string {
	if o == nil {
		return nil
	}
	return o.IdpErrorResponseMessage
}

func (o *Behavior) GetIdpErrorResponseStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.IdpErrorResponseStatusCode
}

func (o *Behavior) GetPurgeTokenOnUpstreamStatusCodes() []int64 {
	if o == nil {
		return nil
	}
	return o.PurgeTokenOnUpstreamStatusCodes
}

func (o *Behavior) GetUpstreamAccessTokenHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamAccessTokenHeaderName
}

type UpstreamOauthPluginMemory struct {
	// The shared dictionary used by the plugin to cache tokens if `config.cache.strategy` is set to `memory`.
	DictionaryName *string `json:"dictionary_name,omitempty"`
}

func (o *UpstreamOauthPluginMemory) GetDictionaryName() *string {
	if o == nil {
		return nil
	}
	return o.DictionaryName
}

type UpstreamOauthPluginClusterNodes struct {
	// A string representing a host name, such as example.com.
	IP *string `json:"ip,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *UpstreamOauthPluginClusterNodes) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *UpstreamOauthPluginClusterNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type UpstreamOauthPluginSentinelNodes struct {
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *UpstreamOauthPluginSentinelNodes) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *UpstreamOauthPluginSentinelNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

// UpstreamOauthPluginSentinelRole - Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
type UpstreamOauthPluginSentinelRole string

const (
	UpstreamOauthPluginSentinelRoleAny    UpstreamOauthPluginSentinelRole = "any"
	UpstreamOauthPluginSentinelRoleMaster UpstreamOauthPluginSentinelRole = "master"
	UpstreamOauthPluginSentinelRoleSlave  UpstreamOauthPluginSentinelRole = "slave"
)

func (e UpstreamOauthPluginSentinelRole) ToPointer() *UpstreamOauthPluginSentinelRole {
	return &e
}
func (e *UpstreamOauthPluginSentinelRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "master":
		fallthrough
	case "slave":
		*e = UpstreamOauthPluginSentinelRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpstreamOauthPluginSentinelRole: %v", v)
	}
}

type UpstreamOauthPluginRedis struct {
	// Maximum retry attempts for redirection.
	ClusterMaxRedirections *int64 `json:"cluster_max_redirections,omitempty"`
	// Cluster addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.
	ClusterNodes []UpstreamOauthPluginClusterNodes `json:"cluster_nodes,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// If the connection to Redis is proxied (e.g. Envoy), set it `true`. Set the `host` and `port` to point to the proxy address.
	ConnectionIsProxied *bool `json:"connection_is_proxied,omitempty"`
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return `nil`. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than `keepalive_pool_size`. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than `keepalive_pool_size`.
	KeepaliveBacklog *int64 `json:"keepalive_backlog,omitempty"`
	// The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither `keepalive_pool_size` nor `keepalive_backlog` is specified, no pool is created. If `keepalive_pool_size` isn't specified but `keepalive_backlog` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.
	KeepalivePoolSize *int64 `json:"keepalive_pool_size,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.
	SentinelMaster *string `json:"sentinel_master,omitempty"`
	// Sentinel node addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.
	SentinelNodes []UpstreamOauthPluginSentinelNodes `json:"sentinel_nodes,omitempty"`
	// Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.
	SentinelPassword *string `json:"sentinel_password,omitempty"`
	// Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
	SentinelRole *UpstreamOauthPluginSentinelRole `json:"sentinel_role,omitempty"`
	// Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.
	SentinelUsername *string `json:"sentinel_username,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *UpstreamOauthPluginRedis) GetClusterMaxRedirections() *int64 {
	if o == nil {
		return nil
	}
	return o.ClusterMaxRedirections
}

func (o *UpstreamOauthPluginRedis) GetClusterNodes() []UpstreamOauthPluginClusterNodes {
	if o == nil {
		return nil
	}
	return o.ClusterNodes
}

func (o *UpstreamOauthPluginRedis) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *UpstreamOauthPluginRedis) GetConnectionIsProxied() *bool {
	if o == nil {
		return nil
	}
	return o.ConnectionIsProxied
}

func (o *UpstreamOauthPluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *UpstreamOauthPluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *UpstreamOauthPluginRedis) GetKeepaliveBacklog() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepaliveBacklog
}

func (o *UpstreamOauthPluginRedis) GetKeepalivePoolSize() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepalivePoolSize
}

func (o *UpstreamOauthPluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *UpstreamOauthPluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *UpstreamOauthPluginRedis) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *UpstreamOauthPluginRedis) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

func (o *UpstreamOauthPluginRedis) GetSentinelMaster() *string {
	if o == nil {
		return nil
	}
	return o.SentinelMaster
}

func (o *UpstreamOauthPluginRedis) GetSentinelNodes() []UpstreamOauthPluginSentinelNodes {
	if o == nil {
		return nil
	}
	return o.SentinelNodes
}

func (o *UpstreamOauthPluginRedis) GetSentinelPassword() *string {
	if o == nil {
		return nil
	}
	return o.SentinelPassword
}

func (o *UpstreamOauthPluginRedis) GetSentinelRole() *UpstreamOauthPluginSentinelRole {
	if o == nil {
		return nil
	}
	return o.SentinelRole
}

func (o *UpstreamOauthPluginRedis) GetSentinelUsername() *string {
	if o == nil {
		return nil
	}
	return o.SentinelUsername
}

func (o *UpstreamOauthPluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *UpstreamOauthPluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *UpstreamOauthPluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *UpstreamOauthPluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// UpstreamOauthPluginStrategy - The method Kong should use to cache tokens issued by the IdP.
type UpstreamOauthPluginStrategy string

const (
	UpstreamOauthPluginStrategyMemory UpstreamOauthPluginStrategy = "memory"
	UpstreamOauthPluginStrategyRedis  UpstreamOauthPluginStrategy = "redis"
)

func (e UpstreamOauthPluginStrategy) ToPointer() *UpstreamOauthPluginStrategy {
	return &e
}
func (e *UpstreamOauthPluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "memory":
		fallthrough
	case "redis":
		*e = UpstreamOauthPluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpstreamOauthPluginStrategy: %v", v)
	}
}

type Cache struct {
	// The lifetime of a token without an explicit `expires_in` value.
	DefaultTTL *float64 `json:"default_ttl,omitempty"`
	// The number of seconds to eagerly expire a cached token. By default, a cached token expires 5 seconds before its lifetime as defined in `expires_in`.
	EagerlyExpire *int64                     `json:"eagerly_expire,omitempty"`
	Memory        *UpstreamOauthPluginMemory `json:"memory,omitempty"`
	Redis         *UpstreamOauthPluginRedis  `json:"redis,omitempty"`
	// The method Kong should use to cache tokens issued by the IdP.
	Strategy *UpstreamOauthPluginStrategy `json:"strategy,omitempty"`
}

func (o *Cache) GetDefaultTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.DefaultTTL
}

func (o *Cache) GetEagerlyExpire() *int64 {
	if o == nil {
		return nil
	}
	return o.EagerlyExpire
}

func (o *Cache) GetMemory() *UpstreamOauthPluginMemory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *Cache) GetRedis() *UpstreamOauthPluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *Cache) GetStrategy() *UpstreamOauthPluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

// AuthMethod - The authentication method used in client requests to the IdP. Supported values are: `client_secret_basic` to send `client_id` and `client_secret` in the `Authorization: Basic` header, `client_secret_post` to send `client_id` and `client_secret` as part of the request body, or `client_secret_jwt` to send a JWT signed with the `client_secret` using the client assertion as part of the body.
type AuthMethod string

const (
	AuthMethodClientSecretBasic AuthMethod = "client_secret_basic"
	AuthMethodClientSecretJwt   AuthMethod = "client_secret_jwt"
	AuthMethodClientSecretPost  AuthMethod = "client_secret_post"
	AuthMethodNone              AuthMethod = "none"
)

func (e AuthMethod) ToPointer() *AuthMethod {
	return &e
}
func (e *AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_secret_basic":
		fallthrough
	case "client_secret_jwt":
		fallthrough
	case "client_secret_post":
		fallthrough
	case "none":
		*e = AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthMethod: %v", v)
	}
}

// ClientSecretJwtAlg - The algorithm to use with JWT when using `client_secret_jwt` authentication.
type ClientSecretJwtAlg string

const (
	ClientSecretJwtAlgHs256 ClientSecretJwtAlg = "HS256"
	ClientSecretJwtAlgHs512 ClientSecretJwtAlg = "HS512"
)

func (e ClientSecretJwtAlg) ToPointer() *ClientSecretJwtAlg {
	return &e
}
func (e *ClientSecretJwtAlg) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HS256":
		fallthrough
	case "HS512":
		*e = ClientSecretJwtAlg(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClientSecretJwtAlg: %v", v)
	}
}

type Client struct {
	// The authentication method used in client requests to the IdP. Supported values are: `client_secret_basic` to send `client_id` and `client_secret` in the `Authorization: Basic` header, `client_secret_post` to send `client_id` and `client_secret` as part of the request body, or `client_secret_jwt` to send a JWT signed with the `client_secret` using the client assertion as part of the body.
	AuthMethod *AuthMethod `json:"auth_method,omitempty"`
	// The algorithm to use with JWT when using `client_secret_jwt` authentication.
	ClientSecretJwtAlg *ClientSecretJwtAlg `json:"client_secret_jwt_alg,omitempty"`
	// The proxy to use when making HTTP requests to the IdP.
	HTTPProxy *string `json:"http_proxy,omitempty"`
	// The `Proxy-Authorization` header value to be used with `http_proxy`.
	HTTPProxyAuthorization *string `json:"http_proxy_authorization,omitempty"`
	// The HTTP version used for requests made by this plugin. Supported values: `1.1` for HTTP 1.1 and `1.0` for HTTP 1.0.
	HTTPVersion *float64 `json:"http_version,omitempty"`
	// The proxy to use when making HTTPS requests to the IdP.
	HTTPSProxy *string `json:"https_proxy,omitempty"`
	// The `Proxy-Authorization` header value to be used with `https_proxy`.
	HTTPSProxyAuthorization *string `json:"https_proxy_authorization,omitempty"`
	// Whether to use keepalive connections to the IdP.
	KeepAlive *bool `json:"keep_alive,omitempty"`
	// A comma-separated list of hosts that should not be proxied.
	NoProxy *string `json:"no_proxy,omitempty"`
	// Whether to verify the certificate presented by the IdP when using HTTPS.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// Network I/O timeout for requests to the IdP in milliseconds.
	Timeout *int64 `json:"timeout,omitempty"`
}

func (o *Client) GetAuthMethod() *AuthMethod {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *Client) GetClientSecretJwtAlg() *ClientSecretJwtAlg {
	if o == nil {
		return nil
	}
	return o.ClientSecretJwtAlg
}

func (o *Client) GetHTTPProxy() *string {
	if o == nil {
		return nil
	}
	return o.HTTPProxy
}

func (o *Client) GetHTTPProxyAuthorization() *string {
	if o == nil {
		return nil
	}
	return o.HTTPProxyAuthorization
}

func (o *Client) GetHTTPVersion() *float64 {
	if o == nil {
		return nil
	}
	return o.HTTPVersion
}

func (o *Client) GetHTTPSProxy() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSProxy
}

func (o *Client) GetHTTPSProxyAuthorization() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSProxyAuthorization
}

func (o *Client) GetKeepAlive() *bool {
	if o == nil {
		return nil
	}
	return o.KeepAlive
}

func (o *Client) GetNoProxy() *string {
	if o == nil {
		return nil
	}
	return o.NoProxy
}

func (o *Client) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *Client) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

// GrantType - The OAuth grant type to be used.
type GrantType string

const (
	GrantTypeClientCredentials GrantType = "client_credentials"
	GrantTypePassword          GrantType = "password"
)

func (e GrantType) ToPointer() *GrantType {
	return &e
}
func (e *GrantType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_credentials":
		fallthrough
	case "password":
		*e = GrantType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GrantType: %v", v)
	}
}

type Oauth struct {
	// List of audiences passed to the IdP when obtaining a new token.
	Audience []string `json:"audience,omitempty"`
	// The client ID for the application registration in the IdP.
	ClientID *string `json:"client_id,omitempty"`
	// The client secret for the application registration in the IdP.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The OAuth grant type to be used.
	GrantType *GrantType `json:"grant_type,omitempty"`
	// The password to use if `config.oauth.grant_type` is set to `password`.
	Password *string `json:"password,omitempty"`
	// List of scopes to request from the IdP when obtaining a new token.
	Scopes []string `json:"scopes,omitempty"`
	// The token endpoint URI.
	TokenEndpoint *string `json:"token_endpoint,omitempty"`
	// Extra headers to be passed in the token endpoint request.
	TokenHeaders map[string]any `json:"token_headers,omitempty"`
	// Extra post arguments to be passed in the token endpoint request.
	TokenPostArgs map[string]any `json:"token_post_args,omitempty"`
	// The username to use if `config.oauth.grant_type` is set to `password`.
	Username *string `json:"username,omitempty"`
}

func (o *Oauth) GetAudience() []string {
	if o == nil {
		return nil
	}
	return o.Audience
}

func (o *Oauth) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *Oauth) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *Oauth) GetGrantType() *GrantType {
	if o == nil {
		return nil
	}
	return o.GrantType
}

func (o *Oauth) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *Oauth) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *Oauth) GetTokenEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.TokenEndpoint
}

func (o *Oauth) GetTokenHeaders() map[string]any {
	if o == nil {
		return nil
	}
	return o.TokenHeaders
}

func (o *Oauth) GetTokenPostArgs() map[string]any {
	if o == nil {
		return nil
	}
	return o.TokenPostArgs
}

func (o *Oauth) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type UpstreamOauthPluginConfig struct {
	Behavior *Behavior `json:"behavior,omitempty"`
	Cache    *Cache    `json:"cache,omitempty"`
	Client   *Client   `json:"client,omitempty"`
	Oauth    *Oauth    `json:"oauth,omitempty"`
}

func (o *UpstreamOauthPluginConfig) GetBehavior() *Behavior {
	if o == nil {
		return nil
	}
	return o.Behavior
}

func (o *UpstreamOauthPluginConfig) GetCache() *Cache {
	if o == nil {
		return nil
	}
	return o.Cache
}

func (o *UpstreamOauthPluginConfig) GetClient() *Client {
	if o == nil {
		return nil
	}
	return o.Client
}

func (o *UpstreamOauthPluginConfig) GetOauth() *Oauth {
	if o == nil {
		return nil
	}
	return o.Oauth
}

// UpstreamOauthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type UpstreamOauthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *UpstreamOauthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// UpstreamOauthPluginConsumerGroup - If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
type UpstreamOauthPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *UpstreamOauthPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type UpstreamOauthPluginProtocols string

const (
	UpstreamOauthPluginProtocolsGrpc  UpstreamOauthPluginProtocols = "grpc"
	UpstreamOauthPluginProtocolsGrpcs UpstreamOauthPluginProtocols = "grpcs"
	UpstreamOauthPluginProtocolsHTTP  UpstreamOauthPluginProtocols = "http"
	UpstreamOauthPluginProtocolsHTTPS UpstreamOauthPluginProtocols = "https"
)

func (e UpstreamOauthPluginProtocols) ToPointer() *UpstreamOauthPluginProtocols {
	return &e
}
func (e *UpstreamOauthPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = UpstreamOauthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpstreamOauthPluginProtocols: %v", v)
	}
}

// UpstreamOauthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type UpstreamOauthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *UpstreamOauthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// UpstreamOauthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type UpstreamOauthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *UpstreamOauthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// UpstreamOauthPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type UpstreamOauthPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"upstream-oauth" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                    `json:"updated_at,omitempty"`
	Config    UpstreamOauthPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *UpstreamOauthPluginConsumer `json:"consumer"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *UpstreamOauthPluginConsumerGroup `json:"consumer_group"`
	// A set of strings representing HTTP protocols.
	Protocols []UpstreamOauthPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *UpstreamOauthPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *UpstreamOauthPluginService `json:"service"`
}

func (u UpstreamOauthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpstreamOauthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpstreamOauthPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UpstreamOauthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *UpstreamOauthPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpstreamOauthPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *UpstreamOauthPlugin) GetName() string {
	return "upstream-oauth"
}

func (o *UpstreamOauthPlugin) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *UpstreamOauthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *UpstreamOauthPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *UpstreamOauthPlugin) GetConfig() UpstreamOauthPluginConfig {
	if o == nil {
		return UpstreamOauthPluginConfig{}
	}
	return o.Config
}

func (o *UpstreamOauthPlugin) GetConsumer() *UpstreamOauthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *UpstreamOauthPlugin) GetConsumerGroup() *UpstreamOauthPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *UpstreamOauthPlugin) GetProtocols() []UpstreamOauthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *UpstreamOauthPlugin) GetRoute() *UpstreamOauthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *UpstreamOauthPlugin) GetService() *UpstreamOauthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// UpstreamOauthPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type UpstreamOauthPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"upstream-oauth" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string                  `json:"tags,omitempty"`
	Config UpstreamOauthPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *UpstreamOauthPluginConsumer `json:"consumer"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *UpstreamOauthPluginConsumerGroup `json:"consumer_group"`
	// A set of strings representing HTTP protocols.
	Protocols []UpstreamOauthPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *UpstreamOauthPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *UpstreamOauthPluginService `json:"service"`
}

func (u UpstreamOauthPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpstreamOauthPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpstreamOauthPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *UpstreamOauthPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpstreamOauthPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *UpstreamOauthPluginInput) GetName() string {
	return "upstream-oauth"
}

func (o *UpstreamOauthPluginInput) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *UpstreamOauthPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *UpstreamOauthPluginInput) GetConfig() UpstreamOauthPluginConfig {
	if o == nil {
		return UpstreamOauthPluginConfig{}
	}
	return o.Config
}

func (o *UpstreamOauthPluginInput) GetConsumer() *UpstreamOauthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *UpstreamOauthPluginInput) GetConsumerGroup() *UpstreamOauthPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *UpstreamOauthPluginInput) GetProtocols() []UpstreamOauthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *UpstreamOauthPluginInput) GetRoute() *UpstreamOauthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *UpstreamOauthPluginInput) GetService() *UpstreamOauthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
