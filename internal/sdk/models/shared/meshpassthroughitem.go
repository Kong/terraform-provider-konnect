// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// MeshPassthroughItemType - the type of the resource
type MeshPassthroughItemType string

const (
	MeshPassthroughItemTypeMeshPassthrough MeshPassthroughItemType = "MeshPassthrough"
)

func (e MeshPassthroughItemType) ToPointer() *MeshPassthroughItemType {
	return &e
}
func (e *MeshPassthroughItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshPassthrough":
		*e = MeshPassthroughItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshPassthroughItemType: %v", v)
	}
}

// MeshPassthroughItemProtocol - Protocol defines the communication protocol. Possible values: `tcp`, `tls`, `grpc`, `http`, `http2`.
type MeshPassthroughItemProtocol string

const (
	MeshPassthroughItemProtocolTCP   MeshPassthroughItemProtocol = "tcp"
	MeshPassthroughItemProtocolTLS   MeshPassthroughItemProtocol = "tls"
	MeshPassthroughItemProtocolGrpc  MeshPassthroughItemProtocol = "grpc"
	MeshPassthroughItemProtocolHTTP  MeshPassthroughItemProtocol = "http"
	MeshPassthroughItemProtocolHttp2 MeshPassthroughItemProtocol = "http2"
)

func (e MeshPassthroughItemProtocol) ToPointer() *MeshPassthroughItemProtocol {
	return &e
}
func (e *MeshPassthroughItemProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "grpc":
		fallthrough
	case "http":
		fallthrough
	case "http2":
		*e = MeshPassthroughItemProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshPassthroughItemProtocol: %v", v)
	}
}

// MeshPassthroughItemSpecType - Type of the match, one of `Domain`, `IP` or `CIDR` is available.
type MeshPassthroughItemSpecType string

const (
	MeshPassthroughItemSpecTypeDomain MeshPassthroughItemSpecType = "Domain"
	MeshPassthroughItemSpecTypeIP     MeshPassthroughItemSpecType = "IP"
	MeshPassthroughItemSpecTypeCidr   MeshPassthroughItemSpecType = "CIDR"
)

func (e MeshPassthroughItemSpecType) ToPointer() *MeshPassthroughItemSpecType {
	return &e
}
func (e *MeshPassthroughItemSpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Domain":
		fallthrough
	case "IP":
		fallthrough
	case "CIDR":
		*e = MeshPassthroughItemSpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshPassthroughItemSpecType: %v", v)
	}
}

type AppendMatch struct {
	// Port defines the port to which a user makes a request.
	Port *int `json:"port,omitempty"`
	// Protocol defines the communication protocol. Possible values: `tcp`, `tls`, `grpc`, `http`, `http2`.
	Protocol *MeshPassthroughItemProtocol `json:"protocol,omitempty"`
	// Type of the match, one of `Domain`, `IP` or `CIDR` is available.
	Type *MeshPassthroughItemSpecType `json:"type,omitempty"`
	// Value for the specified Type.
	Value *string `json:"value,omitempty"`
}

func (o *AppendMatch) GetPort() *int {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *AppendMatch) GetProtocol() *MeshPassthroughItemProtocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *AppendMatch) GetType() *MeshPassthroughItemSpecType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AppendMatch) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// PassthroughMode - Defines the passthrough behavior. Possible values: `All`, `None`, `Matched`
// When `All` or `None` `appendMatch` has no effect.
type PassthroughMode string

const (
	PassthroughModeAll     PassthroughMode = "All"
	PassthroughModeMatched PassthroughMode = "Matched"
	PassthroughModeNone    PassthroughMode = "None"
)

func (e PassthroughMode) ToPointer() *PassthroughMode {
	return &e
}
func (e *PassthroughMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "All":
		fallthrough
	case "Matched":
		fallthrough
	case "None":
		*e = PassthroughMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PassthroughMode: %v", v)
	}
}

// MeshPassthroughItemDefault - MeshPassthrough configuration.
type MeshPassthroughItemDefault struct {
	// AppendMatch is a list of destinations that should be allowed through the sidecar.
	AppendMatch []AppendMatch `json:"appendMatch,omitempty"`
	// Defines the passthrough behavior. Possible values: `All`, `None`, `Matched`
	// When `All` or `None` `appendMatch` has no effect.
	PassthroughMode *PassthroughMode `json:"passthroughMode,omitempty"`
}

func (o *MeshPassthroughItemDefault) GetAppendMatch() []AppendMatch {
	if o == nil {
		return nil
	}
	return o.AppendMatch
}

func (o *MeshPassthroughItemDefault) GetPassthroughMode() *PassthroughMode {
	if o == nil {
		return nil
	}
	return o.PassthroughMode
}

// MeshPassthroughItemKind - Kind of the referenced resource
type MeshPassthroughItemKind string

const (
	MeshPassthroughItemKindMesh                 MeshPassthroughItemKind = "Mesh"
	MeshPassthroughItemKindMeshSubset           MeshPassthroughItemKind = "MeshSubset"
	MeshPassthroughItemKindMeshGateway          MeshPassthroughItemKind = "MeshGateway"
	MeshPassthroughItemKindMeshService          MeshPassthroughItemKind = "MeshService"
	MeshPassthroughItemKindMeshExternalService  MeshPassthroughItemKind = "MeshExternalService"
	MeshPassthroughItemKindMeshMultiZoneService MeshPassthroughItemKind = "MeshMultiZoneService"
	MeshPassthroughItemKindMeshServiceSubset    MeshPassthroughItemKind = "MeshServiceSubset"
	MeshPassthroughItemKindMeshHTTPRoute        MeshPassthroughItemKind = "MeshHTTPRoute"
	MeshPassthroughItemKindDataplane            MeshPassthroughItemKind = "Dataplane"
)

func (e MeshPassthroughItemKind) ToPointer() *MeshPassthroughItemKind {
	return &e
}
func (e *MeshPassthroughItemKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshPassthroughItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshPassthroughItemKind: %v", v)
	}
}

type MeshPassthroughItemProxyTypes string

const (
	MeshPassthroughItemProxyTypesSidecar MeshPassthroughItemProxyTypes = "Sidecar"
	MeshPassthroughItemProxyTypesGateway MeshPassthroughItemProxyTypes = "Gateway"
)

func (e MeshPassthroughItemProxyTypes) ToPointer() *MeshPassthroughItemProxyTypes {
	return &e
}
func (e *MeshPassthroughItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshPassthroughItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshPassthroughItemProxyTypes: %v", v)
	}
}

// MeshPassthroughItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined in-place.
type MeshPassthroughItemTargetRef struct {
	// Kind of the referenced resource
	Kind *MeshPassthroughItemKind `json:"kind,omitempty"`
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
	ProxyTypes []MeshPassthroughItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshPassthroughItemTargetRef) GetKind() *MeshPassthroughItemKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *MeshPassthroughItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshPassthroughItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshPassthroughItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshPassthroughItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshPassthroughItemTargetRef) GetProxyTypes() []MeshPassthroughItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshPassthroughItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshPassthroughItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshPassthroughItemSpec - Spec is the specification of the Kuma MeshPassthrough resource.
type MeshPassthroughItemSpec struct {
	// MeshPassthrough configuration.
	Default *MeshPassthroughItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined in-place.
	TargetRef *MeshPassthroughItemTargetRef `json:"targetRef,omitempty"`
}

func (o *MeshPassthroughItemSpec) GetDefault() *MeshPassthroughItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshPassthroughItemSpec) GetTargetRef() *MeshPassthroughItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

type MeshPassthroughItem struct {
	// the type of the resource
	Type MeshPassthroughItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshPassthrough resource.
	Spec MeshPassthroughItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshPassthroughItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshPassthroughItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshPassthroughItem) GetType() MeshPassthroughItemType {
	if o == nil {
		return MeshPassthroughItemType("")
	}
	return o.Type
}

func (o *MeshPassthroughItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshPassthroughItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshPassthroughItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshPassthroughItem) GetSpec() MeshPassthroughItemSpec {
	if o == nil {
		return MeshPassthroughItemSpec{}
	}
	return o.Spec
}

func (o *MeshPassthroughItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshPassthroughItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}