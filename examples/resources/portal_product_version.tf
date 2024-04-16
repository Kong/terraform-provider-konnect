resource "konnect_portal_product_version" "my_portalproductversion" {
  application_registration_enabled = true
  auto_approve_registration        = true
  deprecated                       = false
  publish_status                   = "published"

  portal_id          = data.konnect_portal_list.my_portallist.data[0].id
  product_version_id = konnect_api_product_version.httpbin_v1.id
  auth_strategy_ids = [
    konnect_application_auth_strategy.my_applicationauthstrategy.id
  ]
}
