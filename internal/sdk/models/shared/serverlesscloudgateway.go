// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"time"
)

// ServerlessCloudGateway - A serverless cloud gateway
type ServerlessCloudGateway struct {
	ControlPlane ServerlessControlPlane `json:"control_plane"`
	CreatedAt    time.Time              `json:"created_at"`
	// Endpoint for the serverless cloud gateway.
	GatewayEndpoint string `json:"gateway_endpoint"`
	// Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'.
	Labels    map[string]string `json:"labels,omitempty"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func (s ServerlessCloudGateway) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServerlessCloudGateway) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServerlessCloudGateway) GetControlPlane() ServerlessControlPlane {
	if o == nil {
		return ServerlessControlPlane{}
	}
	return o.ControlPlane
}

func (o *ServerlessCloudGateway) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ServerlessCloudGateway) GetGatewayEndpoint() string {
	if o == nil {
		return ""
	}
	return o.GatewayEndpoint
}

func (o *ServerlessCloudGateway) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ServerlessCloudGateway) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}