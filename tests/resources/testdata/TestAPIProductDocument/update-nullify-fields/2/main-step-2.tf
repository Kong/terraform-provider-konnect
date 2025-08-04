resource "konnect_portal" "my_portal" {
  name                      = "My New Portal with API product document Null Test"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "api-product-doc-null-test.example.com"
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

resource "konnect_api_product_document" "my_apiproduct_parent_document" {
  api_product_id = konnect_api_product.my_apiproduct.id
  content        = "YmFzZTY0LWVuY29kZWQgdGV4dCBzdHJpbmc="
  slug               = "path-for-seo"
  status             = "unpublished"
  title              = "My Konnect API Product Parent Document"
  metadata = {
    author = "John Doe"
  }
}

resource "konnect_api_product_document" "my_apiproduct_child_document" {
  api_product_id = konnect_api_product.my_apiproduct.id
  content        = "YmFzZTY0LWVuY29kZWQgdGV4dCBzdHJpbmc="
  slug               = "path-for-seo"
  status             = "unpublished"
  title              = "My Konnect API Product Child Document"
  metadata = {
    author = "John Doe"
  }
}