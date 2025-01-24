resource "konnect_gateway_plugin_proxy_cache" "my_gatewaypluginproxycache" {
  config = {
    cache_control = false
    cache_ttl     = 8
    content_type = [
      "..."
    ]
    ignore_uri_case = false
    memory = {
      dictionary_name = "...my_dictionary_name..."
    }
    request_method = [
      "POST"
    ]
    response_code = [
      411
    ]
    response_headers = {
      age            = false
      x_cache_key    = true
      x_cache_status = false
    }
    storage_ttl = 2
    strategy    = "memory"
    vary_headers = [
      "..."
    ]
    vary_query_params = [
      "..."
    ]
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
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
}