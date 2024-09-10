// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateAWSLambdaPluginAWSImdsProtocolVersion - Identifier to select the IMDS protocol version to use: `v1` or `v2`.
type CreateAWSLambdaPluginAWSImdsProtocolVersion string

const (
	CreateAWSLambdaPluginAWSImdsProtocolVersionV1 CreateAWSLambdaPluginAWSImdsProtocolVersion = "v1"
	CreateAWSLambdaPluginAWSImdsProtocolVersionV2 CreateAWSLambdaPluginAWSImdsProtocolVersion = "v2"
)

func (e CreateAWSLambdaPluginAWSImdsProtocolVersion) ToPointer() *CreateAWSLambdaPluginAWSImdsProtocolVersion {
	return &e
}
func (e *CreateAWSLambdaPluginAWSImdsProtocolVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v1":
		fallthrough
	case "v2":
		*e = CreateAWSLambdaPluginAWSImdsProtocolVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAWSLambdaPluginAWSImdsProtocolVersion: %v", v)
	}
}

// CreateAWSLambdaPluginInvocationType - The InvocationType to use when invoking the function. Available types are RequestResponse, Event, DryRun.
type CreateAWSLambdaPluginInvocationType string

const (
	CreateAWSLambdaPluginInvocationTypeRequestResponse CreateAWSLambdaPluginInvocationType = "RequestResponse"
	CreateAWSLambdaPluginInvocationTypeEvent           CreateAWSLambdaPluginInvocationType = "Event"
	CreateAWSLambdaPluginInvocationTypeDryRun          CreateAWSLambdaPluginInvocationType = "DryRun"
)

func (e CreateAWSLambdaPluginInvocationType) ToPointer() *CreateAWSLambdaPluginInvocationType {
	return &e
}
func (e *CreateAWSLambdaPluginInvocationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RequestResponse":
		fallthrough
	case "Event":
		fallthrough
	case "DryRun":
		*e = CreateAWSLambdaPluginInvocationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAWSLambdaPluginInvocationType: %v", v)
	}
}

// CreateAWSLambdaPluginLogType - The LogType to use when invoking the function. By default, None and Tail are supported.
type CreateAWSLambdaPluginLogType string

const (
	CreateAWSLambdaPluginLogTypeTail CreateAWSLambdaPluginLogType = "Tail"
	CreateAWSLambdaPluginLogTypeNone CreateAWSLambdaPluginLogType = "None"
)

func (e CreateAWSLambdaPluginLogType) ToPointer() *CreateAWSLambdaPluginLogType {
	return &e
}
func (e *CreateAWSLambdaPluginLogType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Tail":
		fallthrough
	case "None":
		*e = CreateAWSLambdaPluginLogType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAWSLambdaPluginLogType: %v", v)
	}
}

type CreateAWSLambdaPluginConfig struct {
	// The target AWS IAM role ARN used to invoke the Lambda function.
	AwsAssumeRoleArn *string `json:"aws_assume_role_arn,omitempty"`
	// Identifier to select the IMDS protocol version to use: `v1` or `v2`.
	AwsImdsProtocolVersion *CreateAWSLambdaPluginAWSImdsProtocolVersion `json:"aws_imds_protocol_version,omitempty"`
	// The AWS key credential to be used when invoking the function.
	AwsKey *string `json:"aws_key,omitempty"`
	// A string representing a host name, such as example.com.
	AwsRegion *string `json:"aws_region,omitempty"`
	// The identifier of the assumed role session.
	AwsRoleSessionName *string `json:"aws_role_session_name,omitempty"`
	// The AWS secret credential to be used when invoking the function.
	AwsSecret *string `json:"aws_secret,omitempty"`
	// An optional value that defines whether the plugin should wrap requests into the Amazon API gateway.
	AwsgatewayCompatible *bool `json:"awsgateway_compatible,omitempty"`
	// An optional value that Base64-encodes the request body.
	Base64EncodeBody *bool `json:"base64_encode_body,omitempty"`
	DisableHTTPS     *bool `json:"disable_https,omitempty"`
	// An optional value that defines whether the request body is sent in the request_body field of the JSON-encoded request. If the body arguments can be parsed, they are sent in the separate request_body_args field of the request.
	ForwardRequestBody *bool `json:"forward_request_body,omitempty"`
	// An optional value that defines whether the original HTTP request headers are sent as a map in the request_headers field of the JSON-encoded request.
	ForwardRequestHeaders *bool `json:"forward_request_headers,omitempty"`
	// An optional value that defines whether the original HTTP request method verb is sent in the request_method field of the JSON-encoded request.
	ForwardRequestMethod *bool `json:"forward_request_method,omitempty"`
	// An optional value that defines whether the original HTTP request URI is sent in the request_uri field of the JSON-encoded request.
	ForwardRequestURI *bool `json:"forward_request_uri,omitempty"`
	// The AWS Lambda function to invoke. Both function name and function ARN (including partial) are supported.
	FunctionName *string `json:"function_name,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// The InvocationType to use when invoking the function. Available types are RequestResponse, Event, DryRun.
	InvocationType *CreateAWSLambdaPluginInvocationType `json:"invocation_type,omitempty"`
	// An optional value that defines whether the response format to receive from the Lambda to this format.
	IsProxyIntegration *bool `json:"is_proxy_integration,omitempty"`
	// An optional value in milliseconds that defines how long an idle connection lives before being closed.
	Keepalive *float64 `json:"keepalive,omitempty"`
	// The LogType to use when invoking the function. By default, None and Tail are supported.
	LogType *CreateAWSLambdaPluginLogType `json:"log_type,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	ProxyURL *string `json:"proxy_url,omitempty"`
	// The qualifier to use when invoking the function.
	Qualifier *string `json:"qualifier,omitempty"`
	// An optional value that defines whether Kong should send large bodies that are buffered to disk
	SkipLargeBodies *bool `json:"skip_large_bodies,omitempty"`
	// An optional timeout in milliseconds when invoking the function.
	Timeout *float64 `json:"timeout,omitempty"`
	// The response status code to use (instead of the default 200, 202, or 204) in the case of an Unhandled Function Error.
	UnhandledStatus *int64 `json:"unhandled_status,omitempty"`
}

func (o *CreateAWSLambdaPluginConfig) GetAwsAssumeRoleArn() *string {
	if o == nil {
		return nil
	}
	return o.AwsAssumeRoleArn
}

func (o *CreateAWSLambdaPluginConfig) GetAwsImdsProtocolVersion() *CreateAWSLambdaPluginAWSImdsProtocolVersion {
	if o == nil {
		return nil
	}
	return o.AwsImdsProtocolVersion
}

func (o *CreateAWSLambdaPluginConfig) GetAwsKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsKey
}

func (o *CreateAWSLambdaPluginConfig) GetAwsRegion() *string {
	if o == nil {
		return nil
	}
	return o.AwsRegion
}

func (o *CreateAWSLambdaPluginConfig) GetAwsRoleSessionName() *string {
	if o == nil {
		return nil
	}
	return o.AwsRoleSessionName
}

func (o *CreateAWSLambdaPluginConfig) GetAwsSecret() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecret
}

func (o *CreateAWSLambdaPluginConfig) GetAwsgatewayCompatible() *bool {
	if o == nil {
		return nil
	}
	return o.AwsgatewayCompatible
}

func (o *CreateAWSLambdaPluginConfig) GetBase64EncodeBody() *bool {
	if o == nil {
		return nil
	}
	return o.Base64EncodeBody
}

func (o *CreateAWSLambdaPluginConfig) GetDisableHTTPS() *bool {
	if o == nil {
		return nil
	}
	return o.DisableHTTPS
}

func (o *CreateAWSLambdaPluginConfig) GetForwardRequestBody() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardRequestBody
}

func (o *CreateAWSLambdaPluginConfig) GetForwardRequestHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardRequestHeaders
}

func (o *CreateAWSLambdaPluginConfig) GetForwardRequestMethod() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardRequestMethod
}

func (o *CreateAWSLambdaPluginConfig) GetForwardRequestURI() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardRequestURI
}

func (o *CreateAWSLambdaPluginConfig) GetFunctionName() *string {
	if o == nil {
		return nil
	}
	return o.FunctionName
}

func (o *CreateAWSLambdaPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateAWSLambdaPluginConfig) GetInvocationType() *CreateAWSLambdaPluginInvocationType {
	if o == nil {
		return nil
	}
	return o.InvocationType
}

func (o *CreateAWSLambdaPluginConfig) GetIsProxyIntegration() *bool {
	if o == nil {
		return nil
	}
	return o.IsProxyIntegration
}

func (o *CreateAWSLambdaPluginConfig) GetKeepalive() *float64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *CreateAWSLambdaPluginConfig) GetLogType() *CreateAWSLambdaPluginLogType {
	if o == nil {
		return nil
	}
	return o.LogType
}

func (o *CreateAWSLambdaPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateAWSLambdaPluginConfig) GetProxyURL() *string {
	if o == nil {
		return nil
	}
	return o.ProxyURL
}

func (o *CreateAWSLambdaPluginConfig) GetQualifier() *string {
	if o == nil {
		return nil
	}
	return o.Qualifier
}

func (o *CreateAWSLambdaPluginConfig) GetSkipLargeBodies() *bool {
	if o == nil {
		return nil
	}
	return o.SkipLargeBodies
}

func (o *CreateAWSLambdaPluginConfig) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *CreateAWSLambdaPluginConfig) GetUnhandledStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.UnhandledStatus
}

type CreateAWSLambdaPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateAWSLambdaPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateAWSLambdaPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateAWSLambdaPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateAWSLambdaPluginOrdering struct {
	After  *CreateAWSLambdaPluginAfter  `json:"after,omitempty"`
	Before *CreateAWSLambdaPluginBefore `json:"before,omitempty"`
}

func (o *CreateAWSLambdaPluginOrdering) GetAfter() *CreateAWSLambdaPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateAWSLambdaPluginOrdering) GetBefore() *CreateAWSLambdaPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateAWSLambdaPluginProtocols string

const (
	CreateAWSLambdaPluginProtocolsGrpc           CreateAWSLambdaPluginProtocols = "grpc"
	CreateAWSLambdaPluginProtocolsGrpcs          CreateAWSLambdaPluginProtocols = "grpcs"
	CreateAWSLambdaPluginProtocolsHTTP           CreateAWSLambdaPluginProtocols = "http"
	CreateAWSLambdaPluginProtocolsHTTPS          CreateAWSLambdaPluginProtocols = "https"
	CreateAWSLambdaPluginProtocolsTCP            CreateAWSLambdaPluginProtocols = "tcp"
	CreateAWSLambdaPluginProtocolsTLS            CreateAWSLambdaPluginProtocols = "tls"
	CreateAWSLambdaPluginProtocolsTLSPassthrough CreateAWSLambdaPluginProtocols = "tls_passthrough"
	CreateAWSLambdaPluginProtocolsUDP            CreateAWSLambdaPluginProtocols = "udp"
	CreateAWSLambdaPluginProtocolsWs             CreateAWSLambdaPluginProtocols = "ws"
	CreateAWSLambdaPluginProtocolsWss            CreateAWSLambdaPluginProtocols = "wss"
)

func (e CreateAWSLambdaPluginProtocols) ToPointer() *CreateAWSLambdaPluginProtocols {
	return &e
}
func (e *CreateAWSLambdaPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateAWSLambdaPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAWSLambdaPluginProtocols: %v", v)
	}
}

// CreateAWSLambdaPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateAWSLambdaPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAWSLambdaPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAWSLambdaPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAWSLambdaPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAWSLambdaPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateAWSLambdaPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAWSLambdaPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAWSLambdaPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateAWSLambdaPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAWSLambdaPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAWSLambdaPlugin struct {
	Config *CreateAWSLambdaPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                          `json:"enabled,omitempty"`
	InstanceName *string                        `json:"instance_name,omitempty"`
	name         *string                        `const:"aws-lambda" json:"name,omitempty"`
	Ordering     *CreateAWSLambdaPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateAWSLambdaPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateAWSLambdaPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateAWSLambdaPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateAWSLambdaPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateAWSLambdaPluginService `json:"service,omitempty"`
}

func (c CreateAWSLambdaPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAWSLambdaPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAWSLambdaPlugin) GetConfig() *CreateAWSLambdaPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateAWSLambdaPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateAWSLambdaPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateAWSLambdaPlugin) GetName() *string {
	return types.String("aws-lambda")
}

func (o *CreateAWSLambdaPlugin) GetOrdering() *CreateAWSLambdaPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateAWSLambdaPlugin) GetProtocols() []CreateAWSLambdaPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateAWSLambdaPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateAWSLambdaPlugin) GetConsumer() *CreateAWSLambdaPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateAWSLambdaPlugin) GetConsumerGroup() *CreateAWSLambdaPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateAWSLambdaPlugin) GetRoute() *CreateAWSLambdaPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateAWSLambdaPlugin) GetService() *CreateAWSLambdaPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
