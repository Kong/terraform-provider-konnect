---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_rate_limit Data Source - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshRateLimit DataSource
---

# konnect_mesh_rate_limit (Data Source)

MeshRateLimit DataSource

## Example Usage

```terraform
data "konnect_mesh_rate_limit" "my_meshratelimit" {
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
- `name` (String) name of the MeshRateLimit

### Read-Only

- `creation_time` (String) Time at which the resource was created
- `labels` (Map of String) The labels to help identity resources
- `modification_time` (String) Time at which the resource was updated
- `spec` (Attributes) Spec is the specification of the Kuma MeshRateLimit resource. (see [below for nested schema](#nestedatt--spec))
- `type` (String) the type of the resource

<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Read-Only:

- `from` (Attributes List) From list makes a match between clients and corresponding configurations (see [below for nested schema](#nestedatt--spec--from))
- `target_ref` (Attributes) TargetRef is a reference to the resource the policy takes an effect on.
The resource could be either a real store object or virtual resource
defined inplace. (see [below for nested schema](#nestedatt--spec--target_ref))
- `to` (Attributes List) To list makes a match between clients and corresponding configurations (see [below for nested schema](#nestedatt--spec--to))

<a id="nestedatt--spec--from"></a>
### Nested Schema for `spec.from`

Read-Only:

- `default` (Attributes) Default is a configuration specific to the group of clients referenced in
'targetRef' (see [below for nested schema](#nestedatt--spec--from--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource that represents a group of
clients. (see [below for nested schema](#nestedatt--spec--from--target_ref))

<a id="nestedatt--spec--from--default"></a>
### Nested Schema for `spec.from.default`

Read-Only:

- `local` (Attributes) LocalConf defines local http or/and tcp rate limit configuration (see [below for nested schema](#nestedatt--spec--from--default--local))

<a id="nestedatt--spec--from--default--local"></a>
### Nested Schema for `spec.from.default.local`

Read-Only:

- `http` (Attributes) LocalHTTP defines configuration of local HTTP rate limiting
https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter (see [below for nested schema](#nestedatt--spec--from--default--local--http))
- `tcp` (Attributes) LocalTCP defines confguration of local TCP rate limiting
https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter (see [below for nested schema](#nestedatt--spec--from--default--local--tcp))

<a id="nestedatt--spec--from--default--local--http"></a>
### Nested Schema for `spec.from.default.local.http`

Read-Only:

- `disabled` (Boolean) Define if rate limiting should be disabled.
- `on_rate_limit` (Attributes) Describes the actions to take on a rate limit event (see [below for nested schema](#nestedatt--spec--from--default--local--http--on_rate_limit))
- `request_rate` (Attributes) Defines how many requests are allowed per interval. (see [below for nested schema](#nestedatt--spec--from--default--local--http--request_rate))

<a id="nestedatt--spec--from--default--local--http--on_rate_limit"></a>
### Nested Schema for `spec.from.default.local.http.on_rate_limit`

Read-Only:

- `headers` (Attributes) The Headers to be added to the HTTP response on a rate limit event (see [below for nested schema](#nestedatt--spec--from--default--local--http--on_rate_limit--headers))
- `status` (Number) The HTTP status code to be set on a rate limit event

<a id="nestedatt--spec--from--default--local--http--on_rate_limit--headers"></a>
### Nested Schema for `spec.from.default.local.http.on_rate_limit.headers`

Read-Only:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--from--default--local--http--on_rate_limit--headers--add))
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--from--default--local--http--on_rate_limit--headers--set))

<a id="nestedatt--spec--from--default--local--http--on_rate_limit--headers--add"></a>
### Nested Schema for `spec.from.default.local.http.on_rate_limit.headers.add`

Read-Only:

- `name` (String)
- `value` (String)


<a id="nestedatt--spec--from--default--local--http--on_rate_limit--headers--set"></a>
### Nested Schema for `spec.from.default.local.http.on_rate_limit.headers.set`

Read-Only:

- `name` (String)
- `value` (String)




<a id="nestedatt--spec--from--default--local--http--request_rate"></a>
### Nested Schema for `spec.from.default.local.http.request_rate`

Read-Only:

- `interval` (String) The interval the number of units is accounted for.
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).



<a id="nestedatt--spec--from--default--local--tcp"></a>
### Nested Schema for `spec.from.default.local.tcp`

Read-Only:

- `connection_rate` (Attributes) Defines how many connections are allowed per interval. (see [below for nested schema](#nestedatt--spec--from--default--local--tcp--connection_rate))
- `disabled` (Boolean) Define if rate limiting should be disabled.
Default: false

<a id="nestedatt--spec--from--default--local--tcp--connection_rate"></a>
### Nested Schema for `spec.from.default.local.tcp.connection_rate`

Read-Only:

- `interval` (String) The interval the number of units is accounted for.
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).





<a id="nestedatt--spec--from--target_ref"></a>
### Nested Schema for `spec.from.target_ref`

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


<a id="nestedatt--spec--to"></a>
### Nested Schema for `spec.to`

Read-Only:

- `default` (Attributes) Default is a configuration specific to the group of clients referenced in
'targetRef' (see [below for nested schema](#nestedatt--spec--to--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource that represents a group of
clients. (see [below for nested schema](#nestedatt--spec--to--target_ref))

<a id="nestedatt--spec--to--default"></a>
### Nested Schema for `spec.to.default`

Read-Only:

- `local` (Attributes) LocalConf defines local http or/and tcp rate limit configuration (see [below for nested schema](#nestedatt--spec--to--default--local))

<a id="nestedatt--spec--to--default--local"></a>
### Nested Schema for `spec.to.default.local`

Read-Only:

- `http` (Attributes) LocalHTTP defines configuration of local HTTP rate limiting
https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter (see [below for nested schema](#nestedatt--spec--to--default--local--http))
- `tcp` (Attributes) LocalTCP defines confguration of local TCP rate limiting
https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter (see [below for nested schema](#nestedatt--spec--to--default--local--tcp))

<a id="nestedatt--spec--to--default--local--http"></a>
### Nested Schema for `spec.to.default.local.http`

Read-Only:

- `disabled` (Boolean) Define if rate limiting should be disabled.
- `on_rate_limit` (Attributes) Describes the actions to take on a rate limit event (see [below for nested schema](#nestedatt--spec--to--default--local--http--on_rate_limit))
- `request_rate` (Attributes) Defines how many requests are allowed per interval. (see [below for nested schema](#nestedatt--spec--to--default--local--http--request_rate))

<a id="nestedatt--spec--to--default--local--http--on_rate_limit"></a>
### Nested Schema for `spec.to.default.local.http.on_rate_limit`

Read-Only:

- `headers` (Attributes) The Headers to be added to the HTTP response on a rate limit event (see [below for nested schema](#nestedatt--spec--to--default--local--http--on_rate_limit--headers))
- `status` (Number) The HTTP status code to be set on a rate limit event

<a id="nestedatt--spec--to--default--local--http--on_rate_limit--headers"></a>
### Nested Schema for `spec.to.default.local.http.on_rate_limit.headers`

Read-Only:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--to--default--local--http--on_rate_limit--headers--add))
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--to--default--local--http--on_rate_limit--headers--set))

<a id="nestedatt--spec--to--default--local--http--on_rate_limit--headers--add"></a>
### Nested Schema for `spec.to.default.local.http.on_rate_limit.headers.add`

Read-Only:

- `name` (String)
- `value` (String)


<a id="nestedatt--spec--to--default--local--http--on_rate_limit--headers--set"></a>
### Nested Schema for `spec.to.default.local.http.on_rate_limit.headers.set`

Read-Only:

- `name` (String)
- `value` (String)




<a id="nestedatt--spec--to--default--local--http--request_rate"></a>
### Nested Schema for `spec.to.default.local.http.request_rate`

Read-Only:

- `interval` (String) The interval the number of units is accounted for.
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).



<a id="nestedatt--spec--to--default--local--tcp"></a>
### Nested Schema for `spec.to.default.local.tcp`

Read-Only:

- `connection_rate` (Attributes) Defines how many connections are allowed per interval. (see [below for nested schema](#nestedatt--spec--to--default--local--tcp--connection_rate))
- `disabled` (Boolean) Define if rate limiting should be disabled.
Default: false

<a id="nestedatt--spec--to--default--local--tcp--connection_rate"></a>
### Nested Schema for `spec.to.default.local.tcp.connection_rate`

Read-Only:

- `interval` (String) The interval the number of units is accounted for.
- `num` (Number) Number of units per interval (depending on usage it can be a number of requests,
or a number of connections).





<a id="nestedatt--spec--to--target_ref"></a>
### Nested Schema for `spec.to.target_ref`

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