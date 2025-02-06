resource "konnect_mesh_http_route" "my_meshhttproute" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "Dataplane"
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
        hostnames = [
          "..."
        ]
        rules = [
          {
            default = {
              backend_refs = [
                {
                  kind = "Dataplane"
                  labels = {
                    key = "value",
                  }
                  mesh      = "...my_mesh..."
                  name      = "...my_name..."
                  namespace = "...my_namespace..."
                  port      = 10
                  proxy_types = [
                    "Gateway"
                  ]
                  section_name = "...my_section_name..."
                  tags = {
                    key = "value",
                  }
                  weight = 5
                }
              ]
              filters = [
                {
                  request_header_modifier = {
                    add = [
                      {
                        name  = "...my_name..."
                        value = "...my_value..."
                      }
                    ]
                    remove = [
                      "..."
                    ]
                    set = [
                      {
                        name  = "...my_name..."
                        value = "...my_value..."
                      }
                    ]
                  }
                  request_mirror = {
                    backend_ref = {
                      kind = "Mesh"
                      labels = {
                        key = "value",
                      }
                      mesh      = "...my_mesh..."
                      name      = "...my_name..."
                      namespace = "...my_namespace..."
                      port      = 4
                      proxy_types = [
                        "Sidecar"
                      ]
                      section_name = "...my_section_name..."
                      tags = {
                        key = "value",
                      }
                      weight = 2
                    }
                    percentage = {
                      integer = 0
                      str     = "...my_str..."
                    }
                  }
                  request_redirect = {
                    hostname = "...my_hostname..."
                    path = {
                      replace_full_path    = "...my_replace_full_path..."
                      replace_prefix_match = "...my_replace_prefix_match..."
                      type                 = "ReplaceFullPath"
                    }
                    port        = 46600
                    scheme      = "https"
                    status_code = 308
                  }
                  response_header_modifier = {
                    add = [
                      {
                        name  = "...my_name..."
                        value = "...my_value..."
                      }
                    ]
                    remove = [
                      "..."
                    ]
                    set = [
                      {
                        name  = "...my_name..."
                        value = "...my_value..."
                      }
                    ]
                  }
                  type = "URLRewrite"
                  url_rewrite = {
                    host_to_backend_hostname = false
                    hostname                 = "...my_hostname..."
                    path = {
                      replace_full_path    = "...my_replace_full_path..."
                      replace_prefix_match = "...my_replace_prefix_match..."
                      type                 = "ReplacePrefixMatch"
                    }
                  }
                }
              ]
            }
            matches = [
              {
                headers = [
                  {
                    name  = "...my_name..."
                    type  = "Absent"
                    value = "...my_value..."
                  }
                ]
                method = "PATCH"
                path = {
                  type  = "Exact"
                  value = "...my_value..."
                }
                query_params = [
                  {
                    name  = "...my_name..."
                    type  = "Exact"
                    value = "...my_value..."
                  }
                ]
              }
            ]
          }
        ]
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
      }
    ]
  }
  type = "MeshHTTPRoute"
}