// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

// PartialAppAuthStrategyConfigOpenIDConnect - A more advanced mode to configure an API Product Version’s Application Auth Strategy.
// Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests.
// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy.
// An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.
type PartialAppAuthStrategyConfigOpenIDConnect struct {
	AuthMethods          []string `json:"auth_methods,omitempty"`
	CredentialClaim      []string `json:"credential_claim,omitempty"`
	Issuer               *string  `json:"issuer,omitempty"`
	Scopes               []string `json:"scopes,omitempty"`
	AdditionalProperties any      `additionalProperties:"true" json:"-"`
}

func (p PartialAppAuthStrategyConfigOpenIDConnect) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PartialAppAuthStrategyConfigOpenIDConnect) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PartialAppAuthStrategyConfigOpenIDConnect) GetAuthMethods() []string {
	if o == nil {
		return nil
	}
	return o.AuthMethods
}

func (o *PartialAppAuthStrategyConfigOpenIDConnect) GetCredentialClaim() []string {
	if o == nil {
		return nil
	}
	return o.CredentialClaim
}

func (o *PartialAppAuthStrategyConfigOpenIDConnect) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *PartialAppAuthStrategyConfigOpenIDConnect) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *PartialAppAuthStrategyConfigOpenIDConnect) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
