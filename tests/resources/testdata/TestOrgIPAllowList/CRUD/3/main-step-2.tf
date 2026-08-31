resource "konnect_organization_ip_allow_list" "my_organizationipallowlist" {
  allowed_ips     = ["192.0.2.1", "192.0.2.2"]
  enabled         = false
  organization_id = "72ab1dcd-3e4a-4f9b-9e9d-94217898055a"
}
