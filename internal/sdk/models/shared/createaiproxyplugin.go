// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateAIProxyPluginParamLocation - Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.
type CreateAIProxyPluginParamLocation string

const (
	CreateAIProxyPluginParamLocationQuery CreateAIProxyPluginParamLocation = "query"
	CreateAIProxyPluginParamLocationBody  CreateAIProxyPluginParamLocation = "body"
)

func (e CreateAIProxyPluginParamLocation) ToPointer() *CreateAIProxyPluginParamLocation {
	return &e
}
func (e *CreateAIProxyPluginParamLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "query":
		fallthrough
	case "body":
		*e = CreateAIProxyPluginParamLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIProxyPluginParamLocation: %v", v)
	}
}

type CreateAIProxyPluginAuth struct {
	// If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client ID.
	AzureClientID *string `json:"azure_client_id,omitempty"`
	// If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client secret.
	AzureClientSecret *string `json:"azure_client_secret,omitempty"`
	// If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the tenant ID.
	AzureTenantID *string `json:"azure_tenant_id,omitempty"`
	// Set true to use the Azure Cloud Managed Identity (or user-assigned identity) to authenticate with Azure-provider models.
	AzureUseManagedIdentity *bool `json:"azure_use_managed_identity,omitempty"`
	// If AI model requires authentication via Authorization or API key header, specify its name here.
	HeaderName *string `json:"header_name,omitempty"`
	// Specify the full auth header value for 'header_name', for example 'Bearer key' or just 'key'.
	HeaderValue *string `json:"header_value,omitempty"`
	// Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.
	ParamLocation *CreateAIProxyPluginParamLocation `json:"param_location,omitempty"`
	// If AI model requires authentication via query parameter, specify its name here.
	ParamName *string `json:"param_name,omitempty"`
	// Specify the full parameter value for 'param_name'.
	ParamValue *string `json:"param_value,omitempty"`
}

func (o *CreateAIProxyPluginAuth) GetAzureClientID() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientID
}

func (o *CreateAIProxyPluginAuth) GetAzureClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientSecret
}

func (o *CreateAIProxyPluginAuth) GetAzureTenantID() *string {
	if o == nil {
		return nil
	}
	return o.AzureTenantID
}

func (o *CreateAIProxyPluginAuth) GetAzureUseManagedIdentity() *bool {
	if o == nil {
		return nil
	}
	return o.AzureUseManagedIdentity
}

func (o *CreateAIProxyPluginAuth) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *CreateAIProxyPluginAuth) GetHeaderValue() *string {
	if o == nil {
		return nil
	}
	return o.HeaderValue
}

func (o *CreateAIProxyPluginAuth) GetParamLocation() *CreateAIProxyPluginParamLocation {
	if o == nil {
		return nil
	}
	return o.ParamLocation
}

func (o *CreateAIProxyPluginAuth) GetParamName() *string {
	if o == nil {
		return nil
	}
	return o.ParamName
}

func (o *CreateAIProxyPluginAuth) GetParamValue() *string {
	if o == nil {
		return nil
	}
	return o.ParamValue
}

type CreateAIProxyPluginLogging struct {
	// If enabled, will log the request and response body into the Kong log plugin(s) output.
	LogPayloads *bool `json:"log_payloads,omitempty"`
	// If enabled and supported by the driver, will add model usage and token metrics into the Kong log plugin(s) output.
	LogStatistics *bool `json:"log_statistics,omitempty"`
}

func (o *CreateAIProxyPluginLogging) GetLogPayloads() *bool {
	if o == nil {
		return nil
	}
	return o.LogPayloads
}

func (o *CreateAIProxyPluginLogging) GetLogStatistics() *bool {
	if o == nil {
		return nil
	}
	return o.LogStatistics
}

// CreateAIProxyPluginLlama2Format - If using llama2 provider, select the upstream message format.
type CreateAIProxyPluginLlama2Format string

const (
	CreateAIProxyPluginLlama2FormatRaw    CreateAIProxyPluginLlama2Format = "raw"
	CreateAIProxyPluginLlama2FormatOpenai CreateAIProxyPluginLlama2Format = "openai"
	CreateAIProxyPluginLlama2FormatOllama CreateAIProxyPluginLlama2Format = "ollama"
)

func (e CreateAIProxyPluginLlama2Format) ToPointer() *CreateAIProxyPluginLlama2Format {
	return &e
}
func (e *CreateAIProxyPluginLlama2Format) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "raw":
		fallthrough
	case "openai":
		fallthrough
	case "ollama":
		*e = CreateAIProxyPluginLlama2Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIProxyPluginLlama2Format: %v", v)
	}
}

// CreateAIProxyPluginMistralFormat - If using mistral provider, select the upstream message format.
type CreateAIProxyPluginMistralFormat string

const (
	CreateAIProxyPluginMistralFormatOpenai CreateAIProxyPluginMistralFormat = "openai"
	CreateAIProxyPluginMistralFormatOllama CreateAIProxyPluginMistralFormat = "ollama"
)

func (e CreateAIProxyPluginMistralFormat) ToPointer() *CreateAIProxyPluginMistralFormat {
	return &e
}
func (e *CreateAIProxyPluginMistralFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		fallthrough
	case "ollama":
		*e = CreateAIProxyPluginMistralFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIProxyPluginMistralFormat: %v", v)
	}
}

// CreateAIProxyPluginOptions - Key/value settings for the model
type CreateAIProxyPluginOptions struct {
	// Defines the schema/API version, if using Anthropic provider.
	AnthropicVersion *string `json:"anthropic_version,omitempty"`
	// 'api-version' for Azure OpenAI instances.
	AzureAPIVersion *string `json:"azure_api_version,omitempty"`
	// Deployment ID for Azure OpenAI instances.
	AzureDeploymentID *string `json:"azure_deployment_id,omitempty"`
	// Instance name for Azure OpenAI hosted models.
	AzureInstance *string `json:"azure_instance,omitempty"`
	// If using llama2 provider, select the upstream message format.
	Llama2Format *CreateAIProxyPluginLlama2Format `json:"llama2_format,omitempty"`
	// Defines the max_tokens, if using chat or completion models.
	MaxTokens *int64 `json:"max_tokens,omitempty"`
	// If using mistral provider, select the upstream message format.
	MistralFormat *CreateAIProxyPluginMistralFormat `json:"mistral_format,omitempty"`
	// Defines the matching temperature, if using chat or completion models.
	Temperature *float64 `json:"temperature,omitempty"`
	// Defines the top-k most likely tokens, if supported.
	TopK *int64 `json:"top_k,omitempty"`
	// Defines the top-p probability mass, if supported.
	TopP *float64 `json:"top_p,omitempty"`
	// Manually specify or override the AI operation path, used when e.g. using the 'preserve' route_type.
	UpstreamPath *string `json:"upstream_path,omitempty"`
	// Manually specify or override the full URL to the AI operation endpoints, when calling (self-)hosted models, or for running via a private endpoint.
	UpstreamURL *string `json:"upstream_url,omitempty"`
}

func (o *CreateAIProxyPluginOptions) GetAnthropicVersion() *string {
	if o == nil {
		return nil
	}
	return o.AnthropicVersion
}

func (o *CreateAIProxyPluginOptions) GetAzureAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.AzureAPIVersion
}

func (o *CreateAIProxyPluginOptions) GetAzureDeploymentID() *string {
	if o == nil {
		return nil
	}
	return o.AzureDeploymentID
}

func (o *CreateAIProxyPluginOptions) GetAzureInstance() *string {
	if o == nil {
		return nil
	}
	return o.AzureInstance
}

func (o *CreateAIProxyPluginOptions) GetLlama2Format() *CreateAIProxyPluginLlama2Format {
	if o == nil {
		return nil
	}
	return o.Llama2Format
}

func (o *CreateAIProxyPluginOptions) GetMaxTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTokens
}

func (o *CreateAIProxyPluginOptions) GetMistralFormat() *CreateAIProxyPluginMistralFormat {
	if o == nil {
		return nil
	}
	return o.MistralFormat
}

func (o *CreateAIProxyPluginOptions) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *CreateAIProxyPluginOptions) GetTopK() *int64 {
	if o == nil {
		return nil
	}
	return o.TopK
}

func (o *CreateAIProxyPluginOptions) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}

func (o *CreateAIProxyPluginOptions) GetUpstreamPath() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamPath
}

func (o *CreateAIProxyPluginOptions) GetUpstreamURL() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamURL
}

// CreateAIProxyPluginProvider - AI provider request format - Kong translates requests to and from the specified backend compatible formats.
type CreateAIProxyPluginProvider string

const (
	CreateAIProxyPluginProviderOpenai    CreateAIProxyPluginProvider = "openai"
	CreateAIProxyPluginProviderAzure     CreateAIProxyPluginProvider = "azure"
	CreateAIProxyPluginProviderAnthropic CreateAIProxyPluginProvider = "anthropic"
	CreateAIProxyPluginProviderCohere    CreateAIProxyPluginProvider = "cohere"
	CreateAIProxyPluginProviderMistral   CreateAIProxyPluginProvider = "mistral"
	CreateAIProxyPluginProviderLlama2    CreateAIProxyPluginProvider = "llama2"
)

func (e CreateAIProxyPluginProvider) ToPointer() *CreateAIProxyPluginProvider {
	return &e
}
func (e *CreateAIProxyPluginProvider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		fallthrough
	case "azure":
		fallthrough
	case "anthropic":
		fallthrough
	case "cohere":
		fallthrough
	case "mistral":
		fallthrough
	case "llama2":
		*e = CreateAIProxyPluginProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIProxyPluginProvider: %v", v)
	}
}

type CreateAIProxyPluginModel struct {
	// Model name to execute.
	Name *string `json:"name,omitempty"`
	// Key/value settings for the model
	Options *CreateAIProxyPluginOptions `json:"options,omitempty"`
	// AI provider request format - Kong translates requests to and from the specified backend compatible formats.
	Provider *CreateAIProxyPluginProvider `json:"provider,omitempty"`
}

func (o *CreateAIProxyPluginModel) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateAIProxyPluginModel) GetOptions() *CreateAIProxyPluginOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *CreateAIProxyPluginModel) GetProvider() *CreateAIProxyPluginProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

// CreateAIProxyPluginResponseStreaming - Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events.
type CreateAIProxyPluginResponseStreaming string

const (
	CreateAIProxyPluginResponseStreamingAllow  CreateAIProxyPluginResponseStreaming = "allow"
	CreateAIProxyPluginResponseStreamingDeny   CreateAIProxyPluginResponseStreaming = "deny"
	CreateAIProxyPluginResponseStreamingAlways CreateAIProxyPluginResponseStreaming = "always"
)

func (e CreateAIProxyPluginResponseStreaming) ToPointer() *CreateAIProxyPluginResponseStreaming {
	return &e
}
func (e *CreateAIProxyPluginResponseStreaming) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		fallthrough
	case "deny":
		fallthrough
	case "always":
		*e = CreateAIProxyPluginResponseStreaming(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIProxyPluginResponseStreaming: %v", v)
	}
}

// CreateAIProxyPluginRouteType - The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
type CreateAIProxyPluginRouteType string

const (
	CreateAIProxyPluginRouteTypeLlmV1Chat        CreateAIProxyPluginRouteType = "llm/v1/chat"
	CreateAIProxyPluginRouteTypeLlmV1Completions CreateAIProxyPluginRouteType = "llm/v1/completions"
	CreateAIProxyPluginRouteTypePreserve         CreateAIProxyPluginRouteType = "preserve"
)

func (e CreateAIProxyPluginRouteType) ToPointer() *CreateAIProxyPluginRouteType {
	return &e
}
func (e *CreateAIProxyPluginRouteType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "llm/v1/chat":
		fallthrough
	case "llm/v1/completions":
		fallthrough
	case "preserve":
		*e = CreateAIProxyPluginRouteType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIProxyPluginRouteType: %v", v)
	}
}

type CreateAIProxyPluginConfig struct {
	Auth    *CreateAIProxyPluginAuth    `json:"auth,omitempty"`
	Logging *CreateAIProxyPluginLogging `json:"logging,omitempty"`
	Model   *CreateAIProxyPluginModel   `json:"model,omitempty"`
	// Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events.
	ResponseStreaming *CreateAIProxyPluginResponseStreaming `json:"response_streaming,omitempty"`
	// The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
	RouteType *CreateAIProxyPluginRouteType `json:"route_type,omitempty"`
}

func (o *CreateAIProxyPluginConfig) GetAuth() *CreateAIProxyPluginAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *CreateAIProxyPluginConfig) GetLogging() *CreateAIProxyPluginLogging {
	if o == nil {
		return nil
	}
	return o.Logging
}

func (o *CreateAIProxyPluginConfig) GetModel() *CreateAIProxyPluginModel {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *CreateAIProxyPluginConfig) GetResponseStreaming() *CreateAIProxyPluginResponseStreaming {
	if o == nil {
		return nil
	}
	return o.ResponseStreaming
}

func (o *CreateAIProxyPluginConfig) GetRouteType() *CreateAIProxyPluginRouteType {
	if o == nil {
		return nil
	}
	return o.RouteType
}

type CreateAIProxyPluginProtocols string

const (
	CreateAIProxyPluginProtocolsGrpc           CreateAIProxyPluginProtocols = "grpc"
	CreateAIProxyPluginProtocolsGrpcs          CreateAIProxyPluginProtocols = "grpcs"
	CreateAIProxyPluginProtocolsHTTP           CreateAIProxyPluginProtocols = "http"
	CreateAIProxyPluginProtocolsHTTPS          CreateAIProxyPluginProtocols = "https"
	CreateAIProxyPluginProtocolsTCP            CreateAIProxyPluginProtocols = "tcp"
	CreateAIProxyPluginProtocolsTLS            CreateAIProxyPluginProtocols = "tls"
	CreateAIProxyPluginProtocolsTLSPassthrough CreateAIProxyPluginProtocols = "tls_passthrough"
	CreateAIProxyPluginProtocolsUDP            CreateAIProxyPluginProtocols = "udp"
	CreateAIProxyPluginProtocolsWs             CreateAIProxyPluginProtocols = "ws"
	CreateAIProxyPluginProtocolsWss            CreateAIProxyPluginProtocols = "wss"
)

func (e CreateAIProxyPluginProtocols) ToPointer() *CreateAIProxyPluginProtocols {
	return &e
}
func (e *CreateAIProxyPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateAIProxyPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIProxyPluginProtocols: %v", v)
	}
}

// CreateAIProxyPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateAIProxyPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIProxyPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAIProxyPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIProxyPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAIProxyPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateAIProxyPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIProxyPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAIProxyPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateAIProxyPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIProxyPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAIProxyPlugin struct {
	Config *CreateAIProxyPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"ai-proxy" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateAIProxyPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateAIProxyPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateAIProxyPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateAIProxyPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateAIProxyPluginService `json:"service,omitempty"`
}

func (c CreateAIProxyPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAIProxyPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAIProxyPlugin) GetConfig() *CreateAIProxyPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateAIProxyPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateAIProxyPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateAIProxyPlugin) GetName() *string {
	return types.String("ai-proxy")
}

func (o *CreateAIProxyPlugin) GetProtocols() []CreateAIProxyPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateAIProxyPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateAIProxyPlugin) GetConsumer() *CreateAIProxyPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateAIProxyPlugin) GetConsumerGroup() *CreateAIProxyPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateAIProxyPlugin) GetRoute() *CreateAIProxyPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateAIProxyPlugin) GetService() *CreateAIProxyPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}