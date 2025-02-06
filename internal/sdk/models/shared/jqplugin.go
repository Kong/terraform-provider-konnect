// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type JqPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *JqPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JqPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *JqPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JqPluginOrdering struct {
	After  *JqPluginAfter  `json:"after,omitempty"`
	Before *JqPluginBefore `json:"before,omitempty"`
}

func (o *JqPluginOrdering) GetAfter() *JqPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *JqPluginOrdering) GetBefore() *JqPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type RequestJqProgramOptions struct {
	ASCIIOutput   *bool `json:"ascii_output,omitempty"`
	CompactOutput *bool `json:"compact_output,omitempty"`
	JoinOutput    *bool `json:"join_output,omitempty"`
	RawOutput     *bool `json:"raw_output,omitempty"`
	SortKeys      *bool `json:"sort_keys,omitempty"`
}

func (o *RequestJqProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *RequestJqProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *RequestJqProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *RequestJqProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *RequestJqProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type ResponseJqProgramOptions struct {
	ASCIIOutput   *bool `json:"ascii_output,omitempty"`
	CompactOutput *bool `json:"compact_output,omitempty"`
	JoinOutput    *bool `json:"join_output,omitempty"`
	RawOutput     *bool `json:"raw_output,omitempty"`
	SortKeys      *bool `json:"sort_keys,omitempty"`
}

func (o *ResponseJqProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *ResponseJqProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *ResponseJqProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *ResponseJqProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *ResponseJqProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type JqPluginConfig struct {
	RequestIfMediaType       []string                  `json:"request_if_media_type,omitempty"`
	RequestJqProgram         *string                   `json:"request_jq_program,omitempty"`
	RequestJqProgramOptions  *RequestJqProgramOptions  `json:"request_jq_program_options,omitempty"`
	ResponseIfMediaType      []string                  `json:"response_if_media_type,omitempty"`
	ResponseIfStatusCode     []int64                   `json:"response_if_status_code,omitempty"`
	ResponseJqProgram        *string                   `json:"response_jq_program,omitempty"`
	ResponseJqProgramOptions *ResponseJqProgramOptions `json:"response_jq_program_options,omitempty"`
}

func (o *JqPluginConfig) GetRequestIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.RequestIfMediaType
}

func (o *JqPluginConfig) GetRequestJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.RequestJqProgram
}

func (o *JqPluginConfig) GetRequestJqProgramOptions() *RequestJqProgramOptions {
	if o == nil {
		return nil
	}
	return o.RequestJqProgramOptions
}

func (o *JqPluginConfig) GetResponseIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.ResponseIfMediaType
}

func (o *JqPluginConfig) GetResponseIfStatusCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseIfStatusCode
}

func (o *JqPluginConfig) GetResponseJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgram
}

func (o *JqPluginConfig) GetResponseJqProgramOptions() *ResponseJqProgramOptions {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgramOptions
}

// JqPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type JqPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *JqPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JqPluginProtocols string

const (
	JqPluginProtocolsGrpc  JqPluginProtocols = "grpc"
	JqPluginProtocolsGrpcs JqPluginProtocols = "grpcs"
	JqPluginProtocolsHTTP  JqPluginProtocols = "http"
	JqPluginProtocolsHTTPS JqPluginProtocols = "https"
)

func (e JqPluginProtocols) ToPointer() *JqPluginProtocols {
	return &e
}
func (e *JqPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = JqPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JqPluginProtocols: %v", v)
	}
}

// JqPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type JqPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *JqPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JqPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type JqPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *JqPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JqPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type JqPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"jq" json:"name"`
	Ordering     *JqPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64         `json:"updated_at,omitempty"`
	Config    JqPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *JqPluginConsumer `json:"consumer,omitempty"`
	// A set of strings representing HTTP protocols.
	Protocols []JqPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *JqPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *JqPluginService `json:"service,omitempty"`
}

func (j JqPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JqPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JqPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *JqPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *JqPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JqPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *JqPlugin) GetName() string {
	return "jq"
}

func (o *JqPlugin) GetOrdering() *JqPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *JqPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *JqPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *JqPlugin) GetConfig() JqPluginConfig {
	if o == nil {
		return JqPluginConfig{}
	}
	return o.Config
}

func (o *JqPlugin) GetConsumer() *JqPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *JqPlugin) GetProtocols() []JqPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *JqPlugin) GetRoute() *JqPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *JqPlugin) GetService() *JqPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// JqPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type JqPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"jq" json:"name"`
	Ordering     *JqPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string       `json:"tags,omitempty"`
	Config JqPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *JqPluginConsumer `json:"consumer,omitempty"`
	// A set of strings representing HTTP protocols.
	Protocols []JqPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *JqPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *JqPluginService `json:"service,omitempty"`
}

func (j JqPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JqPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JqPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *JqPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JqPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *JqPluginInput) GetName() string {
	return "jq"
}

func (o *JqPluginInput) GetOrdering() *JqPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *JqPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *JqPluginInput) GetConfig() JqPluginConfig {
	if o == nil {
		return JqPluginConfig{}
	}
	return o.Config
}

func (o *JqPluginInput) GetConsumer() *JqPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *JqPluginInput) GetProtocols() []JqPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *JqPluginInput) GetRoute() *JqPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *JqPluginInput) GetService() *JqPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
