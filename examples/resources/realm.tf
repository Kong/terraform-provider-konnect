resource "konnect_realm" "my_realm" {
  name                     = "Production"
  ttl                      = 60
  negative_ttl             = 10
  allow_all_control_planes = false
  allowed_control_planes   = ["a354cfdf-3483-4e22-9b8e-20938de0bfa6"]
  consumer_groups = [
    "prod-user"
  ]
  force_destroy = "true"
}
