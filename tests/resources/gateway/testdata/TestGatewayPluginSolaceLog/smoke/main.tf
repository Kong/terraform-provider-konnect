resource "konnect_gateway_control_plane" "my_konnect_cp" {
  name         = "Terraform Control Plane For Solace Log Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}

resource "konnect_gateway_plugin_solace_log" "my_plugin" {
  enabled       = true
  instance_name = "my-test-plugin"

  config = {

    session = {
      host     = "tcp://192.168.88.20:55555"
      vpn_name = "default"

      authentication = {
        scheme   = "BASIC"
        username = "admin"
        password = "admin"
      }
    }

    message = {
      destinations = [
        {
          name = "test"
          type = "QUEUE"
      }]
      delivery_mode = "PERSISTENT"

      custom_fields_by_lua = {
        broker = "return \"solace\""
      }
    }
  }

  control_plane_id = konnect_gateway_control_plane.my_konnect_cp.id
}
