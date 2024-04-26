// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type AIPromptDecoratorPluginProtocols string

const (
	AIPromptDecoratorPluginProtocolsGrpc           AIPromptDecoratorPluginProtocols = "grpc"
	AIPromptDecoratorPluginProtocolsGrpcs          AIPromptDecoratorPluginProtocols = "grpcs"
	AIPromptDecoratorPluginProtocolsHTTP           AIPromptDecoratorPluginProtocols = "http"
	AIPromptDecoratorPluginProtocolsHTTPS          AIPromptDecoratorPluginProtocols = "https"
	AIPromptDecoratorPluginProtocolsTCP            AIPromptDecoratorPluginProtocols = "tcp"
	AIPromptDecoratorPluginProtocolsTLS            AIPromptDecoratorPluginProtocols = "tls"
	AIPromptDecoratorPluginProtocolsTLSPassthrough AIPromptDecoratorPluginProtocols = "tls_passthrough"
	AIPromptDecoratorPluginProtocolsUDP            AIPromptDecoratorPluginProtocols = "udp"
	AIPromptDecoratorPluginProtocolsWs             AIPromptDecoratorPluginProtocols = "ws"
	AIPromptDecoratorPluginProtocolsWss            AIPromptDecoratorPluginProtocols = "wss"
)

func (e AIPromptDecoratorPluginProtocols) ToPointer() *AIPromptDecoratorPluginProtocols {
	return &e
}

func (e *AIPromptDecoratorPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AIPromptDecoratorPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AIPromptDecoratorPluginProtocols: %v", v)
	}
}

// AIPromptDecoratorPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AIPromptDecoratorPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptDecoratorPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AIPromptDecoratorPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AIPromptDecoratorPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptDecoratorPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AIPromptDecoratorPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AIPromptDecoratorPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptDecoratorPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Role string

const (
	RoleSystem    Role = "system"
	RoleAssistant Role = "assistant"
	RoleUser      Role = "user"
)

func (e Role) ToPointer() *Role {
	return &e
}

func (e *Role) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system":
		fallthrough
	case "assistant":
		fallthrough
	case "user":
		*e = Role(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Role: %v", v)
	}
}

type AIPromptDecoratorPluginAppend struct {
	Content *string `json:"content,omitempty"`
	Role    *Role   `default:"system" json:"role"`
}

func (a AIPromptDecoratorPluginAppend) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AIPromptDecoratorPluginAppend) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AIPromptDecoratorPluginAppend) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *AIPromptDecoratorPluginAppend) GetRole() *Role {
	if o == nil {
		return nil
	}
	return o.Role
}

type AIPromptDecoratorPluginRole string

const (
	AIPromptDecoratorPluginRoleSystem    AIPromptDecoratorPluginRole = "system"
	AIPromptDecoratorPluginRoleAssistant AIPromptDecoratorPluginRole = "assistant"
	AIPromptDecoratorPluginRoleUser      AIPromptDecoratorPluginRole = "user"
)

func (e AIPromptDecoratorPluginRole) ToPointer() *AIPromptDecoratorPluginRole {
	return &e
}

func (e *AIPromptDecoratorPluginRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system":
		fallthrough
	case "assistant":
		fallthrough
	case "user":
		*e = AIPromptDecoratorPluginRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AIPromptDecoratorPluginRole: %v", v)
	}
}

type Prepend struct {
	Content *string                      `json:"content,omitempty"`
	Role    *AIPromptDecoratorPluginRole `default:"system" json:"role"`
}

func (p Prepend) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Prepend) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Prepend) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *Prepend) GetRole() *AIPromptDecoratorPluginRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type Prompts struct {
	// Insert chat messages at the end of the chat message array. This array preserves exact order when adding messages.
	Append []AIPromptDecoratorPluginAppend `json:"append,omitempty"`
	// Insert chat messages at the beginning of the chat message array. This array preserves exact order when adding messages.
	Prepend []Prepend `json:"prepend,omitempty"`
}

func (o *Prompts) GetAppend() []AIPromptDecoratorPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *Prompts) GetPrepend() []Prepend {
	if o == nil {
		return nil
	}
	return o.Prepend
}

type AIPromptDecoratorPluginConfig struct {
	Prompts *Prompts `json:"prompts,omitempty"`
}

func (o *AIPromptDecoratorPluginConfig) GetPrompts() *Prompts {
	if o == nil {
		return nil
	}
	return o.Prompts
}

// AIPromptDecoratorPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AIPromptDecoratorPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"ai-prompt-decorator" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AIPromptDecoratorPluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AIPromptDecoratorPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AIPromptDecoratorPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AIPromptDecoratorPluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64                        `json:"created_at,omitempty"`
	ID        *string                       `json:"id,omitempty"`
	Config    AIPromptDecoratorPluginConfig `json:"config"`
}

func (a AIPromptDecoratorPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AIPromptDecoratorPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AIPromptDecoratorPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AIPromptDecoratorPlugin) GetName() string {
	return "ai-prompt-decorator"
}

func (o *AIPromptDecoratorPlugin) GetProtocols() []AIPromptDecoratorPluginProtocols {
	if o == nil {
		return []AIPromptDecoratorPluginProtocols{}
	}
	return o.Protocols
}

func (o *AIPromptDecoratorPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AIPromptDecoratorPlugin) GetConsumer() *AIPromptDecoratorPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AIPromptDecoratorPlugin) GetRoute() *AIPromptDecoratorPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AIPromptDecoratorPlugin) GetService() *AIPromptDecoratorPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *AIPromptDecoratorPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AIPromptDecoratorPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AIPromptDecoratorPlugin) GetConfig() AIPromptDecoratorPluginConfig {
	if o == nil {
		return AIPromptDecoratorPluginConfig{}
	}
	return o.Config
}
