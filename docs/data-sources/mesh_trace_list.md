---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_trace_list Data Source - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshTraceList DataSource
---

# konnect_mesh_trace_list (Data Source)

MeshTraceList DataSource

## Example Usage

```terraform
data "konnect_mesh_trace_list" "my_meshtracelist" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  mesh  = "...my_mesh..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cp_id` (String) Id of the Konnect resource
- `mesh` (String) name of the mesh

### Read-Only

- `items` (Attributes List) (see [below for nested schema](#nestedatt--items))
- `next` (String) URL to the next page
- `total` (Number) The total number of entities

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `creation_time` (String) Time at which the resource was created
- `labels` (Map of String) The labels to help identity resources
- `mesh` (String) Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
- `modification_time` (String) Time at which the resource was updated
- `name` (String) Name of the Kuma resource
- `spec` (Attributes) Spec is the specification of the Kuma MeshTrace resource. (see [below for nested schema](#nestedatt--items--spec))
- `type` (String) the type of the resource

<a id="nestedatt--items--spec"></a>
### Nested Schema for `items.spec`

Read-Only:

- `default` (Attributes) MeshTrace configuration. (see [below for nested schema](#nestedatt--items--spec--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource the policy takes an effect on.
The resource could be either a real store object or virtual resource
defined inplace. (see [below for nested schema](#nestedatt--items--spec--target_ref))

<a id="nestedatt--items--spec--default"></a>
### Nested Schema for `items.spec.default`

Read-Only:

- `backends` (Attributes List) A one element array of backend definition.
Envoy allows configuring only 1 backend, so the natural way of
representing that would be just one object. Unfortunately due to the
reasons explained in MADR 009-tracing-policy this has to be a one element
array for now. (see [below for nested schema](#nestedatt--items--spec--default--backends))
- `sampling` (Attributes) Sampling configuration.
Sampling is the process by which a decision is made on whether to
process/export a span or not. (see [below for nested schema](#nestedatt--items--spec--default--sampling))
- `tags` (Attributes List) Custom tags configuration. You can add custom tags to traces based on
headers or literal values. (see [below for nested schema](#nestedatt--items--spec--default--tags))

<a id="nestedatt--items--spec--default--backends"></a>
### Nested Schema for `items.spec.default.backends`

Read-Only:

- `datadog` (Attributes) Datadog backend configuration. (see [below for nested schema](#nestedatt--items--spec--default--backends--datadog))
- `open_telemetry` (Attributes) OpenTelemetry backend configuration. (see [below for nested schema](#nestedatt--items--spec--default--backends--open_telemetry))
- `type` (String)
- `zipkin` (Attributes) Zipkin backend configuration. (see [below for nested schema](#nestedatt--items--spec--default--backends--zipkin))

<a id="nestedatt--items--spec--default--backends--datadog"></a>
### Nested Schema for `items.spec.default.backends.datadog`

Read-Only:

- `split_service` (Boolean) Determines if datadog service name should be split based on traffic
direction and destination. For example, with `splitService: true` and a
`backend` service that communicates with a couple of databases, you would
get service names like `backend_INBOUND`, `backend_OUTBOUND_db1`, and
`backend_OUTBOUND_db2` in Datadog.
- `url` (String) Address of Datadog collector, only host and port are allowed (no paths,
fragments etc.)


<a id="nestedatt--items--spec--default--backends--open_telemetry"></a>
### Nested Schema for `items.spec.default.backends.open_telemetry`

Read-Only:

- `endpoint` (String) Address of OpenTelemetry collector.


<a id="nestedatt--items--spec--default--backends--zipkin"></a>
### Nested Schema for `items.spec.default.backends.zipkin`

Read-Only:

- `api_version` (String) Version of the API.
https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66
- `shared_span_context` (Boolean) Determines whether client and server spans will share the same span
context.
https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63
- `trace_id128bit` (Boolean) Generate 128bit traces.
- `url` (String) Address of Zipkin collector.



<a id="nestedatt--items--spec--default--sampling"></a>
### Nested Schema for `items.spec.default.sampling`

Read-Only:

- `client` (Attributes) Target percentage of requests that will be force traced if the
'x-client-trace-id' header is set. Mirror of client_sampling in Envoy
https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133
Either int or decimal represented as string. (see [below for nested schema](#nestedatt--items--spec--default--sampling--client))
- `overall` (Attributes) Target percentage of requests will be traced
after all other sampling checks have been applied (client, force tracing,
random sampling). This field functions as an upper limit on the total
configured sampling rate. For instance, setting client to 100
but overall to 1 will result in only 1% of client requests with
the appropriate headers to be force traced. Mirror of
overall_sampling in Envoy
https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150
Either int or decimal represented as string. (see [below for nested schema](#nestedatt--items--spec--default--sampling--overall))
- `random` (Attributes) Target percentage of requests that will be randomly selected for trace
generation, if not requested by the client or not forced.
Mirror of random_sampling in Envoy
https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140
Either int or decimal represented as string. (see [below for nested schema](#nestedatt--items--spec--default--sampling--random))

<a id="nestedatt--items--spec--default--sampling--client"></a>
### Nested Schema for `items.spec.default.sampling.client`

Read-Only:

- `integer` (Number)
- `str` (String)


<a id="nestedatt--items--spec--default--sampling--overall"></a>
### Nested Schema for `items.spec.default.sampling.overall`

Read-Only:

- `integer` (Number)
- `str` (String)


<a id="nestedatt--items--spec--default--sampling--random"></a>
### Nested Schema for `items.spec.default.sampling.random`

Read-Only:

- `integer` (Number)
- `str` (String)



<a id="nestedatt--items--spec--default--tags"></a>
### Nested Schema for `items.spec.default.tags`

Read-Only:

- `header` (Attributes) Tag taken from a header. (see [below for nested schema](#nestedatt--items--spec--default--tags--header))
- `literal` (String) Tag taken from literal value.
- `name` (String) Name of the tag.

<a id="nestedatt--items--spec--default--tags--header"></a>
### Nested Schema for `items.spec.default.tags.header`

Read-Only:

- `default` (String) Default value to use if header is missing.
If the default is missing and there is no value the tag will not be
included.
- `name` (String) Name of the header.




<a id="nestedatt--items--spec--target_ref"></a>
### Nested Schema for `items.spec.target_ref`

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