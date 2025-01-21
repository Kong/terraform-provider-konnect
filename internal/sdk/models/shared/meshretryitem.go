// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// MeshRetryItemType - the type of the resource
type MeshRetryItemType string

const (
	MeshRetryItemTypeMeshRetry MeshRetryItemType = "MeshRetry"
)

func (e MeshRetryItemType) ToPointer() *MeshRetryItemType {
	return &e
}
func (e *MeshRetryItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshRetry":
		*e = MeshRetryItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemType: %v", v)
	}
}

// MeshRetryItemKind - Kind of the referenced resource
type MeshRetryItemKind string

const (
	MeshRetryItemKindMesh                 MeshRetryItemKind = "Mesh"
	MeshRetryItemKindMeshSubset           MeshRetryItemKind = "MeshSubset"
	MeshRetryItemKindMeshGateway          MeshRetryItemKind = "MeshGateway"
	MeshRetryItemKindMeshService          MeshRetryItemKind = "MeshService"
	MeshRetryItemKindMeshExternalService  MeshRetryItemKind = "MeshExternalService"
	MeshRetryItemKindMeshMultiZoneService MeshRetryItemKind = "MeshMultiZoneService"
	MeshRetryItemKindMeshServiceSubset    MeshRetryItemKind = "MeshServiceSubset"
	MeshRetryItemKindMeshHTTPRoute        MeshRetryItemKind = "MeshHTTPRoute"
	MeshRetryItemKindDataplane            MeshRetryItemKind = "Dataplane"
)

func (e MeshRetryItemKind) ToPointer() *MeshRetryItemKind {
	return &e
}
func (e *MeshRetryItemKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Mesh":
		fallthrough
	case "MeshSubset":
		fallthrough
	case "MeshGateway":
		fallthrough
	case "MeshService":
		fallthrough
	case "MeshExternalService":
		fallthrough
	case "MeshMultiZoneService":
		fallthrough
	case "MeshServiceSubset":
		fallthrough
	case "MeshHTTPRoute":
		fallthrough
	case "Dataplane":
		*e = MeshRetryItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemKind: %v", v)
	}
}

type MeshRetryItemProxyTypes string

const (
	MeshRetryItemProxyTypesSidecar MeshRetryItemProxyTypes = "Sidecar"
	MeshRetryItemProxyTypesGateway MeshRetryItemProxyTypes = "Gateway"
)

func (e MeshRetryItemProxyTypes) ToPointer() *MeshRetryItemProxyTypes {
	return &e
}
func (e *MeshRetryItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshRetryItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemProxyTypes: %v", v)
	}
}

// MeshRetryItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined inplace.
type MeshRetryItemTargetRef struct {
	// Kind of the referenced resource
	Kind *MeshRetryItemKind `json:"kind,omitempty"`
	// Labels are used to select group of MeshServices that match labels. Either Labels or
	// Name and Namespace can be used.
	Labels map[string]string `json:"labels,omitempty"`
	// Mesh is reserved for future use to identify cross mesh resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the referenced resource. Can only be used with kinds: `MeshService`,
	// `MeshServiceSubset` and `MeshGatewayRoute`
	Name *string `json:"name,omitempty"`
	// Namespace specifies the namespace of target resource. If empty only resources in policy namespace
	// will be targeted.
	Namespace *string `json:"namespace,omitempty"`
	// ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
	// all data plane types are targeted by the policy.
	ProxyTypes []MeshRetryItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshRetryItemTargetRef) GetKind() *MeshRetryItemKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *MeshRetryItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRetryItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRetryItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshRetryItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshRetryItemTargetRef) GetProxyTypes() []MeshRetryItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshRetryItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshRetryItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// BackOff is a configuration of durations which will be used in an exponential
// backoff strategy between retries.
type BackOff struct {
	// BaseInterval is an amount of time which should be taken between retries.
	// Must be greater than zero. Values less than 1 ms are rounded up to 1 ms.
	BaseInterval *string `json:"baseInterval,omitempty"`
	// MaxInterval is a maximal amount of time which will be taken between retries.
	// Default is 10 times the "BaseInterval".
	MaxInterval *string `json:"maxInterval,omitempty"`
}

func (o *BackOff) GetBaseInterval() *string {
	if o == nil {
		return nil
	}
	return o.BaseInterval
}

func (o *BackOff) GetMaxInterval() *string {
	if o == nil {
		return nil
	}
	return o.MaxInterval
}

// MeshRetryItemFormat - The format of the reset header.
type MeshRetryItemFormat string

const (
	MeshRetryItemFormatSeconds       MeshRetryItemFormat = "Seconds"
	MeshRetryItemFormatUnixTimestamp MeshRetryItemFormat = "UnixTimestamp"
)

func (e MeshRetryItemFormat) ToPointer() *MeshRetryItemFormat {
	return &e
}
func (e *MeshRetryItemFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Seconds":
		fallthrough
	case "UnixTimestamp":
		*e = MeshRetryItemFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemFormat: %v", v)
	}
}

type ResetHeaders struct {
	// The format of the reset header.
	Format MeshRetryItemFormat `json:"format"`
	// The Name of the reset header.
	Name string `json:"name"`
}

func (o *ResetHeaders) GetFormat() MeshRetryItemFormat {
	if o == nil {
		return MeshRetryItemFormat("")
	}
	return o.Format
}

func (o *ResetHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// RateLimitedBackOff is a configuration of backoff which will be used when
// the upstream returns one of the headers configured.
type RateLimitedBackOff struct {
	// MaxInterval is a maximal amount of time which will be taken between retries.
	MaxInterval *string `json:"maxInterval,omitempty"`
	// ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset)
	// to match against the response. Headers are tried in order, and matched
	// case-insensitive. The first header to be parsed successfully is used.
	// If no headers match the default exponential BackOff is used instead.
	ResetHeaders []ResetHeaders `json:"resetHeaders,omitempty"`
}

func (o *RateLimitedBackOff) GetMaxInterval() *string {
	if o == nil {
		return nil
	}
	return o.MaxInterval
}

func (o *RateLimitedBackOff) GetResetHeaders() []ResetHeaders {
	if o == nil {
		return nil
	}
	return o.ResetHeaders
}

type RetryOn string

const (
	RetryOnCanceled          RetryOn = "Canceled"
	RetryOnDeadlineExceeded  RetryOn = "DeadlineExceeded"
	RetryOnInternal          RetryOn = "Internal"
	RetryOnResourceExhausted RetryOn = "ResourceExhausted"
	RetryOnUnavailable       RetryOn = "Unavailable"
)

func (e RetryOn) ToPointer() *RetryOn {
	return &e
}
func (e *RetryOn) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Canceled":
		fallthrough
	case "DeadlineExceeded":
		fallthrough
	case "Internal":
		fallthrough
	case "ResourceExhausted":
		fallthrough
	case "Unavailable":
		*e = RetryOn(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetryOn: %v", v)
	}
}

// MeshRetryItemGrpc - GRPC defines a configuration of retries for GRPC traffic
type MeshRetryItemGrpc struct {
	// BackOff is a configuration of durations which will be used in an exponential
	// backoff strategy between retries.
	BackOff *BackOff `json:"backOff,omitempty"`
	// NumRetries is the number of attempts that will be made on failed (and
	// retriable) requests. If not set, the default value is 1.
	NumRetries *int `json:"numRetries,omitempty"`
	// PerTryTimeout is the maximum amount of time each retry attempt can take
	// before it times out. If not set, the global request timeout for the route
	// will be used. Setting this value to 0 will disable the per-try timeout.
	PerTryTimeout *string `json:"perTryTimeout,omitempty"`
	// RateLimitedBackOff is a configuration of backoff which will be used when
	// the upstream returns one of the headers configured.
	RateLimitedBackOff *RateLimitedBackOff `json:"rateLimitedBackOff,omitempty"`
	// RetryOn is a list of conditions which will cause a retry.
	RetryOn []RetryOn `json:"retryOn,omitempty"`
}

func (o *MeshRetryItemGrpc) GetBackOff() *BackOff {
	if o == nil {
		return nil
	}
	return o.BackOff
}

func (o *MeshRetryItemGrpc) GetNumRetries() *int {
	if o == nil {
		return nil
	}
	return o.NumRetries
}

func (o *MeshRetryItemGrpc) GetPerTryTimeout() *string {
	if o == nil {
		return nil
	}
	return o.PerTryTimeout
}

func (o *MeshRetryItemGrpc) GetRateLimitedBackOff() *RateLimitedBackOff {
	if o == nil {
		return nil
	}
	return o.RateLimitedBackOff
}

func (o *MeshRetryItemGrpc) GetRetryOn() []RetryOn {
	if o == nil {
		return nil
	}
	return o.RetryOn
}

// MeshRetryItemBackOff - BackOff is a configuration of durations which will be used in exponential
// backoff strategy between retries.
type MeshRetryItemBackOff struct {
	// BaseInterval is an amount of time which should be taken between retries.
	// Must be greater than zero. Values less than 1 ms are rounded up to 1 ms.
	BaseInterval *string `json:"baseInterval,omitempty"`
	// MaxInterval is a maximal amount of time which will be taken between retries.
	// Default is 10 times the "BaseInterval".
	MaxInterval *string `json:"maxInterval,omitempty"`
}

func (o *MeshRetryItemBackOff) GetBaseInterval() *string {
	if o == nil {
		return nil
	}
	return o.BaseInterval
}

func (o *MeshRetryItemBackOff) GetMaxInterval() *string {
	if o == nil {
		return nil
	}
	return o.MaxInterval
}

// Predicate - Type is requested predicate mode.
type Predicate string

const (
	PredicateOmitPreviousHosts      Predicate = "OmitPreviousHosts"
	PredicateOmitHostsWithTags      Predicate = "OmitHostsWithTags"
	PredicateOmitPreviousPriorities Predicate = "OmitPreviousPriorities"
)

func (e Predicate) ToPointer() *Predicate {
	return &e
}
func (e *Predicate) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OmitPreviousHosts":
		fallthrough
	case "OmitHostsWithTags":
		fallthrough
	case "OmitPreviousPriorities":
		*e = Predicate(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Predicate: %v", v)
	}
}

type HostSelection struct {
	// Type is requested predicate mode.
	Predicate Predicate `json:"predicate"`
	// Tags is a map of metadata to match against for selecting the omitted hosts. Required if Type is
	// OmitHostsWithTags
	Tags map[string]string `json:"tags,omitempty"`
	// UpdateFrequency is how often the priority load should be updated based on previously attempted priorities.
	// Used for OmitPreviousPriorities.
	UpdateFrequency *int `json:"updateFrequency,omitempty"`
}

func (o *HostSelection) GetPredicate() Predicate {
	if o == nil {
		return Predicate("")
	}
	return o.Predicate
}

func (o *HostSelection) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *HostSelection) GetUpdateFrequency() *int {
	if o == nil {
		return nil
	}
	return o.UpdateFrequency
}

// MeshRetryItemSpecFormat - The format of the reset header.
type MeshRetryItemSpecFormat string

const (
	MeshRetryItemSpecFormatSeconds       MeshRetryItemSpecFormat = "Seconds"
	MeshRetryItemSpecFormatUnixTimestamp MeshRetryItemSpecFormat = "UnixTimestamp"
)

func (e MeshRetryItemSpecFormat) ToPointer() *MeshRetryItemSpecFormat {
	return &e
}
func (e *MeshRetryItemSpecFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Seconds":
		fallthrough
	case "UnixTimestamp":
		*e = MeshRetryItemSpecFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemSpecFormat: %v", v)
	}
}

type MeshRetryItemResetHeaders struct {
	// The format of the reset header.
	Format MeshRetryItemSpecFormat `json:"format"`
	// The Name of the reset header.
	Name string `json:"name"`
}

func (o *MeshRetryItemResetHeaders) GetFormat() MeshRetryItemSpecFormat {
	if o == nil {
		return MeshRetryItemSpecFormat("")
	}
	return o.Format
}

func (o *MeshRetryItemResetHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// MeshRetryItemRateLimitedBackOff - RateLimitedBackOff is a configuration of backoff which will be used
// when the upstream returns one of the headers configured.
type MeshRetryItemRateLimitedBackOff struct {
	// MaxInterval is a maximal amount of time which will be taken between retries.
	MaxInterval *string `json:"maxInterval,omitempty"`
	// ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset)
	// to match against the response. Headers are tried in order, and matched
	// case-insensitive. The first header to be parsed successfully is used.
	// If no headers match the default exponential BackOff is used instead.
	ResetHeaders []MeshRetryItemResetHeaders `json:"resetHeaders,omitempty"`
}

func (o *MeshRetryItemRateLimitedBackOff) GetMaxInterval() *string {
	if o == nil {
		return nil
	}
	return o.MaxInterval
}

func (o *MeshRetryItemRateLimitedBackOff) GetResetHeaders() []MeshRetryItemResetHeaders {
	if o == nil {
		return nil
	}
	return o.ResetHeaders
}

// MeshRetryItemSpecType - Type specifies how to match against the value of the header.
type MeshRetryItemSpecType string

const (
	MeshRetryItemSpecTypeExact             MeshRetryItemSpecType = "Exact"
	MeshRetryItemSpecTypePresent           MeshRetryItemSpecType = "Present"
	MeshRetryItemSpecTypeRegularExpression MeshRetryItemSpecType = "RegularExpression"
	MeshRetryItemSpecTypeAbsent            MeshRetryItemSpecType = "Absent"
	MeshRetryItemSpecTypePrefix            MeshRetryItemSpecType = "Prefix"
)

func (e MeshRetryItemSpecType) ToPointer() *MeshRetryItemSpecType {
	return &e
}
func (e *MeshRetryItemSpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Exact":
		fallthrough
	case "Present":
		fallthrough
	case "RegularExpression":
		fallthrough
	case "Absent":
		fallthrough
	case "Prefix":
		*e = MeshRetryItemSpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemSpecType: %v", v)
	}
}

// RetriableRequestHeaders - HeaderMatch describes how to select an HTTP route by matching HTTP request
// headers.
type RetriableRequestHeaders struct {
	// Name is the name of the HTTP Header to be matched. Name MUST be lower case
	// as they will be handled with case insensitivity (See https://tools.ietf.org/html/rfc7230#section-3.2).
	Name string `json:"name"`
	// Type specifies how to match against the value of the header.
	Type *MeshRetryItemSpecType `json:"type,omitempty"`
	// Value is the value of HTTP Header to be matched.
	Value *string `json:"value,omitempty"`
}

func (o *RetriableRequestHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RetriableRequestHeaders) GetType() *MeshRetryItemSpecType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *RetriableRequestHeaders) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// MeshRetryItemSpecToType - Type specifies how to match against the value of the header.
type MeshRetryItemSpecToType string

const (
	MeshRetryItemSpecToTypeExact             MeshRetryItemSpecToType = "Exact"
	MeshRetryItemSpecToTypePresent           MeshRetryItemSpecToType = "Present"
	MeshRetryItemSpecToTypeRegularExpression MeshRetryItemSpecToType = "RegularExpression"
	MeshRetryItemSpecToTypeAbsent            MeshRetryItemSpecToType = "Absent"
	MeshRetryItemSpecToTypePrefix            MeshRetryItemSpecToType = "Prefix"
)

func (e MeshRetryItemSpecToType) ToPointer() *MeshRetryItemSpecToType {
	return &e
}
func (e *MeshRetryItemSpecToType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Exact":
		fallthrough
	case "Present":
		fallthrough
	case "RegularExpression":
		fallthrough
	case "Absent":
		fallthrough
	case "Prefix":
		*e = MeshRetryItemSpecToType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemSpecToType: %v", v)
	}
}

// RetriableResponseHeaders - HeaderMatch describes how to select an HTTP route by matching HTTP request
// headers.
type RetriableResponseHeaders struct {
	// Name is the name of the HTTP Header to be matched. Name MUST be lower case
	// as they will be handled with case insensitivity (See https://tools.ietf.org/html/rfc7230#section-3.2).
	Name string `json:"name"`
	// Type specifies how to match against the value of the header.
	Type *MeshRetryItemSpecToType `json:"type,omitempty"`
	// Value is the value of HTTP Header to be matched.
	Value *string `json:"value,omitempty"`
}

func (o *RetriableResponseHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RetriableResponseHeaders) GetType() *MeshRetryItemSpecToType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *RetriableResponseHeaders) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// MeshRetryItemHTTP - HTTP defines a configuration of retries for HTTP traffic
type MeshRetryItemHTTP struct {
	// BackOff is a configuration of durations which will be used in exponential
	// backoff strategy between retries.
	BackOff *MeshRetryItemBackOff `json:"backOff,omitempty"`
	// HostSelection is a list of predicates that dictate how hosts should be selected
	// when requests are retried.
	HostSelection []HostSelection `json:"hostSelection,omitempty"`
	// HostSelectionMaxAttempts is the maximum number of times host selection will be
	// reattempted before giving up, at which point the host that was last selected will
	// be routed to. If unspecified, this will default to retrying once.
	HostSelectionMaxAttempts *int64 `json:"hostSelectionMaxAttempts,omitempty"`
	// NumRetries is the number of attempts that will be made on failed (and
	// retriable) requests.  If not set, the default value is 1.
	NumRetries *int `json:"numRetries,omitempty"`
	// PerTryTimeout is the amount of time after which retry attempt should time out.
	// If left unspecified, the global route timeout for the request will be used.
	// Consequently, when using a 5xx based retry policy, a request that times out
	// will not be retried as the total timeout budget would have been exhausted.
	// Setting this timeout to 0 will disable it.
	PerTryTimeout *string `json:"perTryTimeout,omitempty"`
	// RateLimitedBackOff is a configuration of backoff which will be used
	// when the upstream returns one of the headers configured.
	RateLimitedBackOff *MeshRetryItemRateLimitedBackOff `json:"rateLimitedBackOff,omitempty"`
	// RetriableRequestHeaders is an HTTP headers which must be present in the request
	// for retries to be attempted.
	RetriableRequestHeaders []RetriableRequestHeaders `json:"retriableRequestHeaders,omitempty"`
	// RetriableResponseHeaders is an HTTP response headers that trigger a retry
	// if present in the response. A retry will be triggered if any of the header
	// matches the upstream response headers.
	RetriableResponseHeaders []RetriableResponseHeaders `json:"retriableResponseHeaders,omitempty"`
	// RetryOn is a list of conditions which will cause a retry. Available values are:
	// [5XX, GatewayError, Reset, Retriable4xx, ConnectFailure, EnvoyRatelimited,
	// RefusedStream, Http3PostConnectFailure, HttpMethodConnect, HttpMethodDelete,
	// HttpMethodGet, HttpMethodHead, HttpMethodOptions, HttpMethodPatch,
	// HttpMethodPost, HttpMethodPut, HttpMethodTrace].
	// Also, any HTTP status code (500, 503, etc.).
	RetryOn []string `json:"retryOn,omitempty"`
}

func (o *MeshRetryItemHTTP) GetBackOff() *MeshRetryItemBackOff {
	if o == nil {
		return nil
	}
	return o.BackOff
}

func (o *MeshRetryItemHTTP) GetHostSelection() []HostSelection {
	if o == nil {
		return nil
	}
	return o.HostSelection
}

func (o *MeshRetryItemHTTP) GetHostSelectionMaxAttempts() *int64 {
	if o == nil {
		return nil
	}
	return o.HostSelectionMaxAttempts
}

func (o *MeshRetryItemHTTP) GetNumRetries() *int {
	if o == nil {
		return nil
	}
	return o.NumRetries
}

func (o *MeshRetryItemHTTP) GetPerTryTimeout() *string {
	if o == nil {
		return nil
	}
	return o.PerTryTimeout
}

func (o *MeshRetryItemHTTP) GetRateLimitedBackOff() *MeshRetryItemRateLimitedBackOff {
	if o == nil {
		return nil
	}
	return o.RateLimitedBackOff
}

func (o *MeshRetryItemHTTP) GetRetriableRequestHeaders() []RetriableRequestHeaders {
	if o == nil {
		return nil
	}
	return o.RetriableRequestHeaders
}

func (o *MeshRetryItemHTTP) GetRetriableResponseHeaders() []RetriableResponseHeaders {
	if o == nil {
		return nil
	}
	return o.RetriableResponseHeaders
}

func (o *MeshRetryItemHTTP) GetRetryOn() []string {
	if o == nil {
		return nil
	}
	return o.RetryOn
}

// MeshRetryItemTCP - TCP defines a configuration of retries for TCP traffic
type MeshRetryItemTCP struct {
	// MaxConnectAttempt is a maximal amount of TCP connection attempts
	// which will be made before giving up
	MaxConnectAttempt *int `json:"maxConnectAttempt,omitempty"`
}

func (o *MeshRetryItemTCP) GetMaxConnectAttempt() *int {
	if o == nil {
		return nil
	}
	return o.MaxConnectAttempt
}

// MeshRetryItemDefault - Default is a configuration specific to the group of destinations referenced in
// 'targetRef'
type MeshRetryItemDefault struct {
	// GRPC defines a configuration of retries for GRPC traffic
	Grpc *MeshRetryItemGrpc `json:"grpc,omitempty"`
	// HTTP defines a configuration of retries for HTTP traffic
	HTTP *MeshRetryItemHTTP `json:"http,omitempty"`
	// TCP defines a configuration of retries for TCP traffic
	TCP *MeshRetryItemTCP `json:"tcp,omitempty"`
}

func (o *MeshRetryItemDefault) GetGrpc() *MeshRetryItemGrpc {
	if o == nil {
		return nil
	}
	return o.Grpc
}

func (o *MeshRetryItemDefault) GetHTTP() *MeshRetryItemHTTP {
	if o == nil {
		return nil
	}
	return o.HTTP
}

func (o *MeshRetryItemDefault) GetTCP() *MeshRetryItemTCP {
	if o == nil {
		return nil
	}
	return o.TCP
}

// MeshRetryItemSpecKind - Kind of the referenced resource
type MeshRetryItemSpecKind string

const (
	MeshRetryItemSpecKindMesh                 MeshRetryItemSpecKind = "Mesh"
	MeshRetryItemSpecKindMeshSubset           MeshRetryItemSpecKind = "MeshSubset"
	MeshRetryItemSpecKindMeshGateway          MeshRetryItemSpecKind = "MeshGateway"
	MeshRetryItemSpecKindMeshService          MeshRetryItemSpecKind = "MeshService"
	MeshRetryItemSpecKindMeshExternalService  MeshRetryItemSpecKind = "MeshExternalService"
	MeshRetryItemSpecKindMeshMultiZoneService MeshRetryItemSpecKind = "MeshMultiZoneService"
	MeshRetryItemSpecKindMeshServiceSubset    MeshRetryItemSpecKind = "MeshServiceSubset"
	MeshRetryItemSpecKindMeshHTTPRoute        MeshRetryItemSpecKind = "MeshHTTPRoute"
	MeshRetryItemSpecKindDataplane            MeshRetryItemSpecKind = "Dataplane"
)

func (e MeshRetryItemSpecKind) ToPointer() *MeshRetryItemSpecKind {
	return &e
}
func (e *MeshRetryItemSpecKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Mesh":
		fallthrough
	case "MeshSubset":
		fallthrough
	case "MeshGateway":
		fallthrough
	case "MeshService":
		fallthrough
	case "MeshExternalService":
		fallthrough
	case "MeshMultiZoneService":
		fallthrough
	case "MeshServiceSubset":
		fallthrough
	case "MeshHTTPRoute":
		fallthrough
	case "Dataplane":
		*e = MeshRetryItemSpecKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemSpecKind: %v", v)
	}
}

type MeshRetryItemSpecProxyTypes string

const (
	MeshRetryItemSpecProxyTypesSidecar MeshRetryItemSpecProxyTypes = "Sidecar"
	MeshRetryItemSpecProxyTypesGateway MeshRetryItemSpecProxyTypes = "Gateway"
)

func (e MeshRetryItemSpecProxyTypes) ToPointer() *MeshRetryItemSpecProxyTypes {
	return &e
}
func (e *MeshRetryItemSpecProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshRetryItemSpecProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRetryItemSpecProxyTypes: %v", v)
	}
}

// MeshRetryItemSpecTargetRef - TargetRef is a reference to the resource that represents a group of
// destinations.
type MeshRetryItemSpecTargetRef struct {
	// Kind of the referenced resource
	Kind *MeshRetryItemSpecKind `json:"kind,omitempty"`
	// Labels are used to select group of MeshServices that match labels. Either Labels or
	// Name and Namespace can be used.
	Labels map[string]string `json:"labels,omitempty"`
	// Mesh is reserved for future use to identify cross mesh resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the referenced resource. Can only be used with kinds: `MeshService`,
	// `MeshServiceSubset` and `MeshGatewayRoute`
	Name *string `json:"name,omitempty"`
	// Namespace specifies the namespace of target resource. If empty only resources in policy namespace
	// will be targeted.
	Namespace *string `json:"namespace,omitempty"`
	// ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
	// all data plane types are targeted by the policy.
	ProxyTypes []MeshRetryItemSpecProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshRetryItemSpecTargetRef) GetKind() *MeshRetryItemSpecKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *MeshRetryItemSpecTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRetryItemSpecTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRetryItemSpecTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshRetryItemSpecTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshRetryItemSpecTargetRef) GetProxyTypes() []MeshRetryItemSpecProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshRetryItemSpecTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshRetryItemSpecTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshRetryItemTo struct {
	// Default is a configuration specific to the group of destinations referenced in
	// 'targetRef'
	Default *MeshRetryItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource that represents a group of
	// destinations.
	TargetRef MeshRetryItemSpecTargetRef `json:"targetRef"`
}

func (o *MeshRetryItemTo) GetDefault() *MeshRetryItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshRetryItemTo) GetTargetRef() MeshRetryItemSpecTargetRef {
	if o == nil {
		return MeshRetryItemSpecTargetRef{}
	}
	return o.TargetRef
}

// MeshRetryItemSpec - Spec is the specification of the Kuma MeshRetry resource.
type MeshRetryItemSpec struct {
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *MeshRetryItemTargetRef `json:"targetRef,omitempty"`
	// To list makes a match between the consumed services and corresponding configurations
	To []MeshRetryItemTo `json:"to,omitempty"`
}

func (o *MeshRetryItemSpec) GetTargetRef() *MeshRetryItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

func (o *MeshRetryItemSpec) GetTo() []MeshRetryItemTo {
	if o == nil {
		return nil
	}
	return o.To
}

type MeshRetryItem struct {
	// the type of the resource
	Type MeshRetryItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshRetry resource.
	Spec MeshRetryItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshRetryItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshRetryItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshRetryItem) GetType() MeshRetryItemType {
	if o == nil {
		return MeshRetryItemType("")
	}
	return o.Type
}

func (o *MeshRetryItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRetryItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRetryItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRetryItem) GetSpec() MeshRetryItemSpec {
	if o == nil {
		return MeshRetryItemSpec{}
	}
	return o.Spec
}

func (o *MeshRetryItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshRetryItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}
