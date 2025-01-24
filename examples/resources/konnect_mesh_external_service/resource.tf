resource "konnect_mesh_external_service" "my_meshexternalservice" {
  cp_id         = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  creation_time = "0001-01-01T00:00:00Z"
  labels = {
    key = "value",
  }
  mesh              = "...my_mesh..."
  modification_time = "0001-01-01T00:00:00Z"
  name              = "...my_name..."
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
  status = {
    addresses = [
      {
        hostname = "...my_hostname..."
        hostname_generator_ref = {
          core_name = "...my_core_name..."
        }
        origin = "...my_origin..."
      }
    ]
    hostname_generators = [
      {
        conditions = [
          {
            message = "...my_message..."
            reason  = "...my_reason..."
            status  = "True"
            type    = "...my_type..."
          }
        ]
        hostname_generator_ref = {
          core_name = "...my_core_name..."
        }
      }
    ]
    vip = {
      ip = "...my_ip..."
    }
  }
  type = "MeshExternalService"
}