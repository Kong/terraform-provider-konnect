resource "konnect_portal" "my_portal" {
  name                      = "My Portal for API Product Version"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "tf-example-domain-api-products.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_api_product" "my_apiproduct" {
  name = "My API Product"
  portal_ids = [
    konnect_portal.my_portal.id
  ]
}

resource "konnect_api_product_version" "my_apiproductversion" {
  api_product_id = konnect_api_product.my_apiproduct.id

  name = "v1"
}

resource "konnect_portal_product_version" "my_portalproductversion" {
  application_registration_enabled = false
  auto_approve_registration        = true
  deprecated                       = false
  publish_status                   = "published"
  auth_strategy_ids = [
  ]
  portal_id          = konnect_portal.my_portal.id
  product_version_id = konnect_api_product_version.my_apiproductversion.id
}
