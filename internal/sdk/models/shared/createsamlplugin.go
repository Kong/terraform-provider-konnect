// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateSamlPluginNameidFormat - The requested `NameId` format. Options available are: - `Unspecified` - `EmailAddress` - `Persistent` - `Transient`
type CreateSamlPluginNameidFormat string

const (
	CreateSamlPluginNameidFormatUnspecified  CreateSamlPluginNameidFormat = "Unspecified"
	CreateSamlPluginNameidFormatEmailAddress CreateSamlPluginNameidFormat = "EmailAddress"
	CreateSamlPluginNameidFormatPersistent   CreateSamlPluginNameidFormat = "Persistent"
	CreateSamlPluginNameidFormatTransient    CreateSamlPluginNameidFormat = "Transient"
)

func (e CreateSamlPluginNameidFormat) ToPointer() *CreateSamlPluginNameidFormat {
	return &e
}
func (e *CreateSamlPluginNameidFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unspecified":
		fallthrough
	case "EmailAddress":
		fallthrough
	case "Persistent":
		fallthrough
	case "Transient":
		*e = CreateSamlPluginNameidFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginNameidFormat: %v", v)
	}
}

type CreateSamlPluginClusterNodes struct {
	// A string representing a host name, such as example.com.
	IP *string `json:"ip,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *CreateSamlPluginClusterNodes) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *CreateSamlPluginClusterNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type CreateSamlPluginSentinelNodes struct {
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *CreateSamlPluginSentinelNodes) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateSamlPluginSentinelNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

// CreateSamlPluginSentinelRole - Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
type CreateSamlPluginSentinelRole string

const (
	CreateSamlPluginSentinelRoleMaster CreateSamlPluginSentinelRole = "master"
	CreateSamlPluginSentinelRoleSlave  CreateSamlPluginSentinelRole = "slave"
	CreateSamlPluginSentinelRoleAny    CreateSamlPluginSentinelRole = "any"
)

func (e CreateSamlPluginSentinelRole) ToPointer() *CreateSamlPluginSentinelRole {
	return &e
}
func (e *CreateSamlPluginSentinelRole) UnmarshalJSON(data []byte) error {
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
		*e = CreateSamlPluginSentinelRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginSentinelRole: %v", v)
	}
}

type CreateSamlPluginRedis struct {
	// Maximum retry attempts for redirection.
	ClusterMaxRedirections *int64 `json:"cluster_max_redirections,omitempty"`
	// Cluster addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.
	ClusterNodes []CreateSamlPluginClusterNodes `json:"cluster_nodes,omitempty"`
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
	// The Redis session key prefix.
	Prefix *string `json:"prefix,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.
	SentinelMaster *string `json:"sentinel_master,omitempty"`
	// Sentinel node addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.
	SentinelNodes []CreateSamlPluginSentinelNodes `json:"sentinel_nodes,omitempty"`
	// Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.
	SentinelPassword *string `json:"sentinel_password,omitempty"`
	// Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
	SentinelRole *CreateSamlPluginSentinelRole `json:"sentinel_role,omitempty"`
	// Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.
	SentinelUsername *string `json:"sentinel_username,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// The Redis unix socket path.
	Socket *string `json:"socket,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *CreateSamlPluginRedis) GetClusterMaxRedirections() *int64 {
	if o == nil {
		return nil
	}
	return o.ClusterMaxRedirections
}

func (o *CreateSamlPluginRedis) GetClusterNodes() []CreateSamlPluginClusterNodes {
	if o == nil {
		return nil
	}
	return o.ClusterNodes
}

func (o *CreateSamlPluginRedis) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *CreateSamlPluginRedis) GetConnectionIsProxied() *bool {
	if o == nil {
		return nil
	}
	return o.ConnectionIsProxied
}

func (o *CreateSamlPluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *CreateSamlPluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateSamlPluginRedis) GetKeepaliveBacklog() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepaliveBacklog
}

func (o *CreateSamlPluginRedis) GetKeepalivePoolSize() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepalivePoolSize
}

func (o *CreateSamlPluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateSamlPluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateSamlPluginRedis) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *CreateSamlPluginRedis) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *CreateSamlPluginRedis) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

func (o *CreateSamlPluginRedis) GetSentinelMaster() *string {
	if o == nil {
		return nil
	}
	return o.SentinelMaster
}

func (o *CreateSamlPluginRedis) GetSentinelNodes() []CreateSamlPluginSentinelNodes {
	if o == nil {
		return nil
	}
	return o.SentinelNodes
}

func (o *CreateSamlPluginRedis) GetSentinelPassword() *string {
	if o == nil {
		return nil
	}
	return o.SentinelPassword
}

func (o *CreateSamlPluginRedis) GetSentinelRole() *CreateSamlPluginSentinelRole {
	if o == nil {
		return nil
	}
	return o.SentinelRole
}

func (o *CreateSamlPluginRedis) GetSentinelUsername() *string {
	if o == nil {
		return nil
	}
	return o.SentinelUsername
}

func (o *CreateSamlPluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *CreateSamlPluginRedis) GetSocket() *string {
	if o == nil {
		return nil
	}
	return o.Socket
}

func (o *CreateSamlPluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *CreateSamlPluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *CreateSamlPluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// CreateSamlPluginRequestDigestAlgorithm - The digest algorithm for Authn requests: - `SHA256` - `SHA1`
type CreateSamlPluginRequestDigestAlgorithm string

const (
	CreateSamlPluginRequestDigestAlgorithmSha256 CreateSamlPluginRequestDigestAlgorithm = "SHA256"
	CreateSamlPluginRequestDigestAlgorithmSha1   CreateSamlPluginRequestDigestAlgorithm = "SHA1"
)

func (e CreateSamlPluginRequestDigestAlgorithm) ToPointer() *CreateSamlPluginRequestDigestAlgorithm {
	return &e
}
func (e *CreateSamlPluginRequestDigestAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHA256":
		fallthrough
	case "SHA1":
		*e = CreateSamlPluginRequestDigestAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginRequestDigestAlgorithm: %v", v)
	}
}

// CreateSamlPluginRequestSignatureAlgorithm - The signature algorithm for signing Authn requests. Options available are: - `SHA256` - `SHA384` - `SHA512`
type CreateSamlPluginRequestSignatureAlgorithm string

const (
	CreateSamlPluginRequestSignatureAlgorithmSha256 CreateSamlPluginRequestSignatureAlgorithm = "SHA256"
	CreateSamlPluginRequestSignatureAlgorithmSha384 CreateSamlPluginRequestSignatureAlgorithm = "SHA384"
	CreateSamlPluginRequestSignatureAlgorithmSha512 CreateSamlPluginRequestSignatureAlgorithm = "SHA512"
)

func (e CreateSamlPluginRequestSignatureAlgorithm) ToPointer() *CreateSamlPluginRequestSignatureAlgorithm {
	return &e
}
func (e *CreateSamlPluginRequestSignatureAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHA256":
		fallthrough
	case "SHA384":
		fallthrough
	case "SHA512":
		*e = CreateSamlPluginRequestSignatureAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginRequestSignatureAlgorithm: %v", v)
	}
}

// CreateSamlPluginResponseDigestAlgorithm - The algorithm for verifying digest in SAML responses: - `SHA256` - `SHA1`
type CreateSamlPluginResponseDigestAlgorithm string

const (
	CreateSamlPluginResponseDigestAlgorithmSha256 CreateSamlPluginResponseDigestAlgorithm = "SHA256"
	CreateSamlPluginResponseDigestAlgorithmSha1   CreateSamlPluginResponseDigestAlgorithm = "SHA1"
)

func (e CreateSamlPluginResponseDigestAlgorithm) ToPointer() *CreateSamlPluginResponseDigestAlgorithm {
	return &e
}
func (e *CreateSamlPluginResponseDigestAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHA256":
		fallthrough
	case "SHA1":
		*e = CreateSamlPluginResponseDigestAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginResponseDigestAlgorithm: %v", v)
	}
}

// CreateSamlPluginResponseSignatureAlgorithm - The algorithm for validating signatures in SAML responses. Options available are: - `SHA256` - `SHA384` - `SHA512`
type CreateSamlPluginResponseSignatureAlgorithm string

const (
	CreateSamlPluginResponseSignatureAlgorithmSha256 CreateSamlPluginResponseSignatureAlgorithm = "SHA256"
	CreateSamlPluginResponseSignatureAlgorithmSha384 CreateSamlPluginResponseSignatureAlgorithm = "SHA384"
	CreateSamlPluginResponseSignatureAlgorithmSha512 CreateSamlPluginResponseSignatureAlgorithm = "SHA512"
)

func (e CreateSamlPluginResponseSignatureAlgorithm) ToPointer() *CreateSamlPluginResponseSignatureAlgorithm {
	return &e
}
func (e *CreateSamlPluginResponseSignatureAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHA256":
		fallthrough
	case "SHA384":
		fallthrough
	case "SHA512":
		*e = CreateSamlPluginResponseSignatureAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginResponseSignatureAlgorithm: %v", v)
	}
}

// CreateSamlPluginSessionCookieSameSite - Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.
type CreateSamlPluginSessionCookieSameSite string

const (
	CreateSamlPluginSessionCookieSameSiteStrict  CreateSamlPluginSessionCookieSameSite = "Strict"
	CreateSamlPluginSessionCookieSameSiteLax     CreateSamlPluginSessionCookieSameSite = "Lax"
	CreateSamlPluginSessionCookieSameSiteNone    CreateSamlPluginSessionCookieSameSite = "None"
	CreateSamlPluginSessionCookieSameSiteDefault CreateSamlPluginSessionCookieSameSite = "Default"
)

func (e CreateSamlPluginSessionCookieSameSite) ToPointer() *CreateSamlPluginSessionCookieSameSite {
	return &e
}
func (e *CreateSamlPluginSessionCookieSameSite) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Strict":
		fallthrough
	case "Lax":
		fallthrough
	case "None":
		fallthrough
	case "Default":
		*e = CreateSamlPluginSessionCookieSameSite(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginSessionCookieSameSite: %v", v)
	}
}

type CreateSamlPluginSessionRequestHeaders string

const (
	CreateSamlPluginSessionRequestHeadersID              CreateSamlPluginSessionRequestHeaders = "id"
	CreateSamlPluginSessionRequestHeadersAudience        CreateSamlPluginSessionRequestHeaders = "audience"
	CreateSamlPluginSessionRequestHeadersSubject         CreateSamlPluginSessionRequestHeaders = "subject"
	CreateSamlPluginSessionRequestHeadersTimeout         CreateSamlPluginSessionRequestHeaders = "timeout"
	CreateSamlPluginSessionRequestHeadersIdlingTimeout   CreateSamlPluginSessionRequestHeaders = "idling-timeout"
	CreateSamlPluginSessionRequestHeadersRollingTimeout  CreateSamlPluginSessionRequestHeaders = "rolling-timeout"
	CreateSamlPluginSessionRequestHeadersAbsoluteTimeout CreateSamlPluginSessionRequestHeaders = "absolute-timeout"
)

func (e CreateSamlPluginSessionRequestHeaders) ToPointer() *CreateSamlPluginSessionRequestHeaders {
	return &e
}
func (e *CreateSamlPluginSessionRequestHeaders) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "audience":
		fallthrough
	case "subject":
		fallthrough
	case "timeout":
		fallthrough
	case "idling-timeout":
		fallthrough
	case "rolling-timeout":
		fallthrough
	case "absolute-timeout":
		*e = CreateSamlPluginSessionRequestHeaders(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginSessionRequestHeaders: %v", v)
	}
}

type CreateSamlPluginSessionResponseHeaders string

const (
	CreateSamlPluginSessionResponseHeadersID              CreateSamlPluginSessionResponseHeaders = "id"
	CreateSamlPluginSessionResponseHeadersAudience        CreateSamlPluginSessionResponseHeaders = "audience"
	CreateSamlPluginSessionResponseHeadersSubject         CreateSamlPluginSessionResponseHeaders = "subject"
	CreateSamlPluginSessionResponseHeadersTimeout         CreateSamlPluginSessionResponseHeaders = "timeout"
	CreateSamlPluginSessionResponseHeadersIdlingTimeout   CreateSamlPluginSessionResponseHeaders = "idling-timeout"
	CreateSamlPluginSessionResponseHeadersRollingTimeout  CreateSamlPluginSessionResponseHeaders = "rolling-timeout"
	CreateSamlPluginSessionResponseHeadersAbsoluteTimeout CreateSamlPluginSessionResponseHeaders = "absolute-timeout"
)

func (e CreateSamlPluginSessionResponseHeaders) ToPointer() *CreateSamlPluginSessionResponseHeaders {
	return &e
}
func (e *CreateSamlPluginSessionResponseHeaders) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "audience":
		fallthrough
	case "subject":
		fallthrough
	case "timeout":
		fallthrough
	case "idling-timeout":
		fallthrough
	case "rolling-timeout":
		fallthrough
	case "absolute-timeout":
		*e = CreateSamlPluginSessionResponseHeaders(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginSessionResponseHeaders: %v", v)
	}
}

// CreateSamlPluginSessionStorage - The session storage for session data: - `cookie`: stores session data with the session cookie. The session cannot be invalidated or revoked without changing the session secret, but is stateless, and doesn't require a database. - `memcached`: stores session data in memcached - `redis`: stores session data in Redis
type CreateSamlPluginSessionStorage string

const (
	CreateSamlPluginSessionStorageCookie    CreateSamlPluginSessionStorage = "cookie"
	CreateSamlPluginSessionStorageMemcache  CreateSamlPluginSessionStorage = "memcache"
	CreateSamlPluginSessionStorageMemcached CreateSamlPluginSessionStorage = "memcached"
	CreateSamlPluginSessionStorageRedis     CreateSamlPluginSessionStorage = "redis"
)

func (e CreateSamlPluginSessionStorage) ToPointer() *CreateSamlPluginSessionStorage {
	return &e
}
func (e *CreateSamlPluginSessionStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cookie":
		fallthrough
	case "memcache":
		fallthrough
	case "memcached":
		fallthrough
	case "redis":
		*e = CreateSamlPluginSessionStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginSessionStorage: %v", v)
	}
}

type CreateSamlPluginConfig struct {
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer. If not set, a Kong Consumer must exist for the SAML IdP user credentials, mapping the username format to the Kong Consumer username.
	Anonymous *string `json:"anonymous,omitempty"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	AssertionConsumerPath *string `json:"assertion_consumer_path,omitempty"`
	// The public certificate provided by the IdP. This is used to validate responses from the IdP.  Only include the contents of the certificate. Do not include the header (`BEGIN CERTIFICATE`) and footer (`END CERTIFICATE`) lines.
	IdpCertificate *string `json:"idp_certificate,omitempty"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	IdpSsoURL *string `json:"idp_sso_url,omitempty"`
	// The unique identifier of the IdP application. Formatted as a URL containing information about the IdP so the SP can validate that the SAML assertions it receives are issued from the correct IdP.
	Issuer *string `json:"issuer,omitempty"`
	// The requested `NameId` format. Options available are: - `Unspecified` - `EmailAddress` - `Persistent` - `Transient`
	NameidFormat *CreateSamlPluginNameidFormat `json:"nameid_format,omitempty"`
	Redis        *CreateSamlPluginRedis        `json:"redis,omitempty"`
	// The digest algorithm for Authn requests: - `SHA256` - `SHA1`
	RequestDigestAlgorithm *CreateSamlPluginRequestDigestAlgorithm `json:"request_digest_algorithm,omitempty"`
	// The signature algorithm for signing Authn requests. Options available are: - `SHA256` - `SHA384` - `SHA512`
	RequestSignatureAlgorithm *CreateSamlPluginRequestSignatureAlgorithm `json:"request_signature_algorithm,omitempty"`
	// The certificate for signing requests.
	RequestSigningCertificate *string `json:"request_signing_certificate,omitempty"`
	// The private key for signing requests.  If this parameter is set, requests sent to the IdP are signed.  The `request_signing_certificate` parameter must be set as well.
	RequestSigningKey *string `json:"request_signing_key,omitempty"`
	// The algorithm for verifying digest in SAML responses: - `SHA256` - `SHA1`
	ResponseDigestAlgorithm *CreateSamlPluginResponseDigestAlgorithm `json:"response_digest_algorithm,omitempty"`
	// The private encryption key required to decrypt encrypted assertions.
	ResponseEncryptionKey *string `json:"response_encryption_key,omitempty"`
	// The algorithm for validating signatures in SAML responses. Options available are: - `SHA256` - `SHA384` - `SHA512`
	ResponseSignatureAlgorithm *CreateSamlPluginResponseSignatureAlgorithm `json:"response_signature_algorithm,omitempty"`
	// The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid.
	SessionAbsoluteTimeout *float64 `json:"session_absolute_timeout,omitempty"`
	// The session audience, for example "my-application"
	SessionAudience *string `json:"session_audience,omitempty"`
	// The session cookie domain flag.
	SessionCookieDomain *string `json:"session_cookie_domain,omitempty"`
	// Forbids JavaScript from accessing the cookie, for example, through the `Document.cookie` property.
	SessionCookieHTTPOnly *bool `json:"session_cookie_http_only,omitempty"`
	// The session cookie name.
	SessionCookieName *string `json:"session_cookie_name,omitempty"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	SessionCookiePath *string `json:"session_cookie_path,omitempty"`
	// Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.
	SessionCookieSameSite *CreateSamlPluginSessionCookieSameSite `json:"session_cookie_same_site,omitempty"`
	// The cookie is only sent to the server when a request is made with the https:scheme (except on localhost), and therefore is more resistant to man-in-the-middle attacks.
	SessionCookieSecure *bool `json:"session_cookie_secure,omitempty"`
	// When set to `true`, audiences are forced to share the same subject.
	SessionEnforceSameSubject *bool `json:"session_enforce_same_subject,omitempty"`
	// When set to `true`, the storage key (session ID) is hashed for extra security. Hashing the storage key means it is impossible to decrypt data from the storage without a cookie.
	SessionHashStorageKey *bool `json:"session_hash_storage_key,omitempty"`
	// When set to `true`, the value of subject is hashed before being stored. Only applies when `session_store_metadata` is enabled.
	SessionHashSubject *bool `json:"session_hash_subject,omitempty"`
	// The session cookie idle time in seconds.
	SessionIdlingTimeout *float64 `json:"session_idling_timeout,omitempty"`
	// The memcached host.
	SessionMemcachedHost *string `json:"session_memcached_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	SessionMemcachedPort *int64 `json:"session_memcached_port,omitempty"`
	// The memcached session key prefix.
	SessionMemcachedPrefix *string `json:"session_memcached_prefix,omitempty"`
	// The memcached unix socket path.
	SessionMemcachedSocket *string `json:"session_memcached_socket,omitempty"`
	// Enables or disables persistent sessions
	SessionRemember *bool `json:"session_remember,omitempty"`
	// Persistent session absolute timeout in seconds.
	SessionRememberAbsoluteTimeout *float64 `json:"session_remember_absolute_timeout,omitempty"`
	// Persistent session cookie name
	SessionRememberCookieName *string `json:"session_remember_cookie_name,omitempty"`
	// Persistent session rolling timeout in seconds.
	SessionRememberRollingTimeout *float64                                 `json:"session_remember_rolling_timeout,omitempty"`
	SessionRequestHeaders         []CreateSamlPluginSessionRequestHeaders  `json:"session_request_headers,omitempty"`
	SessionResponseHeaders        []CreateSamlPluginSessionResponseHeaders `json:"session_response_headers,omitempty"`
	// The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid.
	SessionRollingTimeout *float64 `json:"session_rolling_timeout,omitempty"`
	// The session secret. This must be a random string of 32 characters from the base64 alphabet (letters, numbers, `/`, `_` and `+`). It is used as the secret key for encrypting session data as well as state information that is sent to the IdP in the authentication exchange.
	SessionSecret *string `json:"session_secret,omitempty"`
	// The session storage for session data: - `cookie`: stores session data with the session cookie. The session cannot be invalidated or revoked without changing the session secret, but is stateless, and doesn't require a database. - `memcached`: stores session data in memcached - `redis`: stores session data in Redis
	SessionStorage *CreateSamlPluginSessionStorage `json:"session_storage,omitempty"`
	// Configures whether or not session metadata should be stored. This includes information about the active sessions for the `specific_audience` belonging to a specific subject.
	SessionStoreMetadata *bool `json:"session_store_metadata,omitempty"`
	// Enable signature validation for SAML responses.
	ValidateAssertionSignature *bool `json:"validate_assertion_signature,omitempty"`
}

func (o *CreateSamlPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *CreateSamlPluginConfig) GetAssertionConsumerPath() *string {
	if o == nil {
		return nil
	}
	return o.AssertionConsumerPath
}

func (o *CreateSamlPluginConfig) GetIdpCertificate() *string {
	if o == nil {
		return nil
	}
	return o.IdpCertificate
}

func (o *CreateSamlPluginConfig) GetIdpSsoURL() *string {
	if o == nil {
		return nil
	}
	return o.IdpSsoURL
}

func (o *CreateSamlPluginConfig) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *CreateSamlPluginConfig) GetNameidFormat() *CreateSamlPluginNameidFormat {
	if o == nil {
		return nil
	}
	return o.NameidFormat
}

func (o *CreateSamlPluginConfig) GetRedis() *CreateSamlPluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *CreateSamlPluginConfig) GetRequestDigestAlgorithm() *CreateSamlPluginRequestDigestAlgorithm {
	if o == nil {
		return nil
	}
	return o.RequestDigestAlgorithm
}

func (o *CreateSamlPluginConfig) GetRequestSignatureAlgorithm() *CreateSamlPluginRequestSignatureAlgorithm {
	if o == nil {
		return nil
	}
	return o.RequestSignatureAlgorithm
}

func (o *CreateSamlPluginConfig) GetRequestSigningCertificate() *string {
	if o == nil {
		return nil
	}
	return o.RequestSigningCertificate
}

func (o *CreateSamlPluginConfig) GetRequestSigningKey() *string {
	if o == nil {
		return nil
	}
	return o.RequestSigningKey
}

func (o *CreateSamlPluginConfig) GetResponseDigestAlgorithm() *CreateSamlPluginResponseDigestAlgorithm {
	if o == nil {
		return nil
	}
	return o.ResponseDigestAlgorithm
}

func (o *CreateSamlPluginConfig) GetResponseEncryptionKey() *string {
	if o == nil {
		return nil
	}
	return o.ResponseEncryptionKey
}

func (o *CreateSamlPluginConfig) GetResponseSignatureAlgorithm() *CreateSamlPluginResponseSignatureAlgorithm {
	if o == nil {
		return nil
	}
	return o.ResponseSignatureAlgorithm
}

func (o *CreateSamlPluginConfig) GetSessionAbsoluteTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionAbsoluteTimeout
}

func (o *CreateSamlPluginConfig) GetSessionAudience() *string {
	if o == nil {
		return nil
	}
	return o.SessionAudience
}

func (o *CreateSamlPluginConfig) GetSessionCookieDomain() *string {
	if o == nil {
		return nil
	}
	return o.SessionCookieDomain
}

func (o *CreateSamlPluginConfig) GetSessionCookieHTTPOnly() *bool {
	if o == nil {
		return nil
	}
	return o.SessionCookieHTTPOnly
}

func (o *CreateSamlPluginConfig) GetSessionCookieName() *string {
	if o == nil {
		return nil
	}
	return o.SessionCookieName
}

func (o *CreateSamlPluginConfig) GetSessionCookiePath() *string {
	if o == nil {
		return nil
	}
	return o.SessionCookiePath
}

func (o *CreateSamlPluginConfig) GetSessionCookieSameSite() *CreateSamlPluginSessionCookieSameSite {
	if o == nil {
		return nil
	}
	return o.SessionCookieSameSite
}

func (o *CreateSamlPluginConfig) GetSessionCookieSecure() *bool {
	if o == nil {
		return nil
	}
	return o.SessionCookieSecure
}

func (o *CreateSamlPluginConfig) GetSessionEnforceSameSubject() *bool {
	if o == nil {
		return nil
	}
	return o.SessionEnforceSameSubject
}

func (o *CreateSamlPluginConfig) GetSessionHashStorageKey() *bool {
	if o == nil {
		return nil
	}
	return o.SessionHashStorageKey
}

func (o *CreateSamlPluginConfig) GetSessionHashSubject() *bool {
	if o == nil {
		return nil
	}
	return o.SessionHashSubject
}

func (o *CreateSamlPluginConfig) GetSessionIdlingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionIdlingTimeout
}

func (o *CreateSamlPluginConfig) GetSessionMemcachedHost() *string {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedHost
}

func (o *CreateSamlPluginConfig) GetSessionMemcachedPort() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedPort
}

func (o *CreateSamlPluginConfig) GetSessionMemcachedPrefix() *string {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedPrefix
}

func (o *CreateSamlPluginConfig) GetSessionMemcachedSocket() *string {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedSocket
}

func (o *CreateSamlPluginConfig) GetSessionRemember() *bool {
	if o == nil {
		return nil
	}
	return o.SessionRemember
}

func (o *CreateSamlPluginConfig) GetSessionRememberAbsoluteTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionRememberAbsoluteTimeout
}

func (o *CreateSamlPluginConfig) GetSessionRememberCookieName() *string {
	if o == nil {
		return nil
	}
	return o.SessionRememberCookieName
}

func (o *CreateSamlPluginConfig) GetSessionRememberRollingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionRememberRollingTimeout
}

func (o *CreateSamlPluginConfig) GetSessionRequestHeaders() []CreateSamlPluginSessionRequestHeaders {
	if o == nil {
		return nil
	}
	return o.SessionRequestHeaders
}

func (o *CreateSamlPluginConfig) GetSessionResponseHeaders() []CreateSamlPluginSessionResponseHeaders {
	if o == nil {
		return nil
	}
	return o.SessionResponseHeaders
}

func (o *CreateSamlPluginConfig) GetSessionRollingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionRollingTimeout
}

func (o *CreateSamlPluginConfig) GetSessionSecret() *string {
	if o == nil {
		return nil
	}
	return o.SessionSecret
}

func (o *CreateSamlPluginConfig) GetSessionStorage() *CreateSamlPluginSessionStorage {
	if o == nil {
		return nil
	}
	return o.SessionStorage
}

func (o *CreateSamlPluginConfig) GetSessionStoreMetadata() *bool {
	if o == nil {
		return nil
	}
	return o.SessionStoreMetadata
}

func (o *CreateSamlPluginConfig) GetValidateAssertionSignature() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateAssertionSignature
}

type CreateSamlPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateSamlPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateSamlPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateSamlPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateSamlPluginOrdering struct {
	After  *CreateSamlPluginAfter  `json:"after,omitempty"`
	Before *CreateSamlPluginBefore `json:"before,omitempty"`
}

func (o *CreateSamlPluginOrdering) GetAfter() *CreateSamlPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateSamlPluginOrdering) GetBefore() *CreateSamlPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateSamlPluginProtocols string

const (
	CreateSamlPluginProtocolsGrpc           CreateSamlPluginProtocols = "grpc"
	CreateSamlPluginProtocolsGrpcs          CreateSamlPluginProtocols = "grpcs"
	CreateSamlPluginProtocolsHTTP           CreateSamlPluginProtocols = "http"
	CreateSamlPluginProtocolsHTTPS          CreateSamlPluginProtocols = "https"
	CreateSamlPluginProtocolsTCP            CreateSamlPluginProtocols = "tcp"
	CreateSamlPluginProtocolsTLS            CreateSamlPluginProtocols = "tls"
	CreateSamlPluginProtocolsTLSPassthrough CreateSamlPluginProtocols = "tls_passthrough"
	CreateSamlPluginProtocolsUDP            CreateSamlPluginProtocols = "udp"
	CreateSamlPluginProtocolsWs             CreateSamlPluginProtocols = "ws"
	CreateSamlPluginProtocolsWss            CreateSamlPluginProtocols = "wss"
)

func (e CreateSamlPluginProtocols) ToPointer() *CreateSamlPluginProtocols {
	return &e
}
func (e *CreateSamlPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateSamlPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSamlPluginProtocols: %v", v)
	}
}

// CreateSamlPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateSamlPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateSamlPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateSamlPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateSamlPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateSamlPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateSamlPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateSamlPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateSamlPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateSamlPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateSamlPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateSamlPlugin struct {
	Config *CreateSamlPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                     `json:"enabled,omitempty"`
	InstanceName *string                   `json:"instance_name,omitempty"`
	name         *string                   `const:"saml" json:"name,omitempty"`
	Ordering     *CreateSamlPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateSamlPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateSamlPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateSamlPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateSamlPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateSamlPluginService `json:"service,omitempty"`
}

func (c CreateSamlPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateSamlPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateSamlPlugin) GetConfig() *CreateSamlPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateSamlPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateSamlPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateSamlPlugin) GetName() *string {
	return types.String("saml")
}

func (o *CreateSamlPlugin) GetOrdering() *CreateSamlPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateSamlPlugin) GetProtocols() []CreateSamlPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateSamlPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateSamlPlugin) GetConsumer() *CreateSamlPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateSamlPlugin) GetConsumerGroup() *CreateSamlPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateSamlPlugin) GetRoute() *CreateSamlPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateSamlPlugin) GetService() *CreateSamlPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
