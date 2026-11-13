resource "konnect_api" "my_api" {
  description = "My API Description"
  name         = "MyAPI for Documents"
  slug         = "my-api-v1-document"
  version      = "v1"
}

resource "konnect_api_document" "my_parent_apidocument" {
  api_id             = konnect_api.my_api.id
  content            = "#content"
  slug               = "my-parent-api-document"
  status             = "published"
  title              = "Parent API Document"
}

resource "konnect_api_document" "my_child_apidocument" {
  api_id             = konnect_api.my_api.id
  content            = "#content"
  parent_document_id = konnect_api_document.my_parent_apidocument.id
  slug               = "my-child-api-document"
  status             = "unpublished"
  title              = "Child API Document"
}