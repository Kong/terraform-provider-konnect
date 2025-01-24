resource "konnect_gateway_plugin_canary" "my_gatewayplugincanary" {
  config = {
    canary_by_header_name = "...my_canary_by_header_name..."
    duration              = 0.88
    groups = [
      "..."
    ]
    hash              = "deny"
    hash_header       = "...my_hash_header..."
    percentage        = 35.35
    start             = 7.39
    steps             = 5.98
    upstream_fallback = false
    upstream_host     = "...my_upstream_host..."
    upstream_port     = 15742
    upstream_uri      = "...my_upstream_uri..."
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
    "grpcs"
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