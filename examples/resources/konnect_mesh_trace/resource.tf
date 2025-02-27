resource "konnect_mesh_trace" "my_meshtrace" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value",
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      backends = [
        {
          datadog = {
            split_service = false
            url           = "...my_url..."
          }
          open_telemetry = {
            endpoint = "otel-collector:4317"
          }
          type = "OpenTelemetry"
          zipkin = {
            api_version         = "httpProto"
            shared_span_context = false
            trace_id128bit      = false
            url                 = "...my_url..."
          }
        }
      ]
      sampling = {
        client = {
          integer = 2
          str     = "...my_str..."
        }
        overall = {
          integer = 4
          str     = "...my_str..."
        }
        random = {
          integer = 4
          str     = "...my_str..."
        }
      }
      tags = [
        {
          header = {
            default = "...my_default..."
            name    = "...my_name..."
          }
          literal = "...my_literal..."
          name    = "...my_name..."
        }
      ]
    }
    target_ref = {
      kind = "Dataplane"
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
  type = "MeshTrace"
}