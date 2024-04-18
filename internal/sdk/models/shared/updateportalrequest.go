// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdatePortalRequest struct {
	// The name of the portal, used to distinguish it from other portals. Name must be unique.
	Name *string `json:"name,omitempty"`
	// The description of the portal.
	Description *string `json:"description,omitempty"`
	// Whether the portal catalog can be accessed publicly without any developer authentication. Developer accounts and applications cannot be created if the portal is public.
	IsPublic *bool `json:"is_public,omitempty"`
	// Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for products until unless assigned to teams with access to view and consume specific products.
	RbacEnabled *bool `json:"rbac_enabled,omitempty"`
	// Whether the requests from applications to register for products will be automatically approved, or if they will be set to pending until approved by an admin.
	AutoApproveApplications *bool `json:"auto_approve_applications,omitempty"`
	// Whether the developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.
	AutoApproveDevelopers *bool `json:"auto_approve_developers,omitempty"`
	// The custom domain to access the developer portal. A CNAME for the portal's default domain must be able to be set for the custom domain for it to be valid. After setting a valid CNAME, an SSL/TLS certificate will be automatically manged for the custom domain, and traffic will be able to use the custom domain to route to the portal's web client and API.
	CustomDomain *string `json:"custom_domain,omitempty"`
	// The custom domain to access a self-hosted customized developer portal client. If this is set, the Konnect-hosted portal will no longer be available.  `custom_domain` must be also set for this value to be set. See https://github.com/Kong/konnect-portal for information on how to get started deploying and customizing your own Konnect portal.
	CustomClientDomain *string `json:"custom_client_domain,omitempty"`
	// Default strategy ID applied on applications for the portal
	DefaultApplicationAuthStrategyID *string `json:"default_application_auth_strategy_id,omitempty"`
}

func (o *UpdatePortalRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdatePortalRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdatePortalRequest) GetIsPublic() *bool {
	if o == nil {
		return nil
	}
	return o.IsPublic
}

func (o *UpdatePortalRequest) GetRbacEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.RbacEnabled
}

func (o *UpdatePortalRequest) GetAutoApproveApplications() *bool {
	if o == nil {
		return nil
	}
	return o.AutoApproveApplications
}

func (o *UpdatePortalRequest) GetAutoApproveDevelopers() *bool {
	if o == nil {
		return nil
	}
	return o.AutoApproveDevelopers
}

func (o *UpdatePortalRequest) GetCustomDomain() *string {
	if o == nil {
		return nil
	}
	return o.CustomDomain
}

func (o *UpdatePortalRequest) GetCustomClientDomain() *string {
	if o == nil {
		return nil
	}
	return o.CustomClientDomain
}

func (o *UpdatePortalRequest) GetDefaultApplicationAuthStrategyID() *string {
	if o == nil {
		return nil
	}
	return o.DefaultApplicationAuthStrategyID
}
