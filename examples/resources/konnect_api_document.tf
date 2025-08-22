resource "konnect_api_document" "my_apidocument" {
  api_id = konnect_api.my_api.id
  content = "#content"
  slug    = "my-api-document"
  status  = "published"
  title   = "API Document Title"
}
