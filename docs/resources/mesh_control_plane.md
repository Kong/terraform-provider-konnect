---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_control_plane Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshControlPlane Resource
---

# konnect_mesh_control_plane (Resource)

MeshControlPlane Resource

## Example Usage

```terraform
resource "konnect_mesh_control_plane" "my_meshcontrolplane" {
  description = "A control plane to handle traffic on development environment."
  features = [
    {
      hostname_generator_creation = {
        enabled = false
      }
      mesh_creation = {
        enabled = false
      }
      type = "MeshCreation"
    }
  ]
  labels = {
    key = "value"
  }
  name = "Test control plane"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the control plane.

### Optional

- `description` (String)
- `features` (Attributes List) Requires replacement if changed. (see [below for nested schema](#nestedatt--features))
- `labels` (Map of String) Labels to facilitate tagged search on control planes. Keys must be of length 1-63 characters.

### Read-Only

- `created_at` (String)
- `id` (String) ID of the control plane.
- `updated_at` (String)

<a id="nestedatt--features"></a>
### Nested Schema for `features`

Optional:

- `hostname_generator_creation` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--features--hostname_generator_creation))
- `mesh_creation` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--features--mesh_creation))
- `type` (String) Not Null; must be one of ["MeshCreation", "HostnameGeneratorCreation"]; Requires replacement if changed.

<a id="nestedatt--features--hostname_generator_creation"></a>
### Nested Schema for `features.hostname_generator_creation`

Optional:

- `enabled` (Boolean) Not Null; Requires replacement if changed.


<a id="nestedatt--features--mesh_creation"></a>
### Nested Schema for `features.mesh_creation`

Optional:

- `enabled` (Boolean) Not Null; Requires replacement if changed.

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_mesh_control_plane.my_konnect_mesh_control_plane "d32d905a-ed33-46a3-a093-d8f536af9a8a"
```
