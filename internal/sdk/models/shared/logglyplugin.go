// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type ClientErrorsSeverity string

const (
	ClientErrorsSeverityDebug   ClientErrorsSeverity = "debug"
	ClientErrorsSeverityInfo    ClientErrorsSeverity = "info"
	ClientErrorsSeverityNotice  ClientErrorsSeverity = "notice"
	ClientErrorsSeverityWarning ClientErrorsSeverity = "warning"
	ClientErrorsSeverityErr     ClientErrorsSeverity = "err"
	ClientErrorsSeverityCrit    ClientErrorsSeverity = "crit"
	ClientErrorsSeverityAlert   ClientErrorsSeverity = "alert"
	ClientErrorsSeverityEmerg   ClientErrorsSeverity = "emerg"
)

func (e ClientErrorsSeverity) ToPointer() *ClientErrorsSeverity {
	return &e
}
func (e *ClientErrorsSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = ClientErrorsSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClientErrorsSeverity: %v", v)
	}
}

type LogLevel string

const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelNotice  LogLevel = "notice"
	LogLevelWarning LogLevel = "warning"
	LogLevelErr     LogLevel = "err"
	LogLevelCrit    LogLevel = "crit"
	LogLevelAlert   LogLevel = "alert"
	LogLevelEmerg   LogLevel = "emerg"
)

func (e LogLevel) ToPointer() *LogLevel {
	return &e
}
func (e *LogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = LogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogLevel: %v", v)
	}
}

type ServerErrorsSeverity string

const (
	ServerErrorsSeverityDebug   ServerErrorsSeverity = "debug"
	ServerErrorsSeverityInfo    ServerErrorsSeverity = "info"
	ServerErrorsSeverityNotice  ServerErrorsSeverity = "notice"
	ServerErrorsSeverityWarning ServerErrorsSeverity = "warning"
	ServerErrorsSeverityErr     ServerErrorsSeverity = "err"
	ServerErrorsSeverityCrit    ServerErrorsSeverity = "crit"
	ServerErrorsSeverityAlert   ServerErrorsSeverity = "alert"
	ServerErrorsSeverityEmerg   ServerErrorsSeverity = "emerg"
)

func (e ServerErrorsSeverity) ToPointer() *ServerErrorsSeverity {
	return &e
}
func (e *ServerErrorsSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = ServerErrorsSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerErrorsSeverity: %v", v)
	}
}

type SuccessfulSeverity string

const (
	SuccessfulSeverityDebug   SuccessfulSeverity = "debug"
	SuccessfulSeverityInfo    SuccessfulSeverity = "info"
	SuccessfulSeverityNotice  SuccessfulSeverity = "notice"
	SuccessfulSeverityWarning SuccessfulSeverity = "warning"
	SuccessfulSeverityErr     SuccessfulSeverity = "err"
	SuccessfulSeverityCrit    SuccessfulSeverity = "crit"
	SuccessfulSeverityAlert   SuccessfulSeverity = "alert"
	SuccessfulSeverityEmerg   SuccessfulSeverity = "emerg"
)

func (e SuccessfulSeverity) ToPointer() *SuccessfulSeverity {
	return &e
}
func (e *SuccessfulSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = SuccessfulSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SuccessfulSeverity: %v", v)
	}
}

type LogglyPluginConfig struct {
	ClientErrorsSeverity *ClientErrorsSeverity `json:"client_errors_severity,omitempty"`
	// Lua code as a key-value map
	CustomFieldsByLua map[string]any `json:"custom_fields_by_lua,omitempty"`
	// A string representing a host name, such as example.com.
	Host     *string   `json:"host,omitempty"`
	Key      *string   `json:"key,omitempty"`
	LogLevel *LogLevel `json:"log_level,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port                 *int64                `json:"port,omitempty"`
	ServerErrorsSeverity *ServerErrorsSeverity `json:"server_errors_severity,omitempty"`
	SuccessfulSeverity   *SuccessfulSeverity   `json:"successful_severity,omitempty"`
	Tags                 []string              `json:"tags,omitempty"`
	Timeout              *float64              `json:"timeout,omitempty"`
}

func (o *LogglyPluginConfig) GetClientErrorsSeverity() *ClientErrorsSeverity {
	if o == nil {
		return nil
	}
	return o.ClientErrorsSeverity
}

func (o *LogglyPluginConfig) GetCustomFieldsByLua() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldsByLua
}

func (o *LogglyPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *LogglyPluginConfig) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *LogglyPluginConfig) GetLogLevel() *LogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *LogglyPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *LogglyPluginConfig) GetServerErrorsSeverity() *ServerErrorsSeverity {
	if o == nil {
		return nil
	}
	return o.ServerErrorsSeverity
}

func (o *LogglyPluginConfig) GetSuccessfulSeverity() *SuccessfulSeverity {
	if o == nil {
		return nil
	}
	return o.SuccessfulSeverity
}

func (o *LogglyPluginConfig) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LogglyPluginConfig) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

// LogglyPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type LogglyPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *LogglyPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type LogglyPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *LogglyPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type LogglyPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *LogglyPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type LogglyPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *LogglyPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type LogglyPluginOrdering struct {
	After  *LogglyPluginAfter  `json:"after,omitempty"`
	Before *LogglyPluginBefore `json:"before,omitempty"`
}

func (o *LogglyPluginOrdering) GetAfter() *LogglyPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *LogglyPluginOrdering) GetBefore() *LogglyPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type LogglyPluginProtocols string

const (
	LogglyPluginProtocolsGrpc           LogglyPluginProtocols = "grpc"
	LogglyPluginProtocolsGrpcs          LogglyPluginProtocols = "grpcs"
	LogglyPluginProtocolsHTTP           LogglyPluginProtocols = "http"
	LogglyPluginProtocolsHTTPS          LogglyPluginProtocols = "https"
	LogglyPluginProtocolsTCP            LogglyPluginProtocols = "tcp"
	LogglyPluginProtocolsTLS            LogglyPluginProtocols = "tls"
	LogglyPluginProtocolsTLSPassthrough LogglyPluginProtocols = "tls_passthrough"
	LogglyPluginProtocolsUDP            LogglyPluginProtocols = "udp"
	LogglyPluginProtocolsWs             LogglyPluginProtocols = "ws"
	LogglyPluginProtocolsWss            LogglyPluginProtocols = "wss"
)

func (e LogglyPluginProtocols) ToPointer() *LogglyPluginProtocols {
	return &e
}
func (e *LogglyPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = LogglyPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogglyPluginProtocols: %v", v)
	}
}

// LogglyPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type LogglyPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *LogglyPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// LogglyPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type LogglyPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *LogglyPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// LogglyPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type LogglyPlugin struct {
	Config LogglyPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *LogglyPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *LogglyPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                 `json:"enabled,omitempty"`
	ID           *string               `json:"id,omitempty"`
	InstanceName *string               `json:"instance_name,omitempty"`
	name         string                `const:"loggly" json:"name"`
	Ordering     *LogglyPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []LogglyPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *LogglyPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *LogglyPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (l LogglyPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LogglyPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LogglyPlugin) GetConfig() LogglyPluginConfig {
	if o == nil {
		return LogglyPluginConfig{}
	}
	return o.Config
}

func (o *LogglyPlugin) GetConsumer() *LogglyPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *LogglyPlugin) GetConsumerGroup() *LogglyPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *LogglyPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LogglyPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *LogglyPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LogglyPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *LogglyPlugin) GetName() string {
	return "loggly"
}

func (o *LogglyPlugin) GetOrdering() *LogglyPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *LogglyPlugin) GetProtocols() []LogglyPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *LogglyPlugin) GetRoute() *LogglyPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *LogglyPlugin) GetService() *LogglyPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *LogglyPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LogglyPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// LogglyPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type LogglyPluginInput struct {
	Config LogglyPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *LogglyPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *LogglyPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                 `json:"enabled,omitempty"`
	ID           *string               `json:"id,omitempty"`
	InstanceName *string               `json:"instance_name,omitempty"`
	name         string                `const:"loggly" json:"name"`
	Ordering     *LogglyPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []LogglyPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *LogglyPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *LogglyPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (l LogglyPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LogglyPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LogglyPluginInput) GetConfig() LogglyPluginConfig {
	if o == nil {
		return LogglyPluginConfig{}
	}
	return o.Config
}

func (o *LogglyPluginInput) GetConsumer() *LogglyPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *LogglyPluginInput) GetConsumerGroup() *LogglyPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *LogglyPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *LogglyPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LogglyPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *LogglyPluginInput) GetName() string {
	return "loggly"
}

func (o *LogglyPluginInput) GetOrdering() *LogglyPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *LogglyPluginInput) GetProtocols() []LogglyPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *LogglyPluginInput) GetRoute() *LogglyPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *LogglyPluginInput) GetService() *LogglyPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *LogglyPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}