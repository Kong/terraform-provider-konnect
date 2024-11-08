// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type AzureFunctionsPluginConfig struct {
	// The apikey to access the Azure resources. If provided, it is injected as the `x-functions-key` header.
	Apikey *string `json:"apikey,omitempty"`
	// The Azure app name.
	Appname *string `json:"appname,omitempty"`
	// The `clientid` to access the Azure resources. If provided, it is injected as the `x-functions-clientid` header.
	Clientid *string `json:"clientid,omitempty"`
	// Name of the Azure function to invoke.
	Functionname *string `json:"functionname,omitempty"`
	// The domain where the function resides.
	Hostdomain *string `json:"hostdomain,omitempty"`
	// Use of HTTPS to connect with the Azure Functions server.
	HTTPS *bool `json:"https,omitempty"`
	// Set to `true` to authenticate the Azure Functions server.
	HTTPSVerify *bool `json:"https_verify,omitempty"`
	// Time in milliseconds during which an idle connection to the Azure Functions server lives before being closed.
	Keepalive *float64 `json:"keepalive,omitempty"`
	// Route prefix to use.
	Routeprefix *string `json:"routeprefix,omitempty"`
	// Timeout in milliseconds before closing a connection to the Azure Functions server.
	Timeout *float64 `json:"timeout,omitempty"`
}

func (o *AzureFunctionsPluginConfig) GetApikey() *string {
	if o == nil {
		return nil
	}
	return o.Apikey
}

func (o *AzureFunctionsPluginConfig) GetAppname() *string {
	if o == nil {
		return nil
	}
	return o.Appname
}

func (o *AzureFunctionsPluginConfig) GetClientid() *string {
	if o == nil {
		return nil
	}
	return o.Clientid
}

func (o *AzureFunctionsPluginConfig) GetFunctionname() *string {
	if o == nil {
		return nil
	}
	return o.Functionname
}

func (o *AzureFunctionsPluginConfig) GetHostdomain() *string {
	if o == nil {
		return nil
	}
	return o.Hostdomain
}

func (o *AzureFunctionsPluginConfig) GetHTTPS() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPS
}

func (o *AzureFunctionsPluginConfig) GetHTTPSVerify() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPSVerify
}

func (o *AzureFunctionsPluginConfig) GetKeepalive() *float64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *AzureFunctionsPluginConfig) GetRouteprefix() *string {
	if o == nil {
		return nil
	}
	return o.Routeprefix
}

func (o *AzureFunctionsPluginConfig) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

// AzureFunctionsPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AzureFunctionsPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AzureFunctionsPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AzureFunctionsPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AzureFunctionsPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AzureFunctionsPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *AzureFunctionsPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AzureFunctionsPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *AzureFunctionsPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AzureFunctionsPluginOrdering struct {
	After  *AzureFunctionsPluginAfter  `json:"after,omitempty"`
	Before *AzureFunctionsPluginBefore `json:"before,omitempty"`
}

func (o *AzureFunctionsPluginOrdering) GetAfter() *AzureFunctionsPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *AzureFunctionsPluginOrdering) GetBefore() *AzureFunctionsPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type AzureFunctionsPluginProtocols string

const (
	AzureFunctionsPluginProtocolsGrpc           AzureFunctionsPluginProtocols = "grpc"
	AzureFunctionsPluginProtocolsGrpcs          AzureFunctionsPluginProtocols = "grpcs"
	AzureFunctionsPluginProtocolsHTTP           AzureFunctionsPluginProtocols = "http"
	AzureFunctionsPluginProtocolsHTTPS          AzureFunctionsPluginProtocols = "https"
	AzureFunctionsPluginProtocolsTCP            AzureFunctionsPluginProtocols = "tcp"
	AzureFunctionsPluginProtocolsTLS            AzureFunctionsPluginProtocols = "tls"
	AzureFunctionsPluginProtocolsTLSPassthrough AzureFunctionsPluginProtocols = "tls_passthrough"
	AzureFunctionsPluginProtocolsUDP            AzureFunctionsPluginProtocols = "udp"
	AzureFunctionsPluginProtocolsWs             AzureFunctionsPluginProtocols = "ws"
	AzureFunctionsPluginProtocolsWss            AzureFunctionsPluginProtocols = "wss"
)

func (e AzureFunctionsPluginProtocols) ToPointer() *AzureFunctionsPluginProtocols {
	return &e
}
func (e *AzureFunctionsPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AzureFunctionsPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AzureFunctionsPluginProtocols: %v", v)
	}
}

// AzureFunctionsPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AzureFunctionsPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AzureFunctionsPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AzureFunctionsPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AzureFunctionsPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AzureFunctionsPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AzureFunctionsPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AzureFunctionsPlugin struct {
	Config AzureFunctionsPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *AzureFunctionsPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *AzureFunctionsPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                         `json:"enabled,omitempty"`
	ID           *string                       `json:"id,omitempty"`
	InstanceName *string                       `json:"instance_name,omitempty"`
	name         string                        `const:"azure-functions" json:"name"`
	Ordering     *AzureFunctionsPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AzureFunctionsPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AzureFunctionsPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AzureFunctionsPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (a AzureFunctionsPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AzureFunctionsPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AzureFunctionsPlugin) GetConfig() AzureFunctionsPluginConfig {
	if o == nil {
		return AzureFunctionsPluginConfig{}
	}
	return o.Config
}

func (o *AzureFunctionsPlugin) GetConsumer() *AzureFunctionsPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AzureFunctionsPlugin) GetConsumerGroup() *AzureFunctionsPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AzureFunctionsPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AzureFunctionsPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AzureFunctionsPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AzureFunctionsPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AzureFunctionsPlugin) GetName() string {
	return "azure-functions"
}

func (o *AzureFunctionsPlugin) GetOrdering() *AzureFunctionsPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AzureFunctionsPlugin) GetProtocols() []AzureFunctionsPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AzureFunctionsPlugin) GetRoute() *AzureFunctionsPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AzureFunctionsPlugin) GetService() *AzureFunctionsPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *AzureFunctionsPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AzureFunctionsPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// AzureFunctionsPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AzureFunctionsPluginInput struct {
	Config AzureFunctionsPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *AzureFunctionsPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *AzureFunctionsPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                         `json:"enabled,omitempty"`
	ID           *string                       `json:"id,omitempty"`
	InstanceName *string                       `json:"instance_name,omitempty"`
	name         string                        `const:"azure-functions" json:"name"`
	Ordering     *AzureFunctionsPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AzureFunctionsPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AzureFunctionsPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AzureFunctionsPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (a AzureFunctionsPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AzureFunctionsPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AzureFunctionsPluginInput) GetConfig() AzureFunctionsPluginConfig {
	if o == nil {
		return AzureFunctionsPluginConfig{}
	}
	return o.Config
}

func (o *AzureFunctionsPluginInput) GetConsumer() *AzureFunctionsPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AzureFunctionsPluginInput) GetConsumerGroup() *AzureFunctionsPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AzureFunctionsPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AzureFunctionsPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AzureFunctionsPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AzureFunctionsPluginInput) GetName() string {
	return "azure-functions"
}

func (o *AzureFunctionsPluginInput) GetOrdering() *AzureFunctionsPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AzureFunctionsPluginInput) GetProtocols() []AzureFunctionsPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AzureFunctionsPluginInput) GetRoute() *AzureFunctionsPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AzureFunctionsPluginInput) GetService() *AzureFunctionsPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *AzureFunctionsPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
