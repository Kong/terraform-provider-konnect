// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CredentialType string

const (
	CredentialTypeKeyAuth CredentialType = "key_auth"
)

func (e CredentialType) ToPointer() *CredentialType {
	return &e
}
func (e *CredentialType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "key_auth":
		*e = CredentialType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CredentialType: %v", v)
	}
}

// AuthStrategyKeyAuth - KeyAuth Auth strategy that the application uses.
type AuthStrategyKeyAuth struct {
	// The Application Auth Strategy ID.
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	CredentialType CredentialType `json:"credential_type"`
}

func (o *AuthStrategyKeyAuth) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthStrategyKeyAuth) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AuthStrategyKeyAuth) GetCredentialType() CredentialType {
	if o == nil {
		return CredentialType("")
	}
	return o.CredentialType
}