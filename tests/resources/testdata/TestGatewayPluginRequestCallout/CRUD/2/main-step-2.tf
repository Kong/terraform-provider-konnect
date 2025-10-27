resource "konnect_gateway_control_plane" "plugincp" {
  name         = "Terraform Control Plane For Request Callout Plugin"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
}


resource "konnect_gateway_plugin_request_callout" "my_request_callout" {
  enabled = true

  config = {
    callouts = [
      {
        cache = {
          bypass = true
        }
        name = "c1"

        request = {
          url = "http://example.com/uuid"
          method = "GET"
          headers = {
            custom = {
              key = jsonencode("value")
            }
            forward = false
          }
          body = {
            custom = {
              key = jsonencode("value")
            }
            decode  = true
            forward = false
          }
          query = {
            custom = {
              key = jsonencode("value")
            }
            forward = true
          }
          error = {
            error_response_code = 500
            error_response_msg  = "server error"
            http_statuses = [
              400,
              500
            ]
            on_error = "fail"
            retries  = 1
          }
          http_opts = {
            ssl_verify      = false
            timeouts = {
              connect = 30000
              read    = 3000
              write   = 3000
            }
          }
        }

        response = {
          headers = {
            store = false
          }
          body = {
            decode = true
          }
        }


      }, 
    ]
    upstream = {
    }
  }

  control_plane_id = konnect_gateway_control_plane.plugincp.id
}
