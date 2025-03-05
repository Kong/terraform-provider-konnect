// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type RouteTransformerAdvancedPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *RouteTransformerAdvancedPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RouteTransformerAdvancedPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *RouteTransformerAdvancedPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RouteTransformerAdvancedPluginOrdering struct {
	After  *RouteTransformerAdvancedPluginAfter  `json:"after,omitempty"`
	Before *RouteTransformerAdvancedPluginBefore `json:"before,omitempty"`
}

func (o *RouteTransformerAdvancedPluginOrdering) GetAfter() *RouteTransformerAdvancedPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *RouteTransformerAdvancedPluginOrdering) GetBefore() *RouteTransformerAdvancedPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type RouteTransformerAdvancedPluginConfig struct {
	EscapePath *bool   `json:"escape_path,omitempty"`
	Host       *string `json:"host,omitempty"`
	Path       *string `json:"path,omitempty"`
	Port       *string `json:"port,omitempty"`
}

func (o *RouteTransformerAdvancedPluginConfig) GetEscapePath() *bool {
	if o == nil {
		return nil
	}
	return o.EscapePath
}

func (o *RouteTransformerAdvancedPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *RouteTransformerAdvancedPluginConfig) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *RouteTransformerAdvancedPluginConfig) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

// RouteTransformerAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RouteTransformerAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RouteTransformerAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RouteTransformerAdvancedPluginProtocols string

const (
	RouteTransformerAdvancedPluginProtocolsGrpc  RouteTransformerAdvancedPluginProtocols = "grpc"
	RouteTransformerAdvancedPluginProtocolsGrpcs RouteTransformerAdvancedPluginProtocols = "grpcs"
	RouteTransformerAdvancedPluginProtocolsHTTP  RouteTransformerAdvancedPluginProtocols = "http"
	RouteTransformerAdvancedPluginProtocolsHTTPS RouteTransformerAdvancedPluginProtocols = "https"
)

func (e RouteTransformerAdvancedPluginProtocols) ToPointer() *RouteTransformerAdvancedPluginProtocols {
	return &e
}
func (e *RouteTransformerAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = RouteTransformerAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RouteTransformerAdvancedPluginProtocols: %v", v)
	}
}

// RouteTransformerAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type RouteTransformerAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RouteTransformerAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RouteTransformerAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RouteTransformerAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RouteTransformerAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RouteTransformerAdvancedPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RouteTransformerAdvancedPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                                   `json:"enabled,omitempty"`
	ID           *string                                 `json:"id,omitempty"`
	InstanceName *string                                 `json:"instance_name,omitempty"`
	name         string                                  `const:"route-transformer-advanced" json:"name"`
	Ordering     *RouteTransformerAdvancedPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                               `json:"updated_at,omitempty"`
	Config    RouteTransformerAdvancedPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *RouteTransformerAdvancedPluginConsumer `json:"consumer"`
	// A set of strings representing HTTP protocols.
	Protocols []RouteTransformerAdvancedPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *RouteTransformerAdvancedPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RouteTransformerAdvancedPluginService `json:"service"`
}

func (r RouteTransformerAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RouteTransformerAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RouteTransformerAdvancedPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RouteTransformerAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RouteTransformerAdvancedPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RouteTransformerAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RouteTransformerAdvancedPlugin) GetName() string {
	return "route-transformer-advanced"
}

func (o *RouteTransformerAdvancedPlugin) GetOrdering() *RouteTransformerAdvancedPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RouteTransformerAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RouteTransformerAdvancedPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RouteTransformerAdvancedPlugin) GetConfig() RouteTransformerAdvancedPluginConfig {
	if o == nil {
		return RouteTransformerAdvancedPluginConfig{}
	}
	return o.Config
}

func (o *RouteTransformerAdvancedPlugin) GetConsumer() *RouteTransformerAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RouteTransformerAdvancedPlugin) GetProtocols() []RouteTransformerAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RouteTransformerAdvancedPlugin) GetRoute() *RouteTransformerAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RouteTransformerAdvancedPlugin) GetService() *RouteTransformerAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// RouteTransformerAdvancedPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type RouteTransformerAdvancedPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool                                   `json:"enabled,omitempty"`
	ID           *string                                 `json:"id,omitempty"`
	InstanceName *string                                 `json:"instance_name,omitempty"`
	name         string                                  `const:"route-transformer-advanced" json:"name"`
	Ordering     *RouteTransformerAdvancedPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string                             `json:"tags,omitempty"`
	Config RouteTransformerAdvancedPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *RouteTransformerAdvancedPluginConsumer `json:"consumer"`
	// A set of strings representing HTTP protocols.
	Protocols []RouteTransformerAdvancedPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *RouteTransformerAdvancedPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RouteTransformerAdvancedPluginService `json:"service"`
}

func (r RouteTransformerAdvancedPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RouteTransformerAdvancedPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RouteTransformerAdvancedPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RouteTransformerAdvancedPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RouteTransformerAdvancedPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RouteTransformerAdvancedPluginInput) GetName() string {
	return "route-transformer-advanced"
}

func (o *RouteTransformerAdvancedPluginInput) GetOrdering() *RouteTransformerAdvancedPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RouteTransformerAdvancedPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RouteTransformerAdvancedPluginInput) GetConfig() RouteTransformerAdvancedPluginConfig {
	if o == nil {
		return RouteTransformerAdvancedPluginConfig{}
	}
	return o.Config
}

func (o *RouteTransformerAdvancedPluginInput) GetConsumer() *RouteTransformerAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RouteTransformerAdvancedPluginInput) GetProtocols() []RouteTransformerAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RouteTransformerAdvancedPluginInput) GetRoute() *RouteTransformerAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RouteTransformerAdvancedPluginInput) GetService() *RouteTransformerAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
