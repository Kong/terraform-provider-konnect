resource "konnect_gateway_plugin_azure_functions" "my_gatewaypluginazurefunctions" {
  config = {
    apikey       = "...my_apikey..."
    appname      = "...my_appname..."
    clientid     = "...my_clientid..."
    functionname = "...my_functionname..."
    hostdomain   = "...my_hostdomain..."
    https        = false
    https_verify = false
    keepalive    = 4.24
    routeprefix  = "...my_routeprefix..."
    timeout      = 0.71
  }
  consumer = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 5
  enabled          = true
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  partials = [
    {
      id   = "...my_id..."
      name = "...my_name..."
      path = "...my_path..."
    }
  ]
  protocols = [
    "http"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
  updated_at = 2
}