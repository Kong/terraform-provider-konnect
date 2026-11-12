resource "konnect_api" "my_api" {
  description = "My API Description"
  name         = "MyAPI for Specification"
  slug         = "my-api-v1-version"
  version      = "v1"
}

resource "konnect_api_version" "my_apiversion" {
  api_id = konnect_api.my_api.id
  spec = {
    content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"
  }
  version = "1.0.0"
}