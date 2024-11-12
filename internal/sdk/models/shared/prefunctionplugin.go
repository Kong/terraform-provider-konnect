// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type PreFunctionPluginConfig struct {
	Access          []string `json:"access,omitempty"`
	BodyFilter      []string `json:"body_filter,omitempty"`
	Certificate     []string `json:"certificate,omitempty"`
	HeaderFilter    []string `json:"header_filter,omitempty"`
	Log             []string `json:"log,omitempty"`
	Rewrite         []string `json:"rewrite,omitempty"`
	WsClientFrame   []string `json:"ws_client_frame,omitempty"`
	WsClose         []string `json:"ws_close,omitempty"`
	WsHandshake     []string `json:"ws_handshake,omitempty"`
	WsUpstreamFrame []string `json:"ws_upstream_frame,omitempty"`
}

func (o *PreFunctionPluginConfig) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

func (o *PreFunctionPluginConfig) GetBodyFilter() []string {
	if o == nil {
		return nil
	}
	return o.BodyFilter
}

func (o *PreFunctionPluginConfig) GetCertificate() []string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *PreFunctionPluginConfig) GetHeaderFilter() []string {
	if o == nil {
		return nil
	}
	return o.HeaderFilter
}

func (o *PreFunctionPluginConfig) GetLog() []string {
	if o == nil {
		return nil
	}
	return o.Log
}

func (o *PreFunctionPluginConfig) GetRewrite() []string {
	if o == nil {
		return nil
	}
	return o.Rewrite
}

func (o *PreFunctionPluginConfig) GetWsClientFrame() []string {
	if o == nil {
		return nil
	}
	return o.WsClientFrame
}

func (o *PreFunctionPluginConfig) GetWsClose() []string {
	if o == nil {
		return nil
	}
	return o.WsClose
}

func (o *PreFunctionPluginConfig) GetWsHandshake() []string {
	if o == nil {
		return nil
	}
	return o.WsHandshake
}

func (o *PreFunctionPluginConfig) GetWsUpstreamFrame() []string {
	if o == nil {
		return nil
	}
	return o.WsUpstreamFrame
}

// PreFunctionPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type PreFunctionPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *PreFunctionPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PreFunctionPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *PreFunctionPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PreFunctionPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *PreFunctionPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type PreFunctionPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *PreFunctionPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type PreFunctionPluginOrdering struct {
	After  *PreFunctionPluginAfter  `json:"after,omitempty"`
	Before *PreFunctionPluginBefore `json:"before,omitempty"`
}

func (o *PreFunctionPluginOrdering) GetAfter() *PreFunctionPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *PreFunctionPluginOrdering) GetBefore() *PreFunctionPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type PreFunctionPluginProtocols string

const (
	PreFunctionPluginProtocolsGrpc           PreFunctionPluginProtocols = "grpc"
	PreFunctionPluginProtocolsGrpcs          PreFunctionPluginProtocols = "grpcs"
	PreFunctionPluginProtocolsHTTP           PreFunctionPluginProtocols = "http"
	PreFunctionPluginProtocolsHTTPS          PreFunctionPluginProtocols = "https"
	PreFunctionPluginProtocolsTCP            PreFunctionPluginProtocols = "tcp"
	PreFunctionPluginProtocolsTLS            PreFunctionPluginProtocols = "tls"
	PreFunctionPluginProtocolsTLSPassthrough PreFunctionPluginProtocols = "tls_passthrough"
	PreFunctionPluginProtocolsUDP            PreFunctionPluginProtocols = "udp"
	PreFunctionPluginProtocolsWs             PreFunctionPluginProtocols = "ws"
	PreFunctionPluginProtocolsWss            PreFunctionPluginProtocols = "wss"
)

func (e PreFunctionPluginProtocols) ToPointer() *PreFunctionPluginProtocols {
	return &e
}
func (e *PreFunctionPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = PreFunctionPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PreFunctionPluginProtocols: %v", v)
	}
}

// PreFunctionPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type PreFunctionPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *PreFunctionPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PreFunctionPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type PreFunctionPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *PreFunctionPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PreFunctionPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type PreFunctionPlugin struct {
	Config PreFunctionPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *PreFunctionPluginConsumer      `json:"consumer"`
	ConsumerGroup *PreFunctionPluginConsumerGroup `json:"consumer_group"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                      `json:"enabled,omitempty"`
	ID           *string                    `json:"id,omitempty"`
	InstanceName *string                    `json:"instance_name,omitempty"`
	name         string                     `const:"pre-function" json:"name"`
	Ordering     *PreFunctionPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []PreFunctionPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *PreFunctionPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *PreFunctionPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (p PreFunctionPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PreFunctionPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PreFunctionPlugin) GetConfig() PreFunctionPluginConfig {
	if o == nil {
		return PreFunctionPluginConfig{}
	}
	return o.Config
}

func (o *PreFunctionPlugin) GetConsumer() *PreFunctionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *PreFunctionPlugin) GetConsumerGroup() *PreFunctionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *PreFunctionPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PreFunctionPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *PreFunctionPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PreFunctionPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *PreFunctionPlugin) GetName() string {
	return "pre-function"
}

func (o *PreFunctionPlugin) GetOrdering() *PreFunctionPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *PreFunctionPlugin) GetProtocols() []PreFunctionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *PreFunctionPlugin) GetRoute() *PreFunctionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *PreFunctionPlugin) GetService() *PreFunctionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *PreFunctionPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PreFunctionPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// PreFunctionPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type PreFunctionPluginInput struct {
	Config PreFunctionPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *PreFunctionPluginConsumer      `json:"consumer"`
	ConsumerGroup *PreFunctionPluginConsumerGroup `json:"consumer_group"`
	// Whether the plugin is applied.
	Enabled      *bool                      `json:"enabled,omitempty"`
	ID           *string                    `json:"id,omitempty"`
	InstanceName *string                    `json:"instance_name,omitempty"`
	name         string                     `const:"pre-function" json:"name"`
	Ordering     *PreFunctionPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []PreFunctionPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *PreFunctionPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *PreFunctionPluginService `json:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (p PreFunctionPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PreFunctionPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PreFunctionPluginInput) GetConfig() PreFunctionPluginConfig {
	if o == nil {
		return PreFunctionPluginConfig{}
	}
	return o.Config
}

func (o *PreFunctionPluginInput) GetConsumer() *PreFunctionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *PreFunctionPluginInput) GetConsumerGroup() *PreFunctionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *PreFunctionPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *PreFunctionPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PreFunctionPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *PreFunctionPluginInput) GetName() string {
	return "pre-function"
}

func (o *PreFunctionPluginInput) GetOrdering() *PreFunctionPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *PreFunctionPluginInput) GetProtocols() []PreFunctionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *PreFunctionPluginInput) GetRoute() *PreFunctionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *PreFunctionPluginInput) GetService() *PreFunctionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *PreFunctionPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
