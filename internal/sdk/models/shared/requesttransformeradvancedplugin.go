// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type JSONTypes string

const (
	JSONTypesBoolean JSONTypes = "boolean"
	JSONTypesNumber  JSONTypes = "number"
	JSONTypesString  JSONTypes = "string"
)

func (e JSONTypes) ToPointer() *JSONTypes {
	return &e
}
func (e *JSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = JSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JSONTypes: %v", v)
	}
}

type RequestTransformerAdvancedPluginAdd struct {
	Body        []string    `json:"body,omitempty"`
	Headers     []string    `json:"headers,omitempty"`
	JSONTypes   []JSONTypes `json:"json_types,omitempty"`
	Querystring []string    `json:"querystring,omitempty"`
}

func (o *RequestTransformerAdvancedPluginAdd) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerAdvancedPluginAdd) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerAdvancedPluginAdd) GetJSONTypes() []JSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

func (o *RequestTransformerAdvancedPluginAdd) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Allow struct {
	Body []string `json:"body,omitempty"`
}

func (o *Allow) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

type RequestTransformerAdvancedPluginJSONTypes string

const (
	RequestTransformerAdvancedPluginJSONTypesBoolean RequestTransformerAdvancedPluginJSONTypes = "boolean"
	RequestTransformerAdvancedPluginJSONTypesNumber  RequestTransformerAdvancedPluginJSONTypes = "number"
	RequestTransformerAdvancedPluginJSONTypesString  RequestTransformerAdvancedPluginJSONTypes = "string"
)

func (e RequestTransformerAdvancedPluginJSONTypes) ToPointer() *RequestTransformerAdvancedPluginJSONTypes {
	return &e
}
func (e *RequestTransformerAdvancedPluginJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = RequestTransformerAdvancedPluginJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestTransformerAdvancedPluginJSONTypes: %v", v)
	}
}

type RequestTransformerAdvancedPluginAppend struct {
	Body        []string                                    `json:"body,omitempty"`
	Headers     []string                                    `json:"headers,omitempty"`
	JSONTypes   []RequestTransformerAdvancedPluginJSONTypes `json:"json_types,omitempty"`
	Querystring []string                                    `json:"querystring,omitempty"`
}

func (o *RequestTransformerAdvancedPluginAppend) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerAdvancedPluginAppend) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerAdvancedPluginAppend) GetJSONTypes() []RequestTransformerAdvancedPluginJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

func (o *RequestTransformerAdvancedPluginAppend) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type RequestTransformerAdvancedPluginRemove struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *RequestTransformerAdvancedPluginRemove) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerAdvancedPluginRemove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerAdvancedPluginRemove) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type RequestTransformerAdvancedPluginRename struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *RequestTransformerAdvancedPluginRename) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerAdvancedPluginRename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerAdvancedPluginRename) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type RequestTransformerAdvancedPluginConfigJSONTypes string

const (
	RequestTransformerAdvancedPluginConfigJSONTypesBoolean RequestTransformerAdvancedPluginConfigJSONTypes = "boolean"
	RequestTransformerAdvancedPluginConfigJSONTypesNumber  RequestTransformerAdvancedPluginConfigJSONTypes = "number"
	RequestTransformerAdvancedPluginConfigJSONTypesString  RequestTransformerAdvancedPluginConfigJSONTypes = "string"
)

func (e RequestTransformerAdvancedPluginConfigJSONTypes) ToPointer() *RequestTransformerAdvancedPluginConfigJSONTypes {
	return &e
}
func (e *RequestTransformerAdvancedPluginConfigJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = RequestTransformerAdvancedPluginConfigJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestTransformerAdvancedPluginConfigJSONTypes: %v", v)
	}
}

type RequestTransformerAdvancedPluginReplace struct {
	Body        []string                                          `json:"body,omitempty"`
	Headers     []string                                          `json:"headers,omitempty"`
	JSONTypes   []RequestTransformerAdvancedPluginConfigJSONTypes `json:"json_types,omitempty"`
	Querystring []string                                          `json:"querystring,omitempty"`
	URI         *string                                           `json:"uri,omitempty"`
}

func (o *RequestTransformerAdvancedPluginReplace) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RequestTransformerAdvancedPluginReplace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *RequestTransformerAdvancedPluginReplace) GetJSONTypes() []RequestTransformerAdvancedPluginConfigJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

func (o *RequestTransformerAdvancedPluginReplace) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

func (o *RequestTransformerAdvancedPluginReplace) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}

type RequestTransformerAdvancedPluginConfig struct {
	Add    *RequestTransformerAdvancedPluginAdd    `json:"add,omitempty"`
	Allow  *Allow                                  `json:"allow,omitempty"`
	Append *RequestTransformerAdvancedPluginAppend `json:"append,omitempty"`
	// Specify whether dots (for example, `customers.info.phone`) should be treated as part of a property name or used to descend into nested JSON objects.  See [Arrays and nested objects](#arrays-and-nested-objects).
	DotsInKeys *bool `json:"dots_in_keys,omitempty"`
	// A string representing an HTTP method, such as GET, POST, PUT, or DELETE. The string must contain only uppercase letters.
	HTTPMethod *string                                  `json:"http_method,omitempty"`
	Remove     *RequestTransformerAdvancedPluginRemove  `json:"remove,omitempty"`
	Rename     *RequestTransformerAdvancedPluginRename  `json:"rename,omitempty"`
	Replace    *RequestTransformerAdvancedPluginReplace `json:"replace,omitempty"`
}

func (o *RequestTransformerAdvancedPluginConfig) GetAdd() *RequestTransformerAdvancedPluginAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *RequestTransformerAdvancedPluginConfig) GetAllow() *Allow {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *RequestTransformerAdvancedPluginConfig) GetAppend() *RequestTransformerAdvancedPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *RequestTransformerAdvancedPluginConfig) GetDotsInKeys() *bool {
	if o == nil {
		return nil
	}
	return o.DotsInKeys
}

func (o *RequestTransformerAdvancedPluginConfig) GetHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *RequestTransformerAdvancedPluginConfig) GetRemove() *RequestTransformerAdvancedPluginRemove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *RequestTransformerAdvancedPluginConfig) GetRename() *RequestTransformerAdvancedPluginRename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *RequestTransformerAdvancedPluginConfig) GetReplace() *RequestTransformerAdvancedPluginReplace {
	if o == nil {
		return nil
	}
	return o.Replace
}

type RequestTransformerAdvancedPluginProtocols string

const (
	RequestTransformerAdvancedPluginProtocolsGrpc           RequestTransformerAdvancedPluginProtocols = "grpc"
	RequestTransformerAdvancedPluginProtocolsGrpcs          RequestTransformerAdvancedPluginProtocols = "grpcs"
	RequestTransformerAdvancedPluginProtocolsHTTP           RequestTransformerAdvancedPluginProtocols = "http"
	RequestTransformerAdvancedPluginProtocolsHTTPS          RequestTransformerAdvancedPluginProtocols = "https"
	RequestTransformerAdvancedPluginProtocolsTCP            RequestTransformerAdvancedPluginProtocols = "tcp"
	RequestTransformerAdvancedPluginProtocolsTLS            RequestTransformerAdvancedPluginProtocols = "tls"
	RequestTransformerAdvancedPluginProtocolsTLSPassthrough RequestTransformerAdvancedPluginProtocols = "tls_passthrough"
	RequestTransformerAdvancedPluginProtocolsUDP            RequestTransformerAdvancedPluginProtocols = "udp"
	RequestTransformerAdvancedPluginProtocolsWs             RequestTransformerAdvancedPluginProtocols = "ws"
	RequestTransformerAdvancedPluginProtocolsWss            RequestTransformerAdvancedPluginProtocols = "wss"
)

func (e RequestTransformerAdvancedPluginProtocols) ToPointer() *RequestTransformerAdvancedPluginProtocols {
	return &e
}
func (e *RequestTransformerAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = RequestTransformerAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestTransformerAdvancedPluginProtocols: %v", v)
	}
}

// RequestTransformerAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RequestTransformerAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestTransformerAdvancedPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerAdvancedPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type RequestTransformerAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RequestTransformerAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RequestTransformerAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RequestTransformerAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestTransformerAdvancedPlugin struct {
	Config *RequestTransformerAdvancedPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"request-transformer-advanced" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []RequestTransformerAdvancedPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *RequestTransformerAdvancedPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *RequestTransformerAdvancedPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *RequestTransformerAdvancedPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RequestTransformerAdvancedPluginService `json:"service,omitempty"`
}

func (r RequestTransformerAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestTransformerAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestTransformerAdvancedPlugin) GetConfig() *RequestTransformerAdvancedPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *RequestTransformerAdvancedPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestTransformerAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestTransformerAdvancedPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestTransformerAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RequestTransformerAdvancedPlugin) GetName() *string {
	return types.String("request-transformer-advanced")
}

func (o *RequestTransformerAdvancedPlugin) GetProtocols() []RequestTransformerAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RequestTransformerAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RequestTransformerAdvancedPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RequestTransformerAdvancedPlugin) GetConsumer() *RequestTransformerAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RequestTransformerAdvancedPlugin) GetConsumerGroup() *RequestTransformerAdvancedPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *RequestTransformerAdvancedPlugin) GetRoute() *RequestTransformerAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RequestTransformerAdvancedPlugin) GetService() *RequestTransformerAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}