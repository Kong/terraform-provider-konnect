resource "konnect_mesh_retry" "my_meshretry" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "MeshServiceSubset"
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
          grpc = {
            back_off = {
              base_interval = "...my_base_interval..."
              max_interval  = "...my_max_interval..."
            }
            num_retries     = 2
            per_try_timeout = "...my_per_try_timeout..."
            rate_limited_back_off = {
              max_interval = "...my_max_interval..."
              reset_headers = [
                {
                  format = "UnixTimestamp"
                  name   = "...my_name..."
                }
              ]
            }
            retry_on = [
              "DeadlineExceeded"
            ]
          }
          http = {
            back_off = {
              base_interval = "...my_base_interval..."
              max_interval  = "...my_max_interval..."
            }
            host_selection = [
              {
                predicate = "OmitPreviousHosts"
                tags = {
                  key = "value",
                }
                update_frequency = 6
              }
            ]
            host_selection_max_attempts = 1
            num_retries                 = 6
            per_try_timeout             = "...my_per_try_timeout..."
            rate_limited_back_off = {
              max_interval = "...my_max_interval..."
              reset_headers = [
                {
                  format = "Seconds"
                  name   = "...my_name..."
                }
              ]
            }
            retriable_request_headers = [
              {
                name  = "...my_name..."
                type  = "RegularExpression"
                value = "...my_value..."
              }
            ]
            retriable_response_headers = [
              {
                name  = "...my_name..."
                type  = "RegularExpression"
                value = "...my_value..."
              }
            ]
            retry_on = [
              "..."
            ]
          }
          tcp = {
            max_connect_attempt = 1
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
            "Gateway"
          ]
          section_name = "...my_section_name..."
          tags = {
            key = "value",
          }
        }
      }
    ]
  }
  type = "MeshRetry"
}