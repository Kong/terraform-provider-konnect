---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_plugin_oauth2_introspection Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayPluginOauth2Introspection Resource
---

# konnect_gateway_plugin_oauth2_introspection (Resource)

GatewayPluginOauth2Introspection Resource

## Example Usage

```terraform
resource "konnect_gateway_plugin_oauth2_introspection" "my_gatewaypluginoauth2introspection" {
  config = {
    anonymous           = "...my_anonymous..."
    authorization_value = "...my_authorization_value..."
    consumer_by         = "client_id"
    custom_claims_forward = [
      "..."
    ]
    custom_introspection_headers = {
      "see" : jsonencode("documentation"),
    }
    hide_credentials   = true
    introspect_request = false
    introspection_url  = "...my_introspection_url..."
    keepalive          = 2
    run_on_preflight   = true
    timeout            = 8
    token_type_hint    = "...my_token_type_hint..."
    ttl                = 1.29
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  protocols = [
    "grpcs"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.

### Optional

- `consumer` (Attributes) If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer. (see [below for nested schema](#nestedatt--consumer))
- `consumer_group` (Attributes) (see [below for nested schema](#nestedatt--consumer_group))
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (Attributes) (see [below for nested schema](#nestedatt--ordering))
- `protocols` (List of String) A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.

### Read-Only

- `created_at` (Number) Unix epoch when the resource was created.
- `id` (String) The ID of this resource.
- `updated_at` (Number) Unix epoch when the resource was last updated.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `anonymous` (String) An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure `4xx`. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
- `authorization_value` (String) The value to set as the `Authorization` header when querying the introspection endpoint. This depends on the OAuth 2.0 server, but usually is the `client_id` and `client_secret` as a Base64-encoded Basic Auth string (`Basic MG9hNWl...`).
- `consumer_by` (String) A string indicating whether to associate OAuth2 `username` or `client_id` with the consumer's username. OAuth2 `username` is mapped to a consumer's `username` field, while an OAuth2 `client_id` maps to a consumer's `custom_id`. must be one of ["username", "client_id"]
- `custom_claims_forward` (List of String) A list of custom claims to be forwarded from the introspection response to the upstream request. Claims are forwarded in headers with prefix `X-Credential-{claim-name}`.
- `custom_introspection_headers` (Map of String) A list of custom headers to be added in the introspection request.
- `hide_credentials` (Boolean) An optional boolean value telling the plugin to hide the credential to the upstream API server. It will be removed by Kong before proxying the request.
- `introspect_request` (Boolean) A boolean indicating whether to forward information about the current downstream request to the introspect endpoint. If true, headers `X-Request-Path` and `X-Request-Http-Method` will be inserted into the introspect request.
- `introspection_url` (String) A string representing a URL, such as https://example.com/path/to/resource?q=search.
- `keepalive` (Number) An optional value in milliseconds that defines how long an idle connection lives before being closed.
- `run_on_preflight` (Boolean) A boolean value that indicates whether the plugin should run (and try to authenticate) on `OPTIONS` preflight requests. If set to `false`, then `OPTIONS` requests will always be allowed.
- `timeout` (Number) An optional timeout in milliseconds when sending data to the upstream server.
- `token_type_hint` (String) The `token_type_hint` value to associate to introspection requests.
- `ttl` (Number) The TTL in seconds for the introspection response. Set to 0 to disable the expiration.


<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

Optional:

- `id` (String)


<a id="nestedatt--consumer_group"></a>
### Nested Schema for `consumer_group`

Optional:

- `id` (String)


<a id="nestedatt--ordering"></a>
### Nested Schema for `ordering`

Optional:

- `after` (Attributes) (see [below for nested schema](#nestedatt--ordering--after))
- `before` (Attributes) (see [below for nested schema](#nestedatt--ordering--before))

<a id="nestedatt--ordering--after"></a>
### Nested Schema for `ordering.after`

Optional:

- `access` (List of String)


<a id="nestedatt--ordering--before"></a>
### Nested Schema for `ordering.before`

Optional:

- `access` (List of String)



<a id="nestedatt--route"></a>
### Nested Schema for `route`

Optional:

- `id` (String)


<a id="nestedatt--service"></a>
### Nested Schema for `service`

Optional:

- `id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_gateway_plugin_oauth2_introspection.my_konnect_gateway_plugin_oauth2_introspection "{ \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"plugin_id\": \"3473c251-5b6c-4f45-b1ff-7ede735a366d\"}"
```