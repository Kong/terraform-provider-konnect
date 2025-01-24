resource "konnect_mesh_multi_zone_service" "my_meshmultizoneservice" {
  cp_id         = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  creation_time = "0001-01-01T00:00:00Z"
  labels = {
    key = "value",
  }
  mesh              = "...my_mesh..."
  modification_time = "0001-01-01T00:00:00Z"
  name              = "...my_name..."
  spec = {
    ports = [
      {
        app_protocol = "...my_app_protocol..."
        name         = "...my_name..."
        port         = 5
      }
    ]
    selector = {
      mesh_service = {
        match_labels = {
          key = "value",
        }
      }
    }
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
    hostname_generators = [
      {
        conditions = [
          {
            message = "...my_message..."
            reason  = "...my_reason..."
            status  = "False"
            type    = "...my_type..."
          }
        ]
        hostname_generator_ref = {
          core_name = "...my_core_name..."
        }
      }
    ]
    mesh_services = [
      {
        mesh      = "...my_mesh..."
        name      = "...my_name..."
        namespace = "...my_namespace..."
        zone      = "...my_zone..."
      }
    ]
    vips = [
      {
        ip = "...my_ip..."
      }
    ]
  }
  type = "MeshMultiZoneService"
}