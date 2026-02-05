resource "konnect_api_document" "my_apidocument" {
  api_id             = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
  content            = "...my_content..."
  parent_document_id = "b689d9da-f357-4687-8303-ec1c14d44e37"
  slug               = "api-document"
  status             = "published"
  title              = "API Document"
}