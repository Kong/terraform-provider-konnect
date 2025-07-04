---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_plugin_vault_auth Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayPluginVaultAuth Resource
---

# konnect_gateway_plugin_vault_auth (Resource)

GatewayPluginVaultAuth Resource

## Example Usage

```terraform
resource "konnect_gateway_plugin_vault_auth" "my_gatewaypluginvaultauth" {
  config = {
    access_token_name = "...my_access_token_name..."
    anonymous         = "...my_anonymous..."
    hide_credentials  = false
    run_on_preflight  = false
    secret_token_name = "...my_secret_token_name..."
    tokens_in_body    = true
    vault             = "...my_vault..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 2
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
  updated_at = 1
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
- `protocols` (Set of String) A set of strings representing HTTP protocols.
- `route` (Attributes) If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used. (see [below for nested schema](#nestedatt--route))
- `service` (Attributes) If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched. (see [below for nested schema](#nestedatt--service))
- `tags` (List of String) An optional set of strings associated with the Plugin for grouping and filtering.
- `updated_at` (Number) Unix epoch when the resource was last updated.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `access_token_name` (String) Describes an array of comma-separated parameter names where the plugin looks for an access token. The client must send the access token in one of those key names, and the plugin will try to read the credential from a header or the querystring parameter with the same name. The key names can only contain [a-z], [A-Z], [0-9], [_], and [-].
- `anonymous` (String) An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure `4xx`. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
- `hide_credentials` (Boolean) An optional boolean value telling the plugin to show or hide the credential from the upstream service. If `true`, the plugin will strip the credential from the request (i.e. the header or querystring containing the key) before proxying it.
- `run_on_preflight` (Boolean) A boolean value that indicates whether the plugin should run (and try to authenticate) on `OPTIONS` preflight requests. If set to `false`, then `OPTIONS` requests will always be allowed.
- `secret_token_name` (String) Describes an array of comma-separated parameter names where the plugin looks for a secret token. The client must send the secret in one of those key names, and the plugin will try to read the credential from a header or the querystring parameter with the same name. The key names can only contain [a-z], [A-Z], [0-9], [_], and [-].
- `tokens_in_body` (Boolean) If enabled, the plugin will read the request body (if said request has one and its MIME type is supported) and try to find the key in it. Supported MIME types are `application/www-form-urlencoded`, `application/json`, and `multipart/form-data`.
- `vault` (String) A reference to an existing `vault` object within the database. `vault` entities define the connection and authentication parameters used to connect to a Vault HTTP(S) API.


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
terraform import konnect_gateway_plugin_vault_auth.my_konnect_gateway_plugin_vault_auth "{ \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"id\": \"3473c251-5b6c-4f45-b1ff-7ede735a366d\"}"
```
