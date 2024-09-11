// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateAiProxyPluginParamLocation - Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.
type CreateAiProxyPluginParamLocation string

const (
	CreateAiProxyPluginParamLocationQuery CreateAiProxyPluginParamLocation = "query"
	CreateAiProxyPluginParamLocationBody  CreateAiProxyPluginParamLocation = "body"
)

func (e CreateAiProxyPluginParamLocation) ToPointer() *CreateAiProxyPluginParamLocation {
	return &e
}
func (e *CreateAiProxyPluginParamLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "query":
		fallthrough
	case "body":
		*e = CreateAiProxyPluginParamLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAiProxyPluginParamLocation: %v", v)
	}
}

type CreateAiProxyPluginAuth struct {
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
	ParamLocation *CreateAiProxyPluginParamLocation `json:"param_location,omitempty"`
	// If AI model requires authentication via query parameter, specify its name here.
	ParamName *string `json:"param_name,omitempty"`
	// Specify the full parameter value for 'param_name'.
	ParamValue *string `json:"param_value,omitempty"`
}

func (o *CreateAiProxyPluginAuth) GetAllowOverride() *bool {
	if o == nil {
		return nil
	}
	return o.AllowOverride
}

func (o *CreateAiProxyPluginAuth) GetAwsAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccessKeyID
}

func (o *CreateAiProxyPluginAuth) GetAwsSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretAccessKey
}

func (o *CreateAiProxyPluginAuth) GetAzureClientID() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientID
}

func (o *CreateAiProxyPluginAuth) GetAzureClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientSecret
}

func (o *CreateAiProxyPluginAuth) GetAzureTenantID() *string {
	if o == nil {
		return nil
	}
	return o.AzureTenantID
}

func (o *CreateAiProxyPluginAuth) GetAzureUseManagedIdentity() *bool {
	if o == nil {
		return nil
	}
	return o.AzureUseManagedIdentity
}

func (o *CreateAiProxyPluginAuth) GetGcpServiceAccountJSON() *string {
	if o == nil {
		return nil
	}
	return o.GcpServiceAccountJSON
}

func (o *CreateAiProxyPluginAuth) GetGcpUseServiceAccount() *bool {
	if o == nil {
		return nil
	}
	return o.GcpUseServiceAccount
}

func (o *CreateAiProxyPluginAuth) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *CreateAiProxyPluginAuth) GetHeaderValue() *string {
	if o == nil {
		return nil
	}
	return o.HeaderValue
}

func (o *CreateAiProxyPluginAuth) GetParamLocation() *CreateAiProxyPluginParamLocation {
	if o == nil {
		return nil
	}
	return o.ParamLocation
}

func (o *CreateAiProxyPluginAuth) GetParamName() *string {
	if o == nil {
		return nil
	}
	return o.ParamName
}

func (o *CreateAiProxyPluginAuth) GetParamValue() *string {
	if o == nil {
		return nil
	}
	return o.ParamValue
}

type CreateAiProxyPluginLogging struct {
	// If enabled, will log the request and response body into the Kong log plugin(s) output.
	LogPayloads *bool `json:"log_payloads,omitempty"`
	// If enabled and supported by the driver, will add model usage and token metrics into the Kong log plugin(s) output.
	LogStatistics *bool `json:"log_statistics,omitempty"`
}

func (o *CreateAiProxyPluginLogging) GetLogPayloads() *bool {
	if o == nil {
		return nil
	}
	return o.LogPayloads
}

func (o *CreateAiProxyPluginLogging) GetLogStatistics() *bool {
	if o == nil {
		return nil
	}
	return o.LogStatistics
}

type CreateAiProxyPluginBedrock struct {
	// If using AWS providers (Bedrock) you can override the `AWS_REGION` environment variable by setting this option.
	AwsRegion *string `json:"aws_region,omitempty"`
}

func (o *CreateAiProxyPluginBedrock) GetAwsRegion() *string {
	if o == nil {
		return nil
	}
	return o.AwsRegion
}

type CreateAiProxyPluginGemini struct {
	// If running Gemini on Vertex, specify the regional API endpoint (hostname only).
	APIEndpoint *string `json:"api_endpoint,omitempty"`
	// If running Gemini on Vertex, specify the location ID.
	LocationID *string `json:"location_id,omitempty"`
	// If running Gemini on Vertex, specify the project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

func (o *CreateAiProxyPluginGemini) GetAPIEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.APIEndpoint
}

func (o *CreateAiProxyPluginGemini) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *CreateAiProxyPluginGemini) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

// CreateAiProxyPluginLlama2Format - If using llama2 provider, select the upstream message format.
type CreateAiProxyPluginLlama2Format string

const (
	CreateAiProxyPluginLlama2FormatRaw    CreateAiProxyPluginLlama2Format = "raw"
	CreateAiProxyPluginLlama2FormatOpenai CreateAiProxyPluginLlama2Format = "openai"
	CreateAiProxyPluginLlama2FormatOllama CreateAiProxyPluginLlama2Format = "ollama"
)

func (e CreateAiProxyPluginLlama2Format) ToPointer() *CreateAiProxyPluginLlama2Format {
	return &e
}
func (e *CreateAiProxyPluginLlama2Format) UnmarshalJSON(data []byte) error {
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
		*e = CreateAiProxyPluginLlama2Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAiProxyPluginLlama2Format: %v", v)
	}
}

// CreateAiProxyPluginMistralFormat - If using mistral provider, select the upstream message format.
type CreateAiProxyPluginMistralFormat string

const (
	CreateAiProxyPluginMistralFormatOpenai CreateAiProxyPluginMistralFormat = "openai"
	CreateAiProxyPluginMistralFormatOllama CreateAiProxyPluginMistralFormat = "ollama"
)

func (e CreateAiProxyPluginMistralFormat) ToPointer() *CreateAiProxyPluginMistralFormat {
	return &e
}
func (e *CreateAiProxyPluginMistralFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		fallthrough
	case "ollama":
		*e = CreateAiProxyPluginMistralFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAiProxyPluginMistralFormat: %v", v)
	}
}

// CreateAiProxyPluginOptions - Key/value settings for the model
type CreateAiProxyPluginOptions struct {
	// Defines the schema/API version, if using Anthropic provider.
	AnthropicVersion *string `json:"anthropic_version,omitempty"`
	// 'api-version' for Azure OpenAI instances.
	AzureAPIVersion *string `json:"azure_api_version,omitempty"`
	// Deployment ID for Azure OpenAI instances.
	AzureDeploymentID *string `json:"azure_deployment_id,omitempty"`
	// Instance name for Azure OpenAI hosted models.
	AzureInstance *string                     `json:"azure_instance,omitempty"`
	Bedrock       *CreateAiProxyPluginBedrock `json:"bedrock,omitempty"`
	Gemini        *CreateAiProxyPluginGemini  `json:"gemini,omitempty"`
	// Defines the cost per 1M tokens in your prompt.
	InputCost *float64 `json:"input_cost,omitempty"`
	// If using llama2 provider, select the upstream message format.
	Llama2Format *CreateAiProxyPluginLlama2Format `json:"llama2_format,omitempty"`
	// Defines the max_tokens, if using chat or completion models.
	MaxTokens *int64 `json:"max_tokens,omitempty"`
	// If using mistral provider, select the upstream message format.
	MistralFormat *CreateAiProxyPluginMistralFormat `json:"mistral_format,omitempty"`
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

func (o *CreateAiProxyPluginOptions) GetAnthropicVersion() *string {
	if o == nil {
		return nil
	}
	return o.AnthropicVersion
}

func (o *CreateAiProxyPluginOptions) GetAzureAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.AzureAPIVersion
}

func (o *CreateAiProxyPluginOptions) GetAzureDeploymentID() *string {
	if o == nil {
		return nil
	}
	return o.AzureDeploymentID
}

func (o *CreateAiProxyPluginOptions) GetAzureInstance() *string {
	if o == nil {
		return nil
	}
	return o.AzureInstance
}

func (o *CreateAiProxyPluginOptions) GetBedrock() *CreateAiProxyPluginBedrock {
	if o == nil {
		return nil
	}
	return o.Bedrock
}

func (o *CreateAiProxyPluginOptions) GetGemini() *CreateAiProxyPluginGemini {
	if o == nil {
		return nil
	}
	return o.Gemini
}

func (o *CreateAiProxyPluginOptions) GetInputCost() *float64 {
	if o == nil {
		return nil
	}
	return o.InputCost
}

func (o *CreateAiProxyPluginOptions) GetLlama2Format() *CreateAiProxyPluginLlama2Format {
	if o == nil {
		return nil
	}
	return o.Llama2Format
}

func (o *CreateAiProxyPluginOptions) GetMaxTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTokens
}

func (o *CreateAiProxyPluginOptions) GetMistralFormat() *CreateAiProxyPluginMistralFormat {
	if o == nil {
		return nil
	}
	return o.MistralFormat
}

func (o *CreateAiProxyPluginOptions) GetOutputCost() *float64 {
	if o == nil {
		return nil
	}
	return o.OutputCost
}

func (o *CreateAiProxyPluginOptions) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *CreateAiProxyPluginOptions) GetTopK() *int64 {
	if o == nil {
		return nil
	}
	return o.TopK
}

func (o *CreateAiProxyPluginOptions) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}

func (o *CreateAiProxyPluginOptions) GetUpstreamPath() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamPath
}

func (o *CreateAiProxyPluginOptions) GetUpstreamURL() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamURL
}

// CreateAiProxyPluginProvider - AI provider request format - Kong translates requests to and from the specified backend compatible formats.
type CreateAiProxyPluginProvider string

const (
	CreateAiProxyPluginProviderOpenai    CreateAiProxyPluginProvider = "openai"
	CreateAiProxyPluginProviderAzure     CreateAiProxyPluginProvider = "azure"
	CreateAiProxyPluginProviderAnthropic CreateAiProxyPluginProvider = "anthropic"
	CreateAiProxyPluginProviderCohere    CreateAiProxyPluginProvider = "cohere"
	CreateAiProxyPluginProviderMistral   CreateAiProxyPluginProvider = "mistral"
	CreateAiProxyPluginProviderLlama2    CreateAiProxyPluginProvider = "llama2"
	CreateAiProxyPluginProviderGemini    CreateAiProxyPluginProvider = "gemini"
	CreateAiProxyPluginProviderBedrock   CreateAiProxyPluginProvider = "bedrock"
)

func (e CreateAiProxyPluginProvider) ToPointer() *CreateAiProxyPluginProvider {
	return &e
}
func (e *CreateAiProxyPluginProvider) UnmarshalJSON(data []byte) error {
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
		*e = CreateAiProxyPluginProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAiProxyPluginProvider: %v", v)
	}
}

type CreateAiProxyPluginModel struct {
	// Model name to execute.
	Name *string `json:"name,omitempty"`
	// Key/value settings for the model
	Options *CreateAiProxyPluginOptions `json:"options,omitempty"`
	// AI provider request format - Kong translates requests to and from the specified backend compatible formats.
	Provider *CreateAiProxyPluginProvider `json:"provider,omitempty"`
}

func (o *CreateAiProxyPluginModel) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateAiProxyPluginModel) GetOptions() *CreateAiProxyPluginOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *CreateAiProxyPluginModel) GetProvider() *CreateAiProxyPluginProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

// CreateAiProxyPluginResponseStreaming - Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events.
type CreateAiProxyPluginResponseStreaming string

const (
	CreateAiProxyPluginResponseStreamingAllow  CreateAiProxyPluginResponseStreaming = "allow"
	CreateAiProxyPluginResponseStreamingDeny   CreateAiProxyPluginResponseStreaming = "deny"
	CreateAiProxyPluginResponseStreamingAlways CreateAiProxyPluginResponseStreaming = "always"
)

func (e CreateAiProxyPluginResponseStreaming) ToPointer() *CreateAiProxyPluginResponseStreaming {
	return &e
}
func (e *CreateAiProxyPluginResponseStreaming) UnmarshalJSON(data []byte) error {
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
		*e = CreateAiProxyPluginResponseStreaming(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAiProxyPluginResponseStreaming: %v", v)
	}
}

// CreateAiProxyPluginRouteType - The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
type CreateAiProxyPluginRouteType string

const (
	CreateAiProxyPluginRouteTypeLlmV1Chat        CreateAiProxyPluginRouteType = "llm/v1/chat"
	CreateAiProxyPluginRouteTypeLlmV1Completions CreateAiProxyPluginRouteType = "llm/v1/completions"
	CreateAiProxyPluginRouteTypePreserve         CreateAiProxyPluginRouteType = "preserve"
)

func (e CreateAiProxyPluginRouteType) ToPointer() *CreateAiProxyPluginRouteType {
	return &e
}
func (e *CreateAiProxyPluginRouteType) UnmarshalJSON(data []byte) error {
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
		*e = CreateAiProxyPluginRouteType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAiProxyPluginRouteType: %v", v)
	}
}

type CreateAiProxyPluginConfig struct {
	Auth    *CreateAiProxyPluginAuth    `json:"auth,omitempty"`
	Logging *CreateAiProxyPluginLogging `json:"logging,omitempty"`
	// max allowed body size allowed to be introspected
	MaxRequestBodySize *int64                    `json:"max_request_body_size,omitempty"`
	Model              *CreateAiProxyPluginModel `json:"model,omitempty"`
	// Display the model name selected in the X-Kong-LLM-Model response header
	ModelNameHeader *bool `json:"model_name_header,omitempty"`
	// Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events.
	ResponseStreaming *CreateAiProxyPluginResponseStreaming `json:"response_streaming,omitempty"`
	// The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
	RouteType *CreateAiProxyPluginRouteType `json:"route_type,omitempty"`
}

func (o *CreateAiProxyPluginConfig) GetAuth() *CreateAiProxyPluginAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *CreateAiProxyPluginConfig) GetLogging() *CreateAiProxyPluginLogging {
	if o == nil {
		return nil
	}
	return o.Logging
}

func (o *CreateAiProxyPluginConfig) GetMaxRequestBodySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestBodySize
}

func (o *CreateAiProxyPluginConfig) GetModel() *CreateAiProxyPluginModel {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *CreateAiProxyPluginConfig) GetModelNameHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ModelNameHeader
}

func (o *CreateAiProxyPluginConfig) GetResponseStreaming() *CreateAiProxyPluginResponseStreaming {
	if o == nil {
		return nil
	}
	return o.ResponseStreaming
}

func (o *CreateAiProxyPluginConfig) GetRouteType() *CreateAiProxyPluginRouteType {
	if o == nil {
		return nil
	}
	return o.RouteType
}

type CreateAiProxyPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateAiProxyPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateAiProxyPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateAiProxyPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateAiProxyPluginOrdering struct {
	After  *CreateAiProxyPluginAfter  `json:"after,omitempty"`
	Before *CreateAiProxyPluginBefore `json:"before,omitempty"`
}

func (o *CreateAiProxyPluginOrdering) GetAfter() *CreateAiProxyPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateAiProxyPluginOrdering) GetBefore() *CreateAiProxyPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateAiProxyPluginProtocols string

const (
	CreateAiProxyPluginProtocolsGrpc           CreateAiProxyPluginProtocols = "grpc"
	CreateAiProxyPluginProtocolsGrpcs          CreateAiProxyPluginProtocols = "grpcs"
	CreateAiProxyPluginProtocolsHTTP           CreateAiProxyPluginProtocols = "http"
	CreateAiProxyPluginProtocolsHTTPS          CreateAiProxyPluginProtocols = "https"
	CreateAiProxyPluginProtocolsTCP            CreateAiProxyPluginProtocols = "tcp"
	CreateAiProxyPluginProtocolsTLS            CreateAiProxyPluginProtocols = "tls"
	CreateAiProxyPluginProtocolsTLSPassthrough CreateAiProxyPluginProtocols = "tls_passthrough"
	CreateAiProxyPluginProtocolsUDP            CreateAiProxyPluginProtocols = "udp"
	CreateAiProxyPluginProtocolsWs             CreateAiProxyPluginProtocols = "ws"
	CreateAiProxyPluginProtocolsWss            CreateAiProxyPluginProtocols = "wss"
)

func (e CreateAiProxyPluginProtocols) ToPointer() *CreateAiProxyPluginProtocols {
	return &e
}
func (e *CreateAiProxyPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateAiProxyPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAiProxyPluginProtocols: %v", v)
	}
}

// CreateAiProxyPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateAiProxyPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAiProxyPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAiProxyPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAiProxyPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAiProxyPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateAiProxyPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAiProxyPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAiProxyPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateAiProxyPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAiProxyPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAiProxyPlugin struct {
	Config *CreateAiProxyPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         *string                      `const:"ai-proxy" json:"name,omitempty"`
	Ordering     *CreateAiProxyPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateAiProxyPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateAiProxyPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateAiProxyPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateAiProxyPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateAiProxyPluginService `json:"service,omitempty"`
}

func (c CreateAiProxyPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAiProxyPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAiProxyPlugin) GetConfig() *CreateAiProxyPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateAiProxyPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateAiProxyPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateAiProxyPlugin) GetName() *string {
	return types.String("ai-proxy")
}

func (o *CreateAiProxyPlugin) GetOrdering() *CreateAiProxyPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateAiProxyPlugin) GetProtocols() []CreateAiProxyPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateAiProxyPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateAiProxyPlugin) GetConsumer() *CreateAiProxyPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateAiProxyPlugin) GetConsumerGroup() *CreateAiProxyPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateAiProxyPlugin) GetRoute() *CreateAiProxyPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateAiProxyPlugin) GetService() *CreateAiProxyPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
