// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type Two struct {
	// The most basic mode to configure an Application Auth Strategy for an API Product Version.
	// Using this mode will allow developers to generate API keys that will authenticate their application requests.
	// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for Key Auth.
	//
	KeyAuth AppAuthStrategyConfigKeyAuth `json:"key-auth"`
}

func (o *Two) GetKeyAuth() AppAuthStrategyConfigKeyAuth {
	if o == nil {
		return AppAuthStrategyConfigKeyAuth{}
	}
	return o.KeyAuth
}

type One struct {
	// A more advanced mode to configure an API Product Version’s Application Auth Strategy.
	// Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests.
	// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy.
	// An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.
	//
	OpenidConnect PartialAppAuthStrategyConfigOpenIDConnect `json:"openid-connect"`
}

func (o *One) GetOpenidConnect() PartialAppAuthStrategyConfigOpenIDConnect {
	if o == nil {
		return PartialAppAuthStrategyConfigOpenIDConnect{}
	}
	return o.OpenidConnect
}

type ConfigsType string

const (
	ConfigsTypeOne ConfigsType = "1"
	ConfigsTypeTwo ConfigsType = "2"
)

// Configs - JSON-B object containing the configuration for the OIDC strategy under the key 'openid-connect' or the configuration for the Key Auth strategy under the key 'key-auth'
type Configs struct {
	One *One `queryParam:"inline"`
	Two *Two `queryParam:"inline"`

	Type ConfigsType
}

func CreateConfigsOne(one One) Configs {
	typ := ConfigsTypeOne

	return Configs{
		One:  &one,
		Type: typ,
	}
}

func CreateConfigsTwo(two Two) Configs {
	typ := ConfigsTypeTwo

	return Configs{
		Two:  &two,
		Type: typ,
	}
}

func (u *Configs) UnmarshalJSON(data []byte) error {

	var one One = One{}
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = ConfigsTypeOne
		return nil
	}

	var two Two = Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = ConfigsTypeTwo
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Configs", string(data))
}

func (u Configs) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type Configs: all fields are null")
}

// UpdateAppAuthStrategyRequest - Request body for updating an Application Auth Strategy
type UpdateAppAuthStrategyRequest struct {
	// The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
	//
	Name *string `json:"name,omitempty"`
	// The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
	//
	DisplayName *string `json:"display_name,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Labels are intended to store **INTERNAL** metadata.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels        map[string]*string `json:"labels,omitempty"`
	DcrProviderID *string            `json:"dcr_provider_id,omitempty"`
	// JSON-B object containing the configuration for the OIDC strategy under the key 'openid-connect' or the configuration for the Key Auth strategy under the key 'key-auth'
	Configs *Configs `json:"configs,omitempty"`
}

func (o *UpdateAppAuthStrategyRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateAppAuthStrategyRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *UpdateAppAuthStrategyRequest) GetLabels() map[string]*string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *UpdateAppAuthStrategyRequest) GetDcrProviderID() *string {
	if o == nil {
		return nil
	}
	return o.DcrProviderID
}

func (o *UpdateAppAuthStrategyRequest) GetConfigs() *Configs {
	if o == nil {
		return nil
	}
	return o.Configs
}
