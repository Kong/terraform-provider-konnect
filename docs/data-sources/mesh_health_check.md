---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_health_check Data Source - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshHealthCheck DataSource
---

# konnect_mesh_health_check (Data Source)

MeshHealthCheck DataSource

## Example Usage

```terraform
data "konnect_mesh_health_check" "my_meshhealthcheck" {
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
- `name` (String) name of the MeshHealthCheck

### Read-Only

- `creation_time` (String) Time at which the resource was created
- `labels` (Map of String) The labels to help identity resources
- `modification_time` (String) Time at which the resource was updated
- `spec` (Attributes) Spec is the specification of the Kuma MeshHealthCheck resource. (see [below for nested schema](#nestedatt--spec))
- `type` (String) the type of the resource

<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Read-Only:

- `target_ref` (Attributes) TargetRef is a reference to the resource the policy takes an effect on.
The resource could be either a real store object or virtual resource
defined inplace. (see [below for nested schema](#nestedatt--spec--target_ref))
- `to` (Attributes List) To list makes a match between the consumed services and corresponding configurations (see [below for nested schema](#nestedatt--spec--to))

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

- `default` (Attributes) Default is a configuration specific to the group of destinations referenced in
'targetRef' (see [below for nested schema](#nestedatt--spec--to--default))
- `target_ref` (Attributes) TargetRef is a reference to the resource that represents a group of
destinations. (see [below for nested schema](#nestedatt--spec--to--target_ref))

<a id="nestedatt--spec--to--default"></a>
### Nested Schema for `spec.to.default`

Read-Only:

- `always_log_health_check_failures` (Boolean) If set to true, health check failure events will always be logged. If set
to false, only the initial health check failure event will be logged. The
default value is false.
- `event_log_path` (String) Specifies the path to the file where Envoy can log health check events.
If empty, no event log will be written.
- `fail_traffic_on_panic` (Boolean) If set to true, Envoy will not consider any hosts when the cluster is in
'panic mode'. Instead, the cluster will fail all requests as if all hosts
are unhealthy. This can help avoid potentially overwhelming a failing
service.
- `grpc` (Attributes) GrpcHealthCheck defines gRPC configuration which will instruct the service
the health check will be made for is a gRPC service. (see [below for nested schema](#nestedatt--spec--to--default--grpc))
- `healthy_panic_threshold` (Attributes) Allows to configure panic threshold for Envoy cluster. If not specified,
the default is 50%. To disable panic mode, set to 0%.
Either int or decimal represented as string. (see [below for nested schema](#nestedatt--spec--to--default--healthy_panic_threshold))
- `healthy_threshold` (Number) Number of consecutive healthy checks before considering a host healthy.
- `http` (Attributes) HttpHealthCheck defines HTTP configuration which will instruct the service
the health check will be made for is an HTTP service. (see [below for nested schema](#nestedatt--spec--to--default--http))
- `initial_jitter` (String) If specified, Envoy will start health checking after a random time in
ms between 0 and initialJitter. This only applies to the first health
check.
- `interval` (String) Interval between consecutive health checks.
- `interval_jitter` (String) If specified, during every interval Envoy will add IntervalJitter to the
wait time.
- `interval_jitter_percent` (Number) If specified, during every interval Envoy will add IntervalJitter *
IntervalJitterPercent / 100 to the wait time. If IntervalJitter and
IntervalJitterPercent are both set, both of them will be used to
increase the wait time.
- `no_traffic_interval` (String) The "no traffic interval" is a special health check interval that is used
when a cluster has never had traffic routed to it. This lower interval
allows cluster information to be kept up to date, without sending a
potentially large amount of active health checking traffic for no reason.
Once a cluster has been used for traffic routing, Envoy will shift back
to using the standard health check interval that is defined. Note that
this interval takes precedence over any other. The default value for "no
traffic interval" is 60 seconds.
- `reuse_connection` (Boolean) Reuse health check connection between health checks. Default is true.
- `tcp` (Attributes) TcpHealthCheck defines configuration for specifying bytes to send and
expected response during the health check (see [below for nested schema](#nestedatt--spec--to--default--tcp))
- `timeout` (String) Maximum time to wait for a health check response.
- `unhealthy_threshold` (Number) Number of consecutive unhealthy checks before considering a host
unhealthy.

<a id="nestedatt--spec--to--default--grpc"></a>
### Nested Schema for `spec.to.default.grpc`

Read-Only:

- `authority` (String) The value of the :authority header in the gRPC health check request,
by default name of the cluster this health check is associated with
- `disabled` (Boolean) If true the GrpcHealthCheck is disabled
- `service_name` (String) Service name parameter which will be sent to gRPC service


<a id="nestedatt--spec--to--default--healthy_panic_threshold"></a>
### Nested Schema for `spec.to.default.healthy_panic_threshold`

Read-Only:

- `integer` (Number)
- `str` (String)


<a id="nestedatt--spec--to--default--http"></a>
### Nested Schema for `spec.to.default.http`

Read-Only:

- `disabled` (Boolean) If true the HttpHealthCheck is disabled
- `expected_statuses` (List of Number) List of HTTP response statuses which are considered healthy
- `path` (String) The HTTP path which will be requested during the health check
(ie. /health)
- `request_headers_to_add` (Attributes) The list of HTTP headers which should be added to each health check
request (see [below for nested schema](#nestedatt--spec--to--default--http--request_headers_to_add))

<a id="nestedatt--spec--to--default--http--request_headers_to_add"></a>
### Nested Schema for `spec.to.default.http.request_headers_to_add`

Read-Only:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--to--default--http--request_headers_to_add--add))
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--to--default--http--request_headers_to_add--set))

<a id="nestedatt--spec--to--default--http--request_headers_to_add--add"></a>
### Nested Schema for `spec.to.default.http.request_headers_to_add.add`

Read-Only:

- `name` (String)
- `value` (String)


<a id="nestedatt--spec--to--default--http--request_headers_to_add--set"></a>
### Nested Schema for `spec.to.default.http.request_headers_to_add.set`

Read-Only:

- `name` (String)
- `value` (String)




<a id="nestedatt--spec--to--default--tcp"></a>
### Nested Schema for `spec.to.default.tcp`

Read-Only:

- `disabled` (Boolean) If true the TcpHealthCheck is disabled
- `receive` (List of String) List of Base64 encoded blocks of strings expected as a response. When checking the response,
"fuzzy" matching is performed such that each block must be found, and
in the order specified, but not necessarily contiguous.
If not provided or empty, checks will be performed as "connect only" and be marked as successful when TCP connection is successfully established.
- `send` (String) Base64 encoded content of the message which will be sent during the health check to the target



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