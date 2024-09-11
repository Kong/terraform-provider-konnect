// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// AiResponseTransformerPluginParamLocation - Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.
type AiResponseTransformerPluginParamLocation string

const (
	AiResponseTransformerPluginParamLocationQuery AiResponseTransformerPluginParamLocation = "query"
	AiResponseTransformerPluginParamLocationBody  AiResponseTransformerPluginParamLocation = "body"
)

func (e AiResponseTransformerPluginParamLocation) ToPointer() *AiResponseTransformerPluginParamLocation {
	return &e
}
func (e *AiResponseTransformerPluginParamLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "query":
		fallthrough
	case "body":
		*e = AiResponseTransformerPluginParamLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiResponseTransformerPluginParamLocation: %v", v)
	}
}

type AiResponseTransformerPluginAuth struct {
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
	ParamLocation *AiResponseTransformerPluginParamLocation `json:"param_location,omitempty"`
	// If AI model requires authentication via query parameter, specify its name here.
	ParamName *string `json:"param_name,omitempty"`
	// Specify the full parameter value for 'param_name'.
	ParamValue *string `json:"param_value,omitempty"`
}

func (o *AiResponseTransformerPluginAuth) GetAllowOverride() *bool {
	if o == nil {
		return nil
	}
	return o.AllowOverride
}

func (o *AiResponseTransformerPluginAuth) GetAwsAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccessKeyID
}

func (o *AiResponseTransformerPluginAuth) GetAwsSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretAccessKey
}

func (o *AiResponseTransformerPluginAuth) GetAzureClientID() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientID
}

func (o *AiResponseTransformerPluginAuth) GetAzureClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientSecret
}

func (o *AiResponseTransformerPluginAuth) GetAzureTenantID() *string {
	if o == nil {
		return nil
	}
	return o.AzureTenantID
}

func (o *AiResponseTransformerPluginAuth) GetAzureUseManagedIdentity() *bool {
	if o == nil {
		return nil
	}
	return o.AzureUseManagedIdentity
}

func (o *AiResponseTransformerPluginAuth) GetGcpServiceAccountJSON() *string {
	if o == nil {
		return nil
	}
	return o.GcpServiceAccountJSON
}

func (o *AiResponseTransformerPluginAuth) GetGcpUseServiceAccount() *bool {
	if o == nil {
		return nil
	}
	return o.GcpUseServiceAccount
}

func (o *AiResponseTransformerPluginAuth) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *AiResponseTransformerPluginAuth) GetHeaderValue() *string {
	if o == nil {
		return nil
	}
	return o.HeaderValue
}

func (o *AiResponseTransformerPluginAuth) GetParamLocation() *AiResponseTransformerPluginParamLocation {
	if o == nil {
		return nil
	}
	return o.ParamLocation
}

func (o *AiResponseTransformerPluginAuth) GetParamName() *string {
	if o == nil {
		return nil
	}
	return o.ParamName
}

func (o *AiResponseTransformerPluginAuth) GetParamValue() *string {
	if o == nil {
		return nil
	}
	return o.ParamValue
}

type AiResponseTransformerPluginLogging struct {
	// If enabled, will log the request and response body into the Kong log plugin(s) output.
	LogPayloads *bool `json:"log_payloads,omitempty"`
	// If enabled and supported by the driver, will add model usage and token metrics into the Kong log plugin(s) output.
	LogStatistics *bool `json:"log_statistics,omitempty"`
}

func (o *AiResponseTransformerPluginLogging) GetLogPayloads() *bool {
	if o == nil {
		return nil
	}
	return o.LogPayloads
}

func (o *AiResponseTransformerPluginLogging) GetLogStatistics() *bool {
	if o == nil {
		return nil
	}
	return o.LogStatistics
}

type AiResponseTransformerPluginBedrock struct {
	// If using AWS providers (Bedrock) you can override the `AWS_REGION` environment variable by setting this option.
	AwsRegion *string `json:"aws_region,omitempty"`
}

func (o *AiResponseTransformerPluginBedrock) GetAwsRegion() *string {
	if o == nil {
		return nil
	}
	return o.AwsRegion
}

type AiResponseTransformerPluginGemini struct {
	// If running Gemini on Vertex, specify the regional API endpoint (hostname only).
	APIEndpoint *string `json:"api_endpoint,omitempty"`
	// If running Gemini on Vertex, specify the location ID.
	LocationID *string `json:"location_id,omitempty"`
	// If running Gemini on Vertex, specify the project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

func (o *AiResponseTransformerPluginGemini) GetAPIEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.APIEndpoint
}

func (o *AiResponseTransformerPluginGemini) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *AiResponseTransformerPluginGemini) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

// AiResponseTransformerPluginLlama2Format - If using llama2 provider, select the upstream message format.
type AiResponseTransformerPluginLlama2Format string

const (
	AiResponseTransformerPluginLlama2FormatRaw    AiResponseTransformerPluginLlama2Format = "raw"
	AiResponseTransformerPluginLlama2FormatOpenai AiResponseTransformerPluginLlama2Format = "openai"
	AiResponseTransformerPluginLlama2FormatOllama AiResponseTransformerPluginLlama2Format = "ollama"
)

func (e AiResponseTransformerPluginLlama2Format) ToPointer() *AiResponseTransformerPluginLlama2Format {
	return &e
}
func (e *AiResponseTransformerPluginLlama2Format) UnmarshalJSON(data []byte) error {
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
		*e = AiResponseTransformerPluginLlama2Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiResponseTransformerPluginLlama2Format: %v", v)
	}
}

// AiResponseTransformerPluginMistralFormat - If using mistral provider, select the upstream message format.
type AiResponseTransformerPluginMistralFormat string

const (
	AiResponseTransformerPluginMistralFormatOpenai AiResponseTransformerPluginMistralFormat = "openai"
	AiResponseTransformerPluginMistralFormatOllama AiResponseTransformerPluginMistralFormat = "ollama"
)

func (e AiResponseTransformerPluginMistralFormat) ToPointer() *AiResponseTransformerPluginMistralFormat {
	return &e
}
func (e *AiResponseTransformerPluginMistralFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		fallthrough
	case "ollama":
		*e = AiResponseTransformerPluginMistralFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiResponseTransformerPluginMistralFormat: %v", v)
	}
}

// AiResponseTransformerPluginOptions - Key/value settings for the model
type AiResponseTransformerPluginOptions struct {
	// Defines the schema/API version, if using Anthropic provider.
	AnthropicVersion *string `json:"anthropic_version,omitempty"`
	// 'api-version' for Azure OpenAI instances.
	AzureAPIVersion *string `json:"azure_api_version,omitempty"`
	// Deployment ID for Azure OpenAI instances.
	AzureDeploymentID *string `json:"azure_deployment_id,omitempty"`
	// Instance name for Azure OpenAI hosted models.
	AzureInstance *string                             `json:"azure_instance,omitempty"`
	Bedrock       *AiResponseTransformerPluginBedrock `json:"bedrock,omitempty"`
	Gemini        *AiResponseTransformerPluginGemini  `json:"gemini,omitempty"`
	// Defines the cost per 1M tokens in your prompt.
	InputCost *float64 `json:"input_cost,omitempty"`
	// If using llama2 provider, select the upstream message format.
	Llama2Format *AiResponseTransformerPluginLlama2Format `json:"llama2_format,omitempty"`
	// Defines the max_tokens, if using chat or completion models.
	MaxTokens *int64 `json:"max_tokens,omitempty"`
	// If using mistral provider, select the upstream message format.
	MistralFormat *AiResponseTransformerPluginMistralFormat `json:"mistral_format,omitempty"`
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

func (o *AiResponseTransformerPluginOptions) GetAnthropicVersion() *string {
	if o == nil {
		return nil
	}
	return o.AnthropicVersion
}

func (o *AiResponseTransformerPluginOptions) GetAzureAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.AzureAPIVersion
}

func (o *AiResponseTransformerPluginOptions) GetAzureDeploymentID() *string {
	if o == nil {
		return nil
	}
	return o.AzureDeploymentID
}

func (o *AiResponseTransformerPluginOptions) GetAzureInstance() *string {
	if o == nil {
		return nil
	}
	return o.AzureInstance
}

func (o *AiResponseTransformerPluginOptions) GetBedrock() *AiResponseTransformerPluginBedrock {
	if o == nil {
		return nil
	}
	return o.Bedrock
}

func (o *AiResponseTransformerPluginOptions) GetGemini() *AiResponseTransformerPluginGemini {
	if o == nil {
		return nil
	}
	return o.Gemini
}

func (o *AiResponseTransformerPluginOptions) GetInputCost() *float64 {
	if o == nil {
		return nil
	}
	return o.InputCost
}

func (o *AiResponseTransformerPluginOptions) GetLlama2Format() *AiResponseTransformerPluginLlama2Format {
	if o == nil {
		return nil
	}
	return o.Llama2Format
}

func (o *AiResponseTransformerPluginOptions) GetMaxTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTokens
}

func (o *AiResponseTransformerPluginOptions) GetMistralFormat() *AiResponseTransformerPluginMistralFormat {
	if o == nil {
		return nil
	}
	return o.MistralFormat
}

func (o *AiResponseTransformerPluginOptions) GetOutputCost() *float64 {
	if o == nil {
		return nil
	}
	return o.OutputCost
}

func (o *AiResponseTransformerPluginOptions) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *AiResponseTransformerPluginOptions) GetTopK() *int64 {
	if o == nil {
		return nil
	}
	return o.TopK
}

func (o *AiResponseTransformerPluginOptions) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}

func (o *AiResponseTransformerPluginOptions) GetUpstreamPath() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamPath
}

func (o *AiResponseTransformerPluginOptions) GetUpstreamURL() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamURL
}

// AiResponseTransformerPluginProvider - AI provider request format - Kong translates requests to and from the specified backend compatible formats.
type AiResponseTransformerPluginProvider string

const (
	AiResponseTransformerPluginProviderOpenai    AiResponseTransformerPluginProvider = "openai"
	AiResponseTransformerPluginProviderAzure     AiResponseTransformerPluginProvider = "azure"
	AiResponseTransformerPluginProviderAnthropic AiResponseTransformerPluginProvider = "anthropic"
	AiResponseTransformerPluginProviderCohere    AiResponseTransformerPluginProvider = "cohere"
	AiResponseTransformerPluginProviderMistral   AiResponseTransformerPluginProvider = "mistral"
	AiResponseTransformerPluginProviderLlama2    AiResponseTransformerPluginProvider = "llama2"
	AiResponseTransformerPluginProviderGemini    AiResponseTransformerPluginProvider = "gemini"
	AiResponseTransformerPluginProviderBedrock   AiResponseTransformerPluginProvider = "bedrock"
)

func (e AiResponseTransformerPluginProvider) ToPointer() *AiResponseTransformerPluginProvider {
	return &e
}
func (e *AiResponseTransformerPluginProvider) UnmarshalJSON(data []byte) error {
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
		*e = AiResponseTransformerPluginProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiResponseTransformerPluginProvider: %v", v)
	}
}

type AiResponseTransformerPluginModel struct {
	// Model name to execute.
	Name *string `json:"name,omitempty"`
	// Key/value settings for the model
	Options *AiResponseTransformerPluginOptions `json:"options,omitempty"`
	// AI provider request format - Kong translates requests to and from the specified backend compatible formats.
	Provider *AiResponseTransformerPluginProvider `json:"provider,omitempty"`
}

func (o *AiResponseTransformerPluginModel) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AiResponseTransformerPluginModel) GetOptions() *AiResponseTransformerPluginOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *AiResponseTransformerPluginModel) GetProvider() *AiResponseTransformerPluginProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

// AiResponseTransformerPluginRouteType - The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
type AiResponseTransformerPluginRouteType string

const (
	AiResponseTransformerPluginRouteTypeLlmV1Chat        AiResponseTransformerPluginRouteType = "llm/v1/chat"
	AiResponseTransformerPluginRouteTypeLlmV1Completions AiResponseTransformerPluginRouteType = "llm/v1/completions"
	AiResponseTransformerPluginRouteTypePreserve         AiResponseTransformerPluginRouteType = "preserve"
)

func (e AiResponseTransformerPluginRouteType) ToPointer() *AiResponseTransformerPluginRouteType {
	return &e
}
func (e *AiResponseTransformerPluginRouteType) UnmarshalJSON(data []byte) error {
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
		*e = AiResponseTransformerPluginRouteType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiResponseTransformerPluginRouteType: %v", v)
	}
}

type AiResponseTransformerPluginLlm struct {
	Auth    *AiResponseTransformerPluginAuth    `json:"auth,omitempty"`
	Logging *AiResponseTransformerPluginLogging `json:"logging,omitempty"`
	Model   *AiResponseTransformerPluginModel   `json:"model,omitempty"`
	// The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation.
	RouteType *AiResponseTransformerPluginRouteType `json:"route_type,omitempty"`
}

func (o *AiResponseTransformerPluginLlm) GetAuth() *AiResponseTransformerPluginAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *AiResponseTransformerPluginLlm) GetLogging() *AiResponseTransformerPluginLogging {
	if o == nil {
		return nil
	}
	return o.Logging
}

func (o *AiResponseTransformerPluginLlm) GetModel() *AiResponseTransformerPluginModel {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *AiResponseTransformerPluginLlm) GetRouteType() *AiResponseTransformerPluginRouteType {
	if o == nil {
		return nil
	}
	return o.RouteType
}

type AiResponseTransformerPluginConfig struct {
	// A string representing a host name, such as example.com.
	HTTPProxyHost *string `json:"http_proxy_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	HTTPProxyPort *int64 `json:"http_proxy_port,omitempty"`
	// Timeout in milliseconds for the AI upstream service.
	HTTPTimeout *int64 `json:"http_timeout,omitempty"`
	// A string representing a host name, such as example.com.
	HTTPSProxyHost *string `json:"https_proxy_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	HTTPSProxyPort *int64 `json:"https_proxy_port,omitempty"`
	// Verify the TLS certificate of the AI upstream service.
	HTTPSVerify *bool                           `json:"https_verify,omitempty"`
	Llm         *AiResponseTransformerPluginLlm `json:"llm,omitempty"`
	// max allowed body size allowed to be introspected
	MaxRequestBodySize *int64 `json:"max_request_body_size,omitempty"`
	// Set true to read specific response format from the LLM, and accordingly set the status code / body / headers that proxy back to the client. You need to engineer your LLM prompt to return the correct format, see plugin docs 'Overview' page for usage instructions.
	ParseLlmResponseJSONInstructions *bool `json:"parse_llm_response_json_instructions,omitempty"`
	// Use this prompt to tune the LLM system/assistant message for the returning proxy response (from the upstream), adn what response format you are expecting.
	Prompt *string `json:"prompt,omitempty"`
	// Defines the regular expression that must match to indicate a successful AI transformation at the response phase. The first match will be set as the returning body. If the AI service's response doesn't match this pattern, a failure is returned to the client.
	TransformationExtractPattern *string `json:"transformation_extract_pattern,omitempty"`
}

func (o *AiResponseTransformerPluginConfig) GetHTTPProxyHost() *string {
	if o == nil {
		return nil
	}
	return o.HTTPProxyHost
}

func (o *AiResponseTransformerPluginConfig) GetHTTPProxyPort() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPProxyPort
}

func (o *AiResponseTransformerPluginConfig) GetHTTPTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPTimeout
}

func (o *AiResponseTransformerPluginConfig) GetHTTPSProxyHost() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSProxyHost
}

func (o *AiResponseTransformerPluginConfig) GetHTTPSProxyPort() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPSProxyPort
}

func (o *AiResponseTransformerPluginConfig) GetHTTPSVerify() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPSVerify
}

func (o *AiResponseTransformerPluginConfig) GetLlm() *AiResponseTransformerPluginLlm {
	if o == nil {
		return nil
	}
	return o.Llm
}

func (o *AiResponseTransformerPluginConfig) GetMaxRequestBodySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestBodySize
}

func (o *AiResponseTransformerPluginConfig) GetParseLlmResponseJSONInstructions() *bool {
	if o == nil {
		return nil
	}
	return o.ParseLlmResponseJSONInstructions
}

func (o *AiResponseTransformerPluginConfig) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *AiResponseTransformerPluginConfig) GetTransformationExtractPattern() *string {
	if o == nil {
		return nil
	}
	return o.TransformationExtractPattern
}

type AiResponseTransformerPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiResponseTransformerPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiResponseTransformerPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiResponseTransformerPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiResponseTransformerPluginOrdering struct {
	After  *AiResponseTransformerPluginAfter  `json:"after,omitempty"`
	Before *AiResponseTransformerPluginBefore `json:"before,omitempty"`
}

func (o *AiResponseTransformerPluginOrdering) GetAfter() *AiResponseTransformerPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *AiResponseTransformerPluginOrdering) GetBefore() *AiResponseTransformerPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type AiResponseTransformerPluginProtocols string

const (
	AiResponseTransformerPluginProtocolsGrpc           AiResponseTransformerPluginProtocols = "grpc"
	AiResponseTransformerPluginProtocolsGrpcs          AiResponseTransformerPluginProtocols = "grpcs"
	AiResponseTransformerPluginProtocolsHTTP           AiResponseTransformerPluginProtocols = "http"
	AiResponseTransformerPluginProtocolsHTTPS          AiResponseTransformerPluginProtocols = "https"
	AiResponseTransformerPluginProtocolsTCP            AiResponseTransformerPluginProtocols = "tcp"
	AiResponseTransformerPluginProtocolsTLS            AiResponseTransformerPluginProtocols = "tls"
	AiResponseTransformerPluginProtocolsTLSPassthrough AiResponseTransformerPluginProtocols = "tls_passthrough"
	AiResponseTransformerPluginProtocolsUDP            AiResponseTransformerPluginProtocols = "udp"
	AiResponseTransformerPluginProtocolsWs             AiResponseTransformerPluginProtocols = "ws"
	AiResponseTransformerPluginProtocolsWss            AiResponseTransformerPluginProtocols = "wss"
)

func (e AiResponseTransformerPluginProtocols) ToPointer() *AiResponseTransformerPluginProtocols {
	return &e
}
func (e *AiResponseTransformerPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AiResponseTransformerPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiResponseTransformerPluginProtocols: %v", v)
	}
}

// AiResponseTransformerPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AiResponseTransformerPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiResponseTransformerPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiResponseTransformerPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiResponseTransformerPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiResponseTransformerPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AiResponseTransformerPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiResponseTransformerPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiResponseTransformerPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AiResponseTransformerPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiResponseTransformerPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiResponseTransformerPlugin struct {
	Config *AiResponseTransformerPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                                `json:"enabled,omitempty"`
	ID           *string                              `json:"id,omitempty"`
	InstanceName *string                              `json:"instance_name,omitempty"`
	name         *string                              `const:"ai-response-transformer" json:"name,omitempty"`
	Ordering     *AiResponseTransformerPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AiResponseTransformerPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *AiResponseTransformerPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *AiResponseTransformerPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AiResponseTransformerPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiResponseTransformerPluginService `json:"service,omitempty"`
}

func (a AiResponseTransformerPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiResponseTransformerPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiResponseTransformerPlugin) GetConfig() *AiResponseTransformerPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *AiResponseTransformerPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AiResponseTransformerPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiResponseTransformerPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiResponseTransformerPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiResponseTransformerPlugin) GetName() *string {
	return types.String("ai-response-transformer")
}

func (o *AiResponseTransformerPlugin) GetOrdering() *AiResponseTransformerPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiResponseTransformerPlugin) GetProtocols() []AiResponseTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiResponseTransformerPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiResponseTransformerPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AiResponseTransformerPlugin) GetConsumer() *AiResponseTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiResponseTransformerPlugin) GetConsumerGroup() *AiResponseTransformerPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiResponseTransformerPlugin) GetRoute() *AiResponseTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiResponseTransformerPlugin) GetService() *AiResponseTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
