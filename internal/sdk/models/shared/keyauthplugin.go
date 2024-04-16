// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type KeyAuthPluginProtocols string

const (
	KeyAuthPluginProtocolsGrpc           KeyAuthPluginProtocols = "grpc"
	KeyAuthPluginProtocolsGrpcs          KeyAuthPluginProtocols = "grpcs"
	KeyAuthPluginProtocolsHTTP           KeyAuthPluginProtocols = "http"
	KeyAuthPluginProtocolsHTTPS          KeyAuthPluginProtocols = "https"
	KeyAuthPluginProtocolsTCP            KeyAuthPluginProtocols = "tcp"
	KeyAuthPluginProtocolsTLS            KeyAuthPluginProtocols = "tls"
	KeyAuthPluginProtocolsTLSPassthrough KeyAuthPluginProtocols = "tls_passthrough"
	KeyAuthPluginProtocolsUDP            KeyAuthPluginProtocols = "udp"
	KeyAuthPluginProtocolsWs             KeyAuthPluginProtocols = "ws"
	KeyAuthPluginProtocolsWss            KeyAuthPluginProtocols = "wss"
)

func (e KeyAuthPluginProtocols) ToPointer() *KeyAuthPluginProtocols {
	return &e
}

func (e *KeyAuthPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = KeyAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KeyAuthPluginProtocols: %v", v)
	}
}

// KeyAuthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type KeyAuthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KeyAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type KeyAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KeyAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type KeyAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KeyAuthPluginConfig struct {
	// Describes an array of parameter names where the plugin will look for a key. The key names may only contain [a-z], [A-Z], [0-9], [_] underscore, and [-] hyphen.
	KeyNames []string `json:"key_names,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service. If `true`, the plugin strips the credential from the request.
	HideCredentials *bool `default:"false" json:"hide_credentials"`
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request will fail with an authentication failure `4xx`.
	Anonymous *string `json:"anonymous,omitempty"`
	// If enabled (default), the plugin reads the request header and tries to find the key in it.
	KeyInHeader *bool `default:"true" json:"key_in_header"`
	// If enabled (default), the plugin reads the query parameter in the request and tries to find the key in it.
	KeyInQuery *bool `default:"true" json:"key_in_query"`
	// If enabled, the plugin reads the request body. Supported MIME types: `application/www-form-urlencoded`, `application/json`, and `multipart/form-data`.
	KeyInBody *bool `default:"false" json:"key_in_body"`
	// A boolean value that indicates whether the plugin should run (and try to authenticate) on `OPTIONS` preflight requests. If set to `false`, then `OPTIONS` requests are always allowed.
	RunOnPreflight *bool `default:"true" json:"run_on_preflight"`
}

func (k KeyAuthPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeyAuthPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KeyAuthPluginConfig) GetKeyNames() []string {
	if o == nil {
		return nil
	}
	return o.KeyNames
}

func (o *KeyAuthPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *KeyAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *KeyAuthPluginConfig) GetKeyInHeader() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInHeader
}

func (o *KeyAuthPluginConfig) GetKeyInQuery() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInQuery
}

func (o *KeyAuthPluginConfig) GetKeyInBody() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInBody
}

func (o *KeyAuthPluginConfig) GetRunOnPreflight() *bool {
	if o == nil {
		return nil
	}
	return o.RunOnPreflight
}

// KeyAuthPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type KeyAuthPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"key-auth" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []KeyAuthPluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *KeyAuthPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *KeyAuthPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *KeyAuthPluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64              `json:"created_at,omitempty"`
	ID        *string             `json:"id,omitempty"`
	Config    KeyAuthPluginConfig `json:"config"`
}

func (k KeyAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeyAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KeyAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *KeyAuthPlugin) GetName() string {
	return "key-auth"
}

func (o *KeyAuthPlugin) GetProtocols() []KeyAuthPluginProtocols {
	if o == nil {
		return []KeyAuthPluginProtocols{}
	}
	return o.Protocols
}

func (o *KeyAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *KeyAuthPlugin) GetConsumer() *KeyAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *KeyAuthPlugin) GetRoute() *KeyAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *KeyAuthPlugin) GetService() *KeyAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *KeyAuthPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *KeyAuthPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeyAuthPlugin) GetConfig() KeyAuthPluginConfig {
	if o == nil {
		return KeyAuthPluginConfig{}
	}
	return o.Config
}
