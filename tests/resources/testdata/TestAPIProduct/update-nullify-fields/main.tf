resource "konnect_portal" "my_portal" {
  name                      = "APIProductForNullTest"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "api-product-null-test.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_api_product" "my_null_desc_apiproduct" {
  description = "API product description"

  name = "API Product for null description"
  portal_ids = [
    konnect_portal.my_portal.id
  ]
}