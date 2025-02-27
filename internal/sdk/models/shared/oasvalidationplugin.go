// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type OasValidationPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *OasValidationPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type OasValidationPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *OasValidationPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type OasValidationPluginOrdering struct {
	After  *OasValidationPluginAfter  `json:"after,omitempty"`
	Before *OasValidationPluginBefore `json:"before,omitempty"`
}

func (o *OasValidationPluginOrdering) GetAfter() *OasValidationPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *OasValidationPluginOrdering) GetBefore() *OasValidationPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type OasValidationPluginConfig struct {
	// List of header parameters in the request that will be ignored when performing HTTP header validation. These are additional headers added to an API request beyond those defined in the API specification.  For example, you might include the HTTP header `User-Agent`, which lets servers and network peers identify the application, operating system, vendor, and/or version of the requesting user agent.
	AllowedHeaderParameters *string `json:"allowed_header_parameters,omitempty"`
	// The API specification defined using either Swagger or the OpenAPI. This can be either a JSON or YAML based file. If using a YAML file, the spec needs to be URI-Encoded to preserve the YAML format.
	APISpec *string `json:"api_spec,omitempty"`
	// Indicates whether the api_spec is URI-Encoded.
	APISpecEncoded *bool `json:"api_spec_encoded,omitempty"`
	// The base path to be used for path match evaluation. This value is ignored if `include_base_path` is set to `false`.
	CustomBasePath *string `json:"custom_base_path,omitempty"`
	// If set to true, checks if HTTP header parameters in the request exist in the API specification.
	HeaderParameterCheck *bool `json:"header_parameter_check,omitempty"`
	// Indicates whether to include the base path when performing path match evaluation.
	IncludeBasePath *bool `json:"include_base_path,omitempty"`
	// If set to true, notifications via event hooks are enabled, but request based validation failures don't affect the request flow.
	NotifyOnlyRequestValidationFailure *bool `json:"notify_only_request_validation_failure,omitempty"`
	// If set to true, notifications via event hooks are enabled, but response validation failures don't affect the response flow.
	NotifyOnlyResponseBodyValidationFailure *bool `json:"notify_only_response_body_validation_failure,omitempty"`
	// If set to true, checks if query parameters in the request exist in the API specification.
	QueryParameterCheck *bool `json:"query_parameter_check,omitempty"`
	// If set to true, validates the request body content against the API specification.
	ValidateRequestBody *bool `json:"validate_request_body,omitempty"`
	// If set to true, validates HTTP header parameters against the API specification.
	ValidateRequestHeaderParams *bool `json:"validate_request_header_params,omitempty"`
	// If set to true, validates query parameters against the API specification.
	ValidateRequestQueryParams *bool `json:"validate_request_query_params,omitempty"`
	// If set to true, validates URI parameters in the request against the API specification.
	ValidateRequestURIParams *bool `json:"validate_request_uri_params,omitempty"`
	// If set to true, validates the response from the upstream services against the API specification. If validation fails, it results in an `HTTP 406 Not Acceptable` status code.
	ValidateResponseBody *bool `json:"validate_response_body,omitempty"`
	// If set to true, returns a detailed error message for invalid requests & responses. This is useful while testing.
	VerboseResponse *bool `json:"verbose_response,omitempty"`
}

func (o *OasValidationPluginConfig) GetAllowedHeaderParameters() *string {
	if o == nil {
		return nil
	}
	return o.AllowedHeaderParameters
}

func (o *OasValidationPluginConfig) GetAPISpec() *string {
	if o == nil {
		return nil
	}
	return o.APISpec
}

func (o *OasValidationPluginConfig) GetAPISpecEncoded() *bool {
	if o == nil {
		return nil
	}
	return o.APISpecEncoded
}

func (o *OasValidationPluginConfig) GetCustomBasePath() *string {
	if o == nil {
		return nil
	}
	return o.CustomBasePath
}

func (o *OasValidationPluginConfig) GetHeaderParameterCheck() *bool {
	if o == nil {
		return nil
	}
	return o.HeaderParameterCheck
}

func (o *OasValidationPluginConfig) GetIncludeBasePath() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeBasePath
}

func (o *OasValidationPluginConfig) GetNotifyOnlyRequestValidationFailure() *bool {
	if o == nil {
		return nil
	}
	return o.NotifyOnlyRequestValidationFailure
}

func (o *OasValidationPluginConfig) GetNotifyOnlyResponseBodyValidationFailure() *bool {
	if o == nil {
		return nil
	}
	return o.NotifyOnlyResponseBodyValidationFailure
}

func (o *OasValidationPluginConfig) GetQueryParameterCheck() *bool {
	if o == nil {
		return nil
	}
	return o.QueryParameterCheck
}

func (o *OasValidationPluginConfig) GetValidateRequestBody() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateRequestBody
}

func (o *OasValidationPluginConfig) GetValidateRequestHeaderParams() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateRequestHeaderParams
}

func (o *OasValidationPluginConfig) GetValidateRequestQueryParams() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateRequestQueryParams
}

func (o *OasValidationPluginConfig) GetValidateRequestURIParams() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateRequestURIParams
}

func (o *OasValidationPluginConfig) GetValidateResponseBody() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateResponseBody
}

func (o *OasValidationPluginConfig) GetVerboseResponse() *bool {
	if o == nil {
		return nil
	}
	return o.VerboseResponse
}

// OasValidationPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type OasValidationPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *OasValidationPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type OasValidationPluginProtocols string

const (
	OasValidationPluginProtocolsGrpc  OasValidationPluginProtocols = "grpc"
	OasValidationPluginProtocolsGrpcs OasValidationPluginProtocols = "grpcs"
	OasValidationPluginProtocolsHTTP  OasValidationPluginProtocols = "http"
	OasValidationPluginProtocolsHTTPS OasValidationPluginProtocols = "https"
)

func (e OasValidationPluginProtocols) ToPointer() *OasValidationPluginProtocols {
	return &e
}
func (e *OasValidationPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = OasValidationPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OasValidationPluginProtocols: %v", v)
	}
}

// OasValidationPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type OasValidationPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *OasValidationPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// OasValidationPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type OasValidationPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *OasValidationPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// OasValidationPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type OasValidationPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	ID           *string                      `json:"id,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         string                       `const:"oas-validation" json:"name"`
	Ordering     *OasValidationPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                    `json:"updated_at,omitempty"`
	Config    OasValidationPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *OasValidationPluginConsumer `json:"consumer"`
	// A set of strings representing HTTP protocols.
	Protocols []OasValidationPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *OasValidationPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *OasValidationPluginService `json:"service"`
}

func (o OasValidationPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OasValidationPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OasValidationPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *OasValidationPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *OasValidationPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OasValidationPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *OasValidationPlugin) GetName() string {
	return "oas-validation"
}

func (o *OasValidationPlugin) GetOrdering() *OasValidationPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *OasValidationPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *OasValidationPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *OasValidationPlugin) GetConfig() OasValidationPluginConfig {
	if o == nil {
		return OasValidationPluginConfig{}
	}
	return o.Config
}

func (o *OasValidationPlugin) GetConsumer() *OasValidationPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *OasValidationPlugin) GetProtocols() []OasValidationPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *OasValidationPlugin) GetRoute() *OasValidationPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *OasValidationPlugin) GetService() *OasValidationPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// OasValidationPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type OasValidationPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	ID           *string                      `json:"id,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         string                       `const:"oas-validation" json:"name"`
	Ordering     *OasValidationPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string                  `json:"tags,omitempty"`
	Config OasValidationPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *OasValidationPluginConsumer `json:"consumer"`
	// A set of strings representing HTTP protocols.
	Protocols []OasValidationPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *OasValidationPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *OasValidationPluginService `json:"service"`
}

func (o OasValidationPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OasValidationPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OasValidationPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *OasValidationPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OasValidationPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *OasValidationPluginInput) GetName() string {
	return "oas-validation"
}

func (o *OasValidationPluginInput) GetOrdering() *OasValidationPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *OasValidationPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *OasValidationPluginInput) GetConfig() OasValidationPluginConfig {
	if o == nil {
		return OasValidationPluginConfig{}
	}
	return o.Config
}

func (o *OasValidationPluginInput) GetConsumer() *OasValidationPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *OasValidationPluginInput) GetProtocols() []OasValidationPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *OasValidationPluginInput) GetRoute() *OasValidationPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *OasValidationPluginInput) GetService() *OasValidationPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
