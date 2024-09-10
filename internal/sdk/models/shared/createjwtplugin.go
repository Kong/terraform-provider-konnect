// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateJWTPluginClaimsToVerify string

const (
	CreateJWTPluginClaimsToVerifyExp CreateJWTPluginClaimsToVerify = "exp"
	CreateJWTPluginClaimsToVerifyNbf CreateJWTPluginClaimsToVerify = "nbf"
)

func (e CreateJWTPluginClaimsToVerify) ToPointer() *CreateJWTPluginClaimsToVerify {
	return &e
}
func (e *CreateJWTPluginClaimsToVerify) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "exp":
		fallthrough
	case "nbf":
		*e = CreateJWTPluginClaimsToVerify(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateJWTPluginClaimsToVerify: %v", v)
	}
}

type CreateJWTPluginConfig struct {
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
	Anonymous *string `json:"anonymous,omitempty"`
	// A list of registered claims (according to RFC 7519) that Kong can verify as well. Accepted values: one of exp or nbf.
	ClaimsToVerify []CreateJWTPluginClaimsToVerify `json:"claims_to_verify,omitempty"`
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

func (o *CreateJWTPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *CreateJWTPluginConfig) GetClaimsToVerify() []CreateJWTPluginClaimsToVerify {
	if o == nil {
		return nil
	}
	return o.ClaimsToVerify
}

func (o *CreateJWTPluginConfig) GetCookieNames() []string {
	if o == nil {
		return nil
	}
	return o.CookieNames
}

func (o *CreateJWTPluginConfig) GetHeaderNames() []string {
	if o == nil {
		return nil
	}
	return o.HeaderNames
}

func (o *CreateJWTPluginConfig) GetKeyClaimName() *string {
	if o == nil {
		return nil
	}
	return o.KeyClaimName
}

func (o *CreateJWTPluginConfig) GetMaximumExpiration() *float64 {
	if o == nil {
		return nil
	}
	return o.MaximumExpiration
}

func (o *CreateJWTPluginConfig) GetRunOnPreflight() *bool {
	if o == nil {
		return nil
	}
	return o.RunOnPreflight
}

func (o *CreateJWTPluginConfig) GetSecretIsBase64() *bool {
	if o == nil {
		return nil
	}
	return o.SecretIsBase64
}

func (o *CreateJWTPluginConfig) GetURIParamNames() []string {
	if o == nil {
		return nil
	}
	return o.URIParamNames
}

type CreateJWTPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateJWTPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateJWTPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateJWTPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateJWTPluginOrdering struct {
	After  *CreateJWTPluginAfter  `json:"after,omitempty"`
	Before *CreateJWTPluginBefore `json:"before,omitempty"`
}

func (o *CreateJWTPluginOrdering) GetAfter() *CreateJWTPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateJWTPluginOrdering) GetBefore() *CreateJWTPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateJWTPluginProtocols string

const (
	CreateJWTPluginProtocolsGrpc           CreateJWTPluginProtocols = "grpc"
	CreateJWTPluginProtocolsGrpcs          CreateJWTPluginProtocols = "grpcs"
	CreateJWTPluginProtocolsHTTP           CreateJWTPluginProtocols = "http"
	CreateJWTPluginProtocolsHTTPS          CreateJWTPluginProtocols = "https"
	CreateJWTPluginProtocolsTCP            CreateJWTPluginProtocols = "tcp"
	CreateJWTPluginProtocolsTLS            CreateJWTPluginProtocols = "tls"
	CreateJWTPluginProtocolsTLSPassthrough CreateJWTPluginProtocols = "tls_passthrough"
	CreateJWTPluginProtocolsUDP            CreateJWTPluginProtocols = "udp"
	CreateJWTPluginProtocolsWs             CreateJWTPluginProtocols = "ws"
	CreateJWTPluginProtocolsWss            CreateJWTPluginProtocols = "wss"
)

func (e CreateJWTPluginProtocols) ToPointer() *CreateJWTPluginProtocols {
	return &e
}
func (e *CreateJWTPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateJWTPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateJWTPluginProtocols: %v", v)
	}
}

// CreateJWTPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateJWTPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJWTPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateJWTPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJWTPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJWTPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateJWTPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJWTPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJWTPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateJWTPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJWTPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateJWTPlugin struct {
	Config *CreateJWTPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                    `json:"enabled,omitempty"`
	InstanceName *string                  `json:"instance_name,omitempty"`
	name         *string                  `const:"jwt" json:"name,omitempty"`
	Ordering     *CreateJWTPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateJWTPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateJWTPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateJWTPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateJWTPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateJWTPluginService `json:"service,omitempty"`
}

func (c CreateJWTPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateJWTPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateJWTPlugin) GetConfig() *CreateJWTPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateJWTPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateJWTPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateJWTPlugin) GetName() *string {
	return types.String("jwt")
}

func (o *CreateJWTPlugin) GetOrdering() *CreateJWTPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateJWTPlugin) GetProtocols() []CreateJWTPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateJWTPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateJWTPlugin) GetConsumer() *CreateJWTPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateJWTPlugin) GetConsumerGroup() *CreateJWTPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateJWTPlugin) GetRoute() *CreateJWTPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateJWTPlugin) GetService() *CreateJWTPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
