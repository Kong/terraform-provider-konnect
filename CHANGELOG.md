# Changelog

## 1.0.0
> Released on 2024/09/11

### Features

* Released `terraform-provider-konnect` as GA 🎉

## 0.7.0
> Released on 2024/09/11

### Features
* Add the `konnect_gateway_custom_plugin` resource for managing non-bundled plugins
* Add support for all bundled Kong plugins
* Add support for `labels` to `konnect_api_product_version`
* Add support for `ordering` to `konnect_gateway_plugin`

### Bug fixes
* Force recreation of `konnect_gateway_*` entities if `control_plane_id` changes
* `proxy_urls` on `konnect_control_plane` and `dataplane_groups` on `konnect_cloud_gateway` are now marked as sets rather than list. This makes `terraform diff` ignore the order of items in the list.

## 0.6.3
> Released on 2024/08/25

### Features
* Add support for `labels` to `application_auth_strategies`

### Bug fixes
* The `cluster_cert_key` in `serverless_cloud_gateway` is now marked as `sensitive`

## 0.6.2
> Released on 2024/07/24

### Features
* Add the ability for the Kong team to use the provider against development environments

### Bug fixes
* Generate up to date documentation

## 0.6.1
> Released on 2024/07/23

### Features
* Add list of `enums` for available roles when using the Identity API
* Add pattern validation based on OpenAPI spec when running `terraform validate`

### Bug Fixes
* Fixed `terraform import` for the `konnect_portal` resource

## 0.6.0
> Released on 2024/07/11

### BREAKING CHANGES
* The `konnect_portal` resource can no longer be used to adopt the existing default portal without first running `terraform import`

### Features
* Add support for the  `konnect_api_product_document`, `konnect_portal_appearance`,  `konnect_gateway_plugin_pre_function`, `konnect_gateway_plugin_post_function` and `konnect_gateway_plugin_statsd` resources

## 0.5.1
> Released on 2024/07/03

### Features
* Add `tfdocs` for all existing resources

## 0.5.0
> Released on 2024/07/03

### Features
* Add support for the `konnect_serverless_cloud_gateway`, `konnect_gateway_plugin_exit_transformer`, `konnect_gateway_consumer_group` and `konnect_gateway_consumer_group_member` resources

### Bug Fixes
* `konnect_gateway_*` resources are now updated in place rather than any change causing the entity to be recreated

## 0.4.0
> Released on 2024/06/24

### Features
* Add support for the `konnect_api_product_specification`, `konnect_gateway_custom_plugin_schema`, `konnect_team`, `konnect_team_role` and `konnect_team_user` resources

### Bug Fixes
* Make Cloud Gateways + Identity APIs point at `global.api.konghq.com` rather than the provided `server_url`

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