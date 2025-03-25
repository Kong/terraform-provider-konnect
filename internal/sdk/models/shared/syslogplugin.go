// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type SyslogPluginClientErrorsSeverity string

const (
	SyslogPluginClientErrorsSeverityAlert   SyslogPluginClientErrorsSeverity = "alert"
	SyslogPluginClientErrorsSeverityCrit    SyslogPluginClientErrorsSeverity = "crit"
	SyslogPluginClientErrorsSeverityDebug   SyslogPluginClientErrorsSeverity = "debug"
	SyslogPluginClientErrorsSeverityEmerg   SyslogPluginClientErrorsSeverity = "emerg"
	SyslogPluginClientErrorsSeverityErr     SyslogPluginClientErrorsSeverity = "err"
	SyslogPluginClientErrorsSeverityInfo    SyslogPluginClientErrorsSeverity = "info"
	SyslogPluginClientErrorsSeverityNotice  SyslogPluginClientErrorsSeverity = "notice"
	SyslogPluginClientErrorsSeverityWarning SyslogPluginClientErrorsSeverity = "warning"
)

func (e SyslogPluginClientErrorsSeverity) ToPointer() *SyslogPluginClientErrorsSeverity {
	return &e
}
func (e *SyslogPluginClientErrorsSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "alert":
		fallthrough
	case "crit":
		fallthrough
	case "debug":
		fallthrough
	case "emerg":
		fallthrough
	case "err":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		*e = SyslogPluginClientErrorsSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SyslogPluginClientErrorsSeverity: %v", v)
	}
}

// Facility - The facility is used by the operating system to decide how to handle each log message.
type Facility string

const (
	FacilityAuth     Facility = "auth"
	FacilityAuthpriv Facility = "authpriv"
	FacilityCron     Facility = "cron"
	FacilityDaemon   Facility = "daemon"
	FacilityFtp      Facility = "ftp"
	FacilityKern     Facility = "kern"
	FacilityLocal0   Facility = "local0"
	FacilityLocal1   Facility = "local1"
	FacilityLocal2   Facility = "local2"
	FacilityLocal3   Facility = "local3"
	FacilityLocal4   Facility = "local4"
	FacilityLocal5   Facility = "local5"
	FacilityLocal6   Facility = "local6"
	FacilityLocal7   Facility = "local7"
	FacilityLpr      Facility = "lpr"
	FacilityMail     Facility = "mail"
	FacilityNews     Facility = "news"
	FacilitySyslog   Facility = "syslog"
	FacilityUser     Facility = "user"
	FacilityUucp     Facility = "uucp"
)

func (e Facility) ToPointer() *Facility {
	return &e
}
func (e *Facility) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auth":
		fallthrough
	case "authpriv":
		fallthrough
	case "cron":
		fallthrough
	case "daemon":
		fallthrough
	case "ftp":
		fallthrough
	case "kern":
		fallthrough
	case "local0":
		fallthrough
	case "local1":
		fallthrough
	case "local2":
		fallthrough
	case "local3":
		fallthrough
	case "local4":
		fallthrough
	case "local5":
		fallthrough
	case "local6":
		fallthrough
	case "local7":
		fallthrough
	case "lpr":
		fallthrough
	case "mail":
		fallthrough
	case "news":
		fallthrough
	case "syslog":
		fallthrough
	case "user":
		fallthrough
	case "uucp":
		*e = Facility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Facility: %v", v)
	}
}

type SyslogPluginLogLevel string

const (
	SyslogPluginLogLevelAlert   SyslogPluginLogLevel = "alert"
	SyslogPluginLogLevelCrit    SyslogPluginLogLevel = "crit"
	SyslogPluginLogLevelDebug   SyslogPluginLogLevel = "debug"
	SyslogPluginLogLevelEmerg   SyslogPluginLogLevel = "emerg"
	SyslogPluginLogLevelErr     SyslogPluginLogLevel = "err"
	SyslogPluginLogLevelInfo    SyslogPluginLogLevel = "info"
	SyslogPluginLogLevelNotice  SyslogPluginLogLevel = "notice"
	SyslogPluginLogLevelWarning SyslogPluginLogLevel = "warning"
)

func (e SyslogPluginLogLevel) ToPointer() *SyslogPluginLogLevel {
	return &e
}
func (e *SyslogPluginLogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "alert":
		fallthrough
	case "crit":
		fallthrough
	case "debug":
		fallthrough
	case "emerg":
		fallthrough
	case "err":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		*e = SyslogPluginLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SyslogPluginLogLevel: %v", v)
	}
}

type SyslogPluginServerErrorsSeverity string

const (
	SyslogPluginServerErrorsSeverityAlert   SyslogPluginServerErrorsSeverity = "alert"
	SyslogPluginServerErrorsSeverityCrit    SyslogPluginServerErrorsSeverity = "crit"
	SyslogPluginServerErrorsSeverityDebug   SyslogPluginServerErrorsSeverity = "debug"
	SyslogPluginServerErrorsSeverityEmerg   SyslogPluginServerErrorsSeverity = "emerg"
	SyslogPluginServerErrorsSeverityErr     SyslogPluginServerErrorsSeverity = "err"
	SyslogPluginServerErrorsSeverityInfo    SyslogPluginServerErrorsSeverity = "info"
	SyslogPluginServerErrorsSeverityNotice  SyslogPluginServerErrorsSeverity = "notice"
	SyslogPluginServerErrorsSeverityWarning SyslogPluginServerErrorsSeverity = "warning"
)

func (e SyslogPluginServerErrorsSeverity) ToPointer() *SyslogPluginServerErrorsSeverity {
	return &e
}
func (e *SyslogPluginServerErrorsSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "alert":
		fallthrough
	case "crit":
		fallthrough
	case "debug":
		fallthrough
	case "emerg":
		fallthrough
	case "err":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		*e = SyslogPluginServerErrorsSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SyslogPluginServerErrorsSeverity: %v", v)
	}
}

type SyslogPluginSuccessfulSeverity string

const (
	SyslogPluginSuccessfulSeverityAlert   SyslogPluginSuccessfulSeverity = "alert"
	SyslogPluginSuccessfulSeverityCrit    SyslogPluginSuccessfulSeverity = "crit"
	SyslogPluginSuccessfulSeverityDebug   SyslogPluginSuccessfulSeverity = "debug"
	SyslogPluginSuccessfulSeverityEmerg   SyslogPluginSuccessfulSeverity = "emerg"
	SyslogPluginSuccessfulSeverityErr     SyslogPluginSuccessfulSeverity = "err"
	SyslogPluginSuccessfulSeverityInfo    SyslogPluginSuccessfulSeverity = "info"
	SyslogPluginSuccessfulSeverityNotice  SyslogPluginSuccessfulSeverity = "notice"
	SyslogPluginSuccessfulSeverityWarning SyslogPluginSuccessfulSeverity = "warning"
)

func (e SyslogPluginSuccessfulSeverity) ToPointer() *SyslogPluginSuccessfulSeverity {
	return &e
}
func (e *SyslogPluginSuccessfulSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "alert":
		fallthrough
	case "crit":
		fallthrough
	case "debug":
		fallthrough
	case "emerg":
		fallthrough
	case "err":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		*e = SyslogPluginSuccessfulSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SyslogPluginSuccessfulSeverity: %v", v)
	}
}

type SyslogPluginConfig struct {
	ClientErrorsSeverity *SyslogPluginClientErrorsSeverity `json:"client_errors_severity,omitempty"`
	// Lua code as a key-value map
	CustomFieldsByLua map[string]any `json:"custom_fields_by_lua,omitempty"`
	// The facility is used by the operating system to decide how to handle each log message.
	Facility             *Facility                         `json:"facility,omitempty"`
	LogLevel             *SyslogPluginLogLevel             `json:"log_level,omitempty"`
	ServerErrorsSeverity *SyslogPluginServerErrorsSeverity `json:"server_errors_severity,omitempty"`
	SuccessfulSeverity   *SyslogPluginSuccessfulSeverity   `json:"successful_severity,omitempty"`
}

func (o *SyslogPluginConfig) GetClientErrorsSeverity() *SyslogPluginClientErrorsSeverity {
	if o == nil {
		return nil
	}
	return o.ClientErrorsSeverity
}

func (o *SyslogPluginConfig) GetCustomFieldsByLua() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldsByLua
}

func (o *SyslogPluginConfig) GetFacility() *Facility {
	if o == nil {
		return nil
	}
	return o.Facility
}

func (o *SyslogPluginConfig) GetLogLevel() *SyslogPluginLogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *SyslogPluginConfig) GetServerErrorsSeverity() *SyslogPluginServerErrorsSeverity {
	if o == nil {
		return nil
	}
	return o.ServerErrorsSeverity
}

func (o *SyslogPluginConfig) GetSuccessfulSeverity() *SyslogPluginSuccessfulSeverity {
	if o == nil {
		return nil
	}
	return o.SuccessfulSeverity
}

// SyslogPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type SyslogPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *SyslogPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SyslogPluginProtocols - A string representing a protocol, such as HTTP or HTTPS.
type SyslogPluginProtocols string

const (
	SyslogPluginProtocolsGrpc           SyslogPluginProtocols = "grpc"
	SyslogPluginProtocolsGrpcs          SyslogPluginProtocols = "grpcs"
	SyslogPluginProtocolsHTTP           SyslogPluginProtocols = "http"
	SyslogPluginProtocolsHTTPS          SyslogPluginProtocols = "https"
	SyslogPluginProtocolsTCP            SyslogPluginProtocols = "tcp"
	SyslogPluginProtocolsTLS            SyslogPluginProtocols = "tls"
	SyslogPluginProtocolsTLSPassthrough SyslogPluginProtocols = "tls_passthrough"
	SyslogPluginProtocolsUDP            SyslogPluginProtocols = "udp"
	SyslogPluginProtocolsWs             SyslogPluginProtocols = "ws"
	SyslogPluginProtocolsWss            SyslogPluginProtocols = "wss"
)

func (e SyslogPluginProtocols) ToPointer() *SyslogPluginProtocols {
	return &e
}
func (e *SyslogPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = SyslogPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SyslogPluginProtocols: %v", v)
	}
}

// SyslogPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type SyslogPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *SyslogPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SyslogPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type SyslogPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *SyslogPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SyslogPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type SyslogPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"syslog" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64             `json:"updated_at,omitempty"`
	Config    SyslogPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *SyslogPluginConsumer `json:"consumer"`
	// A set of strings representing protocols.
	Protocols []SyslogPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *SyslogPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *SyslogPluginService `json:"service"`
}

func (s SyslogPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SyslogPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SyslogPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SyslogPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SyslogPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SyslogPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *SyslogPlugin) GetName() string {
	return "syslog"
}

func (o *SyslogPlugin) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *SyslogPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *SyslogPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *SyslogPlugin) GetConfig() SyslogPluginConfig {
	if o == nil {
		return SyslogPluginConfig{}
	}
	return o.Config
}

func (o *SyslogPlugin) GetConsumer() *SyslogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *SyslogPlugin) GetProtocols() []SyslogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *SyslogPlugin) GetRoute() *SyslogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *SyslogPlugin) GetService() *SyslogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// SyslogPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type SyslogPluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"syslog" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string           `json:"tags,omitempty"`
	Config SyslogPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *SyslogPluginConsumer `json:"consumer"`
	// A set of strings representing protocols.
	Protocols []SyslogPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *SyslogPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *SyslogPluginService `json:"service"`
}

func (s SyslogPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SyslogPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SyslogPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SyslogPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SyslogPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *SyslogPluginInput) GetName() string {
	return "syslog"
}

func (o *SyslogPluginInput) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *SyslogPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *SyslogPluginInput) GetConfig() SyslogPluginConfig {
	if o == nil {
		return SyslogPluginConfig{}
	}
	return o.Config
}

func (o *SyslogPluginInput) GetConsumer() *SyslogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *SyslogPluginInput) GetProtocols() []SyslogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *SyslogPluginInput) GetRoute() *SyslogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *SyslogPluginInput) GetService() *SyslogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
