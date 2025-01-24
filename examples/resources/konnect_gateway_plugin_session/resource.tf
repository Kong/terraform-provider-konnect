resource "konnect_gateway_plugin_session" "my_gatewaypluginsession" {
  config = {
    absolute_timeout = 1.93
    audience         = "...my_audience..."
    cookie_domain    = "...my_cookie_domain..."
    cookie_http_only = true
    cookie_name      = "...my_cookie_name..."
    cookie_path      = "...my_cookie_path..."
    cookie_same_site = "Default"
    cookie_secure    = true
    idling_timeout   = 6.27
    logout_methods = [
      "POST"
    ]
    logout_post_arg           = "...my_logout_post_arg..."
    logout_query_arg          = "...my_logout_query_arg..."
    read_body_for_logout      = true
    remember                  = false
    remember_absolute_timeout = 7.47
    remember_cookie_name      = "...my_remember_cookie_name..."
    remember_rolling_timeout  = 6.77
    request_headers = [
      "timeout"
    ]
    response_headers = [
      "audience"
    ]
    rolling_timeout = 9.29
    secret          = "...my_secret..."
    stale_ttl       = 3.88
    storage         = "cookie"
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
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
    "wss"
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