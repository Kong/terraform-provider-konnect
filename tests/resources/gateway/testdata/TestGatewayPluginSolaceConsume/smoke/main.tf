resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For Solace Consume Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_solace_consume" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {
    mode = "AUTO"

    session = {
      host     = "tcp://192.168.88.20:55555"

      authentication = {
        scheme   = "BASIC"
        username = "admin"
        password = "admin"
      }
    }

    flow = {
      binds = [
        {
          name = "queue_name"
      }]
    }
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
