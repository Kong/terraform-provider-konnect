resource "konnect_mesh_tls" "my_meshtls" {
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
          mode = "Permissive"
          tls_ciphers = [
            "ECDHE-RSA-CHACHA20-POLY1305"
          ]
          tls_version = {
            max = "TLS11"
            min = "TLSAuto"
          }
        }
        target_ref = {
          kind = "MeshService"
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
          mode = "Strict"
          tls_ciphers = [
            "ECDHE-RSA-AES256-GCM-SHA384"
          ]
          tls_version = {
            max = "TLSAuto"
            min = "TLS10"
          }
        }
      }
    ]
    target_ref = {
      kind = "MeshService"
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
  type = "MeshTLS"
}