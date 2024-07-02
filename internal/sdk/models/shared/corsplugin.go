// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type Methods string

const (
	MethodsGet     Methods = "GET"
	MethodsHead    Methods = "HEAD"
	MethodsPut     Methods = "PUT"
	MethodsPatch   Methods = "PATCH"
	MethodsPost    Methods = "POST"
	MethodsDelete  Methods = "DELETE"
	MethodsOptions Methods = "OPTIONS"
	MethodsTrace   Methods = "TRACE"
	MethodsConnect Methods = "CONNECT"
)

func (e Methods) ToPointer() *Methods {
	return &e
}
func (e *Methods) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "HEAD":
		fallthrough
	case "PUT":
		fallthrough
	case "PATCH":
		fallthrough
	case "POST":
		fallthrough
	case "DELETE":
		fallthrough
	case "OPTIONS":
		fallthrough
	case "TRACE":
		fallthrough
	case "CONNECT":
		*e = Methods(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Methods: %v", v)
	}
}

type CORSPluginConfig struct {
	// Flag to determine whether the `Access-Control-Allow-Credentials` header should be sent with `true` as the value.
	Credentials *bool `json:"credentials,omitempty"`
	// Value for the `Access-Control-Expose-Headers` header. If not specified, no custom headers are exposed.
	ExposedHeaders []string `json:"exposed_headers,omitempty"`
	// Value for the `Access-Control-Allow-Headers` header.
	Headers []string `json:"headers,omitempty"`
	// Indicates how long the results of the preflight request can be cached, in `seconds`.
	MaxAge *float64 `json:"max_age,omitempty"`
	// 'Value for the `Access-Control-Allow-Methods` header. Available options include `GET`, `HEAD`, `PUT`, `PATCH`, `POST`, `DELETE`, `OPTIONS`, `TRACE`, `CONNECT`. By default, all options are allowed.'
	Methods []Methods `json:"methods,omitempty"`
	// List of allowed domains for the `Access-Control-Allow-Origin` header. If you want to allow all origins, add `*` as a single value to this configuration field. The accepted values can either be flat strings or PCRE regexes.
	Origins []string `json:"origins,omitempty"`
	// A boolean value that instructs the plugin to proxy the `OPTIONS` preflight request to the Upstream service.
	PreflightContinue *bool `json:"preflight_continue,omitempty"`
	// Flag to determine whether the `Access-Control-Allow-Private-Network` header should be sent with `true` as the value.
	PrivateNetwork *bool `json:"private_network,omitempty"`
}

func (o *CORSPluginConfig) GetCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *CORSPluginConfig) GetExposedHeaders() []string {
	if o == nil {
		return nil
	}
	return o.ExposedHeaders
}

func (o *CORSPluginConfig) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CORSPluginConfig) GetMaxAge() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxAge
}

func (o *CORSPluginConfig) GetMethods() []Methods {
	if o == nil {
		return nil
	}
	return o.Methods
}

func (o *CORSPluginConfig) GetOrigins() []string {
	if o == nil {
		return nil
	}
	return o.Origins
}

func (o *CORSPluginConfig) GetPreflightContinue() *bool {
	if o == nil {
		return nil
	}
	return o.PreflightContinue
}

func (o *CORSPluginConfig) GetPrivateNetwork() *bool {
	if o == nil {
		return nil
	}
	return o.PrivateNetwork
}

type CORSPluginProtocols string

const (
	CORSPluginProtocolsGrpc           CORSPluginProtocols = "grpc"
	CORSPluginProtocolsGrpcs          CORSPluginProtocols = "grpcs"
	CORSPluginProtocolsHTTP           CORSPluginProtocols = "http"
	CORSPluginProtocolsHTTPS          CORSPluginProtocols = "https"
	CORSPluginProtocolsTCP            CORSPluginProtocols = "tcp"
	CORSPluginProtocolsTLS            CORSPluginProtocols = "tls"
	CORSPluginProtocolsTLSPassthrough CORSPluginProtocols = "tls_passthrough"
	CORSPluginProtocolsUDP            CORSPluginProtocols = "udp"
	CORSPluginProtocolsWs             CORSPluginProtocols = "ws"
	CORSPluginProtocolsWss            CORSPluginProtocols = "wss"
)

func (e CORSPluginProtocols) ToPointer() *CORSPluginProtocols {
	return &e
}
func (e *CORSPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CORSPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CORSPluginProtocols: %v", v)
	}
}

// CORSPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CORSPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CORSPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CORSPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CORSPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CORSPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CORSPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CORSPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CORSPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CORSPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CORSPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CORSPlugin struct {
	Config *CORSPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"cors" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CORSPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CORSPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CORSPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CORSPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CORSPluginService `json:"service,omitempty"`
}

func (c CORSPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CORSPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CORSPlugin) GetConfig() *CORSPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CORSPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CORSPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CORSPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CORSPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CORSPlugin) GetName() *string {
	return types.String("cors")
}

func (o *CORSPlugin) GetProtocols() []CORSPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CORSPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CORSPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CORSPlugin) GetConsumer() *CORSPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CORSPlugin) GetConsumerGroup() *CORSPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CORSPlugin) GetRoute() *CORSPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CORSPlugin) GetService() *CORSPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}