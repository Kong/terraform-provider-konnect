resource "konnect_mesh_external_service" "my_meshexternalservice" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    endpoints = [
      {
        address = "example.com"
        port    = 9478
      }
    ]
    extension = {
      config = "{ \"see\": \"documentation\" }"
      type   = "...my_type..."
    }
    match = {
      port     = 1444
      protocol = "tcp"
      type     = "HostnameGenerator"
    }
    tls = {
      allow_renegotiation = false
      enabled             = true
      verification = {
        ca_cert = {
          inline        = "...my_inline..."
          inline_string = "...my_inline_string..."
          secret        = "...my_secret..."
        }
        client_cert = {
          inline        = "...my_inline..."
          inline_string = "...my_inline_string..."
          secret        = "...my_secret..."
        }
        client_key = {
          inline        = "...my_inline..."
          inline_string = "...my_inline_string..."
          secret        = "...my_secret..."
        }
        mode        = "SkipAll"
        server_name = "...my_server_name..."
        subject_alt_names = [
          {
            type  = "Exact"
            value = "...my_value..."
          }
        ]
      }
      version = {
        max = "TLS12"
        min = "TLS10"
      }
    }
  }
  type = "MeshExternalService"
}