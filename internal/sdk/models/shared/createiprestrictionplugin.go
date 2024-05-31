// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateIPRestrictionPluginProtocols string

const (
	CreateIPRestrictionPluginProtocolsGrpc           CreateIPRestrictionPluginProtocols = "grpc"
	CreateIPRestrictionPluginProtocolsGrpcs          CreateIPRestrictionPluginProtocols = "grpcs"
	CreateIPRestrictionPluginProtocolsHTTP           CreateIPRestrictionPluginProtocols = "http"
	CreateIPRestrictionPluginProtocolsHTTPS          CreateIPRestrictionPluginProtocols = "https"
	CreateIPRestrictionPluginProtocolsTCP            CreateIPRestrictionPluginProtocols = "tcp"
	CreateIPRestrictionPluginProtocolsTLS            CreateIPRestrictionPluginProtocols = "tls"
	CreateIPRestrictionPluginProtocolsTLSPassthrough CreateIPRestrictionPluginProtocols = "tls_passthrough"
	CreateIPRestrictionPluginProtocolsUDP            CreateIPRestrictionPluginProtocols = "udp"
	CreateIPRestrictionPluginProtocolsWs             CreateIPRestrictionPluginProtocols = "ws"
	CreateIPRestrictionPluginProtocolsWss            CreateIPRestrictionPluginProtocols = "wss"
)

func (e CreateIPRestrictionPluginProtocols) ToPointer() *CreateIPRestrictionPluginProtocols {
	return &e
}
func (e *CreateIPRestrictionPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateIPRestrictionPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateIPRestrictionPluginProtocols: %v", v)
	}
}

// CreateIPRestrictionPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateIPRestrictionPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateIPRestrictionPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateIPRestrictionPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateIPRestrictionPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateIPRestrictionPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateIPRestrictionPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateIPRestrictionPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateIPRestrictionPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateIPRestrictionPluginConfig struct {
	// List of IPs or CIDR ranges to allow. One of `config.allow` or `config.deny` must be specified.
	Allow []string `json:"allow,omitempty"`
	// List of IPs or CIDR ranges to deny. One of `config.allow` or `config.deny` must be specified.
	Deny []string `json:"deny,omitempty"`
	// The message to send as a response body to rejected requests.
	Message *string `json:"message,omitempty"`
	// The HTTP status of the requests that will be rejected by the plugin.
	Status *float64 `json:"status,omitempty"`
}

func (o *CreateIPRestrictionPluginConfig) GetAllow() []string {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *CreateIPRestrictionPluginConfig) GetDeny() []string {
	if o == nil {
		return nil
	}
	return o.Deny
}

func (o *CreateIPRestrictionPluginConfig) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateIPRestrictionPluginConfig) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// CreateIPRestrictionPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateIPRestrictionPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"ip-restriction" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateIPRestrictionPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateIPRestrictionPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateIPRestrictionPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateIPRestrictionPluginService `json:"service,omitempty"`
	Config  CreateIPRestrictionPluginConfig   `json:"config"`
}

func (c CreateIPRestrictionPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateIPRestrictionPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateIPRestrictionPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateIPRestrictionPlugin) GetName() string {
	return "ip-restriction"
}

func (o *CreateIPRestrictionPlugin) GetProtocols() []CreateIPRestrictionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateIPRestrictionPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateIPRestrictionPlugin) GetConsumer() *CreateIPRestrictionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateIPRestrictionPlugin) GetRoute() *CreateIPRestrictionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateIPRestrictionPlugin) GetService() *CreateIPRestrictionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateIPRestrictionPlugin) GetConfig() CreateIPRestrictionPluginConfig {
	if o == nil {
		return CreateIPRestrictionPluginConfig{}
	}
	return o.Config
}
