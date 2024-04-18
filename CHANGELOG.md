# Changelog

## 0.2.0
> Released on 2024/04/18

### BREAKING CHANGES
* The provider `server_url` no longer contains an API version. Update your provider configuration to use `https://us.api.konghq.com` (or your chosen region)
* `labels` on the `konnect_gateway_control_plane` resource now accept a map of values rather than using `jsonencode()`

### Fixes
* Mesh control planes can be created in regions other than NA
* Gateway control planes now accept `CLUSTER_TYPE_CONTROL_PLANE` in the `cluster_type` field

## 0.1.0
> Released on 2024/04/16

* Initial release