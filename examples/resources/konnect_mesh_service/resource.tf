resource "konnect_mesh_service" "my_meshservice" {
  cp_id         = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  creation_time = "0001-01-01T00:00:00Z"
  labels = {
    key = "value",
  }
  mesh              = "...my_mesh..."
  modification_time = "0001-01-01T00:00:00Z"
  name              = "...my_name..."
  spec = {
    identities = [
      {
        type  = "ServiceTag"
        value = "...my_value..."
      }
    ]
    ports = [
      {
        app_protocol = "...my_app_protocol..."
        name         = "...my_name..."
        port         = 8
        target_port = {
          integer = 7
          str     = "...my_str..."
        }
      }
    ]
    selector = {
      dataplane_ref = {
        name = "...my_name..."
      }
      dataplane_tags = {
        key = "value",
      }
    }
    state = "Available"
  }
  status = {
    addresses = [
      {
        hostname = "...my_hostname..."
        hostname_generator_ref = {
          core_name = "...my_core_name..."
        }
        origin = "...my_origin..."
      }
    ]
    dataplane_proxies = {
      connected = 8
      healthy   = 6
      total     = 2
    }
    hostname_generators = [
      {
        conditions = [
          {
            message = "...my_message..."
            reason  = "...my_reason..."
            status  = "True"
            type    = "...my_type..."
          }
        ]
        hostname_generator_ref = {
          core_name = "...my_core_name..."
        }
      }
    ]
    tls = {
      status = "Ready"
    }
    vips = [
      {
        ip = "...my_ip..."
      }
    ]
  }
  type = "MeshService"
}