// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type Add struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Add) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Add) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Add) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Append struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Append) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Append) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Append) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Remove struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Remove) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Remove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Remove) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Rename struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Rename) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Rename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Rename) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Replace struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
	URI         *string  `json:"uri,omitempty"`
}

func (o *Replace) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Replace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Replace) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

func (o *Replace) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}

type RequestTransformerPluginConfig struct {
	Add    *Add    `json:"add,omitempty"`
	Append *Append `json:"append,omitempty"`
	// A string representing an HTTP method, such as GET, POST, PUT, or DELETE. The string must contain only uppercase letters.
	HTTPMethod *string  `json:"http_method,omitempty"`
	Remove     *Remove  `json:"remove,omitempty"`
	Rename     *Rename  `json:"rename,omitempty"`
	Replace    *Replace `json:"replace,omitempty"`
}

func (o *RequestTransformerPluginConfig) GetAdd() *Add {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *RequestTransformerPluginConfig) GetAppend() *Append {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *RequestTransformerPluginConfig) GetHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *RequestTransformerPluginConfig) GetRemove() *Remove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *RequestTransformerPluginConfig) GetRename() *Rename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *RequestTransformerPluginConfig) GetReplace() *Replace {
	if o == nil {
		return nil
	}
	return o.Replace
}

// RequestTransformerPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RequestTransformerPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerPluginConsumerGroup - If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
type RequestTransformerPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerPluginProtocols - A string representing a protocol, such as HTTP or HTTPS.
type RequestTransformerPluginProtocols string

const (
	RequestTransformerPluginProtocolsGrpc           RequestTransformerPluginProtocols = "grpc"
	RequestTransformerPluginProtocolsGrpcs          RequestTransformerPluginProtocols = "grpcs"
	RequestTransformerPluginProtocolsHTTP           RequestTransformerPluginProtocols = "http"
	RequestTransformerPluginProtocolsHTTPS          RequestTransformerPluginProtocols = "https"
	RequestTransformerPluginProtocolsTCP            RequestTransformerPluginProtocols = "tcp"
	RequestTransformerPluginProtocolsTLS            RequestTransformerPluginProtocols = "tls"
	RequestTransformerPluginProtocolsTLSPassthrough RequestTransformerPluginProtocols = "tls_passthrough"
	RequestTransformerPluginProtocolsUDP            RequestTransformerPluginProtocols = "udp"
	RequestTransformerPluginProtocolsWs             RequestTransformerPluginProtocols = "ws"
	RequestTransformerPluginProtocolsWss            RequestTransformerPluginProtocols = "wss"
)

func (e RequestTransformerPluginProtocols) ToPointer() *RequestTransformerPluginProtocols {
	return &e
}
func (e *RequestTransformerPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = RequestTransformerPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestTransformerPluginProtocols: %v", v)
	}
}

// RequestTransformerPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type RequestTransformerPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RequestTransformerPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RequestTransformerPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"request-transformer" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                         `json:"updated_at,omitempty"`
	Config    RequestTransformerPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *RequestTransformerPluginConsumer `json:"consumer"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *RequestTransformerPluginConsumerGroup `json:"consumer_group"`
	// A set of strings representing protocols.
	Protocols []RequestTransformerPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *RequestTransformerPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RequestTransformerPluginService `json:"service"`
}

func (r RequestTransformerPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestTransformerPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestTransformerPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestTransformerPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestTransformerPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestTransformerPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RequestTransformerPlugin) GetName() string {
	return "request-transformer"
}

func (o *RequestTransformerPlugin) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RequestTransformerPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RequestTransformerPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RequestTransformerPlugin) GetConfig() RequestTransformerPluginConfig {
	if o == nil {
		return RequestTransformerPluginConfig{}
	}
	return o.Config
}

func (o *RequestTransformerPlugin) GetConsumer() *RequestTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RequestTransformerPlugin) GetConsumerGroup() *RequestTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *RequestTransformerPlugin) GetProtocols() []RequestTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RequestTransformerPlugin) GetRoute() *RequestTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RequestTransformerPlugin) GetService() *RequestTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// RequestTransformerPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RequestTransformerPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"request-transformer" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string                       `json:"tags,omitempty"`
	Config RequestTransformerPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *RequestTransformerPluginConsumer `json:"consumer"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *RequestTransformerPluginConsumerGroup `json:"consumer_group"`
	// A set of strings representing protocols.
	Protocols []RequestTransformerPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *RequestTransformerPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RequestTransformerPluginService `json:"service"`
}

func (r RequestTransformerPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestTransformerPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestTransformerPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestTransformerPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestTransformerPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RequestTransformerPluginInput) GetName() string {
	return "request-transformer"
}

func (o *RequestTransformerPluginInput) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RequestTransformerPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RequestTransformerPluginInput) GetConfig() RequestTransformerPluginConfig {
	if o == nil {
		return RequestTransformerPluginConfig{}
	}
	return o.Config
}

func (o *RequestTransformerPluginInput) GetConsumer() *RequestTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RequestTransformerPluginInput) GetConsumerGroup() *RequestTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *RequestTransformerPluginInput) GetProtocols() []RequestTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RequestTransformerPluginInput) GetRoute() *RequestTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RequestTransformerPluginInput) GetService() *RequestTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
