resource "konnect_mesh_tcp_route" "my_meshtcproute" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "Mesh"
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
        rules = [
          {
            default = {
              backend_refs = [
                {
                  kind = "MeshService"
                  labels = {
                    key = "value"
                  }
                  mesh      = "...my_mesh..."
                  name      = "...my_name..."
                  namespace = "...my_namespace..."
                  port      = 6
                  proxy_types = [
                    "Gateway"
                  ]
                  section_name = "...my_section_name..."
                  tags = {
                    key = "value"
                  }
                  weight = 10
                }
              ]
            }
          }
        ]
        target_ref = {
          kind = "Mesh"
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
  type = "MeshTCPRoute"
}