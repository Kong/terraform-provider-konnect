resource "konnect_api" "my_api" {
  attributes  = jsonencode({
    env = ["staging"]
  })
  name         = "MyAPI"
  slug         = "my-api-v1"
  version      = "v1"
}