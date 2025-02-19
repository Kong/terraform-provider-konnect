// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DcrProviderCurityProviderType string

const (
	DcrProviderCurityProviderTypeCurity DcrProviderCurityProviderType = "curity"
)

func (e DcrProviderCurityProviderType) ToPointer() *DcrProviderCurityProviderType {
	return &e
}
func (e *DcrProviderCurityProviderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "curity":
		*e = DcrProviderCurityProviderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DcrProviderCurityProviderType: %v", v)
	}
}

// DcrProviderCurity - A DCR provider for Curity -- only properties not included in DcrProviderBase
type DcrProviderCurity struct {
	// Contains a unique identifier used by the API for this resource.
	ID *string `json:"id,omitempty"`
	// The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI.
	//
	Name *string `json:"name,omitempty"`
	// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
	//
	DisplayName *string `json:"display_name,omitempty"`
	// The issuer of the DCR provider.
	Issuer *string `json:"issuer,omitempty"`
	// At least one active auth strategy is using this DCR provider.
	Active       *bool                         `json:"active,omitempty"`
	ProviderType DcrProviderCurityProviderType `json:"provider_type"`
	// A DCR provider configuration for Curity
	DcrConfig DcrConfigCurityInResponse `json:"dcr_config"`
}

func (o *DcrProviderCurity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DcrProviderCurity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DcrProviderCurity) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *DcrProviderCurity) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *DcrProviderCurity) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *DcrProviderCurity) GetProviderType() DcrProviderCurityProviderType {
	if o == nil {
		return DcrProviderCurityProviderType("")
	}
	return o.ProviderType
}

func (o *DcrProviderCurity) GetDcrConfig() DcrConfigCurityInResponse {
	if o == nil {
		return DcrConfigCurityInResponse{}
	}
	return o.DcrConfig
}
