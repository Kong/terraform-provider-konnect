resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Demo cp For Datakit Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}


resource "konnect_gateway_plugin_datakit" "my_datakit" {
  enabled = true

  config = {
    nodes = [
      { static = {
        name = "STATIC_INPUTS"
        type = "static"

        values = {

          headers = jsonencode({
            Content-Type = "application/x-www-form-urlencoded"
          })
        }
      }}, 
    ]
  }
  tags = []

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
