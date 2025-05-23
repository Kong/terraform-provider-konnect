---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_api_product_version Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  APIProductVersion Resource
---

# konnect_api_product_version (Resource)

APIProductVersion Resource

## Example Usage

```terraform
resource "konnect_api_product_version" "my_apiproductversion" {
  api_product_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  deprecated     = false
  gateway_service = {
    control_plane_id = "e4d9ebb1-26b4-426a-b00e-cb67044f3baf"
    id               = "09b4786a-3e48-4631-8f6b-62d1d8e1a7f3"
  }
  labels = {
    key = "value"
  }
  name = "v1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_product_id` (String) The API Product ID
- `name` (String) The version name of the API product version.

### Optional

- `deprecated` (Boolean, Deprecated) Indicates if the version of the API product is deprecated. Applies deprecation or removes deprecation from all related portal product versions. This field is deprecated: Use [PortalProductVersion.deprecated](https://docs.konghq.com/konnect/api/portal-management/v2/#/operations/create-portal-product-version) instead.
- `gateway_service` (Attributes) (see [below for nested schema](#nestedatt--gateway_service))
- `labels` (Map of String) Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. 

Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".

### Read-Only

- `created_at` (String) An ISO-8601 timestamp representation of entity creation date.
- `id` (String) The API product version identifier.
- `portals` (Attributes List) The list of portals which this API product version is configured for (see [below for nested schema](#nestedatt--portals))
- `updated_at` (String) An ISO-8601 timestamp representation of entity update date.

<a id="nestedatt--gateway_service"></a>
### Nested Schema for `gateway_service`

Optional:

- `control_plane_id` (String) The identifier of the control plane that the gateway service resides in. Not Null
- `id` (String) The identifier of a gateway service associated with the version of the API product. Not Null

Read-Only:

- `runtime_group_id` (String, Deprecated) This field is deprecated, please use `control_plane_id` instead. The identifier of the control plane that the gateway service resides in


<a id="nestedatt--portals"></a>
### Nested Schema for `portals`

Read-Only:

- `application_registration_enabled` (Boolean)
- `auth_strategies` (Attributes List) (see [below for nested schema](#nestedatt--portals--auth_strategies))
- `auto_approve_registration` (Boolean)
- `deprecated` (Boolean)
- `portal_id` (String)
- `portal_name` (String)
- `portal_product_version_id` (String)
- `publish_status` (String) must be one of ["published", "unpublished"]

<a id="nestedatt--portals--auth_strategies"></a>
### Nested Schema for `portals.auth_strategies`

Read-Only:

- `id` (String)
- `name` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_api_product_version.my_konnect_api_product_version "{ \"api_product_id\": \"d32d905a-ed33-46a3-a093-d8f536af9a8a\",  \"id\": \"9f5061ce-78f6-4452-9108-ad7c02821fd5\"}"
```
