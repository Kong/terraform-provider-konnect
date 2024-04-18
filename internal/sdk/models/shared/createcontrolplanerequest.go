// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// The ClusterType value of the cluster associated with the Control Plane.
type ClusterType string

const (
	ClusterTypeClusterTypeControlPlane         ClusterType = "CLUSTER_TYPE_CONTROL_PLANE"
	ClusterTypeClusterTypeHybrid               ClusterType = "CLUSTER_TYPE_HYBRID"
	ClusterTypeClusterTypeK8SIngressController ClusterType = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
	ClusterTypeClusterTypeControlPlaneGroup    ClusterType = "CLUSTER_TYPE_CONTROL_PLANE_GROUP"
)

func (e ClusterType) ToPointer() *ClusterType {
	return &e
}

func (e *ClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CLUSTER_TYPE_CONTROL_PLANE":
		fallthrough
	case "CLUSTER_TYPE_HYBRID":
		fallthrough
	case "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER":
		fallthrough
	case "CLUSTER_TYPE_CONTROL_PLANE_GROUP":
		*e = ClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClusterType: %v", v)
	}
}

// AuthType - The auth type value of the cluster associated with the Runtime Group.
type AuthType string

const (
	AuthTypePinnedClientCerts AuthType = "pinned_client_certs"
	AuthTypePkiClientCerts    AuthType = "pki_client_certs"
)

func (e AuthType) ToPointer() *AuthType {
	return &e
}

func (e *AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinned_client_certs":
		fallthrough
	case "pki_client_certs":
		*e = AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthType: %v", v)
	}
}

// CreateControlPlaneRequest - The request schema for the create control plane request.
type CreateControlPlaneRequest struct {
	// The name of the control plane.
	Name string `json:"name"`
	// The description of the control plane in Konnect.
	Description *string `json:"description,omitempty"`
	// The ClusterType value of the cluster associated with the Control Plane.
	ClusterType *ClusterType `json:"cluster_type,omitempty"`
	// The auth type value of the cluster associated with the Runtime Group.
	AuthType *AuthType `json:"auth_type,omitempty"`
	// Whether this control-plane can be used for cloud-gateways.
	CloudGateway *bool `json:"cloud_gateway,omitempty"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls []ProxyURL `json:"proxy_urls,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
}

func (o *CreateControlPlaneRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateControlPlaneRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateControlPlaneRequest) GetClusterType() *ClusterType {
	if o == nil {
		return nil
	}
	return o.ClusterType
}

func (o *CreateControlPlaneRequest) GetAuthType() *AuthType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *CreateControlPlaneRequest) GetCloudGateway() *bool {
	if o == nil {
		return nil
	}
	return o.CloudGateway
}

func (o *CreateControlPlaneRequest) GetProxyUrls() []ProxyURL {
	if o == nil {
		return nil
	}
	return o.ProxyUrls
}

func (o *CreateControlPlaneRequest) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}
