resource "konnect_gateway_control_plane" "plugin_acl_cp" {
  name         = "Terraform Control Plane For ACL Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_acl" "my_acl" {
  enabled = true

  config = {
    include_consumer_groups = true
    allow = ["dev", "admin"]
    deny = ["test"]
  }

  control_plane_id = konnect_gateway_control_plane.plugin_acl_cp.id
}
