// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type IPRestrictionPluginConfig struct {
	// List of IPs or CIDR ranges to allow. One of `config.allow` or `config.deny` must be specified.
	Allow []string `json:"allow,omitempty"`
	// List of IPs or CIDR ranges to deny. One of `config.allow` or `config.deny` must be specified.
	Deny []string `json:"deny,omitempty"`
	// The message to send as a response body to rejected requests.
	Message *string `json:"message,omitempty"`
	// The HTTP status of the requests that will be rejected by the plugin.
	Status *float64 `json:"status,omitempty"`
}

func (o *IPRestrictionPluginConfig) GetAllow() []string {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *IPRestrictionPluginConfig) GetDeny() []string {
	if o == nil {
		return nil
	}
	return o.Deny
}

func (o *IPRestrictionPluginConfig) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *IPRestrictionPluginConfig) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type IPRestrictionPluginProtocols string

const (
	IPRestrictionPluginProtocolsGrpc           IPRestrictionPluginProtocols = "grpc"
	IPRestrictionPluginProtocolsGrpcs          IPRestrictionPluginProtocols = "grpcs"
	IPRestrictionPluginProtocolsHTTP           IPRestrictionPluginProtocols = "http"
	IPRestrictionPluginProtocolsHTTPS          IPRestrictionPluginProtocols = "https"
	IPRestrictionPluginProtocolsTCP            IPRestrictionPluginProtocols = "tcp"
	IPRestrictionPluginProtocolsTLS            IPRestrictionPluginProtocols = "tls"
	IPRestrictionPluginProtocolsTLSPassthrough IPRestrictionPluginProtocols = "tls_passthrough"
	IPRestrictionPluginProtocolsUDP            IPRestrictionPluginProtocols = "udp"
	IPRestrictionPluginProtocolsWs             IPRestrictionPluginProtocols = "ws"
	IPRestrictionPluginProtocolsWss            IPRestrictionPluginProtocols = "wss"
)

func (e IPRestrictionPluginProtocols) ToPointer() *IPRestrictionPluginProtocols {
	return &e
}
func (e *IPRestrictionPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = IPRestrictionPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IPRestrictionPluginProtocols: %v", v)
	}
}

// IPRestrictionPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type IPRestrictionPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *IPRestrictionPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type IPRestrictionPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *IPRestrictionPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// IPRestrictionPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type IPRestrictionPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *IPRestrictionPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// IPRestrictionPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type IPRestrictionPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *IPRestrictionPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type IPRestrictionPlugin struct {
	Config *IPRestrictionPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"ip-restriction" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []IPRestrictionPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *IPRestrictionPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *IPRestrictionPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *IPRestrictionPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *IPRestrictionPluginService `json:"service,omitempty"`
}

func (i IPRestrictionPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IPRestrictionPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IPRestrictionPlugin) GetConfig() *IPRestrictionPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *IPRestrictionPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IPRestrictionPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *IPRestrictionPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IPRestrictionPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *IPRestrictionPlugin) GetName() *string {
	return types.String("ip-restriction")
}

func (o *IPRestrictionPlugin) GetProtocols() []IPRestrictionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *IPRestrictionPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *IPRestrictionPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *IPRestrictionPlugin) GetConsumer() *IPRestrictionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *IPRestrictionPlugin) GetConsumerGroup() *IPRestrictionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *IPRestrictionPlugin) GetRoute() *IPRestrictionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *IPRestrictionPlugin) GetService() *IPRestrictionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
