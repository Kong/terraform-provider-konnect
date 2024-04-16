// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

// CreateUpstreamAlgorithm - Which load balancing algorithm to use.
type CreateUpstreamAlgorithm string

const (
	CreateUpstreamAlgorithmConsistentHashing CreateUpstreamAlgorithm = "consistent-hashing"
	CreateUpstreamAlgorithmLeastConnections  CreateUpstreamAlgorithm = "least-connections"
	CreateUpstreamAlgorithmRoundRobin        CreateUpstreamAlgorithm = "round-robin"
)

func (e CreateUpstreamAlgorithm) ToPointer() *CreateUpstreamAlgorithm {
	return &e
}

func (e *CreateUpstreamAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consistent-hashing":
		fallthrough
	case "least-connections":
		fallthrough
	case "round-robin":
		*e = CreateUpstreamAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateUpstreamAlgorithm: %v", v)
	}
}

// CreateUpstreamClientCertificate - If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
type CreateUpstreamClientCertificate struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateUpstreamClientCertificate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// HashFallback - What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.
type HashFallback string

const (
	HashFallbackNone       HashFallback = "none"
	HashFallbackConsumer   HashFallback = "consumer"
	HashFallbackIP         HashFallback = "ip"
	HashFallbackHeader     HashFallback = "header"
	HashFallbackCookie     HashFallback = "cookie"
	HashFallbackPath       HashFallback = "path"
	HashFallbackQueryArg   HashFallback = "query_arg"
	HashFallbackURICapture HashFallback = "uri_capture"
)

func (e HashFallback) ToPointer() *HashFallback {
	return &e
}

func (e *HashFallback) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "consumer":
		fallthrough
	case "ip":
		fallthrough
	case "header":
		fallthrough
	case "cookie":
		fallthrough
	case "path":
		fallthrough
	case "query_arg":
		fallthrough
	case "uri_capture":
		*e = HashFallback(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HashFallback: %v", v)
	}
}

// HashOn - What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.
type HashOn string

const (
	HashOnNone       HashOn = "none"
	HashOnConsumer   HashOn = "consumer"
	HashOnIP         HashOn = "ip"
	HashOnHeader     HashOn = "header"
	HashOnCookie     HashOn = "cookie"
	HashOnPath       HashOn = "path"
	HashOnQueryArg   HashOn = "query_arg"
	HashOnURICapture HashOn = "uri_capture"
)

func (e HashOn) ToPointer() *HashOn {
	return &e
}

func (e *HashOn) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "consumer":
		fallthrough
	case "ip":
		fallthrough
	case "header":
		fallthrough
	case "cookie":
		fallthrough
	case "path":
		fallthrough
	case "query_arg":
		fallthrough
	case "uri_capture":
		*e = HashOn(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HashOn: %v", v)
	}
}

type Healthy struct {
	HTTPStatuses []int64  `json:"http_statuses,omitempty"`
	Interval     *float64 `default:"0" json:"interval"`
	Successes    *int64   `default:"0" json:"successes"`
}

func (h Healthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *Healthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Healthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *Healthy) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *Healthy) GetSuccesses() *int64 {
	if o == nil {
		return nil
	}
	return o.Successes
}

type Type string

const (
	TypeTCP   Type = "tcp"
	TypeHTTP  Type = "http"
	TypeHTTPS Type = "https"
	TypeGrpc  Type = "grpc"
	TypeGrpcs Type = "grpcs"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "grpc":
		fallthrough
	case "grpcs":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Unhealthy struct {
	HTTPFailures *int64   `default:"0" json:"http_failures"`
	HTTPStatuses []int64  `json:"http_statuses,omitempty"`
	Interval     *float64 `default:"0" json:"interval"`
	TCPFailures  *int64   `default:"0" json:"tcp_failures"`
	Timeouts     *int64   `default:"0" json:"timeouts"`
}

func (u Unhealthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *Unhealthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Unhealthy) GetHTTPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPFailures
}

func (o *Unhealthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *Unhealthy) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *Unhealthy) GetTCPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.TCPFailures
}

func (o *Unhealthy) GetTimeouts() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeouts
}

type Active struct {
	Concurrency            *int64                 `default:"10" json:"concurrency"`
	Headers                map[string]interface{} `json:"headers,omitempty"`
	Healthy                *Healthy               `json:"healthy,omitempty"`
	HTTPPath               *string                `default:"/" json:"http_path"`
	HTTPSSni               *string                `json:"https_sni,omitempty"`
	HTTPSVerifyCertificate *bool                  `default:"true" json:"https_verify_certificate"`
	Timeout                *float64               `default:"1" json:"timeout"`
	Type                   *Type                  `default:"http" json:"type"`
	Unhealthy              *Unhealthy             `json:"unhealthy,omitempty"`
}

func (a Active) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Active) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Active) GetConcurrency() *int64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *Active) GetHeaders() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Active) GetHealthy() *Healthy {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *Active) GetHTTPPath() *string {
	if o == nil {
		return nil
	}
	return o.HTTPPath
}

func (o *Active) GetHTTPSSni() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSSni
}

func (o *Active) GetHTTPSVerifyCertificate() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPSVerifyCertificate
}

func (o *Active) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *Active) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Active) GetUnhealthy() *Unhealthy {
	if o == nil {
		return nil
	}
	return o.Unhealthy
}

type CreateUpstreamHealthy struct {
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	Successes    *int64  `default:"0" json:"successes"`
}

func (c CreateUpstreamHealthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateUpstreamHealthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateUpstreamHealthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *CreateUpstreamHealthy) GetSuccesses() *int64 {
	if o == nil {
		return nil
	}
	return o.Successes
}

type CreateUpstreamType string

const (
	CreateUpstreamTypeTCP   CreateUpstreamType = "tcp"
	CreateUpstreamTypeHTTP  CreateUpstreamType = "http"
	CreateUpstreamTypeHTTPS CreateUpstreamType = "https"
	CreateUpstreamTypeGrpc  CreateUpstreamType = "grpc"
	CreateUpstreamTypeGrpcs CreateUpstreamType = "grpcs"
)

func (e CreateUpstreamType) ToPointer() *CreateUpstreamType {
	return &e
}

func (e *CreateUpstreamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "grpc":
		fallthrough
	case "grpcs":
		*e = CreateUpstreamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateUpstreamType: %v", v)
	}
}

type CreateUpstreamUnhealthy struct {
	HTTPFailures *int64  `default:"0" json:"http_failures"`
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	TCPFailures  *int64  `default:"0" json:"tcp_failures"`
	Timeouts     *int64  `default:"0" json:"timeouts"`
}

func (c CreateUpstreamUnhealthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateUpstreamUnhealthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateUpstreamUnhealthy) GetHTTPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPFailures
}

func (o *CreateUpstreamUnhealthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *CreateUpstreamUnhealthy) GetTCPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.TCPFailures
}

func (o *CreateUpstreamUnhealthy) GetTimeouts() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeouts
}

type Passive struct {
	Healthy   *CreateUpstreamHealthy   `json:"healthy,omitempty"`
	Type      *CreateUpstreamType      `default:"http" json:"type"`
	Unhealthy *CreateUpstreamUnhealthy `json:"unhealthy,omitempty"`
}

func (p Passive) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Passive) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Passive) GetHealthy() *CreateUpstreamHealthy {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *Passive) GetType() *CreateUpstreamType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Passive) GetUnhealthy() *CreateUpstreamUnhealthy {
	if o == nil {
		return nil
	}
	return o.Unhealthy
}

type Healthchecks struct {
	Active    *Active  `json:"active,omitempty"`
	Passive   *Passive `json:"passive,omitempty"`
	Threshold *float64 `default:"0" json:"threshold"`
}

func (h Healthchecks) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *Healthchecks) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Healthchecks) GetActive() *Active {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *Healthchecks) GetPassive() *Passive {
	if o == nil {
		return nil
	}
	return o.Passive
}

func (o *Healthchecks) GetThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.Threshold
}

// CreateUpstream - The upstream object represents a virtual hostname and can be used to loadbalance incoming requests over multiple services (targets). So for example an upstream named `service.v1.xyz` for a Service object whose `host` is `service.v1.xyz`. Requests for this Service would be proxied to the targets defined within the upstream. An upstream also includes a [health checker][healthchecks], which is able to enable and disable targets based on their ability or inability to serve requests. The configuration for the health checker is stored in the upstream object, and applies to all of its targets.
type CreateUpstream struct {
	// Which load balancing algorithm to use.
	Algorithm *CreateUpstreamAlgorithm `default:"round-robin" json:"algorithm"`
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *CreateUpstreamClientCertificate `json:"client_certificate,omitempty"`
	// What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.
	HashFallback *HashFallback `default:"none" json:"hash_fallback"`
	// The header name to take the value from as hash input. Only required when `hash_fallback` is set to `header`.
	HashFallbackHeader *string `json:"hash_fallback_header,omitempty"`
	// The name of the query string argument to take the value from as hash input. Only required when `hash_fallback` is set to `query_arg`.
	HashFallbackQueryArg *string `json:"hash_fallback_query_arg,omitempty"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hash_fallback` is set to `uri_capture`.
	HashFallbackURICapture *string `json:"hash_fallback_uri_capture,omitempty"`
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.
	HashOn *HashOn `default:"none" json:"hash_on"`
	// The cookie name to take the value from as hash input. Only required when `hash_on` or `hash_fallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie *string `json:"hash_on_cookie,omitempty"`
	// The cookie path to set in the response headers. Only required when `hash_on` or `hash_fallback` is set to `cookie`.
	HashOnCookiePath *string `default:"/" json:"hash_on_cookie_path"`
	// The header name to take the value from as hash input. Only required when `hash_on` is set to `header`.
	HashOnHeader *string `json:"hash_on_header,omitempty"`
	// The name of the query string argument to take the value from as hash input. Only required when `hash_on` is set to `query_arg`.
	HashOnQueryArg *string `json:"hash_on_query_arg,omitempty"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hash_on` is set to `uri_capture`.
	HashOnURICapture *string       `json:"hash_on_uri_capture,omitempty"`
	Healthchecks     *Healthchecks `json:"healthchecks,omitempty"`
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader *string `json:"host_header,omitempty"`
	// This is a hostname, which must be equal to the `host` of a Service.
	Name string `json:"name"`
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots *int64 `default:"10000" json:"slots"`
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName *bool `default:"false" json:"use_srv_name"`
}

func (c CreateUpstream) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateUpstream) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateUpstream) GetAlgorithm() *CreateUpstreamAlgorithm {
	if o == nil {
		return nil
	}
	return o.Algorithm
}

func (o *CreateUpstream) GetClientCertificate() *CreateUpstreamClientCertificate {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *CreateUpstream) GetHashFallback() *HashFallback {
	if o == nil {
		return nil
	}
	return o.HashFallback
}

func (o *CreateUpstream) GetHashFallbackHeader() *string {
	if o == nil {
		return nil
	}
	return o.HashFallbackHeader
}

func (o *CreateUpstream) GetHashFallbackQueryArg() *string {
	if o == nil {
		return nil
	}
	return o.HashFallbackQueryArg
}

func (o *CreateUpstream) GetHashFallbackURICapture() *string {
	if o == nil {
		return nil
	}
	return o.HashFallbackURICapture
}

func (o *CreateUpstream) GetHashOn() *HashOn {
	if o == nil {
		return nil
	}
	return o.HashOn
}

func (o *CreateUpstream) GetHashOnCookie() *string {
	if o == nil {
		return nil
	}
	return o.HashOnCookie
}

func (o *CreateUpstream) GetHashOnCookiePath() *string {
	if o == nil {
		return nil
	}
	return o.HashOnCookiePath
}

func (o *CreateUpstream) GetHashOnHeader() *string {
	if o == nil {
		return nil
	}
	return o.HashOnHeader
}

func (o *CreateUpstream) GetHashOnQueryArg() *string {
	if o == nil {
		return nil
	}
	return o.HashOnQueryArg
}

func (o *CreateUpstream) GetHashOnURICapture() *string {
	if o == nil {
		return nil
	}
	return o.HashOnURICapture
}

func (o *CreateUpstream) GetHealthchecks() *Healthchecks {
	if o == nil {
		return nil
	}
	return o.Healthchecks
}

func (o *CreateUpstream) GetHostHeader() *string {
	if o == nil {
		return nil
	}
	return o.HostHeader
}

func (o *CreateUpstream) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateUpstream) GetSlots() *int64 {
	if o == nil {
		return nil
	}
	return o.Slots
}

func (o *CreateUpstream) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateUpstream) GetUseSrvName() *bool {
	if o == nil {
		return nil
	}
	return o.UseSrvName
}
