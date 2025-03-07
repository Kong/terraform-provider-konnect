resource "konnect_gateway_upstream" "my_gatewayupstream" {
  algorithm = "latency"
  client_certificate = {
    id = "...my_id..."
  }
  control_plane_id          = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  hash_fallback             = "path"
  hash_fallback_header      = "...my_hash_fallback_header..."
  hash_fallback_query_arg   = "...my_hash_fallback_query_arg..."
  hash_fallback_uri_capture = "...my_hash_fallback_uri_capture..."
  hash_on                   = "header"
  hash_on_cookie            = "...my_hash_on_cookie..."
  hash_on_cookie_path       = "...my_hash_on_cookie_path..."
  hash_on_header            = "...my_hash_on_header..."
  hash_on_query_arg         = "...my_hash_on_query_arg..."
  hash_on_uri_capture       = "...my_hash_on_uri_capture..."
  healthchecks = {
    active = {
      concurrency = 6
      headers = {
        key = "value"
      }
      healthy = {
        http_statuses = [
          0
        ]
        interval  = 8.45
        successes = 8
      }
      http_path                = "...my_http_path..."
      https_sni                = "...my_https_sni..."
      https_verify_certificate = true
      timeout                  = 9.02
      type                     = "https"
      unhealthy = {
        http_failures = 8
        http_statuses = [
          3
        ]
        interval     = 4.15
        tcp_failures = 3
        timeouts     = 7
      }
    }
    passive = {
      healthy = {
        http_statuses = [
          0
        ]
        successes = 8
      }
      type = "grpcs"
      unhealthy = {
        http_failures = 10
        http_statuses = [
          9
        ]
        tcp_failures = 10
        timeouts     = 0
      }
    }
    threshold = 3.63
  }
  host_header = "...my_host_header..."
  id          = "...my_id..."
  name        = "...my_name..."
  slots       = 0
  tags = [
    "..."
  ]
  use_srv_name = true
}