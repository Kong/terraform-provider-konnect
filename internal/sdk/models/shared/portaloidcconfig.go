// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PortalOIDCConfig - Configuration properties for an OpenID Connect Identity Provider.
type PortalOIDCConfig struct {
	Issuer   string   `json:"issuer"`
	ClientID string   `json:"client_id"`
	Scopes   []string `json:"scopes,omitempty"`
	// Mappings from a portal developer atribute to an Identity Provider claim.
	ClaimMappings *PortalClaimMappings `json:"claim_mappings,omitempty"`
}

func (o *PortalOIDCConfig) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *PortalOIDCConfig) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *PortalOIDCConfig) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *PortalOIDCConfig) GetClaimMappings() *PortalClaimMappings {
	if o == nil {
		return nil
	}
	return o.ClaimMappings
}
