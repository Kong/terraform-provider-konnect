---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_control_planes Data Source - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshControlPlanes DataSource
---

# konnect_mesh_control_planes (Data Source)

MeshControlPlanes DataSource

## Example Usage

```terraform
data "konnect_mesh_control_planes" "my_meshcontrolplanes" {
  page_number = 1
  page_size   = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `page_number` (Number) Determines which page of the entities to retrieve.
- `page_size` (Number) The maximum number of items to include per page. The last page of a collection may include fewer items.

### Read-Only

- `data` (Attributes List) (see [below for nested schema](#nestedatt--data))
- `meta` (Attributes) returns the pagination information (see [below for nested schema](#nestedatt--meta))

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `created_at` (String)
- `description` (String)
- `features` (Attributes List) (see [below for nested schema](#nestedatt--data--features))
- `id` (String) ID of the control plane.
- `labels` (Map of String) Labels to facilitate tagged search on control planes. Keys must be of length 1-63 characters.
- `name` (String) The name of the control plane.
- `updated_at` (String)

<a id="nestedatt--data--features"></a>
### Nested Schema for `data.features`

Read-Only:

- `hostname_generator_creation` (Attributes) (see [below for nested schema](#nestedatt--data--features--hostname_generator_creation))
- `mesh_creation` (Attributes) (see [below for nested schema](#nestedatt--data--features--mesh_creation))
- `type` (String)

<a id="nestedatt--data--features--hostname_generator_creation"></a>
### Nested Schema for `data.features.hostname_generator_creation`

Read-Only:

- `enabled` (Boolean)


<a id="nestedatt--data--features--mesh_creation"></a>
### Nested Schema for `data.features.mesh_creation`

Read-Only:

- `enabled` (Boolean)




<a id="nestedatt--meta"></a>
### Nested Schema for `meta`

Read-Only:

- `page` (Attributes) Contains pagination query parameters and the total number of objects returned. (see [below for nested schema](#nestedatt--meta--page))

<a id="nestedatt--meta--page"></a>
### Nested Schema for `meta.page`

Read-Only:

- `number` (Number)
- `size` (Number)
- `total` (Number)
