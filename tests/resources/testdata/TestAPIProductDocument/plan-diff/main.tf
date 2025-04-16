resource "konnect_portal" "my_portal" {
  name                      = "My New Portal with API product document"
  auto_approve_applications = false
  auto_approve_developers   = false
  custom_domain             = "tf-example-domain.example.com"
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

resource "konnect_api_product_document" "my_apiproductdocument" {
  api_product_id = konnect_api_product.my_apiproduct.id
  content        = "YmFzZTY0LWVuY29kZWQgdGV4dCBzdHJpbmc="
  slug               = "path-for-seo"
  status             = "unpublished"
  title              = "My Konnect API Product Document"
  metadata = {
    author = "John Doe"
  }
}