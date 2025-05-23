// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type BasicAuthPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *BasicAuthPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type BasicAuthPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *BasicAuthPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type BasicAuthPluginOrdering struct {
	After  *BasicAuthPluginAfter  `json:"after,omitempty"`
	Before *BasicAuthPluginBefore `json:"before,omitempty"`
}

func (o *BasicAuthPluginOrdering) GetAfter() *BasicAuthPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *BasicAuthPluginOrdering) GetBefore() *BasicAuthPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type BasicAuthPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *BasicAuthPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BasicAuthPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *BasicAuthPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

type BasicAuthPluginConfig struct {
	// An optional string (Consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request will fail with an authentication failure `4xx`. Please note that this value must refer to the Consumer `id` or `username` attribute, and **not** its `custom_id`.
	Anonymous *string `json:"anonymous,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service. If `true`, the plugin will strip the credential from the request (i.e. the `Authorization` header) before proxying it.
	HideCredentials *bool `json:"hide_credentials,omitempty"`
	// When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
	Realm *string `json:"realm,omitempty"`
}

func (o *BasicAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *BasicAuthPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *BasicAuthPluginConfig) GetRealm() *string {
	if o == nil {
		return nil
	}
	return o.Realm
}

type BasicAuthPluginProtocols string

const (
	BasicAuthPluginProtocolsGrpc  BasicAuthPluginProtocols = "grpc"
	BasicAuthPluginProtocolsGrpcs BasicAuthPluginProtocols = "grpcs"
	BasicAuthPluginProtocolsHTTP  BasicAuthPluginProtocols = "http"
	BasicAuthPluginProtocolsHTTPS BasicAuthPluginProtocols = "https"
	BasicAuthPluginProtocolsWs    BasicAuthPluginProtocols = "ws"
	BasicAuthPluginProtocolsWss   BasicAuthPluginProtocols = "wss"
)

func (e BasicAuthPluginProtocols) ToPointer() *BasicAuthPluginProtocols {
	return &e
}
func (e *BasicAuthPluginProtocols) UnmarshalJSON(data []byte) error {
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
	case "ws":
		fallthrough
	case "wss":
		*e = BasicAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BasicAuthPluginProtocols: %v", v)
	}
}

// BasicAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type BasicAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *BasicAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// BasicAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type BasicAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *BasicAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// BasicAuthPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type BasicAuthPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                     `json:"enabled,omitempty"`
	ID           *string                   `json:"id,omitempty"`
	InstanceName *string                   `json:"instance_name,omitempty"`
	name         string                    `const:"basic-auth" json:"name"`
	Ordering     *BasicAuthPluginOrdering  `json:"ordering,omitempty"`
	Partials     []BasicAuthPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                 `json:"updated_at,omitempty"`
	Config    *BasicAuthPluginConfig `json:"config,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support tcp and tls.
	Protocols []BasicAuthPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *BasicAuthPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *BasicAuthPluginService `json:"service"`
}

func (b BasicAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BasicAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BasicAuthPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BasicAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *BasicAuthPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BasicAuthPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *BasicAuthPlugin) GetName() string {
	return "basic-auth"
}

func (o *BasicAuthPlugin) GetOrdering() *BasicAuthPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *BasicAuthPlugin) GetPartials() []BasicAuthPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *BasicAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BasicAuthPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *BasicAuthPlugin) GetConfig() *BasicAuthPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *BasicAuthPlugin) GetProtocols() []BasicAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *BasicAuthPlugin) GetRoute() *BasicAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *BasicAuthPlugin) GetService() *BasicAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
