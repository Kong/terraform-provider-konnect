// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type PrometheusPluginConfig struct {
	// A boolean value that determines if bandwidth metrics should be collected. If enabled, `bandwidth_bytes` and `stream_sessions_total` metrics will be exported.
	BandwidthMetrics *bool `json:"bandwidth_metrics,omitempty"`
	// A boolean value that determines if latency metrics should be collected. If enabled, `kong_latency_ms`, `upstream_latency_ms` and `request_latency_ms` metrics will be exported.
	LatencyMetrics *bool `json:"latency_metrics,omitempty"`
	// A boolean value that determines if per-consumer metrics should be collected. If enabled, the `kong_http_requests_total` and `kong_bandwidth_bytes` metrics fill in the consumer label when available.
	PerConsumer *bool `json:"per_consumer,omitempty"`
	// A boolean value that determines if status code metrics should be collected. If enabled, `http_requests_total`, `stream_sessions_total` metrics will be exported.
	StatusCodeMetrics *bool `json:"status_code_metrics,omitempty"`
	// A boolean value that determines if upstream metrics should be collected. If enabled, `upstream_target_health` metric will be exported.
	UpstreamHealthMetrics *bool `json:"upstream_health_metrics,omitempty"`
}

func (o *PrometheusPluginConfig) GetBandwidthMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.BandwidthMetrics
}

func (o *PrometheusPluginConfig) GetLatencyMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.LatencyMetrics
}

func (o *PrometheusPluginConfig) GetPerConsumer() *bool {
	if o == nil {
		return nil
	}
	return o.PerConsumer
}

func (o *PrometheusPluginConfig) GetStatusCodeMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.StatusCodeMetrics
}

func (o *PrometheusPluginConfig) GetUpstreamHealthMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.UpstreamHealthMetrics
}

type PrometheusPluginProtocols string

const (
	PrometheusPluginProtocolsGrpc           PrometheusPluginProtocols = "grpc"
	PrometheusPluginProtocolsGrpcs          PrometheusPluginProtocols = "grpcs"
	PrometheusPluginProtocolsHTTP           PrometheusPluginProtocols = "http"
	PrometheusPluginProtocolsHTTPS          PrometheusPluginProtocols = "https"
	PrometheusPluginProtocolsTCP            PrometheusPluginProtocols = "tcp"
	PrometheusPluginProtocolsTLS            PrometheusPluginProtocols = "tls"
	PrometheusPluginProtocolsTLSPassthrough PrometheusPluginProtocols = "tls_passthrough"
	PrometheusPluginProtocolsUDP            PrometheusPluginProtocols = "udp"
	PrometheusPluginProtocolsWs             PrometheusPluginProtocols = "ws"
	PrometheusPluginProtocolsWss            PrometheusPluginProtocols = "wss"
)

func (e PrometheusPluginProtocols) ToPointer() *PrometheusPluginProtocols {
	return &e
}
func (e *PrometheusPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = PrometheusPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PrometheusPluginProtocols: %v", v)
	}
}

// PrometheusPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type PrometheusPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *PrometheusPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PrometheusPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *PrometheusPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PrometheusPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type PrometheusPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *PrometheusPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PrometheusPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type PrometheusPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *PrometheusPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PrometheusPlugin struct {
	Config *PrometheusPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"prometheus" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []PrometheusPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *PrometheusPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *PrometheusPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *PrometheusPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *PrometheusPluginService `json:"service,omitempty"`
}

func (p PrometheusPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PrometheusPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PrometheusPlugin) GetConfig() *PrometheusPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *PrometheusPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PrometheusPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *PrometheusPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PrometheusPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *PrometheusPlugin) GetName() *string {
	return types.String("prometheus")
}

func (o *PrometheusPlugin) GetProtocols() []PrometheusPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *PrometheusPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PrometheusPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *PrometheusPlugin) GetConsumer() *PrometheusPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *PrometheusPlugin) GetConsumerGroup() *PrometheusPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *PrometheusPlugin) GetRoute() *PrometheusPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *PrometheusPlugin) GetService() *PrometheusPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
