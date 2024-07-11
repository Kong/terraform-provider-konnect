resource "konnect_api_product_document" "my_apiproductdocument" {
  title          = "How to create a document in Konnect DocumentHub"
  content        = "## How To"
  slug           = "create-a-document"
  status         = "published"
  api_product_id = konnect_api_product.httpbin.id
}

resource "konnect_api_product_document" "my_child_apiproductdocument" {
  title              = "Child Page"
  content            = "## Child Page"
  slug               = "child-page"
  status             = "published"
  api_product_id     = konnect_api_product.httpbin.id
  parent_document_id = konnect_api_product_document.my_apiproductdocument.id
}
