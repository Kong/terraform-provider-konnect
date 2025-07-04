---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_gateway_hmac_auth Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  GatewayHMACAuth Resource
---

# konnect_gateway_hmac_auth (Resource)

GatewayHMACAuth Resource

## Example Usage

```terraform
resource "konnect_gateway_hmac_auth" "my_gatewayhmacauth" {
  consumer_id      = "f28acbfa-c866-4587-b688-0208ac24df21"
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 9
  id               = "...my_id..."
  secret           = "...my_secret..."
  tags = [
    "..."
  ]
  username = "...my_username..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `consumer_id` (String) Consumer ID for nested entities. Requires replacement if changed.
- `control_plane_id` (String) The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
- `username` (String) Requires replacement if changed.

### Optional

- `created_at` (Number) Unix epoch when the resource was created. Requires replacement if changed.
- `id` (String) Requires replacement if changed.
- `secret` (String) Requires replacement if changed.
- `tags` (List of String) Requires replacement if changed.

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_gateway_hmac_auth.my_konnect_gateway_hmac_auth "{ \"consumer_id\": \"f28acbfa-c866-4587-b688-0208ac24df21\",  \"control_plane_id\": \"9524ec7d-36d9-465d-a8c5-83a3c9390458\",  \"id\": \"70e7b00b-72f2-471b-a5ce-9c4171775360\"}"
```
