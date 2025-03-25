// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type Templates struct {
	// Unique name for the template, can be called with `{template://NAME}`
	Name string `json:"name"`
	// Template string for this request, supports mustache-style `{{"{{"}}placeholders}}`
	Template string `json:"template"`
}

func (o *Templates) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Templates) GetTemplate() string {
	if o == nil {
		return ""
	}
	return o.Template
}

type AiPromptTemplatePluginConfig struct {
	// Set true to allow requests that don't call or match any template.
	AllowUntemplatedRequests *bool `json:"allow_untemplated_requests,omitempty"`
	// Set true to add the original request to the Kong log plugin(s) output.
	LogOriginalRequest *bool `json:"log_original_request,omitempty"`
	// max allowed body size allowed to be introspected
	MaxRequestBodySize *int64 `json:"max_request_body_size,omitempty"`
	// Array of templates available to the request context.
	Templates []Templates `json:"templates,omitempty"`
}

func (o *AiPromptTemplatePluginConfig) GetAllowUntemplatedRequests() *bool {
	if o == nil {
		return nil
	}
	return o.AllowUntemplatedRequests
}

func (o *AiPromptTemplatePluginConfig) GetLogOriginalRequest() *bool {
	if o == nil {
		return nil
	}
	return o.LogOriginalRequest
}

func (o *AiPromptTemplatePluginConfig) GetMaxRequestBodySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestBodySize
}

func (o *AiPromptTemplatePluginConfig) GetTemplates() []Templates {
	if o == nil {
		return nil
	}
	return o.Templates
}

// AiPromptTemplatePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AiPromptTemplatePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptTemplatePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptTemplatePluginConsumerGroup - If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
type AiPromptTemplatePluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptTemplatePluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiPromptTemplatePluginProtocols string

const (
	AiPromptTemplatePluginProtocolsGrpc  AiPromptTemplatePluginProtocols = "grpc"
	AiPromptTemplatePluginProtocolsGrpcs AiPromptTemplatePluginProtocols = "grpcs"
	AiPromptTemplatePluginProtocolsHTTP  AiPromptTemplatePluginProtocols = "http"
	AiPromptTemplatePluginProtocolsHTTPS AiPromptTemplatePluginProtocols = "https"
)

func (e AiPromptTemplatePluginProtocols) ToPointer() *AiPromptTemplatePluginProtocols {
	return &e
}
func (e *AiPromptTemplatePluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AiPromptTemplatePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiPromptTemplatePluginProtocols: %v", v)
	}
}

// AiPromptTemplatePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type AiPromptTemplatePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptTemplatePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptTemplatePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AiPromptTemplatePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptTemplatePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptTemplatePlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AiPromptTemplatePlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"ai-prompt-template" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                       `json:"updated_at,omitempty"`
	Config    AiPromptTemplatePluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AiPromptTemplatePluginConsumer `json:"consumer"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *AiPromptTemplatePluginConsumerGroup `json:"consumer_group"`
	// A set of strings representing HTTP protocols.
	Protocols []AiPromptTemplatePluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *AiPromptTemplatePluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiPromptTemplatePluginService `json:"service"`
}

func (a AiPromptTemplatePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiPromptTemplatePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiPromptTemplatePlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AiPromptTemplatePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiPromptTemplatePlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiPromptTemplatePlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiPromptTemplatePlugin) GetName() string {
	return "ai-prompt-template"
}

func (o *AiPromptTemplatePlugin) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiPromptTemplatePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiPromptTemplatePlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AiPromptTemplatePlugin) GetConfig() AiPromptTemplatePluginConfig {
	if o == nil {
		return AiPromptTemplatePluginConfig{}
	}
	return o.Config
}

func (o *AiPromptTemplatePlugin) GetConsumer() *AiPromptTemplatePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiPromptTemplatePlugin) GetConsumerGroup() *AiPromptTemplatePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiPromptTemplatePlugin) GetProtocols() []AiPromptTemplatePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiPromptTemplatePlugin) GetRoute() *AiPromptTemplatePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiPromptTemplatePlugin) GetService() *AiPromptTemplatePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// AiPromptTemplatePluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AiPromptTemplatePluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"ai-prompt-template" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string                     `json:"tags,omitempty"`
	Config AiPromptTemplatePluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AiPromptTemplatePluginConsumer `json:"consumer"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *AiPromptTemplatePluginConsumerGroup `json:"consumer_group"`
	// A set of strings representing HTTP protocols.
	Protocols []AiPromptTemplatePluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *AiPromptTemplatePluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiPromptTemplatePluginService `json:"service"`
}

func (a AiPromptTemplatePluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiPromptTemplatePluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiPromptTemplatePluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiPromptTemplatePluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiPromptTemplatePluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiPromptTemplatePluginInput) GetName() string {
	return "ai-prompt-template"
}

func (o *AiPromptTemplatePluginInput) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiPromptTemplatePluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiPromptTemplatePluginInput) GetConfig() AiPromptTemplatePluginConfig {
	if o == nil {
		return AiPromptTemplatePluginConfig{}
	}
	return o.Config
}

func (o *AiPromptTemplatePluginInput) GetConsumer() *AiPromptTemplatePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiPromptTemplatePluginInput) GetConsumerGroup() *AiPromptTemplatePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiPromptTemplatePluginInput) GetProtocols() []AiPromptTemplatePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiPromptTemplatePluginInput) GetRoute() *AiPromptTemplatePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiPromptTemplatePluginInput) GetService() *AiPromptTemplatePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
