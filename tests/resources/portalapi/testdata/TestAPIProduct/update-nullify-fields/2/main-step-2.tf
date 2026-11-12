resource "konnect_portal_classic" "my_portal" {
  name                      = "APIProductNullPortal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "api-product-null-test.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_api_product" "my_null_desc_apiproduct" {
  name = "API Product for null description"
  portal_ids = [
    konnect_portal_classic.my_portal.id
  ]
}