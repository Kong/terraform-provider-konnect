resource "konnect_gateway_control_plane" "tfdemo" {
  name = "Terraform Control Plane"
}

# Clone the built-in ACL plugin and give it a custom name
resource "konnect_gateway_cloned_plugin" "my_custom_acl" {
  name = "custom-acl"
  ref  = "acl"

  control_plane_id = konnect_gateway_control_plane.tfdemo.id
}

# Now you can use the cloned plugin by name
resource "konnect_gateway_custom_plugin" "acl_instance" {
  name = "custom-acl"

  config = {
    allow = ["mygroup"]
  }

  control_plane_id = konnect_gateway_control_plane.tfdemo.id

  depends_on = [konnect_gateway_cloned_plugin.my_custom_acl]
}
