// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"time"
)

// CreateSystemAccountAccessToken - The request body to create a system account access token.
type CreateSystemAccountAccessToken struct {
	Name      *string    `json:"name,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

func (c CreateSystemAccountAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateSystemAccountAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateSystemAccountAccessToken) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateSystemAccountAccessToken) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}