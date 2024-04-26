// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type FileLogPluginProtocols string

const (
	FileLogPluginProtocolsGrpc           FileLogPluginProtocols = "grpc"
	FileLogPluginProtocolsGrpcs          FileLogPluginProtocols = "grpcs"
	FileLogPluginProtocolsHTTP           FileLogPluginProtocols = "http"
	FileLogPluginProtocolsHTTPS          FileLogPluginProtocols = "https"
	FileLogPluginProtocolsTCP            FileLogPluginProtocols = "tcp"
	FileLogPluginProtocolsTLS            FileLogPluginProtocols = "tls"
	FileLogPluginProtocolsTLSPassthrough FileLogPluginProtocols = "tls_passthrough"
	FileLogPluginProtocolsUDP            FileLogPluginProtocols = "udp"
	FileLogPluginProtocolsWs             FileLogPluginProtocols = "ws"
	FileLogPluginProtocolsWss            FileLogPluginProtocols = "wss"
)

func (e FileLogPluginProtocols) ToPointer() *FileLogPluginProtocols {
	return &e
}

func (e *FileLogPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = FileLogPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileLogPluginProtocols: %v", v)
	}
}

// FileLogPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type FileLogPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *FileLogPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// FileLogPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type FileLogPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *FileLogPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// FileLogPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type FileLogPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *FileLogPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type FileLogPluginConfig struct {
	// Lua code as a key-value map
	CustomFieldsByLua map[string]interface{} `json:"custom_fields_by_lua,omitempty"`
	// The file path of the output log file. The plugin creates the log file if it doesn't exist yet.
	Path *string `json:"path,omitempty"`
	// Determines whether the log file is closed and reopened on every request.
	Reopen *bool `default:"false" json:"reopen"`
}

func (f FileLogPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FileLogPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FileLogPluginConfig) GetCustomFieldsByLua() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.CustomFieldsByLua
}

func (o *FileLogPluginConfig) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *FileLogPluginConfig) GetReopen() *bool {
	if o == nil {
		return nil
	}
	return o.Reopen
}

// FileLogPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type FileLogPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"file-log" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []FileLogPluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *FileLogPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *FileLogPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *FileLogPluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64              `json:"created_at,omitempty"`
	ID        *string             `json:"id,omitempty"`
	Config    FileLogPluginConfig `json:"config"`
}

func (f FileLogPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FileLogPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FileLogPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *FileLogPlugin) GetName() string {
	return "file-log"
}

func (o *FileLogPlugin) GetProtocols() []FileLogPluginProtocols {
	if o == nil {
		return []FileLogPluginProtocols{}
	}
	return o.Protocols
}

func (o *FileLogPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *FileLogPlugin) GetConsumer() *FileLogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *FileLogPlugin) GetRoute() *FileLogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *FileLogPlugin) GetService() *FileLogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *FileLogPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *FileLogPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FileLogPlugin) GetConfig() FileLogPluginConfig {
	if o == nil {
		return FileLogPluginConfig{}
	}
	return o.Config
}
