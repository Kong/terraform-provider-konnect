// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type ClaimsToVerify string

const (
	ClaimsToVerifyExp ClaimsToVerify = "exp"
	ClaimsToVerifyNbf ClaimsToVerify = "nbf"
)

func (e ClaimsToVerify) ToPointer() *ClaimsToVerify {
	return &e
}
func (e *ClaimsToVerify) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "exp":
		fallthrough
	case "nbf":
		*e = ClaimsToVerify(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClaimsToVerify: %v", v)
	}
}

type JWTPluginConfig struct {
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
	Anonymous *string `json:"anonymous,omitempty"`
	// A list of registered claims (according to RFC 7519) that Kong can verify as well. Accepted values: one of exp or nbf.
	ClaimsToVerify []ClaimsToVerify `json:"claims_to_verify,omitempty"`
	// A list of cookie names that Kong will inspect to retrieve JWTs.
	CookieNames []string `json:"cookie_names,omitempty"`
	// A list of HTTP header names that Kong will inspect to retrieve JWTs.
	HeaderNames []string `json:"header_names,omitempty"`
	// The name of the claim in which the key identifying the secret must be passed. The plugin will attempt to read this claim from the JWT payload and the header, in that order.
	KeyClaimName *string `json:"key_claim_name,omitempty"`
	// A value between 0 and 31536000 (365 days) limiting the lifetime of the JWT to maximum_expiration seconds in the future.
	MaximumExpiration *float64 `json:"maximum_expiration,omitempty"`
	// A boolean value that indicates whether the plugin should run (and try to authenticate) on OPTIONS preflight requests. If set to false, then OPTIONS requests will always be allowed.
	RunOnPreflight *bool `json:"run_on_preflight,omitempty"`
	// If true, the plugin assumes the credential’s secret to be base64 encoded. You will need to create a base64-encoded secret for your Consumer, and sign your JWT with the original secret.
	SecretIsBase64 *bool `json:"secret_is_base64,omitempty"`
	// A list of querystring parameters that Kong will inspect to retrieve JWTs.
	URIParamNames []string `json:"uri_param_names,omitempty"`
}

func (o *JWTPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *JWTPluginConfig) GetClaimsToVerify() []ClaimsToVerify {
	if o == nil {
		return nil
	}
	return o.ClaimsToVerify
}

func (o *JWTPluginConfig) GetCookieNames() []string {
	if o == nil {
		return nil
	}
	return o.CookieNames
}

func (o *JWTPluginConfig) GetHeaderNames() []string {
	if o == nil {
		return nil
	}
	return o.HeaderNames
}

func (o *JWTPluginConfig) GetKeyClaimName() *string {
	if o == nil {
		return nil
	}
	return o.KeyClaimName
}

func (o *JWTPluginConfig) GetMaximumExpiration() *float64 {
	if o == nil {
		return nil
	}
	return o.MaximumExpiration
}

func (o *JWTPluginConfig) GetRunOnPreflight() *bool {
	if o == nil {
		return nil
	}
	return o.RunOnPreflight
}

func (o *JWTPluginConfig) GetSecretIsBase64() *bool {
	if o == nil {
		return nil
	}
	return o.SecretIsBase64
}

func (o *JWTPluginConfig) GetURIParamNames() []string {
	if o == nil {
		return nil
	}
	return o.URIParamNames
}

type JWTPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *JWTPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JWTPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *JWTPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JWTPluginOrdering struct {
	After  *JWTPluginAfter  `json:"after,omitempty"`
	Before *JWTPluginBefore `json:"before,omitempty"`
}

func (o *JWTPluginOrdering) GetAfter() *JWTPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *JWTPluginOrdering) GetBefore() *JWTPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type JWTPluginProtocols string

const (
	JWTPluginProtocolsGrpc           JWTPluginProtocols = "grpc"
	JWTPluginProtocolsGrpcs          JWTPluginProtocols = "grpcs"
	JWTPluginProtocolsHTTP           JWTPluginProtocols = "http"
	JWTPluginProtocolsHTTPS          JWTPluginProtocols = "https"
	JWTPluginProtocolsTCP            JWTPluginProtocols = "tcp"
	JWTPluginProtocolsTLS            JWTPluginProtocols = "tls"
	JWTPluginProtocolsTLSPassthrough JWTPluginProtocols = "tls_passthrough"
	JWTPluginProtocolsUDP            JWTPluginProtocols = "udp"
	JWTPluginProtocolsWs             JWTPluginProtocols = "ws"
	JWTPluginProtocolsWss            JWTPluginProtocols = "wss"
)

func (e JWTPluginProtocols) ToPointer() *JWTPluginProtocols {
	return &e
}
func (e *JWTPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = JWTPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JWTPluginProtocols: %v", v)
	}
}

// JWTPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type JWTPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *JWTPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JWTPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *JWTPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JWTPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type JWTPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *JWTPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JWTPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type JWTPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *JWTPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JWTPlugin struct {
	Config *JWTPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool              `json:"enabled,omitempty"`
	ID           *string            `json:"id,omitempty"`
	InstanceName *string            `json:"instance_name,omitempty"`
	name         *string            `const:"jwt" json:"name,omitempty"`
	Ordering     *JWTPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []JWTPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *JWTPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *JWTPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *JWTPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *JWTPluginService `json:"service,omitempty"`
}

func (j JWTPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JWTPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JWTPlugin) GetConfig() *JWTPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *JWTPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *JWTPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *JWTPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JWTPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *JWTPlugin) GetName() *string {
	return types.String("jwt")
}

func (o *JWTPlugin) GetOrdering() *JWTPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *JWTPlugin) GetProtocols() []JWTPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *JWTPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *JWTPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *JWTPlugin) GetConsumer() *JWTPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *JWTPlugin) GetConsumerGroup() *JWTPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *JWTPlugin) GetRoute() *JWTPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *JWTPlugin) GetService() *JWTPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
