---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_proxy_patch Data Source - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshProxyPatch DataSource
---

# konnect_mesh_proxy_patch (Data Source)

MeshProxyPatch DataSource

## Example Usage

```terraform
data "konnect_mesh_proxy_patch" "my_meshproxypatch" {
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
- `name` (String) name of the MeshProxyPatch

### Read-Only

- `creation_time` (String) Time at which the resource was created
- `labels` (Map of String) The labels to help identity resources
- `modification_time` (String) Time at which the resource was updated
- `spec` (Attributes) Spec is the specification of the Kuma MeshProxyPatch resource. (see [below for nested schema](#nestedatt--spec))
- `type` (String) the type of the resource

<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Read-Only:

- `default` (Attributes) Default is a configuration specific to the group of destinations
referenced in 'targetRef'. (see [below for nested schema](#nestedatt--spec--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource the policy takes an effect on.
The resource could be either a real store object or virtual resource
defined inplace. (see [below for nested schema](#nestedatt--spec--target_ref))

<a id="nestedatt--spec--default"></a>
### Nested Schema for `spec.default`

Read-Only:

- `append_modifications` (Attributes List) AppendModifications is a list of modifications applied on the selected proxy. (see [below for nested schema](#nestedatt--spec--default--append_modifications))

<a id="nestedatt--spec--default--append_modifications"></a>
### Nested Schema for `spec.default.append_modifications`

Read-Only:

- `cluster` (Attributes) Cluster is a modification of Envoy's Cluster resource. (see [below for nested schema](#nestedatt--spec--default--append_modifications--cluster))
- `http_filter` (Attributes) HTTPFilter is a modification of Envoy HTTP Filter
available in HTTP Connection Manager in a Listener resource. (see [below for nested schema](#nestedatt--spec--default--append_modifications--http_filter))
- `listener` (Attributes) Listener is a modification of Envoy's Listener resource. (see [below for nested schema](#nestedatt--spec--default--append_modifications--listener))
- `network_filter` (Attributes) NetworkFilter is a modification of Envoy Listener's filter. (see [below for nested schema](#nestedatt--spec--default--append_modifications--network_filter))
- `virtual_host` (Attributes) VirtualHost is a modification of Envoy's VirtualHost
referenced in HTTP Connection Manager in a Listener resource. (see [below for nested schema](#nestedatt--spec--default--append_modifications--virtual_host))

<a id="nestedatt--spec--default--append_modifications--cluster"></a>
### Nested Schema for `spec.default.append_modifications.cluster`

Read-Only:

- `json_patches` (Attributes List) JsonPatches specifies list of jsonpatches to apply to on Envoy's Cluster
resource (see [below for nested schema](#nestedatt--spec--default--append_modifications--cluster--json_patches))
- `match` (Attributes) Match is a set of conditions that have to be matched for modification operation to happen. (see [below for nested schema](#nestedatt--spec--default--append_modifications--cluster--match))
- `operation` (String) Operation to execute on matched cluster.
- `value` (String) Value of xDS resource in YAML format to add or patch.

<a id="nestedatt--spec--default--append_modifications--cluster--json_patches"></a>
### Nested Schema for `spec.default.append_modifications.cluster.json_patches`

Read-Only:

- `from` (String) From is a jsonpatch from string, used by move and copy operations.
- `op` (String) Op is a jsonpatch operation string.
- `path` (String) Path is a jsonpatch path string.
- `value` (String) Value must be a valid json value used by replace and add operations. Parsed as JSON.


<a id="nestedatt--spec--default--append_modifications--cluster--match"></a>
### Nested Schema for `spec.default.append_modifications.cluster.match`

Read-Only:

- `name` (String) Name of the cluster to match.
- `origin` (String) Origin is the name of the component or plugin that generated the resource.

Here is the list of well-known origins:
inbound - resources generated for handling incoming traffic.
outbound - resources generated for handling outgoing traffic.
transparent - resources generated for transparent proxy functionality.
prometheus - resources generated when Prometheus metrics are enabled.
direct-access - resources generated for Direct Access functionality.
ingress - resources generated for Zone Ingress.
egress - resources generated for Zone Egress.
gateway - resources generated for MeshGateway.

The list is not complete, because policy plugins can introduce new resources.
For example MeshTrace plugin can create Cluster with "mesh-trace" origin.



<a id="nestedatt--spec--default--append_modifications--http_filter"></a>
### Nested Schema for `spec.default.append_modifications.http_filter`

Read-Only:

- `json_patches` (Attributes List) JsonPatches specifies list of jsonpatches to apply to on Envoy's
HTTP Filter available in HTTP Connection Manager in a Listener resource. (see [below for nested schema](#nestedatt--spec--default--append_modifications--http_filter--json_patches))
- `match` (Attributes) Match is a set of conditions that have to be matched for modification operation to happen. (see [below for nested schema](#nestedatt--spec--default--append_modifications--http_filter--match))
- `operation` (String) Operation to execute on matched listener.
- `value` (String) Value of xDS resource in YAML format to add or patch.

<a id="nestedatt--spec--default--append_modifications--http_filter--json_patches"></a>
### Nested Schema for `spec.default.append_modifications.http_filter.json_patches`

Read-Only:

- `from` (String) From is a jsonpatch from string, used by move and copy operations.
- `op` (String) Op is a jsonpatch operation string.
- `path` (String) Path is a jsonpatch path string.
- `value` (String) Value must be a valid json value used by replace and add operations. Parsed as JSON.


<a id="nestedatt--spec--default--append_modifications--http_filter--match"></a>
### Nested Schema for `spec.default.append_modifications.http_filter.match`

Read-Only:

- `listener_name` (String) Name of the listener to match.
- `listener_tags` (Map of String) Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]
- `name` (String) Name of the HTTP filter. For example "envoy.filters.http.local_ratelimit"
- `origin` (String) Origin is the name of the component or plugin that generated the resource.

Here is the list of well-known origins:
inbound - resources generated for handling incoming traffic.
outbound - resources generated for handling outgoing traffic.
transparent - resources generated for transparent proxy functionality.
prometheus - resources generated when Prometheus metrics are enabled.
direct-access - resources generated for Direct Access functionality.
ingress - resources generated for Zone Ingress.
egress - resources generated for Zone Egress.
gateway - resources generated for MeshGateway.

The list is not complete, because policy plugins can introduce new resources.
For example MeshTrace plugin can create Cluster with "mesh-trace" origin.



<a id="nestedatt--spec--default--append_modifications--listener"></a>
### Nested Schema for `spec.default.append_modifications.listener`

Read-Only:

- `json_patches` (Attributes List) JsonPatches specifies list of jsonpatches to apply to on Envoy's Listener
resource (see [below for nested schema](#nestedatt--spec--default--append_modifications--listener--json_patches))
- `match` (Attributes) Match is a set of conditions that have to be matched for modification operation to happen. (see [below for nested schema](#nestedatt--spec--default--append_modifications--listener--match))
- `operation` (String) Operation to execute on matched listener.
- `value` (String) Value of xDS resource in YAML format to add or patch.

<a id="nestedatt--spec--default--append_modifications--listener--json_patches"></a>
### Nested Schema for `spec.default.append_modifications.listener.json_patches`

Read-Only:

- `from` (String) From is a jsonpatch from string, used by move and copy operations.
- `op` (String) Op is a jsonpatch operation string.
- `path` (String) Path is a jsonpatch path string.
- `value` (String) Value must be a valid json value used by replace and add operations. Parsed as JSON.


<a id="nestedatt--spec--default--append_modifications--listener--match"></a>
### Nested Schema for `spec.default.append_modifications.listener.match`

Read-Only:

- `name` (String) Name of the listener to match.
- `origin` (String) Origin is the name of the component or plugin that generated the resource.

Here is the list of well-known origins:
inbound - resources generated for handling incoming traffic.
outbound - resources generated for handling outgoing traffic.
transparent - resources generated for transparent proxy functionality.
prometheus - resources generated when Prometheus metrics are enabled.
direct-access - resources generated for Direct Access functionality.
ingress - resources generated for Zone Ingress.
egress - resources generated for Zone Egress.
gateway - resources generated for MeshGateway.

The list is not complete, because policy plugins can introduce new resources.
For example MeshTrace plugin can create Cluster with "mesh-trace" origin.
- `tags` (Map of String) Tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]



<a id="nestedatt--spec--default--append_modifications--network_filter"></a>
### Nested Schema for `spec.default.append_modifications.network_filter`

Read-Only:

- `json_patches` (Attributes List) JsonPatches specifies list of jsonpatches to apply to on Envoy Listener's
filter. (see [below for nested schema](#nestedatt--spec--default--append_modifications--network_filter--json_patches))
- `match` (Attributes) Match is a set of conditions that have to be matched for modification operation to happen. (see [below for nested schema](#nestedatt--spec--default--append_modifications--network_filter--match))
- `operation` (String) Operation to execute on matched listener.
- `value` (String) Value of xDS resource in YAML format to add or patch.

<a id="nestedatt--spec--default--append_modifications--network_filter--json_patches"></a>
### Nested Schema for `spec.default.append_modifications.network_filter.json_patches`

Read-Only:

- `from` (String) From is a jsonpatch from string, used by move and copy operations.
- `op` (String) Op is a jsonpatch operation string.
- `path` (String) Path is a jsonpatch path string.
- `value` (String) Value must be a valid json value used by replace and add operations. Parsed as JSON.


<a id="nestedatt--spec--default--append_modifications--network_filter--match"></a>
### Nested Schema for `spec.default.append_modifications.network_filter.match`

Read-Only:

- `listener_name` (String) Name of the listener to match.
- `listener_tags` (Map of String) Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]
- `name` (String) Name of the network filter. For example "envoy.filters.network.ratelimit"
- `origin` (String) Origin is the name of the component or plugin that generated the resource.

Here is the list of well-known origins:
inbound - resources generated for handling incoming traffic.
outbound - resources generated for handling outgoing traffic.
transparent - resources generated for transparent proxy functionality.
prometheus - resources generated when Prometheus metrics are enabled.
direct-access - resources generated for Direct Access functionality.
ingress - resources generated for Zone Ingress.
egress - resources generated for Zone Egress.
gateway - resources generated for MeshGateway.

The list is not complete, because policy plugins can introduce new resources.
For example MeshTrace plugin can create Cluster with "mesh-trace" origin.



<a id="nestedatt--spec--default--append_modifications--virtual_host"></a>
### Nested Schema for `spec.default.append_modifications.virtual_host`

Read-Only:

- `json_patches` (Attributes List) JsonPatches specifies list of jsonpatches to apply to on Envoy's
VirtualHost resource (see [below for nested schema](#nestedatt--spec--default--append_modifications--virtual_host--json_patches))
- `match` (Attributes) Match is a set of conditions that have to be matched for modification operation to happen. (see [below for nested schema](#nestedatt--spec--default--append_modifications--virtual_host--match))
- `operation` (String) Operation to execute on matched listener.
- `value` (String) Value of xDS resource in YAML format to add or patch.

<a id="nestedatt--spec--default--append_modifications--virtual_host--json_patches"></a>
### Nested Schema for `spec.default.append_modifications.virtual_host.json_patches`

Read-Only:

- `from` (String) From is a jsonpatch from string, used by move and copy operations.
- `op` (String) Op is a jsonpatch operation string.
- `path` (String) Path is a jsonpatch path string.
- `value` (String) Value must be a valid json value used by replace and add operations. Parsed as JSON.


<a id="nestedatt--spec--default--append_modifications--virtual_host--match"></a>
### Nested Schema for `spec.default.append_modifications.virtual_host.match`

Read-Only:

- `name` (String) Name of the VirtualHost to match.
- `origin` (String) Origin is the name of the component or plugin that generated the resource.

Here is the list of well-known origins:
inbound - resources generated for handling incoming traffic.
outbound - resources generated for handling outgoing traffic.
transparent - resources generated for transparent proxy functionality.
prometheus - resources generated when Prometheus metrics are enabled.
direct-access - resources generated for Direct Access functionality.
ingress - resources generated for Zone Ingress.
egress - resources generated for Zone Egress.
gateway - resources generated for MeshGateway.

The list is not complete, because policy plugins can introduce new resources.
For example MeshTrace plugin can create Cluster with "mesh-trace" origin.
- `route_configuration_name` (String) Name of the RouteConfiguration resource to match.





<a id="nestedatt--spec--target_ref"></a>
### Nested Schema for `spec.target_ref`

Read-Only:

- `kind` (String) Kind of the referenced resource
- `labels` (Map of String) Labels are used to select group of MeshServices that match labels. Either Labels or
Name and Namespace can be used.
- `mesh` (String) Mesh is reserved for future use to identify cross mesh resources.
- `name` (String) Name of the referenced resource. Can only be used with kinds: `MeshService`,
`MeshServiceSubset` and `MeshGatewayRoute`
- `namespace` (String) Namespace specifies the namespace of target resource. If empty only resources in policy namespace
will be targeted.
- `proxy_types` (List of String) ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
all data plane types are targeted by the policy.
- `section_name` (String) SectionName is used to target specific section of resource.
For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
- `tags` (Map of String) Tags used to select a subset of proxies by tags. Can only be used with kinds
`MeshSubset` and `MeshServiceSubset`