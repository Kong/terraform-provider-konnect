// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// MeshTCPRouteItemType - the type of the resource
type MeshTCPRouteItemType string

const (
	MeshTCPRouteItemTypeMeshTCPRoute MeshTCPRouteItemType = "MeshTCPRoute"
)

func (e MeshTCPRouteItemType) ToPointer() *MeshTCPRouteItemType {
	return &e
}
func (e *MeshTCPRouteItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshTCPRoute":
		*e = MeshTCPRouteItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTCPRouteItemType: %v", v)
	}
}

// MeshTCPRouteItemKind - Kind of the referenced resource
type MeshTCPRouteItemKind string

const (
	MeshTCPRouteItemKindMesh                 MeshTCPRouteItemKind = "Mesh"
	MeshTCPRouteItemKindMeshSubset           MeshTCPRouteItemKind = "MeshSubset"
	MeshTCPRouteItemKindMeshGateway          MeshTCPRouteItemKind = "MeshGateway"
	MeshTCPRouteItemKindMeshService          MeshTCPRouteItemKind = "MeshService"
	MeshTCPRouteItemKindMeshExternalService  MeshTCPRouteItemKind = "MeshExternalService"
	MeshTCPRouteItemKindMeshMultiZoneService MeshTCPRouteItemKind = "MeshMultiZoneService"
	MeshTCPRouteItemKindMeshServiceSubset    MeshTCPRouteItemKind = "MeshServiceSubset"
	MeshTCPRouteItemKindMeshHTTPRoute        MeshTCPRouteItemKind = "MeshHTTPRoute"
	MeshTCPRouteItemKindDataplane            MeshTCPRouteItemKind = "Dataplane"
)

func (e MeshTCPRouteItemKind) ToPointer() *MeshTCPRouteItemKind {
	return &e
}
func (e *MeshTCPRouteItemKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshTCPRouteItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTCPRouteItemKind: %v", v)
	}
}

type MeshTCPRouteItemProxyTypes string

const (
	MeshTCPRouteItemProxyTypesSidecar MeshTCPRouteItemProxyTypes = "Sidecar"
	MeshTCPRouteItemProxyTypesGateway MeshTCPRouteItemProxyTypes = "Gateway"
)

func (e MeshTCPRouteItemProxyTypes) ToPointer() *MeshTCPRouteItemProxyTypes {
	return &e
}
func (e *MeshTCPRouteItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTCPRouteItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTCPRouteItemProxyTypes: %v", v)
	}
}

// MeshTCPRouteItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined in-place.
type MeshTCPRouteItemTargetRef struct {
	// Kind of the referenced resource
	Kind *MeshTCPRouteItemKind `json:"kind,omitempty"`
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
	ProxyTypes []MeshTCPRouteItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshTCPRouteItemTargetRef) GetKind() *MeshTCPRouteItemKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *MeshTCPRouteItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTCPRouteItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTCPRouteItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTCPRouteItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTCPRouteItemTargetRef) GetProxyTypes() []MeshTCPRouteItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTCPRouteItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTCPRouteItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshTCPRouteItemSpecToKind - Kind of the referenced resource
type MeshTCPRouteItemSpecToKind string

const (
	MeshTCPRouteItemSpecToKindMesh                 MeshTCPRouteItemSpecToKind = "Mesh"
	MeshTCPRouteItemSpecToKindMeshSubset           MeshTCPRouteItemSpecToKind = "MeshSubset"
	MeshTCPRouteItemSpecToKindMeshGateway          MeshTCPRouteItemSpecToKind = "MeshGateway"
	MeshTCPRouteItemSpecToKindMeshService          MeshTCPRouteItemSpecToKind = "MeshService"
	MeshTCPRouteItemSpecToKindMeshExternalService  MeshTCPRouteItemSpecToKind = "MeshExternalService"
	MeshTCPRouteItemSpecToKindMeshMultiZoneService MeshTCPRouteItemSpecToKind = "MeshMultiZoneService"
	MeshTCPRouteItemSpecToKindMeshServiceSubset    MeshTCPRouteItemSpecToKind = "MeshServiceSubset"
	MeshTCPRouteItemSpecToKindMeshHTTPRoute        MeshTCPRouteItemSpecToKind = "MeshHTTPRoute"
	MeshTCPRouteItemSpecToKindDataplane            MeshTCPRouteItemSpecToKind = "Dataplane"
)

func (e MeshTCPRouteItemSpecToKind) ToPointer() *MeshTCPRouteItemSpecToKind {
	return &e
}
func (e *MeshTCPRouteItemSpecToKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshTCPRouteItemSpecToKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTCPRouteItemSpecToKind: %v", v)
	}
}

type MeshTCPRouteItemSpecToProxyTypes string

const (
	MeshTCPRouteItemSpecToProxyTypesSidecar MeshTCPRouteItemSpecToProxyTypes = "Sidecar"
	MeshTCPRouteItemSpecToProxyTypesGateway MeshTCPRouteItemSpecToProxyTypes = "Gateway"
)

func (e MeshTCPRouteItemSpecToProxyTypes) ToPointer() *MeshTCPRouteItemSpecToProxyTypes {
	return &e
}
func (e *MeshTCPRouteItemSpecToProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTCPRouteItemSpecToProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTCPRouteItemSpecToProxyTypes: %v", v)
	}
}

// MeshTCPRouteItemBackendRefs - BackendRef defines where to forward traffic.
type MeshTCPRouteItemBackendRefs struct {
	// Kind of the referenced resource
	Kind *MeshTCPRouteItemSpecToKind `json:"kind,omitempty"`
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
	// Port is only supported when this ref refers to a real MeshService object
	Port *int `json:"port,omitempty"`
	// ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
	// all data plane types are targeted by the policy.
	ProxyTypes []MeshTCPRouteItemSpecToProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags   map[string]string `json:"tags,omitempty"`
	Weight *int64            `json:"weight,omitempty"`
}

func (o *MeshTCPRouteItemBackendRefs) GetKind() *MeshTCPRouteItemSpecToKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *MeshTCPRouteItemBackendRefs) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTCPRouteItemBackendRefs) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTCPRouteItemBackendRefs) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTCPRouteItemBackendRefs) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTCPRouteItemBackendRefs) GetPort() *int {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *MeshTCPRouteItemBackendRefs) GetProxyTypes() []MeshTCPRouteItemSpecToProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTCPRouteItemBackendRefs) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTCPRouteItemBackendRefs) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *MeshTCPRouteItemBackendRefs) GetWeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Weight
}

// MeshTCPRouteItemDefault - Default holds routing rules that can be merged with rules from other
// policies.
type MeshTCPRouteItemDefault struct {
	BackendRefs []MeshTCPRouteItemBackendRefs `json:"backendRefs"`
}

func (o *MeshTCPRouteItemDefault) GetBackendRefs() []MeshTCPRouteItemBackendRefs {
	if o == nil {
		return []MeshTCPRouteItemBackendRefs{}
	}
	return o.BackendRefs
}

type MeshTCPRouteItemRules struct {
	// Default holds routing rules that can be merged with rules from other
	// policies.
	Default MeshTCPRouteItemDefault `json:"default"`
}

func (o *MeshTCPRouteItemRules) GetDefault() MeshTCPRouteItemDefault {
	if o == nil {
		return MeshTCPRouteItemDefault{}
	}
	return o.Default
}

// MeshTCPRouteItemSpecKind - Kind of the referenced resource
type MeshTCPRouteItemSpecKind string

const (
	MeshTCPRouteItemSpecKindMesh                 MeshTCPRouteItemSpecKind = "Mesh"
	MeshTCPRouteItemSpecKindMeshSubset           MeshTCPRouteItemSpecKind = "MeshSubset"
	MeshTCPRouteItemSpecKindMeshGateway          MeshTCPRouteItemSpecKind = "MeshGateway"
	MeshTCPRouteItemSpecKindMeshService          MeshTCPRouteItemSpecKind = "MeshService"
	MeshTCPRouteItemSpecKindMeshExternalService  MeshTCPRouteItemSpecKind = "MeshExternalService"
	MeshTCPRouteItemSpecKindMeshMultiZoneService MeshTCPRouteItemSpecKind = "MeshMultiZoneService"
	MeshTCPRouteItemSpecKindMeshServiceSubset    MeshTCPRouteItemSpecKind = "MeshServiceSubset"
	MeshTCPRouteItemSpecKindMeshHTTPRoute        MeshTCPRouteItemSpecKind = "MeshHTTPRoute"
	MeshTCPRouteItemSpecKindDataplane            MeshTCPRouteItemSpecKind = "Dataplane"
)

func (e MeshTCPRouteItemSpecKind) ToPointer() *MeshTCPRouteItemSpecKind {
	return &e
}
func (e *MeshTCPRouteItemSpecKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshTCPRouteItemSpecKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTCPRouteItemSpecKind: %v", v)
	}
}

type MeshTCPRouteItemSpecProxyTypes string

const (
	MeshTCPRouteItemSpecProxyTypesSidecar MeshTCPRouteItemSpecProxyTypes = "Sidecar"
	MeshTCPRouteItemSpecProxyTypesGateway MeshTCPRouteItemSpecProxyTypes = "Gateway"
)

func (e MeshTCPRouteItemSpecProxyTypes) ToPointer() *MeshTCPRouteItemSpecProxyTypes {
	return &e
}
func (e *MeshTCPRouteItemSpecProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTCPRouteItemSpecProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTCPRouteItemSpecProxyTypes: %v", v)
	}
}

// MeshTCPRouteItemSpecTargetRef - TargetRef is a reference to the resource that represents a group of
// destinations.
type MeshTCPRouteItemSpecTargetRef struct {
	// Kind of the referenced resource
	Kind *MeshTCPRouteItemSpecKind `json:"kind,omitempty"`
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
	ProxyTypes []MeshTCPRouteItemSpecProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshTCPRouteItemSpecTargetRef) GetKind() *MeshTCPRouteItemSpecKind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *MeshTCPRouteItemSpecTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTCPRouteItemSpecTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTCPRouteItemSpecTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTCPRouteItemSpecTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTCPRouteItemSpecTargetRef) GetProxyTypes() []MeshTCPRouteItemSpecProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTCPRouteItemSpecTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTCPRouteItemSpecTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshTCPRouteItemTo struct {
	// Rules contains the routing rules applies to a combination of top-level
	// targetRef and the targetRef in this entry.
	Rules []MeshTCPRouteItemRules `json:"rules,omitempty"`
	// TargetRef is a reference to the resource that represents a group of
	// destinations.
	TargetRef MeshTCPRouteItemSpecTargetRef `json:"targetRef"`
}

func (o *MeshTCPRouteItemTo) GetRules() []MeshTCPRouteItemRules {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *MeshTCPRouteItemTo) GetTargetRef() MeshTCPRouteItemSpecTargetRef {
	if o == nil {
		return MeshTCPRouteItemSpecTargetRef{}
	}
	return o.TargetRef
}

// MeshTCPRouteItemSpec - Spec is the specification of the Kuma MeshTCPRoute resource.
type MeshTCPRouteItemSpec struct {
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined in-place.
	TargetRef *MeshTCPRouteItemTargetRef `json:"targetRef,omitempty"`
	// To list makes a match between the consumed services and corresponding
	// configurations
	To []MeshTCPRouteItemTo `json:"to,omitempty"`
}

func (o *MeshTCPRouteItemSpec) GetTargetRef() *MeshTCPRouteItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

func (o *MeshTCPRouteItemSpec) GetTo() []MeshTCPRouteItemTo {
	if o == nil {
		return nil
	}
	return o.To
}

type MeshTCPRouteItem struct {
	// the type of the resource
	Type MeshTCPRouteItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshTCPRoute resource.
	Spec MeshTCPRouteItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshTCPRouteItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTCPRouteItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTCPRouteItem) GetType() MeshTCPRouteItemType {
	if o == nil {
		return MeshTCPRouteItemType("")
	}
	return o.Type
}

func (o *MeshTCPRouteItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTCPRouteItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshTCPRouteItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTCPRouteItem) GetSpec() MeshTCPRouteItemSpec {
	if o == nil {
		return MeshTCPRouteItemSpec{}
	}
	return o.Spec
}

func (o *MeshTCPRouteItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshTCPRouteItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}