resource "konnect_api" "my_api" {
  description = "My API Description"
  name        = "MyAPI"
  slug        = "my-api-v1-implementation"
  version     = "v1"
  attributes = jsonencode({
    environment = ["staging"],
  })
}
