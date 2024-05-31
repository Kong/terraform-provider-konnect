// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type SamlPluginProtocols string

const (
	SamlPluginProtocolsGrpc           SamlPluginProtocols = "grpc"
	SamlPluginProtocolsGrpcs          SamlPluginProtocols = "grpcs"
	SamlPluginProtocolsHTTP           SamlPluginProtocols = "http"
	SamlPluginProtocolsHTTPS          SamlPluginProtocols = "https"
	SamlPluginProtocolsTCP            SamlPluginProtocols = "tcp"
	SamlPluginProtocolsTLS            SamlPluginProtocols = "tls"
	SamlPluginProtocolsTLSPassthrough SamlPluginProtocols = "tls_passthrough"
	SamlPluginProtocolsUDP            SamlPluginProtocols = "udp"
	SamlPluginProtocolsWs             SamlPluginProtocols = "ws"
	SamlPluginProtocolsWss            SamlPluginProtocols = "wss"
)

func (e SamlPluginProtocols) ToPointer() *SamlPluginProtocols {
	return &e
}
func (e *SamlPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = SamlPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SamlPluginProtocols: %v", v)
	}
}

// SamlPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type SamlPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *SamlPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SamlPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type SamlPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *SamlPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SamlPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type SamlPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *SamlPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// NameidFormat - The requested `NameId` format. Options available are: - `Unspecified` - `EmailAddress` - `Persistent` - `Transient`
type NameidFormat string

const (
	NameidFormatUnspecified  NameidFormat = "Unspecified"
	NameidFormatEmailAddress NameidFormat = "EmailAddress"
	NameidFormatPersistent   NameidFormat = "Persistent"
	NameidFormatTransient    NameidFormat = "Transient"
)

func (e NameidFormat) ToPointer() *NameidFormat {
	return &e
}
func (e *NameidFormat) UnmarshalJSON(data []byte) error {
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
		*e = NameidFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NameidFormat: %v", v)
	}
}

// RequestDigestAlgorithm - The digest algorithm for Authn requests: - `SHA256` - `SHA1`
type RequestDigestAlgorithm string

const (
	RequestDigestAlgorithmSha256 RequestDigestAlgorithm = "SHA256"
	RequestDigestAlgorithmSha1   RequestDigestAlgorithm = "SHA1"
)

func (e RequestDigestAlgorithm) ToPointer() *RequestDigestAlgorithm {
	return &e
}
func (e *RequestDigestAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHA256":
		fallthrough
	case "SHA1":
		*e = RequestDigestAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestDigestAlgorithm: %v", v)
	}
}

// RequestSignatureAlgorithm - The signature algorithm for signing Authn requests. Options available are: - `SHA256` - `SHA384` - `SHA512`
type RequestSignatureAlgorithm string

const (
	RequestSignatureAlgorithmSha256 RequestSignatureAlgorithm = "SHA256"
	RequestSignatureAlgorithmSha384 RequestSignatureAlgorithm = "SHA384"
	RequestSignatureAlgorithmSha512 RequestSignatureAlgorithm = "SHA512"
)

func (e RequestSignatureAlgorithm) ToPointer() *RequestSignatureAlgorithm {
	return &e
}
func (e *RequestSignatureAlgorithm) UnmarshalJSON(data []byte) error {
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
		*e = RequestSignatureAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestSignatureAlgorithm: %v", v)
	}
}

// ResponseDigestAlgorithm - The algorithm for verifying digest in SAML responses: - `SHA256` - `SHA1`
type ResponseDigestAlgorithm string

const (
	ResponseDigestAlgorithmSha256 ResponseDigestAlgorithm = "SHA256"
	ResponseDigestAlgorithmSha1   ResponseDigestAlgorithm = "SHA1"
)

func (e ResponseDigestAlgorithm) ToPointer() *ResponseDigestAlgorithm {
	return &e
}
func (e *ResponseDigestAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHA256":
		fallthrough
	case "SHA1":
		*e = ResponseDigestAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseDigestAlgorithm: %v", v)
	}
}

// ResponseSignatureAlgorithm - The algorithm for validating signatures in SAML responses. Options available are: - `SHA256` - `SHA384` - `SHA512`
type ResponseSignatureAlgorithm string

const (
	ResponseSignatureAlgorithmSha256 ResponseSignatureAlgorithm = "SHA256"
	ResponseSignatureAlgorithmSha384 ResponseSignatureAlgorithm = "SHA384"
	ResponseSignatureAlgorithmSha512 ResponseSignatureAlgorithm = "SHA512"
)

func (e ResponseSignatureAlgorithm) ToPointer() *ResponseSignatureAlgorithm {
	return &e
}
func (e *ResponseSignatureAlgorithm) UnmarshalJSON(data []byte) error {
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
		*e = ResponseSignatureAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseSignatureAlgorithm: %v", v)
	}
}

// SamlPluginSessionCookieSameSite - Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.
type SamlPluginSessionCookieSameSite string

const (
	SamlPluginSessionCookieSameSiteStrict  SamlPluginSessionCookieSameSite = "Strict"
	SamlPluginSessionCookieSameSiteLax     SamlPluginSessionCookieSameSite = "Lax"
	SamlPluginSessionCookieSameSiteNone    SamlPluginSessionCookieSameSite = "None"
	SamlPluginSessionCookieSameSiteDefault SamlPluginSessionCookieSameSite = "Default"
)

func (e SamlPluginSessionCookieSameSite) ToPointer() *SamlPluginSessionCookieSameSite {
	return &e
}
func (e *SamlPluginSessionCookieSameSite) UnmarshalJSON(data []byte) error {
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
		*e = SamlPluginSessionCookieSameSite(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SamlPluginSessionCookieSameSite: %v", v)
	}
}

type SamlPluginSessionRedisClusterNodes struct {
	// A string representing a host name, such as example.com.
	IP *string `default:"127.0.0.1" json:"ip"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `default:"6379" json:"port"`
}

func (s SamlPluginSessionRedisClusterNodes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SamlPluginSessionRedisClusterNodes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SamlPluginSessionRedisClusterNodes) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *SamlPluginSessionRedisClusterNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type SamlPluginSessionRequestHeaders string

const (
	SamlPluginSessionRequestHeadersID              SamlPluginSessionRequestHeaders = "id"
	SamlPluginSessionRequestHeadersAudience        SamlPluginSessionRequestHeaders = "audience"
	SamlPluginSessionRequestHeadersSubject         SamlPluginSessionRequestHeaders = "subject"
	SamlPluginSessionRequestHeadersTimeout         SamlPluginSessionRequestHeaders = "timeout"
	SamlPluginSessionRequestHeadersIdlingTimeout   SamlPluginSessionRequestHeaders = "idling-timeout"
	SamlPluginSessionRequestHeadersRollingTimeout  SamlPluginSessionRequestHeaders = "rolling-timeout"
	SamlPluginSessionRequestHeadersAbsoluteTimeout SamlPluginSessionRequestHeaders = "absolute-timeout"
)

func (e SamlPluginSessionRequestHeaders) ToPointer() *SamlPluginSessionRequestHeaders {
	return &e
}
func (e *SamlPluginSessionRequestHeaders) UnmarshalJSON(data []byte) error {
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
		*e = SamlPluginSessionRequestHeaders(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SamlPluginSessionRequestHeaders: %v", v)
	}
}

type SamlPluginSessionResponseHeaders string

const (
	SamlPluginSessionResponseHeadersID              SamlPluginSessionResponseHeaders = "id"
	SamlPluginSessionResponseHeadersAudience        SamlPluginSessionResponseHeaders = "audience"
	SamlPluginSessionResponseHeadersSubject         SamlPluginSessionResponseHeaders = "subject"
	SamlPluginSessionResponseHeadersTimeout         SamlPluginSessionResponseHeaders = "timeout"
	SamlPluginSessionResponseHeadersIdlingTimeout   SamlPluginSessionResponseHeaders = "idling-timeout"
	SamlPluginSessionResponseHeadersRollingTimeout  SamlPluginSessionResponseHeaders = "rolling-timeout"
	SamlPluginSessionResponseHeadersAbsoluteTimeout SamlPluginSessionResponseHeaders = "absolute-timeout"
)

func (e SamlPluginSessionResponseHeaders) ToPointer() *SamlPluginSessionResponseHeaders {
	return &e
}
func (e *SamlPluginSessionResponseHeaders) UnmarshalJSON(data []byte) error {
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
		*e = SamlPluginSessionResponseHeaders(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SamlPluginSessionResponseHeaders: %v", v)
	}
}

// SamlPluginSessionStorage - The session storage for session data: - `cookie`: stores session data with the session cookie. The session cannot be invalidated or revoked without changing the session secret, but is stateless, and doesn't require a database. - `memcached`: stores session data in memcached - `redis`: stores session data in Redis
type SamlPluginSessionStorage string

const (
	SamlPluginSessionStorageCookie    SamlPluginSessionStorage = "cookie"
	SamlPluginSessionStorageMemcache  SamlPluginSessionStorage = "memcache"
	SamlPluginSessionStorageMemcached SamlPluginSessionStorage = "memcached"
	SamlPluginSessionStorageRedis     SamlPluginSessionStorage = "redis"
)

func (e SamlPluginSessionStorage) ToPointer() *SamlPluginSessionStorage {
	return &e
}
func (e *SamlPluginSessionStorage) UnmarshalJSON(data []byte) error {
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
		*e = SamlPluginSessionStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SamlPluginSessionStorage: %v", v)
	}
}

type SamlPluginConfig struct {
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
	NameidFormat *NameidFormat `default:"EmailAddress" json:"nameid_format"`
	// The digest algorithm for Authn requests: - `SHA256` - `SHA1`
	RequestDigestAlgorithm *RequestDigestAlgorithm `default:"SHA256" json:"request_digest_algorithm"`
	// The signature algorithm for signing Authn requests. Options available are: - `SHA256` - `SHA384` - `SHA512`
	RequestSignatureAlgorithm *RequestSignatureAlgorithm `default:"SHA256" json:"request_signature_algorithm"`
	// The certificate for signing requests.
	RequestSigningCertificate *string `json:"request_signing_certificate,omitempty"`
	// The private key for signing requests.  If this parameter is set, requests sent to the IdP are signed.  The `request_signing_certificate` parameter must be set as well.
	RequestSigningKey *string `json:"request_signing_key,omitempty"`
	// The algorithm for verifying digest in SAML responses: - `SHA256` - `SHA1`
	ResponseDigestAlgorithm *ResponseDigestAlgorithm `default:"SHA256" json:"response_digest_algorithm"`
	// The private encryption key required to decrypt encrypted assertions.
	ResponseEncryptionKey *string `json:"response_encryption_key,omitempty"`
	// The algorithm for validating signatures in SAML responses. Options available are: - `SHA256` - `SHA384` - `SHA512`
	ResponseSignatureAlgorithm *ResponseSignatureAlgorithm `default:"SHA256" json:"response_signature_algorithm"`
	// The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid.
	SessionAbsoluteTimeout *float64 `default:"86400" json:"session_absolute_timeout"`
	// The session audience, for example "my-application"
	SessionAudience *string `default:"default" json:"session_audience"`
	// The session cookie domain flag.
	SessionCookieDomain *string `json:"session_cookie_domain,omitempty"`
	// Forbids JavaScript from accessing the cookie, for example, through the `Document.cookie` property.
	SessionCookieHTTPOnly *bool `default:"true" json:"session_cookie_http_only"`
	// The session cookie name.
	SessionCookieName *string `default:"session" json:"session_cookie_name"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	SessionCookiePath *string `default:"/" json:"session_cookie_path"`
	// Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.
	SessionCookieSameSite *SamlPluginSessionCookieSameSite `default:"Lax" json:"session_cookie_same_site"`
	// The cookie is only sent to the server when a request is made with the https:scheme (except on localhost), and therefore is more resistant to man-in-the-middle attacks.
	SessionCookieSecure *bool `json:"session_cookie_secure,omitempty"`
	// When set to `true`, audiences are forced to share the same subject.
	SessionEnforceSameSubject *bool `default:"false" json:"session_enforce_same_subject"`
	// When set to `true`, the storage key (session ID) is hashed for extra security. Hashing the storage key means it is impossible to decrypt data from the storage without a cookie.
	SessionHashStorageKey *bool `default:"false" json:"session_hash_storage_key"`
	// When set to `true`, the value of subject is hashed before being stored. Only applies when `session_store_metadata` is enabled.
	SessionHashSubject *bool `default:"false" json:"session_hash_subject"`
	// The session cookie idle time in seconds.
	SessionIdlingTimeout *float64 `default:"900" json:"session_idling_timeout"`
	// The memcached host.
	SessionMemcachedHost *string `default:"127.0.0.1" json:"session_memcached_host"`
	// An integer representing a port number between 0 and 65535, inclusive.
	SessionMemcachedPort *int64 `default:"11211" json:"session_memcached_port"`
	// The memcached session key prefix.
	SessionMemcachedPrefix *string `json:"session_memcached_prefix,omitempty"`
	// The memcached unix socket path.
	SessionMemcachedSocket *string `json:"session_memcached_socket,omitempty"`
	// The Redis cluster maximum redirects.
	SessionRedisClusterMaxRedirections *int64 `json:"session_redis_cluster_max_redirections,omitempty"`
	// The Redis cluster node host. Takes an array of host records, with either `ip` or `host`, and `port` values.
	SessionRedisClusterNodes []SamlPluginSessionRedisClusterNodes `json:"session_redis_cluster_nodes,omitempty"`
	// The Redis connection timeout in milliseconds.
	SessionRedisConnectTimeout *int64 `json:"session_redis_connect_timeout,omitempty"`
	// The Redis host IP.
	SessionRedisHost *string `default:"127.0.0.1" json:"session_redis_host"`
	// Password to use for Redis connection when the `redis` session storage is defined. If undefined, no auth commands are sent to Redis. This value is pulled from
	SessionRedisPassword *string `json:"session_redis_password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	SessionRedisPort *int64 `default:"6379" json:"session_redis_port"`
	// The Redis session key prefix.
	SessionRedisPrefix *string `json:"session_redis_prefix,omitempty"`
	// The Redis read timeout in milliseconds.
	SessionRedisReadTimeout *int64 `json:"session_redis_read_timeout,omitempty"`
	// The Redis send timeout in milliseconds.
	SessionRedisSendTimeout *int64 `json:"session_redis_send_timeout,omitempty"`
	// The SNI used for connecting to the Redis server.
	SessionRedisServerName *string `json:"session_redis_server_name,omitempty"`
	// The Redis unix socket path.
	SessionRedisSocket *string `json:"session_redis_socket,omitempty"`
	// Use SSL/TLS for the Redis connection.
	SessionRedisSsl *bool `default:"false" json:"session_redis_ssl"`
	// Verify the Redis server certificate.
	SessionRedisSslVerify *bool `default:"false" json:"session_redis_ssl_verify"`
	// Redis username if the `redis` session storage is defined and ACL authentication is desired.If undefined, ACL authentication will not be performed.  This requires Redis v6.0.0+. The username **cannot** be set to `default`.
	SessionRedisUsername *string `json:"session_redis_username,omitempty"`
	// Enables or disables persistent sessions
	SessionRemember *bool `default:"false" json:"session_remember"`
	// Persistent session absolute timeout in seconds.
	SessionRememberAbsoluteTimeout *float64 `default:"2592000" json:"session_remember_absolute_timeout"`
	// Persistent session cookie name
	SessionRememberCookieName *string `default:"remember" json:"session_remember_cookie_name"`
	// Persistent session rolling timeout in seconds.
	SessionRememberRollingTimeout *float64                           `default:"604800" json:"session_remember_rolling_timeout"`
	SessionRequestHeaders         []SamlPluginSessionRequestHeaders  `json:"session_request_headers,omitempty"`
	SessionResponseHeaders        []SamlPluginSessionResponseHeaders `json:"session_response_headers,omitempty"`
	// The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid.
	SessionRollingTimeout *float64 `default:"3600" json:"session_rolling_timeout"`
	// The session secret. This must be a random string of 32 characters from the base64 alphabet (letters, numbers, `/`, `_` and `+`). It is used as the secret key for encrypting session data as well as state information that is sent to the IdP in the authentication exchange.
	SessionSecret *string `json:"session_secret,omitempty"`
	// The session storage for session data: - `cookie`: stores session data with the session cookie. The session cannot be invalidated or revoked without changing the session secret, but is stateless, and doesn't require a database. - `memcached`: stores session data in memcached - `redis`: stores session data in Redis
	SessionStorage *SamlPluginSessionStorage `default:"cookie" json:"session_storage"`
	// Configures whether or not session metadata should be stored. This includes information about the active sessions for the `specific_audience` belonging to a specific subject.
	SessionStoreMetadata *bool `default:"false" json:"session_store_metadata"`
	// Enable signature validation for SAML responses.
	ValidateAssertionSignature *bool `default:"true" json:"validate_assertion_signature"`
}

func (s SamlPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SamlPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SamlPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *SamlPluginConfig) GetAssertionConsumerPath() *string {
	if o == nil {
		return nil
	}
	return o.AssertionConsumerPath
}

func (o *SamlPluginConfig) GetIdpCertificate() *string {
	if o == nil {
		return nil
	}
	return o.IdpCertificate
}

func (o *SamlPluginConfig) GetIdpSsoURL() *string {
	if o == nil {
		return nil
	}
	return o.IdpSsoURL
}

func (o *SamlPluginConfig) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *SamlPluginConfig) GetNameidFormat() *NameidFormat {
	if o == nil {
		return nil
	}
	return o.NameidFormat
}

func (o *SamlPluginConfig) GetRequestDigestAlgorithm() *RequestDigestAlgorithm {
	if o == nil {
		return nil
	}
	return o.RequestDigestAlgorithm
}

func (o *SamlPluginConfig) GetRequestSignatureAlgorithm() *RequestSignatureAlgorithm {
	if o == nil {
		return nil
	}
	return o.RequestSignatureAlgorithm
}

func (o *SamlPluginConfig) GetRequestSigningCertificate() *string {
	if o == nil {
		return nil
	}
	return o.RequestSigningCertificate
}

func (o *SamlPluginConfig) GetRequestSigningKey() *string {
	if o == nil {
		return nil
	}
	return o.RequestSigningKey
}

func (o *SamlPluginConfig) GetResponseDigestAlgorithm() *ResponseDigestAlgorithm {
	if o == nil {
		return nil
	}
	return o.ResponseDigestAlgorithm
}

func (o *SamlPluginConfig) GetResponseEncryptionKey() *string {
	if o == nil {
		return nil
	}
	return o.ResponseEncryptionKey
}

func (o *SamlPluginConfig) GetResponseSignatureAlgorithm() *ResponseSignatureAlgorithm {
	if o == nil {
		return nil
	}
	return o.ResponseSignatureAlgorithm
}

func (o *SamlPluginConfig) GetSessionAbsoluteTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionAbsoluteTimeout
}

func (o *SamlPluginConfig) GetSessionAudience() *string {
	if o == nil {
		return nil
	}
	return o.SessionAudience
}

func (o *SamlPluginConfig) GetSessionCookieDomain() *string {
	if o == nil {
		return nil
	}
	return o.SessionCookieDomain
}

func (o *SamlPluginConfig) GetSessionCookieHTTPOnly() *bool {
	if o == nil {
		return nil
	}
	return o.SessionCookieHTTPOnly
}

func (o *SamlPluginConfig) GetSessionCookieName() *string {
	if o == nil {
		return nil
	}
	return o.SessionCookieName
}

func (o *SamlPluginConfig) GetSessionCookiePath() *string {
	if o == nil {
		return nil
	}
	return o.SessionCookiePath
}

func (o *SamlPluginConfig) GetSessionCookieSameSite() *SamlPluginSessionCookieSameSite {
	if o == nil {
		return nil
	}
	return o.SessionCookieSameSite
}

func (o *SamlPluginConfig) GetSessionCookieSecure() *bool {
	if o == nil {
		return nil
	}
	return o.SessionCookieSecure
}

func (o *SamlPluginConfig) GetSessionEnforceSameSubject() *bool {
	if o == nil {
		return nil
	}
	return o.SessionEnforceSameSubject
}

func (o *SamlPluginConfig) GetSessionHashStorageKey() *bool {
	if o == nil {
		return nil
	}
	return o.SessionHashStorageKey
}

func (o *SamlPluginConfig) GetSessionHashSubject() *bool {
	if o == nil {
		return nil
	}
	return o.SessionHashSubject
}

func (o *SamlPluginConfig) GetSessionIdlingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionIdlingTimeout
}

func (o *SamlPluginConfig) GetSessionMemcachedHost() *string {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedHost
}

func (o *SamlPluginConfig) GetSessionMemcachedPort() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedPort
}

func (o *SamlPluginConfig) GetSessionMemcachedPrefix() *string {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedPrefix
}

func (o *SamlPluginConfig) GetSessionMemcachedSocket() *string {
	if o == nil {
		return nil
	}
	return o.SessionMemcachedSocket
}

func (o *SamlPluginConfig) GetSessionRedisClusterMaxRedirections() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionRedisClusterMaxRedirections
}

func (o *SamlPluginConfig) GetSessionRedisClusterNodes() []SamlPluginSessionRedisClusterNodes {
	if o == nil {
		return nil
	}
	return o.SessionRedisClusterNodes
}

func (o *SamlPluginConfig) GetSessionRedisConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionRedisConnectTimeout
}

func (o *SamlPluginConfig) GetSessionRedisHost() *string {
	if o == nil {
		return nil
	}
	return o.SessionRedisHost
}

func (o *SamlPluginConfig) GetSessionRedisPassword() *string {
	if o == nil {
		return nil
	}
	return o.SessionRedisPassword
}

func (o *SamlPluginConfig) GetSessionRedisPort() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionRedisPort
}

func (o *SamlPluginConfig) GetSessionRedisPrefix() *string {
	if o == nil {
		return nil
	}
	return o.SessionRedisPrefix
}

func (o *SamlPluginConfig) GetSessionRedisReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionRedisReadTimeout
}

func (o *SamlPluginConfig) GetSessionRedisSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionRedisSendTimeout
}

func (o *SamlPluginConfig) GetSessionRedisServerName() *string {
	if o == nil {
		return nil
	}
	return o.SessionRedisServerName
}

func (o *SamlPluginConfig) GetSessionRedisSocket() *string {
	if o == nil {
		return nil
	}
	return o.SessionRedisSocket
}

func (o *SamlPluginConfig) GetSessionRedisSsl() *bool {
	if o == nil {
		return nil
	}
	return o.SessionRedisSsl
}

func (o *SamlPluginConfig) GetSessionRedisSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SessionRedisSslVerify
}

func (o *SamlPluginConfig) GetSessionRedisUsername() *string {
	if o == nil {
		return nil
	}
	return o.SessionRedisUsername
}

func (o *SamlPluginConfig) GetSessionRemember() *bool {
	if o == nil {
		return nil
	}
	return o.SessionRemember
}

func (o *SamlPluginConfig) GetSessionRememberAbsoluteTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionRememberAbsoluteTimeout
}

func (o *SamlPluginConfig) GetSessionRememberCookieName() *string {
	if o == nil {
		return nil
	}
	return o.SessionRememberCookieName
}

func (o *SamlPluginConfig) GetSessionRememberRollingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionRememberRollingTimeout
}

func (o *SamlPluginConfig) GetSessionRequestHeaders() []SamlPluginSessionRequestHeaders {
	if o == nil {
		return nil
	}
	return o.SessionRequestHeaders
}

func (o *SamlPluginConfig) GetSessionResponseHeaders() []SamlPluginSessionResponseHeaders {
	if o == nil {
		return nil
	}
	return o.SessionResponseHeaders
}

func (o *SamlPluginConfig) GetSessionRollingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionRollingTimeout
}

func (o *SamlPluginConfig) GetSessionSecret() *string {
	if o == nil {
		return nil
	}
	return o.SessionSecret
}

func (o *SamlPluginConfig) GetSessionStorage() *SamlPluginSessionStorage {
	if o == nil {
		return nil
	}
	return o.SessionStorage
}

func (o *SamlPluginConfig) GetSessionStoreMetadata() *bool {
	if o == nil {
		return nil
	}
	return o.SessionStoreMetadata
}

func (o *SamlPluginConfig) GetValidateAssertionSignature() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateAssertionSignature
}

// SamlPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type SamlPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"saml" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []SamlPluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *SamlPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *SamlPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *SamlPluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64           `json:"created_at,omitempty"`
	ID        *string          `json:"id,omitempty"`
	Config    SamlPluginConfig `json:"config"`
}

func (s SamlPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SamlPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SamlPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SamlPlugin) GetName() string {
	return "saml"
}

func (o *SamlPlugin) GetProtocols() []SamlPluginProtocols {
	if o == nil {
		return []SamlPluginProtocols{}
	}
	return o.Protocols
}

func (o *SamlPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *SamlPlugin) GetConsumer() *SamlPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *SamlPlugin) GetRoute() *SamlPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *SamlPlugin) GetService() *SamlPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *SamlPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SamlPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SamlPlugin) GetConfig() SamlPluginConfig {
	if o == nil {
		return SamlPluginConfig{}
	}
	return o.Config
}
