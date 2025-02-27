resource "konnect_mesh_proxy_patch" "my_meshproxypatch" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      append_modifications = [
        {
          cluster = {
            json_patches = [
              {
                from  = "...my_from..."
                op    = "remove"
                path  = "...my_path..."
                value = "{ \"see\": \"documentation\" }"
              }
            ]
            match = {
              name   = "...my_name..."
              origin = "...my_origin..."
            }
            operation = "Patch"
            value     = "...my_value..."
          }
          http_filter = {
            json_patches = [
              {
                from  = "...my_from..."
                op    = "move"
                path  = "...my_path..."
                value = "{ \"see\": \"documentation\" }"
              }
            ]
            match = {
              listener_name = "...my_listener_name..."
              listener_tags = {
                key = "value",
              }
              name   = "...my_name..."
              origin = "...my_origin..."
            }
            operation = "AddFirst"
            value     = "...my_value..."
          }
          listener = {
            json_patches = [
              {
                from  = "...my_from..."
                op    = "add"
                path  = "...my_path..."
                value = "{ \"see\": \"documentation\" }"
              }
            ]
            match = {
              name   = "...my_name..."
              origin = "...my_origin..."
              tags = {
                key = "value",
              }
            }
            operation = "Add"
            value     = "...my_value..."
          }
          network_filter = {
            json_patches = [
              {
                from  = "...my_from..."
                op    = "replace"
                path  = "...my_path..."
                value = "{ \"see\": \"documentation\" }"
              }
            ]
            match = {
              listener_name = "...my_listener_name..."
              listener_tags = {
                key = "value",
              }
              name   = "...my_name..."
              origin = "...my_origin..."
            }
            operation = "AddBefore"
            value     = "...my_value..."
          }
          virtual_host = {
            json_patches = [
              {
                from  = "...my_from..."
                op    = "move"
                path  = "...my_path..."
                value = "{ \"see\": \"documentation\" }"
              }
            ]
            match = {
              name                     = "...my_name..."
              origin                   = "...my_origin..."
              route_configuration_name = "...my_route_configuration_name..."
            }
            operation = "Remove"
            value     = "...my_value..."
          }
        }
      ]
    }
    target_ref = {
      kind = "MeshExternalService"
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
  type = "MeshProxyPatch"
}