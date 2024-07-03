---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_acl Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayACL Resource
---

# konnect_gateway_acl (Resource)

GatewayACL Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `consumer_id` (String) Consumer ID for nested entities. Requires replacement if changed.
- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.

### Optional

- `group` (String) Requires replacement if changed.
- `tags` (List of String) Requires replacement if changed.

### Read-Only

- `consumer` (Attributes) (see [below for nested schema](#nestedatt--consumer))
- `created_at` (Number) Unix epoch when the resource was created.
- `id` (String) The ID of this resource.

<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

Read-Only:

- `id` (String)

