// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

// ParamLocation - Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.
type ParamLocation string

const (
	ParamLocationQuery ParamLocation = "query"
	ParamLocationBody  ParamLocation = "body"
)

func (e ParamLocation) ToPointer() *ParamLocation {
	return &e
}
func (e *ParamLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "query":
		fallthrough
	case "body":
		*e = ParamLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ParamLocation: %v", v)
	}
}

type Auth struct {
	// If enabled, the authorization header or parameter can be overridden in the request by the value configured in the plugin.
	AllowOverride *bool `json:"allow_override,omitempty"`
	// Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_ACCESS_KEY_ID environment variable for this plugin instance.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`
	// Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_SECRET_ACCESS_KEY environment variable for this plugin instance.
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`
	// If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client ID.
	AzureClientID *string `json:"azure_client_id,omitempty"`
	// If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client secret.
	AzureClientSecret *string `json:"azure_client_secret,omitempty"`
	// If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the tenant ID.
	AzureTenantID *string `json:"azure_tenant_id,omitempty"`
	// Set true to use the Azure Cloud Managed Identity (or user-assigned identity) to authenticate with Azure-provider models.
	AzureUseManagedIdentity *bool `json:"azure_use_managed_identity,omitempty"`
	// Set this field to the full JSON of the GCP service account to authenticate, if required. If null (and gcp_use_service_account is true), Kong will attempt to read from environment variable `GCP_SERVICE_ACCOUNT`.
	GcpServiceAccountJSON *string `json:"gcp_service_account_json,omitempty"`
	// Use service account auth for GCP-based providers and models.
	GcpUseServiceAccount *bool `json:"gcp_use_service_account,omitempty"`
	// If AI model requires authentication via Authorization or API key header, specify its name here.
	HeaderName *string `json:"header_name,omitempty"`
	// Specify the full auth header value for 'header_name', for example 'Bearer key' or just 'key'.
	HeaderValue *string `json:"header_value,omitempty"`
	// Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.
	ParamLocation *ParamLocation `json:"param_location,omitempty"`
	// If AI model requires authentication via query parameter, specify its name here.
	ParamName *string `json:"param_name,omitempty"`
	// Specify the full parameter value for 'param_name'.
	ParamValue *string `json:"param_value,omitempty"`
}

func (o *Auth) GetAllowOverride() *bool {
	if o == nil {
		return nil
	}
	return o.AllowOverride
}

func (o *Auth) GetAwsAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccessKeyID
}

func (o *Auth) GetAwsSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretAccessKey
}

func (o *Auth) GetAzureClientID() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientID
}

func (o *Auth) GetAzureClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientSecret
}

func (o *Auth) GetAzureTenantID() *string {
	if o == nil {
		return nil
	}
	return o.AzureTenantID
}

func (o *Auth) GetAzureUseManagedIdentity() *bool {
	if o == nil {
		return nil
	}
	return o.AzureUseManagedIdentity
}

func (o *Auth) GetGcpServiceAccountJSON() *string {
	if o == nil {
		return nil
	}
	return o.GcpServiceAccountJSON
}

func (o *Auth) GetGcpUseServiceAccount() *bool {
	if o == nil {
		return nil
	}
	return o.GcpUseServiceAccount
}

func (o *Auth) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *Auth) GetHeaderValue() *string {
	if o == nil {
		return nil
	}
	return o.HeaderValue
}

func (o *Auth) GetParamLocation() *ParamLocation {
	if o == nil {
		return nil
	}
	return o.ParamLocation
}

func (o *Auth) GetParamName() *string {
	if o == nil {
		return nil
	}
	return o.ParamName
}

func (o *Auth) GetParamValue() *string {
	if o == nil {
		return nil
	}
	return o.ParamValue
}

type Logging struct {
	// If enabled, will log the request and response body into the Kong log plugin(s) output.
	LogPayloads *bool `json:"log_payloads,omitempty"`
	// If enabled and supported by the driver, will add model usage and token metrics into the Kong log plugin(s) output.
	LogStatistics *bool `json:"log_statistics,omitempty"`
}

func (o *Logging) GetLogPayloads() *bool {
	if o == nil {
		return nil
	}
	return o.LogPayloads
}

func (o *Logging) GetLogStatistics() *bool {
	if o == nil {
		return nil
	}
	return o.LogStatistics
}

type Bedrock struct {
	// If using AWS providers (Bedrock) you can override the `AWS_REGION` environment variable by setting this option.
	AwsRegion *string `json:"aws_region,omitempty"`
}

func (o *Bedrock) GetAwsRegion() *string {
	if o == nil {
		return nil
	}
	return o.AwsRegion
}

type Gemini struct {
	// If running Gemini on Vertex, specify the regional API endpoint (hostname only).
	APIEndpoint *string `json:"api_endpoint,omitempty"`
	// If running Gemini on Vertex, specify the location ID.
	LocationID *string `json:"location_id,omitempty"`
	// If running Gemini on Vertex, specify the project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

func (o *Gemini) GetAPIEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.APIEndpoint
}

func (o *Gemini) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *Gemini) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

// Llama2Format - If using llama2 provider, select the upstream message format.
type Llama2Format string

const (
	Llama2FormatRaw    Llama2Format = "raw"
	Llama2FormatOpenai Llama2Format = "openai"
	Llama2FormatOllama Llama2Format = "ollama"
)

func (e Llama2Format) ToPointer() *Llama2Format {
	return &e
}
func (e *Llama2Format) UnmarshalJSON(data []byte) error {
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
		*e = Llama2Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Llama2Format: %v", v)
	}
}

// MistralFormat - If using mistral provider, select the upstream message format.
type MistralFormat string

const (
	MistralFormatOpenai MistralFormat = "openai"
	MistralFormatOllama MistralFormat = "ollama"
)

func (e MistralFormat) ToPointer() *MistralFormat {
	return &e
}
func (e *MistralFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		fallthrough
	case "ollama":
		*e = MistralFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MistralFormat: %v", v)
	}
}

// OptionsObj - Key/value settings for the model
type OptionsObj struct {
	// Defines the schema/API version, if using Anthropic provider.
	AnthropicVersion *string `json:"anthropic_version,omitempty"`
	// 'api-version' for Azure OpenAI instances.
	AzureAPIVersion *string `json:"azure_api_version,omitempty"`
	// Deployment ID for Azure OpenAI instances.
	AzureDeploymentID *string `json:"azure_deployment_id,omitempty"`
	// Instance name for Azure OpenAI hosted models.
	AzureInstance *string  `json:"azure_instance,omitempty"`
	Bedrock       *Bedrock `json:"bedrock,omitempty"`
	Gemini        *Gemini  `json:"gemini,omitempty"`
	// Defines the cost per 1M tokens in your prompt.
	InputCost *float64 `json:"input_cost,omitempty"`
	// If using llama2 provider, select the upstream message format.
	Llama2Format *Llama2Format `json:"llama2_format,omitempty"`
	// Defines the max_tokens, if using chat or completion models.
	MaxTokens *int64 `json:"max_tokens,omitempty"`
	// If using mistral provider, select the upstream message format.
	MistralFormat *MistralFormat `json:"mistral_format,omitempty"`
	// Defines the cost per 1M tokens in the output of the AI.
	OutputCost *float64 `json:"output_cost,omitempty"`
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

func (o *OptionsObj) GetAnthropicVersion() *string {
	if o == nil {
		return nil
	}
	return o.AnthropicVersion
}

func (o *OptionsObj) GetAzureAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.AzureAPIVersion
}

func (o *OptionsObj) GetAzureDeploymentID() *string {
	if o == nil {
		return nil
	}
	return o.AzureDeploymentID
}

func (o *OptionsObj) GetAzureInstance() *string {
	if o == nil {
		return nil
	}
	return o.AzureInstance
}

func (o *OptionsObj) GetBedrock() *Bedrock {
	if o == nil {
		return nil
	}
	return o.Bedrock
}

func (o *OptionsObj) GetGemini() *Gemini {
	if o == nil {
		return nil
	}
	return o.Gemini
}

func (o *OptionsObj) GetInputCost() *float64 {
	if o == nil {
		return nil
	}
	return o.InputCost
}

func (o *OptionsObj) GetLlama2Format() *Llama2Format {
	if o == nil {
		return nil
	}
	return o.Llama2Format
}

func (o *OptionsObj) GetMaxTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTokens
}

func (o *OptionsObj) GetMistralFormat() *MistralFormat {
	if o == nil {
		return nil
	}
	return o.MistralFormat
}

func (o *OptionsObj) GetOutputCost() *float64 {
	if o == nil {
		return nil
	}
	return o.OutputCost
}

func (o *OptionsObj) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *OptionsObj) GetTopK() *int64 {
	if o == nil {
		return nil
	}
	return o.TopK
}

func (o *OptionsObj) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}

func (o *OptionsObj) GetUpstreamPath() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamPath
}

func (o *OptionsObj) GetUpstreamURL() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamURL
}

// Provider - AI provider request format - Kong translates requests to and from the specified backend compatible formats.
type Provider string

const (
	ProviderOpenai    Provider = "openai"
	ProviderAzure     Provider = "azure"
	ProviderAnthropic Provider = "anthropic"
	ProviderCohere    Provider = "cohere"
	ProviderMistral   Provider = "mistral"
	ProviderLlama2    Provider = "llama2"
	ProviderGemini    Provider = "gemini"
	ProviderBedrock   Provider = "bedrock"
)

func (e Provider) ToPointer() *Provider {
	return &e
}
func (e *Provider) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "gemini":
		fallthrough
	case "bedrock":
		*e = Provider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Provider: %v", v)
	}
}

type Model struct {
	// Model name to execute.
	Name *string `json:"name,omitempty"`
	// Key/value settings for the model
	Options *OptionsObj `json:"options,omitempty"`
	// AI provider request format - Kong translates requests to and from the specified backend compatible formats.
	Provider *Provider `json:"provider,omitempty"`
}

func (o *Model) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Model) GetOptions() *OptionsObj {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *Model) GetProvider() *Provider {
	if o == nil {
		return nil
	}
	return o.Provider
}

// ResponseStreaming - Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events.
type ResponseStreaming string

const (
	ResponseStreamingAllow  ResponseStreaming = "allow"
	ResponseStreamingDeny   ResponseStreaming = "deny"
	ResponseStreamingAlways ResponseStreaming = "always"
)

func (e ResponseStreaming) ToPointer() *ResponseStreaming {
	return &e
}
func (e *ResponseStreaming) UnmarshalJSON(data []byte) error {
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
		*e = ResponseStreaming(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseStreaming: %v", v)
	}
}

// RouteType - The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
type RouteType string

const (
	RouteTypeLlmV1Chat        RouteType = "llm/v1/chat"
	RouteTypeLlmV1Completions RouteType = "llm/v1/completions"
	RouteTypePreserve         RouteType = "preserve"
)

func (e RouteType) ToPointer() *RouteType {
	return &e
}
func (e *RouteType) UnmarshalJSON(data []byte) error {
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
		*e = RouteType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RouteType: %v", v)
	}
}

type AiProxyPluginConfig struct {
	Auth    *Auth    `json:"auth,omitempty"`
	Logging *Logging `json:"logging,omitempty"`
	// max allowed body size allowed to be introspected
	MaxRequestBodySize *int64 `json:"max_request_body_size,omitempty"`
	Model              *Model `json:"model,omitempty"`
	// Display the model name selected in the X-Kong-LLM-Model response header
	ModelNameHeader *bool `json:"model_name_header,omitempty"`
	// Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events.
	ResponseStreaming *ResponseStreaming `json:"response_streaming,omitempty"`
	// The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
	RouteType *RouteType `json:"route_type,omitempty"`
}

func (o *AiProxyPluginConfig) GetAuth() *Auth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *AiProxyPluginConfig) GetLogging() *Logging {
	if o == nil {
		return nil
	}
	return o.Logging
}

func (o *AiProxyPluginConfig) GetMaxRequestBodySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestBodySize
}

func (o *AiProxyPluginConfig) GetModel() *Model {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *AiProxyPluginConfig) GetModelNameHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ModelNameHeader
}

func (o *AiProxyPluginConfig) GetResponseStreaming() *ResponseStreaming {
	if o == nil {
		return nil
	}
	return o.ResponseStreaming
}

func (o *AiProxyPluginConfig) GetRouteType() *RouteType {
	if o == nil {
		return nil
	}
	return o.RouteType
}

// AiProxyPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AiProxyPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiProxyPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiProxyPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiProxyPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiProxyPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiProxyPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiProxyPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiProxyPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiProxyPluginOrdering struct {
	After  *AiProxyPluginAfter  `json:"after,omitempty"`
	Before *AiProxyPluginBefore `json:"before,omitempty"`
}

func (o *AiProxyPluginOrdering) GetAfter() *AiProxyPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *AiProxyPluginOrdering) GetBefore() *AiProxyPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type AiProxyPluginProtocols string

const (
	AiProxyPluginProtocolsGrpc           AiProxyPluginProtocols = "grpc"
	AiProxyPluginProtocolsGrpcs          AiProxyPluginProtocols = "grpcs"
	AiProxyPluginProtocolsHTTP           AiProxyPluginProtocols = "http"
	AiProxyPluginProtocolsHTTPS          AiProxyPluginProtocols = "https"
	AiProxyPluginProtocolsTCP            AiProxyPluginProtocols = "tcp"
	AiProxyPluginProtocolsTLS            AiProxyPluginProtocols = "tls"
	AiProxyPluginProtocolsTLSPassthrough AiProxyPluginProtocols = "tls_passthrough"
	AiProxyPluginProtocolsUDP            AiProxyPluginProtocols = "udp"
	AiProxyPluginProtocolsWs             AiProxyPluginProtocols = "ws"
	AiProxyPluginProtocolsWss            AiProxyPluginProtocols = "wss"
)

func (e AiProxyPluginProtocols) ToPointer() *AiProxyPluginProtocols {
	return &e
}
func (e *AiProxyPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AiProxyPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiProxyPluginProtocols: %v", v)
	}
}

// AiProxyPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AiProxyPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiProxyPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiProxyPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AiProxyPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiProxyPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiProxyPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AiProxyPlugin struct {
	Config AiProxyPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *AiProxyPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *AiProxyPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                  `json:"enabled,omitempty"`
	ID           *string                `json:"id,omitempty"`
	InstanceName *string                `json:"instance_name,omitempty"`
	name         string                 `const:"ai-proxy" json:"name"`
	Ordering     *AiProxyPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AiProxyPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AiProxyPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiProxyPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (a AiProxyPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiProxyPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiProxyPlugin) GetConfig() AiProxyPluginConfig {
	if o == nil {
		return AiProxyPluginConfig{}
	}
	return o.Config
}

func (o *AiProxyPlugin) GetConsumer() *AiProxyPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiProxyPlugin) GetConsumerGroup() *AiProxyPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiProxyPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AiProxyPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiProxyPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiProxyPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiProxyPlugin) GetName() string {
	return "ai-proxy"
}

func (o *AiProxyPlugin) GetOrdering() *AiProxyPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiProxyPlugin) GetProtocols() []AiProxyPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiProxyPlugin) GetRoute() *AiProxyPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiProxyPlugin) GetService() *AiProxyPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *AiProxyPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiProxyPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// AiProxyPluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AiProxyPluginInput struct {
	Config AiProxyPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *AiProxyPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *AiProxyPluginConsumerGroup `json:"consumer_group,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                  `json:"enabled,omitempty"`
	ID           *string                `json:"id,omitempty"`
	InstanceName *string                `json:"instance_name,omitempty"`
	name         string                 `const:"ai-proxy" json:"name"`
	Ordering     *AiProxyPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AiProxyPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AiProxyPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiProxyPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (a AiProxyPluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiProxyPluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiProxyPluginInput) GetConfig() AiProxyPluginConfig {
	if o == nil {
		return AiProxyPluginConfig{}
	}
	return o.Config
}

func (o *AiProxyPluginInput) GetConsumer() *AiProxyPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiProxyPluginInput) GetConsumerGroup() *AiProxyPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiProxyPluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiProxyPluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiProxyPluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiProxyPluginInput) GetName() string {
	return "ai-proxy"
}

func (o *AiProxyPluginInput) GetOrdering() *AiProxyPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiProxyPluginInput) GetProtocols() []AiProxyPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiProxyPluginInput) GetRoute() *AiProxyPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiProxyPluginInput) GetService() *AiProxyPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *AiProxyPluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
