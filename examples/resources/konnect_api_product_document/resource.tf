resource "konnect_api_product_document" "my_apiproductdocument" {
  api_product_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  content        = "## My Markdown"
  metadata = {
    # ...
  }
  parent_document_id = "dd4e1b98-3629-4dd3-acc0-759a726ffee2"
  slug               = "path-for-seo"
  status             = "unpublished"
  title              = "How to create a document in Konnect DocumentHub"
}