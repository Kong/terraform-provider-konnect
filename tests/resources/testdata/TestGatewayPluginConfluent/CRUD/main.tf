resource "konnect_gateway_control_plane" "tfdemo" {
  name         = "Terraform Control Plane - Plugin Confluent"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_confluent" "my_confluent" {
  enabled = true
  config = {
    topic = "kong-tf-test"
    cluster_api_key = "random_key"
    cluster_api_secret = "random_secret"
  }
  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}
