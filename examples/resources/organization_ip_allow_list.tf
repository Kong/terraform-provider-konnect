resource "konnect_organization_ip_allow_list" "my_organizationipallowlist" {
  allowed_ips = [
    "192.168.1.1",
    "192.168.1.0/22",
  ]
  enabled         = true
  organization_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
}
