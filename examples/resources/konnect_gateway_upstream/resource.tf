resource "konnect_gateway_upstream" "my_gatewayupstream" {
  algorithm                 = "consistent-hashing"
  control_plane_id          = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  hash_fallback             = "query_arg"
  hash_fallback_header      = "...my_hash_fallback_header..."
  hash_fallback_query_arg   = "...my_hash_fallback_query_arg..."
  hash_fallback_uri_capture = "...my_hash_fallback_uri_capture..."
  hash_on                   = "cookie"
  hash_on_cookie            = "...my_hash_on_cookie..."
  hash_on_cookie_path       = "...my_hash_on_cookie_path..."
  hash_on_header            = "...my_hash_on_header..."
  hash_on_query_arg         = "...my_hash_on_query_arg..."
  hash_on_uri_capture       = "...my_hash_on_uri_capture..."
  host_header               = "...my_host_header..."
  name                      = "Daisy Daniel"
  slots                     = 2
  upstream_id               = "426d620c-7058-4ae6-aacc-f85a3204a2c5"
  use_srv_name              = true
}