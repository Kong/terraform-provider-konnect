# Check that portal list works
data "konnect_portal_list" "my_portallist" {
  page_number = 1
  page_size   = 1
}

output "portal_list" {
  value = data.konnect_portal_list.my_portallist
}
