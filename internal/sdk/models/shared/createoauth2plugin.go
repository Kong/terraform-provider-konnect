// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateOauth2PluginPkce - Specifies a mode of how the Proof Key for Code Exchange (PKCE) should be handled by the plugin.
type CreateOauth2PluginPkce string

const (
	CreateOauth2PluginPkceNone   CreateOauth2PluginPkce = "none"
	CreateOauth2PluginPkceLax    CreateOauth2PluginPkce = "lax"
	CreateOauth2PluginPkceStrict CreateOauth2PluginPkce = "strict"
)

func (e CreateOauth2PluginPkce) ToPointer() *CreateOauth2PluginPkce {
	return &e
}
func (e *CreateOauth2PluginPkce) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "lax":
		fallthrough
	case "strict":
		*e = CreateOauth2PluginPkce(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOauth2PluginPkce: %v", v)
	}
}

type CreateOauth2PluginConfig struct {
	// Accepts HTTPs requests that have already been terminated by a proxy or load balancer.
	AcceptHTTPIfAlreadyTerminated *bool `json:"accept_http_if_already_terminated,omitempty"`
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
	Anonymous *string `json:"anonymous,omitempty"`
	// The name of the header that is supposed to carry the access token.
	AuthHeaderName *string `json:"auth_header_name,omitempty"`
	// An optional boolean value to enable the three-legged Authorization Code flow (RFC 6742 Section 4.1).
	EnableAuthorizationCode *bool `json:"enable_authorization_code,omitempty"`
	// An optional boolean value to enable the Client Credentials Grant flow (RFC 6742 Section 4.4).
	EnableClientCredentials *bool `json:"enable_client_credentials,omitempty"`
	// An optional boolean value to enable the Implicit Grant flow which allows to provision a token as a result of the authorization process (RFC 6742 Section 4.2).
	EnableImplicitGrant *bool `json:"enable_implicit_grant,omitempty"`
	// An optional boolean value to enable the Resource Owner Password Credentials Grant flow (RFC 6742 Section 4.3).
	EnablePasswordGrant *bool `json:"enable_password_grant,omitempty"`
	// An optional boolean value that allows using the same OAuth credentials generated by the plugin with any other service whose OAuth 2.0 plugin configuration also has `config.global_credentials=true`.
	GlobalCredentials *bool `json:"global_credentials,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service.
	HideCredentials *bool `json:"hide_credentials,omitempty"`
	// An optional boolean value telling the plugin to require at least one `scope` to be authorized by the end user.
	MandatoryScope         *bool `json:"mandatory_scope,omitempty"`
	PersistentRefreshToken *bool `json:"persistent_refresh_token,omitempty"`
	// Specifies a mode of how the Proof Key for Code Exchange (PKCE) should be handled by the plugin.
	Pkce *CreateOauth2PluginPkce `json:"pkce,omitempty"`
	// The unique key the plugin has generated when it has been added to the Service.
	ProvisionKey *string `json:"provision_key,omitempty"`
	// When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
	Realm *string `json:"realm,omitempty"`
	// Time-to-live value for data
	RefreshTokenTTL *float64 `json:"refresh_token_ttl,omitempty"`
	// An optional boolean value that indicates whether an OAuth refresh token is reused when refreshing an access token.
	ReuseRefreshToken *bool `json:"reuse_refresh_token,omitempty"`
	// Describes an array of scope names that will be available to the end user. If `mandatory_scope` is set to `true`, then `scopes` are required.
	Scopes []string `json:"scopes,omitempty"`
	// An optional integer value telling the plugin how many seconds a token should last, after which the client will need to refresh the token. Set to `0` to disable the expiration.
	TokenExpiration *float64 `json:"token_expiration,omitempty"`
}

func (o *CreateOauth2PluginConfig) GetAcceptHTTPIfAlreadyTerminated() *bool {
	if o == nil {
		return nil
	}
	return o.AcceptHTTPIfAlreadyTerminated
}

func (o *CreateOauth2PluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *CreateOauth2PluginConfig) GetAuthHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.AuthHeaderName
}

func (o *CreateOauth2PluginConfig) GetEnableAuthorizationCode() *bool {
	if o == nil {
		return nil
	}
	return o.EnableAuthorizationCode
}

func (o *CreateOauth2PluginConfig) GetEnableClientCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.EnableClientCredentials
}

func (o *CreateOauth2PluginConfig) GetEnableImplicitGrant() *bool {
	if o == nil {
		return nil
	}
	return o.EnableImplicitGrant
}

func (o *CreateOauth2PluginConfig) GetEnablePasswordGrant() *bool {
	if o == nil {
		return nil
	}
	return o.EnablePasswordGrant
}

func (o *CreateOauth2PluginConfig) GetGlobalCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.GlobalCredentials
}

func (o *CreateOauth2PluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *CreateOauth2PluginConfig) GetMandatoryScope() *bool {
	if o == nil {
		return nil
	}
	return o.MandatoryScope
}

func (o *CreateOauth2PluginConfig) GetPersistentRefreshToken() *bool {
	if o == nil {
		return nil
	}
	return o.PersistentRefreshToken
}

func (o *CreateOauth2PluginConfig) GetPkce() *CreateOauth2PluginPkce {
	if o == nil {
		return nil
	}
	return o.Pkce
}

func (o *CreateOauth2PluginConfig) GetProvisionKey() *string {
	if o == nil {
		return nil
	}
	return o.ProvisionKey
}

func (o *CreateOauth2PluginConfig) GetRealm() *string {
	if o == nil {
		return nil
	}
	return o.Realm
}

func (o *CreateOauth2PluginConfig) GetRefreshTokenTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.RefreshTokenTTL
}

func (o *CreateOauth2PluginConfig) GetReuseRefreshToken() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseRefreshToken
}

func (o *CreateOauth2PluginConfig) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *CreateOauth2PluginConfig) GetTokenExpiration() *float64 {
	if o == nil {
		return nil
	}
	return o.TokenExpiration
}

type CreateOauth2PluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateOauth2PluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateOauth2PluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateOauth2PluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateOauth2PluginOrdering struct {
	After  *CreateOauth2PluginAfter  `json:"after,omitempty"`
	Before *CreateOauth2PluginBefore `json:"before,omitempty"`
}

func (o *CreateOauth2PluginOrdering) GetAfter() *CreateOauth2PluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateOauth2PluginOrdering) GetBefore() *CreateOauth2PluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateOauth2PluginProtocols string

const (
	CreateOauth2PluginProtocolsGrpc           CreateOauth2PluginProtocols = "grpc"
	CreateOauth2PluginProtocolsGrpcs          CreateOauth2PluginProtocols = "grpcs"
	CreateOauth2PluginProtocolsHTTP           CreateOauth2PluginProtocols = "http"
	CreateOauth2PluginProtocolsHTTPS          CreateOauth2PluginProtocols = "https"
	CreateOauth2PluginProtocolsTCP            CreateOauth2PluginProtocols = "tcp"
	CreateOauth2PluginProtocolsTLS            CreateOauth2PluginProtocols = "tls"
	CreateOauth2PluginProtocolsTLSPassthrough CreateOauth2PluginProtocols = "tls_passthrough"
	CreateOauth2PluginProtocolsUDP            CreateOauth2PluginProtocols = "udp"
	CreateOauth2PluginProtocolsWs             CreateOauth2PluginProtocols = "ws"
	CreateOauth2PluginProtocolsWss            CreateOauth2PluginProtocols = "wss"
)

func (e CreateOauth2PluginProtocols) ToPointer() *CreateOauth2PluginProtocols {
	return &e
}
func (e *CreateOauth2PluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateOauth2PluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOauth2PluginProtocols: %v", v)
	}
}

// CreateOauth2PluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateOauth2PluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOauth2PluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateOauth2PluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOauth2PluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateOauth2PluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateOauth2PluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOauth2PluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateOauth2PluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateOauth2PluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOauth2PluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateOauth2Plugin struct {
	Config *CreateOauth2PluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                       `json:"enabled,omitempty"`
	InstanceName *string                     `json:"instance_name,omitempty"`
	name         *string                     `const:"oauth2" json:"name,omitempty"`
	Ordering     *CreateOauth2PluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateOauth2PluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateOauth2PluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateOauth2PluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateOauth2PluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateOauth2PluginService `json:"service,omitempty"`
}

func (c CreateOauth2Plugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateOauth2Plugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateOauth2Plugin) GetConfig() *CreateOauth2PluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateOauth2Plugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateOauth2Plugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateOauth2Plugin) GetName() *string {
	return types.String("oauth2")
}

func (o *CreateOauth2Plugin) GetOrdering() *CreateOauth2PluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateOauth2Plugin) GetProtocols() []CreateOauth2PluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateOauth2Plugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateOauth2Plugin) GetConsumer() *CreateOauth2PluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateOauth2Plugin) GetConsumerGroup() *CreateOauth2PluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateOauth2Plugin) GetRoute() *CreateOauth2PluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateOauth2Plugin) GetService() *CreateOauth2PluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
