# WARNING: this resource is risky. When enabled, only the listed IPs/CIDRs can
# reach your organization's Konnect Admin API, including this provider. Include
# the address of every machine and CI runner that applies Terraform, or you can
# lock yourself out. Destroying the resource disables the allow list but does
# not remove the entries.
resource "konnect_organization_ip_allow_list" "my_organizationipallowlist" {
  allowed_ips = [
    "192.0.2.1",
    "198.51.100.0/24",
  ]
  enabled         = false # set to true only once the list is correct
  organization_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
}
