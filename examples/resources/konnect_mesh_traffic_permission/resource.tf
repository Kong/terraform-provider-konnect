resource "konnect_mesh_traffic_permission" "my_meshtrafficpermission" {
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
          action = "Deny"
        }
        target_ref = {
          kind = "MeshService"
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
  type = "MeshTrafficPermission"
}