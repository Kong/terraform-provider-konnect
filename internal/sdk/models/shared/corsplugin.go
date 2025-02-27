// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type CorsPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CorsPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CorsPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CorsPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CorsPluginOrdering struct {
	After  *CorsPluginAfter  `json:"after,omitempty"`
	Before *CorsPluginBefore `json:"before,omitempty"`
}

func (o *CorsPluginOrdering) GetAfter() *CorsPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CorsPluginOrdering) GetBefore() *CorsPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type Methods string

const (
	MethodsConnect Methods = "CONNECT"
	MethodsDelete  Methods = "DELETE"
	MethodsGet     Methods = "GET"
	MethodsHead    Methods = "HEAD"
	MethodsOptions Methods = "OPTIONS"
	MethodsPatch   Methods = "PATCH"
	MethodsPost    Methods = "POST"
	MethodsPut     Methods = "PUT"
	MethodsTrace   Methods = "TRACE"
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
	case "CONNECT":
		fallthrough
	case "DELETE":
		fallthrough
	case "GET":
		fallthrough
	case "HEAD":
		fallthrough
	case "OPTIONS":
		fallthrough
	case "PATCH":
		fallthrough
	case "POST":
		fallthrough
	case "PUT":
		fallthrough
	case "TRACE":
		*e = Methods(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Methods: %v", v)
	}
}

type CorsPluginConfig struct {
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

func (o *CorsPluginConfig) GetCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *CorsPluginConfig) GetExposedHeaders() []string {
	if o == nil {
		return nil
	}
	return o.ExposedHeaders
}

func (o *CorsPluginConfig) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CorsPluginConfig) GetMaxAge() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxAge
}

func (o *CorsPluginConfig) GetMethods() []Methods {
	if o == nil {
		return nil
	}
	return o.Methods
}

func (o *CorsPluginConfig) GetOrigins() []string {
	if o == nil {
		return nil
	}
	return o.Origins
}

func (o *CorsPluginConfig) GetPreflightContinue() *bool {
	if o == nil {
		return nil
	}
	return o.PreflightContinue
}

func (o *CorsPluginConfig) GetPrivateNetwork() *bool {
	if o == nil {
		return nil
	}
	return o.PrivateNetwork
}

type CorsPluginProtocols string

const (
	CorsPluginProtocolsGrpc  CorsPluginProtocols = "grpc"
	CorsPluginProtocolsGrpcs CorsPluginProtocols = "grpcs"
	CorsPluginProtocolsHTTP  CorsPluginProtocols = "http"
	CorsPluginProtocolsHTTPS CorsPluginProtocols = "https"
)

func (e CorsPluginProtocols) ToPointer() *CorsPluginProtocols {
	return &e
}
func (e *CorsPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CorsPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CorsPluginProtocols: %v", v)
	}
}

// CorsPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type CorsPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CorsPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CorsPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CorsPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CorsPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CorsPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CorsPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool               `json:"enabled,omitempty"`
	ID           *string             `json:"id,omitempty"`
	InstanceName *string             `json:"instance_name,omitempty"`
	name         string              `const:"cors" json:"name"`
	Ordering     *CorsPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64           `json:"updated_at,omitempty"`
	Config    CorsPluginConfig `json:"config"`
	// A set of strings representing HTTP protocols.
	Protocols []CorsPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *CorsPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CorsPluginService `json:"service"`
}

func (c CorsPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CorsPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CorsPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CorsPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CorsPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CorsPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CorsPlugin) GetName() string {
	return "cors"
}

func (o *CorsPlugin) GetOrdering() *CorsPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CorsPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CorsPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CorsPlugin) GetConfig() CorsPluginConfig {
	if o == nil {
		return CorsPluginConfig{}
	}
	return o.Config
}

func (o *CorsPlugin) GetProtocols() []CorsPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CorsPlugin) GetRoute() *CorsPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CorsPlugin) GetService() *CorsPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// CorsPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CorsPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool               `json:"enabled,omitempty"`
	ID           *string             `json:"id,omitempty"`
	InstanceName *string             `json:"instance_name,omitempty"`
	name         string              `const:"cors" json:"name"`
	Ordering     *CorsPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string         `json:"tags,omitempty"`
	Config CorsPluginConfig `json:"config"`
	// A set of strings representing HTTP protocols.
	Protocols []CorsPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *CorsPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CorsPluginService `json:"service"`
}

func (c CorsPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CorsPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CorsPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CorsPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CorsPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CorsPluginInput) GetName() string {
	return "cors"
}

func (o *CorsPluginInput) GetOrdering() *CorsPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CorsPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CorsPluginInput) GetConfig() CorsPluginConfig {
	if o == nil {
		return CorsPluginConfig{}
	}
	return o.Config
}

func (o *CorsPluginInput) GetProtocols() []CorsPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CorsPluginInput) GetRoute() *CorsPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CorsPluginInput) GetService() *CorsPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
