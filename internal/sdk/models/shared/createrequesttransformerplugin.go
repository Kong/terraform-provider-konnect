// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateRequestTransformerPluginAdd struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *CreateRequestTransformerPluginAdd) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateRequestTransformerPluginAdd) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateRequestTransformerPluginAdd) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type CreateRequestTransformerPluginAppend struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *CreateRequestTransformerPluginAppend) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateRequestTransformerPluginAppend) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateRequestTransformerPluginAppend) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type CreateRequestTransformerPluginRemove struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *CreateRequestTransformerPluginRemove) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateRequestTransformerPluginRemove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateRequestTransformerPluginRemove) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type CreateRequestTransformerPluginRename struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *CreateRequestTransformerPluginRename) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateRequestTransformerPluginRename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateRequestTransformerPluginRename) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type CreateRequestTransformerPluginReplace struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
	URI         *string  `json:"uri,omitempty"`
}

func (o *CreateRequestTransformerPluginReplace) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateRequestTransformerPluginReplace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateRequestTransformerPluginReplace) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

func (o *CreateRequestTransformerPluginReplace) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}

type CreateRequestTransformerPluginConfig struct {
	Add    *CreateRequestTransformerPluginAdd    `json:"add,omitempty"`
	Append *CreateRequestTransformerPluginAppend `json:"append,omitempty"`
	// A string representing an HTTP method, such as GET, POST, PUT, or DELETE. The string must contain only uppercase letters.
	HTTPMethod *string                                `json:"http_method,omitempty"`
	Remove     *CreateRequestTransformerPluginRemove  `json:"remove,omitempty"`
	Rename     *CreateRequestTransformerPluginRename  `json:"rename,omitempty"`
	Replace    *CreateRequestTransformerPluginReplace `json:"replace,omitempty"`
}

func (o *CreateRequestTransformerPluginConfig) GetAdd() *CreateRequestTransformerPluginAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *CreateRequestTransformerPluginConfig) GetAppend() *CreateRequestTransformerPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *CreateRequestTransformerPluginConfig) GetHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *CreateRequestTransformerPluginConfig) GetRemove() *CreateRequestTransformerPluginRemove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *CreateRequestTransformerPluginConfig) GetRename() *CreateRequestTransformerPluginRename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *CreateRequestTransformerPluginConfig) GetReplace() *CreateRequestTransformerPluginReplace {
	if o == nil {
		return nil
	}
	return o.Replace
}

type CreateRequestTransformerPluginProtocols string

const (
	CreateRequestTransformerPluginProtocolsGrpc           CreateRequestTransformerPluginProtocols = "grpc"
	CreateRequestTransformerPluginProtocolsGrpcs          CreateRequestTransformerPluginProtocols = "grpcs"
	CreateRequestTransformerPluginProtocolsHTTP           CreateRequestTransformerPluginProtocols = "http"
	CreateRequestTransformerPluginProtocolsHTTPS          CreateRequestTransformerPluginProtocols = "https"
	CreateRequestTransformerPluginProtocolsTCP            CreateRequestTransformerPluginProtocols = "tcp"
	CreateRequestTransformerPluginProtocolsTLS            CreateRequestTransformerPluginProtocols = "tls"
	CreateRequestTransformerPluginProtocolsTLSPassthrough CreateRequestTransformerPluginProtocols = "tls_passthrough"
	CreateRequestTransformerPluginProtocolsUDP            CreateRequestTransformerPluginProtocols = "udp"
	CreateRequestTransformerPluginProtocolsWs             CreateRequestTransformerPluginProtocols = "ws"
	CreateRequestTransformerPluginProtocolsWss            CreateRequestTransformerPluginProtocols = "wss"
)

func (e CreateRequestTransformerPluginProtocols) ToPointer() *CreateRequestTransformerPluginProtocols {
	return &e
}
func (e *CreateRequestTransformerPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateRequestTransformerPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestTransformerPluginProtocols: %v", v)
	}
}

// CreateRequestTransformerPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateRequestTransformerPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTransformerPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestTransformerPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTransformerPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestTransformerPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateRequestTransformerPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTransformerPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestTransformerPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateRequestTransformerPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTransformerPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestTransformerPlugin struct {
	Config *CreateRequestTransformerPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"request-transformer" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateRequestTransformerPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateRequestTransformerPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateRequestTransformerPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateRequestTransformerPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateRequestTransformerPluginService `json:"service,omitempty"`
}

func (c CreateRequestTransformerPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRequestTransformerPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRequestTransformerPlugin) GetConfig() *CreateRequestTransformerPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateRequestTransformerPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateRequestTransformerPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateRequestTransformerPlugin) GetName() *string {
	return types.String("request-transformer")
}

func (o *CreateRequestTransformerPlugin) GetProtocols() []CreateRequestTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRequestTransformerPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateRequestTransformerPlugin) GetConsumer() *CreateRequestTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateRequestTransformerPlugin) GetConsumerGroup() *CreateRequestTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateRequestTransformerPlugin) GetRoute() *CreateRequestTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRequestTransformerPlugin) GetService() *CreateRequestTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}