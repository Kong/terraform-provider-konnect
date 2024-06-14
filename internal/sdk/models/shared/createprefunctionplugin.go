// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreatePreFunctionPluginConfig struct {
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

func (o *CreatePreFunctionPluginConfig) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

func (o *CreatePreFunctionPluginConfig) GetBodyFilter() []string {
	if o == nil {
		return nil
	}
	return o.BodyFilter
}

func (o *CreatePreFunctionPluginConfig) GetCertificate() []string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *CreatePreFunctionPluginConfig) GetHeaderFilter() []string {
	if o == nil {
		return nil
	}
	return o.HeaderFilter
}

func (o *CreatePreFunctionPluginConfig) GetLog() []string {
	if o == nil {
		return nil
	}
	return o.Log
}

func (o *CreatePreFunctionPluginConfig) GetRewrite() []string {
	if o == nil {
		return nil
	}
	return o.Rewrite
}

func (o *CreatePreFunctionPluginConfig) GetWsClientFrame() []string {
	if o == nil {
		return nil
	}
	return o.WsClientFrame
}

func (o *CreatePreFunctionPluginConfig) GetWsClose() []string {
	if o == nil {
		return nil
	}
	return o.WsClose
}

func (o *CreatePreFunctionPluginConfig) GetWsHandshake() []string {
	if o == nil {
		return nil
	}
	return o.WsHandshake
}

func (o *CreatePreFunctionPluginConfig) GetWsUpstreamFrame() []string {
	if o == nil {
		return nil
	}
	return o.WsUpstreamFrame
}

type CreatePreFunctionPluginProtocols string

const (
	CreatePreFunctionPluginProtocolsGrpc           CreatePreFunctionPluginProtocols = "grpc"
	CreatePreFunctionPluginProtocolsGrpcs          CreatePreFunctionPluginProtocols = "grpcs"
	CreatePreFunctionPluginProtocolsHTTP           CreatePreFunctionPluginProtocols = "http"
	CreatePreFunctionPluginProtocolsHTTPS          CreatePreFunctionPluginProtocols = "https"
	CreatePreFunctionPluginProtocolsTCP            CreatePreFunctionPluginProtocols = "tcp"
	CreatePreFunctionPluginProtocolsTLS            CreatePreFunctionPluginProtocols = "tls"
	CreatePreFunctionPluginProtocolsTLSPassthrough CreatePreFunctionPluginProtocols = "tls_passthrough"
	CreatePreFunctionPluginProtocolsUDP            CreatePreFunctionPluginProtocols = "udp"
	CreatePreFunctionPluginProtocolsWs             CreatePreFunctionPluginProtocols = "ws"
	CreatePreFunctionPluginProtocolsWss            CreatePreFunctionPluginProtocols = "wss"
)

func (e CreatePreFunctionPluginProtocols) ToPointer() *CreatePreFunctionPluginProtocols {
	return &e
}
func (e *CreatePreFunctionPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreatePreFunctionPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePreFunctionPluginProtocols: %v", v)
	}
}

// CreatePreFunctionPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreatePreFunctionPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePreFunctionPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreatePreFunctionPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePreFunctionPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreatePreFunctionPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreatePreFunctionPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePreFunctionPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreatePreFunctionPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreatePreFunctionPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePreFunctionPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreatePreFunctionPlugin struct {
	Config *CreatePreFunctionPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"pre-function" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreatePreFunctionPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreatePreFunctionPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreatePreFunctionPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreatePreFunctionPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreatePreFunctionPluginService `json:"service,omitempty"`
}

func (c CreatePreFunctionPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePreFunctionPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePreFunctionPlugin) GetConfig() *CreatePreFunctionPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreatePreFunctionPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreatePreFunctionPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreatePreFunctionPlugin) GetName() *string {
	return types.String("pre-function")
}

func (o *CreatePreFunctionPlugin) GetProtocols() []CreatePreFunctionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreatePreFunctionPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreatePreFunctionPlugin) GetConsumer() *CreatePreFunctionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreatePreFunctionPlugin) GetConsumerGroup() *CreatePreFunctionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreatePreFunctionPlugin) GetRoute() *CreatePreFunctionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreatePreFunctionPlugin) GetService() *CreatePreFunctionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
