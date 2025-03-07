resource "konnect_mesh_access_log" "my_meshaccesslog" {
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
          backends = [
            {
              file = {
                format = {
                  json = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  omit_empty_values = false
                  plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                  type              = "Plain"
                }
                path = "/tmp/access.log"
              }
              open_telemetry = {
                attributes = [
                  {
                    key   = "...my_key..."
                    value = "...my_value..."
                  }
                ]
                body     = { "kvlistValue" : { "values" : [{ "key" : "mesh", "value" : { "stringValue" : "%KUMA_MESH%" } }] } }
                endpoint = "otel-collector:4317"
              }
              tcp = {
                address = "127.0.0.1:5000"
                format = {
                  json = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  omit_empty_values = false
                  plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                  type              = "Json"
                }
              }
              type = "Tcp"
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
            "Sidecar"
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
          backends = [
            {
              file = {
                format = {
                  json = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  omit_empty_values = true
                  plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                  type              = "Json"
                }
                path = "/tmp/access.log"
              }
              open_telemetry = {
                attributes = [
                  {
                    key   = "...my_key..."
                    value = "...my_value..."
                  }
                ]
                body     = { "kvlistValue" : { "values" : [{ "key" : "mesh", "value" : { "stringValue" : "%KUMA_MESH%" } }] } }
                endpoint = "otel-collector:4317"
              }
              tcp = {
                address = "127.0.0.1:5000"
                format = {
                  json = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  omit_empty_values = true
                  plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                  type              = "Json"
                }
              }
              type = "OpenTelemetry"
            }
          ]
        }
      }
    ]
    target_ref = {
      kind = "MeshExternalService"
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
    to = [
      {
        default = {
          backends = [
            {
              file = {
                format = {
                  json = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  omit_empty_values = false
                  plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                  type              = "Json"
                }
                path = "/tmp/access.log"
              }
              open_telemetry = {
                attributes = [
                  {
                    key   = "...my_key..."
                    value = "...my_value..."
                  }
                ]
                body     = { "kvlistValue" : { "values" : [{ "key" : "mesh", "value" : { "stringValue" : "%KUMA_MESH%" } }] } }
                endpoint = "otel-collector:4317"
              }
              tcp = {
                address = "127.0.0.1:5000"
                format = {
                  json = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  omit_empty_values = false
                  plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                  type              = "Json"
                }
              }
              type = "Tcp"
            }
          ]
        }
        target_ref = {
          kind = "MeshExternalService"
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
  type = "MeshAccessLog"
}