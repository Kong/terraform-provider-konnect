// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"time"
)

// UpdatePortalResponse - Details about the portal being updated.
type UpdatePortalResponse struct {
	// Contains a unique identifier used by the API for this resource.
	ID string `json:"id"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
	// The name of the portal, used to distinguish it from other portals. Name must be unique.
	Name string `json:"name"`
	// The display name of the portal. This value will be the portal's `name` in Portal API.
	DisplayName string `json:"display_name"`
	// The description of the portal.
	Description *string `json:"description"`
	// The domain assigned to the portal by Konnect. This is the default place to access the portal and its API if not using a `custom_domain``.
	DefaultDomain string `json:"default_domain"`
	// Whether the portal catalog can be accessed publicly without any developer authentication. Developer accounts and applications cannot be created if the portal is public.
	IsPublic bool `json:"is_public"`
	// Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for products until unless assigned to teams with access to view and consume specific products.
	RbacEnabled bool `json:"rbac_enabled"`
	// Whether the requests from applications to register for products will be automatically approved, or if they will be set to pending until approved by an admin.
	AutoApproveApplications bool `json:"auto_approve_applications"`
	// Whether the developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.
	AutoApproveDevelopers bool `json:"auto_approve_developers"`
	// The custom domain to access the developer portal. A CNAME for the portal's default domain must be able to be set for the custom domain for it to be valid. After setting a valid CNAME, an SSL/TLS certificate will be automatically manged for the custom domain, and traffic will be able to use the custom domain to route to the portal's web client and API.
	CustomDomain *string `json:"custom_domain"`
	// The custom domain to access a self-hosted customized developer portal client. If this is set, the Konnect-hosted portal client will no longer be available. `custom_domain` must be also set for this value to be set. See https://github.com/Kong/konnect-portal for information on how to get started deploying and customizing your own Konnect portal.
	CustomClientDomain *string `json:"custom_client_domain"`
	// Default strategy ID applied on applications for the portal
	DefaultApplicationAuthStrategyID *string `json:"default_application_auth_strategy_id,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
	// Number of applications created in the portal.
	ApplicationCount float64 `json:"application_count"`
	// Number of developers using the portal.
	DeveloperCount float64 `json:"developer_count"`
	// Number of api products published to the portal
	PublishedProductCount float64 `json:"published_product_count"`
}

func (u UpdatePortalResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdatePortalResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdatePortalResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdatePortalResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *UpdatePortalResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *UpdatePortalResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdatePortalResponse) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *UpdatePortalResponse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdatePortalResponse) GetDefaultDomain() string {
	if o == nil {
		return ""
	}
	return o.DefaultDomain
}

func (o *UpdatePortalResponse) GetIsPublic() bool {
	if o == nil {
		return false
	}
	return o.IsPublic
}

func (o *UpdatePortalResponse) GetRbacEnabled() bool {
	if o == nil {
		return false
	}
	return o.RbacEnabled
}

func (o *UpdatePortalResponse) GetAutoApproveApplications() bool {
	if o == nil {
		return false
	}
	return o.AutoApproveApplications
}

func (o *UpdatePortalResponse) GetAutoApproveDevelopers() bool {
	if o == nil {
		return false
	}
	return o.AutoApproveDevelopers
}

func (o *UpdatePortalResponse) GetCustomDomain() *string {
	if o == nil {
		return nil
	}
	return o.CustomDomain
}

func (o *UpdatePortalResponse) GetCustomClientDomain() *string {
	if o == nil {
		return nil
	}
	return o.CustomClientDomain
}

func (o *UpdatePortalResponse) GetDefaultApplicationAuthStrategyID() *string {
	if o == nil {
		return nil
	}
	return o.DefaultApplicationAuthStrategyID
}

func (o *UpdatePortalResponse) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *UpdatePortalResponse) GetApplicationCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.ApplicationCount
}

func (o *UpdatePortalResponse) GetDeveloperCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.DeveloperCount
}

func (o *UpdatePortalResponse) GetPublishedProductCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.PublishedProductCount
}
