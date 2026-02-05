resource "konnect_realm" "my_realm" {
  allow_all_control_planes = true
  allowed_control_planes = [
    "..."
  ]
  consumer_groups = [
    "..."
  ]
  force_destroy = "false"
  name          = "...my_name..."
  negative_ttl  = 11
  ttl           = 15
}