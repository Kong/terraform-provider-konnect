// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CookieSameSite - Determines whether and how a cookie may be sent with cross-site requests.
type CookieSameSite string

const (
	CookieSameSiteStrict  CookieSameSite = "Strict"
	CookieSameSiteLax     CookieSameSite = "Lax"
	CookieSameSiteNone    CookieSameSite = "None"
	CookieSameSiteDefault CookieSameSite = "Default"
)

func (e CookieSameSite) ToPointer() *CookieSameSite {
	return &e
}
func (e *CookieSameSite) UnmarshalJSON(data []byte) error {
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
		*e = CookieSameSite(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CookieSameSite: %v", v)
	}
}

type SessionPluginLogoutMethods string

const (
	SessionPluginLogoutMethodsGet    SessionPluginLogoutMethods = "GET"
	SessionPluginLogoutMethodsPost   SessionPluginLogoutMethods = "POST"
	SessionPluginLogoutMethodsDelete SessionPluginLogoutMethods = "DELETE"
)

func (e SessionPluginLogoutMethods) ToPointer() *SessionPluginLogoutMethods {
	return &e
}
func (e *SessionPluginLogoutMethods) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "DELETE":
		*e = SessionPluginLogoutMethods(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SessionPluginLogoutMethods: %v", v)
	}
}

type RequestHeaders string

const (
	RequestHeadersID              RequestHeaders = "id"
	RequestHeadersAudience        RequestHeaders = "audience"
	RequestHeadersSubject         RequestHeaders = "subject"
	RequestHeadersTimeout         RequestHeaders = "timeout"
	RequestHeadersIdlingTimeout   RequestHeaders = "idling-timeout"
	RequestHeadersRollingTimeout  RequestHeaders = "rolling-timeout"
	RequestHeadersAbsoluteTimeout RequestHeaders = "absolute-timeout"
)

func (e RequestHeaders) ToPointer() *RequestHeaders {
	return &e
}
func (e *RequestHeaders) UnmarshalJSON(data []byte) error {
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
		*e = RequestHeaders(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestHeaders: %v", v)
	}
}

type SessionPluginResponseHeaders string

const (
	SessionPluginResponseHeadersID              SessionPluginResponseHeaders = "id"
	SessionPluginResponseHeadersAudience        SessionPluginResponseHeaders = "audience"
	SessionPluginResponseHeadersSubject         SessionPluginResponseHeaders = "subject"
	SessionPluginResponseHeadersTimeout         SessionPluginResponseHeaders = "timeout"
	SessionPluginResponseHeadersIdlingTimeout   SessionPluginResponseHeaders = "idling-timeout"
	SessionPluginResponseHeadersRollingTimeout  SessionPluginResponseHeaders = "rolling-timeout"
	SessionPluginResponseHeadersAbsoluteTimeout SessionPluginResponseHeaders = "absolute-timeout"
)

func (e SessionPluginResponseHeaders) ToPointer() *SessionPluginResponseHeaders {
	return &e
}
func (e *SessionPluginResponseHeaders) UnmarshalJSON(data []byte) error {
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
		*e = SessionPluginResponseHeaders(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SessionPluginResponseHeaders: %v", v)
	}
}

// SessionPluginStorage - Determines where the session data is stored. `kong`: Stores encrypted session data into Kong's current database strategy; the cookie will not contain any session data. `cookie`: Stores encrypted session data within the cookie itself.
type SessionPluginStorage string

const (
	SessionPluginStorageCookie SessionPluginStorage = "cookie"
	SessionPluginStorageKong   SessionPluginStorage = "kong"
)

func (e SessionPluginStorage) ToPointer() *SessionPluginStorage {
	return &e
}
func (e *SessionPluginStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cookie":
		fallthrough
	case "kong":
		*e = SessionPluginStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SessionPluginStorage: %v", v)
	}
}

type SessionPluginConfig struct {
	// The session cookie absolute timeout, in seconds. Specifies how long the session can be used until it is no longer valid.
	AbsoluteTimeout *float64 `json:"absolute_timeout,omitempty"`
	// The session audience, which is the intended target application. For example `"my-application"`.
	Audience *string `json:"audience,omitempty"`
	// The domain with which the cookie is intended to be exchanged.
	CookieDomain *string `json:"cookie_domain,omitempty"`
	// Applies the `HttpOnly` tag so that the cookie is sent only to a server.
	CookieHTTPOnly *bool `json:"cookie_http_only,omitempty"`
	// The name of the cookie.
	CookieName *string `json:"cookie_name,omitempty"`
	// The resource in the host where the cookie is available.
	CookiePath *string `json:"cookie_path,omitempty"`
	// Determines whether and how a cookie may be sent with cross-site requests.
	CookieSameSite *CookieSameSite `json:"cookie_same_site,omitempty"`
	// Applies the Secure directive so that the cookie may be sent to the server only with an encrypted request over the HTTPS protocol.
	CookieSecure *bool `json:"cookie_secure,omitempty"`
	// The session cookie idle time, in seconds.
	IdlingTimeout *float64 `json:"idling_timeout,omitempty"`
	// A set of HTTP methods that the plugin will respond to.
	LogoutMethods []SessionPluginLogoutMethods `json:"logout_methods,omitempty"`
	// The POST argument passed to logout requests. Do not change this property.
	LogoutPostArg *string `json:"logout_post_arg,omitempty"`
	// The query argument passed to logout requests.
	LogoutQueryArg    *string `json:"logout_query_arg,omitempty"`
	ReadBodyForLogout *bool   `json:"read_body_for_logout,omitempty"`
	// Enables or disables persistent sessions.
	Remember *bool `json:"remember,omitempty"`
	// The persistent session absolute timeout limit, in seconds.
	RememberAbsoluteTimeout *float64 `json:"remember_absolute_timeout,omitempty"`
	// Persistent session cookie name. Use with the `remember` configuration parameter.
	RememberCookieName *string `json:"remember_cookie_name,omitempty"`
	// The persistent session rolling timeout window, in seconds.
	RememberRollingTimeout *float64 `json:"remember_rolling_timeout,omitempty"`
	// List of information to include, as headers, in the response to the downstream.
	RequestHeaders []RequestHeaders `json:"request_headers,omitempty"`
	// List of information to include, as headers, in the response to the downstream.
	ResponseHeaders []SessionPluginResponseHeaders `json:"response_headers,omitempty"`
	// The session cookie rolling timeout, in seconds. Specifies how long the session can be used until it needs to be renewed.
	RollingTimeout *float64 `json:"rolling_timeout,omitempty"`
	// The secret that is used in keyed HMAC generation.
	Secret *string `json:"secret,omitempty"`
	// The duration, in seconds, after which an old cookie is discarded, starting from the moment when the session becomes outdated and is replaced by a new one.
	StaleTTL *float64 `json:"stale_ttl,omitempty"`
	// Determines where the session data is stored. `kong`: Stores encrypted session data into Kong's current database strategy; the cookie will not contain any session data. `cookie`: Stores encrypted session data within the cookie itself.
	Storage *SessionPluginStorage `json:"storage,omitempty"`
}

func (o *SessionPluginConfig) GetAbsoluteTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.AbsoluteTimeout
}

func (o *SessionPluginConfig) GetAudience() *string {
	if o == nil {
		return nil
	}
	return o.Audience
}

func (o *SessionPluginConfig) GetCookieDomain() *string {
	if o == nil {
		return nil
	}
	return o.CookieDomain
}

func (o *SessionPluginConfig) GetCookieHTTPOnly() *bool {
	if o == nil {
		return nil
	}
	return o.CookieHTTPOnly
}

func (o *SessionPluginConfig) GetCookieName() *string {
	if o == nil {
		return nil
	}
	return o.CookieName
}

func (o *SessionPluginConfig) GetCookiePath() *string {
	if o == nil {
		return nil
	}
	return o.CookiePath
}

func (o *SessionPluginConfig) GetCookieSameSite() *CookieSameSite {
	if o == nil {
		return nil
	}
	return o.CookieSameSite
}

func (o *SessionPluginConfig) GetCookieSecure() *bool {
	if o == nil {
		return nil
	}
	return o.CookieSecure
}

func (o *SessionPluginConfig) GetIdlingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.IdlingTimeout
}

func (o *SessionPluginConfig) GetLogoutMethods() []SessionPluginLogoutMethods {
	if o == nil {
		return nil
	}
	return o.LogoutMethods
}

func (o *SessionPluginConfig) GetLogoutPostArg() *string {
	if o == nil {
		return nil
	}
	return o.LogoutPostArg
}

func (o *SessionPluginConfig) GetLogoutQueryArg() *string {
	if o == nil {
		return nil
	}
	return o.LogoutQueryArg
}

func (o *SessionPluginConfig) GetReadBodyForLogout() *bool {
	if o == nil {
		return nil
	}
	return o.ReadBodyForLogout
}

func (o *SessionPluginConfig) GetRemember() *bool {
	if o == nil {
		return nil
	}
	return o.Remember
}

func (o *SessionPluginConfig) GetRememberAbsoluteTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RememberAbsoluteTimeout
}

func (o *SessionPluginConfig) GetRememberCookieName() *string {
	if o == nil {
		return nil
	}
	return o.RememberCookieName
}

func (o *SessionPluginConfig) GetRememberRollingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RememberRollingTimeout
}

func (o *SessionPluginConfig) GetRequestHeaders() []RequestHeaders {
	if o == nil {
		return nil
	}
	return o.RequestHeaders
}

func (o *SessionPluginConfig) GetResponseHeaders() []SessionPluginResponseHeaders {
	if o == nil {
		return nil
	}
	return o.ResponseHeaders
}

func (o *SessionPluginConfig) GetRollingTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RollingTimeout
}

func (o *SessionPluginConfig) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *SessionPluginConfig) GetStaleTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.StaleTTL
}

func (o *SessionPluginConfig) GetStorage() *SessionPluginStorage {
	if o == nil {
		return nil
	}
	return o.Storage
}

type SessionPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *SessionPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type SessionPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *SessionPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type SessionPluginOrdering struct {
	After  *SessionPluginAfter  `json:"after,omitempty"`
	Before *SessionPluginBefore `json:"before,omitempty"`
}

func (o *SessionPluginOrdering) GetAfter() *SessionPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *SessionPluginOrdering) GetBefore() *SessionPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type SessionPluginProtocols string

const (
	SessionPluginProtocolsGrpc           SessionPluginProtocols = "grpc"
	SessionPluginProtocolsGrpcs          SessionPluginProtocols = "grpcs"
	SessionPluginProtocolsHTTP           SessionPluginProtocols = "http"
	SessionPluginProtocolsHTTPS          SessionPluginProtocols = "https"
	SessionPluginProtocolsTCP            SessionPluginProtocols = "tcp"
	SessionPluginProtocolsTLS            SessionPluginProtocols = "tls"
	SessionPluginProtocolsTLSPassthrough SessionPluginProtocols = "tls_passthrough"
	SessionPluginProtocolsUDP            SessionPluginProtocols = "udp"
	SessionPluginProtocolsWs             SessionPluginProtocols = "ws"
	SessionPluginProtocolsWss            SessionPluginProtocols = "wss"
)

func (e SessionPluginProtocols) ToPointer() *SessionPluginProtocols {
	return &e
}
func (e *SessionPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = SessionPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SessionPluginProtocols: %v", v)
	}
}

// SessionPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type SessionPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *SessionPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type SessionPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *SessionPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SessionPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type SessionPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *SessionPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SessionPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type SessionPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *SessionPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type SessionPlugin struct {
	Config *SessionPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                  `json:"enabled,omitempty"`
	ID           *string                `json:"id,omitempty"`
	InstanceName *string                `json:"instance_name,omitempty"`
	name         *string                `const:"session" json:"name,omitempty"`
	Ordering     *SessionPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []SessionPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *SessionPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *SessionPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *SessionPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *SessionPluginService `json:"service,omitempty"`
}

func (s SessionPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SessionPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SessionPlugin) GetConfig() *SessionPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *SessionPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SessionPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SessionPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SessionPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *SessionPlugin) GetName() *string {
	return types.String("session")
}

func (o *SessionPlugin) GetOrdering() *SessionPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *SessionPlugin) GetProtocols() []SessionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *SessionPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *SessionPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *SessionPlugin) GetConsumer() *SessionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *SessionPlugin) GetConsumerGroup() *SessionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *SessionPlugin) GetRoute() *SessionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *SessionPlugin) GetService() *SessionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
