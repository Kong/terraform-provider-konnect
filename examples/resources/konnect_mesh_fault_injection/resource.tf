resource "konnect_mesh_fault_injection" "my_meshfaultinjection" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    from = [
      {
        default = {
          http = [
            {
              abort = {
                http_status = 1
                percentage = {
                  integer = 7
                  str     = "...my_str..."
                }
              }
              delay = {
                percentage = {
                  integer = 8
                  str     = "...my_str..."
                }
                value = "...my_value..."
              }
              response_bandwidth = {
                limit = "...my_limit..."
                percentage = {
                  integer = 4
                  str     = "...my_str..."
                }
              }
            }
          ]
        }
        target_ref = {
          kind = "MeshServiceSubset"
          labels = {
            key = "value"
          }
          mesh      = "...my_mesh..."
          name      = "...my_name..."
          namespace = "...my_namespace..."
          proxy_types = [
            "Gateway"
          ]
          section_name = "...my_section_name..."
          tags = {
            key = "value"
          }
        }
      }
    ]
    target_ref = {
      kind = "Dataplane"
      labels = {
        key = "value"
      }
      mesh      = "...my_mesh..."
      name      = "...my_name..."
      namespace = "...my_namespace..."
      proxy_types = [
        "Gateway"
      ]
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
    to = [
      {
        default = {
          http = [
            {
              abort = {
                http_status = 0
                percentage = {
                  integer = 10
                  str     = "...my_str..."
                }
              }
              delay = {
                percentage = {
                  integer = 2
                  str     = "...my_str..."
                }
                value = "...my_value..."
              }
              response_bandwidth = {
                limit = "...my_limit..."
                percentage = {
                  integer = 6
                  str     = "...my_str..."
                }
              }
            }
          ]
        }
        target_ref = {
          kind = "MeshGateway"
          labels = {
            key = "value"
          }
          mesh      = "...my_mesh..."
          name      = "...my_name..."
          namespace = "...my_namespace..."
          proxy_types = [
            "Sidecar"
          ]
          section_name = "...my_section_name..."
          tags = {
            key = "value"
          }
        }
      }
    ]
  }
  type = "MeshFaultInjection"
}