---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_portal Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  Portal Resource
---

# konnect_portal (Resource)

Portal Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `portal_id` (String) ID of the portal.

### Optional

- `auto_approve_applications` (Boolean) Whether the requests from applications to register for products will be automatically approved, or if they will be set to pending until approved by an admin.
- `auto_approve_developers` (Boolean) Whether the developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.
- `custom_client_domain` (String) The custom domain to access a self-hosted customized developer portal client. If this is set, the Konnect-hosted portal will no longer be available.  `custom_domain` must be also set for this value to be set. See https://github.com/Kong/konnect-portal for information on how to get started deploying and customizing your own Konnect portal.
- `custom_domain` (String) The custom domain to access the developer portal. A CNAME for the portal's default domain must be able to be set for the custom domain for it to be valid. After setting a valid CNAME, an SSL/TLS certificate will be automatically manged for the custom domain, and traffic will be able to use the custom domain to route to the portal's web client and API.
- `default_application_auth_strategy_id` (String) Default strategy ID applied on applications for the portal
- `description` (String) The description of the portal.
- `display_name` (String) The display name of the portal. This value will be the portal's `name` in Portal API.
- `is_public` (Boolean) Whether the portal catalog can be accessed publicly without any developer authentication. Developer accounts and applications cannot be created if the portal is public.
- `name` (String) The name of the portal, used to distinguish it from other portals. Name must be unique.
- `rbac_enabled` (Boolean) Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for products until unless assigned to teams with access to view and consume specific products.

### Read-Only

- `application_count` (Number) Number of applications created in the portal.
- `created_at` (String) An ISO-8601 timestamp representation of entity creation date.
- `default_domain` (String) The domain assigned to the portal by Konnect. This is the default place to access the portal and its API if not using a `custom_domain``.
- `developer_count` (Number) Number of developers using the portal.
- `id` (String) Contains a unique identifier used for this resource.
- `published_product_count` (Number) Number of api products published to the portal
- `updated_at` (String) An ISO-8601 timestamp representation of entity update date.

