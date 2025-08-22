resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "My Control Plane"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_service" "httpbin" {
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

resource "konnect_api" "my_api" {
  description = "My API Description"
  name         = "MyAPI"
  slug         = "my-api-v1-implementation"
  version      = "v1"
}

resource "konnect_api_implementation" "my_apiimplementation" {
  api_id = konnect_api.my_api.id
  service_reference = {
    service = {
      control_plane_id = konnect_gateway_control_plane.tfdemo.id
      id               = konnect_gateway_service.httpbin.id
    }
  }
}

resource "konnect_api_document" "my_apidocument" {
  api_id             = konnect_api.my_api.id
  content            = "#content"
  slug               = "my-api-document"
  status             = "published"
  title              = "API Document Title"
}

resource "konnect_api_specification" "my_apispecification" {
  api_id  = konnect_api.my_api.id
  content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"
  type    = "oas3"
}

resource "konnect_api_version" "my_apiversion" {
  api_id = konnect_api.my_api.id
  spec = {
    content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"
  }
  version = "1.0.0"
}

resource "konnect_portal" "my_portal_api_publication" {
  force_destroy                        = "true"
  name         = "My portal for api publication"
}

resource "konnect_api_publication" "my_apipublication" {
  api_id = konnect_api.my_api.id
  auto_approve_registrations = false
  portal_id                  = konnect_portal.my_portal_api_publication.id
  visibility                 = "private"
}
