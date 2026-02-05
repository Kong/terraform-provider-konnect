resource "konnect_gateway_plugin_request_transformer" "my_gatewaypluginrequesttransformer" {
  config = {
    add = {
      body = [
        "..."
      ]
      headers = [
        "..."
      ]
      querystring = [
        "..."
      ]
    }
    append = {
      body = [
        "..."
      ]
      headers = [
        "..."
      ]
      querystring = [
        "..."
      ]
    }
    http_method = "...my_http_method..."
    remove = {
      body = [
        "..."
      ]
      headers = [
        "..."
      ]
      querystring = [
        "..."
      ]
    }
    rename = {
      body = [
        "..."
      ]
      headers = [
        "..."
      ]
      querystring = [
        "..."
      ]
    }
    replace = {
      body = [
        "..."
      ]
      headers = [
        "..."
      ]
      querystring = [
        "..."
      ]
      uri = "...my_uri..."
    }
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at       = 7
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
    "https"
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
  updated_at = 6
}