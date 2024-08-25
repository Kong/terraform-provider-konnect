// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CreateNetworkRequest - Request schema for creating a network.
type CreateNetworkRequest struct {
	// Human-readable name of the network.
	Name                          string `json:"name"`
	CloudGatewayProviderAccountID string `json:"cloud_gateway_provider_account_id"`
	// Region ID for cloud provider region.
	Region string `json:"region"`
	// List of availability zones that the network is attached to.
	AvailabilityZones []string `json:"availability_zones"`
	// CIDR block configuration for the network.
	CidrBlock string `json:"cidr_block"`
	// Firewall configuration for a network.
	Firewall *NetworkFirewallConfig `json:"firewall,omitempty"`
	// Whether DDOS protection is enabled for the network.
	DdosProtection *bool `json:"ddos_protection,omitempty"`
}

func (o *CreateNetworkRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateNetworkRequest) GetCloudGatewayProviderAccountID() string {
	if o == nil {
		return ""
	}
	return o.CloudGatewayProviderAccountID
}

func (o *CreateNetworkRequest) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}

func (o *CreateNetworkRequest) GetAvailabilityZones() []string {
	if o == nil {
		return []string{}
	}
	return o.AvailabilityZones
}

func (o *CreateNetworkRequest) GetCidrBlock() string {
	if o == nil {
		return ""
	}
	return o.CidrBlock
}

func (o *CreateNetworkRequest) GetFirewall() *NetworkFirewallConfig {
	if o == nil {
		return nil
	}
	return o.Firewall
}

func (o *CreateNetworkRequest) GetDdosProtection() *bool {
	if o == nil {
		return nil
	}
	return o.DdosProtection
}