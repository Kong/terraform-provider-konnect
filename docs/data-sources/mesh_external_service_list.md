---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_mesh_external_service_list Data Source - terraform-provider-konnect"
subcategory: ""
description: |-
  MeshExternalServiceList DataSource
---

# konnect_mesh_external_service_list (Data Source)

MeshExternalServiceList DataSource

## Example Usage

```terraform
data "konnect_mesh_external_service_list" "my_meshexternalservicelist" {
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
- `spec` (Attributes) Spec is the specification of the Kuma MeshExternalService resource. (see [below for nested schema](#nestedatt--items--spec))
- `status` (Attributes) Status is the current status of the Kuma MeshExternalService resource. (see [below for nested schema](#nestedatt--items--status))
- `type` (String) the type of the resource

<a id="nestedatt--items--spec"></a>
### Nested Schema for `items.spec`

Read-Only:

- `endpoints` (Attributes List) Endpoints defines a list of destinations to send traffic to. (see [below for nested schema](#nestedatt--items--spec--endpoints))
- `extension` (Attributes) Extension struct for a plugin configuration, in the presence of an extension `endpoints` and `tls` are not required anymore - it's up to the extension to validate them independently. (see [below for nested schema](#nestedatt--items--spec--extension))
- `match` (Attributes) Match defines traffic that should be routed through the sidecar. (see [below for nested schema](#nestedatt--items--spec--match))
- `tls` (Attributes) Tls provides a TLS configuration when proxy is resposible for a TLS origination (see [below for nested schema](#nestedatt--items--spec--tls))

<a id="nestedatt--items--spec--endpoints"></a>
### Nested Schema for `items.spec.endpoints`

Read-Only:

- `address` (String) Address defines an address to which a user want to send a request. Is possible to provide `domain`, `ip`.
- `port` (Number) Port of the endpoint


<a id="nestedatt--items--spec--extension"></a>
### Nested Schema for `items.spec.extension`

Read-Only:

- `config` (String) Config freeform configuration for the extension. Parsed as JSON.
- `type` (String) Type of the extension.


<a id="nestedatt--items--spec--match"></a>
### Nested Schema for `items.spec.match`

Read-Only:

- `port` (Number) Port defines a port to which a user does request.
- `protocol` (String) Protocol defines a protocol of the communication. Possible values: `tcp`, `grpc`, `http`, `http2`.
- `type` (String) Type of the match, only `HostnameGenerator` is available at the moment.


<a id="nestedatt--items--spec--tls"></a>
### Nested Schema for `items.spec.tls`

Read-Only:

- `allow_renegotiation` (Boolean) AllowRenegotiation defines if TLS sessions will allow renegotiation.
Setting this to true is not recommended for security reasons.
- `enabled` (Boolean) Enabled defines if proxy should originate TLS.
- `verification` (Attributes) Verification section for providing TLS verification details. (see [below for nested schema](#nestedatt--items--spec--tls--verification))
- `version` (Attributes) Version section for providing version specification. (see [below for nested schema](#nestedatt--items--spec--tls--version))

<a id="nestedatt--items--spec--tls--verification"></a>
### Nested Schema for `items.spec.tls.verification`

Read-Only:

- `ca_cert` (Attributes) CaCert defines a certificate of CA. (see [below for nested schema](#nestedatt--items--spec--tls--verification--ca_cert))
- `client_cert` (Attributes) ClientCert defines a certificate of a client. (see [below for nested schema](#nestedatt--items--spec--tls--verification--client_cert))
- `client_key` (Attributes) ClientKey defines a client private key. (see [below for nested schema](#nestedatt--items--spec--tls--verification--client_key))
- `mode` (String) Mode defines if proxy should skip verification, one of `SkipSAN`, `SkipCA`, `Secured`, `SkipAll`. Default `Secured`.
- `server_name` (String) ServerName overrides the default Server Name Indicator set by Kuma.
- `subject_alt_names` (Attributes List) SubjectAltNames list of names to verify in the certificate. (see [below for nested schema](#nestedatt--items--spec--tls--verification--subject_alt_names))

<a id="nestedatt--items--spec--tls--verification--ca_cert"></a>
### Nested Schema for `items.spec.tls.verification.ca_cert`

Read-Only:

- `inline` (String) Data source is inline bytes.
- `inline_string` (String) Data source is inline string`
- `secret` (String) Data source is a secret with given Secret key.


<a id="nestedatt--items--spec--tls--verification--client_cert"></a>
### Nested Schema for `items.spec.tls.verification.client_cert`

Read-Only:

- `inline` (String) Data source is inline bytes.
- `inline_string` (String) Data source is inline string`
- `secret` (String) Data source is a secret with given Secret key.


<a id="nestedatt--items--spec--tls--verification--client_key"></a>
### Nested Schema for `items.spec.tls.verification.client_key`

Read-Only:

- `inline` (String) Data source is inline bytes.
- `inline_string` (String) Data source is inline string`
- `secret` (String) Data source is a secret with given Secret key.


<a id="nestedatt--items--spec--tls--verification--subject_alt_names"></a>
### Nested Schema for `items.spec.tls.verification.subject_alt_names`

Read-Only:

- `type` (String) Type specifies matching type, one of `Exact`, `Prefix`. Default: `Exact`
- `value` (String) Value to match.



<a id="nestedatt--items--spec--tls--version"></a>
### Nested Schema for `items.spec.tls.version`

Read-Only:

- `max` (String) Max defines maximum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
- `min` (String) Min defines minimum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.




<a id="nestedatt--items--status"></a>
### Nested Schema for `items.status`

Read-Only:

- `addresses` (Attributes List) Addresses section for generated domains (see [below for nested schema](#nestedatt--items--status--addresses))
- `hostname_generators` (Attributes List) (see [below for nested schema](#nestedatt--items--status--hostname_generators))
- `vip` (Attributes) Vip section for allocated IP (see [below for nested schema](#nestedatt--items--status--vip))

<a id="nestedatt--items--status--addresses"></a>
### Nested Schema for `items.status.addresses`

Read-Only:

- `hostname` (String)
- `hostname_generator_ref` (Attributes) (see [below for nested schema](#nestedatt--items--status--addresses--hostname_generator_ref))
- `origin` (String)

<a id="nestedatt--items--status--addresses--hostname_generator_ref"></a>
### Nested Schema for `items.status.addresses.hostname_generator_ref`

Read-Only:

- `core_name` (String)



<a id="nestedatt--items--status--hostname_generators"></a>
### Nested Schema for `items.status.hostname_generators`

Read-Only:

- `conditions` (Attributes List) Conditions is an array of hostname generator conditions. (see [below for nested schema](#nestedatt--items--status--hostname_generators--conditions))
- `hostname_generator_ref` (Attributes) (see [below for nested schema](#nestedatt--items--status--hostname_generators--hostname_generator_ref))

<a id="nestedatt--items--status--hostname_generators--conditions"></a>
### Nested Schema for `items.status.hostname_generators.conditions`

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


<a id="nestedatt--items--status--hostname_generators--hostname_generator_ref"></a>
### Nested Schema for `items.status.hostname_generators.hostname_generator_ref`

Read-Only:

- `core_name` (String)



<a id="nestedatt--items--status--vip"></a>
### Nested Schema for `items.status.vip`

Read-Only:

- `ip` (String) Value allocated IP for a provided domain with `HostnameGenerator` type in a match section.