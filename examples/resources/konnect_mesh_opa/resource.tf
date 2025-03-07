resource "konnect_mesh_opa" "my_meshopa" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      agent_config = {
        inline        = "...my_inline..."
        inline_string = "...my_inline_string..."
        secret        = "...my_secret..."
      }
      append_policies = [
        {
          ignore_decision = true
          rego = {
            inline        = "...my_inline..."
            inline_string = "...my_inline_string..."
            secret        = "...my_secret..."
          }
        }
      ]
      auth_config = {
        on_agent_failure = "Deny"
        request_body = {
          max_size      = 6
          send_raw_body = true
        }
        status_on_error = 5
        timeout         = "...my_timeout..."
      }
    }
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
  }
  type = "MeshOPA"
}