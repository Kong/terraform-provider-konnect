# Changelog

## 0.3.1
> Released on 2024/06/14

### Bug Fixes
* Ensure that `konnect_gateway_vault.config` is sent via the provider

## 0.3.0
> Released on 2024/06/14

### Features
* Update all schemas to Kong Gateway 3.7.x

### Bug Fixes
* Do not send default values from the provider as they'll be applied on the server automatically
  * This allows terraform-provider-konnect 0.3.0 to be used with Kong Gateway 3.6 and below

## 0.2.4
> Released on 2024/05/31

### Features
* Allow users to set provider attributes via environment variables
  * `personal_access_token` = `KONNECT_TOKEN`
  * `system_account_access_token` = `KONNECT_SPAT`
  * `server_url` = `KONNECT_SERVER_URL`

## 0.2.3
> Released on 2024/05/14

### Features
* Add support for the `konnect_cloud_gateway_configuration`, `konnect_cloud_gateway_custom_domain`, `konnect_cloud_gateway_network` and `konnect_cloud_gateway_transit_gateway` resources

### Bug Fixes
* Fix `konnect_portal_product_version` creation bug

## 0.2.2
> Released on 2024/05/04

### Features
* Add support for the `konnect_system_account`, `konnect_system_account_access_token`, `konnect_system_account_role` and `konnect_system_account_team` resources

## 0.2.1
> Released on 2024/04/26

### Features
* Add support for the `konnect_gateway_data_plane_client_certificate` resource

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