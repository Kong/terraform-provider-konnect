// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type AiSemanticCachePluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiSemanticCachePluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiSemanticCachePluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *AiSemanticCachePluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type AiSemanticCachePluginOrdering struct {
	After  *AiSemanticCachePluginAfter  `json:"after,omitempty"`
	Before *AiSemanticCachePluginBefore `json:"before,omitempty"`
}

func (o *AiSemanticCachePluginOrdering) GetAfter() *AiSemanticCachePluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *AiSemanticCachePluginOrdering) GetBefore() *AiSemanticCachePluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

// AiSemanticCachePluginParamLocation - Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.
type AiSemanticCachePluginParamLocation string

const (
	AiSemanticCachePluginParamLocationBody  AiSemanticCachePluginParamLocation = "body"
	AiSemanticCachePluginParamLocationQuery AiSemanticCachePluginParamLocation = "query"
)

func (e AiSemanticCachePluginParamLocation) ToPointer() *AiSemanticCachePluginParamLocation {
	return &e
}
func (e *AiSemanticCachePluginParamLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "body":
		fallthrough
	case "query":
		*e = AiSemanticCachePluginParamLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiSemanticCachePluginParamLocation: %v", v)
	}
}

type AiSemanticCachePluginAuth struct {
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
	ParamLocation *AiSemanticCachePluginParamLocation `json:"param_location,omitempty"`
	// If AI model requires authentication via query parameter, specify its name here.
	ParamName *string `json:"param_name,omitempty"`
	// Specify the full parameter value for 'param_name'.
	ParamValue *string `json:"param_value,omitempty"`
}

func (o *AiSemanticCachePluginAuth) GetAllowOverride() *bool {
	if o == nil {
		return nil
	}
	return o.AllowOverride
}

func (o *AiSemanticCachePluginAuth) GetAwsAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccessKeyID
}

func (o *AiSemanticCachePluginAuth) GetAwsSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretAccessKey
}

func (o *AiSemanticCachePluginAuth) GetAzureClientID() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientID
}

func (o *AiSemanticCachePluginAuth) GetAzureClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.AzureClientSecret
}

func (o *AiSemanticCachePluginAuth) GetAzureTenantID() *string {
	if o == nil {
		return nil
	}
	return o.AzureTenantID
}

func (o *AiSemanticCachePluginAuth) GetAzureUseManagedIdentity() *bool {
	if o == nil {
		return nil
	}
	return o.AzureUseManagedIdentity
}

func (o *AiSemanticCachePluginAuth) GetGcpServiceAccountJSON() *string {
	if o == nil {
		return nil
	}
	return o.GcpServiceAccountJSON
}

func (o *AiSemanticCachePluginAuth) GetGcpUseServiceAccount() *bool {
	if o == nil {
		return nil
	}
	return o.GcpUseServiceAccount
}

func (o *AiSemanticCachePluginAuth) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *AiSemanticCachePluginAuth) GetHeaderValue() *string {
	if o == nil {
		return nil
	}
	return o.HeaderValue
}

func (o *AiSemanticCachePluginAuth) GetParamLocation() *AiSemanticCachePluginParamLocation {
	if o == nil {
		return nil
	}
	return o.ParamLocation
}

func (o *AiSemanticCachePluginAuth) GetParamName() *string {
	if o == nil {
		return nil
	}
	return o.ParamName
}

func (o *AiSemanticCachePluginAuth) GetParamValue() *string {
	if o == nil {
		return nil
	}
	return o.ParamValue
}

// AiSemanticCachePluginOptions - Key/value settings for the model
type AiSemanticCachePluginOptions struct {
	// upstream url for the embeddings
	UpstreamURL *string `json:"upstream_url,omitempty"`
}

func (o *AiSemanticCachePluginOptions) GetUpstreamURL() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamURL
}

// AiSemanticCachePluginProvider - AI provider format to use for embeddings API
type AiSemanticCachePluginProvider string

const (
	AiSemanticCachePluginProviderMistral AiSemanticCachePluginProvider = "mistral"
	AiSemanticCachePluginProviderOpenai  AiSemanticCachePluginProvider = "openai"
)

func (e AiSemanticCachePluginProvider) ToPointer() *AiSemanticCachePluginProvider {
	return &e
}
func (e *AiSemanticCachePluginProvider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mistral":
		fallthrough
	case "openai":
		*e = AiSemanticCachePluginProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiSemanticCachePluginProvider: %v", v)
	}
}

type AiSemanticCachePluginModel struct {
	// Model name to execute.
	Name *string `json:"name,omitempty"`
	// Key/value settings for the model
	Options *AiSemanticCachePluginOptions `json:"options,omitempty"`
	// AI provider format to use for embeddings API
	Provider *AiSemanticCachePluginProvider `json:"provider,omitempty"`
}

func (o *AiSemanticCachePluginModel) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AiSemanticCachePluginModel) GetOptions() *AiSemanticCachePluginOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *AiSemanticCachePluginModel) GetProvider() *AiSemanticCachePluginProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

type AiSemanticCachePluginEmbeddings struct {
	Auth  *AiSemanticCachePluginAuth  `json:"auth,omitempty"`
	Model *AiSemanticCachePluginModel `json:"model,omitempty"`
}

func (o *AiSemanticCachePluginEmbeddings) GetAuth() *AiSemanticCachePluginAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *AiSemanticCachePluginEmbeddings) GetModel() *AiSemanticCachePluginModel {
	if o == nil {
		return nil
	}
	return o.Model
}

// AiSemanticCachePluginDistanceMetric - the distance metric to use for vector searches
type AiSemanticCachePluginDistanceMetric string

const (
	AiSemanticCachePluginDistanceMetricCosine    AiSemanticCachePluginDistanceMetric = "cosine"
	AiSemanticCachePluginDistanceMetricEuclidean AiSemanticCachePluginDistanceMetric = "euclidean"
)

func (e AiSemanticCachePluginDistanceMetric) ToPointer() *AiSemanticCachePluginDistanceMetric {
	return &e
}
func (e *AiSemanticCachePluginDistanceMetric) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cosine":
		fallthrough
	case "euclidean":
		*e = AiSemanticCachePluginDistanceMetric(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiSemanticCachePluginDistanceMetric: %v", v)
	}
}

type AiSemanticCachePluginClusterNodes struct {
	// A string representing a host name, such as example.com.
	IP *string `json:"ip,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *AiSemanticCachePluginClusterNodes) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *AiSemanticCachePluginClusterNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type AiSemanticCachePluginSentinelNodes struct {
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *AiSemanticCachePluginSentinelNodes) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *AiSemanticCachePluginSentinelNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

// AiSemanticCachePluginSentinelRole - Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
type AiSemanticCachePluginSentinelRole string

const (
	AiSemanticCachePluginSentinelRoleAny    AiSemanticCachePluginSentinelRole = "any"
	AiSemanticCachePluginSentinelRoleMaster AiSemanticCachePluginSentinelRole = "master"
	AiSemanticCachePluginSentinelRoleSlave  AiSemanticCachePluginSentinelRole = "slave"
)

func (e AiSemanticCachePluginSentinelRole) ToPointer() *AiSemanticCachePluginSentinelRole {
	return &e
}
func (e *AiSemanticCachePluginSentinelRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "master":
		fallthrough
	case "slave":
		*e = AiSemanticCachePluginSentinelRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiSemanticCachePluginSentinelRole: %v", v)
	}
}

type AiSemanticCachePluginRedis struct {
	// Maximum retry attempts for redirection.
	ClusterMaxRedirections *int64 `json:"cluster_max_redirections,omitempty"`
	// Cluster addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.
	ClusterNodes []AiSemanticCachePluginClusterNodes `json:"cluster_nodes,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// If the connection to Redis is proxied (e.g. Envoy), set it `true`. Set the `host` and `port` to point to the proxy address.
	ConnectionIsProxied *bool `json:"connection_is_proxied,omitempty"`
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return `nil`. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than `keepalive_pool_size`. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than `keepalive_pool_size`.
	KeepaliveBacklog *int64 `json:"keepalive_backlog,omitempty"`
	// The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither `keepalive_pool_size` nor `keepalive_backlog` is specified, no pool is created. If `keepalive_pool_size` isn't specified but `keepalive_backlog` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.
	KeepalivePoolSize *int64 `json:"keepalive_pool_size,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.
	SentinelMaster *string `json:"sentinel_master,omitempty"`
	// Sentinel node addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.
	SentinelNodes []AiSemanticCachePluginSentinelNodes `json:"sentinel_nodes,omitempty"`
	// Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.
	SentinelPassword *string `json:"sentinel_password,omitempty"`
	// Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
	SentinelRole *AiSemanticCachePluginSentinelRole `json:"sentinel_role,omitempty"`
	// Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.
	SentinelUsername *string `json:"sentinel_username,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *AiSemanticCachePluginRedis) GetClusterMaxRedirections() *int64 {
	if o == nil {
		return nil
	}
	return o.ClusterMaxRedirections
}

func (o *AiSemanticCachePluginRedis) GetClusterNodes() []AiSemanticCachePluginClusterNodes {
	if o == nil {
		return nil
	}
	return o.ClusterNodes
}

func (o *AiSemanticCachePluginRedis) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *AiSemanticCachePluginRedis) GetConnectionIsProxied() *bool {
	if o == nil {
		return nil
	}
	return o.ConnectionIsProxied
}

func (o *AiSemanticCachePluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *AiSemanticCachePluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *AiSemanticCachePluginRedis) GetKeepaliveBacklog() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepaliveBacklog
}

func (o *AiSemanticCachePluginRedis) GetKeepalivePoolSize() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepalivePoolSize
}

func (o *AiSemanticCachePluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *AiSemanticCachePluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *AiSemanticCachePluginRedis) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *AiSemanticCachePluginRedis) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

func (o *AiSemanticCachePluginRedis) GetSentinelMaster() *string {
	if o == nil {
		return nil
	}
	return o.SentinelMaster
}

func (o *AiSemanticCachePluginRedis) GetSentinelNodes() []AiSemanticCachePluginSentinelNodes {
	if o == nil {
		return nil
	}
	return o.SentinelNodes
}

func (o *AiSemanticCachePluginRedis) GetSentinelPassword() *string {
	if o == nil {
		return nil
	}
	return o.SentinelPassword
}

func (o *AiSemanticCachePluginRedis) GetSentinelRole() *AiSemanticCachePluginSentinelRole {
	if o == nil {
		return nil
	}
	return o.SentinelRole
}

func (o *AiSemanticCachePluginRedis) GetSentinelUsername() *string {
	if o == nil {
		return nil
	}
	return o.SentinelUsername
}

func (o *AiSemanticCachePluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *AiSemanticCachePluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *AiSemanticCachePluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *AiSemanticCachePluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// AiSemanticCachePluginStrategy - which vector database driver to use
type AiSemanticCachePluginStrategy string

const (
	AiSemanticCachePluginStrategyRedis AiSemanticCachePluginStrategy = "redis"
)

func (e AiSemanticCachePluginStrategy) ToPointer() *AiSemanticCachePluginStrategy {
	return &e
}
func (e *AiSemanticCachePluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redis":
		*e = AiSemanticCachePluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiSemanticCachePluginStrategy: %v", v)
	}
}

type AiSemanticCachePluginVectordb struct {
	// the desired dimensionality for the vectors
	Dimensions *int64 `json:"dimensions,omitempty"`
	// the distance metric to use for vector searches
	DistanceMetric *AiSemanticCachePluginDistanceMetric `json:"distance_metric,omitempty"`
	Redis          *AiSemanticCachePluginRedis          `json:"redis,omitempty"`
	// which vector database driver to use
	Strategy *AiSemanticCachePluginStrategy `json:"strategy,omitempty"`
	// the default similarity threshold for accepting semantic search results (float)
	Threshold *float64 `json:"threshold,omitempty"`
}

func (o *AiSemanticCachePluginVectordb) GetDimensions() *int64 {
	if o == nil {
		return nil
	}
	return o.Dimensions
}

func (o *AiSemanticCachePluginVectordb) GetDistanceMetric() *AiSemanticCachePluginDistanceMetric {
	if o == nil {
		return nil
	}
	return o.DistanceMetric
}

func (o *AiSemanticCachePluginVectordb) GetRedis() *AiSemanticCachePluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *AiSemanticCachePluginVectordb) GetStrategy() *AiSemanticCachePluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *AiSemanticCachePluginVectordb) GetThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.Threshold
}

type AiSemanticCachePluginConfig struct {
	// When enabled, respect the Cache-Control behaviors defined in RFC7234.
	CacheControl *bool `json:"cache_control,omitempty"`
	// TTL in seconds of cache entities. Must be a value greater than 0.
	CacheTTL   *int64                           `json:"cache_ttl,omitempty"`
	Embeddings *AiSemanticCachePluginEmbeddings `json:"embeddings,omitempty"`
	// When enabled, a first check for exact query will be done. It will impact DB size
	ExactCaching *bool `json:"exact_caching,omitempty"`
	// Ignore and discard any assistant prompts when Vectorizing the request
	IgnoreAssistantPrompts *bool `json:"ignore_assistant_prompts,omitempty"`
	// Ignore and discard any system prompts when Vectorizing the request
	IgnoreSystemPrompts *bool `json:"ignore_system_prompts,omitempty"`
	// Ignore and discard any tool prompts when Vectorizing the request
	IgnoreToolPrompts *bool `json:"ignore_tool_prompts,omitempty"`
	// Number of messages in the chat history to Vectorize/Cache
	MessageCountback *float64 `json:"message_countback,omitempty"`
	// Halt the LLM request process in case of a caching system failure
	StopOnFailure *bool                          `json:"stop_on_failure,omitempty"`
	Vectordb      *AiSemanticCachePluginVectordb `json:"vectordb,omitempty"`
}

func (o *AiSemanticCachePluginConfig) GetCacheControl() *bool {
	if o == nil {
		return nil
	}
	return o.CacheControl
}

func (o *AiSemanticCachePluginConfig) GetCacheTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *AiSemanticCachePluginConfig) GetEmbeddings() *AiSemanticCachePluginEmbeddings {
	if o == nil {
		return nil
	}
	return o.Embeddings
}

func (o *AiSemanticCachePluginConfig) GetExactCaching() *bool {
	if o == nil {
		return nil
	}
	return o.ExactCaching
}

func (o *AiSemanticCachePluginConfig) GetIgnoreAssistantPrompts() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreAssistantPrompts
}

func (o *AiSemanticCachePluginConfig) GetIgnoreSystemPrompts() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreSystemPrompts
}

func (o *AiSemanticCachePluginConfig) GetIgnoreToolPrompts() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreToolPrompts
}

func (o *AiSemanticCachePluginConfig) GetMessageCountback() *float64 {
	if o == nil {
		return nil
	}
	return o.MessageCountback
}

func (o *AiSemanticCachePluginConfig) GetStopOnFailure() *bool {
	if o == nil {
		return nil
	}
	return o.StopOnFailure
}

func (o *AiSemanticCachePluginConfig) GetVectordb() *AiSemanticCachePluginVectordb {
	if o == nil {
		return nil
	}
	return o.Vectordb
}

// AiSemanticCachePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AiSemanticCachePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiSemanticCachePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiSemanticCachePluginConsumerGroup - If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
type AiSemanticCachePluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiSemanticCachePluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiSemanticCachePluginProtocols string

const (
	AiSemanticCachePluginProtocolsGrpc  AiSemanticCachePluginProtocols = "grpc"
	AiSemanticCachePluginProtocolsGrpcs AiSemanticCachePluginProtocols = "grpcs"
	AiSemanticCachePluginProtocolsHTTP  AiSemanticCachePluginProtocols = "http"
	AiSemanticCachePluginProtocolsHTTPS AiSemanticCachePluginProtocols = "https"
)

func (e AiSemanticCachePluginProtocols) ToPointer() *AiSemanticCachePluginProtocols {
	return &e
}
func (e *AiSemanticCachePluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AiSemanticCachePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiSemanticCachePluginProtocols: %v", v)
	}
}

// AiSemanticCachePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type AiSemanticCachePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiSemanticCachePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiSemanticCachePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AiSemanticCachePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiSemanticCachePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiSemanticCachePlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AiSemanticCachePlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                          `json:"enabled,omitempty"`
	ID           *string                        `json:"id,omitempty"`
	InstanceName *string                        `json:"instance_name,omitempty"`
	name         string                         `const:"ai-semantic-cache" json:"name"`
	Ordering     *AiSemanticCachePluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                      `json:"updated_at,omitempty"`
	Config    AiSemanticCachePluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AiSemanticCachePluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *AiSemanticCachePluginConsumerGroup `json:"consumer_group,omitempty"`
	// A set of strings representing HTTP protocols.
	Protocols []AiSemanticCachePluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *AiSemanticCachePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiSemanticCachePluginService `json:"service,omitempty"`
}

func (a AiSemanticCachePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiSemanticCachePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiSemanticCachePlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AiSemanticCachePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiSemanticCachePlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiSemanticCachePlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiSemanticCachePlugin) GetName() string {
	return "ai-semantic-cache"
}

func (o *AiSemanticCachePlugin) GetOrdering() *AiSemanticCachePluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiSemanticCachePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiSemanticCachePlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AiSemanticCachePlugin) GetConfig() AiSemanticCachePluginConfig {
	if o == nil {
		return AiSemanticCachePluginConfig{}
	}
	return o.Config
}

func (o *AiSemanticCachePlugin) GetConsumer() *AiSemanticCachePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiSemanticCachePlugin) GetConsumerGroup() *AiSemanticCachePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiSemanticCachePlugin) GetProtocols() []AiSemanticCachePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiSemanticCachePlugin) GetRoute() *AiSemanticCachePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiSemanticCachePlugin) GetService() *AiSemanticCachePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// AiSemanticCachePluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AiSemanticCachePluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool                          `json:"enabled,omitempty"`
	ID           *string                        `json:"id,omitempty"`
	InstanceName *string                        `json:"instance_name,omitempty"`
	name         string                         `const:"ai-semantic-cache" json:"name"`
	Ordering     *AiSemanticCachePluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string                    `json:"tags,omitempty"`
	Config AiSemanticCachePluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AiSemanticCachePluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups
	ConsumerGroup *AiSemanticCachePluginConsumerGroup `json:"consumer_group,omitempty"`
	// A set of strings representing HTTP protocols.
	Protocols []AiSemanticCachePluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *AiSemanticCachePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiSemanticCachePluginService `json:"service,omitempty"`
}

func (a AiSemanticCachePluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiSemanticCachePluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiSemanticCachePluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiSemanticCachePluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiSemanticCachePluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiSemanticCachePluginInput) GetName() string {
	return "ai-semantic-cache"
}

func (o *AiSemanticCachePluginInput) GetOrdering() *AiSemanticCachePluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiSemanticCachePluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiSemanticCachePluginInput) GetConfig() AiSemanticCachePluginConfig {
	if o == nil {
		return AiSemanticCachePluginConfig{}
	}
	return o.Config
}

func (o *AiSemanticCachePluginInput) GetConsumer() *AiSemanticCachePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiSemanticCachePluginInput) GetConsumerGroup() *AiSemanticCachePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiSemanticCachePluginInput) GetProtocols() []AiSemanticCachePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiSemanticCachePluginInput) GetRoute() *AiSemanticCachePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiSemanticCachePluginInput) GetService() *AiSemanticCachePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
