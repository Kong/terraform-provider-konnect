// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateRequestValidatorPluginIn - The location of the parameter.
type CreateRequestValidatorPluginIn string

const (
	CreateRequestValidatorPluginInQuery  CreateRequestValidatorPluginIn = "query"
	CreateRequestValidatorPluginInHeader CreateRequestValidatorPluginIn = "header"
	CreateRequestValidatorPluginInPath   CreateRequestValidatorPluginIn = "path"
)

func (e CreateRequestValidatorPluginIn) ToPointer() *CreateRequestValidatorPluginIn {
	return &e
}
func (e *CreateRequestValidatorPluginIn) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "query":
		fallthrough
	case "header":
		fallthrough
	case "path":
		*e = CreateRequestValidatorPluginIn(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestValidatorPluginIn: %v", v)
	}
}

// CreateRequestValidatorPluginStyle - Required when `schema` and `explode` are set. Describes how the parameter value will be deserialized depending on the type of the parameter value.
type CreateRequestValidatorPluginStyle string

const (
	CreateRequestValidatorPluginStyleLabel          CreateRequestValidatorPluginStyle = "label"
	CreateRequestValidatorPluginStyleForm           CreateRequestValidatorPluginStyle = "form"
	CreateRequestValidatorPluginStyleMatrix         CreateRequestValidatorPluginStyle = "matrix"
	CreateRequestValidatorPluginStyleSimple         CreateRequestValidatorPluginStyle = "simple"
	CreateRequestValidatorPluginStyleSpaceDelimited CreateRequestValidatorPluginStyle = "spaceDelimited"
	CreateRequestValidatorPluginStylePipeDelimited  CreateRequestValidatorPluginStyle = "pipeDelimited"
	CreateRequestValidatorPluginStyleDeepObject     CreateRequestValidatorPluginStyle = "deepObject"
)

func (e CreateRequestValidatorPluginStyle) ToPointer() *CreateRequestValidatorPluginStyle {
	return &e
}
func (e *CreateRequestValidatorPluginStyle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "label":
		fallthrough
	case "form":
		fallthrough
	case "matrix":
		fallthrough
	case "simple":
		fallthrough
	case "spaceDelimited":
		fallthrough
	case "pipeDelimited":
		fallthrough
	case "deepObject":
		*e = CreateRequestValidatorPluginStyle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestValidatorPluginStyle: %v", v)
	}
}

type CreateRequestValidatorPluginParameterSchema struct {
	// Required when `schema` and `style` are set. When `explode` is `true`, parameter values of type `array` or `object` generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters, this property has no effect.
	Explode *bool `json:"explode,omitempty"`
	// The location of the parameter.
	In CreateRequestValidatorPluginIn `json:"in"`
	// The name of the parameter. Parameter names are case-sensitive, and correspond to the parameter name used by the `in` property. If `in` is `path`, the `name` field MUST correspond to the named capture group from the configured `route`.
	Name string `json:"name"`
	// Determines whether this parameter is mandatory.
	Required bool `json:"required"`
	// Requred when `style` and `explode` are set. This is the schema defining the type used for the parameter. It is validated using `draft4` for JSON Schema draft 4 compliant validator. In addition to being a valid JSON Schema, the parameter schema MUST have a top-level `type` property to enable proper deserialization before validating.
	Schema *string `json:"schema,omitempty"`
	// Required when `schema` and `explode` are set. Describes how the parameter value will be deserialized depending on the type of the parameter value.
	Style *CreateRequestValidatorPluginStyle `json:"style,omitempty"`
}

func (o *CreateRequestValidatorPluginParameterSchema) GetExplode() *bool {
	if o == nil {
		return nil
	}
	return o.Explode
}

func (o *CreateRequestValidatorPluginParameterSchema) GetIn() CreateRequestValidatorPluginIn {
	if o == nil {
		return CreateRequestValidatorPluginIn("")
	}
	return o.In
}

func (o *CreateRequestValidatorPluginParameterSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateRequestValidatorPluginParameterSchema) GetRequired() bool {
	if o == nil {
		return false
	}
	return o.Required
}

func (o *CreateRequestValidatorPluginParameterSchema) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *CreateRequestValidatorPluginParameterSchema) GetStyle() *CreateRequestValidatorPluginStyle {
	if o == nil {
		return nil
	}
	return o.Style
}

// CreateRequestValidatorPluginVersion - Which validator to use. Supported values are `kong` (default) for using Kong's own schema validator, or `draft4` for using a JSON Schema Draft 4-compliant validator.
type CreateRequestValidatorPluginVersion string

const (
	CreateRequestValidatorPluginVersionKong   CreateRequestValidatorPluginVersion = "kong"
	CreateRequestValidatorPluginVersionDraft4 CreateRequestValidatorPluginVersion = "draft4"
)

func (e CreateRequestValidatorPluginVersion) ToPointer() *CreateRequestValidatorPluginVersion {
	return &e
}
func (e *CreateRequestValidatorPluginVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kong":
		fallthrough
	case "draft4":
		*e = CreateRequestValidatorPluginVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestValidatorPluginVersion: %v", v)
	}
}

type CreateRequestValidatorPluginConfig struct {
	// List of allowed content types. The value can be configured with the `charset` parameter. For example, `application/json; charset=UTF-8`.
	AllowedContentTypes []string `json:"allowed_content_types,omitempty"`
	// The request body schema specification. One of `body_schema` or `parameter_schema` must be specified.
	BodySchema *string `json:"body_schema,omitempty"`
	// Determines whether to enable parameters validation of request content-type.
	ContentTypeParameterValidation *bool `json:"content_type_parameter_validation,omitempty"`
	// Array of parameter validator specification. One of `body_schema` or `parameter_schema` must be specified.
	ParameterSchema []CreateRequestValidatorPluginParameterSchema `json:"parameter_schema,omitempty"`
	// If enabled, the plugin returns more verbose and detailed validation errors.
	VerboseResponse *bool `json:"verbose_response,omitempty"`
	// Which validator to use. Supported values are `kong` (default) for using Kong's own schema validator, or `draft4` for using a JSON Schema Draft 4-compliant validator.
	Version *CreateRequestValidatorPluginVersion `json:"version,omitempty"`
}

func (o *CreateRequestValidatorPluginConfig) GetAllowedContentTypes() []string {
	if o == nil {
		return nil
	}
	return o.AllowedContentTypes
}

func (o *CreateRequestValidatorPluginConfig) GetBodySchema() *string {
	if o == nil {
		return nil
	}
	return o.BodySchema
}

func (o *CreateRequestValidatorPluginConfig) GetContentTypeParameterValidation() *bool {
	if o == nil {
		return nil
	}
	return o.ContentTypeParameterValidation
}

func (o *CreateRequestValidatorPluginConfig) GetParameterSchema() []CreateRequestValidatorPluginParameterSchema {
	if o == nil {
		return nil
	}
	return o.ParameterSchema
}

func (o *CreateRequestValidatorPluginConfig) GetVerboseResponse() *bool {
	if o == nil {
		return nil
	}
	return o.VerboseResponse
}

func (o *CreateRequestValidatorPluginConfig) GetVersion() *CreateRequestValidatorPluginVersion {
	if o == nil {
		return nil
	}
	return o.Version
}

type CreateRequestValidatorPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateRequestValidatorPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateRequestValidatorPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateRequestValidatorPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateRequestValidatorPluginOrdering struct {
	After  *CreateRequestValidatorPluginAfter  `json:"after,omitempty"`
	Before *CreateRequestValidatorPluginBefore `json:"before,omitempty"`
}

func (o *CreateRequestValidatorPluginOrdering) GetAfter() *CreateRequestValidatorPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateRequestValidatorPluginOrdering) GetBefore() *CreateRequestValidatorPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateRequestValidatorPluginProtocols string

const (
	CreateRequestValidatorPluginProtocolsGrpc           CreateRequestValidatorPluginProtocols = "grpc"
	CreateRequestValidatorPluginProtocolsGrpcs          CreateRequestValidatorPluginProtocols = "grpcs"
	CreateRequestValidatorPluginProtocolsHTTP           CreateRequestValidatorPluginProtocols = "http"
	CreateRequestValidatorPluginProtocolsHTTPS          CreateRequestValidatorPluginProtocols = "https"
	CreateRequestValidatorPluginProtocolsTCP            CreateRequestValidatorPluginProtocols = "tcp"
	CreateRequestValidatorPluginProtocolsTLS            CreateRequestValidatorPluginProtocols = "tls"
	CreateRequestValidatorPluginProtocolsTLSPassthrough CreateRequestValidatorPluginProtocols = "tls_passthrough"
	CreateRequestValidatorPluginProtocolsUDP            CreateRequestValidatorPluginProtocols = "udp"
	CreateRequestValidatorPluginProtocolsWs             CreateRequestValidatorPluginProtocols = "ws"
	CreateRequestValidatorPluginProtocolsWss            CreateRequestValidatorPluginProtocols = "wss"
)

func (e CreateRequestValidatorPluginProtocols) ToPointer() *CreateRequestValidatorPluginProtocols {
	return &e
}
func (e *CreateRequestValidatorPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateRequestValidatorPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestValidatorPluginProtocols: %v", v)
	}
}

// CreateRequestValidatorPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateRequestValidatorPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestValidatorPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestValidatorPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestValidatorPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestValidatorPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateRequestValidatorPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestValidatorPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestValidatorPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateRequestValidatorPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestValidatorPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestValidatorPlugin struct {
	Config *CreateRequestValidatorPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                                 `json:"enabled,omitempty"`
	InstanceName *string                               `json:"instance_name,omitempty"`
	name         *string                               `const:"request-validator" json:"name,omitempty"`
	Ordering     *CreateRequestValidatorPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateRequestValidatorPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateRequestValidatorPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateRequestValidatorPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateRequestValidatorPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateRequestValidatorPluginService `json:"service,omitempty"`
}

func (c CreateRequestValidatorPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRequestValidatorPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRequestValidatorPlugin) GetConfig() *CreateRequestValidatorPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateRequestValidatorPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateRequestValidatorPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateRequestValidatorPlugin) GetName() *string {
	return types.String("request-validator")
}

func (o *CreateRequestValidatorPlugin) GetOrdering() *CreateRequestValidatorPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateRequestValidatorPlugin) GetProtocols() []CreateRequestValidatorPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRequestValidatorPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateRequestValidatorPlugin) GetConsumer() *CreateRequestValidatorPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateRequestValidatorPlugin) GetConsumerGroup() *CreateRequestValidatorPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateRequestValidatorPlugin) GetRoute() *CreateRequestValidatorPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRequestValidatorPlugin) GetService() *CreateRequestValidatorPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
