// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// ProviderAccount - Object containing mapping for organization and cloud provider to account ID.
type ProviderAccount struct {
	ID string `json:"id"`
	// Name of cloud provider.
	Provider ProviderName `json:"provider"`
	// ID of the cloud provider account.
	ProviderAccountID string `json:"provider_account_id"`
	// An RFC-3339 timestamp representation of provider account creation date.
	CreatedAt time.Time `json:"created_at"`
	// An RFC-3339 timestamp representation of provider account update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (p ProviderAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProviderAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProviderAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ProviderAccount) GetProvider() ProviderName {
	if o == nil {
		return ProviderName("")
	}
	return o.Provider
}

func (o *ProviderAccount) GetProviderAccountID() string {
	if o == nil {
		return ""
	}
	return o.ProviderAccountID
}

func (o *ProviderAccount) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ProviderAccount) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
