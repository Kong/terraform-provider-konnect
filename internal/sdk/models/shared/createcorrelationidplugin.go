// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateCorrelationIDPluginProtocols string

const (
	CreateCorrelationIDPluginProtocolsGrpc           CreateCorrelationIDPluginProtocols = "grpc"
	CreateCorrelationIDPluginProtocolsGrpcs          CreateCorrelationIDPluginProtocols = "grpcs"
	CreateCorrelationIDPluginProtocolsHTTP           CreateCorrelationIDPluginProtocols = "http"
	CreateCorrelationIDPluginProtocolsHTTPS          CreateCorrelationIDPluginProtocols = "https"
	CreateCorrelationIDPluginProtocolsTCP            CreateCorrelationIDPluginProtocols = "tcp"
	CreateCorrelationIDPluginProtocolsTLS            CreateCorrelationIDPluginProtocols = "tls"
	CreateCorrelationIDPluginProtocolsTLSPassthrough CreateCorrelationIDPluginProtocols = "tls_passthrough"
	CreateCorrelationIDPluginProtocolsUDP            CreateCorrelationIDPluginProtocols = "udp"
	CreateCorrelationIDPluginProtocolsWs             CreateCorrelationIDPluginProtocols = "ws"
	CreateCorrelationIDPluginProtocolsWss            CreateCorrelationIDPluginProtocols = "wss"
)

func (e CreateCorrelationIDPluginProtocols) ToPointer() *CreateCorrelationIDPluginProtocols {
	return &e
}
func (e *CreateCorrelationIDPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateCorrelationIDPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCorrelationIDPluginProtocols: %v", v)
	}
}

// CreateCorrelationIDPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateCorrelationIDPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCorrelationIDPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateCorrelationIDPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateCorrelationIDPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCorrelationIDPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateCorrelationIDPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateCorrelationIDPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCorrelationIDPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateCorrelationIDPluginGenerator - The generator to use for the correlation ID. Accepted values are `uuid`, `uuid#counter`, and `tracker`. See [Generators](#generators).
type CreateCorrelationIDPluginGenerator string

const (
	CreateCorrelationIDPluginGeneratorUUID              CreateCorrelationIDPluginGenerator = "uuid"
	CreateCorrelationIDPluginGeneratorUUIDNumberCounter CreateCorrelationIDPluginGenerator = "uuid#counter"
	CreateCorrelationIDPluginGeneratorTracker           CreateCorrelationIDPluginGenerator = "tracker"
)

func (e CreateCorrelationIDPluginGenerator) ToPointer() *CreateCorrelationIDPluginGenerator {
	return &e
}
func (e *CreateCorrelationIDPluginGenerator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "uuid":
		fallthrough
	case "uuid#counter":
		fallthrough
	case "tracker":
		*e = CreateCorrelationIDPluginGenerator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCorrelationIDPluginGenerator: %v", v)
	}
}

type CreateCorrelationIDPluginConfig struct {
	// Whether to echo the header back to downstream (the client).
	EchoDownstream *bool `default:"false" json:"echo_downstream"`
	// The generator to use for the correlation ID. Accepted values are `uuid`, `uuid#counter`, and `tracker`. See [Generators](#generators).
	Generator *CreateCorrelationIDPluginGenerator `default:"uuid#counter" json:"generator"`
	// The HTTP header name to use for the correlation ID.
	HeaderName *string `default:"Kong-Request-ID" json:"header_name"`
}

func (c CreateCorrelationIDPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCorrelationIDPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCorrelationIDPluginConfig) GetEchoDownstream() *bool {
	if o == nil {
		return nil
	}
	return o.EchoDownstream
}

func (o *CreateCorrelationIDPluginConfig) GetGenerator() *CreateCorrelationIDPluginGenerator {
	if o == nil {
		return nil
	}
	return o.Generator
}

func (o *CreateCorrelationIDPluginConfig) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

// CreateCorrelationIDPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateCorrelationIDPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"correlation-id" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateCorrelationIDPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateCorrelationIDPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateCorrelationIDPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateCorrelationIDPluginService `json:"service,omitempty"`
	Config  CreateCorrelationIDPluginConfig   `json:"config"`
}

func (c CreateCorrelationIDPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCorrelationIDPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCorrelationIDPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateCorrelationIDPlugin) GetName() string {
	return "correlation-id"
}

func (o *CreateCorrelationIDPlugin) GetProtocols() []CreateCorrelationIDPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateCorrelationIDPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateCorrelationIDPlugin) GetConsumer() *CreateCorrelationIDPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateCorrelationIDPlugin) GetRoute() *CreateCorrelationIDPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateCorrelationIDPlugin) GetService() *CreateCorrelationIDPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateCorrelationIDPlugin) GetConfig() CreateCorrelationIDPluginConfig {
	if o == nil {
		return CreateCorrelationIDPluginConfig{}
	}
	return o.Config
}
