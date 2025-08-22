resource "konnect_api" "my_api" {
  attributes  = jsonencode({
    env = ["staging"]
  })
  description = "My API Description"
  name         = "MyAPI"
  slug         = "my-api-v1"
  version      = "v1"
}