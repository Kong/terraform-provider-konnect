resource "konnect_portal" "my_portal" {
  name                      = "My New Test Portal"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "api-product-spec-test.example.com"
  is_public                 = false
  rbac_enabled              = false
}

resource "konnect_api_product" "my_apiproduct" {
  description = "Text describing the API product"
  labels = {
    key = "value"
  }
  name = "API Product"
  portal_ids = [
    konnect_portal.my_portal.id
  ]
}

resource "konnect_api_product_version" "my_apiproductversion" {
  api_product_id = konnect_api_product.my_apiproduct.id
  name = "v1"
}

resource "konnect_api_product_specification" "my_apiproductspecification" {
  api_product_id         = konnect_api_product.my_apiproduct.id
  api_product_version_id = konnect_api_product_version.my_apiproductversion.id
  content                =  "TXkgWUFNTCBvciBKU09OIGZvcm1hdHRlZCBPQVMgY29udGVudA=="
  name                   = "oas.yaml"
}
