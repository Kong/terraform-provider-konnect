resource "konnect_mesh_metric" "my_meshmetric" {
  cp_id         = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  creation_time = "0001-01-01T00:00:00Z"
  labels = {
    key = "value",
  }
  mesh              = "...my_mesh..."
  modification_time = "0001-01-01T00:00:00Z"
  name              = "...my_name..."
  spec = {
    default = {
      applications = [
        {
          address = "...my_address..."
          name    = "...my_name..."
          path    = "...my_path..."
          port    = 8
        }
      ]
      backends = [
        {
          open_telemetry = {
            endpoint         = "...my_endpoint..."
            refresh_interval = "...my_refresh_interval..."
          }
          prometheus = {
            client_id = "...my_client_id..."
            path      = "...my_path..."
            port      = 7
            tls = {
              mode = "Disabled"
            }
          }
          type = "Prometheus"
        }
      ]
      sidecar = {
        include_unused = false
        profiles = {
          append_profiles = [
            {
              name = "Basic"
            }
          ]
          exclude = [
            {
              match = "...my_match..."
              type  = "Contains"
            }
          ]
          include = [
            {
              match = "...my_match..."
              type  = "Regex"
            }
          ]
        }
      }
    }
    target_ref = {
      kind = "MeshSubset"
      labels = {
        key = "value",
      }
      mesh      = "...my_mesh..."
      name      = "...my_name..."
      namespace = "...my_namespace..."
      proxy_types = [
        "Gateway"
      ]
      section_name = "...my_section_name..."
      tags = {
        key = "value",
      }
    }
  }
  type = "MeshMetric"
}