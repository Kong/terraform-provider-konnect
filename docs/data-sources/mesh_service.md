---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_service Data Source - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshService DataSource
---

# konnect_mesh_service (Data Source)

MeshService DataSource

## Example Usage

```terraform
data "konnect_mesh_service" "my_meshservice" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  mesh  = "...my_mesh..."
  name  = "...my_name..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cp_id` (String) Id of the Konnect resource
- `mesh` (String) name of the mesh
- `name` (String) name of the MeshService

### Read-Only

- `creation_time` (String) Time at which the resource was created
- `labels` (Map of String) The labels to help identity resources
- `modification_time` (String) Time at which the resource was updated
- `spec` (Attributes) Spec is the specification of the Kuma MeshService resource. (see [below for nested schema](#nestedatt--spec))
- `status` (Attributes) Status is the current status of the Kuma MeshService resource. (see [below for nested schema](#nestedatt--status))
- `type` (String) the type of the resource

<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Read-Only:

- `identities` (Attributes List) (see [below for nested schema](#nestedatt--spec--identities))
- `ports` (Attributes List) (see [below for nested schema](#nestedatt--spec--ports))
- `selector` (Attributes) (see [below for nested schema](#nestedatt--spec--selector))
- `state` (String) State of MeshService. Available if there is at least one healthy endpoint. Otherwise, Unavailable.
It's used for cross zone communication to check if we should send traffic to it, when MeshService is aggregated into MeshMultiZoneService.

<a id="nestedatt--spec--identities"></a>
### Nested Schema for `spec.identities`

Read-Only:

- `type` (String)
- `value` (String)


<a id="nestedatt--spec--ports"></a>
### Nested Schema for `spec.ports`

Read-Only:

- `app_protocol` (String) Protocol identifies a protocol supported by a service.
- `name` (String)
- `port` (Number)
- `target_port` (Attributes) (see [below for nested schema](#nestedatt--spec--ports--target_port))

<a id="nestedatt--spec--ports--target_port"></a>
### Nested Schema for `spec.ports.target_port`

Read-Only:

- `integer` (Number)
- `str` (String)



<a id="nestedatt--spec--selector"></a>
### Nested Schema for `spec.selector`

Read-Only:

- `dataplane_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--selector--dataplane_ref))
- `dataplane_tags` (Map of String)

<a id="nestedatt--spec--selector--dataplane_ref"></a>
### Nested Schema for `spec.selector.dataplane_ref`

Read-Only:

- `name` (String)




<a id="nestedatt--status"></a>
### Nested Schema for `status`

Read-Only:

- `addresses` (Attributes List) (see [below for nested schema](#nestedatt--status--addresses))
- `dataplane_proxies` (Attributes) Data plane proxies statistics selected by this MeshService. (see [below for nested schema](#nestedatt--status--dataplane_proxies))
- `hostname_generators` (Attributes List) (see [below for nested schema](#nestedatt--status--hostname_generators))
- `tls` (Attributes) (see [below for nested schema](#nestedatt--status--tls))
- `vips` (Attributes List) (see [below for nested schema](#nestedatt--status--vips))

<a id="nestedatt--status--addresses"></a>
### Nested Schema for `status.addresses`

Read-Only:

- `hostname` (String)
- `hostname_generator_ref` (Attributes) (see [below for nested schema](#nestedatt--status--addresses--hostname_generator_ref))
- `origin` (String)

<a id="nestedatt--status--addresses--hostname_generator_ref"></a>
### Nested Schema for `status.addresses.hostname_generator_ref`

Read-Only:

- `core_name` (String)



<a id="nestedatt--status--dataplane_proxies"></a>
### Nested Schema for `status.dataplane_proxies`

Read-Only:

- `connected` (Number) Number of data plane proxies connected to the zone control plane
- `healthy` (Number) Number of data plane proxies with all healthy inbounds selected by this MeshService.
- `total` (Number) Total number of data plane proxies.


<a id="nestedatt--status--hostname_generators"></a>
### Nested Schema for `status.hostname_generators`

Read-Only:

- `conditions` (Attributes List) Conditions is an array of hostname generator conditions. (see [below for nested schema](#nestedatt--status--hostname_generators--conditions))
- `hostname_generator_ref` (Attributes) (see [below for nested schema](#nestedatt--status--hostname_generators--hostname_generator_ref))

<a id="nestedatt--status--hostname_generators--conditions"></a>
### Nested Schema for `status.hostname_generators.conditions`

Read-Only:

- `message` (String) message is a human readable message indicating details about the transition.
This may be an empty string.
- `reason` (String) reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.
- `status` (String) status of the condition, one of True, False, Unknown.
- `type` (String) type of condition in CamelCase or in foo.example.com/CamelCase.


<a id="nestedatt--status--hostname_generators--hostname_generator_ref"></a>
### Nested Schema for `status.hostname_generators.hostname_generator_ref`

Read-Only:

- `core_name` (String)



<a id="nestedatt--status--tls"></a>
### Nested Schema for `status.tls`

Read-Only:

- `status` (String)


<a id="nestedatt--status--vips"></a>
### Nested Schema for `status.vips`

Read-Only:

- `ip` (String)