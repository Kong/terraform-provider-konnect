// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// MeshOPAItemType - the type of the resource
type MeshOPAItemType string

const (
	MeshOPAItemTypeMeshOpa MeshOPAItemType = "MeshOPA"
)

func (e MeshOPAItemType) ToPointer() *MeshOPAItemType {
	return &e
}
func (e *MeshOPAItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshOPA":
		*e = MeshOPAItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshOPAItemType: %v", v)
	}
}

// AgentConfig defines bootstrap OPA agent configuration.
type AgentConfig struct {
	// Data source is inline bytes.
	Inline *string `json:"inline,omitempty"`
	// Data source is inline string`
	InlineString *string `json:"inlineString,omitempty"`
	// Data source is a secret with given Secret key.
	Secret *string `json:"secret,omitempty"`
}

func (o *AgentConfig) GetInline() *string {
	if o == nil {
		return nil
	}
	return o.Inline
}

func (o *AgentConfig) GetInlineString() *string {
	if o == nil {
		return nil
	}
	return o.InlineString
}

func (o *AgentConfig) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

// Rego - OPA Policy written in Rego. Available values: secret, inline, inlineString.
type Rego struct {
	// Data source is inline bytes.
	Inline *string `json:"inline,omitempty"`
	// Data source is inline string`
	InlineString *string `json:"inlineString,omitempty"`
	// Data source is a secret with given Secret key.
	Secret *string `json:"secret,omitempty"`
}

func (o *Rego) GetInline() *string {
	if o == nil {
		return nil
	}
	return o.Inline
}

func (o *Rego) GetInlineString() *string {
	if o == nil {
		return nil
	}
	return o.InlineString
}

func (o *Rego) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

type AppendPolicies struct {
	// If true, then policy won't be taken into account when making a decision.
	IgnoreDecision *bool `json:"ignoreDecision,omitempty"`
	// OPA Policy written in Rego. Available values: secret, inline, inlineString.
	Rego Rego `json:"rego"`
}

func (o *AppendPolicies) GetIgnoreDecision() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreDecision
}

func (o *AppendPolicies) GetRego() Rego {
	if o == nil {
		return Rego{}
	}
	return o.Rego
}

// OnAgentFailure either 'allow' or 'deny' (default to deny) whether
// to allow requests when the authorization agent failed.
type OnAgentFailure string

const (
	OnAgentFailureAllow OnAgentFailure = "Allow"
	OnAgentFailureDeny  OnAgentFailure = "Deny"
)

func (e OnAgentFailure) ToPointer() *OnAgentFailure {
	return &e
}
func (e *OnAgentFailure) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Allow":
		fallthrough
	case "Deny":
		*e = OnAgentFailure(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OnAgentFailure: %v", v)
	}
}

// RequestBody configuration to apply on the request body sent to the
// authorization agent (if absent, the body is not sent).
type RequestBody struct {
	// MaxSize defines the maximum payload size sent to authorization agent. If the payload
	// is larger it will be truncated and there will be a header
	// `x-envoy-auth-partial-body: true`. If it is set to 0 no body will be
	// sent to the agent.
	MaxSize *int `json:"maxSize,omitempty"`
	// SendRawBody enable sending raw body instead of the body encoded into UTF-8
	SendRawBody *bool `json:"sendRawBody,omitempty"`
}

func (o *RequestBody) GetMaxSize() *int {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *RequestBody) GetSendRawBody() *bool {
	if o == nil {
		return nil
	}
	return o.SendRawBody
}

// AuthConfig are configurations specific to the filter.
type AuthConfig struct {
	// OnAgentFailure either 'allow' or 'deny' (default to deny) whether
	// to allow requests when the authorization agent failed.
	OnAgentFailure *OnAgentFailure `json:"onAgentFailure,omitempty"`
	// RequestBody configuration to apply on the request body sent to the
	// authorization agent (if absent, the body is not sent).
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	// StatusOnError is the http status to return when there's a connection
	// failure between the dataplane and the authorization agent
	StatusOnError *int `json:"statusOnError,omitempty"`
	// Timeout for the single gRPC request from Envoy to OPA Agent.
	Timeout *string `json:"timeout,omitempty"`
}

func (o *AuthConfig) GetOnAgentFailure() *OnAgentFailure {
	if o == nil {
		return nil
	}
	return o.OnAgentFailure
}

func (o *AuthConfig) GetRequestBody() *RequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AuthConfig) GetStatusOnError() *int {
	if o == nil {
		return nil
	}
	return o.StatusOnError
}

func (o *AuthConfig) GetTimeout() *string {
	if o == nil {
		return nil
	}
	return o.Timeout
}

type MeshOPAItemDefault struct {
	// AgentConfig defines bootstrap OPA agent configuration.
	AgentConfig *AgentConfig `json:"agentConfig,omitempty"`
	// Policies define OPA policies that will be applied on OPA Agent.
	AppendPolicies []AppendPolicies `json:"appendPolicies,omitempty"`
	// AuthConfig are configurations specific to the filter.
	AuthConfig *AuthConfig `json:"authConfig,omitempty"`
}

func (o *MeshOPAItemDefault) GetAgentConfig() *AgentConfig {
	if o == nil {
		return nil
	}
	return o.AgentConfig
}

func (o *MeshOPAItemDefault) GetAppendPolicies() []AppendPolicies {
	if o == nil {
		return nil
	}
	return o.AppendPolicies
}

func (o *MeshOPAItemDefault) GetAuthConfig() *AuthConfig {
	if o == nil {
		return nil
	}
	return o.AuthConfig
}

// MeshOPAItemKind - Kind of the referenced resource
type MeshOPAItemKind string

const (
	MeshOPAItemKindMesh                 MeshOPAItemKind = "Mesh"
	MeshOPAItemKindMeshSubset           MeshOPAItemKind = "MeshSubset"
	MeshOPAItemKindMeshGateway          MeshOPAItemKind = "MeshGateway"
	MeshOPAItemKindMeshService          MeshOPAItemKind = "MeshService"
	MeshOPAItemKindMeshExternalService  MeshOPAItemKind = "MeshExternalService"
	MeshOPAItemKindMeshMultiZoneService MeshOPAItemKind = "MeshMultiZoneService"
	MeshOPAItemKindMeshServiceSubset    MeshOPAItemKind = "MeshServiceSubset"
	MeshOPAItemKindMeshHTTPRoute        MeshOPAItemKind = "MeshHTTPRoute"
	MeshOPAItemKindDataplane            MeshOPAItemKind = "Dataplane"
)

func (e MeshOPAItemKind) ToPointer() *MeshOPAItemKind {
	return &e
}
func (e *MeshOPAItemKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshOPAItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshOPAItemKind: %v", v)
	}
}

type MeshOPAItemProxyTypes string

const (
	MeshOPAItemProxyTypesSidecar MeshOPAItemProxyTypes = "Sidecar"
	MeshOPAItemProxyTypesGateway MeshOPAItemProxyTypes = "Gateway"
)

func (e MeshOPAItemProxyTypes) ToPointer() *MeshOPAItemProxyTypes {
	return &e
}
func (e *MeshOPAItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshOPAItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshOPAItemProxyTypes: %v", v)
	}
}

// MeshOPAItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined inplace.
type MeshOPAItemTargetRef struct {
	// Kind of the referenced resource
	Kind *MeshOPAItemKind `json:"kind,omitempty"`
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
	ProxyTypes []MeshOPAItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshOPAItemTargetRef) GetKind() *MeshOPAItemKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *MeshOPAItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshOPAItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshOPAItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshOPAItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshOPAItemTargetRef) GetProxyTypes() []MeshOPAItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshOPAItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshOPAItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshOPAItemSpec - Spec is the specification of the Kuma MeshOPA resource.
type MeshOPAItemSpec struct {
	Default *MeshOPAItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *MeshOPAItemTargetRef `json:"targetRef,omitempty"`
}

func (o *MeshOPAItemSpec) GetDefault() *MeshOPAItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshOPAItemSpec) GetTargetRef() *MeshOPAItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

type MeshOPAItem struct {
	// the type of the resource
	Type MeshOPAItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshOPA resource.
	Spec MeshOPAItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshOPAItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshOPAItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshOPAItem) GetType() MeshOPAItemType {
	if o == nil {
		return MeshOPAItemType("")
	}
	return o.Type
}

func (o *MeshOPAItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshOPAItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshOPAItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshOPAItem) GetSpec() MeshOPAItemSpec {
	if o == nil {
		return MeshOPAItemSpec{}
	}
	return o.Spec
}

func (o *MeshOPAItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshOPAItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}