resource "konnect_api_version" "my_apiversion" {
  api_id = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
  spec = {
    content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"
  }
  version = "1.0.0"
}