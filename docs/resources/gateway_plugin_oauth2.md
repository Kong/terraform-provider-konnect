---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_plugin_oauth2 Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayPluginOauth2 Resource
---

# konnect_gateway_plugin_oauth2 (Resource)

GatewayPluginOauth2 Resource

## Example Usage

```terraform
resource "konnect_gateway_plugin_oauth2" "my_gatewaypluginoauth2" {
  config = {
    accept_http_if_already_terminated = true
    anonymous                         = "...my_anonymous..."
    auth_header_name                  = "...my_auth_header_name..."
    enable_authorization_code         = true
    enable_client_credentials         = false
    enable_implicit_grant             = true
    enable_password_grant             = true
    global_credentials                = true
    hide_credentials                  = false
    mandatory_scope                   = false
    persistent_refresh_token          = true
    pkce                              = "strict"
    provision_key                     = "...my_provision_key..."
    realm                             = "...my_realm..."
    refresh_token_ttl                 = 43385286.36
    reuse_refresh_token               = true
    scopes = [
      "..."
    ]
    token_expiration = 5.01
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 4
  enabled          = true
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
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "wss"
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
  updated_at = 9
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.

### Optional

- `config` (Attributes) (see [below for nested schema](#nestedatt--config))
- `created_at` (Number) Unix epoch when the resource was created.
- `enabled` (Boolean) Whether the plugin is applied.
- `instance_name` (String)
- `ordering` (Attributes) (see [below for nested schema](#nestedatt--ordering))
- `partials` (Attributes List) (see [below for nested schema](#nestedatt--partials))
- `protocols` (Set of String) A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support tcp and tls.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.
- `updated_at` (Number) Unix epoch when the resource was last updated.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `accept_http_if_already_terminated` (Boolean) Accepts HTTPs requests that have already been terminated by a proxy or load balancer.
- `anonymous` (String) An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
- `auth_header_name` (String) The name of the header that is supposed to carry the access token.
- `enable_authorization_code` (Boolean) An optional boolean value to enable the three-legged Authorization Code flow (RFC 6742 Section 4.1).
- `enable_client_credentials` (Boolean) An optional boolean value to enable the Client Credentials Grant flow (RFC 6742 Section 4.4).
- `enable_implicit_grant` (Boolean) An optional boolean value to enable the Implicit Grant flow which allows to provision a token as a result of the authorization process (RFC 6742 Section 4.2).
- `enable_password_grant` (Boolean) An optional boolean value to enable the Resource Owner Password Credentials Grant flow (RFC 6742 Section 4.3).
- `global_credentials` (Boolean) An optional boolean value that allows using the same OAuth credentials generated by the plugin with any other service whose OAuth 2.0 plugin configuration also has `config.global_credentials=true`.
- `hide_credentials` (Boolean) An optional boolean value telling the plugin to show or hide the credential from the upstream service.
- `mandatory_scope` (Boolean) An optional boolean value telling the plugin to require at least one `scope` to be authorized by the end user.
- `persistent_refresh_token` (Boolean)
- `pkce` (String) Specifies a mode of how the Proof Key for Code Exchange (PKCE) should be handled by the plugin. must be one of ["lax", "none", "strict"]
- `provision_key` (String) The unique key the plugin has generated when it has been added to the Service.
- `realm` (String) When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
- `refresh_token_ttl` (Number) Time-to-live value for data
- `reuse_refresh_token` (Boolean) An optional boolean value that indicates whether an OAuth refresh token is reused when refreshing an access token.
- `scopes` (List of String) Describes an array of scope names that will be available to the end user. If `mandatory_scope` is set to `true`, then `scopes` are required.
- `token_expiration` (Number) An optional integer value telling the plugin how many seconds a token should last, after which the client will need to refresh the token. Set to `0` to disable the expiration.


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



<a id="nestedatt--partials"></a>
### Nested Schema for `partials`

Optional:

- `id` (String)
- `name` (String)
- `path` (String)


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
terraform import konnect_gateway_plugin_oauth2.my_konnect_gateway_plugin_oauth2 "{ \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"id\": \"3473c251-5b6c-4f45-b1ff-7ede735a366d\"}"
```
