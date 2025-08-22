resource "konnect_api_specification" "my_apispecification" {
  api_id  = konnect_api.my_api.id
  content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"
  type    = "oas3"
}
