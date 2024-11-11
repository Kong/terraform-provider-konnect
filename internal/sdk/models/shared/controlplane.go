// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// ControlPlaneClusterType - The ClusterType value of the cluster associated with the Control Plane.
type ControlPlaneClusterType string

const (
	ControlPlaneClusterTypeClusterTypeControlPlane         ControlPlaneClusterType = "CLUSTER_TYPE_CONTROL_PLANE"
	ControlPlaneClusterTypeClusterTypeK8SIngressController ControlPlaneClusterType = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
	ControlPlaneClusterTypeClusterTypeControlPlaneGroup    ControlPlaneClusterType = "CLUSTER_TYPE_CONTROL_PLANE_GROUP"
	ControlPlaneClusterTypeClusterTypeServerless           ControlPlaneClusterType = "CLUSTER_TYPE_SERVERLESS"
	ControlPlaneClusterTypeClusterTypeHybrid               ControlPlaneClusterType = "CLUSTER_TYPE_HYBRID"
)

func (e ControlPlaneClusterType) ToPointer() *ControlPlaneClusterType {
	return &e
}
func (e *ControlPlaneClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CLUSTER_TYPE_CONTROL_PLANE":
		fallthrough
	case "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER":
		fallthrough
	case "CLUSTER_TYPE_CONTROL_PLANE_GROUP":
		fallthrough
	case "CLUSTER_TYPE_SERVERLESS":
		fallthrough
	case "CLUSTER_TYPE_HYBRID":
		*e = ControlPlaneClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ControlPlaneClusterType: %v", v)
	}
}

// ControlPlaneAuthType - The auth type value of the cluster associated with the Runtime Group.
type ControlPlaneAuthType string

const (
	ControlPlaneAuthTypePinnedClientCerts ControlPlaneAuthType = "pinned_client_certs"
	ControlPlaneAuthTypePkiClientCerts    ControlPlaneAuthType = "pki_client_certs"
)

func (e ControlPlaneAuthType) ToPointer() *ControlPlaneAuthType {
	return &e
}
func (e *ControlPlaneAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinned_client_certs":
		fallthrough
	case "pki_client_certs":
		*e = ControlPlaneAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ControlPlaneAuthType: %v", v)
	}
}

// Config - CP configuration object for related access endpoints.
type Config struct {
	// Control Plane Endpoint.
	ControlPlaneEndpoint string `json:"control_plane_endpoint"`
	// Telemetry Endpoint.
	TelemetryEndpoint string `json:"telemetry_endpoint"`
	// The ClusterType value of the cluster associated with the Control Plane.
	ClusterType ControlPlaneClusterType `json:"cluster_type"`
	// The auth type value of the cluster associated with the Runtime Group.
	AuthType ControlPlaneAuthType `json:"auth_type"`
	// Whether the Control Plane can be used for cloud-gateways.
	CloudGateway bool `json:"cloud_gateway"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls []ProxyURL `json:"proxy_urls,omitempty"`
}

func (o *Config) GetControlPlaneEndpoint() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneEndpoint
}

func (o *Config) GetTelemetryEndpoint() string {
	if o == nil {
		return ""
	}
	return o.TelemetryEndpoint
}

func (o *Config) GetClusterType() ControlPlaneClusterType {
	if o == nil {
		return ControlPlaneClusterType("")
	}
	return o.ClusterType
}

func (o *Config) GetAuthType() ControlPlaneAuthType {
	if o == nil {
		return ControlPlaneAuthType("")
	}
	return o.AuthType
}

func (o *Config) GetCloudGateway() bool {
	if o == nil {
		return false
	}
	return o.CloudGateway
}

func (o *Config) GetProxyUrls() []ProxyURL {
	if o == nil {
		return nil
	}
	return o.ProxyUrls
}

// ControlPlane - The control plane object contains information about a Kong control plane.
type ControlPlane struct {
	// The control plane ID.
	ID string `json:"id"`
	// The name of the control plane.
	Name string `json:"name"`
	// The description of the control plane in Konnect.
	Description *string `json:"description,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
	// CP configuration object for related access endpoints.
	Config Config `json:"config"`
	// An ISO-8604 timestamp representation of control plane creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8604 timestamp representation of control plane update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (c ControlPlane) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ControlPlane) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ControlPlane) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ControlPlane) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ControlPlane) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ControlPlane) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ControlPlane) GetConfig() Config {
	if o == nil {
		return Config{}
	}
	return o.Config
}

func (o *ControlPlane) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ControlPlane) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
