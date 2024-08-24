// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateProxyCachePluginMemory struct {
	// The name of the shared dictionary in which to hold cache entities when the memory strategy is selected. Note that this dictionary currently must be defined manually in the Kong Nginx template.
	DictionaryName *string `json:"dictionary_name,omitempty"`
}

func (o *CreateProxyCachePluginMemory) GetDictionaryName() *string {
	if o == nil {
		return nil
	}
	return o.DictionaryName
}

type CreateProxyCachePluginRequestMethod string

const (
	CreateProxyCachePluginRequestMethodHead  CreateProxyCachePluginRequestMethod = "HEAD"
	CreateProxyCachePluginRequestMethodGet   CreateProxyCachePluginRequestMethod = "GET"
	CreateProxyCachePluginRequestMethodPost  CreateProxyCachePluginRequestMethod = "POST"
	CreateProxyCachePluginRequestMethodPatch CreateProxyCachePluginRequestMethod = "PATCH"
	CreateProxyCachePluginRequestMethodPut   CreateProxyCachePluginRequestMethod = "PUT"
)

func (e CreateProxyCachePluginRequestMethod) ToPointer() *CreateProxyCachePluginRequestMethod {
	return &e
}
func (e *CreateProxyCachePluginRequestMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HEAD":
		fallthrough
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "PATCH":
		fallthrough
	case "PUT":
		*e = CreateProxyCachePluginRequestMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCachePluginRequestMethod: %v", v)
	}
}

// CreateProxyCachePluginResponseHeaders - Caching related diagnostic headers that should be included in cached responses
type CreateProxyCachePluginResponseHeaders struct {
	XCacheKey    *bool `json:"X-Cache-Key,omitempty"`
	XCacheStatus *bool `json:"X-Cache-Status,omitempty"`
	Age          *bool `json:"age,omitempty"`
}

func (o *CreateProxyCachePluginResponseHeaders) GetXCacheKey() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheKey
}

func (o *CreateProxyCachePluginResponseHeaders) GetXCacheStatus() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheStatus
}

func (o *CreateProxyCachePluginResponseHeaders) GetAge() *bool {
	if o == nil {
		return nil
	}
	return o.Age
}

// CreateProxyCachePluginStrategy - The backing data store in which to hold cache entities.
type CreateProxyCachePluginStrategy string

const (
	CreateProxyCachePluginStrategyMemory CreateProxyCachePluginStrategy = "memory"
)

func (e CreateProxyCachePluginStrategy) ToPointer() *CreateProxyCachePluginStrategy {
	return &e
}
func (e *CreateProxyCachePluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "memory":
		*e = CreateProxyCachePluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCachePluginStrategy: %v", v)
	}
}

type CreateProxyCachePluginConfig struct {
	// When enabled, respect the Cache-Control behaviors defined in RFC7234.
	CacheControl *bool `json:"cache_control,omitempty"`
	// TTL, in seconds, of cache entities.
	CacheTTL *int64 `json:"cache_ttl,omitempty"`
	// Upstream response content types considered cacheable. The plugin performs an **exact match** against each specified value.
	ContentType   []string                      `json:"content_type,omitempty"`
	IgnoreURICase *bool                         `json:"ignore_uri_case,omitempty"`
	Memory        *CreateProxyCachePluginMemory `json:"memory,omitempty"`
	// Downstream request methods considered cacheable.
	RequestMethod []CreateProxyCachePluginRequestMethod `json:"request_method,omitempty"`
	// Upstream response status code considered cacheable.
	ResponseCode []int64 `json:"response_code,omitempty"`
	// Caching related diagnostic headers that should be included in cached responses
	ResponseHeaders *CreateProxyCachePluginResponseHeaders `json:"response_headers,omitempty"`
	// Number of seconds to keep resources in the storage backend. This value is independent of `cache_ttl` or resource TTLs defined by Cache-Control behaviors.
	StorageTTL *int64 `json:"storage_ttl,omitempty"`
	// The backing data store in which to hold cache entities.
	Strategy *CreateProxyCachePluginStrategy `json:"strategy,omitempty"`
	// Relevant headers considered for the cache key. If undefined, none of the headers are taken into consideration.
	VaryHeaders []string `json:"vary_headers,omitempty"`
	// Relevant query parameters considered for the cache key. If undefined, all params are taken into consideration.
	VaryQueryParams []string `json:"vary_query_params,omitempty"`
}

func (o *CreateProxyCachePluginConfig) GetCacheControl() *bool {
	if o == nil {
		return nil
	}
	return o.CacheControl
}

func (o *CreateProxyCachePluginConfig) GetCacheTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *CreateProxyCachePluginConfig) GetContentType() []string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *CreateProxyCachePluginConfig) GetIgnoreURICase() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreURICase
}

func (o *CreateProxyCachePluginConfig) GetMemory() *CreateProxyCachePluginMemory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *CreateProxyCachePluginConfig) GetRequestMethod() []CreateProxyCachePluginRequestMethod {
	if o == nil {
		return nil
	}
	return o.RequestMethod
}

func (o *CreateProxyCachePluginConfig) GetResponseCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseCode
}

func (o *CreateProxyCachePluginConfig) GetResponseHeaders() *CreateProxyCachePluginResponseHeaders {
	if o == nil {
		return nil
	}
	return o.ResponseHeaders
}

func (o *CreateProxyCachePluginConfig) GetStorageTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.StorageTTL
}

func (o *CreateProxyCachePluginConfig) GetStrategy() *CreateProxyCachePluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *CreateProxyCachePluginConfig) GetVaryHeaders() []string {
	if o == nil {
		return nil
	}
	return o.VaryHeaders
}

func (o *CreateProxyCachePluginConfig) GetVaryQueryParams() []string {
	if o == nil {
		return nil
	}
	return o.VaryQueryParams
}

type CreateProxyCachePluginProtocols string

const (
	CreateProxyCachePluginProtocolsGrpc           CreateProxyCachePluginProtocols = "grpc"
	CreateProxyCachePluginProtocolsGrpcs          CreateProxyCachePluginProtocols = "grpcs"
	CreateProxyCachePluginProtocolsHTTP           CreateProxyCachePluginProtocols = "http"
	CreateProxyCachePluginProtocolsHTTPS          CreateProxyCachePluginProtocols = "https"
	CreateProxyCachePluginProtocolsTCP            CreateProxyCachePluginProtocols = "tcp"
	CreateProxyCachePluginProtocolsTLS            CreateProxyCachePluginProtocols = "tls"
	CreateProxyCachePluginProtocolsTLSPassthrough CreateProxyCachePluginProtocols = "tls_passthrough"
	CreateProxyCachePluginProtocolsUDP            CreateProxyCachePluginProtocols = "udp"
	CreateProxyCachePluginProtocolsWs             CreateProxyCachePluginProtocols = "ws"
	CreateProxyCachePluginProtocolsWss            CreateProxyCachePluginProtocols = "wss"
)

func (e CreateProxyCachePluginProtocols) ToPointer() *CreateProxyCachePluginProtocols {
	return &e
}
func (e *CreateProxyCachePluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateProxyCachePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCachePluginProtocols: %v", v)
	}
}

// CreateProxyCachePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateProxyCachePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCachePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateProxyCachePluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCachePluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateProxyCachePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateProxyCachePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCachePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateProxyCachePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateProxyCachePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCachePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateProxyCachePlugin struct {
	Config *CreateProxyCachePluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"proxy-cache" json:"name,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateProxyCachePluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateProxyCachePluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateProxyCachePluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateProxyCachePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateProxyCachePluginService `json:"service,omitempty"`
}

func (c CreateProxyCachePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateProxyCachePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateProxyCachePlugin) GetConfig() *CreateProxyCachePluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateProxyCachePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateProxyCachePlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateProxyCachePlugin) GetName() *string {
	return types.String("proxy-cache")
}

func (o *CreateProxyCachePlugin) GetProtocols() []CreateProxyCachePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateProxyCachePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateProxyCachePlugin) GetConsumer() *CreateProxyCachePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateProxyCachePlugin) GetConsumerGroup() *CreateProxyCachePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateProxyCachePlugin) GetRoute() *CreateProxyCachePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateProxyCachePlugin) GetService() *CreateProxyCachePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
