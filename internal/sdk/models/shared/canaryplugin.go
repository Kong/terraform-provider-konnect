// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

// Hash algorithm to be used for canary release.
//
// * `consumer`: The hash will be based on the consumer.
// * `ip`: The hash will be based on the client IP address.
// * `none`: No hash will be applied.
// * `allow`: Allows the specified groups to access the canary release.
// * `deny`: Denies the specified groups from accessing the canary release.
// * `header`: The hash will be based on the specified header value.
type Hash string

const (
	HashConsumer Hash = "consumer"
	HashIP       Hash = "ip"
	HashNone     Hash = "none"
	HashAllow    Hash = "allow"
	HashDeny     Hash = "deny"
	HashHeader   Hash = "header"
)

func (e Hash) ToPointer() *Hash {
	return &e
}
func (e *Hash) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer":
		fallthrough
	case "ip":
		fallthrough
	case "none":
		fallthrough
	case "allow":
		fallthrough
	case "deny":
		fallthrough
	case "header":
		*e = Hash(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Hash: %v", v)
	}
}

type CanaryPluginConfig struct {
	// A string representing an HTTP header name.
	CanaryByHeaderName *string `json:"canary_by_header_name,omitempty"`
	// The duration of the canary release in seconds.
	Duration *float64 `json:"duration,omitempty"`
	// The groups allowed to access the canary release.
	Groups []string `json:"groups,omitempty"`
	// Hash algorithm to be used for canary release.
	//
	// * `consumer`: The hash will be based on the consumer.
	// * `ip`: The hash will be based on the client IP address.
	// * `none`: No hash will be applied.
	// * `allow`: Allows the specified groups to access the canary release.
	// * `deny`: Denies the specified groups from accessing the canary release.
	// * `header`: The hash will be based on the specified header value.
	Hash *Hash `json:"hash,omitempty"`
	// A string representing an HTTP header name.
	HashHeader *string `json:"hash_header,omitempty"`
	// The percentage of traffic to be routed to the canary release.
	Percentage *float64 `json:"percentage,omitempty"`
	// Future time in seconds since epoch, when the canary release will start. Ignored when `percentage` is set, or when using `allow` or `deny` in `hash`.
	Start *float64 `json:"start,omitempty"`
	// The number of steps for the canary release.
	Steps *float64 `json:"steps,omitempty"`
	// Specifies whether to fallback to the upstream server if the canary release fails.
	UpstreamFallback *bool `json:"upstream_fallback,omitempty"`
	// A string representing a host name, such as example.com.
	UpstreamHost *string `json:"upstream_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	UpstreamPort *int64 `json:"upstream_port,omitempty"`
	// The URI of the upstream server to be used for the canary release.
	UpstreamURI *string `json:"upstream_uri,omitempty"`
}

func (o *CanaryPluginConfig) GetCanaryByHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.CanaryByHeaderName
}

func (o *CanaryPluginConfig) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *CanaryPluginConfig) GetGroups() []string {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *CanaryPluginConfig) GetHash() *Hash {
	if o == nil {
		return nil
	}
	return o.Hash
}

func (o *CanaryPluginConfig) GetHashHeader() *string {
	if o == nil {
		return nil
	}
	return o.HashHeader
}

func (o *CanaryPluginConfig) GetPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.Percentage
}

func (o *CanaryPluginConfig) GetStart() *float64 {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *CanaryPluginConfig) GetSteps() *float64 {
	if o == nil {
		return nil
	}
	return o.Steps
}

func (o *CanaryPluginConfig) GetUpstreamFallback() *bool {
	if o == nil {
		return nil
	}
	return o.UpstreamFallback
}

func (o *CanaryPluginConfig) GetUpstreamHost() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamHost
}

func (o *CanaryPluginConfig) GetUpstreamPort() *int64 {
	if o == nil {
		return nil
	}
	return o.UpstreamPort
}

func (o *CanaryPluginConfig) GetUpstreamURI() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamURI
}

// CanaryPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CanaryPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CanaryPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CanaryPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CanaryPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CanaryPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CanaryPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CanaryPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CanaryPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CanaryPluginOrdering struct {
	After  *CanaryPluginAfter  `json:"after,omitempty"`
	Before *CanaryPluginBefore `json:"before,omitempty"`
}

func (o *CanaryPluginOrdering) GetAfter() *CanaryPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CanaryPluginOrdering) GetBefore() *CanaryPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CanaryPluginProtocols string

const (
	CanaryPluginProtocolsGrpc           CanaryPluginProtocols = "grpc"
	CanaryPluginProtocolsGrpcs          CanaryPluginProtocols = "grpcs"
	CanaryPluginProtocolsHTTP           CanaryPluginProtocols = "http"
	CanaryPluginProtocolsHTTPS          CanaryPluginProtocols = "https"
	CanaryPluginProtocolsTCP            CanaryPluginProtocols = "tcp"
	CanaryPluginProtocolsTLS            CanaryPluginProtocols = "tls"
	CanaryPluginProtocolsTLSPassthrough CanaryPluginProtocols = "tls_passthrough"
	CanaryPluginProtocolsUDP            CanaryPluginProtocols = "udp"
	CanaryPluginProtocolsWs             CanaryPluginProtocols = "ws"
	CanaryPluginProtocolsWss            CanaryPluginProtocols = "wss"
)

func (e CanaryPluginProtocols) ToPointer() *CanaryPluginProtocols {
	return &e
}
func (e *CanaryPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CanaryPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CanaryPluginProtocols: %v", v)
	}
}

// CanaryPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CanaryPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CanaryPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CanaryPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CanaryPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CanaryPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CanaryPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CanaryPlugin struct {
	Config CanaryPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CanaryPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CanaryPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                 `json:"enabled,omitempty"`
	ID           *string               `json:"id,omitempty"`
	InstanceName *string               `json:"instance_name,omitempty"`
	name         string                `const:"canary" json:"name"`
	Ordering     *CanaryPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CanaryPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CanaryPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CanaryPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (c CanaryPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CanaryPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CanaryPlugin) GetConfig() CanaryPluginConfig {
	if o == nil {
		return CanaryPluginConfig{}
	}
	return o.Config
}

func (o *CanaryPlugin) GetConsumer() *CanaryPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CanaryPlugin) GetConsumerGroup() *CanaryPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CanaryPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CanaryPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CanaryPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CanaryPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CanaryPlugin) GetName() string {
	return "canary"
}

func (o *CanaryPlugin) GetOrdering() *CanaryPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CanaryPlugin) GetProtocols() []CanaryPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CanaryPlugin) GetRoute() *CanaryPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CanaryPlugin) GetService() *CanaryPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CanaryPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CanaryPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// CanaryPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CanaryPluginInput struct {
	Config CanaryPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CanaryPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CanaryPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                 `json:"enabled,omitempty"`
	ID           *string               `json:"id,omitempty"`
	InstanceName *string               `json:"instance_name,omitempty"`
	name         string                `const:"canary" json:"name"`
	Ordering     *CanaryPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CanaryPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CanaryPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CanaryPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (c CanaryPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CanaryPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CanaryPluginInput) GetConfig() CanaryPluginConfig {
	if o == nil {
		return CanaryPluginConfig{}
	}
	return o.Config
}

func (o *CanaryPluginInput) GetConsumer() *CanaryPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CanaryPluginInput) GetConsumerGroup() *CanaryPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CanaryPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CanaryPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CanaryPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CanaryPluginInput) GetName() string {
	return "canary"
}

func (o *CanaryPluginInput) GetOrdering() *CanaryPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CanaryPluginInput) GetProtocols() []CanaryPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CanaryPluginInput) GetRoute() *CanaryPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CanaryPluginInput) GetService() *CanaryPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CanaryPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
