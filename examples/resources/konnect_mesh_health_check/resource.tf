resource "konnect_mesh_health_check" "my_meshhealthcheck" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "MeshHTTPRoute"
      labels = {
        key = "value",
      }
      mesh      = "...my_mesh..."
      name      = "...my_name..."
      namespace = "...my_namespace..."
      proxy_types = [
        "Sidecar"
      ]
      section_name = "...my_section_name..."
      tags = {
        key = "value",
      }
    }
    to = [
      {
        default = {
          always_log_health_check_failures = true
          event_log_path                   = "...my_event_log_path..."
          fail_traffic_on_panic            = true
          grpc = {
            authority    = "...my_authority..."
            disabled     = true
            service_name = "...my_service_name..."
          }
          healthy_panic_threshold = {
            integer = 6
            str     = "...my_str..."
          }
          healthy_threshold = 1
          http = {
            disabled = true
            expected_statuses = [
              5
            ]
            path = "...my_path..."
            request_headers_to_add = {
              add = [
                {
                  name  = "...my_name..."
                  value = "...my_value..."
                }
              ]
              set = [
                {
                  name  = "...my_name..."
                  value = "...my_value..."
                }
              ]
            }
          }
          initial_jitter          = "...my_initial_jitter..."
          interval                = "...my_interval..."
          interval_jitter         = "...my_interval_jitter..."
          interval_jitter_percent = 0
          no_traffic_interval     = "...my_no_traffic_interval..."
          reuse_connection        = true
          tcp = {
            disabled = true
            receive = [
              "..."
            ]
            send = "...my_send..."
          }
          timeout             = "...my_timeout..."
          unhealthy_threshold = 5
        }
        target_ref = {
          kind = "MeshServiceSubset"
          labels = {
            key = "value",
          }
          mesh      = "...my_mesh..."
          name      = "...my_name..."
          namespace = "...my_namespace..."
          proxy_types = [
            "Sidecar"
          ]
          section_name = "...my_section_name..."
          tags = {
            key = "value",
          }
        }
      }
    ]
  }
  type = "MeshHealthCheck"
}