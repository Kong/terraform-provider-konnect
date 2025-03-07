resource "konnect_mesh_rate_limit" "my_meshratelimit" {
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
          local = {
            http = {
              disabled = true
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
                status = 7
              }
              request_rate = {
                interval = "...my_interval..."
                num      = 6
              }
            }
            tcp = {
              connection_rate = {
                interval = "...my_interval..."
                num      = 1
              }
              disabled = true
            }
          }
        }
        target_ref = {
          kind = "MeshHTTPRoute"
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
    rules = [
      {
        default = {
          local = {
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
                status = 5
              }
              request_rate = {
                interval = "...my_interval..."
                num      = 9
              }
            }
            tcp = {
              connection_rate = {
                interval = "...my_interval..."
                num      = 0
              }
              disabled = false
            }
          }
        }
      }
    ]
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
    to = [
      {
        default = {
          local = {
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
                status = 9
              }
              request_rate = {
                interval = "...my_interval..."
                num      = 10
              }
            }
            tcp = {
              connection_rate = {
                interval = "...my_interval..."
                num      = 5
              }
              disabled = false
            }
          }
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
  }
  type = "MeshRateLimit"
}