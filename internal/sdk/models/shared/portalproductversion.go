// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"time"
)

// PortalProductVersion - A Portal Product Version holds metadata that describes how a Product Version is configured for a specific portal.
//
// It contains:
// - Lifecyle and deprecation statuses
// - Application registration settings like auto approve or whether application registration is enabled
// - The Authentication Strategy (if present) that is enabled for Application Registration
type PortalProductVersion struct {
	// Whether the application registration on this portal for the api product version is enabled
	ApplicationRegistrationEnabled bool `json:"application_registration_enabled"`
	// A list of authentication strategies
	AuthStrategies []AuthStrategy `json:"auth_strategies"`
	// Whether the application registration auto approval on this portal for the api product version is enabled
	AutoApproveRegistration bool `json:"auto_approve_registration"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// Whether the api product version on the portal is deprecated
	Deprecated bool `json:"deprecated"`
	// Contains a unique identifier used by the API for this resource.
	ID string `json:"id"`
	// Contains a unique identifier used by the API for this resource.
	ProductVersionID string `json:"product_version_id"`
	// Publication status of the API product version on the portal
	PublishStatus PortalProductVersionPublishStatus `json:"publish_status"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (p PortalProductVersion) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PortalProductVersion) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PortalProductVersion) GetApplicationRegistrationEnabled() bool {
	if o == nil {
		return false
	}
	return o.ApplicationRegistrationEnabled
}

func (o *PortalProductVersion) GetAuthStrategies() []AuthStrategy {
	if o == nil {
		return []AuthStrategy{}
	}
	return o.AuthStrategies
}

func (o *PortalProductVersion) GetAutoApproveRegistration() bool {
	if o == nil {
		return false
	}
	return o.AutoApproveRegistration
}

func (o *PortalProductVersion) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PortalProductVersion) GetDeprecated() bool {
	if o == nil {
		return false
	}
	return o.Deprecated
}

func (o *PortalProductVersion) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PortalProductVersion) GetProductVersionID() string {
	if o == nil {
		return ""
	}
	return o.ProductVersionID
}

func (o *PortalProductVersion) GetPublishStatus() PortalProductVersionPublishStatus {
	if o == nil {
		return PortalProductVersionPublishStatus("")
	}
	return o.PublishStatus
}

func (o *PortalProductVersion) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
