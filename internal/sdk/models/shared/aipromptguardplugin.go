// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type AiPromptGuardPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiPromptGuardPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiPromptGuardPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiPromptGuardPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiPromptGuardPluginOrdering struct {
	After  *AiPromptGuardPluginAfter  `json:"after,omitempty"`
	Before *AiPromptGuardPluginBefore `json:"before,omitempty"`
}

func (o *AiPromptGuardPluginOrdering) GetAfter() *AiPromptGuardPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *AiPromptGuardPluginOrdering) GetBefore() *AiPromptGuardPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type AiPromptGuardPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *AiPromptGuardPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiPromptGuardPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AiPromptGuardPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// AiPromptGuardPluginLlmFormat - LLM input and output format and schema to use
type AiPromptGuardPluginLlmFormat string

const (
	AiPromptGuardPluginLlmFormatBedrock AiPromptGuardPluginLlmFormat = "bedrock"
	AiPromptGuardPluginLlmFormatGemini  AiPromptGuardPluginLlmFormat = "gemini"
	AiPromptGuardPluginLlmFormatOpenai  AiPromptGuardPluginLlmFormat = "openai"
)

func (e AiPromptGuardPluginLlmFormat) ToPointer() *AiPromptGuardPluginLlmFormat {
	return &e
}
func (e *AiPromptGuardPluginLlmFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bedrock":
		fallthrough
	case "gemini":
		fallthrough
	case "openai":
		*e = AiPromptGuardPluginLlmFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiPromptGuardPluginLlmFormat: %v", v)
	}
}

type AiPromptGuardPluginConfig struct {
	// If true, will ignore all previous chat prompts from the conversation history.
	AllowAllConversationHistory *bool `json:"allow_all_conversation_history,omitempty"`
	// Array of valid regex patterns, or valid questions from the 'user' role in chat.
	AllowPatterns []string `json:"allow_patterns,omitempty"`
	// Array of invalid regex patterns, or invalid questions from the 'user' role in chat.
	DenyPatterns []string `json:"deny_patterns,omitempty"`
	// LLM input and output format and schema to use
	LlmFormat *AiPromptGuardPluginLlmFormat `json:"llm_format,omitempty"`
	// If true, will match all roles in addition to 'user' role in conversation history.
	MatchAllRoles *bool `json:"match_all_roles,omitempty"`
	// max allowed body size allowed to be introspected
	MaxRequestBodySize *int64 `json:"max_request_body_size,omitempty"`
}

func (o *AiPromptGuardPluginConfig) GetAllowAllConversationHistory() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAllConversationHistory
}

func (o *AiPromptGuardPluginConfig) GetAllowPatterns() []string {
	if o == nil {
		return nil
	}
	return o.AllowPatterns
}

func (o *AiPromptGuardPluginConfig) GetDenyPatterns() []string {
	if o == nil {
		return nil
	}
	return o.DenyPatterns
}

func (o *AiPromptGuardPluginConfig) GetLlmFormat() *AiPromptGuardPluginLlmFormat {
	if o == nil {
		return nil
	}
	return o.LlmFormat
}

func (o *AiPromptGuardPluginConfig) GetMatchAllRoles() *bool {
	if o == nil {
		return nil
	}
	return o.MatchAllRoles
}

func (o *AiPromptGuardPluginConfig) GetMaxRequestBodySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestBodySize
}

// AiPromptGuardPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AiPromptGuardPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptGuardPluginConsumerGroup - If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
type AiPromptGuardPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiPromptGuardPluginProtocols string

const (
	AiPromptGuardPluginProtocolsGrpc  AiPromptGuardPluginProtocols = "grpc"
	AiPromptGuardPluginProtocolsGrpcs AiPromptGuardPluginProtocols = "grpcs"
	AiPromptGuardPluginProtocolsHTTP  AiPromptGuardPluginProtocols = "http"
	AiPromptGuardPluginProtocolsHTTPS AiPromptGuardPluginProtocols = "https"
)

func (e AiPromptGuardPluginProtocols) ToPointer() *AiPromptGuardPluginProtocols {
	return &e
}
func (e *AiPromptGuardPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AiPromptGuardPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiPromptGuardPluginProtocols: %v", v)
	}
}

// AiPromptGuardPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type AiPromptGuardPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptGuardPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AiPromptGuardPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptGuardPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AiPromptGuardPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                         `json:"enabled,omitempty"`
	ID           *string                       `json:"id,omitempty"`
	InstanceName *string                       `json:"instance_name,omitempty"`
	name         string                        `const:"ai-prompt-guard" json:"name"`
	Ordering     *AiPromptGuardPluginOrdering  `json:"ordering,omitempty"`
	Partials     []AiPromptGuardPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                     `json:"updated_at,omitempty"`
	Config    *AiPromptGuardPluginConfig `json:"config,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AiPromptGuardPluginConsumer `json:"consumer"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *AiPromptGuardPluginConsumerGroup `json:"consumer_group"`
	// A set of strings representing HTTP protocols.
	Protocols []AiPromptGuardPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *AiPromptGuardPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiPromptGuardPluginService `json:"service"`
}

func (a AiPromptGuardPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiPromptGuardPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiPromptGuardPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AiPromptGuardPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiPromptGuardPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiPromptGuardPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiPromptGuardPlugin) GetName() string {
	return "ai-prompt-guard"
}

func (o *AiPromptGuardPlugin) GetOrdering() *AiPromptGuardPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiPromptGuardPlugin) GetPartials() []AiPromptGuardPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *AiPromptGuardPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiPromptGuardPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AiPromptGuardPlugin) GetConfig() *AiPromptGuardPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *AiPromptGuardPlugin) GetConsumer() *AiPromptGuardPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiPromptGuardPlugin) GetConsumerGroup() *AiPromptGuardPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiPromptGuardPlugin) GetProtocols() []AiPromptGuardPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiPromptGuardPlugin) GetRoute() *AiPromptGuardPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiPromptGuardPlugin) GetService() *AiPromptGuardPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
