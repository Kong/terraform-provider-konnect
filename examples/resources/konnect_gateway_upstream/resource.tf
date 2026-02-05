resource "konnect_gateway_upstream" "my_gatewayupstream" {
  algorithm = "sticky-sessions"
  client_certificate = {
    id = "...my_id..."
  }
  control_plane_id          = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  created_at                = 5
  hash_fallback             = "path"
  hash_fallback_header      = "...my_hash_fallback_header..."
  hash_fallback_query_arg   = "...my_hash_fallback_query_arg..."
  hash_fallback_uri_capture = "...my_hash_fallback_uri_capture..."
  hash_on                   = "ip"
  hash_on_cookie            = "...my_hash_on_cookie..."
  hash_on_cookie_path       = "...my_hash_on_cookie_path..."
  hash_on_header            = "...my_hash_on_header..."
  hash_on_query_arg         = "...my_hash_on_query_arg..."
  hash_on_uri_capture       = "...my_hash_on_uri_capture..."
  healthchecks = {
    active = {
      concurrency = 1225225904
      headers = {
        key = [
          # ...
        ]
      }
      healthy = {
        http_statuses = [
          179
        ]
        interval  = 55332.9
        successes = 196
      }
      http_path                = "...my_http_path..."
      https_sni                = "...my_https_sni..."
      https_verify_certificate = true
      timeout                  = 59064.97
      type                     = "http"
      unhealthy = {
        http_failures = 207
        http_statuses = [
          396
        ]
        interval     = 27179.84
        tcp_failures = 92
        timeouts     = 183
      }
    }
    passive = {
      healthy = {
        http_statuses = [
          135
        ]
        successes = 189
      }
      type = "tcp"
      unhealthy = {
        http_failures = 249
        http_statuses = [
          865
        ]
        tcp_failures = 234
        timeouts     = 21
      }
    }
    threshold = 36.29
  }
  host_header                 = "...my_host_header..."
  id                          = "...my_id..."
  name                        = "...my_name..."
  slots                       = 4503
  sticky_sessions_cookie      = "...my_sticky_sessions_cookie..."
  sticky_sessions_cookie_path = "...my_sticky_sessions_cookie_path..."
  tags = [
    "..."
  ]
  updated_at   = 1
  use_srv_name = true
}