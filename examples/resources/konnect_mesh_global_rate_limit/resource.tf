resource "konnect_mesh_global_rate_limit" "my_meshglobalratelimit" {
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
          backend = {
            rate_limit_service = {
              limit_on_service_fail = true
              timeout               = "...my_timeout..."
              url                   = "...my_url..."
            }
          }
          http = {
            disabled = false
            on_rate_limit = {
              headers = {
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
              status = 10
            }
            ratelimit_on_request = [
              {
                kind = "OnHeader"
                limits = [
                  {
                    request_rate = {
                      interval = "...my_interval..."
                      num      = 7
                    }
                    value = "...my_value..."
                  }
                ]
                name = "...my_name..."
              }
            ]
            request_rate = {
              interval = "...my_interval..."
              num      = 3
            }
          }
          mode = "Limit"
        }
        target_ref = {
          kind = "Mesh"
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
    target_ref = {
      kind = "MeshService"
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
          backend = {
            rate_limit_service = {
              limit_on_service_fail = true
              timeout               = "...my_timeout..."
              url                   = "...my_url..."
            }
          }
          http = {
            disabled = false
            on_rate_limit = {
              headers = {
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
              status = 1
            }
            ratelimit_on_request = [
              {
                kind = "OnHeader"
                limits = [
                  {
                    request_rate = {
                      interval = "...my_interval..."
                      num      = 8
                    }
                    value = "...my_value..."
                  }
                ]
                name = "...my_name..."
              }
            ]
            request_rate = {
              interval = "...my_interval..."
              num      = 8
            }
          }
          mode = "Shadow"
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
  type = "MeshGlobalRateLimit"
}