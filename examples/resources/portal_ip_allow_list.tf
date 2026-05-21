resource "konnect_portal" "my_portal1" {
  authentication_enabled               = true
  auto_approve_applications            = false
  auto_approve_developers              = false
  default_api_visibility               = "public"
  default_page_visibility              = "private"
  force_destroy                        = "false"
  labels = {
    key = "value"
  }
  name         = "my_portal1"
  rbac_enabled = false
}

resource "konnect_portal_ip_allow_list" "my_portalipallowlist" {
  allowed_ips = [
    "192.168.1.1",
    "192.168.1.0/22",
    "192.168.1.2"
  ]
  portal_id = konnect_portal.my_portal1.id
}
