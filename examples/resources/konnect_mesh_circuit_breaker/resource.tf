resource "konnect_mesh_circuit_breaker" "my_meshcircuitbreaker" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    from = [
      {
        default = {
          connection_limits = {
            max_connection_pools = 9
            max_connections      = 1
            max_pending_requests = 3
            max_requests         = 7
            max_retries          = 3
          }
          outlier_detection = {
            base_ejection_time = "...my_base_ejection_time..."
            detectors = {
              failure_percentage = {
                minimum_hosts  = 7
                request_volume = 5
                threshold      = 1
              }
              gateway_failures = {
                consecutive = 2
              }
              local_origin_failures = {
                consecutive = 0
              }
              success_rate = {
                minimum_hosts  = 9
                request_volume = 6
                standard_deviation_factor = {
                  integer = 8
                  str     = "...my_str..."
                }
              }
              total_failures = {
                consecutive = 7
              }
            }
            disabled                        = false
            interval                        = "...my_interval..."
            max_ejection_percent            = 9
            split_external_and_local_errors = true
          }
        }
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
      }
    ]
    rules = [
      {
        default = {
          connection_limits = {
            max_connection_pools = 6
            max_connections      = 10
            max_pending_requests = 0
            max_requests         = 0
            max_retries          = 5
          }
          outlier_detection = {
            base_ejection_time = "...my_base_ejection_time..."
            detectors = {
              failure_percentage = {
                minimum_hosts  = 6
                request_volume = 1
                threshold      = 6
              }
              gateway_failures = {
                consecutive = 0
              }
              local_origin_failures = {
                consecutive = 7
              }
              success_rate = {
                minimum_hosts  = 8
                request_volume = 9
                standard_deviation_factor = {
                  integer = 6
                  str     = "...my_str..."
                }
              }
              total_failures = {
                consecutive = 9
              }
            }
            disabled                        = false
            interval                        = "...my_interval..."
            max_ejection_percent            = 1
            split_external_and_local_errors = false
          }
        }
      }
    ]
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
    to = [
      {
        default = {
          connection_limits = {
            max_connection_pools = 4
            max_connections      = 8
            max_pending_requests = 3
            max_requests         = 2
            max_retries          = 3
          }
          outlier_detection = {
            base_ejection_time = "...my_base_ejection_time..."
            detectors = {
              failure_percentage = {
                minimum_hosts  = 1
                request_volume = 10
                threshold      = 0
              }
              gateway_failures = {
                consecutive = 10
              }
              local_origin_failures = {
                consecutive = 5
              }
              success_rate = {
                minimum_hosts  = 9
                request_volume = 3
                standard_deviation_factor = {
                  integer = 10
                  str     = "...my_str..."
                }
              }
              total_failures = {
                consecutive = 2
              }
            }
            disabled                        = true
            interval                        = "...my_interval..."
            max_ejection_percent            = 4
            split_external_and_local_errors = true
          }
        }
        target_ref = {
          kind = "MeshGateway"
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
  type = "MeshCircuitBreaker"
}