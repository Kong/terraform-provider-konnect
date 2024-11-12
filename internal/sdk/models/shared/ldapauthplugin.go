// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type LdapAuthPluginConfig struct {
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure `4xx`.
	Anonymous *string `json:"anonymous,omitempty"`
	// Attribute to be used to search the user; e.g. cn
	Attribute *string `json:"attribute,omitempty"`
	// Base DN as the starting point for the search; e.g., dc=example,dc=com
	BaseDn *string `json:"base_dn,omitempty"`
	// Cache expiry time in seconds.
	CacheTTL *float64 `json:"cache_ttl,omitempty"`
	// An optional string to use as part of the Authorization header
	HeaderType *string `json:"header_type,omitempty"`
	// An optional boolean value telling the plugin to hide the credential to the upstream server. It will be removed by Kong before proxying the request.
	HideCredentials *bool `json:"hide_credentials,omitempty"`
	// An optional value in milliseconds that defines how long an idle connection to LDAP server will live before being closed.
	Keepalive *float64 `json:"keepalive,omitempty"`
	// A string representing a host name, such as example.com.
	LdapHost *string `json:"ldap_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	LdapPort *int64 `json:"ldap_port,omitempty"`
	// Set to `true` to connect using the LDAPS protocol (LDAP over TLS).  When `ldaps` is configured, you must use port 636. If the `ldap` setting is enabled, ensure the `start_tls` setting is disabled.
	Ldaps *bool `json:"ldaps,omitempty"`
	// When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
	Realm *string `json:"realm,omitempty"`
	// Set it to `true` to issue StartTLS (Transport Layer Security) extended operation over `ldap` connection. If the `start_tls` setting is enabled, ensure the `ldaps` setting is disabled.
	StartTLS *bool `json:"start_tls,omitempty"`
	// An optional timeout in milliseconds when waiting for connection with LDAP server.
	Timeout *float64 `json:"timeout,omitempty"`
	// Set to `true` to authenticate LDAP server. The server certificate will be verified according to the CA certificates specified by the `lua_ssl_trusted_certificate` directive.
	VerifyLdapHost *bool `json:"verify_ldap_host,omitempty"`
}

func (o *LdapAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *LdapAuthPluginConfig) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *LdapAuthPluginConfig) GetBaseDn() *string {
	if o == nil {
		return nil
	}
	return o.BaseDn
}

func (o *LdapAuthPluginConfig) GetCacheTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *LdapAuthPluginConfig) GetHeaderType() *string {
	if o == nil {
		return nil
	}
	return o.HeaderType
}

func (o *LdapAuthPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *LdapAuthPluginConfig) GetKeepalive() *float64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *LdapAuthPluginConfig) GetLdapHost() *string {
	if o == nil {
		return nil
	}
	return o.LdapHost
}

func (o *LdapAuthPluginConfig) GetLdapPort() *int64 {
	if o == nil {
		return nil
	}
	return o.LdapPort
}

func (o *LdapAuthPluginConfig) GetLdaps() *bool {
	if o == nil {
		return nil
	}
	return o.Ldaps
}

func (o *LdapAuthPluginConfig) GetRealm() *string {
	if o == nil {
		return nil
	}
	return o.Realm
}

func (o *LdapAuthPluginConfig) GetStartTLS() *bool {
	if o == nil {
		return nil
	}
	return o.StartTLS
}

func (o *LdapAuthPluginConfig) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *LdapAuthPluginConfig) GetVerifyLdapHost() *bool {
	if o == nil {
		return nil
	}
	return o.VerifyLdapHost
}

// LdapAuthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type LdapAuthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *LdapAuthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type LdapAuthPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *LdapAuthPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type LdapAuthPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *LdapAuthPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type LdapAuthPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *LdapAuthPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type LdapAuthPluginOrdering struct {
	After  *LdapAuthPluginAfter  `json:"after,omitempty"`
	Before *LdapAuthPluginBefore `json:"before,omitempty"`
}

func (o *LdapAuthPluginOrdering) GetAfter() *LdapAuthPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *LdapAuthPluginOrdering) GetBefore() *LdapAuthPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type LdapAuthPluginProtocols string

const (
	LdapAuthPluginProtocolsGrpc           LdapAuthPluginProtocols = "grpc"
	LdapAuthPluginProtocolsGrpcs          LdapAuthPluginProtocols = "grpcs"
	LdapAuthPluginProtocolsHTTP           LdapAuthPluginProtocols = "http"
	LdapAuthPluginProtocolsHTTPS          LdapAuthPluginProtocols = "https"
	LdapAuthPluginProtocolsTCP            LdapAuthPluginProtocols = "tcp"
	LdapAuthPluginProtocolsTLS            LdapAuthPluginProtocols = "tls"
	LdapAuthPluginProtocolsTLSPassthrough LdapAuthPluginProtocols = "tls_passthrough"
	LdapAuthPluginProtocolsUDP            LdapAuthPluginProtocols = "udp"
	LdapAuthPluginProtocolsWs             LdapAuthPluginProtocols = "ws"
	LdapAuthPluginProtocolsWss            LdapAuthPluginProtocols = "wss"
)

func (e LdapAuthPluginProtocols) ToPointer() *LdapAuthPluginProtocols {
	return &e
}
func (e *LdapAuthPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = LdapAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LdapAuthPluginProtocols: %v", v)
	}
}

// LdapAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type LdapAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *LdapAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// LdapAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type LdapAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *LdapAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// LdapAuthPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type LdapAuthPlugin struct {
	Config LdapAuthPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *LdapAuthPluginConsumer      `json:"consumer"`
	ConsumerGroup *LdapAuthPluginConsumerGroup `json:"consumer_group"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                   `json:"enabled,omitempty"`
	ID           *string                 `json:"id,omitempty"`
	InstanceName *string                 `json:"instance_name,omitempty"`
	name         string                  `const:"ldap-auth" json:"name"`
	Ordering     *LdapAuthPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []LdapAuthPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *LdapAuthPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *LdapAuthPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (l LdapAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LdapAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LdapAuthPlugin) GetConfig() LdapAuthPluginConfig {
	if o == nil {
		return LdapAuthPluginConfig{}
	}
	return o.Config
}

func (o *LdapAuthPlugin) GetConsumer() *LdapAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *LdapAuthPlugin) GetConsumerGroup() *LdapAuthPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *LdapAuthPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LdapAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *LdapAuthPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LdapAuthPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *LdapAuthPlugin) GetName() string {
	return "ldap-auth"
}

func (o *LdapAuthPlugin) GetOrdering() *LdapAuthPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *LdapAuthPlugin) GetProtocols() []LdapAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *LdapAuthPlugin) GetRoute() *LdapAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *LdapAuthPlugin) GetService() *LdapAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *LdapAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LdapAuthPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// LdapAuthPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type LdapAuthPluginInput struct {
	Config LdapAuthPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *LdapAuthPluginConsumer      `json:"consumer"`
	ConsumerGroup *LdapAuthPluginConsumerGroup `json:"consumer_group"`
	// Whether the plugin is applied.
	Enabled      *bool                   `json:"enabled,omitempty"`
	ID           *string                 `json:"id,omitempty"`
	InstanceName *string                 `json:"instance_name,omitempty"`
	name         string                  `const:"ldap-auth" json:"name"`
	Ordering     *LdapAuthPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []LdapAuthPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *LdapAuthPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *LdapAuthPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (l LdapAuthPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LdapAuthPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LdapAuthPluginInput) GetConfig() LdapAuthPluginConfig {
	if o == nil {
		return LdapAuthPluginConfig{}
	}
	return o.Config
}

func (o *LdapAuthPluginInput) GetConsumer() *LdapAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *LdapAuthPluginInput) GetConsumerGroup() *LdapAuthPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *LdapAuthPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *LdapAuthPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LdapAuthPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *LdapAuthPluginInput) GetName() string {
	return "ldap-auth"
}

func (o *LdapAuthPluginInput) GetOrdering() *LdapAuthPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *LdapAuthPluginInput) GetProtocols() []LdapAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *LdapAuthPluginInput) GetRoute() *LdapAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *LdapAuthPluginInput) GetService() *LdapAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *LdapAuthPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
