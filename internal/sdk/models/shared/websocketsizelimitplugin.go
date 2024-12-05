// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type WebsocketSizeLimitPluginConfig struct {
	ClientMaxPayload   *int64 `json:"client_max_payload,omitempty"`
	UpstreamMaxPayload *int64 `json:"upstream_max_payload,omitempty"`
}

func (o *WebsocketSizeLimitPluginConfig) GetClientMaxPayload() *int64 {
	if o == nil {
		return nil
	}
	return o.ClientMaxPayload
}

func (o *WebsocketSizeLimitPluginConfig) GetUpstreamMaxPayload() *int64 {
	if o == nil {
		return nil
	}
	return o.UpstreamMaxPayload
}

// WebsocketSizeLimitPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type WebsocketSizeLimitPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketSizeLimitPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type WebsocketSizeLimitPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketSizeLimitPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type WebsocketSizeLimitPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *WebsocketSizeLimitPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type WebsocketSizeLimitPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *WebsocketSizeLimitPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type WebsocketSizeLimitPluginOrdering struct {
	After  *WebsocketSizeLimitPluginAfter  `json:"after,omitempty"`
	Before *WebsocketSizeLimitPluginBefore `json:"before,omitempty"`
}

func (o *WebsocketSizeLimitPluginOrdering) GetAfter() *WebsocketSizeLimitPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *WebsocketSizeLimitPluginOrdering) GetBefore() *WebsocketSizeLimitPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type WebsocketSizeLimitPluginProtocols string

const (
	WebsocketSizeLimitPluginProtocolsGrpc           WebsocketSizeLimitPluginProtocols = "grpc"
	WebsocketSizeLimitPluginProtocolsGrpcs          WebsocketSizeLimitPluginProtocols = "grpcs"
	WebsocketSizeLimitPluginProtocolsHTTP           WebsocketSizeLimitPluginProtocols = "http"
	WebsocketSizeLimitPluginProtocolsHTTPS          WebsocketSizeLimitPluginProtocols = "https"
	WebsocketSizeLimitPluginProtocolsTCP            WebsocketSizeLimitPluginProtocols = "tcp"
	WebsocketSizeLimitPluginProtocolsTLS            WebsocketSizeLimitPluginProtocols = "tls"
	WebsocketSizeLimitPluginProtocolsTLSPassthrough WebsocketSizeLimitPluginProtocols = "tls_passthrough"
	WebsocketSizeLimitPluginProtocolsUDP            WebsocketSizeLimitPluginProtocols = "udp"
	WebsocketSizeLimitPluginProtocolsWs             WebsocketSizeLimitPluginProtocols = "ws"
	WebsocketSizeLimitPluginProtocolsWss            WebsocketSizeLimitPluginProtocols = "wss"
)

func (e WebsocketSizeLimitPluginProtocols) ToPointer() *WebsocketSizeLimitPluginProtocols {
	return &e
}
func (e *WebsocketSizeLimitPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = WebsocketSizeLimitPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebsocketSizeLimitPluginProtocols: %v", v)
	}
}

// WebsocketSizeLimitPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type WebsocketSizeLimitPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketSizeLimitPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// WebsocketSizeLimitPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type WebsocketSizeLimitPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketSizeLimitPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// WebsocketSizeLimitPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type WebsocketSizeLimitPlugin struct {
	Config WebsocketSizeLimitPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *WebsocketSizeLimitPluginConsumer      `json:"consumer"`
	ConsumerGroup *WebsocketSizeLimitPluginConsumerGroup `json:"consumer_group"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                             `json:"enabled,omitempty"`
	ID           *string                           `json:"id,omitempty"`
	InstanceName *string                           `json:"instance_name,omitempty"`
	name         string                            `const:"websocket-size-limit" json:"name"`
	Ordering     *WebsocketSizeLimitPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []WebsocketSizeLimitPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *WebsocketSizeLimitPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *WebsocketSizeLimitPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (w WebsocketSizeLimitPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebsocketSizeLimitPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebsocketSizeLimitPlugin) GetConfig() WebsocketSizeLimitPluginConfig {
	if o == nil {
		return WebsocketSizeLimitPluginConfig{}
	}
	return o.Config
}

func (o *WebsocketSizeLimitPlugin) GetConsumer() *WebsocketSizeLimitPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *WebsocketSizeLimitPlugin) GetConsumerGroup() *WebsocketSizeLimitPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *WebsocketSizeLimitPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WebsocketSizeLimitPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *WebsocketSizeLimitPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WebsocketSizeLimitPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *WebsocketSizeLimitPlugin) GetName() string {
	return "websocket-size-limit"
}

func (o *WebsocketSizeLimitPlugin) GetOrdering() *WebsocketSizeLimitPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *WebsocketSizeLimitPlugin) GetProtocols() []WebsocketSizeLimitPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *WebsocketSizeLimitPlugin) GetRoute() *WebsocketSizeLimitPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *WebsocketSizeLimitPlugin) GetService() *WebsocketSizeLimitPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *WebsocketSizeLimitPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *WebsocketSizeLimitPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// WebsocketSizeLimitPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type WebsocketSizeLimitPluginInput struct {
	Config WebsocketSizeLimitPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *WebsocketSizeLimitPluginConsumer      `json:"consumer"`
	ConsumerGroup *WebsocketSizeLimitPluginConsumerGroup `json:"consumer_group"`
	// Whether the plugin is applied.
	Enabled      *bool                             `json:"enabled,omitempty"`
	ID           *string                           `json:"id,omitempty"`
	InstanceName *string                           `json:"instance_name,omitempty"`
	name         string                            `const:"websocket-size-limit" json:"name"`
	Ordering     *WebsocketSizeLimitPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []WebsocketSizeLimitPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *WebsocketSizeLimitPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *WebsocketSizeLimitPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (w WebsocketSizeLimitPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebsocketSizeLimitPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebsocketSizeLimitPluginInput) GetConfig() WebsocketSizeLimitPluginConfig {
	if o == nil {
		return WebsocketSizeLimitPluginConfig{}
	}
	return o.Config
}

func (o *WebsocketSizeLimitPluginInput) GetConsumer() *WebsocketSizeLimitPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *WebsocketSizeLimitPluginInput) GetConsumerGroup() *WebsocketSizeLimitPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *WebsocketSizeLimitPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *WebsocketSizeLimitPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WebsocketSizeLimitPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *WebsocketSizeLimitPluginInput) GetName() string {
	return "websocket-size-limit"
}

func (o *WebsocketSizeLimitPluginInput) GetOrdering() *WebsocketSizeLimitPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *WebsocketSizeLimitPluginInput) GetProtocols() []WebsocketSizeLimitPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *WebsocketSizeLimitPluginInput) GetRoute() *WebsocketSizeLimitPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *WebsocketSizeLimitPluginInput) GetService() *WebsocketSizeLimitPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *WebsocketSizeLimitPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}